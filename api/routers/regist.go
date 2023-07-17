package routers

import (
	"github.com/Madou-Shinni/gin-quickstart/api/handle"
	"github.com/gin-gonic/gin"
)

// 注册路由
func RegistRouterRegister(r *gin.Engine) {
	registGroup := r.Group("regist")
	registHandle := handle.NewRegistHandle()
	{
		registGroup.POST("", registHandle.Add)
		registGroup.POST("/regist-rollback", registHandle.AddRollBack)
		registGroup.DELETE("", registHandle.Delete)
		registGroup.DELETE("/delete-batch", registHandle.DeleteByIds)
		registGroup.GET("", registHandle.Find)
		registGroup.GET("/list", registHandle.List)
		registGroup.PUT("", registHandle.Update)
		registGroup.POST("/trans", registHandle.Trans)
	}
}
