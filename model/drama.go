package model

import (
	"context"
	"gorm.io/gorm"
	"reflect"
	"strconv"
	"video/global"
	"video/model/doc"
)

// Drama 视频剧集
type Drama struct {
	IDModel
	DramaVideos *[]DramaVideo
	UserModel
	RegionId        int64   `gorm:"type:int not null default 0;comment:区域id;index:idx_region_id"`
	CategoryId      int64   `gorm:"type:int not null default 0;comment:分类id;index:idx_category_id"`
	Name            string  `gorm:"type:varchar(255); not null default '';comment:名称"`
	Introduce       string  `gorm:"type:varchar(2048); not null default '';comment:简介"`
	Cover           string  `gorm:"type:varchar(255); not null default '';comment:横版封面"`
	HorizontalCover string  `gorm:"type:varchar(255); not null default '';comment:纵版封面"`
	Score           float64 `gorm:"type:decimal(1,1);not null default 0;comment:评分"`
	NewEpisode      int64   `gorm:"type:int; not null default 0;comment:剧集"`
	TotalCount      int64   `gorm:"type:int; not null default 0;comment:剧集数"`
	FavoriteCount   int64   `gorm:"type:int; not null default 0;comment:收藏数量"`
	LikeCount       int64   `gorm:"type:int; not null default 0;comment:点赞数量"`
	PlayCount       int64   `gorm:"type:int; not null default 0;comment:播放数量"`
	BarrageCount    int64   `gorm:"type:int; not null default 0;comment:弹幕数量"`
	IsNew           bool    `gorm:"type:tinyint(1); not null default 0;comment:是否最新"`
	IsHot           bool    `gorm:"type:tinyint(1); not null default 0;comment:是否热播"`
	IsEnd           bool    `gorm:"type:tinyint(1); not null default 0;comment:是否完结"`
	Year            int64   `gorm:"type:int; not null default 0;comment:年份"`
	Quarter         int64   `gorm:"type:tinyint(4); not null default 0;comment:季度"`
	Visible
	DateModel
	DeletedModel
}

func (drama *Drama) IsEmpty() bool {
	return reflect.DeepEqual(drama, Drama{})
}

func (drama *Drama) AfterCreate(tx *gorm.DB) error {
	d := drama.ToDoc()
	// 写入es
	_, err := global.ES.
		Index().
		Index(doc.Drama{}.GetIndexName()).
		BodyJson(d).
		Id(strconv.Itoa(int(drama.ID))).
		Do(context.Background())
	return err
}

func (drama *Drama) AfterUpdate(tx *gorm.DB) error {
	//videos := GetDramaVideos(drama.ID)
	d := drama.ToDoc()
	// 更新es. 指定 id 防止重复
	_, err := global.ES.
		Update().
		Index(doc.Drama{}.GetIndexName()).
		Doc(d).
		Id(strconv.Itoa(int(drama.ID))).
		Do(context.Background())

	return err
}

func (drama *Drama) AfterDelete(tx *gorm.DB) error {
	// 删除 es 数据
	_, err := global.ES.
		Delete().
		Index(doc.Drama{}.GetIndexName()).
		Id(strconv.Itoa(int(drama.ID))).
		Do(context.Background())

	return err
}

func (drama *Drama) ToDoc() *doc.UpdateDrama {
	d := &doc.UpdateDrama{
		ID:            drama.ID,
		UserID:        drama.UserID,
		RegionID:      drama.RegionId,
		CategoryID:    drama.CategoryId,
		TotalCount:    drama.TotalCount,
		FavoriteCount: drama.FavoriteCount,
		LikeCount:     drama.LikeCount,
		PlayCount:     drama.PlayCount,
		BarrageCount:  drama.BarrageCount,
		Year:          drama.Year,
		Quarter:       drama.Quarter,
		IsNew:         drama.IsNew,
		IsHot:         drama.IsHot,
		IsEnd:         drama.IsEnd,
		IsVisible:     drama.IsVisible,
		Score:         drama.Score,
		Name:          drama.Name,
		Introduce:     drama.Introduce,
	}
	return d
}
