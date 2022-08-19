package api

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	proto "video/api/pb"
	"video/business"
)

type CategoryServer struct {
	proto.UnimplementedCategoryServer
}

func (s *CategoryServer) Create(ctx context.Context, request *proto.UpdateCategoryRequest) (*proto.CategoryResponse, error) {
	// todo 验证用户权限
	b := business.Category{
		Name:      request.Name,
		Icon:      request.Icon,
		Level:     request.Level,
		ParentId:  &request.ParentId,
		IsVisible: &request.IsVisible,
	}
	id, err := b.Create()
	if err != nil {
		return nil, err
	}
	return &proto.CategoryResponse{Id: id}, nil
}

func (s *CategoryServer) Update(ctx context.Context, request *proto.UpdateCategoryRequest) (*emptypb.Empty, error) {
	// todo 验证用户权限
	b := business.Category{
		Id:        request.Id,
		Name:      request.Name,
		Icon:      request.Icon,
		Level:     request.Level,
		ParentId:  &request.ParentId,
		IsVisible: &request.IsVisible,
	}
	_, err := b.Update()
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (s *CategoryServer) Delete(ctx context.Context, request *proto.UpdateCategoryRequest) (*emptypb.Empty, error) {
	// todo 验证用户权限
	b := business.Category{
		Id: request.Id,
	}
	_, err := b.Delete()
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (s *CategoryServer) Get(ctx context.Context, request *proto.GetCategoryRequest) (*proto.GetCategoryResponse, error) {
	b := business.Category{
		Level:     request.Level,
		ParentId:  &request.ParentId,
		IsVisible: &request.IsVisible,
	}
	category, err := b.List()
	if err != nil {
		return nil, err
	}

	response := proto.GetCategoryResponse{}
	for _, c := range *category {
		response.Category = append(response.Category, &proto.CategoryResponse{
			Id:       c.ID,
			Name:     c.Name,
			Icon:     c.Icon,
			Level:    c.Level,
			ParentId: c.ParentId,
		})
	}

	return &response, nil
}
