package main

import (
	"gorm.io/gorm"
	"video/model"
)

func CreateRegions(db *gorm.DB) {
	models := []model.Region{
		{
			Name: "国产",
			Icon: "",
			Visible: model.Visible{
				IsVisible: true,
			},
		},
		{
			Name: "日本",
			Icon: "",
			Visible: model.Visible{
				IsVisible: true,
			},
		},
		{
			Name: "欧美",
			Icon: "",
			Visible: model.Visible{
				IsVisible: true,
			},
		},
		{
			Name: "其他",
			Icon: "",
			Visible: model.Visible{
				IsVisible: true,
			},
		},
	}

	db.Exec("truncate table region")
	db.CreateInBatches(models, 100)
}
