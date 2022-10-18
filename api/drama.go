package api

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	proto "video/api/qvbilam/video/v1"
	"video/business"
)

type DramaServer struct {
	proto.UnimplementedDramaServer
}

func (s *DramaServer) Create(ctx context.Context, request *proto.UpdateDramaRequest) (*proto.DramaResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (s *DramaServer) Update(ctx context.Context, request *proto.UpdateDramaRequest) (*emptypb.Empty, error) {
	b := business.DramaBusiness{}
	b.Id = 2
	isNew := true
	b.IsNew = &isNew
	_, _ = b.Update()
	return &emptypb.Empty{}, nil
}

func (s *DramaServer) Delete(ctx context.Context, request *proto.UpdateDramaRequest) (*emptypb.Empty, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (s *DramaServer) Get(ctx context.Context, request *proto.SearchDramaRequest) (*proto.DramaListResponse, error) {
	b := business.DramaBusiness{}
	b.Keyword = request.Keyword
	_, err := b.List()
	if err != nil {
		return nil, err
	}
	return &proto.DramaListResponse{
		Total: 0,
		Drama: nil,
	}, nil
}

func (s *DramaServer) GetDetail(ctx context.Context, request *proto.SearchDramaRequest) (*proto.DramaResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}
