package model

import (
	"ginBlog/utils/errmsg"
)

type Category struct {
	Id   int    `gorm:"primary_key;aoto_increment" json:"id"`
	Name string `gorm:"type:varchar(20);not null" json:"name"`
}

// 创建分类
func CreateCate(category *Category) int {
	err := db.Create(category).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//检查分类名是否存在
func CheckCateByName(name string) int {
	var category Category
	db.Where("name = ?", name).First(&category)
	if category.Id > 0 {
		return errmsg.ERROR_CATENAME_USED
	}
	return errmsg.SUCCESS
}

// 查询单个分类信息
func GetCateInfo(id int) (Category, int) {
	var category Category
	err := db.Where("id = ?", id).First(&category).Error
	if err != nil {
		return category, errmsg.ERROR
	}
	return category, errmsg.SUCCESS
}

// 根据id号删除分类
func DeleteCateById(id int) int {
	var category Category
	err := db.Where("id = ?", id).Delete(&category).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 编辑分类
func EditCate(id int, category *Category) int {
	maps := make(map[string]interface{})
	maps["name"] = category.Name
	code := errmsg.SUCCESS
	err := db.Model(category).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		code = errmsg.ERROR
	}
	return code
}
