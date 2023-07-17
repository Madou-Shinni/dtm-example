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
type TicketRepo interface {
	Create(ticket domain.Ticket) error
	Delete(ticket domain.Ticket) error
	Update(ticket map[string]interface{}) error
	Find(ticket domain.Ticket) (domain.Ticket, error)
	List(page domain.PageTicketSearch) ([]domain.Ticket, error)
	Count() (int64, error)
	DeleteByIds(ids request.Ids) error
}

type TicketService struct {
	repo TicketRepo
}

func NewTicketService() *TicketService {
	return &TicketService{repo: &data.TicketRepo{}}
}

func (s *TicketService) Add(ticket domain.Ticket) error {
	// 3.持久化入库
	if err := s.repo.Create(ticket); err != nil {
		// 4.记录日志
		logger.Error("s.repo.Create(ticket)", zap.Error(err), zap.Any("domain.Ticket", ticket))
		return err
	}

	return nil
}

func (s *TicketService) Delete(ticket domain.Ticket) error {
	if err := s.repo.Delete(ticket); err != nil {
		logger.Error("s.repo.Delete(ticket)", zap.Error(err), zap.Any("domain.Ticket", ticket))
		return err
	}

	return nil
}

func (s *TicketService) Update(ticket map[string]interface{}) error {
	if err := s.repo.Update(ticket); err != nil {
		logger.Error("s.repo.Update(ticket)", zap.Error(err), zap.Any("domain.Ticket", ticket))
		return err
	}

	return nil
}

func (s *TicketService) Find(ticket domain.Ticket) (domain.Ticket, error) {
	res, err := s.repo.Find(ticket)

	if err != nil {
		logger.Error("s.repo.Find(ticket)", zap.Error(err), zap.Any("domain.Ticket", ticket))
		return res, err
	}

	return res, nil
}

func (s *TicketService) List(page domain.PageTicketSearch) (response.PageResponse, error) {
	var (
		pageRes response.PageResponse
	)

	data, err := s.repo.List(page)
	if err != nil {
		logger.Error("s.repo.List(page)", zap.Error(err), zap.Any("domain.PageTicketSearch", page))
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

func (s *TicketService) DeleteByIds(ids request.Ids) error {
	if err := s.repo.DeleteByIds(ids); err != nil {
		logger.Error("s.DeleteByIds(ids)", zap.Error(err), zap.Any("ids request.Ids", ids))
		return err
	}

	return nil
}
