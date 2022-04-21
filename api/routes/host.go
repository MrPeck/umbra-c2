package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Hosts(c *gin.Context) {
	c.Status(http.StatusOK)
}

func Host(c *gin.Context) {
	c.Status(http.StatusOK)
}

func HostFile(c *gin.Context) {
	c.Status(http.StatusOK)
}
