package doc

type DramaVideos struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Introduce string `json:"introduce"`
	Episode   int64  `json:"episode"`
}

type Drama struct {
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

	Type      string `json:"type"`
	Name      string `json:"name"`
	Introduce string `json:"introduce"`

	Videos []DramaVideos `json:"videos"`
}

func (Drama) GetIndexName() string {
	return "drama"
}

func (Drama) GetMapping() string {
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
            "type":{
                "type":"text"
            },
            "name":{
                "type":"text",
                "analyzer":"ik_max_word"
            },
            "introduce":{
                "type":"text",
                "analyzer":"ik_max_word"
            },
            "videos":{
                "properties":{
                    "id": {
                        "type":"integer"
                    },
                    "name": {
                        "type":"text",
                        "analyzer":"ik_max_word"
                    },
                    "introduce": {
                        "type":"text",
                        "analyzer":"ik_max_word"
                    },
                    "episode": {
                        "type":"integer"
                    }
                }
            }
        }
    }
}`

	return dramaMapping
}
