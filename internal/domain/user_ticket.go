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
