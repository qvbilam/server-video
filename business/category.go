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
	IsVisible bool
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
