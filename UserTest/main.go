package main

import (
	"UserTest/todo"
	"UserTest/user"
	"github.com/gin-gonic/gin"
)

func main() {
	user.Init()
	router := gin.Default()
	v1 := router.Group("/api/v1")
	{
		v1.POST("user/register", user.Register)
		v1.POST("user/login", user.Login)
		// v1.Use(user.GetCookie())
		// 此处注释了一个刷新cookie的中间件，原因是cookie名设定为用户登录成功传入的name+“_cookie”，在该方法中请求cookie同样需要用户的name。未想到如何传入，遂注释，先写其余部分。
		v2 := v1.Group("/")
		v2.Use(user.JWT())
		{
			v2.POST("add", todo.Add)
			v2.GET("show/all", todo.ShowAll)
			v2.GET("show/finish", todo.ShowFinish)
			v2.GET("show/noFinish", todo.ShowNoFinish)
			v2.GET("show/target", todo.ShowTarget) // 不是很理解输入关键词查询是什么程度的关键词，姑且用title查询
			v2.PUT("update/:id", todo.UpdateOne)
			v2.PUT("update/all/:completed", todo.UpdateALL)
			v2.DELETE("delete/:id", todo.DeleteOne)
			v2.DELETE("delete/all/:completed", todo.DeletePart)
			v2.DELETE("delete/all", todo.DeleteAll)
		}
	}
	router.Run(":8000")
}
