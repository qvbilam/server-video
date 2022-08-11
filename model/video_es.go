package model

type VideoES struct {
	ID            int64 `json:"id"`
	UserID        int64 `json:"user_id"`
	RegionID      int64 `json:"region_id"`
	CategoryID    int64 `json:"category_id"`
	Count         int64 `json:"count"`
	TotalCount    int64 `json:"total_count"`
	FavoriteCount int64 `json:"favorite_count"`
	LikeCount     int64 `json:"like_count"`
	PlayCount     int64 `json:"play_count"`
	BarrageCount  int64 `json:"barrage_count"`

	IsRecommend bool `json:"is_recommend"`
	IsNew       bool `json:"is_new"`
	IsHot       bool `json:"is_hot"`
	IsEnd       bool `json:"is_end"`
	IsVisible   bool `json:"isVisible"`

	Score float64 `json:"score"`

	Name      string `json:"name"`
	Introduce string `json:"introduce"`
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
            "region_id":{
                "type":"integer"
            },
            "category_id":{
                "type":"integer"
            },
            "count":{
                "type":"integer"
            },
            "total_count":{
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

	return videoMapping
}
