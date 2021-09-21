package main

import (
	"github.com/symmetric-project/node-backend/env"
	"github.com/symmetric-project/node-backend/handlers"

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

	GIN.Use(handlers.CORS())
	GIN.POST("/", handlers.GraphQLHandler())
	GIN.GET("/", handlers.GraphQLPlaygroundHandler())
	GIN.Run(":4000")
}
