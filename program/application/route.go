package application

import "go-gin-postgres/program/controllers/post"

func Route() *Application {
	router.POST("/todo", post.Add)
	return &Application{}
}
