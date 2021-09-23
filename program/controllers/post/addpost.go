package post

import (
	"go-gin-postgres/program/inputs"
	postService "go-gin-postgres/program/services/post"
	"net/http"

	"github.com/gin-gonic/gin"
	"gopkg.in/validator.v2"
)

func AddPost(context *gin.Context) {
	var postInput inputs.PostInput

	if error := context.ShouldBindJSON(&postInput); error != nil {
		context.JSON(http.StatusBadRequest, error.Error())
		return
	}

	if error := validator.Validate(postInput); error != nil {
		errors := error.(validator.ErrorMap)
		context.JSON(http.StatusBadRequest, errors)
		return
	}

	if error := postService.AddPost(postInput); error != nil {
		context.JSON(http.StatusBadRequest, error.Error())
		return
	}

	context.JSON(http.StatusAccepted, "post inserted")
}
