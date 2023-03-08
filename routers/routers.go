package routers

import (
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"web_app/controllers"
	"web_app/logger"
	"web_app/middlewares"
)

func SetupRouter()*gin.Engine{
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true), middlewares.RateLimitMiddleware(2*time.Second, 1))

	v1 := r.Group("/api/v1")

	//注册业务路由
	v1.POST("/signup", controllers.SignUpHandler)
	//登录
	v1.POST("/login", controllers.LoginHandler)

	v1.Use(middlewares.JWTAuthMiddleWare())
	{
		v1.GET("/community", controllers.CommunityHandler)
		v1.GET("/community/:id", controllers.CommunityDetailHandler)

		v1.POST("/post", controllers.CreatePostHandler)
		v1.GET("/post/:id", controllers.CreatePostDetailHandler)
		v1.GET("/posts", controllers.CreatePostListHandler)

		//根据时间或分数获取帖子列表
		v1.GET("/posts2", controllers.CreatePostListHandler2)

		//投票
		v1.POST("/vote", controllers.PostVoteController)
	}

	//r.GET("/ping", middlewares.JWTAuthMiddleWare(), func(c *gin.Context){
	//	//如果是登录用户，判断请求头里面是否有有效的JWT
	//	c.String(http.StatusOK, "ok")
	//})
	pprof.Register(r) //注册pprof相关路由
	r.NoRoute(func(c *gin.Context){
		c.JSON(http.StatusOK, gin.H{
			"msg":  "404",
		})
	})

	return r
}


