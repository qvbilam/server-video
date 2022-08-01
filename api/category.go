package api

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	proto "video/api/v1"
)

type CategoryServer struct {
	proto.UnimplementedCategoryServer
}

func (s *CategoryServer) Create(ctx context.Context, request *proto.UpdateCategoryRequest) (*proto.CategoryResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (s *CategoryServer) Update(ctx context.Context, request *proto.UpdateCategoryRequest) (*emptypb.Empty, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (s *CategoryServer) Delete(ctx context.Context, request *proto.UpdateCategoryRequest) (*emptypb.Empty, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (s *CategoryServer) Get(ctx context.Context, request *proto.GetCategoryRequest) (*proto.GetCategoryResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}
