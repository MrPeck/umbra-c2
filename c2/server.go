package c2

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"time"
)

var Clients map[string]net.Conn

type C2Config struct {
	Host string
	Port string
}

func ConnIsClosed(c net.Conn) bool {
	r := bufio.NewReader(c)
	_, err := r.Peek(1)
	return err == io.EOF
}

func Run(c *C2Config) error {
	server, err := net.Listen("tcp", fmt.Sprintf("%s:%s", c.Host, c.Port))

	if err != nil {
		return err
	}

	defer server.Close()

	defer func() {
		for name, conn := range Clients {
			err = conn.Close()
			if err != nil {
				fmt.Printf("failed to close connection to %s: %v\n", name, err)
			}
		}
	}()

	cleanupClientsTicker := time.NewTicker(30 * time.Second)
	go func() {
		for {
			<-cleanupClientsTicker.C
			fmt.Println("client cleanup")
			for name, conn := range Clients {
				if ConnIsClosed(conn) {
					err := conn.Close()

					if err != nil {
						fmt.Println(err)
					}

					delete(Clients, name)
					fmt.Println("cleaned client", name)
				}
			}
		}
	}()

	for {
		client, err := server.Accept()

		if err != nil {
			return err
		}

		Clients[client.RemoteAddr().String()] = client
	}
}
