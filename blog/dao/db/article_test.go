package db

import (
	"fmt"
	"github.com/1284551258/blog/model"
	"testing"
	"time"
)

func init() {
	dns := "root:root@tcp(11.2.2.128:3306)/blogger?parseTime=true"
	err := Init(dns)
	if err != nil {
		panic(err)
	}
	fmt.Println("数据库连接成功")
}

func TestInsertArticle(t *testing.T) {
	article := &model.ArticleDetail{}
	article.Title = "文章1标题"
	article.Content = "文章1内容"
	article.CategoryId = 1
	article.Summary = "文章1摘要"
	article.Username = "sum"
	article.CreateTime = time.Now()
	article.Status = 1
	insertArticleId, err := InsertArticle(article)
	if err != nil {
		panic(err)
	}
	t.Logf("insertArticleId:%v", insertArticleId)
}

func TestGetArticleById(t *testing.T) {
	articleDetail, err := GetArticleById(1)
	if err != nil {
		panic(err)
	}
	t.Logf("articleDetail：%#v", articleDetail)
}

func TestGetArticleList(t *testing.T) {
	articleList, err := GetArticleList(0, 15)
	if err != nil {
		panic(err)
	}
	t.Logf("articleList的数量为：%#v", len(*articleList))
}

func TestGetArticleByCategoryId(t *testing.T) {
	articleByCategoryList, err := GetArticleByCategoryId(1, 0, 15)
	if err != nil {
		panic(err)
	}
	t.Logf("articleList的数量为：%#v", len(*articleByCategoryList))
}
