package db

import (
	"code.oldboyedu.com/go_5_blog/model"
	_ "github.com/go-sql-driver/mysql"
)

// 添加文章
func InsertArticle(article *model.ArticleDetail) (articleId int64, err error) {
	// 加个验证
	if article == nil {
		return
	}
	sqlstr := `insert into
               article(content,summart,title,username,category_id,view_count,comment_count) values(?,?,?,?,?,?,?)`
	result, err := DB.Exec(sqlstr, article.Content, article.Summary, article.Title, article.Username, article.ArticleInfo.CategoryId,
		article.ArticleInfo.ViewCount, article.ArticleInfo.CommentCount)
	if err != nil {
		return
	}
	articleId, err = result.LastInsertId()
	return
}

// 获取文章列表，做个分页
func GetAricleList(pageNum, pageSize int) (articleList []*model.ArticleInfo, err error) {
	if pageNum < 0 || pageSize <= 0 {
		return
	}
	// 时间降序排序
	sqlstr := `select
                    id,summary,title,view_count,create_time,comment_count,username,category_id
               from
                    article
               where
                    status=1
               order by create_time desc
               limit ?,?`
	err = DB.Select(&articleList, sqlstr, pageNum, pageSize)
	return
}
