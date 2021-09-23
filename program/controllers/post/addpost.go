package post

import (
	"go-gin-postgres/program/inputs"
	"go-gin-postgres/program/payloads"
	postService "go-gin-postgres/program/services/post"
	"net/http"

	"github.com/gin-gonic/gin"
	"gopkg.in/validator.v2"
)

// HealthCheck godoc
// @Summary /post
// @Description description over here
// @Tags Post
// @Accept */*
// @Produce json
// @Param body body inputs.PostInput true "Add Post"
// @Success 202 {object} interface{}
// @Success 404 {object} interface{}
// @Router /posts [post]
func AddPost(context *gin.Context) {
	var postInput inputs.PostInput

	if error := context.ShouldBindJSON(&postInput); error != nil {
		context.JSON(http.StatusBadRequest, payloads.ErrorPayload{
			Message: "error binding data",
		})
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
