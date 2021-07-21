package db

import "testing"

func init() {
	// parseTime=true 将mysql中的时间类型，自动解析为go结构体中的时间类型
	// 不加报错
	dns := "root:123456@tcp(127.0.0.1:3306)/test?parseTime=true"
	err := Init(dns)
	if err != nil {
		panic(err)
	}
}

// 获取单个分类信息
func TestGetCategoryById(t *testing.T) {
	category, err := GetCategoryById(2)
	if err != nil {
		panic(err)
	}
	t.Logf("category:%#v", category)
}
