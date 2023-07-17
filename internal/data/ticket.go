package data

import (
    "errors"
    "fmt"
	"github.com/Madou-Shinni/gin-quickstart/internal/domain"
	"github.com/Madou-Shinni/gin-quickstart/pkg/global"
	"github.com/Madou-Shinni/gin-quickstart/pkg/request"
	"github.com/Madou-Shinni/gin-quickstart/pkg/tools/pagelimit"
)

type TicketRepo struct {
}

func (s *TicketRepo) Create(ticket domain.Ticket) error {
	return global.DB.Create(&ticket).Error
}

func (s *TicketRepo) Delete(ticket domain.Ticket) error {
	return global.DB.Delete(&ticket).Error
}

func (s *TicketRepo) DeleteByIds(ids request.Ids) error {
	return global.DB.Delete(&[]domain.Ticket{}, ids.Ids).Error
}

func (s *TicketRepo) Update(ticket map[string]interface{}) error {
    var columns []string
	for key := range ticket {
		columns = append(columns, key)
	}
	if _,ok := ticket["id"];!ok {
        // 不存在id
        return errors.New(fmt.Sprintf("missing %s.id","ticket"))
    }
	model := domain.Ticket{}
	model.ID = uint(ticket["id"].(float64))
	return global.DB.Model(&model).Select(columns).Updates(&ticket).Error
}

func (s *TicketRepo) Find(ticket domain.Ticket) (domain.Ticket, error) {
	db := global.DB.Model(&domain.Ticket{})
	// TODO：条件过滤

	res := db.First(&ticket)

	return ticket, res.Error
}

func (s *TicketRepo) List(page domain.PageTicketSearch) ([]domain.Ticket, error) {
	var (
		ticketList []domain.Ticket
		err      error
	)
	// db
	db := global.DB.Model(&domain.Ticket{})
	// page
	offset, limit := pagelimit.OffsetLimit(page.PageNum, page.PageSize)

	// TODO：条件过滤

	err = db.Offset(offset).Limit(limit).Find(&ticketList).Error

	return ticketList, err
}

func (s *TicketRepo) Count() (int64, error) {
	var (
		count int64
		err   error
	)

	err = global.DB.Model(&domain.Ticket{}).Count(&count).Error

	return count, err
}
