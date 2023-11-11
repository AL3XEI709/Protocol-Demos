package Controller

import (
	"Sharing/mycrt" 
	"fmt" 
	"github.com/gin-gonic/gin" 
	"net/http" 
)

type PubController struct {
}

func SendJsonBack(feedback string, check string, enc string, c *gin.Context) {
	messageMap := map[string]interface{}{
		"msg":   feedback,
		"check": check,
		"enc":   enc,
	}
	c.JSON(http.StatusOK, messageMap)
}

func (pc PubController) Router(pr *gin.Engine) {
	pr.GET("/", pc.Room)
	pr.POST("/api/sec", pc.SecShr)
}

func (pc PubController) Room(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{"ping": "pong"}) 
}

