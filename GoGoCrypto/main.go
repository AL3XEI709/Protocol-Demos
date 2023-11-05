package main

import (

	"GoCode/Controller"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("hello world")
	REngin := gin.Default()
	REngin.Static("/assets", "./assets")
	REngin.LoadHTMLGlob("views/*")
	registRouter(REngin)
	REngin.Run(":7842")
}

func registRouter(r *gin.Engine) {

	new(Controller.PubController).Router(r)

}
