package db

import (
	"code.oldboyedu.com/go_5_blog/model"
	"testing"
	"time"
)

func init() {
	// parseTime=true 将mysql中的时间类型，自动解析为go结构体中的时间类型
	// 不加报错
	dns := "root:123456@tcp(127.0.0.1:3306)/test?parseTime=true"
	err := Init(dns)
	if err != nil {
		panic(err)
	}
}

// 测试 添加文章
func TestInsertArticle(t *testing.T) {
	// 构建对象
	article := &model.ArticleDetail{}
	article.ArticleInfo.CategoryId = 10
	article.ArticleInfo.CommentCount = 10
	article.Content = "shangbuzaigarewro"
	article.ArticleInfo.CreateTime = time.Now()
	article.ArticleInfo.Title = "哈哈哈3"
	article.ArticleInfo.Username = "fsgf4"
	article.ArticleInfo.Summary = "abvfdf4"
	article.ArticleInfo.ViewCount = 2
	articleId, err := InsertArticle(article)
	if err != nil {
		return
	}
	t.Logf("articleId:%d\n", articleId)
}

func TestGetAricleList(t *testing.T) {
	articleList, err := GetAricleList(1, 15)
	if err != nil {
		return
	}
	t.Logf("crticle: %d\n", len(articleList))
}

func TestGetarticleDetail(t *testing.T) {
	detail, err := GetarticleDetail(8)
	if err != nil {
		return
	}
	t.Logf("detail: %#v\n", detail)
}
