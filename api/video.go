package api

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	proto "video/api/v1"
	"video/business"
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

func (s *VideoServer) Get(ctx context.Context, request *proto.GetVideoRequest) (*proto.VideosResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}

func (s *VideoServer) GetDetail(ctx context.Context, request *proto.GetVideoRequest) (*proto.VideoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDetail not implemented")
}
