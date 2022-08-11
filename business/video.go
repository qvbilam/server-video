package business

import (
	"github.com/olivere/elastic/v7"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"video/global"
	"video/model"
)

type Video struct {
	Id             int64
	UserId         int64
	RegionId       int64
	CategoryId     int64
	Name           string
	Introduction   string
	Icon           string
	HorizontalIcon string
	TotalCount     int64
	Keyword        string
	// 以下字段只允许通过剧集修改
	Count         int64
	Score         float64
	FavoriteCount int64
	LikeCount     int64
	PlayCount     int64
	BarrageCount  int64
	IsRecommend   bool
	IsNew         bool
	IsHot         bool
	IsEnd         bool
	IsVisible     bool
}

func (s *Video) Create() (int64, error) {
	// 开启事物
	tx := global.DB.Begin()

	// 视频实例
	entity := model.Video{
		UserModel: model.UserModel{
			UserID: s.UserId,
		},
		RegionId:       s.RegionId,
		CategoryId:     s.CategoryId,
		Name:           s.Name,
		Introduction:   s.Introduction,
		Icon:           s.Icon,
		HorizontalIcon: s.HorizontalIcon,
		TotalCount:     s.TotalCount,
		Visible:        model.Visible{},
	}

	cvRes := tx.Create(&entity) // 同时会在model.After 写入es
	if cvRes.Error != nil {
		tx.Rollback()
		return 0, status.Errorf(codes.Internal, "创建视频异常: %s", cvRes.Error)
	}

	if cvRes.RowsAffected == 0 {
		tx.Rollback()
		return 0, status.Errorf(codes.Internal, "创建视频失败: %s", cvRes.Error)
	}

	tx.Commit()
	return entity.ID, nil
}

func (s *Video) Update() (int64, error) {
	tx := global.DB.Begin()
	videoEntity := model.Video{}
	if res := global.DB.First(&videoEntity, s.Id); res.RowsAffected == 0 {
		tx.Rollback()
		return 0, status.Errorf(codes.NotFound, "视频不存在")
	}
	videoEntity.RegionId = s.RegionId
	videoEntity.CategoryId = s.CategoryId
	videoEntity.Name = s.Name
	videoEntity.Introduction = s.Introduction
	videoEntity.Icon = s.Icon
	videoEntity.HorizontalIcon = s.HorizontalIcon
	videoEntity.Score = s.Score
	videoEntity.TotalCount = s.TotalCount

	res := global.DB.Save(&videoEntity)
	if res.Error != nil {
		tx.Rollback()
		return 0, status.Errorf(codes.Internal, "修改异常: %s", res.Error)
	}

	if res.RowsAffected == 0 {
		tx.Rollback()
		return 0, status.Errorf(codes.Internal, "更新失败")
	}

	tx.Commit()
	return res.RowsAffected, nil
}

func (s *Video) Delete() (int64, error) {
	tx := global.DB.Begin()

	// 注意: 删除实体进入afterDelete是获取不到ID的. 需要在模型中传入请求的参数id
	res := tx.Where(s.Id).Delete(&model.Video{
		IDModel: model.IDModel{ID: s.Id},
	}, s.Id)
	if res.Error != nil {
		tx.Rollback()
		return 0, status.Errorf(codes.Internal, "删除异常: %s", res.Error)
	}

	if res.RowsAffected == 0 {
		tx.Rollback()
		return 0, status.Errorf(codes.Internal, "删除失败")
	}

	tx.Commit()
	return res.RowsAffected, nil
}

func (s *Video) Detail() (*model.Video, error) {
	videoEntity := model.Video{}
	if res := global.DB.First(&videoEntity, s.Id); res.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "视频不存在")
	}

	return &videoEntity, nil
}

func (s *Video) List() {
	// match bool 复合查询
	q := elastic.NewBoolQuery()

	if s.Keyword != "" { // 搜索 名称, 简介
		q = q.Must(elastic.NewMultiMatchQuery(s.Keyword, "name", "introduction"))
	}
	if s.UserId > 0 { // 搜索用户
		q = q.Filter(elastic.NewTermQuery("user_id", s.UserId))
	}
	if s.RegionId > 0 { // 搜索
		q = q.Filter(elastic.NewTermQuery("region_id", s.RegionId))
	}
	if s.CategoryId > 0 { // 搜索分类
		q = q.Filter(elastic.NewTermQuery("category_id", s.CategoryId))
	}
	if s.IsHot { // 搜索热度
		q = q.Filter(elastic.NewTermQuery("is_hot", s.IsHot))
	}
	if s.IsNew { // 搜索新品
		q = q.Filter(elastic.NewTermQuery("is_new", s.IsNew))
	}
	if s.IsEnd { // 搜索是否完结
		q = q.Filter(elastic.NewTermQuery("is_end", s.IsNew))
	}
	if s.IsVisible { // 搜索展示状态
		q = q.Filter(elastic.NewTermQuery("is_visible", s.IsNew))
	}
}
