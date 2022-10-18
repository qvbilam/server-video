package business

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm/clause"
	"strconv"
	"video/global"
	"video/model"
	"video/model/doc"
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
	if b.PerPage <= 0 {
		b.PerPage = 10
	}

	result, err := b.ElasticSearch()

	if err != nil {
		return nil, err
	}

	res := &DramaListResponse{}
	res.Total = result.TotalHits()

	var dramaIds []int64
	for _, hit := range result.Hits.Hits {
		dramaId, _ := strconv.Atoi(hit.Id)
		dramaIds = append(dramaIds, int64(dramaId))
		fmt.Println("================")
		fmt.Println(hit.Highlight)
	}
	// todo 返回数据库查询， 替换高亮字段

	return nil, nil
}

func (b *DramaBusiness) ElasticSearch() (*elastic.SearchResult, error) {
	d := &doc.DramaSearch{
		Keyword:          b.Keyword,
		Type:             b.Type,
		UserId:           b.UserId,
		IsHot:            b.IsHot,
		IsNew:            b.IsNew,
		IsVisible:        b.IsVisible,
		FavoriteCountMin: b.FavoriteCountMin,
		FavoriteCountMax: b.FavoriteCountMax,
		LikeCountMin:     b.LikeCountMin,
		LikeCountMax:     b.LikeCountMax,
		PlayCountMin:     b.PlayCountMin,
		PlayCountMax:     b.PlayCountMax,
		BarrageCountMin:  b.BarrageCountMin,
		BarrageCountMax:  b.BarrageCountMax,
	}

	q := d.GetQuery()
	highlightFields := []string{"name", "introduce", "videos.name", "videos.introduce"}
	highlightPreTags := `<font color="#FF0000">`
	highlightPostTags := `"</font>"`
	h := doc.SetHighlight(highlightFields, highlightPreTags, highlightPostTags)
	s := doc.SetSort(b.Sort)

	client := global.ES.Search()
	client.Index(doc.Drama{}.GetIndexName())
	client.Query(q)
	client.From(int(b.Page))
	client.Size(int(b.PerPage))
	client.Highlight(h)
	if s != nil {
		client.SortWithInfo(*s)
	}
	return client.Do(context.Background())
}
