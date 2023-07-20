package api

import (
	"net"

	"github.com/gin-gonic/gin"
	"github.com/yosa12978/MyDocu/internal/config"
)

func Listen() {
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(gin.Logger())

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "hello world",
		})
	})

	listener, err := net.Listen("tcp", config.Conf.Api.Addr)
	if err != nil {
		panic(err)
	}

	router.RunListener(listener)
}
