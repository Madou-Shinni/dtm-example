package handle

import (
	"github.com/Madou-Shinni/gin-quickstart/internal/domain"
	"github.com/Madou-Shinni/gin-quickstart/internal/service"
	"github.com/Madou-Shinni/gin-quickstart/pkg/constant"
	"github.com/Madou-Shinni/gin-quickstart/pkg/request"
	"github.com/Madou-Shinni/gin-quickstart/pkg/response"
	"github.com/gin-gonic/gin"
)

type UserTicketHandle struct {
	s *service.UserTicketService
}

func NewUserTicketHandle() *UserTicketHandle {
	return &UserTicketHandle{s: service.NewUserTicketService()}
}

// Add 创建UserTicket
// @Tags     UserTicket
// @Summary  创建UserTicket
// @accept   application/json
// @Produce  application/json
// @Param    data body     domain.UserTicket true "创建UserTicket"
// @Success  200  {string} string            "{"code":200,"msg":"","data":{}"}"
// @Router   /userTicket [post]
func (cl *UserTicketHandle) Add(c *gin.Context) {
	var userTicket domain.UserTicket
	if err := c.ShouldBindJSON(&userTicket); err != nil {
		response.Error(c, 409, constant.CODE_INVALID_PARAMETER.Msg())
		return
	}

	if _, err := cl.s.Add(userTicket); err != nil {
		response.Error(c, 409, constant.CODE_ADD_FAILED.Msg())
		return
	} else {
		response.Success(c)
	}
}

// Delete 删除UserTicket
// @Tags     UserTicket
// @Summary  删除UserTicket
// @accept   application/json
// @Produce  application/json
// @Param    data body     domain.UserTicket true "删除UserTicket"
// @Success  200  {string} string            "{"code":200,"msg":"","data":{}"}"
// @Router   /userTicket [delete]
func (cl *UserTicketHandle) Delete(c *gin.Context) {
	var userTicket domain.UserTicket
	if err := c.ShouldBindJSON(&userTicket); err != nil {
		response.Error(c, constant.CODE_INVALID_PARAMETER, constant.CODE_INVALID_PARAMETER.Msg())
		return
	}

	if err := cl.s.Delete(userTicket); err != nil {
		response.Error(c, 409, constant.CODE_DELETE_FAILED.Msg())
		return
	}

	response.Success(c)
}

// DeleteByIds 批量删除UserTicket
// @Tags     UserTicket
// @Summary  批量删除UserTicket
// @accept   application/json
// @Produce  application/json
// @Param    data body     request.Ids true "批量删除UserTicket"
// @Success  200  {string} string            "{"code":200,"msg":"","data":{}"}"
// @Router   /userTicket/delete-batch [delete]
func (cl *UserTicketHandle) DeleteByIds(c *gin.Context) {
	var ids request.Ids
	if err := c.ShouldBindJSON(&ids); err != nil {
		response.Error(c, constant.CODE_INVALID_PARAMETER, constant.CODE_INVALID_PARAMETER.Msg())
		return
	}

	if err := cl.s.DeleteByIds(ids); err != nil {
		response.Error(c, constant.CODE_DELETE_FAILED, constant.CODE_DELETE_FAILED.Msg())
		return
	}

	response.Success(c)
}

// Update 修改UserTicket
// @Tags     UserTicket
// @Summary  修改UserTicket
// @accept   application/json
// @Produce  application/json
// @Param    data body     domain.UserTicket true "修改UserTicket"
// @Success  200  {string} string            "{"code":200,"msg":"","data":{}"}"
// @Router   /userTicket [put]
func (cl *UserTicketHandle) Update(c *gin.Context) {
	var userTicket map[string]interface{}
	if err := c.ShouldBindJSON(&userTicket); err != nil {
		response.Error(c, constant.CODE_INVALID_PARAMETER, constant.CODE_INVALID_PARAMETER.Msg())
		return
	}

	if err := cl.s.Update(userTicket); err != nil {
		response.Error(c, constant.CODE_UPDATE_FAILED, constant.CODE_UPDATE_FAILED.Msg())
		return
	}

	response.Success(c)
}

// Find 查询UserTicket
// @Tags     UserTicket
// @Summary  查询UserTicket
// @accept   application/json
// @Produce  application/json
// @Param    data query     domain.UserTicket true "查询UserTicket"
// @Success  200  {string} string            "{"code":200,"msg":"查询成功","data":{}"}"
// @Router   /userTicket [get]
func (cl *UserTicketHandle) Find(c *gin.Context) {
	var userTicket domain.UserTicket
	if err := c.ShouldBindQuery(&userTicket); err != nil {
		response.Error(c, constant.CODE_INVALID_PARAMETER, constant.CODE_INVALID_PARAMETER.Msg())
		return
	}

	res, err := cl.s.Find(userTicket)

	if err != nil {
		response.Error(c, constant.CODE_FIND_FAILED, constant.CODE_FIND_FAILED.Msg())
		return
	}

	response.Success(c, res)
}

// List 查询UserTicket列表
// @Tags     UserTicket
// @Summary  查询UserTicket列表
// @accept   application/json
// @Produce  application/json
// @Param    data query     domain.UserTicket true "查询UserTicket列表"
// @Success  200  {string} string            "{"code":200,"msg":"查询成功","data":{}"}"
// @Router   /userTicket/list [get]
func (cl *UserTicketHandle) List(c *gin.Context) {
	var userTicket domain.PageUserTicketSearch
	if err := c.ShouldBindQuery(&userTicket); err != nil {
		response.Error(c, constant.CODE_INVALID_PARAMETER, constant.CODE_INVALID_PARAMETER.Msg())
		return
	}

	res, err := cl.s.List(userTicket)

	if err != nil {
		response.Error(c, constant.CODE_FIND_FAILED, constant.CODE_FIND_FAILED.Msg())
		return
	}

	response.Success(c, res)
}
