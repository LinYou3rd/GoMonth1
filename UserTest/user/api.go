package user

import (
	"UserTest/todo"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

const (
	JSON_SUCCESS int = 1
	JSON_ERROR   int = 0
)

func Register(context *gin.Context) {

	var user userModel
	var count int64
	db.Where("name=?", context.PostForm("name")).First(&user).Count(&count)
	if count == 1 {
		context.JSON(http.StatusOK, gin.H{
			"status":  JSON_ERROR,
			"message": "已有该用户，请勿重名",
		})
		return
	}

	user.Name = context.PostForm("name")
	bytes, err := bcrypt.GenerateFromPassword([]byte(context.PostForm("password")), 1)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"status":  JSON_ERROR,
			"message": "账户由于密码原因创建失败",
		})
		return
	}

	user.PasswordDigest = string(bytes)
	user.Email = context.PostForm("email")
	db.Save(&user)

	context.JSON(http.StatusOK, gin.H{
		"status":  JSON_SUCCESS,
		"message": "创建成功",
		"id":      user.ID,
	})
}

func Login(context *gin.Context) {

	var user userModel
	db.Where("name=?", context.PostForm("name")).First(&user)
	if user.ID == 0 {
		context.JSON(http.StatusOK, gin.H{
			"status":  JSON_ERROR,
			"message": "该用户不存在",
		})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordDigest), []byte(context.PostForm("password")))
	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"status":  JSON_ERROR,
			"message": "密码错误",
		})
		return
	}

	token, err := GenerateToken(user)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"status":  JSON_ERROR,
			"message": "Token签发失败",
		})
		return
	}

	context.SetCookie(user.Name+"_cookie", string(user.ID), 3600, "/", "localhost", false, true)
	// 思索有了token还需不需要cookie，百度看得有点懵

	context.JSON(http.StatusOK, gin.H{
		"status":  JSON_SUCCESS,
		"message": "登录成功",
		"token":   token,
	})

	todo.Init() // 连接上事务表
}
