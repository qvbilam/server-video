package api

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	proto "video/api/qvbilam/video/v1"
)

type EpisodesServer struct {
	proto.UnimplementedEpisodesServer
}

func (s *EpisodesServer) Create(ctx context.Context, request *proto.UpdateEpisodesRequest) (*proto.EpisodesResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (s *EpisodesServer) Update(ctx context.Context, request *proto.UpdateEpisodesRequest) (*emptypb.Empty, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (s *EpisodesServer) Delete(ctx context.Context, request *proto.UpdateEpisodesRequest) (*emptypb.Empty, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (s *EpisodesServer) Get(ctx context.Context, request *proto.GetEpisodesRequest) (*proto.EpisodesListResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (s *EpisodesServer) GetDetail(ctx context.Context, request *proto.GetEpisodesRequest) (*proto.EpisodesResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}
