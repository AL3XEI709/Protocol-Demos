# Ajax相关

## Go

**用go的gin框架写一个简单程序 ，证明gin可以实现ajax请求**

当使用Gin框架来实现Ajax请求时，你需要创建一个简单的API端点，该端点将响应Ajax请求，并返回JSON数据。



```go
package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

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

// 设置一个API端点来处理Ajax请求
router.GET("/api/data", func(c *gin.Context) {
	// 模拟一些数据
	data := ResponseData{
		Message: "这是来自服务器的响应。",
	}

	c.JSON(http.StatusOK, data)
})

router.Run(":8080")
}
```
在这个示例中，我们创建了一个名为`/api/data`的API端点，用于处理Ajax请求。当客户端发出GET请求到`/api/data`时，服务器会返回一个JSON响应，其中包含一条消息。

## JS

接下来，创建一个HTML模板文件 `index.html`，用于呈现一个包含Ajax请求的页面：

```html
<!DOCTYPE html>
<html>
<head>
    <title>使用Ajax请求</title>
    <script src="https://code.jquery.com/jquery-3.6.0.min.js"></script>
</head>
<body>
    <h1>使用Ajax请求获取数据</h1>
    <div id="response"></div>
    <button id="fetchData">获取数据</button>
    <script>
        // 当页面加载完成后，附加一个点击事件处理程序到按钮上
        $(document).ready(function () {
            $("#fetchData").click(function () {
                $.ajax({
                    type: "GET",
                    url: "/api/data",
                    success: function (data) {
                        $("#response").html(data.message);
                    },
                    error: function (error) {
                        $("#response").html("出错了：" + error.statusText);
                    }
                });
            });
        });
    </script>
</body>
</html>
```



下面是 `index.html` 内部代码的解释：

1. `<!DOCTYPE html>`：这是HTML文档的文档类型声明，表示这是一个HTML5文档。

2. `<html>`：HTML文档的根元素。

3. `<head>`：这个部分包含了文档的元信息和引用的外部资源。

4. `<title>`：这个元素用于设置浏览器标签栏中显示的标题。

5. `<script>`：这个元素用于在HTML文档中嵌入JavaScript代码。在这个示例中，我们在页面头部引入了jQuery库，以便使用jQuery来处理Ajax请求。

6. `<body>`：这是HTML文档的主体部分，用于显示页面内容。

7. `<h1>`：这个元素定义了一个一级标题，通常用于显示页面的主要标题。

8. `<div id="response">`：这是一个`<div>`元素，带有一个唯一的ID属性`response`。我们将在此元素中显示来自Ajax请求的响应数据。

9. `<button id="fetchData">获取数据</button>`：这是一个按钮元素，带有一个唯一的ID属性`fetchData`。当用户点击此按钮时，将触发一个JavaScript事件，发起Ajax请求。

10. `<script>`：这是用于编写JavaScript代码的元素。

11. `$(document).ready(function () { ... });`：这是jQuery的文档就绪函数，它确保页面的DOM已经完全加载后再执行其中的JavaScript代码。

12. `$("#fetchData").click(function () { ... });`：这是jQuery代码，用于为ID为`fetchData`的按钮添加点击事件处理程序。当用户点击按钮时，下面的函数将被执行。

13. `$.ajax({ ... });`：这是jQuery的`$.ajax()`函数，用于执行Ajax请求。它接受一个包含Ajax请求参数的对象。

14. `type: "GET"`：指定请求的HTTP方法为GET，即获取数据。

15. `url: "/api/data"`：指定要请求数据的URL，即`/api/data`，这是服务器上的一个API端点。

16. `success: function (data) { ... }`：定义了当请求成功完成后的回调函数，这里的`data`参数包含了从服务器返回的数据。

17. `$("#response").html(data.message);`：这行代码将服务器返回的`data.message`（即消息文本）设置为ID为`response`的`<div>`元素的HTML内容，以便将响应数据显示在页面上。

18. `error: function (error) { ... }`：定义了当请求发生错误时的回调函数，这里的`error`参数包含了错误信息。

19. `$("#response").html("出错了：" + error.statusText);`：这行代码将错误信息显示在页面上，告诉用户发生了错误。

总之，`index.html`包含一个简单的HTML页面，其中包括一个按钮，当点击按钮时，通过Ajax请求获取数据，然后将响应数据显示在页面上。这使用户能够与服务器进行交互，而不必刷新整个页面。



# JQuery

jQuery（简写为 `$`）是一个流行的JavaScript库，它提供了许多方便的函数和方法，以简化JavaScript编程和处理DOM操作。在给定的上下文中，`$(document).ready(function () { ... })` 是 jQuery 的一种常见用法，它的含义是：

1. `$(document)`：这部分代码使用 jQuery 选择器 `$` 来选择文档对象（即网页文档）。这是一个特殊的 jQuery 选择器，用于选择整个文档对象，通常代表整个网页。

2. `.ready(function () { ... })`：这是一个事件处理函数，它被附加到文档对象上，以确保在文档完全加载并解析后执行其中的代码。在这个上下文中，函数内的代码将在页面文档准备好后执行。

具体含义是，这段代码告诉浏览器在整个文档准备好之后执行函数内的代码。这是一个非常有用的技巧，因为它确保您的 JavaScript 代码不会在DOM元素尚未准备好时执行，从而避免出现错误。

通常，您会在 `$(document).ready()` 函数内编写初始化代码、事件处理程序和其他与DOM相关的操作，以确保它们在页面加载后立即可用。