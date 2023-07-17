package handle

import (
	"github.com/Madou-Shinni/gin-quickstart/internal/domain"
	"github.com/Madou-Shinni/gin-quickstart/internal/service"
	"github.com/Madou-Shinni/gin-quickstart/pkg/constant"
	"github.com/Madou-Shinni/gin-quickstart/pkg/request"
	"github.com/Madou-Shinni/gin-quickstart/pkg/response"
	"github.com/gin-gonic/gin"
)

type TicketHandle struct {
	s *service.TicketService
}

func NewTicketHandle() *TicketHandle {
	return &TicketHandle{s: service.NewTicketService()}
}

// Add 创建Ticket
// @Tags     Ticket
// @Summary  创建Ticket
// @accept   application/json
// @Produce  application/json
// @Param    data body     domain.Ticket true "创建Ticket"
// @Success  200  {string} string            "{"code":200,"msg":"","data":{}"}"
// @Router   /ticket [post]
func (cl *TicketHandle) Add(c *gin.Context) {
	var ticket domain.Ticket
	if err := c.ShouldBindJSON(&ticket); err != nil {
		response.Error(c, constant.CODE_INVALID_PARAMETER, constant.CODE_INVALID_PARAMETER.Msg())
		return
	}

	if err := cl.s.Add(ticket); err != nil {
		response.Error(c, constant.CODE_ADD_FAILED, constant.CODE_ADD_FAILED.Msg())
		return
	}

	response.Success(c)
}

// Delete 删除Ticket
// @Tags     Ticket
// @Summary  删除Ticket
// @accept   application/json
// @Produce  application/json
// @Param    data body     domain.Ticket true "删除Ticket"
// @Success  200  {string} string            "{"code":200,"msg":"","data":{}"}"
// @Router   /ticket [delete]
func (cl *TicketHandle) Delete(c *gin.Context) {
	var ticket domain.Ticket
	if err := c.ShouldBindJSON(&ticket); err != nil {
		response.Error(c, constant.CODE_INVALID_PARAMETER, constant.CODE_INVALID_PARAMETER.Msg())
		return
	}

	if err := cl.s.Delete(ticket); err != nil {
		response.Error(c, constant.CODE_DELETE_FAILED, constant.CODE_DELETE_FAILED.Msg())
		return
	}

	response.Success(c)
}

func (cl *TicketHandle) AddRollBack(c *gin.Context) {
	var ticket domain.Ticket
	if err := c.ShouldBindJSON(&ticket); err != nil {
		// 409回滚
		response.Error(c, 409, constant.CODE_INVALID_PARAMETER.Msg())
		return
	}

	if err := cl.s.Delete(ticket); err != nil {
		response.Error(c, 409, constant.CODE_DELETE_FAILED.Msg())
		return
	}

	response.Success(c)
}

// DeleteByIds 批量删除Ticket
// @Tags     Ticket
// @Summary  批量删除Ticket
// @accept   application/json
// @Produce  application/json
// @Param    data body     request.Ids true "批量删除Ticket"
// @Success  200  {string} string            "{"code":200,"msg":"","data":{}"}"
// @Router   /ticket/delete-batch [delete]
func (cl *TicketHandle) DeleteByIds(c *gin.Context) {
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

// Update 修改Ticket
// @Tags     Ticket
// @Summary  修改Ticket
// @accept   application/json
// @Produce  application/json
// @Param    data body     domain.Ticket true "修改Ticket"
// @Success  200  {string} string            "{"code":200,"msg":"","data":{}"}"
// @Router   /ticket [put]
func (cl *TicketHandle) Update(c *gin.Context) {
	var ticket map[string]interface{}
	if err := c.ShouldBindJSON(&ticket); err != nil {
		response.Error(c, constant.CODE_INVALID_PARAMETER, constant.CODE_INVALID_PARAMETER.Msg())
		return
	}

	if err := cl.s.Update(ticket); err != nil {
		response.Error(c, constant.CODE_UPDATE_FAILED, constant.CODE_UPDATE_FAILED.Msg())
		return
	}

	response.Success(c)
}

// Find 查询Ticket
// @Tags     Ticket
// @Summary  查询Ticket
// @accept   application/json
// @Produce  application/json
// @Param    data query     domain.Ticket true "查询Ticket"
// @Success  200  {string} string            "{"code":200,"msg":"查询成功","data":{}"}"
// @Router   /ticket [get]
func (cl *TicketHandle) Find(c *gin.Context) {
	var ticket domain.Ticket
	if err := c.ShouldBindQuery(&ticket); err != nil {
		response.Error(c, constant.CODE_INVALID_PARAMETER, constant.CODE_INVALID_PARAMETER.Msg())
		return
	}

	res, err := cl.s.Find(ticket)

	if err != nil {
		response.Error(c, constant.CODE_FIND_FAILED, constant.CODE_FIND_FAILED.Msg())
		return
	}

	response.Success(c, res)
}

// List 查询Ticket列表
// @Tags     Ticket
// @Summary  查询Ticket列表
// @accept   application/json
// @Produce  application/json
// @Param    data query     domain.Ticket true "查询Ticket列表"
// @Success  200  {string} string            "{"code":200,"msg":"查询成功","data":{}"}"
// @Router   /ticket/list [get]
func (cl *TicketHandle) List(c *gin.Context) {
	var ticket domain.PageTicketSearch
	if err := c.ShouldBindQuery(&ticket); err != nil {
		response.Error(c, constant.CODE_INVALID_PARAMETER, constant.CODE_INVALID_PARAMETER.Msg())
		return
	}

	res, err := cl.s.List(ticket)

	if err != nil {
		response.Error(c, constant.CODE_FIND_FAILED, constant.CODE_FIND_FAILED.Msg())
		return
	}

	response.Success(c, res)
}
