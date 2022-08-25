package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

/*
gin可以构建中间件，但它支队注册过的路由函数起作用。
对于分组路由，嵌套使用中间件，可以限定中间件的作用范围。
中间件分为全局中间件，单个路由中间件和群组中间件。
gin中间件必须是一个*gin.HandlerFunc类型
*/

//定义一个全局的中间件
func MiddleWare() gin.HandlerFunc {

	return func(c *gin.Context) {
		t1 := time.Now()
		fmt.Println("中间件开始执行了")
		//c.Set()设置变量到Context的key中，可以通过Get()获取
		c.Set("request", "中间件")
		//执行中间件
		c.Next()
		//中间件执行完后续的一些事
		//获取响应状态
		status := c.Writer.Status()
		//time.Sleep(3 * time.Second)
		fmt.Printf("中间件执行完毕，状态码为：%d\n", status)
		t2 := time.Since(t1)
		fmt.Printf("用时为：%v\n", t2)
	}
}

func main() {
	r := gin.Default()

	//注册全局中间件
	r.Use(MiddleWare())
	//{}规范
	{
		r.GET("/middleware", func(c *gin.Context) {
			request, _ := c.Get("request")
			fmt.Printf("get request:%v\n", request)
			c.JSON(http.StatusOK, gin.H{"request": request})
		})
		//路径后面是定义的局部中间件
		r.GET("/middleware2", MiddleWare(), func(c *gin.Context) {
			request, _ := c.Get("request")
			fmt.Printf("get request:%v\n", request)
			c.JSON(http.StatusOK, gin.H{"request": request})
		})
	}

	r.Run(":8000")
}
