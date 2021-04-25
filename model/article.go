package model

import (
	"ginBlog/utils/errmsg"
	"github.com/jinzhu/gorm"
)

type Article struct {
	Category Category `gorm:"foreignKey:Cid"`

	gorm.Model
	Title string `gorm:"type:varchar(100);not null" json:"title"`
	//Category的id
	Cid int `gorm:"type:int;not null" json:"cid"`
	//描述
	Desc    string `gorm:"type:varchar(200)" json:"desc"`
	Content string `gorm:"type:longtext" json:"content"`
	Img     string `gorm:"type:varchar(100)" json:"img"`
}

// 创建文章
func CreateArticle(article *Article) int {
	err := db.Create(article).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// GetCateArt 查询分类下的所有文章
func GetArticleByCate(cid int, pagesize int, pagenum int) ([]Article, int, int64) {
	var articles []Article
	var total int64

	err := db.Select("article.id, title, img, created_at, updated_at, `desc`").Limit(pagesize).Offset((pagenum-1)*pagesize).Order("Created_At DESC").
		Joins("Category").Where("cid = ?", cid).Find(&articles).Error
	db.Model(&articles).Where("cid = ?", cid).Count(&total)
	if err != nil {
		return nil, errmsg.ERROR_CATE_NOT_EXIST, 0
	}
	return articles, errmsg.SUCCESS, total
}

func GetArticleInfo(id int) (Article, int) {
	var article Article
	code := errmsg.SUCCESS
	err := db.Where("id = ?", id).First(&article).Error
	if err != nil {
		code = errmsg.ERROR_ART_NOT_EXIST
	}
	return article, code
}

// 查询所有文章
func GetArticles(pagesize, pagenum int) ([]Article, int, int64) {
	var articleList []Article
	var err error
	var total int64
	err = db.Select("article.id, title, img, created_at, updated_at, `desc`").Limit(pagesize).Offset((pagenum - 1) * pagesize).Order("Created_At DESC").Joins("Category").Find(&articleList).Error
	db.Model(&articleList).Count(&total)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errmsg.ERROR, 0
	}
	return articleList, errmsg.SUCCESS, total
}

// 根据id号删除文章
func DeleteArticleById(id int) int {
	var article Article
	err := db.Where("id = ?", id).Delete(&article).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 编辑文章
func EditArticle(id int, article *Article) int {
	maps := make(map[string]interface{})
	maps["title"] = article.Title
	maps["cid"] = article.Cid
	maps["desc"] = article.Desc
	maps["content"] = article.Content
	maps["img"] = article.Img

	code := errmsg.SUCCESS
	err := db.Model(article).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		code = errmsg.ERROR
	}
	return code
}
