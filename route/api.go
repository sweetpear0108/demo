package route

import (
	"awesomeProject/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetUser(c *gin.Context) {
	req_account_id := c.Param("id")
	id, err := strconv.Atoi(req_account_id)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"error":   true,
			"message": "the id must be integer",
		})
		return
	}
	user, err := service.QueryById(id)
	//service.A()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"error":   true,
			"message": "user doesn't exist",
		})
		return

	}
	c.IndentedJSON(http.StatusOK, user)

}
