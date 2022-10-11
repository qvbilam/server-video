package business

import (
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"video/global"
	"video/model"
)

type CategoryBusiness struct {
	Id        int64
	Name      string
	Icon      string
	Level     int64
	ParentId  *int64
	IsVisible *bool
}

func (b *CategoryBusiness) Create() (int64, error) {
	entity := model.Category{
		ParentId: *b.ParentId,
		Name:     b.Name,
		Icon:     b.Icon,
		Level:    b.Level,
	}
	if res := global.DB.Save(&entity); res.RowsAffected == 0 {
		return 0, status.Errorf(codes.Internal, "创建失败")
	}
	return entity.ID, nil
}

func (b *CategoryBusiness) Update() (int64, error) {
	updates := model.Category{}
	if b.ParentId != nil {
		updates.ParentId = *b.ParentId
	}
	if b.Level > 0 {
		updates.Level = b.Level
	}
	if b.Icon != "" {
		updates.Icon = b.Icon
	}
	if b.Name != "" {
		updates.Name = b.Name
	}
	if b.IsVisible != nil {
		updates.Visible.IsVisible = *b.IsVisible
	}

	res := global.DB.Where(model.Category{
		IDModel: model.IDModel{ID: b.Id},
	}).Updates(updates)

	if res.RowsAffected == 0 {
		return 0, status.Errorf(codes.Internal, "修改失败")
	}

	return res.RowsAffected, nil
}

func (b *CategoryBusiness) Delete() (int64, error) {
	res := global.DB.Delete(&model.Category{}, b.Id)
	if res.RowsAffected == 0 {
		return 0, status.Errorf(codes.Internal, "删除失败")
	}
	return res.RowsAffected, nil
}

func (b *CategoryBusiness) List() (*[]model.Category, error) {
	var entity []model.Category

	condition := model.Category{}
	if b.Level > 0 {
		condition.Level = b.Level
	}
	if b.ParentId != nil {
		condition.ParentId = *b.ParentId
	}
	if b.IsVisible != nil {
		condition.IsVisible = *b.IsVisible
	}

	if res := global.DB.Where(condition).Find(&entity); res.Error != nil {
		return nil, res.Error
	}
	return &entity, nil
}

// GetMultistageCategory 获取多级分类
func (b *CategoryBusiness) GetMultistageCategory() ([]interface{}, error) {
	entity := model.Category{}
	if res := global.DB.First(&entity, b.Id); res.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "分类不存在")
	}

	var subQuery string
	if entity.Level == 1 {
		subQuery = fmt.Sprintf("SELECT id FROM category WHERE parent_id IN (SELECT id FROM category WHERE parent_id = %d)", b.Id)
	} else if entity.Level == 2 { // 二级分类
		subQuery = fmt.Sprintf("SELECT id FROM category WHERE parent_id = %d", b.Id)
	} else { // 三级分类
		subQuery = fmt.Sprintf("SELECT id FROM category WHERE id = %d", b.Id)
	}

	var categoryIds []interface{}
	global.DB.Model(model.Category{}).Raw(subQuery).Scan(&categoryIds)
	for _, cId := range categoryIds {
		categoryIds = append(categoryIds, cId)
	}

	return categoryIds, nil
}

func (b *CategoryBusiness) Exists() (bool, error) {
	var entity model.Category

	condition := model.Category{}

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
		return false, status.Errorf(codes.NotFound, "分类不存在")
	}

	return true, nil
}
