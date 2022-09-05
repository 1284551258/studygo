package db

import (
	"github.com/1284551258/blog/model"
	"github.com/jmoiron/sqlx"
)

/*
文章分类相关的操作（添加、查询、查询1个分类、查询多个分类、查所有分类）
*/

// InsertCategory 添加分类
func InsertCategory(category *model.Category) (categoryId int64, err error) {
	sqlStr := "INSERT INTO category (category_name, category_no) VALUES(?, ?); "
	result, err := DB.Exec(sqlStr, category.CategoryName, category.CategoryNo)
	if err != nil {
		return 0, err
	}
	categoryId, err = result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return
}

// GetCategoryById 获取单个分类
func GetCategoryById(id int64) (category *model.Category, err error) {
	category = new(model.Category)
	sqlStr := "SELECT id, category_name, category_no FROM category where id = ? ;"

	err = DB.Get(category, sqlStr, id)
	if err != nil {
		return nil, err
	}
	return
}

// GetCategoryList 获取多个分类
func GetCategoryList(ids []int64) (categoryList *[]*model.Category, err error) {
	categoryList = new([]*model.Category)
	//构造sqlStr
	sqlStr, options, err := sqlx.In("SELECT id, category_name, category_no FROM category where id in (?) ",
		ids)
	if err != nil {
		return nil, err
	}
	err = DB.Select(categoryList, sqlStr, options...)
	if err != nil {
		return nil, err
	}
	return

}

// GetAllCategoryList 获取所有分类
func GetAllCategoryList() (categoryList *[]*model.Category, err error) {
	categoryList = new([]*model.Category)
	//构造sqlStr
	sqlStr := "SELECT id, category_name, category_no FROM category "
	err = DB.Select(categoryList, sqlStr)
	if err != nil {
		return nil, err
	}
	return
}
