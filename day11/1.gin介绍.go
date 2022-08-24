package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

//hello world 例子
/*
func main() {
	//1.创建路由引擎
	r := gin.Default()
	//2.绑定路由规则，执行的函数
	//gin.Context 封装了request和response
	r.GET("/", func(c *gin.Context) {
		//返回hello world!
		c.String(http.StatusOK, "hello world!")
	})
	//3.监听端口，默认在8080
	r.Run(":8000")
}
*/

// 获取api参数
/*
func main() {
	//1.创建路由引擎
	r := gin.Default()
	r.GET("/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		c.String(http.StatusOK, name+" is "+action)
	})
	//3.监听端口，默认在8080
	r.Run(":8000")
}

*/
// 获取url参数
/*
func main() {
	//1.创建路由引擎
	r := gin.Default()
	r.GET("/welcome", func(c *gin.Context) {
		name := c.DefaultQuery("name", "jack") //为空返回默认值
		//name := c.Query("name") //为空返回空字符串
		c.String(http.StatusOK, " hello "+name)
		//http://localhost:8000/welcome?name=lucy  -> hello lucy

	})
	//3.监听端口，默认在8080
	r.Run(":8000")
}
*/
// 获取表单参数
func main() {
	//1.创建路由引擎
	r := gin.Default()
	r.POST("/form", func(c *gin.Context) {
		type1 := c.DefaultPostForm("type1", "alert") //为空返回默认值

		username := c.PostForm("username")
		password := c.PostForm("password")
		hobby := c.PostFormArray("hobby")

		c.String(http.StatusOK, fmt.Sprintf("type1:%s,username:%s,password:%s,hobby:%v",
			type1, username, password, hobby))
	})
	//3.监听端口，默认在8080
	r.Run(":8000")
}
