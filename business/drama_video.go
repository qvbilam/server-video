package business

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"video/model"
)

type DramaVideoBusiness struct {
	Id      int64
	DramaId int64
	VideoId int64
	Episode *int64
}

func (b *DramaVideoBusiness) UpdateEpisode(tx *gorm.DB) (int64, error) {
	if b.DramaId == 0 {
		return 0, nil
	}
	dramaEntity := model.Drama{}
	if res := tx.First(&dramaEntity, b.DramaId); res.RowsAffected == 0 {
		return 0, status.Errorf(codes.NotFound, "剧集不存在")
	}

	// 设置查询条件
	episodeCondition := model.DramaVideo{}
	episodeCondition.DramaId = b.DramaId
	if b.Episode != nil {
		episodeCondition.Episode = *b.Episode
	}
	if b.VideoId != 0 {
		episodeCondition.VideoId = b.VideoId
	}

	// 更新剧集
	episodeEntity := model.DramaVideo{}
	res := tx.Where(episodeCondition).First(&episodeEntity)
	if res.RowsAffected == 0 {
		tx.Create(&episodeEntity)
	} else {
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
