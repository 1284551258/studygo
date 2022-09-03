package main

/*
import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	//注册权限验证中间件
	r.Use(admitWare)

	r.GET("/login", func(c *gin.Context) {
		c.SetCookie("is_login", "yes", 60, "/",
			"localhost", false, true)
		fmt.Printf("login successful!\n")
		c.String(http.StatusOK, "login successful！")
	})
	r.GET("/home", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "/home"})
	})
	r.Run(":8000")
}

func admitWare(c *gin.Context) {
	//检查是否携带cookie
	cookie, err := c.Cookie("is_login")
	if err != nil || cookie != "yes" {
		fmt.Printf("Not found cookie,now set cookie or cookie err,cookie:%s,err:%v,path:%v\n", cookie, err, c.Request.URL.Path)
		if c.Request.URL.Path != "/login" {
			fmt.Println("direct return")
			c.String(http.StatusUnauthorized, "not login")
			c.Abort()
			return
		}
	}
	fmt.Printf("start %s handlerFun\n", c.Request.URL.Path)
	c.Next()
	return
}


*/
