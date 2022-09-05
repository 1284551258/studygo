package controller

import (
	"github.com/1284551258/blog/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// IndexHandle 访问主页的控制器
func IndexHandle(c *gin.Context) {
	//从service获取数据
	//获取所有文章数据
	articleRecordList, err := service.GetArticleRecordList(0, 15)
	if err != nil {
		//返回错误页面和空数据
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}
	//获取所有分类数据
	categoryList, err := service.GetAllCategoryList()
	if err != nil {
		//返回错误页面和空数据
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}
	//返回主页和数据
	c.HTML(http.StatusOK, "views/index.html", gin.H{
		"article_list":  *articleRecordList,
		"category_list": *categoryList,
	})
	return
}

// CategoryList 点击分类云进行分类
func CategoryList(c *gin.Context) {
	//从service获取数据
	//获取所有文章数据
	categoryID := c.Query("category_id")
	parseInt, err := strconv.ParseInt(categoryID, 10, 64)
	if err != nil {
		//返回错误页面和空数据
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}
	articleRecordList, err := service.GetArticleRecordListById(parseInt, 0, 15)
	if err != nil {
		//返回错误页面和空数据
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}
	//获取所有分类数据
	categoryList, err := service.GetAllCategoryList()
	if err != nil {
		//返回错误页面和空数据
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}
	if len(*articleRecordList) == 0 {
		//返回错误页面和空数据
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}
	//返回主页和数据
	c.HTML(http.StatusOK, "views/index.html", gin.H{
		"article_list":  *articleRecordList,
		"category_list": *categoryList,
	})
	return
}
