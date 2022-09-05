package service

import (
	"github.com/1284551258/blog/dao/db"
	"github.com/1284551258/blog/model"
)

// GetAllCategoryList 获取所有分类
func GetAllCategoryList() (allCategoryList *[]*model.Category, err error) {
	allCategoryList, err = db.GetAllCategoryList()
	if err != nil {
		return nil, err
	}
	return
}
