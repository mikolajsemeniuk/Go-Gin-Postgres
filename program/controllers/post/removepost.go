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
// @Tags RemovePost
// @Accept */*
// @Produce json
// @Param id path int true "Post ID"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Router /todo [delete]
func RemovePost(context *gin.Context) {
	id, error := strconv.ParseInt(context.Param("id"), 10, 64)
	if error != nil {
		context.JSON(http.StatusBadRequest, "cannot convert id to int64")
		return
	}

	if error := postService.RemovePost(id); error != nil {
		context.JSON(http.StatusBadRequest, error.Error())
		return
	}

	context.JSON(http.StatusAccepted, "post removed")
}
