package model

import (
	"gorm.io/gorm"
	"video/global"
)

type DramaVideo struct {
	IDModel
	DramaId int64 `gorm:"type:int not null default 0;comment:剧id;index:idx_drama_id_video_id_episode"`
	VideoId int64 `gorm:"type:int not null default 0;comment:视频id;index:idx_drama_id_video_id_episode;index:idx_video_id"`
	Episode int64 `gorm:"type:int not null default 0;comment:集;index:idx_drama_id_video_id_episode"`
	Video   Video `gorm:"foreignkey:VideoId;references:ID"`
}

func GetDramaVideos(dramaId int64) *[]DramaVideo {
	var models []DramaVideo
	global.DB.Where(DramaVideo{DramaId: dramaId}).Preload("Video", func(db *gorm.DB) *gorm.DB {
		return db.Select("id", "file_id", "name", "introduce")
	}).Find(&models)

	return &models
}
