package service

import (
	"github.com/1284551258/blog/dao/db"
	"github.com/1284551258/blog/model"
)

// GetArticleRecordList 获取文章和对应分类
func GetArticleRecordList(pageNum, pageSize int) (articleRecordList *[]*model.ArticleRecord, err error) {
	articleRecordList = new([]*model.ArticleRecord)
	//按照分页获取文章
	articleList, err := db.GetArticleList(pageNum, pageSize)
	if err != nil {
		return nil, err
	}
	if len(*articleList) <= 0 {
		return
	}
	//获取所有文章对应的分类
	ids, err := GetCategoryList(articleList)
	if err != nil {
		return nil, err
	}
	categoryList, err := db.GetCategoryList(ids)
	if err != nil {
		return nil, err
	}

	//把文章和分类进行聚合
	for _, articleInfo := range *articleList {
		articleRecord := &model.ArticleRecord{
			ArticleInfo: *articleInfo,
		}

		for _, category := range *categoryList {
			if category.CategoryID == articleInfo.CategoryId {
				articleRecord.Category = *category
				break
			}
		}
		*articleRecordList = append(*articleRecordList, articleRecord)
	}
	return

}

func GetCategoryList(articleList *[]*model.ArticleInfo) (ids []int64, err error) {
LABEL:
	for _, articleInfo := range *articleList {
		for _, id := range ids {
			if articleInfo.CategoryId == id {
				continue LABEL
			}
		}
		ids = append(ids, articleInfo.CategoryId)
	}
	return
}

// GetArticleRecordListById 按照分类获取文章和对应分类
func GetArticleRecordListById(categoryId int64, pageNum, pageSize int) (articleRecordList *[]*model.ArticleRecord, err error) {
	articleRecordList = new([]*model.ArticleRecord)
	//按照分页获取文章
	articleList, err := db.GetArticleByCategoryId(categoryId, pageNum, pageSize)
	if err != nil {
		return nil, err
	}
	if len(*articleList) <= 0 {
		return
	}
	//获取所有文章对应的分类
	ids, err := GetCategoryList(articleList)
	if err != nil {
		return nil, err
	}
	categoryList, err := db.GetCategoryList(ids)
	if err != nil {
		return nil, err
	}

	//把文章和分类进行聚合
	for _, articleInfo := range *articleList {
		articleRecord := &model.ArticleRecord{
			ArticleInfo: *articleInfo,
		}

		for _, category := range *categoryList {
			if category.CategoryID == articleInfo.CategoryId {
				articleRecord.Category = *category
				break
			}
		}
		*articleRecordList = append(*articleRecordList, articleRecord)
	}
	return
}
