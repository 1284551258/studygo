package db

import (
	"fmt"
	"github.com/1284551258/blog/model"
	"testing"
)

func init() {
	dns := "root:root@tcp(11.2.2.128:3306)/blogger?parseTime=true"
	err := Init(dns)
	if err != nil {
		panic(err)
	}
	fmt.Println("数据库连接成功")
}

func TestInsertCategory(t *testing.T) {
	c := &model.Category{
		CategoryName: "python",
		CategoryNo:   3,
	}
	categoryId, err := InsertCategory(c)
	if err != nil {
		panic(err)
	}
	t.Logf("categoryId:%d", categoryId)
}

func TestGetCategoryById(t *testing.T) {
	cate, err := GetCategoryById(1)
	if err != nil {
		panic(err)
	}
	t.Logf("category:%#v", cate)
}

func TestGetCategoryList(t *testing.T) {
	categoryList, err := GetCategoryList([]int64{1, 2})
	if err != nil {
		panic(err)
	}
	t.Logf("category:%d", len(*categoryList))
	for i, category := range *categoryList {
		t.Logf("第%d个category:%#v", i, category)
	}
}

func TestGetAllCategoryList(t *testing.T) {
	categoryList, err := GetAllCategoryList()
	if err != nil {
		panic(err)
	}
	t.Logf("category:%d", len(*categoryList))
	for i, category := range *categoryList {
		t.Logf("第%d个category:%#v", i, category)
	}
}
