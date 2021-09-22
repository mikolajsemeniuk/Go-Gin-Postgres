package post

import (
	postService "go-gin-postgres/program/services/post"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddPost(context *gin.Context) {
	if error := postService.AddPost(); error != nil {
		context.JSON(http.StatusBadRequest, error.Error())
		return
	}
	context.JSON(http.StatusAccepted, "post inserted")
}
