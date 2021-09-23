package post

import (
	postService "go-gin-postgres/program/services/post"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// HealthCheck godoc
// @Summary /post/:id
// @Description description over here
// @Tags Post
// @Accept */*
// @Produce json
// @Param id path int true "Post Id"
// @Success 200 {object} post.Post
// @Failure 400 {object} interface{}
// @Failure 404 {object} interface{}
// @Router /posts/{id} [get]
func GetPost(context *gin.Context) {
	id, error := strconv.ParseInt(context.Param("id"), 10, 64)
	if error != nil {
		context.JSON(http.StatusBadRequest, "cannot convert id to int64")
		return
	}

	post, error := postService.GetPost(id)
	if error != nil {
		context.JSON(http.StatusNotFound, error.Error())
		return
	}
	context.JSON(http.StatusOK, post)
}
