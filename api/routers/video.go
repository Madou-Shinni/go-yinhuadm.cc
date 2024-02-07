package routers

import (
	"github.com/Madou-Shinni/gin-quickstart/api/handle"
	"github.com/gin-gonic/gin"
)

// 注册路由
func VideoRouterRegister(r *gin.RouterGroup) {
	videoGroup := r.Group("video")
	videoHandle := handle.NewVideoHandle()
	{
		videoGroup.POST("", videoHandle.Add)
		videoGroup.DELETE("", videoHandle.Delete)
		videoGroup.DELETE("/delete-batch", videoHandle.DeleteByIds)
		videoGroup.GET("", videoHandle.Find)
		videoGroup.GET("/play", videoHandle.Play)
		videoGroup.GET("/home", videoHandle.Home)
		videoGroup.GET("/list", videoHandle.List)
		videoGroup.PUT("", videoHandle.Update)
	}
}
