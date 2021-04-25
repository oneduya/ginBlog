package v1

import (
	"ginBlog/model"
	"ginBlog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//查询分类名是否存在

//添加分类
func AddCategory(c *gin.Context) {
	var category model.Category
	c.ShouldBindJSON(&category)
	code := model.CheckCateByName(category.Name)
	if code == errmsg.SUCCESS {
		code = model.CreateCate(&category)
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    category,
		"message": errmsg.GetErrMsg(code),
	})
}

//查询单个分类

//查询分类信息
func GetCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	category, code := model.GetCateInfo(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    category,
		"message": errmsg.GetErrMsg(code),
	})
}

//编辑分类
func EditCategory(c *gin.Context) {
	var category model.Category
	id, _ := strconv.Atoi(c.Param("id"))
	c.ShouldBindJSON(&category)
	code := model.CheckCateByName(category.Name)
	if code == errmsg.SUCCESS {
		code = model.EditCate(id, &category)
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    category,
		"message": errmsg.GetErrMsg(code),
	})
}

//删除分类
func DeleteCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code := model.DeleteCateById(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}
