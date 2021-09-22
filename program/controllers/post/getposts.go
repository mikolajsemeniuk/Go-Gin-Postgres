package post

import (
	postService "go-gin-postgres/program/services/post"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetPosts(context *gin.Context) {
	posts, error := postService.GetPosts()
	if error != nil {
		context.JSON(http.StatusBadRequest, error.Error())
		return
	}
	context.JSON(http.StatusAccepted, posts)
}
