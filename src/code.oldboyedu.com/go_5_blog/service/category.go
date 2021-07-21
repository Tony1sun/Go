package service

import (
	"code.oldboyedu.com/go_5_blog/dao/db"
	"code.oldboyedu.com/go_5_blog/model"
)

// 获取所有分类
func GetAllCategoryList() (categoryList []*model.Category, err error) {
	categoryList, err = db.GetAllCategoryList()
	if err != nil {
		return
	}
	return
}
