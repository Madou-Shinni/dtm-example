package data

import (
	"errors"
	"fmt"
	"github.com/Madou-Shinni/gin-quickstart/internal/domain"
	"github.com/Madou-Shinni/gin-quickstart/pkg/global"
	"github.com/Madou-Shinni/gin-quickstart/pkg/request"
	"github.com/Madou-Shinni/gin-quickstart/pkg/tools/pagelimit"
)

type UserTicketRepo struct {
}

func (s *UserTicketRepo) Create(userTicket domain.UserTicket) (uint, error) {
	err := global.DB.Create(&userTicket).Error
	return userTicket.ID, err
}

func (s *UserTicketRepo) Delete(userTicket domain.UserTicket) error {
	tx := global.DB.Delete(&userTicket)
	tx.Commit()
	fmt.Printf("delete row: %d,ID: %d", tx.RowsAffected, userTicket.ID)
	return tx.Error
}

func (s *UserTicketRepo) DeleteByIds(ids request.Ids) error {
	return global.DB.Delete(&[]domain.UserTicket{}, ids.Ids).Error
}

func (s *UserTicketRepo) Update(userTicket map[string]interface{}) error {
	var columns []string
	for key := range userTicket {
		columns = append(columns, key)
	}
	if _, ok := userTicket["id"]; !ok {
		// 不存在id
		return errors.New(fmt.Sprintf("missing %s.id", "userTicket"))
	}
	model := domain.UserTicket{}
	model.ID = uint(userTicket["id"].(float64))
	return global.DB.Model(&model).Select(columns).Updates(&userTicket).Error
}

func (s *UserTicketRepo) Find(userTicket domain.UserTicket) (domain.UserTicket, error) {
	db := global.DB.Model(&domain.UserTicket{})
	// TODO：条件过滤

	res := db.First(&userTicket)

	return userTicket, res.Error
}

func (s *UserTicketRepo) List(page domain.PageUserTicketSearch) ([]domain.UserTicket, error) {
	var (
		userTicketList []domain.UserTicket
		err            error
	)
	// db
	db := global.DB.Model(&domain.UserTicket{})
	// page
	offset, limit := pagelimit.OffsetLimit(page.PageNum, page.PageSize)

	// TODO：条件过滤

	err = db.Offset(offset).Limit(limit).Find(&userTicketList).Error

	return userTicketList, err
}

func (s *UserTicketRepo) Count() (int64, error) {
	var (
		count int64
		err   error
	)

	err = global.DB.Model(&domain.UserTicket{}).Count(&count).Error

	return count, err
}
