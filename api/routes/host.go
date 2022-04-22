package routes

import (
	"net/http"
	"umbra-c2/c2"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Host struct {
	Name string
	Id   uuid.UUID
}

func newHost(c *c2.C2Client) Host {
	return Host{
		Name: c.Conn.RemoteAddr().String(),
		Id:   c.Id,
	}
}

func GetHosts(c *gin.Context) {
	hosts := make([]Host, 0, len(c2.Clients))
	for _, client := range c2.Clients {
		hosts = append(hosts, newHost(client))
	}

	c.JSON(http.StatusOK, hosts)
}

func GetHost(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))

	if err != nil {
		c.Status(http.StatusNotFound)
		return
	}

	if client, ok := c2.Clients[id.String()]; ok {
		c.JSON(http.StatusOK, newHost(client))
	} else {
		c.Status(http.StatusNotFound)
	}
}

func GetHostFile(c *gin.Context) {
	c.Status(http.StatusOK)
}
