package Controller

import (
	"Commitment/myhash" 
	"fmt" 
	"github.com/gin-gonic/gin" 
	"net/http"
	"math/big"
	"crypto/rand"
	"strconv"
)

type PubController struct {
}
var RandX int 
var KeyB string 
var HashC string 

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
	pr.POST("/api/nonce", pc.NceGen)
	pr.POST("/api/guess", pc.GssCheck)
}

func (pc PubController) Room(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{"ping": "pong"}) 

}

func (pc PubController) NceGen(c *gin.Context) {
	var feedback string
	var check string
	var output string 
	Nce_err := c.PostForm("Nonce")
	fmt.Println("Nonce:", Nce_err)
	Nce, err := strconv.Atoi(Nce_err) 
	if err != nil {
		feedback = "Give me an Integer."
		check = "false"
	} else {
		if Nce > (1<<33-1) {
			feedback = "Give me something smaller." 
			check = "false"
		} else {
			X_,_ := rand.Int(rand.Reader, big.NewInt(100))  
			RandX = int(X_.Int64()) 
			KeyB = Myhash.GenKeyB()
			HashC = Myhash.GenHashC(Nce, RandX, KeyB) 
			feedback = "Submit OK." 
			check = "true" 
			output = "Hash (C): "+ HashC
		}
	}
	

	SendJsonBack(feedback, check, output, c) 
}

func (pc PubController) GssCheck(c *gin.Context) {
	var feedback string
	var check string
	var output string 
	guess_err := c.PostForm("Guess") 
	if HashC == "" {
		feedback = "Give me your Nonce first." 
		check = "false" 
	} else {
		fmt.Println("guess:", guess_err) 
		guess, err := strconv.Atoi(guess_err) 
		if err != nil {
			feedback = "Give me an Integer."
			check = "false"
		} else {
			if guess > 100 {
				feedback = "Give me number between 0 and 100." 
				check = "false"
			} else {
				feedback = "Submit OK."
				check = "true"  
				if guess == RandX {
					output = "Congrats! The key (B) is: " + KeyB 
				} else {
					output = "Sorry, the answer is " + strconv.Itoa(RandX) + ", the key (B) is " + KeyB 
				}
			}
		}
	}
	SendJsonBack(feedback, check, output, c) 
}