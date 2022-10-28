package business

import (
	"context"
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
	RegionId       int64
	Name           string
	Introduce      string
	Icon           string
	HorizontalIcon string
	TotalCount     int64
	Keyword        string
	Type           string

	// 以下字段只允许通过剧集修改
	Score         float64
	FavoriteCount *int64
	LikeCount     *int64
	PlayCount     *int64
	BarrageCount  *int64
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

func (b *DramaBusiness) Create() (*model.Drama, error) {
	entity := &model.Drama{
		UserModel:      model.UserModel{UserID: b.UserId},
		Type:           b.Type,
		CategoryId:     b.CategoryId,
		Name:           b.Name,
		Introduce:      b.Introduce,
		Icon:           b.Icon,
		HorizontalIcon: b.HorizontalIcon,
		IsNew:          *b.IsNew,
		IsHot:          *b.IsHot,
		IsEnd:          *b.IsNew,
	}
	if b.IsVisible != nil {
		entity.Visible = model.Visible{IsVisible: *b.IsVisible}
	}

	tx := global.DB.Begin()
	res := tx.Save(entity)
	if res.RowsAffected == 0 || res.Error != nil {
		tx.Rollback()
		return nil, status.Errorf(codes.Internal, "创建失败")
	}

	tx.Commit()
	return entity, nil
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

	tx := global.DB.Begin()

	if res := tx.Clauses(clause.Locking{Strength: "UPDATE"}).Where(condition).First(&entity); res.RowsAffected == 0 {
		tx.Rollback()
		return 0, status.Errorf(codes.NotFound, "剧集不存在")
	}

	if b.RegionId != 0 {
		entity.RegionId = b.RegionId
	}
	if b.CategoryId != 0 {
		entity.CategoryId = b.CategoryId
	}
	if b.Name != "" {
		entity.Name = b.Name
	}
	if b.Introduce != "" {
		entity.Introduce = b.Introduce
	}
	if b.Icon != "" {
		entity.Icon = b.Icon
	}
	if b.HorizontalIcon != "" {
		entity.HorizontalIcon = b.HorizontalIcon
	}
	if b.IsNew != nil {
		entity.IsNew = *b.IsNew
	}
	if b.PlayCount != nil {
		entity.IsHot = *b.IsHot
	}
	if b.IsVisible != nil {
		entity.Visible = model.Visible{IsVisible: *b.IsVisible}
	}

	res := tx.Save(&entity)
	if res.RowsAffected == 0 || res.Error != nil {
		tx.Rollback()
		return 0, status.Errorf(codes.NotFound, "更新失败")
	}
	tx.Commit()
	return res.RowsAffected, nil
}

func (b *DramaBusiness) Delete() (int64, error) {
	if b.Id <= 0 {
		return 0, status.Errorf(codes.InvalidArgument, "参数错误")
	}
	tx := global.DB.Begin()
	res := tx.Where(model.Drama{IDModel: model.IDModel{ID: b.Id}}).Delete(model.Drama{})
	if res.RowsAffected == 0 {
		tx.Rollback()
		return 0, status.Errorf(codes.NotFound, "记录不存在")
	}
	if res.Error != nil {
		tx.Rollback()
		return 0, status.Errorf(codes.Internal, "更新失败")
	}
	tx.Commit()
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
	type highlightMapStruct struct {
		Name      *string
		Introduce *string
	}
	highlightMap := make(map[int64]*highlightMapStruct)
	for _, hit := range result.Hits.Hits {
		idI, _ := strconv.Atoi(hit.Id)
		id := int64(idI)
		dramaIds = append(dramaIds, id)
		highlight := hit.Highlight
		highlightStruct := highlightMapStruct{}
		if highlight != nil {
			if highlight["name"] != nil {
				name := &hit.Highlight["name"][0]
				highlightStruct.Name = name
			}
			if highlight["introduce"] != nil {
				introduce := &hit.Highlight["introduce"][0]
				highlightStruct.Introduce = introduce
			}
		}
		highlightMap[id] = &highlightStruct
	}

	var dramas []model.Drama
	if r := global.DB.Where("id IN ?", dramaIds).Preload("DramaVideos.Video").Find(&dramas); r.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "")
	}
	var entityDramas []model.Drama
	for _, d := range dramas {
		if highlightMap[d.ID] != nil {
			if highlightMap[d.ID].Name != nil {
				d.Name = *highlightMap[d.ID].Name
			}
			if highlightMap[d.ID].Introduce != nil {
				d.Introduce = *highlightMap[d.ID].Introduce
			}
		}
		entityDramas = append(entityDramas, d)
	}
	res.Dramas = &entityDramas
	return res, nil
}

func (b *DramaBusiness) ElasticSearch() (*elastic.SearchResult, error) {
	d := &doc.DramaSearch{
		Keyword:          b.Keyword,
		Type:             b.Type,
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
	if b.UserId != 0 {
		d.UserId = b.UserId
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
