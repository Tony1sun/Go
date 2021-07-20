package db

import "testing"

func init() {
	// parseTime=true 将mysql中的时间类型，自动解析为go结构体中的时间类型
	// 不加报错
	dns := "root:root@tcp(127.0.0.1:3306)/test?parseTime=true"
	err := Init(dns)
	if err != nil {
		panic(err)
	}
}

// 获取单个分类信息
func TestGetCategoryById(t *testing.T) {
	category, err := GetCategoryById(1)
	if err != nil {
		panic(err)
	}
	t.Logf("category:%#v", category)
}

func TestGetCategoryList(t *testing.T) {
	var categoryIds []int64
	categoryIds = append(categoryIds, 1, 2, 3)
	categoryList, err := GetCategoryList(categoryIds)
	if err != nil {
		panic(err)
	}
	for _, v := range categoryList {
		t.Logf("id:%d category:%#v\n", v.CategoryId, v)
	}
}

func TestGetCategoryAll(t *testing.T) {
	categoryList, err := GetCategoryAll()
	if err != nil {
		panic(err)
	}
	for _, v := range categoryList {
		t.Logf("id:%d category:%#v\n", v.CategoryId, v)
	}
}
