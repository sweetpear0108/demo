package route

import (
	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	router := gin.Default()
	router.GET("/user/:id", GetUser)
	return router
}
