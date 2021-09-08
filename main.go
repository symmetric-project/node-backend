package main

import (
	"os"

	"github.com/gin-contrib/cors"
	"github.com/joho/godotenv"
	"github.com/symmetric-project/node-backend/graph"
	"github.com/symmetric-project/node-backend/graph/generated"

	"github.com/gin-gonic/gin"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

func graphqlHandler() gin.HandlerFunc {
	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func graphqlPlaygroundHandler() gin.HandlerFunc {
	h := playground.Handler("graphql", "/")
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

var MODE string
var DOMAIN_DEV string
var DOMAIN_PROD string
var G *gin.Engine

func init() {
	godotenv.Load()
	MODE = os.Getenv("MODE")
	DOMAIN_DEV = os.Getenv("DOMAIN_DEV")
	DOMAIN_PROD = os.Getenv("DOMAIN_PROD")
}

func main() {
	G = gin.Default()
	allowedOrigins := make([]string, 0)
	if MODE == "prod" {
		allowedOrigins = append(allowedOrigins, "https://"+DOMAIN_PROD)
	} else {
		allowedOrigins = append(allowedOrigins, "http://"+DOMAIN_DEV)
	}
	G.Use(cors.New(cors.Config{
		AllowOrigins:     allowedOrigins,
		AllowCredentials: true,
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Content-Type"},
	}))
	G.POST("/", graphqlHandler())
	G.GET("/", graphqlPlaygroundHandler())
	G.Run(":4000")
}
