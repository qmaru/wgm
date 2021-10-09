package apis

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// DataHandler Response 数据结构
func DataHandler(c *gin.Context, status int, message string, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"status":  status,
		"message": message,
		"data":    data,
	})
}

func ParseID(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Panic(err)
	}
	return i
}
