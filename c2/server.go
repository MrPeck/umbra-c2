package c2

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"time"

	"github.com/google/uuid"
)

type C2Client struct {
	Conn net.Conn
	Id   uuid.UUID
}

type C2Config struct {
	Host string
	Port string
}

var Clients map[string]*C2Client

func NewC2Client(c net.Conn) *C2Client {
	return &C2Client{
		Conn: c,
		Id:   uuid.New(),
	}
}

func connIsClosed(c net.Conn) bool {
	r := bufio.NewReader(c)
	_, err := r.Peek(1)
	return err == io.EOF
}

func clientCleanupCronJob(d time.Duration) {
	cleanupClientsTicker := time.NewTicker(d)
	go func() {
		for {
			<-cleanupClientsTicker.C
			fmt.Println("client cleanup")
			for name, client := range Clients {
				if connIsClosed(client.Conn) {
					err := client.Conn.Close()

					if err != nil {
						fmt.Println(err)
					}

					delete(Clients, name)
					fmt.Println("cleaned client", name)
				}
			}
		}
	}()
}

func Run(c *C2Config) error {
	server, err := net.Listen("tcp", fmt.Sprintf("%s:%s", c.Host, c.Port))

	if err != nil {
		return err
	}

	defer server.Close()

	defer func() {
		for name, client := range Clients {
			err = client.Conn.Close()
			if err != nil {
				fmt.Printf("failed to close connection to %s: %v\n", name, err)
			}
		}
	}()

	clientCleanupCronJob(30 * time.Second)

	for {
		conn, err := server.Accept()

		if err != nil {
			return err
		}

		Clients[conn.RemoteAddr().String()] = NewC2Client(conn)
	}
}
