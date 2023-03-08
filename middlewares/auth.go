package middlewares

import (
	"github.com/gin-gonic/gin"
	"strings"
	"web_app/controllers"
	"web_app/pkg/jwt"
)



func JWTAuthMiddleWare() func(c *gin.Context){
	return func(c *gin.Context){
		// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URL
		// 这里假设Token放在Header的Authorization中，并使用Bearer开头
		// 这里的具体实现方式要依据你的实际业务情况决定
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader ==""{
			controllers.ResponseError(c, controllers.CodeNeedLogin)
			//c.JSON(http.StatusOK, gin.H{
			//	"code" : 2003,
			//	"msg" : "请求头auth为空",
			//})
			c.Abort()
			return
		}
		//	按空格分割
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer"){
			controllers.ResponseError(c, controllers.CodeInvalidToken)
			//c.JSON(http.StatusOK, gin.H{
			//	"code" : 2004,
			//	"msg" : "请求头auth格式有误",
			//})
			c.Abort()
			return
		}
		// parts[1]是获取到的token string，我们使用之前定义好的解析JWT的函数解析它
		mc, err := jwt.ParseToken(parts[1])
		if err != nil {
			controllers.ResponseError(c, controllers.CodeInvalidToken)
			//c.JSON(http.StatusOK, gin.H{
			//	"code" : 2005,
			//	"msg" : "无效的token",
			//})
			c.Abort()
			return
		}
		//将当前请求的username信息保存到请求的上下文c上
		c.Set(controllers.CtxUserIDKey, mc.UserId)
		//后续的处理函数可以用 c.Get("username")来获取当前请求的用户信息
		c.Next()
	}
}