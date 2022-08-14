package api

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	proto "video/api/pb"
)

type RegionServer struct {
	proto.UnimplementedRegionServer
}

func (s *RegionServer) Create(ctx context.Context, request *proto.UpdateRegionRequest) (*proto.RegionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (s *RegionServer) Update(ctx context.Context, request *proto.UpdateRegionRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (s *RegionServer) Delete(ctx context.Context, request *proto.DeleteRegionRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (s *RegionServer) Get(ctx context.Context, request *proto.GetRegionRequest) (*proto.GetRegionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
