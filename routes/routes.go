package routes

import (
	"example.com/blog-rest-api/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/posts", getPosts)
	server.GET("/posts/:id", getPost)

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/posts", createPost)
	authenticated.PUT("/posts/:id", updatePost)
	authenticated.DELETE("/posts/:id", deletePost)

	server.POST("/signup", signup)
	server.POST("/login", login)
}
