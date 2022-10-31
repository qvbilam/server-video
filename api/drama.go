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
	userId := 1 // todo 获取用户id
	b := business.DramaBusiness{
		Id:              request.Id,
		UserId:          int64(userId),
		CategoryId:      request.CategoryId,
		RegionId:        request.RegionId,
		Name:            request.Name,
		Introduce:       request.Introduce,
		Cover:           request.Cover,
		HorizontalCover: request.HorizontalCover,
		TotalCount:      request.TotalCount,
		IsRecommend:     &request.IsRecommend,
		IsNew:           &request.IsNew,
		IsHot:           &request.IsHot,
		IsVisible:       &request.IsVisible,
	}

	entity, err := b.Create()
	if err != nil {
		return nil, err
	}

	return &proto.DramaResponse{Id: entity.ID}, nil
}

func (s *DramaServer) Update(ctx context.Context, request *proto.UpdateDramaRequest) (*emptypb.Empty, error) {
	//userId := 1 // todo 获取用户id
	b := business.DramaBusiness{
		Id:              request.Id,
		CategoryId:      request.CategoryId,
		Name:            request.Name,
		Introduce:       request.Introduce,
		Cover:           request.Cover,
		HorizontalCover: request.HorizontalCover,
		TotalCount:      request.TotalCount,
		IsRecommend:     &request.IsRecommend,
		IsNew:           &request.IsNew,
		IsHot:           &request.IsHot,
		IsVisible:       &request.IsVisible,
	}

	_, err := b.Update()
	if err != nil {
		return &emptypb.Empty{}, err
	}
	return &emptypb.Empty{}, nil
}

func (s *DramaServer) Delete(ctx context.Context, request *proto.UpdateDramaRequest) (*emptypb.Empty, error) {
	b := business.DramaBusiness{Id: request.Id}
	_, err := b.Delete()
	if err != nil {
		return &emptypb.Empty{}, err
	}
	return &emptypb.Empty{}, nil
}

func (s *DramaServer) Get(ctx context.Context, request *proto.SearchDramaRequest) (*proto.DramaListResponse, error) {
	b := business.DramaBusiness{
		Keyword: request.Keyword,
		Sort:    request.Sort,
	}
	if request.Page != nil {
		b.Page = request.Page.Page
		b.PerPage = request.Page.PerPage
	}

	model, err := b.List()
	if err != nil {
		return nil, err
	}

	res := &proto.DramaListResponse{}
	res.Total = model.Total
	for _, m := range *model.Dramas {
		var episodes []*proto.EpisodeResponse
		for _, e := range *m.DramaVideos {
			episodes = append(episodes, &proto.EpisodeResponse{
				Id:      e.ID,
				Episode: e.Episode,
				Video: &proto.VideoResponse{
					Id:        e.Video.ID,
					Name:      e.Video.Name,
					Introduce: e.Video.Introduce,
				},
			})
		}

		res.Drama = append(res.Drama, &proto.DramaResponse{
			Id:              m.ID,
			Name:            m.Name,
			Introduce:       m.Introduce,
			Cover:           m.Cover,
			HorizontalCover: m.HorizontalCover,
			Score:           float32(m.Score),
			EpisodeCount:    m.EpisodeCount,
			FavoriteCount:   m.FavoriteCount,
			LikeCount:       m.LikeCount,
			PlayCount:       m.PlayCount,
			BarrageCount:    m.BarrageCount,
			IsNew:           m.IsNew,
			IsHot:           m.IsHot,
			IsEnd:           m.IsEnd,
			Episode:         episodes,
			CreatedTime:     m.CreatedAt.Unix(),
		})
	}

	return res, nil
}

func (s *DramaServer) GetDetail(ctx context.Context, request *proto.SearchDramaRequest) (*proto.DramaResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}
