package model

import "time"

//定义article文章基本信息结构体
//id, category_id, content, title, view_count, comment_count, username, status, summary, create_time, update_time

type ArticleInfo struct {
	ArticleId    int64     `db:"id"`
	CategoryId   int64     `db:"category_id"`
	Title        string    `db:"title"`
	ViewCount    uint32    `db:"view_count"`
	CommentCount uint32    `db:"comment_count"`
	Username     string    `db:"username"`
	Summary      string    `db:"summary"`
	CreateTime   time.Time `db:"create_time"`
}

type ArticleDetail struct {
	ArticleInfo
	Content string `db:"content"`
	Category
}
type ArticleRecord struct {
	ArticleInfo
	Category
}
