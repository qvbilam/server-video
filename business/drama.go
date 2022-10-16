package business

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"video/global"
	"video/model"
)

type DramaListResponse struct {
	Total  int64
	Dramas *[]model.Drama
}

type DramaBusiness struct {
	Id             int64
	UserId         int64
	CategoryId     int64
	Name           string
	Introduce      string
	Icon           string
	HorizontalIcon string
	TotalCount     int64
	Keyword        string
	Type           string

	// 以下字段只允许通过剧集修改
	Score         float64
	FavoriteCount int64
	LikeCount     int64
	PlayCount     int64
	BarrageCount  int64
	IsRecommend   *bool
	IsNew         *bool
	IsHot         *bool
	IsVisible     *bool

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

func (b *DramaBusiness) Update() (int64, error) {
	entity := model.Drama{}
	condition := model.Drama{}
	if b.Id != 0 {
		condition.ID = b.Id
	}
	if condition.IsEmpty() {
		return 0, status.Errorf(codes.InvalidArgument, "参数错误")
	}
	if res := global.DB.Clauses(clause.Locking{Strength: "UPDATE"}).Where(condition).First(&entity); res.RowsAffected == 0 {
		return 0, status.Errorf(codes.NotFound, "剧集不存在")
	}
	if b.Name != "" {
		entity.Name = b.Name
	}
	if b.Introduce != "" {
		entity.Introduce = b.Introduce
	}
	if b.IsNew != nil {
		entity.IsNew = *b.IsNew
	}
	if b.IsHot != nil {
		entity.IsHot = *b.IsHot
	}

	res := global.DB.Save(&entity)
	return res.RowsAffected, nil
}

func (b *DramaBusiness) List() (*DramaListResponse, error) {
	var drams []model.Drama
	global.DB.
		Where(model.Video{IDModel: model.IDModel{ID: 2}}).
		Preload("DramaVideos", func(db *gorm.DB) *gorm.DB {
			return db.Preload("Video", func(db *gorm.DB) *gorm.DB {
				return db.Select("id", "file_id", "name", "introduce")
			})
		}).
		Find(&drams)

	for _, e := range drams {
		for _, vs := range e.DramaVideos {
			fmt.Printf("id: %d, name: %s, episode: %d\n", vs.Video.ID, vs.Video.Name, vs.Episode)
		}
	}

	return nil, nil
	q := b.GetESQuery()
	result, err := global.ES.
		Search().
		Index(model.VideoES{}.GetIndexName()).
		Query(q).
		SortWithInfo(b.GetESSortInfo()).
		From(int(b.Page)).
		Size(int(b.PerPage)).
		Do(context.Background())

	if err != nil {
		return nil, err
	}

	fmt.Printf("%+v\n", result)

	return nil, nil
}

func (b *DramaBusiness) GetESQuery() *elastic.BoolQuery {
	// match bool 复合查询
	q := elastic.NewBoolQuery()

	if b.Keyword != "" { // 搜索 名称, 简介
		q = q.Must(elastic.NewMultiMatchQuery(b.Keyword, "name", "introduction"))
	}

	if b.Type != "" { // 搜索类型
		q = q.Filter(elastic.NewTermQuery("type", b.Type))
	}

	if b.UserId > 0 { // 搜索用户
		q = q.Filter(elastic.NewTermQuery("user_id", b.UserId))
	}

	if b.IsHot != nil { // 搜索热度
		q = q.Filter(elastic.NewTermQuery("is_hot", b.IsHot))
	}
	if b.IsNew != nil { // 搜索新品
		q = q.Filter(elastic.NewTermQuery("is_new", b.IsNew))
	}
	if b.IsVisible != nil { // 搜索展示状态
		q = q.Filter(elastic.NewTermQuery("is_visible", b.IsNew))
	}

	// 范围查询
	if b.FavoriteCountMin > 0 {
		q = q.Filter(elastic.NewRangeQuery("favorite_count").Gte(b.FavoriteCountMin))
	}

	if b.FavoriteCountMax > 0 {
		q = q.Filter(elastic.NewRangeQuery("favorite_count").Lte(b.FavoriteCountMax))
	}

	if b.LikeCountMin > 0 {
		q = q.Filter(elastic.NewRangeQuery("like_count").Gte(b.LikeCountMin))
	}

	if b.LikeCountMax > 0 {
		q = q.Filter(elastic.NewRangeQuery("like_count").Lte(b.LikeCountMax))
	}

	if b.PlayCountMin > 0 {
		q = q.Filter(elastic.NewRangeQuery("play_count").Gte(b.PlayCountMin))
	}

	if b.PlayCountMax > 0 {
		q = q.Filter(elastic.NewRangeQuery("play_count").Lte(b.PlayCountMax))
	}

	if b.BarrageCountMin > 0 {
		q = q.Filter(elastic.NewRangeQuery("barrage_count").Gte(b.BarrageCountMin))
	}

	if b.BarrageCountMax > 0 {
		q = q.Filter(elastic.NewRangeQuery("barrage_count").Lte(b.BarrageCountMax))
	}

	return q
}

func (b *DramaBusiness) GetESSortInfo() elastic.SortInfo {
	sort := elastic.SortInfo{
		Field:     "play_count",
		Ascending: false,
	}

	if b.Sort != "" {
		if string(b.Sort[0]) == "-" {
			sort.Field = b.Sort[0:]
		} else {
			sort.Field = b.Sort
			sort.Ascending = true
		}
	}

	return sort
}
