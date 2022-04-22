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

	router.GET("/hosts", routes.GetHosts)

	host := router.Group("/host")
	{
		specificHost := host.Group("/:id")
		{
			specificHost.GET("/", routes.GetHost)
			specificHost.GET("/file", routes.GetHostFile)
		}
	}

	router.Run(fmt.Sprintf("%s:%s", c.Host, c.Port))
}
