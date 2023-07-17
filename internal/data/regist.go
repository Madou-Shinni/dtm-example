package data

import (
    "errors"
    "fmt"
	"github.com/Madou-Shinni/gin-quickstart/internal/domain"
	"github.com/Madou-Shinni/gin-quickstart/pkg/global"
	"github.com/Madou-Shinni/gin-quickstart/pkg/request"
	"github.com/Madou-Shinni/gin-quickstart/pkg/tools/pagelimit"
)

type RegistRepo struct {
}

func (s *RegistRepo) Create(regist domain.Regist) error {
	return global.DB.Create(&regist).Error
}

func (s *RegistRepo) Delete(regist domain.Regist) error {
	return global.DB.Delete(&regist).Error
}

func (s *RegistRepo) DeleteByIds(ids request.Ids) error {
	return global.DB.Delete(&[]domain.Regist{}, ids.Ids).Error
}

func (s *RegistRepo) Update(regist map[string]interface{}) error {
    var columns []string
	for key := range regist {
		columns = append(columns, key)
	}
	if _,ok := regist["id"];!ok {
        // 不存在id
        return errors.New(fmt.Sprintf("missing %s.id","regist"))
    }
	model := domain.Regist{}
	model.ID = uint(regist["id"].(float64))
	return global.DB.Model(&model).Select(columns).Updates(&regist).Error
}

func (s *RegistRepo) Find(regist domain.Regist) (domain.Regist, error) {
	db := global.DB.Model(&domain.Regist{})
	// TODO：条件过滤

	res := db.First(&regist)

	return regist, res.Error
}

func (s *RegistRepo) List(page domain.PageRegistSearch) ([]domain.Regist, error) {
	var (
		registList []domain.Regist
		err      error
	)
	// db
	db := global.DB.Model(&domain.Regist{})
	// page
	offset, limit := pagelimit.OffsetLimit(page.PageNum, page.PageSize)

	// TODO：条件过滤

	err = db.Offset(offset).Limit(limit).Find(&registList).Error

	return registList, err
}

func (s *RegistRepo) Count() (int64, error) {
	var (
		count int64
		err   error
	)

	err = global.DB.Model(&domain.Regist{}).Count(&count).Error

	return count, err
}
