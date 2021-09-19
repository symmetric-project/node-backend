package main

import (
	"github.com/gin-contrib/cors"
	"github.com/symmetric-project/node-backend/env"

	"github.com/gin-gonic/gin"
)

var MODE string
var DOMAIN_DEV string
var DOMAIN_PROD string
var GIN *gin.Engine

func main() {
	GIN = gin.Default()

	allowedOrigins := make([]string, 0)
	if env.CONFIG.MODE == "dev" {
		allowedOrigins = append(allowedOrigins, "http://"+env.CONFIG.DOMAIN)
	} else {
		allowedOrigins = append(allowedOrigins, "https://"+env.CONFIG.DOMAIN)
	}

	GIN.Use(cors.Default())

	GIN.Use(authMiddlewareHandler)
	GIN.POST("/", graphqlHandler())
	GIN.GET("/", graphqlPlaygroundHandler())
	GIN.Run(":4000")
}
