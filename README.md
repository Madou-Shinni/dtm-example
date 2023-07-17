dtm+http
业务介绍：
1.用户购票 2.用户通过购得的票报名
报名和购票结构
购票
package domain

import (
"github.com/Madou-Shinni/gin-quickstart/pkg/model"
"github.com/Madou-Shinni/gin-quickstart/pkg/request"
)

type UserTicket struct {
model.Model
UserId   uint `json:"userId"`
TicketId uint `json:"ticketId"`
}

type PageUserTicketSearch struct {
UserTicket
request.PageSearch
}

func (UserTicket) TableName() string {
return "user_ticket"
}
报名
package domain

import (
"github.com/Madou-Shinni/gin-quickstart/pkg/model"
"github.com/Madou-Shinni/gin-quickstart/pkg/request"
)

type Regist struct {
model.Model
UserId       uint `json:"userId"`
UserTicketId uint `json:"userTicketId"`
}

type PageRegistSearch struct {
Regist
request.PageSearch
}

func (Regist) TableName() string {
return "regist"
}
路由
购票
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
报名
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
路由处理函数
添加购票和回滚接口，这里我们回滚直接删除用户购票的数据
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
添加报名和报名回滚接口
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
业务实现
购票
func (s *UserTicketService) Add(userTicket domain.UserTicket) (uint, error) {
// 3.持久化入库
if id, err := s.repo.Create(userTicket); err != nil {
// 4.记录日志
logger.Error("s.repo.Create(userTicket)", zap.Error(err), zap.Any("domain.UserTicket", userTicket))
return id, err
}

    return 0, nil
}

func (s *UserTicketService) Delete(userTicket domain.UserTicket) error {
if err := s.repo.Delete(userTicket); err != nil {
logger.Error("s.repo.Delete(userTicket)", zap.Error(err), zap.Any("domain.UserTicket", userTicket))
return err
}

	return nil
}
报名
func (s *RegistService) Add(regist domain.Regist) error {
// 3.持久化入库
if err := s.repo.Create(regist); err != nil {
// 4.记录日志
logger.Error("s.repo.Create(regist)", zap.Error(err), zap.Any("domain.Regist", regist))
return err
}

    return errors.New("用户已报名，请勿重复报名！")
}

func (s *RegistService) AddRollback(regist domain.Regist) error {
res := global.DB.Where("user_id = ? AND user_ticket_id = ?", regist.UserId, regist.UserTicketId).Delete(&domain.Regist{})
if res.Error != nil {
return res.Error
}
if res.RowsAffected < 0 {
return errors.New("回滚")
}

	return nil
}
saga事务实现购票接口和报名接口原子性
注意：这里我们在事务开始前生成了用户购票表的id主键，以便于添加报名的事务分支可以使用到我们用户购买的票
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