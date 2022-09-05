package main

import (
	"github.com/1284551258/blog/controller"
	"github.com/1284551258/blog/dao/db"
	"github.com/gin-gonic/gin"
)

func main() {
	//定义默认路由

	r := gin.Default()
	//初始化数据库
	dns := "root:root@tcp(11.2.2.128:3306)/blogger?parseTime=true"
	err := db.Init(dns)
	if err != nil {
		panic(err)
	}
	//加载静态文件
	r.Static("/static/", "./static")

	//加载模板
	r.LoadHTMLGlob("./views/*")

	r.GET("/", controller.IndexHandle)
	r.GET("/category/", controller.CategoryList)

	//启动
	_ = r.Run(":8000")

}
