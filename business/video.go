package business

import (
	"context"
	"encoding/json"
	"github.com/olivere/elastic/v7"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"video/global"
	"video/model"
)

type VideoListResponse struct {
	Total  int64
	Videos []model.Video
}

type Video struct {
	Id             int64
	AliCloudId     string
	UserId         int64
	CategoryId     int64
	Name           string
	Introduction   string
	Icon           string
	HorizontalIcon string
	TotalCount     int64
	Keyword        string
	// 其他表需要字段
	DramaId int64
	Episode *int64
	// 以下字段只允许通过剧集修改
	Score         float64
	FavoriteCount int64
	LikeCount     int64
	PlayCount     int64
	BarrageCount  int64
	IsRecommend   bool
	IsNew         bool
	IsHot         bool
	IsVisible     bool
	// 多层级查找
	CategoryIds []interface{}
	// 范围查找
	FavoriteCountMin int64
	FavoriteCountMax int64
	LikeCountMin     int64
	LikeCountMax     int64
	PlayCountMin     int64
	PlayCountMax     int64
	BarrageCountMin  int64
	BarrageCountMax  int64
	// 分页
	Page    int64
	PerPage int64
	// 排序方式
	Sort string
}

func (s *Video) Create() (int64, error) {
	// 开启事物
	tx := global.DB.Begin()

	// 视频实例
	entity := model.Video{
		UserModel: model.UserModel{
			UserID: s.UserId,
		},
		CategoryId:     s.CategoryId,
		Name:           s.Name,
		Introduction:   s.Introduction,
		Icon:           s.Icon,
		HorizontalIcon: s.HorizontalIcon,
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

	if s.DramaId != 0 {
		dvBis := DramaVideoBusiness{
			DramaId: s.DramaId,
			VideoId: entity.ID,
			Episode: s.Episode,
		}
		if _, err := dvBis.UpdateEpisode(tx); err != nil {
			tx.Rollback()
			return 0, err
		}
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
	videoEntity.CategoryId = s.CategoryId
	videoEntity.Name = s.Name
	videoEntity.Introduction = s.Introduction
	videoEntity.Icon = s.Icon
	videoEntity.HorizontalIcon = s.HorizontalIcon
	videoEntity.Score = s.Score

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
	// todo 关联
	if res := global.DB.
		First(&videoEntity, s.Id); res.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "视频不存在")
	}

	return &videoEntity, nil
}

func (s *Video) List() (*VideoListResponse, error) {
	switch {
	case s.PerPage <= 0:
		s.PerPage = 10
	case s.PerPage > 1000:
		s.PerPage = 1000
	}
	// 分页数据
	if s.Page == 0 {
		s.Page = 1
	}
	s.Page = (s.Page - 1) * s.PerPage

	// 多级分类
	if s.CategoryId > 0 {
		cs := Category{}
		var err error
		s.CategoryIds, err = cs.GetMultistageCategory()
		if err != nil {
			return nil, err
		}
	}
	// 获取 ES query
	q := s.GetVideosESQuery()

	// 查询
	result, err := global.ES.
		Search().
		Index(model.VideoES{}.GetIndexName()).
		Query(q).
		SortWithInfo(s.GetESVideoSortInfo()).
		From(int(s.Page)).
		Size(int(s.PerPage)).
		Do(context.Background())
	if err != nil {
		return nil, err
	}

	// 获取总数
	total := result.Hits.TotalHits.Value

	// 获取视频 ids
	videoIds := make([]int64, 0)
	for _, video := range result.Hits.Hits {
		videoESModel := model.VideoES{}
		_ = json.Unmarshal(video.Source, &videoESModel)
		videoIds = append(videoIds, videoESModel.ID)
	}

	// 获取视频详细信息 todo 关联分类，区域信息
	var videos []model.Video
	if res := global.DB.
		Model(model.Video{}).
		Find(&videos, videoIds); res.Error != nil {
		return nil, res.Error
	}

	return &VideoListResponse{
		Total:  total,
		Videos: videos,
	}, nil
}

func (s *Video) GetVideosESQuery() *elastic.BoolQuery {
	// match bool 复合查询
	q := elastic.NewBoolQuery()

	if s.Keyword != "" { // 搜索 名称, 简介
		q = q.Must(elastic.NewMultiMatchQuery(s.Keyword, "name", "introduction"))
	}
	if s.UserId > 0 { // 搜索用户
		q = q.Filter(elastic.NewTermQuery("user_id", s.UserId))
	}

	if s.IsHot { // 搜索热度
		q = q.Filter(elastic.NewTermQuery("is_hot", s.IsHot))
	}
	if s.IsNew { // 搜索新品
		q = q.Filter(elastic.NewTermQuery("is_new", s.IsNew))
	}
	if s.IsVisible { // 搜索展示状态
		q = q.Filter(elastic.NewTermQuery("is_visible", s.IsNew))
	}

	// 多级分类查找
	if len(s.CategoryIds) > 0 {
		q = q.Filter(elastic.NewTermsQuery("category_id", s.CategoryIds...))
	}

	// 范围查询
	if s.FavoriteCountMin > 0 {
		q = q.Filter(elastic.NewRangeQuery("favorite_count").Gte(s.FavoriteCountMin))
	}

	if s.FavoriteCountMax > 0 {
		q = q.Filter(elastic.NewRangeQuery("favorite_count").Lte(s.FavoriteCountMax))
	}

	if s.LikeCountMin > 0 {
		q = q.Filter(elastic.NewRangeQuery("like_count").Gte(s.LikeCountMin))
	}

	if s.LikeCountMax > 0 {
		q = q.Filter(elastic.NewRangeQuery("like_count").Lte(s.LikeCountMax))
	}

	if s.PlayCountMin > 0 {
		q = q.Filter(elastic.NewRangeQuery("play_count").Gte(s.PlayCountMin))
	}

	if s.PlayCountMax > 0 {
		q = q.Filter(elastic.NewRangeQuery("play_count").Lte(s.PlayCountMax))
	}

	if s.BarrageCountMin > 0 {
		q = q.Filter(elastic.NewRangeQuery("barrage_count").Gte(s.BarrageCountMin))
	}

	if s.BarrageCountMax > 0 {
		q = q.Filter(elastic.NewRangeQuery("barrage_count").Lte(s.BarrageCountMax))
	}

	return q
}

func (s *Video) GetESVideoSortInfo() elastic.SortInfo {
	sort := elastic.SortInfo{
		Field:     "play_count",
		Ascending: false,
	}

	if s.Sort != "" {
		if string(s.Sort[0]) == "-" {
			sort.Field = s.Sort[0:]
		} else {
			sort.Field = s.Sort
			sort.Ascending = true
		}
	}

	return sort
}
