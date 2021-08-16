package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Response(c *gin.Context, result map[string]interface{}, resultInfo ResultInfo) {
	if result == nil {
		c.JSON(http.StatusOK, gin.H{
			"state":   resultInfo.State,
			"message": resultInfo.MessageInfo,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"state":   resultInfo.State,
		"result":  result,
		"message": resultInfo.MessageInfo,
	})
}
