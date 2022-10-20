package model

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"strconv"
	"video/global"
	"video/model/doc"
)

type DramaVideo struct {
	IDModel
	DramaId int64 `gorm:"type:int not null default 0;comment:剧id;index:idx_drama_id_video_id_episode"`
	VideoId int64 `gorm:"type:int not null default 0;comment:视频id;index:idx_drama_id_video_id_episode;index:idx_video_id"`
	Episode int64 `gorm:"type:int not null default 0;comment:集;index:idx_drama_id_video_id_episode"`
	Video   Video `gorm:"foreignkey:VideoId;references:ID"`
}

func GetDramaVideos(tx *gorm.DB, dramaId int64) *[]DramaVideo {
	var models []DramaVideo
	tx.Where(DramaVideo{DramaId: dramaId}).Preload("Video", func(db *gorm.DB) *gorm.DB {
		return db.Select("id", "file_id", "name", "introduce")
	}).Order("episode").Find(&models)

	return &models
}

func (DramaVideo *DramaVideo) AfterCreate(tx *gorm.DB) error {
	dramaCondition := Drama{
		IDModel: IDModel{ID: DramaVideo.DramaId},
	}
	dramaUpdate := map[string]interface{}{"new_episode": DramaVideo.Episode}
	if res := tx.Model(Drama{}).Where(dramaCondition).Updates(dramaUpdate); res.RowsAffected == 0 {
		tx.Rollback()
		return status.Errorf(codes.NotFound, "")
	}
	// 更新文档
	return DramaVideo.UpdateDoc(tx)
}

func (DramaVideo *DramaVideo) AfterUpdate(tx *gorm.DB) error {
	return DramaVideo.UpdateDoc(tx)
}

func (DramaVideo *DramaVideo) AfterDelete(tx *gorm.DB) error {
	return DramaVideo.UpdateDoc(tx)
}

func (DramaVideo *DramaVideo) UpdateDoc(tx *gorm.DB) error {
	// 更新 dram.video doc
	var ds []doc.DramaVideo
	videos := GetDramaVideos(tx, DramaVideo.DramaId)
	for _, v := range *videos {
		d := v.ToDoc()
		ds = append(ds, *d)
	}
	updateDoc := doc.UpdateDramaVideos{Videos: ds}
	_, err := global.ES.Update().
		Index(doc.Drama{}.GetIndexName()).
		Doc(updateDoc).
		Id(strconv.Itoa(int(DramaVideo.DramaId))).
		Do(context.Background())

	return err
}

func (DramaVideo *DramaVideo) ToDoc() *doc.DramaVideo {
	return &doc.DramaVideo{
		ID:        DramaVideo.Video.ID,
		Name:      DramaVideo.Video.Name,
		Introduce: DramaVideo.Video.Introduce,
		Episode:   DramaVideo.Episode,
	}
}
