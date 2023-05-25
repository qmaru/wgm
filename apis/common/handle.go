package common

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// JSONHandler 返回 JSON 数据
func JSONHandler(c *gin.Context, status int, message string, data any) {
	c.JSON(http.StatusOK, gin.H{
		"status":  status,
		"message": message,
		"data":    data,
	})
}

func IDtoInt(id string) (int, error) {
	idN, err := strconv.Atoi(id)
	if err != nil {
		return 0, err
	}
	return idN, nil
}
