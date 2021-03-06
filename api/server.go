package api

import (
	"fmt"
	"umbra-c2/api/routes"

	"github.com/gin-gonic/gin"
)

type APIConfig struct {
	Host string
	Port string
}

func Run(c *APIConfig) {
	router := gin.Default()

	router.GET("/hosts", routes.Hosts)

	host := router.Group("/host")
	{
		specificHost := host.Group("/:id")
		{
			specificHost.GET("/", routes.Host)
			specificHost.GET("/file", routes.HostFile)
		}
	}

	router.Run(fmt.Sprintf("%s:%s", c.Host, c.Port))
}
