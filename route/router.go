package route

import (
	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	router := gin.Default()

	router.POST("/reg", Verify)
	router.POST("/userReg", RegisterRequest)
	router.GET("/user/:id", GetUser)
	return router
}
