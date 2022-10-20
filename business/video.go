package business

import (
	"context"
	"encoding/json"
	"github.com/olivere/elastic/v7"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"video/global"
	"video/model"
	"video/model/doc"
)

type VideoListResponse struct {
	Total  int64
	Videos []model.Video
}

type VideoBusiness struct {
	Id             int64
	FileId         int64
	UserId         int64
	CategoryId     int64
	Name           string
	Introduce      string
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

func (b *VideoBusiness) Create() (int64, error) {
	// 开启事物
	tx := global.DB.Begin()

	// 视频实例
	entity := model.Video{
		UserModel: model.UserModel{
			UserID: b.UserId,
		},
		CategoryId:     b.CategoryId,
		Name:           b.Name,
		Introduce:      b.Introduce,
		Icon:           b.Icon,
		HorizontalIcon: b.HorizontalIcon,
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

	if b.DramaId != 0 {
		dvBis := DramaVideoBusiness{
			DramaId: b.DramaId,
			VideoId: entity.ID,
			Episode: b.Episode,
		}
		if err := dvBis.Create(tx); err != nil {
			tx.Rollback()
			return 0, err
		}
	}

	tx.Commit()
	return entity.ID, nil
}

func (b *VideoBusiness) Update() (int64, error) {
	tx := global.DB.Begin()
	videoEntity := model.Video{}
	if res := global.DB.First(&videoEntity, b.Id); res.RowsAffected == 0 {
		tx.Rollback()
		return 0, status.Errorf(codes.NotFound, "视频不存在")
	}
	videoEntity.CategoryId = b.CategoryId
	videoEntity.Name = b.Name
	videoEntity.Introduce = b.Introduce
	videoEntity.Icon = b.Icon
	videoEntity.HorizontalIcon = b.HorizontalIcon
	videoEntity.Score = b.Score

	res := global.DB.Save(&videoEntity)
	if res.Error != nil {
		tx.Rollback()
		return 0, status.Errorf(codes.Internal, "修改异常: %s", res.Error)
	}

	if res.RowsAffected == 0 {
		tx.Rollback()
		return 0, status.Errorf(codes.Internal, "更新失败")
	}

	if b.DramaId != 0 && b.Episode != nil {
		dvBis := DramaVideoBusiness{
			DramaId: b.DramaId,
			VideoId: b.Id,
			Episode: b.Episode,
		}
		if _, err := dvBis.Update(tx); err != nil {
			tx.Rollback()
			return 0, err
		}
	}

	tx.Commit()
	return res.RowsAffected, nil
}

func (b *VideoBusiness) Delete() (int64, error) {
	tx := global.DB.Begin()

	// 注意: 删除实体进入afterDelete是获取不到ID的. 需要在模型中传入请求的参数id
	res := tx.Where(b.Id).Delete(&model.Video{
		IDModel: model.IDModel{ID: b.Id},
	}, b.Id)
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

func (b *VideoBusiness) Detail() (*model.Video, error) {
	videoEntity := model.Video{}
	// todo 关联
	if res := global.DB.
		First(&videoEntity, b.Id); res.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "视频不存在")
	}

	return &videoEntity, nil
}

func (b *VideoBusiness) List() (*VideoListResponse, error) {
	switch {
	case b.PerPage <= 0:
		b.PerPage = 10
	case b.PerPage > 1000:
		b.PerPage = 1000
	}
	// 分页数据
	if b.Page == 0 {
		b.Page = 1
	}
	b.Page = (b.Page - 1) * b.PerPage

	// 多级分类
	if b.CategoryId > 0 {
		cs := CategoryBusiness{}
		var err error
		b.CategoryIds, err = cs.GetMultistageCategory()
		if err != nil {
			return nil, err
		}
	}

	result, err := b.ElasticSearch()
	if err != nil {
		return nil, err
	}

	// 获取总数
	total := result.Hits.TotalHits.Value

	// 获取视频 ids
	videoIds := make([]int64, 0)
	for _, video := range result.Hits.Hits {
		d := doc.Video{}
		_ = json.Unmarshal(video.Source, &d)
		videoIds = append(videoIds, d.ID)
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

func (b *VideoBusiness) ElasticSearch() (*elastic.SearchResult, error) {
	d := doc.VideoSearch{
		Keyword:          b.Keyword,
		CategoryIds:      b.CategoryIds,
		UserId:           b.UserId,
		FavoriteCountMin: b.FavoriteCountMin,
		FavoriteCountMax: b.FavoriteCountMax,
		LikeCountMin:     b.LikeCountMin,
		LikeCountMax:     b.LikeCountMax,
		PlayCountMin:     b.PlayCountMin,
		PlayCountMax:     b.PlayCountMax,
		BarrageCountMin:  b.BarrageCountMin,
		BarrageCountMax:  b.BarrageCountMax,
		IsRecommend:      b.IsRecommend,
		IsNew:            b.IsNew,
		IsHot:            b.IsHot,
		IsVisible:        b.IsVisible,
	}
	q := d.GetQuery()
	highlightFields := []string{"name", "introduce"}
	highlightPreTags := `<font color="#FF0000">`
	highlightPostTags := `"</font>"`

	h := doc.SetHighlight(highlightFields, highlightPreTags, highlightPostTags)
	s := doc.SetSort(b.Sort)

	client := global.ES.Search()
	client.Index(doc.Video{}.GetIndexName())
	client.Query(q)
	client.From(int(b.Page))
	client.Size(int(b.PerPage))
	client.Highlight(h)
	if s != nil {
		client.SortWithInfo(*s)
	}
	return client.Do(context.Background())
}
