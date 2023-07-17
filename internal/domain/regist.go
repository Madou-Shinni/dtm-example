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
