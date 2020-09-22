package server

import (
	"net/http"
	"os"
	"wiliwili/api"
	"wiliwili/middleware"

	"github.com/gin-gonic/gin"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()
	r.StaticFS("/static", http.Dir("./static"))

	// 为 multipart forms 设置较低的内存限制 (默认是 32 MiB)
	r.MaxMultipartMemory = 100 << 20 // 8 MiB

	// 中间件, 顺序不能改
	r.Use(middleware.Session(os.Getenv("SESSION_SECRET")))
	r.Use(middleware.Cors())
	r.Use(middleware.CurrentUser())

	// 路由
	v1 := r.Group("/api/v1")
	{
		v1.POST("ping", api.Ping)

		// 用户登录
		v1.POST("user/register", api.UserRegister)

		// 用户登录
		v1.POST("user/login", api.UserLogin)

		// 需要登录保护的
		authed := v1.Group("/")
		authed.Use(middleware.AuthRequired())
		{
			// User Routing
			authed.GET("user/me", api.UserMe)
			authed.DELETE("user/logout", api.UserLogout)
		}

		// 视频操作
		v1.POST("videos", api.CreateVideo)
		v1.GET("video/:id", api.ShowVideo)
		v1.GET("videos", api.ListVideo)
		v1.PUT("video/:id", api.UpdateVideo)
		v1.DELETE("video/:id", api.DeleteVideo)
		// 排行榜
		v1.GET("rank/daily", api.DailyRank)
		// 其他
		v1.POST("upload/token", api.UploadToken)
		// 上传图片
		v1.POST("upload/image", api.UploadImage)
		// 上传视频
		v1.POST("upload/video", api.UploadVideo)
	}

	// swagger文档
	// 游览器打开 http://localhost:3000/swagger/index.html
	r.StaticFile("/swagger.json", "./swagger/swagger.json")
	r.Static("/swagger", "./swagger/dist")

	return r
}
