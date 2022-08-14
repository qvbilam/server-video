package api

import (
	"go.uber.org/zap"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	proto "video/api/pb"
	userProto "video/api/user/pb"
	"video/business"
	"video/global"
	"video/model"
)

type VideoServer struct {
	proto.UnimplementedVideoServer
}

func (s *VideoServer) Create(ctx context.Context, request *proto.UpdateVideoRequest) (*proto.VideoResponse, error) {
	// todo 验证分类 区分用户角色
	categoryBusiness := business.Category{Id: request.CategoryId}
	if _, err := categoryBusiness.Exists(); err != nil {
		//return nil, err
	}

	// todo 验证区域 区分用户角色
	regionBusiness := business.Region{Id: request.RegionId}
	if _, err := regionBusiness.Exists(); err != nil {
		//return nil, err
	}

	videoBusiness := business.Video{
		UserId:         request.UserId,
		RegionId:       request.RegionId,
		CategoryId:     request.CategoryId,
		Name:           request.Name,
		Introduction:   request.Introduction,
		Icon:           request.Icon,
		HorizontalIcon: request.HorizontalIcon,
		TotalCount:     request.TotalCount,
	}
	videoId, err := videoBusiness.Create()
	if err != nil {
		return nil, err
	}

	return &proto.VideoResponse{Id: videoId}, nil
}

func (s *VideoServer) Update(ctx context.Context, request *proto.UpdateVideoRequest) (*emptypb.Empty, error) {
	videoBusiness := business.Video{
		Id:             request.Id,
		UserId:         request.UserId,
		RegionId:       request.RegionId,
		CategoryId:     request.CategoryId,
		Name:           request.Name,
		Introduction:   request.Introduction,
		Icon:           request.Icon,
		HorizontalIcon: request.HorizontalIcon,
		TotalCount:     request.TotalCount,
		Score:          float64(request.Score),
	}
	_, err := videoBusiness.Update()
	if err != nil {
		return nil, status.Errorf(codes.Unimplemented, err.Error())
	}

	return &emptypb.Empty{}, nil
}

func (s *VideoServer) Delete(ctx context.Context, request *proto.UpdateVideoRequest) (*emptypb.Empty, error) {
	videoBusiness := business.Video{Id: request.Id}
	_, err := videoBusiness.Delete()
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return &emptypb.Empty{}, nil
}

func (s *VideoServer) Get(ctx context.Context, request *proto.SearchVideoRequest) (*proto.VideosResponse, error) {
	// 查询列表
	videoBusiness := searchRequestToCondition(request)
	res, err := videoBusiness.List()
	if err != nil {
		return nil, err
	}

	response := proto.VideosResponse{Total: res.Total}
	// 结果转换
	var userIds []int64
	userMaps := make(map[int64]*userProto.UserResponse)
	for _, video := range res.Videos {
		userIds = append(userIds, video.UserID)
		videoResponse := modelToResponse(video)
		response.Videos = append(response.Videos, &videoResponse)
	}

	// 获取用户信息
	users, err := global.UserServerClient.ListByIds(context.Background(), &userProto.ListByIdsRequest{Ids: userIds})
	if err != nil {
		zap.S().Errorf("获取用户信息失败: %s", err)
	}
	for _, user := range users.Users {
		userMaps[user.Id] = user
	}

	// 视频追加用户信息
	for _, video := range response.Videos {
		if userMaps[video.User.Id] != nil {
			video.User = &proto.VideoUserResponse{
				Id:       userMaps[video.User.Id].Id,
				Nickname: userMaps[video.User.Id].Nickname,
				Avatar:   userMaps[video.User.Id].Avatar,
				Gender:   userMaps[video.User.Id].Gender,
			}
		}
	}

	return &response, nil
}

func (s *VideoServer) GetDetail(ctx context.Context, request *proto.GetVideoRequest) (*proto.VideoResponse, error) {
	videoBusiness := business.Video{}
	res, err := videoBusiness.Detail()
	if err != nil {
		return nil, err
	}
	response := modelToResponse(*res)
	return &response, nil
}

func searchRequestToCondition(request *proto.SearchVideoRequest) business.Video {
	return business.Video{
		UserId:           request.UserId,
		RegionId:         request.RegionId,
		CategoryId:       request.CategoryId,
		Keyword:          request.Keyword,
		IsRecommend:      request.IsRecommend,
		IsNew:            request.IsNew,
		IsHot:            request.IsHot,
		IsEnd:            request.IsEnd,
		IsVisible:        request.IsVisible,
		FavoriteCountMin: request.FavoriteCountMin,
		FavoriteCountMax: request.FavoriteCountMax,
		LikeCountMin:     request.LikeCountMin,
		LikeCountMax:     request.LikeCountMax,
		PlayCountMin:     request.PlayCountMin,
		PlayCountMax:     request.PlayCountMax,
		BarrageCountMin:  request.BarrageCountMin,
		BarrageCountMax:  request.BarrageCountMax,
		Page:             request.Page,
		PerPage:          request.PerPage,
	}
}

func modelToResponse(video model.Video) proto.VideoResponse {
	return proto.VideoResponse{
		Id: video.ID,
		User: &proto.VideoUserResponse{
			Id: video.UserID,
		},
		Region: &proto.RegionResponse{
			Id: video.RegionId,
		},
		Category: &proto.CategoryResponse{
			Id: video.CategoryId,
		},
		Name:           video.Name,
		Introduction:   video.Introduction,
		Icon:           video.Icon,
		HorizontalIcon: video.HorizontalIcon,
		Score:          float32(video.Score),
		Count:          video.Count,
		TotalCount:     video.TotalCount,
		FavoriteCount:  video.FavoriteCount,
		LikeCount:      video.LikeCount,
		PlayCount:      video.PlayCount,
		BarrageCount:   video.BarrageCount,
		IsRecommend:    video.IsRecommend,
		IsHot:          video.IsHot,
		IsEnd:          video.IsEnd,
		CreatedAt:      video.CreatedAt.Unix(),
		UpdatedAt:      video.UpdatedAt.Unix(),
	}
}
