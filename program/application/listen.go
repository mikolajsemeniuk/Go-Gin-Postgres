package application

import "github.com/gin-gonic/gin"

var (
	router = gin.Default()
)

func (*Application) Listen() {
	router.Run()
}
