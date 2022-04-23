package commands

const (
	EXF_FILE uint8 = iota
	INF_FILE
	READ_DIR
	SUICIDE
)

const (
	GOOD int = iota
	OPEN_ERROR
	READ_ERROR
	WRITE_ERROR
	STAT_ERROR
)

type C2Req struct {
	ReqLen uint64
	Type   uint8
}

type C2Res struct {
	ResLen uint64
	Type   uint8
}

type ExfReq struct {
	Header  C2Req
	PathLen uint16
	Path    []byte
}

type ExfRes struct {
	Header     C2Res
	Status     uint8
	ContentLen uint64
	Content    []byte
}

type InfReq struct {
	Header     C2Req
	Perm       uint16
	PathLen    uint16
	ContentLen uint64
	Data       []byte
}

type InfRes struct {
	Header C2Res
	Status uint8
}

type DirReq struct {
	Header  C2Res
	PathLen uint16
	Path    []byte
}

type DirRes struct {
	Header    C2Res
	Status    uint8
	LinkCount int32
}
