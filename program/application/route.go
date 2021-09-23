package application

import "go-gin-postgres/program/controllers/post"

func Route() *Application {
	router.GET("/todo", post.GetPosts)
	router.GET("/todo/:id", post.GetPost)
	router.POST("/todo", post.AddPost)
	router.PUT("/todo/:id", post.UpdatePost)
	router.DELETE("/todo/:id", post.RemovePost)
	return &Application{}
}
