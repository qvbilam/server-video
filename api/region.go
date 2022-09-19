package api

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	proto "video/api/qvbilam/video/v1"
	"video/business"
)

type RegionServer struct {
	proto.UnimplementedRegionServer
}

func (s *RegionServer) Create(ctx context.Context, request *proto.UpdateRegionRequest) (*proto.RegionResponse, error) {
	b := business.Region{
		Name:      request.Name,
		Icon:      request.Icon,
		IsVisible: &request.IsVisible,
	}
	id, err := b.Create()
	if err != nil {
		return nil, err
	}

	return &proto.RegionResponse{Id: id}, nil
}

func (s *RegionServer) Update(ctx context.Context, request *proto.UpdateRegionRequest) (*emptypb.Empty, error) {
	b := business.Region{
		Id:        request.Id,
		Name:      request.Name,
		Icon:      request.Icon,
		IsVisible: &request.IsVisible,
	}
	if _, err := b.Update(); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (s *RegionServer) Delete(ctx context.Context, request *proto.DeleteRegionRequest) (*emptypb.Empty, error) {
	b := business.Region{Id: request.Id}
	if _, err := b.Delete(); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (s *RegionServer) Get(ctx context.Context, request *proto.GetRegionRequest) (*proto.GetRegionResponse, error) {
	b := business.Region{
		IsVisible: &request.IsVisible,
	}
	regions, err := b.List()
	if err != nil {
		return nil, err
	}

	response := proto.GetRegionResponse{}
	for _, region := range *regions {
		response.Region = append(response.Region, &proto.RegionResponse{
			Id:   region.ID,
			Name: region.Name,
			Icon: region.Icon,
		})
	}

	return &response, nil
}
