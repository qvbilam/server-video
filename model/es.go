package model

type VideoES struct {
}

type EpisodesES struct {
}

func (e VideoES) GetIndexName() string {
	return "video"
}

func (e EpisodesES) GetIndexName() string {
	return "episodes"
}

func (e VideoES) getMapping() string {
	m := `
{
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
            "is_hot":{
                "type":"boolean"
            },
            "is_end":{
                "type":"boolean"
            },
            "is_new":{
                "type":"boolean"
            },
			"is_visible":{
                "type":"boolean"
			}
            "score":{
                "type":"float"
            },
            "name":{
                "type":"text",
                "analyzer":"ik_max_word"
            },
            "introduction":{
                "type":"text",
                "analyzer":"ik_max_word"
            }
        }
    }
}
`

	return m
}

func (e EpisodesES) getMapping() string {
	m := `
{
    "mappings":{
        "properties":{
            "video_id":{
                "type":"integer"
            },
            "number":{
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
			"is_hot":{
                "type":"boolean"
			},
			"is_new":{
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
            "introduction":{
                "type":"text",
                "analyzer":"ik_max_word"
            }
        }
    }
}
`

	return m
}
