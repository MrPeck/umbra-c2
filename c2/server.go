package c2

import (
	"fmt"
	"net"
)

type C2Config struct {
	Host string
	Port string
}

func Run(c *C2Config) error {
	server, err := net.Listen("tcp", fmt.Sprintf("%s:%s", c.Host, c.Port))

	if err != nil {
		return err
	}

	defer server.Close()

	for {
		client, err := server.Accept()

		if err != nil {
			return err
		}

		defer client.Close()
	}
}
