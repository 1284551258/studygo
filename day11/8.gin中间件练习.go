package main

/*
import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

//定义一个计算程序运行时间中间件
func funTime(c *gin.Context) {
	t1 := time.Now()
	c.Next()
	t2 := time.Since(t1)
	fmt.Println("程序运行时间为：", t2)
}

func main() {
	//定义一个默认路由引擎
	r := gin.Default()

	//注册中间件
	r.Use(funTime)

	//定义组路由
	shoppingGroup := r.Group("/shopping")
	//规范
	{
		shoppingGroup.GET("/index", shoppingIndexhandlerFunc)
		shoppingGroup.GET("/home", shoppingHomehandlerFunc)
	}

	//指定运行端口
	r.Run(":8000")
}

func shoppingHomehandlerFunc(context *gin.Context) {
	time.Sleep(5 * time.Second)
}

func shoppingIndexhandlerFunc(context *gin.Context) {
	time.Sleep(3 * time.Second)
}


*/
