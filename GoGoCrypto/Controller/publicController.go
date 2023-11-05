package Controller

import (
	"GoCode/myfunc"
	"github.com/gin-gonic/gin" 
	"net/http"
	"encoding/base64"
	"fmt" 
)

type PubController struct {
}

var sid []byte 
var key []byte 
var iv []byte 
var nonce []byte 
var pwd string  

func SendJsonBack(feedback string, check string, c *gin.Context) {
	messageMap := map[string]interface{}{
		"msg":   feedback,
		"check": check,
	}
	c.JSON(http.StatusOK, messageMap)
}

func (pc PubController) Router(pr *gin.Engine) {
	pr.GET("/", pc.Room)
	pr.POST("/api/dec", pc.MyDec) 
	pr.POST("/api/check", pc.CheckPwd)
}

func (pc PubController) Room(c *gin.Context) {
	sid = Myfunc.GetRandBytes(16) 
	key = Myfunc.GetRandBytes(16) 
	iv = Myfunc.GetRandBytes(16) 
	nonce = Myfunc.GetRandBytes(16) 

	token, _ := Myfunc.AESEnc(sid, key, iv) 
	c.SetCookie("token", base64.StdEncoding.EncodeToString(token), 0, "/", "", true, true) 
	c.SetCookie("nonce",base64.StdEncoding.EncodeToString(nonce) , 0, "/", "", true, true) 
	fmt.Println("token:", base64.StdEncoding.EncodeToString(token))	
	c.HTML(http.StatusOK, "index.html", gin.H{"ping": "pong"}) 

}

func(pc PubController) MyDec(c *gin.Context) {
	var feedback string 
	var check string 


	Rec_err := c.PostForm("Rec") 
	fmt.Println("Rec:", Rec_err) 
	Rec, err := base64.StdEncoding.DecodeString(Rec_err) 
	if err != nil {
		check = "false" 
		feedback = "Try again." 
		SendJsonBack(feedback, check, c) 
		return 
	}

	Pt_err, err := Myfunc.AESDec(Rec, key, iv) 
	//
	if err != nil || base64.StdEncoding.EncodeToString(Pt_err) == base64.StdEncoding.EncodeToString(sid){
		check = "false" 
		feedback = "Don't try to fool me! Try again." 
		SendJsonBack(feedback, check, c) 
		return 
	}
	check = "true" 
	feedback = "Access Accepted."  
	Pt := base64.StdEncoding.EncodeToString(Pt_err) 
	fmt.Println("(DEBUG)Pt: ", Pt) 
	nonce_hash := Myfunc.Hash(nonce) 
	pwd = Pt+nonce_hash
	fmt.Println("pwd: ",pwd)
	SendJsonBack(feedback, check, c) 
	
}

func(pc PubController) CheckPwd(c *gin.Context) {
	var feedback string 
	var check string 

	Rec := c.PostForm("Password")
	if pwd == "" {
		check = "false" 
		feedback = "Give me your token first." 
		SendJsonBack(feedback, check, c) 
		return  
	}
	if Rec == pwd {
		check = "true" 
		feedback = "Your flag is: SYC{AL3XEI_FAKE_FLAG}" 
	} else {
		check =  "false" 
		feedback = "Oops! Please try again." 
		
	}
	SendJsonBack(feedback, check, c) 

}
