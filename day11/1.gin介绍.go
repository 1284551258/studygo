package main

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
/*
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
*/
/*
//上传单个文件
func main() {
	//1.创建路由引擎
	r := gin.Default()
	r.POST("/upload", func(c *gin.Context) {
		//表单获取文件
		file, err := c.FormFile("file")
		if err != nil {
			return
		}
		log.Printf("get file name:%s", file.Filename)
		err = c.SaveUploadedFile(file, file.Filename)
		if err != nil {
			log.Printf("save file failed,err:%v", err)
			return
		}
		c.String(200, "upload %s successful!", file.Filename)
	})
	//3.监听端口，默认在8080
	r.Run(":8000")
}
*/
/*
//上传多个文件
func main() {
	//1.创建路由引擎
	r := gin.Default()
	//限制表单上传大小为8MB，默认为32MB
	r.POST("/uploads", func(c *gin.Context) {
		//拿到表单
		form, err := c.MultipartForm()
		if err != nil {
			c.String(http.StatusBadRequest, "get form faild,err:%v", err)
			return
		}
		//获取所有文件
		files := form.File["files"]

		//逐个保存文件
		for _, file := range files {
			err := c.SaveUploadedFile(file, file.Filename)
			if err != nil {
				c.String(http.StatusBadRequest, "save file faild,err:%v", err)
				return
			}
		}
		c.String(http.StatusOK, "upload %d file successful!", len(files))
	})
	//3.监听端口，默认在8080
	r.Run(":8000")
}
*/
/*
//路由组
func main() {
	//1.创建路由引擎
	r := gin.Default()
	v1 := r.Group("/v1")
	//{}是规范
	{
		v1.GET("/login", login)
		v1.GET("/submit", submit)
	}
	v2 := r.Group("/v2")
	//{}是规范
	{
		v2.POST("/login", login)
		v2.POST("/submit", submit)
	}
	//3.监听端口，默认在8080
	r.Run(":8000")
}

func submit(c *gin.Context) {
	name := c.DefaultQuery("name", "jack")
	c.String(http.StatusOK, "hello %s", name)
}

func login(c *gin.Context) {
	name := c.DefaultQuery("name", "livy")
	c.String(http.StatusOK, "hello %s", name)
}

*/

//路由原理
/*
httprouter会将所有的路由规则构造一颗前缀树
*/
