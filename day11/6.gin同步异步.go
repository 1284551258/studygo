package main

/*
import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func main() {
	r := gin.Default()

	//1.异步执行
	r.GET("/login_async", func(c *gin.Context) {
		//不能直接使用原来的c，需要创建一个副本cCopy
		cCopy := c.Copy()
		go func() {
			time.Sleep(3 * time.Second)
			log.Println("异步执行：", cCopy.Request.URL.Path)
		}()
	})
	//2.同步执行
	r.GET("/login_sync", func(c *gin.Context) {
		time.Sleep(3 * time.Second)
		log.Println("同步执行：", c.Request.URL.Path)
	})
	r.Run(":8000")
}

*/
