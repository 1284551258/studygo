package db

import (
	"github.com/1284551258/blog/model"
)

// 插入文章

func InsertArticle(article *model.ArticleDetail) (articleId int64, err error) {
	if article == nil {
		return
	}
	sqlStr := `INSERT INTO article
			(category_id,content, title, view_count, comment_count, username, summary,status)
			VALUES(?,?,?,?,?,?,?,?);`
	result, err := DB.Exec(sqlStr, article.CategoryId, article.Content, article.Title, article.ViewCount, article.CommentCount,
		article.Username, article.Summary, article.Status)
	if err != nil {
		return
	}
	articleId, err = result.LastInsertId()
	if err != nil {
		return
	}
	return
}

// GetArticleList 获取文章列表，做个分页
func GetArticleList(pageNum, pageSize int) (articleList *[]*model.ArticleInfo, err error) {
	articleList = new([]*model.ArticleInfo)
	if pageSize <= 0 || pageNum < 0 {
		return
	}
	sqlStr := `SELECT id, category_id, title, view_count, comment_count,
				username, status, summary FROM article where status = 1 order by create_time limit ?,?;`
	err = DB.Select(articleList, sqlStr, pageNum, pageSize)
	if err != nil {
		return nil, err
	}
	return
}

// GetArticleById 根据文章id查询单个文章
func GetArticleById(id int64) (article *model.ArticleDetail, err error) {
	article = new(model.ArticleDetail)
	sqlStr := `SELECT id,content, category_id, content, title, view_count, comment_count,
				username, status, summary FROM article where status = 1 and id = ?;`
	err = DB.Get(article, sqlStr, id)
	if err != nil {
		return nil, err
	}
	return
}

// GetArticleByCategoryId 根据分类id查询一类文章
func GetArticleByCategoryId(categoryId int64, pageNum, pageSize int) (articleList *[]*model.ArticleInfo, err error) {
	articleList = new([]*model.ArticleInfo)
	if pageSize <= 0 || pageNum < 0 {
		return
	}
	sqlStr := `SELECT id, category_id, title, view_count, comment_count,
				username, status, summary FROM article where status = 1 and category_id = ? 
				order by create_time limit ?,?;`
	err = DB.Select(articleList, sqlStr, categoryId, pageNum, pageSize)
	if err != nil {
		return nil, err
	}
	return
}
