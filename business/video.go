package business

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"video/global"
	"video/model"
)

type Video struct {
	Id             int64
	UserId         int64
	RegionId       int64
	CategoryId     int64
	Name           string
	Introduction   string
	Icon           string
	HorizontalIcon string
	TotalCount     int64
}

func (s *Video) Create() (int64, error) {
	// 开启事物
	tx := global.DB.Begin()

	// 视频实例
	entity := model.Video{
		UserModel: model.UserModel{
			UserID: s.UserId,
		},
		RegionId:       s.RegionId,
		CategoryId:     s.CategoryId,
		Name:           s.Name,
		Introduction:   s.Introduction,
		Icon:           s.Icon,
		HorizontalIcon: s.HorizontalIcon,
		TotalCount:     s.TotalCount,
		Visible:        model.Visible{},
	}

	cvRes := tx.Create(&entity) // 同时会在model.After 写入es
	if cvRes.RowsAffected == 0 {
		tx.Rollback()
		return 0, status.Errorf(codes.Internal, "创建视频失败: %s", cvRes.Error)
	}

	if cvRes.Error != nil {
		tx.Rollback()
		return 0, status.Errorf(codes.Internal, "创建视频异常: %s", cvRes.Error)
	}

	tx.Commit()
	return entity.ID, nil
}
