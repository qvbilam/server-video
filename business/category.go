package business

import (
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"video/global"
	"video/model"
)

type Category struct {
	Id        int64
	Name      string
	Icon      string
	Level     int64
	ParentId  *int64
	IsVisible *bool
}

func (s *Category) Create() (int64, error) {
	entity := model.Category{
		ParentId: *s.ParentId,
		Name:     s.Name,
		Icon:     s.Icon,
		Level:    s.Level,
	}
	if res := global.DB.Save(&entity); res.RowsAffected == 0 {
		return 0, status.Errorf(codes.Internal, "创建失败")
	}
	return entity.ID, nil
}

func (s *Category) Update() (int64, error) {
	updates := model.Category{}
	if s.ParentId != nil {
		updates.ParentId = *s.ParentId
	}
	if s.Level > 0 {
		updates.Level = s.Level
	}
	if s.Icon != "" {
		updates.Icon = s.Icon
	}
	if s.Name != "" {
		updates.Name = s.Name
	}
	if s.IsVisible != nil {
		updates.Visible.IsVisible = *s.IsVisible
	}

	res := global.DB.Where(model.Category{
		IDModel: model.IDModel{ID: s.Id},
	}).Updates(updates)

	if res.RowsAffected == 0 {
		return 0, status.Errorf(codes.Internal, "修改失败")
	}

	return res.RowsAffected, nil
}

func (s *Category) Delete() (int64, error) {
	res := global.DB.Delete(&model.Category{}, s.Id)
	if res.RowsAffected == 0 {
		return 0, status.Errorf(codes.Internal, "删除失败")
	}
	return res.RowsAffected, nil
}

func (s *Category) List() (*[]model.Category, error) {
	var entity []model.Category

	condition := model.Category{}
	if s.Level > 0 {
		condition.Level = s.Level
	}
	if s.ParentId != nil {
		condition.ParentId = *s.ParentId
	}
	if s.IsVisible != nil {
		condition.IsVisible = *s.IsVisible
	}

	if res := global.DB.Where(condition).Find(&entity); res.Error != nil {
		return nil, res.Error
	}
	return &entity, nil
}

// GetMultistageCategory 获取多级分类
func (s *Category) GetMultistageCategory() ([]interface{}, error) {
	entity := model.Category{}
	if res := global.DB.First(&entity, s.Id); res.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "分类不存在")
	}

	var subQuery string
	if entity.Level == 1 {
		subQuery = fmt.Sprintf("SELECT id FROM category WHERE parent_id IN (SELECT id FROM category WHERE parent_id = %d)", s.Id)
	} else if entity.Level == 2 { // 二级分类
		subQuery = fmt.Sprintf("SELECT id FROM category WHERE parent_id = %d", s.Id)
	} else { // 三级分类
		subQuery = fmt.Sprintf("SELECT id FROM category WHERE id = %d", s.Id)
	}

	var categoryIds []interface{}
	global.DB.Model(model.Category{}).Raw(subQuery).Scan(&categoryIds)
	for _, cId := range categoryIds {
		categoryIds = append(categoryIds, cId)
	}

	return categoryIds, nil
}

func (s *Category) Exists() (bool, error) {
	var entity model.Category

	condition := model.Category{}

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
		return false, status.Errorf(codes.NotFound, "分类不存在")
	}

	return true, nil
}
