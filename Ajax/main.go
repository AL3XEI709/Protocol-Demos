package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type RequestData struct {
	Text string `json:"text"`
}

type ResponseData struct {
	Message string `json:"message"`
}

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("templates/*")

	// 设置一个路由来提供HTML页面
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	// 设置一个API端点来处理用户提交
	router.POST("/api/submit", func(c *gin.Context) {
		var requestData RequestData
		if err := c.ShouldBindJSON(&requestData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "无效的请求数据"})
			return
		}

		// 在这里，您可以使用 requestData.Text 进行处理，然后构建响应
		// 为示例，直接返回输入文本
		data := ResponseData{
			Message: "您输入的文本是: " + requestData.Text,
		}

		c.JSON(http.StatusOK, data)
	})

	router.Run(":8080")
}
