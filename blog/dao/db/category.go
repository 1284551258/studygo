package db

import "github.com/1284551258/blog/model"

/*
文章分类相关的操作（添加、查询、查询1个分类、查询多个分类、查所有分类）
*/

// InsertCategory 添加函数
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

func QueryCategoryById(id int64) (category *model.Category, err error) {
	category = new(model.Category)
	sqlStr := "SELECT id, category_name, category_no, create_time, update_time FROM blogger.category where id = ? ;"

	err = DB.Get(category, sqlStr, id)
	if err != nil {
		return nil, err
	}
	return
}
