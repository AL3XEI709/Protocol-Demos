package main

import (

	"Sharing/Controller"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("hello world")
	REngin := gin.Default()
	REngin.LoadHTMLGlob("views/*")
	registRouter(REngin)
	REngin.Run(":9999")
}

func registRouter(r *gin.Engine) {

	new(Controller.PubController).Router(r)

}
