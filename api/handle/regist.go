package handle

import (
	"encoding/json"
	"fmt"
	"github.com/Madou-Shinni/gin-quickstart/internal/domain"
	"github.com/Madou-Shinni/gin-quickstart/internal/service"
	"github.com/Madou-Shinni/gin-quickstart/pkg/constant"
	"github.com/Madou-Shinni/gin-quickstart/pkg/request"
	"github.com/Madou-Shinni/gin-quickstart/pkg/response"
	"github.com/Madou-Shinni/gin-quickstart/pkg/tools/snowflake"
	"github.com/dtm-labs/client/dtmcli"
	"github.com/gin-gonic/gin"
	"regexp"
)

type RegistHandle struct {
	s *service.RegistService
}

func NewRegistHandle() *RegistHandle {
	return &RegistHandle{s: service.NewRegistService()}
}

// Add 创建Regist
// @Tags     Regist
// @Summary  创建Regist
// @accept   application/json
// @Produce  application/json
// @Param    data body     domain.Regist true "创建Regist"
// @Success  200  {string} string            "{"code":200,"msg":"","data":{}"}"
// @Router   /regist [post]
func (cl *RegistHandle) Add(c *gin.Context) {
	var regist domain.Regist
	if err := c.ShouldBindJSON(&regist); err != nil {
		response.Error(c, 409, constant.CODE_INVALID_PARAMETER.Msg())
		return
	}

	if err := cl.s.Add(regist); err != nil {
		response.Error(c, 409, err.Error())
		return
	}

	response.Success(c)
}

func (cl *RegistHandle) AddRollBack(c *gin.Context) {
	var regist domain.Regist
	if err := c.ShouldBindJSON(&regist); err != nil {
		response.Error(c, 409, constant.CODE_INVALID_PARAMETER.Msg())
		return
	}

	if err := cl.s.AddRollback(regist); err != nil {
		response.Error(c, 409, constant.CODE_ADD_FAILED.Msg())
		return
	}

	response.Success(c)
}

// Delete 删除Regist
// @Tags     Regist
// @Summary  删除Regist
// @accept   application/json
// @Produce  application/json
// @Param    data body     domain.Regist true "删除Regist"
// @Success  200  {string} string            "{"code":200,"msg":"","data":{}"}"
// @Router   /regist [delete]
func (cl *RegistHandle) Delete(c *gin.Context) {
	var regist domain.Regist
	if err := c.ShouldBindJSON(&regist); err != nil {
		response.Error(c, constant.CODE_INVALID_PARAMETER, constant.CODE_INVALID_PARAMETER.Msg())
		return
	}

	if err := cl.s.Delete(regist); err != nil {
		response.Error(c, constant.CODE_DELETE_FAILED, constant.CODE_DELETE_FAILED.Msg())
		return
	}

	response.Success(c)
}

// DeleteByIds 批量删除Regist
// @Tags     Regist
// @Summary  批量删除Regist
// @accept   application/json
// @Produce  application/json
// @Param    data body     request.Ids true "批量删除Regist"
// @Success  200  {string} string            "{"code":200,"msg":"","data":{}"}"
// @Router   /regist/delete-batch [delete]
func (cl *RegistHandle) DeleteByIds(c *gin.Context) {
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

// Update 修改Regist
// @Tags     Regist
// @Summary  修改Regist
// @accept   application/json
// @Produce  application/json
// @Param    data body     domain.Regist true "修改Regist"
// @Success  200  {string} string            "{"code":200,"msg":"","data":{}"}"
// @Router   /regist [put]
func (cl *RegistHandle) Update(c *gin.Context) {
	var regist map[string]interface{}
	if err := c.ShouldBindJSON(&regist); err != nil {
		response.Error(c, constant.CODE_INVALID_PARAMETER, constant.CODE_INVALID_PARAMETER.Msg())
		return
	}

	if err := cl.s.Update(regist); err != nil {
		response.Error(c, constant.CODE_UPDATE_FAILED, constant.CODE_UPDATE_FAILED.Msg())
		return
	}

	response.Success(c)
}

type DtmResp struct {
	DtmResult string `json:"dtm_result"`
	Message   string `json:"message"`
}

type RollbackReason struct {
	Code int    `json:"code"`
	Msg  string `json:"message"`
}

// Trans 报名事务
// @Tags     Regist
// @Summary  报名事务
// @accept   application/json
// @Produce  application/json
// @Param    data body     domain.Regist true "修改Regist"
// @Success  200  {string} string            "{"code":200,"msg":"","data":{}"}"
// @Router   /regist/trans [post]
func (cl *RegistHandle) Trans(c *gin.Context) {
	const qsBusi = "http://localhost:8080/"
	dtmServer := "http://localhost:36789/api/dtmsvr"
	// dtm
	req := domain.UserTicket{UserId: 1, TicketId: 1}
	req.ID = uint(snowflake.GenerateID())
	req2 := domain.Regist{UserId: 1, UserTicketId: req.ID}
	saga := dtmcli.NewSaga(dtmServer, dtmcli.MustGenGid(dtmServer)).
		// 购票
		Add(qsBusi+"userTicket", qsBusi+"userTicket/add-rollback", req).
		// 用户报名
		Add(qsBusi+"regist", qsBusi+"regist/regist-rollback", req2)
	saga.WaitResult = true
	err := saga.Submit()
	if err != nil {
		errorMsg := DtmResp{}
		reason := RollbackReason{}
		json.Unmarshal([]byte(err.Error()), &errorMsg)
		re := regexp.MustCompile(`\{.*\}`)
		findString := re.FindString(errorMsg.Message)
		json.Unmarshal([]byte(findString), &reason)
		fmt.Printf("dtm_result: %s, message: %s", errorMsg.DtmResult, errorMsg.Message)
		response.Error(c, constant.RspCode(reason.Code), reason.Msg)
		return
	}

	response.Success(c)
}

// Find 查询Regist
// @Tags     Regist
// @Summary  查询Regist
// @accept   application/json
// @Produce  application/json
// @Param    data query     domain.Regist true "查询Regist"
// @Success  200  {string} string            "{"code":200,"msg":"查询成功","data":{}"}"
// @Router   /regist [get]
func (cl *RegistHandle) Find(c *gin.Context) {
	var regist domain.Regist
	if err := c.ShouldBindQuery(&regist); err != nil {
		response.Error(c, constant.CODE_INVALID_PARAMETER, constant.CODE_INVALID_PARAMETER.Msg())
		return
	}

	res, err := cl.s.Find(regist)

	if err != nil {
		response.Error(c, constant.CODE_FIND_FAILED, constant.CODE_FIND_FAILED.Msg())
		return
	}

	response.Success(c, res)
}

// List 查询Regist列表
// @Tags     Regist
// @Summary  查询Regist列表
// @accept   application/json
// @Produce  application/json
// @Param    data query     domain.Regist true "查询Regist列表"
// @Success  200  {string} string            "{"code":200,"msg":"查询成功","data":{}"}"
// @Router   /regist/list [get]
func (cl *RegistHandle) List(c *gin.Context) {
	var regist domain.PageRegistSearch
	if err := c.ShouldBindQuery(&regist); err != nil {
		response.Error(c, constant.CODE_INVALID_PARAMETER, constant.CODE_INVALID_PARAMETER.Msg())
		return
	}

	res, err := cl.s.List(regist)

	if err != nil {
		response.Error(c, constant.CODE_FIND_FAILED, constant.CODE_FIND_FAILED.Msg())
		return
	}

	response.Success(c, res)
}
