package doc

import "github.com/olivere/elastic/v7"

type Video struct {
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

	Type      string `json:"type"`
	Name      string `json:"name"`
	Introduce string `json:"introduce"`
}

type VideoSearch struct {
	Keyword string // 搜索
	// 多层级查找
	CategoryIds      []interface{}
	UserId           int64
	FavoriteCountMin int64
	FavoriteCountMax int64
	LikeCountMin     int64
	LikeCountMax     int64
	PlayCountMin     int64
	PlayCountMax     int64
	BarrageCountMin  int64
	BarrageCountMax  int64

	IsRecommend *bool
	IsNew       *bool
	IsHot       *bool
	IsVisible   *bool
}

func (Video) GetIndexName() string {
	return "video"
}

func (Video) GetMapping() string {
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
            }
        }
    }
}`

	return videoMapping
}

func (s *VideoSearch) GetQuery() *elastic.BoolQuery {
	// match bool 复合查询
	q := elastic.NewBoolQuery()

	if s.Keyword != "" { // 搜索 名称, 简介
		q = q.Must(elastic.NewMultiMatchQuery(s.Keyword, "name", "introduction"))
	}
	if s.UserId > 0 { // 搜索用户
		q = q.Filter(elastic.NewTermQuery("user_id", s.UserId))
	}

	if s.IsHot != nil { // 搜索热度
		q = q.Filter(elastic.NewTermQuery("is_hot", s.IsHot))
	}
	if s.IsNew != nil { // 搜索新品
		q = q.Filter(elastic.NewTermQuery("is_new", s.IsNew))
	}
	if s.IsVisible != nil { // 搜索展示状态
		q = q.Filter(elastic.NewTermQuery("is_visible", s.IsNew))
	}

	// 多级分类查找
	if len(s.CategoryIds) > 0 {
		q = q.Filter(elastic.NewTermsQuery("category_id", s.CategoryIds...))
	}

	// 范围查询
	if s.FavoriteCountMin > 0 {
		q = q.Filter(elastic.NewRangeQuery("favorite_count").Gte(s.FavoriteCountMin))
	}

	if s.FavoriteCountMax > 0 {
		q = q.Filter(elastic.NewRangeQuery("favorite_count").Lte(s.FavoriteCountMax))
	}

	if s.LikeCountMin > 0 {
		q = q.Filter(elastic.NewRangeQuery("like_count").Gte(s.LikeCountMin))
	}

	if s.LikeCountMax > 0 {
		q = q.Filter(elastic.NewRangeQuery("like_count").Lte(s.LikeCountMax))
	}

	if s.PlayCountMin > 0 {
		q = q.Filter(elastic.NewRangeQuery("play_count").Gte(s.PlayCountMin))
	}

	if s.PlayCountMax > 0 {
		q = q.Filter(elastic.NewRangeQuery("play_count").Lte(s.PlayCountMax))
	}

	if s.BarrageCountMin > 0 {
		q = q.Filter(elastic.NewRangeQuery("barrage_count").Gte(s.BarrageCountMin))
	}

	if s.BarrageCountMax > 0 {
		q = q.Filter(elastic.NewRangeQuery("barrage_count").Lte(s.BarrageCountMax))
	}

	return q
}
