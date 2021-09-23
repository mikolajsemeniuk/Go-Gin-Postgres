package post

import (
	"go-gin-postgres/program/inputs"
	postService "go-gin-postgres/program/services/post"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gopkg.in/validator.v2"
)

// HealthCheck godoc
// @Summary Show the status of server.
// @Description get the status of server.
// @Tags UpdatePost
// @Accept */*
// @Produce json
// @Param id path int true "Post ID"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Router /todo/:id [put]
func UpdatePost(context *gin.Context) {
	id, error := strconv.ParseInt(context.Param("id"), 10, 64)
	if error != nil {
		context.JSON(http.StatusBadRequest, "cannot convert id to int64")
		return
	}

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

	if error := postService.UpdatePost(id, postInput); error != nil {
		context.JSON(http.StatusBadRequest, error.Error())
		return
	}

	context.JSON(http.StatusAccepted, "post updated")
}
