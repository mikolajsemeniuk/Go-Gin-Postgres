package post

import (
	postService "go-gin-postgres/program/services/post"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// HealthCheck godoc
// @Summary Show the status of server.
// @Description get the status of server.
// @Tags GetPost
// @Accept */*
// @Produce json
// @Param id path int true "Post ID"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Router /todo/:id [get]
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
	context.JSON(http.StatusAccepted, post)
}
