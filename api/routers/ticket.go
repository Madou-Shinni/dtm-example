package routers

import (
	"github.com/Madou-Shinni/gin-quickstart/api/handle"
	"github.com/gin-gonic/gin"
)

// 注册路由
func TicketRouterRegister(r *gin.Engine) {
	ticketGroup := r.Group("ticket")
	ticketHandle := handle.NewTicketHandle()
	{
		ticketGroup.POST("", ticketHandle.Add)
		ticketGroup.DELETE("", ticketHandle.Delete)
		ticketGroup.DELETE("/delete-batch", ticketHandle.DeleteByIds)
		ticketGroup.GET("", ticketHandle.Find)
		ticketGroup.GET("/list", ticketHandle.List)
		ticketGroup.PUT("", ticketHandle.Update)
	}
}