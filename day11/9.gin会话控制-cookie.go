package main

/*
import (
	"fmt"
	"github.com/gin-gonic/gin"
)

//1.cookie是什么？
//HTTP是无状态协议，服务器不能记录浏览器的访问状态，也就是说服务器
//不能区分两次请求是否由同一个客户端发出。
//Cookie就是解决HTTP协议无状态的方案之一，中文就是小甜饼的意思。
//Cookie实际上就是服务器保存至浏览器上的一段信息，浏览器有了Cookie
//之后，每次向服务器发送请求时都会同时将该信息发送给服务器，服务器
//收到请求后，就可以根据该信息处理请求。
//Cookie由服务器创建，并发送给浏览器，最终由浏览器保存。



//cookie使用

func main() {
	r := gin.Default()

	r.GET("/cookie", func(c *gin.Context) {
		//拿到cookie
		cookie, err := c.Cookie("key_cookie")
		if err != nil {
			cookie = "Not Set"
			//若不存在，则设置cookie
			//name string, value string,
			//maxAge int, 过期时间，单位为s
			//path string, cookie所在目录
			//domain string, 域名 若设置为localhost则127.0.0.1不会设置cookie
			//secure bool, 是否只能通过https访问
			//httpOnly bool 是否允许别人通过js获取自己的cookie true为允许
			c.SetCookie("key_cookie", "value_cookie", 60, "/",
				"localhost", false, true)
		}
		fmt.Println("cookie值为：", cookie)
	})

	r.Run(":8000")
}
//cookie的缺点
不安全，明文
增加带宽消耗
可以被禁用
cookie有上限


*/
