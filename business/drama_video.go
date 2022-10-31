package business

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"video/model"
)

type DramaVideoBusiness struct {
	Id      int64
	DramaId int64
	VideoId int64
	Episode *int64
}

func (b *DramaVideoBusiness) Create(tx *gorm.DB) error {
	if b.Episode == nil || b.DramaId == 0 || b.VideoId == 0 {
		return status.Errorf(codes.InvalidArgument, "缺少参数")
	}

	if err := b.Exists(tx); err != nil {
		return err
	}

	entity := model.DramaVideo{
		DramaId: b.DramaId,
		VideoId: b.VideoId,
		Episode: *b.Episode,
	}

	// Skip all associations when creating a DramaVideo
	if res := tx.Omit(clause.Associations).Create(&entity); res.RowsAffected == 0 || res.Error != nil {
		return status.Errorf(codes.Internal, "创建失败")
	}
	return nil
}

func (b *DramaVideoBusiness) Update(tx *gorm.DB) (int64, error) {
	if b.DramaId == 0 {
		return 0, status.Errorf(codes.InvalidArgument, "错误参数")
	}
	dramaEntity := model.Drama{}
	if res := tx.First(&dramaEntity, b.DramaId); res.RowsAffected == 0 {
		return 0, status.Errorf(codes.NotFound, "剧集不存在")
	}

	// 验证剧是否有重复集
	if err := b.Exists(tx); err != nil {
		return 0, err
	}

	// 更新剧集(查询剧是否包含视频
	episodeEntity := model.DramaVideo{}
	res := tx.Where(model.DramaVideo{DramaId: b.DramaId, VideoId: b.VideoId}).First(&episodeEntity)
	if res.RowsAffected == 0 {
		tx.Create(&model.DramaVideo{DramaId: b.DramaId, VideoId: b.VideoId, Episode: *b.Episode})
	} else {
		episodeEntity.Episode = *b.Episode
		tx.Save(&episodeEntity)
	}

	// 获取视频最大的剧集
	lastEpisodeEntity := model.DramaVideo{}
	tx.Where(model.DramaVideo{DramaId: b.DramaId}).Order("episode desc").Last(&lastEpisodeEntity)
	if lastEpisodeEntity.Episode > dramaEntity.NewEpisode {
		dramaEntity.NewEpisode = lastEpisodeEntity.Episode
		if res := tx.Save(&dramaEntity); res.RowsAffected == 0 {
			return 0, status.Errorf(codes.Internal, "更新剧集失败")
		}
	}

	return res.RowsAffected, nil
}

func (b *DramaVideoBusiness) Exists(tx *gorm.DB) error {
	condition := model.DramaVideo{
		DramaId: b.DramaId,
		Episode: *b.Episode,
	}

	exists := model.DramaVideo{}
	res := tx.Where(condition).First(&exists)
	if b.Id != 0 {
		if exists.ID != b.Id && res.RowsAffected != 0 {
			return status.Errorf(codes.AlreadyExists, "剧集已存在")
		}
	}

	return nil
}
