package main

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/symmetric-project/node-backend/graph"
	"github.com/symmetric-project/node-backend/graph/generated"
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
