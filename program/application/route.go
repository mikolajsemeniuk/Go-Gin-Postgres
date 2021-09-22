package application

import "go-gin-postgres/program/controllers/post"

func Route() *Application {
	router.GET("/todo", post.GetPosts)
	router.GET("/todo/:id", post.GetPost)
	router.POST("/todo", post.AddPost)
	return &Application{}
}
