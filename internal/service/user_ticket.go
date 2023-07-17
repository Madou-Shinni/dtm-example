package service

import (
	"github.com/Madou-Shinni/gin-quickstart/internal/data"
	"github.com/Madou-Shinni/gin-quickstart/internal/domain"
	"github.com/Madou-Shinni/gin-quickstart/pkg/request"
	"github.com/Madou-Shinni/gin-quickstart/pkg/response"
	"github.com/Madou-Shinni/go-logger"
	"go.uber.org/zap"
)

// 定义接口
type UserTicketRepo interface {
	Create(userTicket domain.UserTicket) (uint, error)
	Delete(userTicket domain.UserTicket) error
	Update(userTicket map[string]interface{}) error
	Find(userTicket domain.UserTicket) (domain.UserTicket, error)
	List(page domain.PageUserTicketSearch) ([]domain.UserTicket, error)
	Count() (int64, error)
	DeleteByIds(ids request.Ids) error
}

type UserTicketService struct {
	repo UserTicketRepo
}

func NewUserTicketService() *UserTicketService {
	return &UserTicketService{repo: &data.UserTicketRepo{}}
}

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

func (s *UserTicketService) Update(userTicket map[string]interface{}) error {
	if err := s.repo.Update(userTicket); err != nil {
		logger.Error("s.repo.Update(userTicket)", zap.Error(err), zap.Any("domain.UserTicket", userTicket))
		return err
	}

	return nil
}

func (s *UserTicketService) Find(userTicket domain.UserTicket) (domain.UserTicket, error) {
	res, err := s.repo.Find(userTicket)

	if err != nil {
		logger.Error("s.repo.Find(userTicket)", zap.Error(err), zap.Any("domain.UserTicket", userTicket))
		return res, err
	}

	return res, nil
}

func (s *UserTicketService) List(page domain.PageUserTicketSearch) (response.PageResponse, error) {
	var (
		pageRes response.PageResponse
	)

	data, err := s.repo.List(page)
	if err != nil {
		logger.Error("s.repo.List(page)", zap.Error(err), zap.Any("domain.PageUserTicketSearch", page))
		return pageRes, err
	}

	count, err := s.repo.Count()
	if err != nil {
		logger.Error("s.repo.Count()", zap.Error(err))
		return pageRes, err
	}

	pageRes.Data = data
	pageRes.Total = count

	return pageRes, nil
}

func (s *UserTicketService) DeleteByIds(ids request.Ids) error {
	if err := s.repo.DeleteByIds(ids); err != nil {
		logger.Error("s.DeleteByIds(ids)", zap.Error(err), zap.Any("ids request.Ids", ids))
		return err
	}

	return nil
}
