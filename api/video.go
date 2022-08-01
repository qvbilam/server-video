package api

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	proto "video/api/v1"
)

type VideoServer struct {
	proto.UnimplementedVideoServer
}

func (s *VideoServer) Create(ctx context.Context, request *proto.UpdateVideoRequest) (*proto.VideoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}

func (s *VideoServer) Update(ctx context.Context, request *proto.UpdateVideoRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}

func (s *VideoServer) Delete(ctx context.Context, request *proto.UpdateVideoRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

func (s *VideoServer) Get(ctx context.Context, request *proto.GetVideoRequest) (*proto.VideosResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}

func (s *VideoServer) GetDetail(ctx context.Context, request *proto.GetVideoRequest) (*proto.VideoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDetail not implemented")
}
