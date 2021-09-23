package post

import (
	postService "go-gin-postgres/program/services/post"
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthCheck godoc
// @Summary /post
// @Description description over here
// @Tags Post
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /posts [get]
func GetPosts(context *gin.Context) {
	posts, error := postService.GetPosts()
	if error != nil {
		context.JSON(http.StatusBadRequest, error.Error())
		return
	}
	context.JSON(http.StatusOK, posts)
}
