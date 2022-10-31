package model

import (
	"context"
	"gorm.io/gorm"
	"strconv"
	"video/global"
	"video/model/doc"
)

// Video 视频
type Video struct {
	IDModel
	UserModel
	FileId          int64   `gorm:"type:int(11); not null default 0';comment:视频文件id"`
	CategoryId      int64   `gorm:"type:int not null default 0;comment:分类id;index:idx_category_id"`
	Name            string  `gorm:"type:varchar(255); not null default '';comment:名称"`
	Introduce       string  `gorm:"type:varchar(2048); not null default '';comment:简介"`
	Cover           string  `gorm:"type:varchar(255); not null default '';comment:横版封面"`
	HorizontalCover string  `gorm:"type:varchar(255); not null default '';comment:纵版封面"`
	Score           float64 `gorm:"type:decimal(1,1);not null default 0;comment:评分"`
	FavoriteCount   int64   `gorm:"type:int; not null default 0;comment:收藏数量"`
	LikeCount       int64   `gorm:"type:int; not null default 0;comment:点赞数量"`
	PlayCount       int64   `gorm:"type:int; not null default 0;comment:播放数量"`
	BarrageCount    int64   `gorm:"type:int; not null default 0;comment:弹幕数量"`
	IsRecommend     bool    `gorm:"type:tinyint(1); not null default 0;comment:是否推荐"`
	IsNew           bool    `gorm:"type:tinyint(1); not null default 0;comment:是否最新"`
	IsHot           bool    `gorm:"type:tinyint(1); not null default 0;comment:是否热播"`
	Visible
	DateModel
	DeletedModel
}

func (video *Video) AfterCreate(tx *gorm.DB) error {
	d := video.ToDoc()
	// 写入es
	_, err := global.ES.
		Index().
		Index(doc.Video{}.GetIndexName()).
		BodyJson(d).
		Id(strconv.Itoa(int(video.ID))).
		Do(context.Background())
	return err
}

func (video *Video) AfterUpdate(tx *gorm.DB) error {
	d := video.ToDoc()
	// 更新es. 指定 id 防止重复
	_, err := global.ES.
		Update().
		Index(doc.Video{}.GetIndexName()).
		Doc(d).
		Id(strconv.Itoa(int(video.ID))).
		Do(context.Background())

	return err
}

func (video *Video) AfterDelete(tx *gorm.DB) error {
	// 删除 es 数据
	_, err := global.ES.
		Delete().
		Index(doc.Video{}.GetIndexName()).
		Id(strconv.Itoa(int(video.ID))).
		Do(context.Background())

	return err
}

func (video *Video) ToDoc() *doc.Video {
	return &doc.Video{
		ID:            video.ID,
		UserID:        video.UserID,
		CategoryID:    video.CategoryId,
		FavoriteCount: video.FavoriteCount,
		LikeCount:     video.LikeCount,
		PlayCount:     video.PlayCount,
		BarrageCount:  video.BarrageCount,
		IsRecommend:   video.IsRecommend,
		IsNew:         video.IsNew,
		IsHot:         video.IsHot,
		IsVisible:     video.IsVisible,
		Score:         video.Score,
		Name:          video.Name,
		Introduce:     video.Introduce,
	}
}
