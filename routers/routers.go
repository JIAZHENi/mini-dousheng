package routers

import (
	"Git/mini-dousheng/controller"
	"Git/mini-dousheng/middlewares"
	"Git/mini-dousheng/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupRouter(mode string) *gin.Engine {

	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()

	// public directory is used to serve static resources
	r.Static("/static", "./public")
	// r.StaticFS("/static", http.Dir("./static"))

	apiRouter := r.Group("/douyin")

	// basic apis

	apiRouter.POST("/user/register/", controller.Register) // 用户注册
	apiRouter.POST("/user/login/", controller.Login)       // 用户登录
	apiRouter.Use(middlewares.JWTAuth())
	{
		apiRouter.GET("/feed/", controller.Feed)                // 视频流接口
		apiRouter.GET("/user/", controller.UserInfo)            // 用户信息
		apiRouter.POST("/publish/action/", controller.Publish)  // 视频投稿
		apiRouter.GET("/publish/list/", controller.PublishList) // 发布列表

		// extra apis - I
		apiRouter.POST("/favorite/action/", controller.FavoriteAction) // 点赞操作
		apiRouter.GET("/favorite/list/", controller.FavoriteList)      // 点赞列表
		apiRouter.POST("/comment/action/", controller.CommentAction)   // 评论操作
		apiRouter.GET("/comment/list/", controller.CommentList)        // 评论列表

		// extra apis - II
		apiRouter.POST("/relation/action/", controller.RelationAction)     // 关注操作
		apiRouter.GET("/relation/follow/list/", controller.FollowList)     // 关注列表
		apiRouter.GET("/relation/follower/list/", controller.FollowerList) // 粉丝列表
	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, model.Response{StatusCode: 1, StatusMsg: "你请求的页面不存在"})
	})

	return r
}
