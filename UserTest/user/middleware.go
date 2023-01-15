package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func JWT() gin.HandlerFunc {
	return func(context *gin.Context) {

		token := context.GetHeader("token")
		if token == "" {
			context.JSON(http.StatusBadRequest, gin.H{
				"status":  JSON_ERROR,
				"message": "未带有token",
			})
			context.Abort()
			return
		} else {
			claims, err := ParseToken(token)
			if err != nil {
				context.JSON(http.StatusBadRequest, gin.H{
					"status":  JSON_ERROR,
					"message": "token错误",
				})
				context.Abort()
				return
			} else if time.Now().Unix() > claims.ExpiresAt {
				context.JSON(http.StatusBadRequest, gin.H{
					"status":  JSON_ERROR,
					"message": "token已过期",
				})
				context.Abort()
				return
			}
		}
		context.Next()
	}
}

/*
 func GetCookie(user userModel) gin.HandlerFunc {
	 return func(context *gin.Context) {
	 	cookie, err := context.Request.Cookie(user.Name + "_cookie")
		 if err == nil {
		 	context.SetCookie(cookie.Name, cookie.Value, 1000, cookie.Path, cookie.Domain, cookie.Secure, cookie.HttpOnly)
		 	context.Next()
		 } else {
		 	context.Abort()
		 }
	 }
 }
注释原因见main函数
*/
