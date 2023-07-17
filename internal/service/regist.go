package service

import (
	"errors"
	"github.com/Madou-Shinni/gin-quickstart/internal/data"
	"github.com/Madou-Shinni/gin-quickstart/internal/domain"
	"github.com/Madou-Shinni/gin-quickstart/pkg/global"
	"github.com/Madou-Shinni/gin-quickstart/pkg/request"
	"github.com/Madou-Shinni/gin-quickstart/pkg/response"
	"github.com/Madou-Shinni/go-logger"
	"go.uber.org/zap"
)

// 定义接口
type RegistRepo interface {
	Create(regist domain.Regist) error
	Delete(regist domain.Regist) error
	Update(regist map[string]interface{}) error
	Find(regist domain.Regist) (domain.Regist, error)
	List(page domain.PageRegistSearch) ([]domain.Regist, error)
	Count() (int64, error)
	DeleteByIds(ids request.Ids) error
}

type RegistService struct {
	repo RegistRepo
}

func NewRegistService() *RegistService {
	return &RegistService{repo: &data.RegistRepo{}}
}

func (s *RegistService) Add(regist domain.Regist) error {
	// 3.持久化入库
	if err := s.repo.Create(regist); err != nil {
		// 4.记录日志
		logger.Error("s.repo.Create(regist)", zap.Error(err), zap.Any("domain.Regist", regist))
		return err
	}

	return errors.New("用户已报名，请勿重复报名！")
}

func (s *RegistService) Delete(regist domain.Regist) error {
	if err := s.repo.Delete(regist); err != nil {
		logger.Error("s.repo.Delete(regist)", zap.Error(err), zap.Any("domain.Regist", regist))
		return err
	}

	return nil
}

func (s *RegistService) Update(regist map[string]interface{}) error {
	if err := s.repo.Update(regist); err != nil {
		logger.Error("s.repo.Update(regist)", zap.Error(err), zap.Any("domain.Regist", regist))
		return err
	}

	return nil
}

func (s *RegistService) Find(regist domain.Regist) (domain.Regist, error) {
	res, err := s.repo.Find(regist)

	if err != nil {
		logger.Error("s.repo.Find(regist)", zap.Error(err), zap.Any("domain.Regist", regist))
		return res, err
	}

	return res, nil
}

func (s *RegistService) List(page domain.PageRegistSearch) (response.PageResponse, error) {
	var (
		pageRes response.PageResponse
	)

	data, err := s.repo.List(page)
	if err != nil {
		logger.Error("s.repo.List(page)", zap.Error(err), zap.Any("domain.PageRegistSearch", page))
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

func (s *RegistService) DeleteByIds(ids request.Ids) error {
	if err := s.repo.DeleteByIds(ids); err != nil {
		logger.Error("s.DeleteByIds(ids)", zap.Error(err), zap.Any("ids request.Ids", ids))
		return err
	}

	return nil
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
