package v1

import (
	"ginBlog/model"
	"ginBlog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//添加文章
func AddArticle(c *gin.Context) {
	var article model.Article
	c.ShouldBindJSON(&article)
	code := model.CreateArticle(&article)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    article,
		"message": errmsg.GetErrMsg(code),
	})
}

// GetCateArt 查询分类下的所有文章
func GetCateArt(c *gin.Context) {
	pagesize, _ := strconv.Atoi(c.Query("pagesize"))
	pagenum, _ := strconv.Atoi(c.Query("pagenum"))
	cid, _ := strconv.Atoi(c.Param("cid"))

	switch {
	case pagesize > 100:
		pagesize = 100
	case pagesize <= 0:
		pagesize = 10
	}
	if pagenum == 0 {
		pagenum = 1
	}
	articles, code, total := model.GetArticleByCate(cid, pagesize, pagenum)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    articles,
		"total":   total,
		"message": errmsg.GetErrMsg(code),
	})
}

//查询单个文章
func GetArticleInfo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	article, code := model.GetArticleInfo(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    article,
		"message": errmsg.GetErrMsg(code),
	})
}

//查询文章列表

func GetArticles(c *gin.Context) {
	pagesize, _ := strconv.Atoi(c.Query("pagesize"))
	pagenum, _ := strconv.Atoi(c.Query("pagenum"))

	switch {
	case pagesize > 100:
		pagesize = 100
	case pagesize <= 0:
		pagesize = 10
	}
	if pagenum == 0 {
		pagenum = 1
	}
	articles, code, total := model.GetArticles(pagesize, pagenum)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    articles,
		"total":   total,
		"message": errmsg.GetErrMsg(code),
	})
}

//编辑文章
func EditArticle(c *gin.Context) {
	var article model.Article
	id, _ := strconv.Atoi(c.Param("id"))
	c.ShouldBindJSON(&article)
	code := model.EditArticle(id, &article)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    article,
		"message": errmsg.GetErrMsg(code),
	})
}

//删除文章
func DeleteArticle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code := model.DeleteArticleById(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}
