package main

import (
	"gorm.io/gorm"
	"video/model"
)

const defaultUserId = 1
const categoryAnime = 1

const videoTypeAnime = "anime"

func CreateDramaVideo(db *gorm.DB) {
	db.Exec("truncate table drama")
	db.Exec("truncate table video")
	db.Exec("truncate table drama_video")

	// 创建动漫
	cartoonDrama, cartoonVideo, cartoonDramaVideo := getAnime()
	db.CreateInBatches(cartoonDrama, 100)
	db.CreateInBatches(cartoonVideo, 100)
	db.CreateInBatches(cartoonDramaVideo, 100)
}

func getAnime() (*[]model.Drama, *[]model.Video, *[]model.DramaVideo) {
	dramaModel := []model.Drama{
		{
			UserModel: model.UserModel{
				UserID: defaultUserId,
			},
			CategoryId: categoryAnime,
			Name:       "某科学的超电磁炮",
			Introduce:  "某科学的电磁炮",
			Cover:      "http://damowang.test.upcdn.net/video/images/we5310pi2yur2to539q2.gif",
			Visible: model.Visible{
				IsVisible: true,
			},
		},
		{
			UserModel: model.UserModel{
				UserID: defaultUserId,
			},
			CategoryId: categoryAnime,
			Name:       "fate zero 第一季",
			Introduce:  "传说中，圣杯是能够实现拥有者愿望的宝物。为了追求圣杯的力量，7位魔术师各自召唤英灵，展开争夺圣杯的战斗，这就是圣杯战争。\n时间退回到第五次圣杯战争的10年前，第四次圣杯战争，参与者正是士郎他们的父辈。为了得到圣杯不择手段的卫宫切嗣，年轻时代的言峰绮礼，间桐家与远坂家的关系，同样身为王却意志不同的三位英灵。第四次圣杯之战就此爆发。",
			Cover:      "http://damowang.test.upcdn.net/video/dm/FZ/f101.png",
			Visible: model.Visible{
				IsVisible: true,
			},
		},
	}

	videoModel := []model.Video{
		{
			UserModel:  model.UserModel{UserID: defaultUserId},
			FileId:     1,
			CategoryId: categoryAnime,
			Name:       "超电磁炮（Railgun)",
			Introduce:  "第一集",
			Cover:      "http://damowang.test.upcdn.net/video/images/qo3ui5rpw2e130212t6y.jpeg",
		},
		{
			UserModel:  model.UserModel{UserID: defaultUserId},
			FileId:     2,
			CategoryId: categoryAnime,
			Name:       "英灵召唤",
			Introduce:  "第一集",
			Cover:      "http://damowang.test.upcdn.net/video/dm/FZ/f101.png",
		},
		{
			UserModel:  model.UserModel{UserID: defaultUserId},
			FileId:     3,
			CategoryId: categoryAnime,
			Name:       "虚假的战争",
			Introduce:  "第二集",
			Cover:      "http://damowang.test.upcdn.net/video/images/y42iw56491otuqpre641.png",
		},
	}

	dramaVideoModel := []model.DramaVideo{
		{
			DramaId: 1,
			VideoId: 1,
			Episode: 1,
		},
		{
			DramaId: 2,
			VideoId: 2,
			Episode: 1,
		},
		{
			DramaId: 2,
			VideoId: 3,
			Episode: 2,
		},
	}

	return &dramaModel, &videoModel, &dramaVideoModel
}
