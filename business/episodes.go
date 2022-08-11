package business

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"video/global"
	"video/model"
)

type Episodes struct {
	Id           int64
	VideoId      int64
	AliCloudId   string
	Name         string
	Introduction string
	Icon         string
	Url          string
	Number       int64
}

func (s *Episodes) Create() (int64, error) {
	entity := model.Episodes{
		VideoId:      s.VideoId,
		AliCloudId:   s.AliCloudId,
		Name:         s.Name,
		Introduction: s.Introduction,
		Icon:         s.Icon,
		Url:          s.Url,
		Number:       s.Number,
	}

	tx := global.DB.Begin()

	// 1. 创建剧集
	res := tx.Create(&entity)
	if res.Error != nil {
		tx.Rollback()
		return 0, status.Errorf(codes.Internal, "创建剧集异常: %s", res.Error)
	}

	if res.RowsAffected == 0 {
		tx.Rollback()
		return 0, status.Errorf(codes.Internal, "创建剧集失败")
	}

	// 2. 增加视频数量
	videoEntity := model.Video{}
	videoRes := tx.Model(&videoEntity).
		Where(model.Video{IDModel: model.IDModel{ID: s.VideoId}}).
		Update("count", gorm.Expr("count + ?", 1))
	if videoRes.Error != nil {
		tx.Rollback()
		return 0, status.Errorf(codes.Internal, "创建视频异常: %s", res.Error)
	}

	if videoRes.RowsAffected == 0 {
		tx.Rollback()
		return 0, status.Errorf(codes.Internal, "创建视频失败")
	}

	// todo 3. 增加用户视频数量
	// todo 4. 消息队列(关注用户通知等...)

	tx.Commit()
	return entity.ID, nil
}
