package main

import (
	"book/book"
	"fmt"
	"net/http"
	"strconv"
)

import "github.com/gin-gonic/gin"

func main() {

	//初始化数据库
	err := book.IniDB()
	if err != nil {
		fmt.Println("初始化数据库失败，err:", err)
		return
	}
	r := gin.Default()
	//加载模板文件
	r.LoadHTMLGlob("./templates/*")

	r.GET("/book/list", bookQueryHandlerFun)
	r.GET("/book/new", bookNewHandlerFun)
	r.POST("/book/add", bookAddHandlerFun)
	r.GET("/book/delete", bookDeleteHandlerFun)

	_ = r.Run(":8000")

}

func bookDeleteHandlerFun(c *gin.Context) {
	value := c.Query("id")
	valueInt, err := strconv.Atoi(value)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err,
		})
	}
	err = book.DeleteBook(int64(valueInt))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "删除成功，删除的id为" + value,
	})
}

func bookAddHandlerFun(c *gin.Context) {
	title := c.PostForm("title")
	price := c.PostForm("price")
	priceInt, err := strconv.Atoi(price)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err,
		})
		return
	}

	err = book.InsertBook(title, int64(priceInt))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "插入成功",
	})
}

func bookNewHandlerFun(c *gin.Context) {
	c.HTML(http.StatusOK, "new_book.html", gin.H{
		"data": "ok",
	})
}

func bookQueryHandlerFun(c *gin.Context) {
	allBook, err := book.QueryAllBook()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err,
		})
		return
	}

	c.HTML(http.StatusOK, "book_list.html", gin.H{
		"data": allBook,
	})

}
