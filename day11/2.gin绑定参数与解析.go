package main

/*
//json参数解析和绑定
type User struct {
	UserName string `json:"user" xml:"username" form:"username" url:"username" binding:"required"`
	PassWord string `json:"password" xml:"password" form:"password" url:"password" binding:"required"`
}

func main() {
	//创建默认路由引擎
	r := gin.Default()

	r.POST("/loginJson", func(c *gin.Context) {

		//定义一个接收变量
		var j User
		err := c.ShouldBindJSON(&j)
		if err != nil {
			//gin.H封装了生成json数据的工具
			c.JSON(http.StatusBadRequest, gin.H{"data": err.Error()})
			return
		}
		if j.UserName != "root" || j.PassWord != "admin" {
			c.JSON(http.StatusBadRequest, gin.H{"data": "304"})
		} else {
			c.JSON(http.StatusOK, gin.H{"data": "200"})
		}
		return
	})
	//运行
	r.Run(":8000")

	//测试 curl 127.0.0.1:8000/loginJson -X POST -H 'content-type:application/json' -d "{\"user\":\"root\",\"password\":\"admin\"}"
}
*/

/*
//form参数解析与绑定
type User struct {
	UserName string `json:"user" xml:"username" form:"username" url:"username" binding:"required"`
	PassWord string `json:"password" xml:"password" form:"password" url:"password" binding:"required"`
}

func main() {
	//创建默认路由引擎
	r := gin.Default()

	r.POST("/loginForm", func(c *gin.Context) {

		//定义一个接收变量
		var f User
		//Bind()默认解析并绑定form格式
		//根据请求头中的content-type自动推断
		err := c.Bind(&f)
		if err != nil {
			//gin.H封装了生成json数据的工具
			c.JSON(http.StatusBadRequest, gin.H{"data": err.Error()})
			return
		}
		if f.UserName != "root" || f.PassWord != "admin" {
			c.JSON(http.StatusBadRequest, gin.H{"data": "304"})
		} else {
			c.JSON(http.StatusOK, gin.H{"data": "200"})
		}
		return
	})
	//运行
	r.Run(":8000")

}

*/

// uri参数解析与绑定

/*
type User struct {
	UserName string `json:"user" xml:"username" form:"username" uri:"username" binding:"required"`
	PassWord string `json:"password" xml:"password" form:"password" uri:"password" binding:"required"`
}

func main() {
	//创建默认路由引擎
	r := gin.Default()

	r.GET("/:username/:password", func(c *gin.Context) {

		//定义一个接收变量
		var u User
		err := c.ShouldBindUri(&u)
		if err != nil {
			//gin.H封装了生成json数据的工具
			c.JSON(http.StatusBadRequest, gin.H{"data": err.Error()})
			return
		}
		if u.UserName != "root" || u.PassWord != "admin" {
			c.JSON(http.StatusBadRequest, gin.H{"data": "304"})
		} else {
			c.JSON(http.StatusOK, gin.H{"data": "200"})
		}
		return
	})
	//运行
	r.Run(":8000")
	//测试 curl 127.0.0.1:8000/root/admin

}
*/
