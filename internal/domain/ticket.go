package domain

import (
	"github.com/Madou-Shinni/gin-quickstart/pkg/model"
	"github.com/Madou-Shinni/gin-quickstart/pkg/request"
)

type Ticket struct {
	model.Model
}

type PageTicketSearch struct {
	Ticket
	request.PageSearch
}

func (Ticket) TableName() string {
	return "ticket"
}