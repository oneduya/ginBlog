package model

import (
	"ginBlog/utils/errmsg"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"log"
)

//加密使用的参数，决定了加密的复杂性
const hashCost = 4

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20)" json:"username"`
	Password string `gorm:"type:varchar(100)" json:"password"`
	Role     int    `gorm:"type:int" json:"role"`
}

// 创建用户
func CreateUser(user *User) int {
	err := db.Create(user).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//创建之前先加密
func (u *User) BeforeCreate(_ *gorm.DB) (err error) {
	u.Password = GenerateSecret(u.Password)
	return
}

//检查是否重复
func CheckUserByName(name string) int {
	var user User
	db.Where("username = ?", name).First(&user)
	if user.ID > 0 {
		return errmsg.ERROR_USERNAME_USED
	}
	return errmsg.SUCCESS
}

//获取用户列表
func GetUsers(pagesize, pagenum int) []User {
	users := []User{}
	err = db.Limit(pagesize).Offset((pagenum - 1) * pagesize).Find(&users).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return users
}

//删除用户
func DeleteUserById(id int) int {
	var user User
	err := db.Where("id = ?", id).Delete(&user).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//编辑用户
func EditUser(id int, user *User) int {
	maps := make(map[string]interface{})
	maps["username"] = user.Username
	maps["role"] = user.Role
	code := errmsg.SUCCESS
	err := db.Model(user).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		code = errmsg.ERROR
	}
	return code
}

//用golang自带的bcrypt来创建加密的密码
func GenerateSecret(password string) string {
	pw, err := bcrypt.GenerateFromPassword([]byte(password), hashCost)
	if err != nil {
		log.Fatal(err)
	}
	return string(pw)
}
