package model

import (
	"context"
	"gorm.io/gorm"
	"strconv"
	"video/global"
)

type VideoES struct {
	ID            int64 `json:"id"`
	UserID        int64 `json:"user_id"`
	CategoryID    int64 `json:"category_id"`
	FavoriteCount int64 `json:"favorite_count"`
	LikeCount     int64 `json:"like_count"`
	PlayCount     int64 `json:"play_count"`
	BarrageCount  int64 `json:"barrage_count"`

	IsRecommend bool `json:"is_recommend"`
	IsNew       bool `json:"is_new"`
	IsHot       bool `json:"is_hot"`
	IsVisible   bool `json:"isVisible"`

	Score float64 `json:"score"`

	AliCloudId string `json:"ali_cloud_id"`
	Name       string `json:"name"`
	Introduce  string `json:"introduce"`
}

// Video 视频
type Video struct {
	IDModel
	UserModel
	AliCloudId     string  `gorm:"type:varchar(255); not null default '';comment:阿里视频id"`
	CategoryId     int64   `gorm:"type:int not null default 0;comment:分类id;index:idx_category_id"`
	Name           string  `gorm:"type:varchar(255); not null default '';comment:名称"`
	Introduction   string  `gorm:"type:varchar(2048); not null default '';comment:简介"`
	Icon           string  `gorm:"type:varchar(255); not null default '';comment:横版图标"`
	HorizontalIcon string  `gorm:"type:varchar(255); not null default '';comment:纵版图标"`
	Score          float64 `gorm:"type:decimal(1,1);not null default 0;comment:评分"`
	FavoriteCount  int64   `gorm:"type:int; not null default 0;comment:收藏数量"`
	LikeCount      int64   `gorm:"type:int; not null default 0;comment:点赞数量"`
	PlayCount      int64   `gorm:"type:int; not null default 0;comment:播放数量"`
	BarrageCount   int64   `gorm:"type:int; not null default 0;comment:弹幕数量"`
	IsRecommend    bool    `gorm:"type:tinyint(1); not null default 0;comment:是否推荐"`
	IsNew          bool    `gorm:"type:tinyint(1); not null default 0;comment:是否最新"`
	IsHot          bool    `gorm:"type:tinyint(1); not null default 0;comment:是否热播"`
	Visible
	DateModel
	DeletedModel
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
		Introduce:     video.Introduction,
	}
}

func (VideoES) GetIndexName() string {
	return "video"
}

func (VideoES) GetMapping() string {
	videoMapping := `{
    "mappings":{
        "properties":{
            "user_id":{
                "type":"integer"
            },
            "category_id":{
                "type":"integer"
            },
            "favorite_count":{
                "type":"integer"
            },
            "like_count":{
                "type":"integer"
            },
            "play_count":{
                "type":"integer"
            },
            "barrage_count":{
                "type":"integer"
            },
            "is_recommend":{
                "type":"boolean"
            },
            "is_new":{
                "type":"boolean"
            },
            "is_hot":{
                "type":"boolean"
            },
            "is_visible":{
                "type":"boolean"
            },
            "score":{
                "type":"float"
            },
            "ali_cloud_id":{
                "type":"text"
            },
            "name":{
                "type":"text",
                "analyzer":"ik_max_word"
            },
            "introduce":{
                "type":"text",
                "analyzer":"ik_max_word"
            }
        }
    }
}`

	return videoMapping
}
