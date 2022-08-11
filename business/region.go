package business

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"video/global"
	"video/model"
)

type Region struct {
	Id        int64
	IsVisible bool
}

func (s *Region) Exists() (bool, error) {
	entity := model.Region{}

	condition := model.Region{}

	if s.Id > 0 {
		condition.IDModel.ID = s.Id
	}

	if s.IsVisible == true {
		condition.Visible.IsVisible = true
	}

	res := global.DB.Where(condition).Select("id").First(&entity)

	if res.Error != nil {
		return false, res.Error
	}
	if res.RowsAffected == 0 {
		return false, status.Errorf(codes.NotFound, "区域不存在")
	}

	return true, nil
}
