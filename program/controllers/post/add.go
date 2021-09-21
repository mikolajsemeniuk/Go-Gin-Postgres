package post

import (
	"fmt"
	postService "go-gin-postgres/program/services/post"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Add(context *gin.Context) {
	if error := postService.Add(); error != nil {
		fmt.Println(error)
		context.JSON(http.StatusBadRequest, error.Error())
	}
	context.JSON(http.StatusOK, "ok")
}
