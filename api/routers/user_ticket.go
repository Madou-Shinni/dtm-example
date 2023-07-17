package routers

import (
	"github.com/Madou-Shinni/gin-quickstart/api/handle"
	"github.com/gin-gonic/gin"
)

// 注册路由
func UserTicketRouterRegister(r *gin.Engine) {
	userTicketGroup := r.Group("userTicket")
	userTicketHandle := handle.NewUserTicketHandle()
	{
		userTicketGroup.POST("", userTicketHandle.Add)
		userTicketGroup.POST("/add-rollback", userTicketHandle.Delete)
		userTicketGroup.DELETE("", userTicketHandle.Delete)
		userTicketGroup.DELETE("/delete-batch", userTicketHandle.DeleteByIds)
		userTicketGroup.GET("", userTicketHandle.Find)
		userTicketGroup.GET("/list", userTicketHandle.List)
		userTicketGroup.PUT("", userTicketHandle.Update)
	}
}
