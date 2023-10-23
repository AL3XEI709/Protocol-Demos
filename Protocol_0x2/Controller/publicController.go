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

func (pc PubController) SecShr(c *gin.Context) {
	var output string 

	secret_err := c.PostForm("secret") 
	sec, SecFeedback, SecCheck := Mycrt.CheckAtoi(secret_err)
	fmt.Println("Secret:", sec) 
	share_err := c.PostForm("share") 
	shr, ShrFeedback, ShrCheck := Mycrt.CheckAtoi(share_err)
	fmt.Println("Share up limit:", shr)  
	
	

}