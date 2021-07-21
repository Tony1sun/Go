package service

import (
	"code.oldboyedu.com/go_5_blog/dao/db"
	"code.oldboyedu.com/go_5_blog/model"
)

// 获取文章和对应的分类
func GetArticleRecordList(pageNum, pageSize int) (articleRecordList []model.ArticleRecord, err error) {
	// 获取文章列表
	articleInfolist, err := db.GetAricleList(pageNum, pageSize)
	if err != nil {
		return
	}
	if len(articleInfolist) <= 0 {
		return
	}
	// 获取文章对应的分类(多个)
}
