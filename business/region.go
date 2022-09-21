package business

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"video/global"
	"video/model"
)

type RegionBusiness struct {
	Id        int64
	Name      string
	Icon      string
	IsVisible *bool
}

func (b *RegionBusiness) Create() (int64, error) {
	entity := model.Region{
		Name: b.Name,
		Icon: b.Icon,
		Visible: model.Visible{
			IsVisible: *b.IsVisible,
		},
	}

	if res := global.DB.Save(&entity); res.RowsAffected == 0 {
		return 0, status.Errorf(codes.Internal, "创建失败")
	}

	return entity.ID, nil

}

func (b *RegionBusiness) Update() (int64, error) {
	updates := model.Region{}
	if b.Name != "" {
		updates.Name = b.Name
	}

	if b.Icon != "" {
		updates.Icon = b.Icon
	}

	if b.IsVisible != nil {
		updates.IsVisible = *b.IsVisible
	}

	res := global.DB.Where(model.Region{IDModel: model.IDModel{ID: b.Id}}, updates)
	if res.RowsAffected == 0 {
		return 0, status.Errorf(codes.Internal, "更新失败")
	}

	return res.RowsAffected, nil
}

func (b *RegionBusiness) Delete() (int64, error) {
	res := global.DB.Delete(model.Region{}, b.Id)
	if res.RowsAffected == 0 {
		return 0, status.Errorf(codes.Internal, "删除失败")
	}

	return res.RowsAffected, nil
}

func (b *RegionBusiness) List() (*[]model.Region, error) {
	var entity []model.Region
	condition := model.Region{}
	if b.IsVisible != nil {
		condition.Visible.IsVisible = *b.IsVisible
	}
	if b.Name != "" {
		condition.Name = b.Name
	}
	global.DB.Where(condition).Find(&entity)
	return &entity, nil
}

func (b *RegionBusiness) Exists() (bool, error) {
	entity := model.Region{}

	condition := model.Region{}

	if b.Id > 0 {
		condition.IDModel.ID = b.Id
	}

	if b.IsVisible != nil {
		condition.Visible.IsVisible = *b.IsVisible
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
