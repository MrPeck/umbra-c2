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

func GetHosts(c *gin.Context) {
	hosts := make([]Host, 0, len(c2.Clients))
	for name, client := range c2.Clients {
		hosts = append(hosts, Host{
			Name: name,
			Id:   client.Id,
		})
	}

	c.JSON(http.StatusOK, hosts)
}

func GetHost(c *gin.Context) {
	c.Status(http.StatusOK)
}

func GetHostFile(c *gin.Context) {
	c.Status(http.StatusOK)
}
