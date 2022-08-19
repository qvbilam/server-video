package model

import (
	"context"
	"gorm.io/gorm"
	"strconv"
	"video/global"
)

// Region 区域
type Region struct {
	IDModel
	Name string `gorm:"type:varchar(255); not null default '';comment:名称"`
	Icon string `gorm:"type:varchar(255); not null default '';comment:横版图标"`
	Visible
	DateModel
	DeletedModel
}

// Category 分类
type Category struct {
	IDModel
	ParentId int64  `gorm:"type:int not null default 0;comment:父级id;index:idx_parent_id"`
	Name     string `gorm:"type:varchar(255); not null default '';comment:名称"`
	Icon     string `gorm:"type:varchar(255); not null default '';comment:横版图标"`
	Level    int64  `gorm:"type:int(11);not null;default:0;comment:分类等级"`
	Visible
	DateModel
	DeletedModel
}

// Style 风格
type Style struct {
	IDModel
	CategoryId int64  `gorm:"type:int not null default 0;comment:分类id;index:idx_category_id"`
	Name       string `gorm:"type:varchar(255); not null default '';comment:名称"`
	Icon       string `gorm:"type:varchar(255); not null default '';comment:横版图标"`
	Visible
	DateModel
	DeletedModel
}

// StyleVideo 风格视频关联
type StyleVideo struct {
	IDModel
	StyleId int64 `gorm:"type:int not null default 0;comment:风格id;index:idx_style_video"`
	VideoId int64 `gorm:"type:int not null default 0;comment:视频id;index:idx_style_video"`
	DateModel
}

// Video 视频
type Video struct {
	IDModel
	UserModel
	RegionId       int64   `gorm:"type:int not null default 0;comment:区域id;index:idx_region_id"`
	CategoryId     int64   `gorm:"type:int not null default 0;comment:分类id;index:idx_category_id"`
	Name           string  `gorm:"type:varchar(255); not null default '';comment:名称"`
	Introduction   string  `gorm:"type:varchar(2048); not null default '';comment:简介"`
	Icon           string  `gorm:"type:varchar(255); not null default '';comment:横版图标"`
	HorizontalIcon string  `gorm:"type:varchar(255); not null default '';comment:纵版图标"`
	Score          float64 `gorm:"type:decimal(1,1);not null default 0;comment:评分"`
	Count          int64   `gorm:"not null default 0;comment:当前集数"`
	TotalCount     int64   `gorm:"not null default 0;comment:总集数"`
	FavoriteCount  int64   `gorm:"not null default 0;comment:收藏数量"`
	LikeCount      int64   `gorm:"not null default 0;comment:点赞数量"`
	PlayCount      int64   `gorm:"not null default 0;comment:播放数量"`
	BarrageCount   int64   `gorm:"not null default 0;comment:弹幕数量"`
	IsRecommend    bool    `gorm:"not null default 0;comment:是否推荐"`
	IsNew          bool    `gorm:"not null default 0;comment:是否最新"`
	IsHot          bool    `gorm:"not null default 0;comment:是否热播"`
	IsEnd          bool    `gorm:"not null default 0;comment:是否完结"`
	Visible
	DateModel
	DeletedModel
}

// Episodes 视频剧集
type Episodes struct {
	IDModel
	VideoId       int64   `gorm:"type:int not null default 0;comment:区域id;index:idx_video_id"`
	AliCloudId    string  `gorm:"type:varchar(255); not null default '';comment:阿里视频id"`
	Name          string  `gorm:"type:varchar(255); not null default '';comment:名称"`
	Introduction  string  `gorm:"type:varchar(2048); not null default '';comment:简介"`
	Icon          string  `gorm:"type:varchar(255); not null default '';comment:横版图标"`
	Score         float64 `gorm:"type:decimal(1,1);not null default 0;comment:评分"`
	Url           string  `gorm:"type:varchar(255); not null default '';comment:播放地址"`
	Number        int64   `gorm:"not null default 0;comment:集数编号"`
	FavoriteCount int64   `gorm:"not null default 0;comment:收藏数量"`
	LikeCount     int64   `gorm:"not null default 0;comment:点赞数量"`
	PlayCount     int64   `gorm:"not null default 0;comment:播放数量"`
	BarrageCount  int64   `gorm:"not null default 0;comment:弹幕数量"`
	IsNew         bool    `gorm:"not null default 0;comment:是否最新"`
	IsHot         bool    `gorm:"not null default 0;comment:是否热播"`
	Visible
	DateModel
	DeletedModel
}

// Barrage 视频弹幕
type Barrage struct {
	IDModel
	UserModel
	VeId    int64  `gorm:"type:int not null default 0;comment:剧集id;index:idx_ve_second"`
	Second  int64  `gorm:"not null default 0;comment:视频当前时间;index:idx_ve_second"`
	Content string `gorm:"type:varchar(255); not null default '';comment:内容"`
	Visible
	DateModel
}

func (video *Video) AfterCreate(tx *gorm.DB) error {
	esModel := videoModelToEsIndex(video)
	// 写入es
	_, err := global.ES.
		Index().
		Index(VideoES{}.GetIndexName()).
		BodyJson(esModel).
		Id(strconv.Itoa(int(video.ID))).
		Do(context.Background())
	return err
}

func (video *Video) AfterUpdate(tx *gorm.DB) error {
	esModel := videoModelToEsIndex(video)
	// 更新es. 指定 id 防止重复
	_, err := global.ES.
		Update().
		Index(VideoES{}.GetIndexName()).
		Doc(esModel).
		Id(strconv.Itoa(int(video.ID))).
		Do(context.Background())

	return err
}

func (video *Video) AfterDelete(tx *gorm.DB) error {
	// 删除 es 数据
	_, err := global.ES.
		Delete().
		Index(VideoES{}.GetIndexName()).
		Id(strconv.Itoa(int(video.ID))).
		Do(context.Background())

	return err
}

func videoModelToEsIndex(video *Video) *VideoES {
	return &VideoES{
		ID:            video.ID,
		UserID:        video.UserID,
		RegionID:      video.RegionId,
		CategoryID:    video.CategoryId,
		Count:         video.Count,
		TotalCount:    video.TotalCount,
		FavoriteCount: video.FavoriteCount,
		LikeCount:     video.LikeCount,
		PlayCount:     video.PlayCount,
		BarrageCount:  video.BarrageCount,
		IsRecommend:   video.IsRecommend,
		IsNew:         video.IsNew,
		IsHot:         video.IsHot,
		IsEnd:         video.IsEnd,
		IsVisible:     video.IsVisible,
		Score:         video.Score,
		Name:          video.Name,
		Introduce:     video.Introduction,
	}
}
