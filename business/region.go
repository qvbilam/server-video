package business

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"video/global"
	"video/model"
)

type Region struct {
	Id        int64
	Name      string
	Icon      string
	IsVisible *bool
}

func (s *Region) Create() (int64, error) {
	entity := model.Region{
		Name: s.Name,
		Icon: s.Icon,
		Visible: model.Visible{
			IsVisible: *s.IsVisible,
		},
	}

	if res := global.DB.Save(&entity); res.RowsAffected == 0 {
		return 0, status.Errorf(codes.Internal, "创建失败")
	}

	return entity.ID, nil

}

func (s *Region) Update() (int64, error) {
	updates := model.Region{}
	if s.Name != "" {
		updates.Name = s.Name
	}

	if s.Icon != "" {
		updates.Icon = s.Icon
	}

	if s.IsVisible != nil {
		updates.IsVisible = *s.IsVisible
	}

	res := global.DB.Where(model.Region{IDModel: model.IDModel{ID: s.Id}}, updates)
	if res.RowsAffected == 0 {
		return 0, status.Errorf(codes.Internal, "更新失败")
	}

	return res.RowsAffected, nil
}

func (s *Region) Delete() (int64, error) {
	res := global.DB.Delete(model.Region{}, s.Id)
	if res.RowsAffected == 0 {
		return 0, status.Errorf(codes.Internal, "删除失败")
	}

	return res.RowsAffected, nil
}

func (s *Region) List() (*[]model.Region, error) {
	var entity []model.Region
	condition := model.Region{}
	if s.IsVisible != nil {
		condition.Visible.IsVisible = *s.IsVisible
	}
	if s.Name != "" {
		condition.Name = s.Name
	}
	global.DB.Where(condition).Find(&entity)
	return &entity, nil
}

func (s *Region) Exists() (bool, error) {
	entity := model.Region{}

	condition := model.Region{}

	if s.Id > 0 {
		condition.IDModel.ID = s.Id
	}

	if s.IsVisible != nil {
		condition.Visible.IsVisible = *s.IsVisible
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
