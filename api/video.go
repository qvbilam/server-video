package api

import (
	"go.uber.org/zap"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	userProto "video/api/qvbilam/user/v1"
	proto "video/api/qvbilam/video/v1"
	"video/business"
	"video/global"
	"video/model"
)

type VideoServer struct {
	proto.UnimplementedVideoServer
}

func (s *VideoServer) Create(ctx context.Context, request *proto.UpdateVideoRequest) (*proto.VideoResponse, error) {
	// todo 验证分类 区分用户角色
	categoryBusiness := business.CategoryBusiness{Id: request.CategoryId}
	if _, err := categoryBusiness.Exists(); err != nil {
		//return nil, err
	}
	if request.FileId == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "缺少必要参数")
	}

	videoBusiness := business.VideoBusiness{
		DramaId:         request.DramaId,
		Episode:         &request.Episode,
		FileId:          request.FileId,
		UserId:          request.UserId,
		CategoryId:      request.CategoryId,
		Name:            request.Name,
		Introduce:       request.Introduce,
		Cover:           request.Cover,
		HorizontalCover: request.HorizontalCover,
	}
	videoId, err := videoBusiness.Create()
	if err != nil {
		return nil, err
	}

	return &proto.VideoResponse{Id: videoId}, nil
}

func (s *VideoServer) Update(ctx context.Context, request *proto.UpdateVideoRequest) (*emptypb.Empty, error) {
	videoBusiness := business.VideoBusiness{
		DramaId:         request.DramaId,
		Episode:         &request.Episode,
		Id:              request.Id,
		UserId:          request.UserId,
		FileId:          request.FileId,
		CategoryId:      request.CategoryId,
		Name:            request.Name,
		Introduce:       request.Introduce,
		Cover:           request.Cover,
		HorizontalCover: request.HorizontalCover,
		Score:           float64(request.Score),
	}
	_, err := videoBusiness.Update()
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *VideoServer) Delete(ctx context.Context, request *proto.UpdateVideoRequest) (*emptypb.Empty, error) {
	videoBusiness := business.VideoBusiness{Id: request.Id, DramaId: request.DramaId}
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

	response := proto.VideosResponse{Total: res.Total, Videos: nil}
	// 结果转换
	var userIds []int64
	userMaps := make(map[int64]*userProto.UserResponse)
	for _, video := range res.Videos {
		userIds = append(userIds, video.UserID)
		videoResponse := modelToResponse(video)
		response.Videos = append(response.Videos, &videoResponse)
	}

	// 获取用户信息
	users, err := global.UserServerClient.List(context.Background(), &userProto.SearchRequest{Id: userIds})
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
	videoBusiness := business.VideoBusiness{Id: request.Id}
	res, err := videoBusiness.Detail()
	if err != nil {
		return nil, err
	}
	response := modelToResponse(*res)
	return &response, nil
}

func (s *VideoServer) Play(ctx context.Context, request *proto.UpdateVideoRequest) (*emptypb.Empty, error) {
	videoBusiness := business.VideoBusiness{
		Id:        request.Id,
		IP:        request.Ip,
		PlayCount: request.PlayCount,
	}
	// todo 验证是否增加播放量
	if err := videoBusiness.Play(); err != nil {
		return nil, err
	}
	return nil, nil
}

func (s VideoServer) Barrage(ctx context.Context, request *proto.UpdateVideoRequest) (*emptypb.Empty, error) {
	videoBusiness := business.VideoBusiness{
		Id:        request.Id,
		IP:        request.Ip,
		PlayCount: request.PlayCount,
	}
	if err := videoBusiness.Barrage(); err != nil {
		return nil, err
	}
	return nil, nil
}

func searchRequestToCondition(request *proto.SearchVideoRequest) business.VideoBusiness {
	return business.VideoBusiness{
		UserId:           request.UserId,
		CategoryId:       request.CategoryId,
		Keyword:          request.Keyword,
		IsRecommend:      &request.IsRecommend,
		IsNew:            &request.IsNew,
		IsHot:            &request.IsHot,
		IsVisible:        &request.IsVisible,
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
		Id:     video.ID,
		FileId: video.FileId,
		User: &proto.VideoUserResponse{
			Id: video.UserID,
		},
		Category: &proto.CategoryResponse{
			Id: video.CategoryId,
		},
		Name:            video.Name,
		Introduce:       video.Introduce,
		Cover:           video.Cover,
		HorizontalCover: video.HorizontalCover,
		Score:           float32(video.Score),
		FavoriteCount:   video.FavoriteCount,
		LikeCount:       video.LikeCount,
		PlayCount:       video.PlayCount,
		BarrageCount:    video.BarrageCount,
		IsRecommend:     video.IsRecommend,
		IsHot:           video.IsHot,
		CreatedTime:     video.CreatedAt.Unix(),
	}
}
