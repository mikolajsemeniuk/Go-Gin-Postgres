package application

import (
	"go-gin-postgres/program/controllers/post"

	_ "go-gin-postgres/program/docs" // swag init generates this folder

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Route() *Application {
	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json") // The url pointing to API definition
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	router.GET("/todo", post.GetPosts)
	router.GET("/todo/:id", post.GetPost)
	router.POST("/todo", post.AddPost)
	router.PUT("/todo/:id", post.UpdatePost)
	router.DELETE("/todo/:id", post.RemovePost)
	return &Application{}
}
