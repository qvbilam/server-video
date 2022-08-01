package main

import (
	"gorm.io/gorm"
	"video/model"
)

func CreateCategory(db *gorm.DB) {
	models := []model.Category{
		{
			Name: "番剧",
			Icon: "",
			Visible: model.Visible{
				IsVisible: true,
			},
		},
		{
			Name: "动画",
			Icon: "",
			Visible: model.Visible{
				IsVisible: true,
			},
		},
		{
			Name: "电影",
			Icon: "",
			Visible: model.Visible{
				IsVisible: true,
			},
		},
		{
			Name: "电视剧",
			Icon: "",
			Visible: model.Visible{
				IsVisible: true,
			},
		},
		{
			Name: "音乐",
			Icon: "",
			Visible: model.Visible{
				IsVisible: true,
			},
		},
	}

	db.Exec("truncate table category")
	db.CreateInBatches(models, 100)
}
