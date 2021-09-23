package main

import (
	"github.com/symmetric-project/node-backend/handlers"

	"github.com/gin-gonic/gin"
)

var MODE string
var DOMAIN_DEV string
var DOMAIN_PROD string
var GIN *gin.Engine

func main() {
	GIN = gin.Default()
	GIN.Use(handlers.CORS())
	GIN.POST("/", handlers.GraphQLHandler())
	GIN.GET("/", handlers.GraphQLPlaygroundHandler())
	GIN.Run(":4000")
}
