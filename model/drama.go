package model

import (
	"context"
	"gorm.io/gorm"
	"strconv"
	"video/global"
)

type DramaES struct {
	ID            int64 `json:"id"`
	UserID        int64 `json:"user_id"`
	RegionID      int64 `json:"region_id"`
	CategoryID    int64 `json:"category_id"`
	EpisodeCount  int64 `json:"episode_count"`
	FavoriteCount int64 `json:"favorite_count"`
	LikeCount     int64 `json:"like_count"`
	PlayCount     int64 `json:"play_count"`
	BarrageCount  int64 `json:"barrage_count"`
	Year          int64 `json:"year"`
	Quarter       int64 `json:"quarter"`

	IsNew     bool `json:"is_new"`
	IsHot     bool `json:"is_hot"`
	IsEnd     bool `json:"is_end"`
	IsVisible bool `json:"isVisible"`

	Score float64 `json:"score"`

	Name      string `json:"name"`
	Introduce string `json:"introduce"`
}

// Drama 视频剧集
type Drama struct {
	IDModel
	UserModel
	RegionId      int64   `gorm:"type:int not null default 0;comment:区域id;index:idx_region_id"`
	CategoryId    int64   `gorm:"type:int not null default 0;comment:分类id;index:idx_category_id"`
	Name          string  `gorm:"type:varchar(255); not null default '';comment:名称"`
	Introduction  string  `gorm:"type:varchar(2048); not null default '';comment:简介"`
	Icon          string  `gorm:"type:varchar(255); not null default '';comment:横版图标"`
	Score         float64 `gorm:"type:decimal(1,1);not null default 0;comment:评分"`
	NewEpisode    int64   `gorm:"type:int; not null default 0;comment:剧集"`
	EpisodeCount  int64   `gorm:"type:int; not null default 0;comment:剧集数"`
	FavoriteCount int64   `gorm:"type:int; not null default 0;comment:收藏数量"`
	LikeCount     int64   `gorm:"type:int; not null default 0;comment:点赞数量"`
	PlayCount     int64   `gorm:"type:int; not null default 0;comment:播放数量"`
	BarrageCount  int64   `gorm:"type:int; not null default 0;comment:弹幕数量"`
	IsNew         bool    `gorm:"type:tinyint(1); not null default 0;comment:是否最新"`
	IsHot         bool    `gorm:"type:tinyint(1); not null default 0;comment:是否热播"`
	IsEnd         bool    `gorm:"type:tinyint(1); not null default 0;comment:是否完结"`
	Year          int64   `gorm:"type:int; not null default 0;comment:年份"`
	Quarter       int64   `gorm:"type:tinyint(4); not null default 0;comment:季度"`
	Visible
	DateModel
	DeletedModel
}

type DramaVideo struct {
	IDModel
	DramaId int64 `gorm:"type:int not null default 0;comment:剧id;index:idx_drama_id_video_id_episode"`
	VideoId int64 `gorm:"type:int not null default 0;comment:视频id;index:idx_drama_id_video_id_episode;index:idx_video_id"`
	Episode int64 `gorm:"type:int not null default 0;comment:集;index:idx_drama_id_video_id_episode"`
}

func (drama *Drama) AfterCreate(tx *gorm.DB) error {
	esModel := dramaModelToEsIndex(drama)
	// 写入es
	_, err := global.ES.
		Index().
		Index(DramaES{}.GetIndexName()).
		BodyJson(esModel).
		Id(strconv.Itoa(int(drama.ID))).
		Do(context.Background())
	return err
}

func (drama *Drama) AfterUpdate(tx *gorm.DB) error {
	esModel := dramaModelToEsIndex(drama)
	// 更新es. 指定 id 防止重复
	_, err := global.ES.
		Update().
		Index(DramaES{}.GetIndexName()).
		Doc(esModel).
		Id(strconv.Itoa(int(drama.ID))).
		Do(context.Background())

	return err
}

func (drama *Drama) AfterDelete(tx *gorm.DB) error {
	// 删除 es 数据
	_, err := global.ES.
		Delete().
		Index(DramaES{}.GetIndexName()).
		Id(strconv.Itoa(int(drama.ID))).
		Do(context.Background())

	return err
}

func dramaModelToEsIndex(drama *Drama) *DramaES {
	return &DramaES{
		ID:            drama.ID,
		UserID:        drama.UserID,
		RegionID:      drama.RegionId,
		CategoryID:    drama.CategoryId,
		EpisodeCount:  drama.EpisodeCount,
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
		Introduce:     drama.Introduction,
	}
}

func (DramaES) GetIndexName() string {
	return "drama"
}

func (DramaES) GetMapping() string {
	dramaMapping := `{
    "mappings":{
        "properties":{
            "user_id":{
                "type":"integer"
            },
            "region_id":{
                "type":"integer"
            },
            "category_id":{
                "type":"integer"
            },
            "episode_count":{
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
            "is_new":{
                "type":"boolean"
            },
            "is_hot":{
                "type":"boolean"
            },
            "is_end":{
                "type":"boolean"
            },
            "is_visible":{
                "type":"boolean"
            },
            "score":{
                "type":"float"
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

	return dramaMapping
}
