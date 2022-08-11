package business

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"video/global"
	"video/model"
)

type Category struct {
	Id        int64
	IsVisible bool
}

func (s *Category) Exists() (bool, error) {
	var entity model.Category

	condition := model.Category{}

	if s.Id > 0 {
		condition.IDModel.ID = s.Id
	}

	if s.IsVisible == true {
		condition.Visible.IsVisible = s.IsVisible
	}

	res := global.DB.Where(condition).Select("id").First(&entity)

	if res.Error != nil {
		return false, res.Error
	}
	if res.RowsAffected == 0 {
		return false, status.Errorf(codes.NotFound, "分类不存在")
	}

	return true, nil
}
