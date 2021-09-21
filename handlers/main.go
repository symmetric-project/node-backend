package handlers

import (
	"context"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/symmetric-project/node-backend/graph"
	"github.com/symmetric-project/node-backend/graph/generated"
	"github.com/symmetric-project/node-backend/middleware"
)

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func GraphQLHandler() gin.HandlerFunc {
	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	return func(c *gin.Context) {
		resolverContext := middleware.ResolverContext{}

		// Add http.Request for the purpose of setting cookies in gqlgen resolvers
		resolverContext.Writer = &c.Writer

		// Add jwt if it is present in the request
		jwt, err := c.Cookie("jwt")
		if err == nil {
			resolverContext.JWT = &jwt
		}

		ctx := context.WithValue(c.Request.Context(), "resolverContext", resolverContext)

		// Serve the request with the new context
		req := c.Request.WithContext(ctx)
		h.ServeHTTP(c.Writer, req)
	}
}

func GraphQLPlaygroundHandler() gin.HandlerFunc {
	h := playground.Handler("graphql", "/")
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
