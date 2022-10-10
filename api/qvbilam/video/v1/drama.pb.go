// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0-devel
// 	protoc        v3.19.4
// source: drama.proto

package videoV1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
	v1 "video/api/qvbilam/page/v1"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type DramaResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             int64             `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Category       *CategoryResponse `protobuf:"bytes,2,opt,name=category,proto3" json:"category,omitempty"`
	Region         *RegionResponse   `protobuf:"bytes,3,opt,name=region,proto3" json:"region,omitempty"`
	Name           string            `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Introduction   string            `protobuf:"bytes,5,opt,name=introduction,proto3" json:"introduction,omitempty"`
	Icon           string            `protobuf:"bytes,6,opt,name=icon,proto3" json:"icon,omitempty"`
	HorizontalIcon string            `protobuf:"bytes,7,opt,name=horizontalIcon,proto3" json:"horizontalIcon,omitempty"`
	Score          float32           `protobuf:"fixed32,8,opt,name=score,proto3" json:"score,omitempty"`
	EpisodeCount   int64             `protobuf:"varint,9,opt,name=episodeCount,proto3" json:"episodeCount,omitempty"`
	FavoriteCount  int64             `protobuf:"varint,10,opt,name=favoriteCount,proto3" json:"favoriteCount,omitempty"`
	LikeCount      int64             `protobuf:"varint,11,opt,name=likeCount,proto3" json:"likeCount,omitempty"`
	PlayCount      int64             `protobuf:"varint,12,opt,name=playCount,proto3" json:"playCount,omitempty"`
	BarrageCount   int64             `protobuf:"varint,13,opt,name=barrageCount,proto3" json:"barrageCount,omitempty"`
	IsNew          bool              `protobuf:"varint,14,opt,name=is_new,json=isNew,proto3" json:"is_new,omitempty"`
	IsHot          bool              `protobuf:"varint,15,opt,name=is_hot,json=isHot,proto3" json:"is_hot,omitempty"`
	IsEnd          bool              `protobuf:"varint,16,opt,name=is_end,json=isEnd,proto3" json:"is_end,omitempty"`
	IsVisible      bool              `protobuf:"varint,17,opt,name=isVisible,proto3" json:"isVisible,omitempty"`
}

func (x *DramaResponse) Reset() {
	*x = DramaResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_drama_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DramaResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DramaResponse) ProtoMessage() {}

func (x *DramaResponse) ProtoReflect() protoreflect.Message {
	mi := &file_drama_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DramaResponse.ProtoReflect.Descriptor instead.
func (*DramaResponse) Descriptor() ([]byte, []int) {
	return file_drama_proto_rawDescGZIP(), []int{0}
}

func (x *DramaResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *DramaResponse) GetCategory() *CategoryResponse {
	if x != nil {
		return x.Category
	}
	return nil
}

func (x *DramaResponse) GetRegion() *RegionResponse {
	if x != nil {
		return x.Region
	}
	return nil
}

func (x *DramaResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DramaResponse) GetIntroduction() string {
	if x != nil {
		return x.Introduction
	}
	return ""
}

func (x *DramaResponse) GetIcon() string {
	if x != nil {
		return x.Icon
	}
	return ""
}

func (x *DramaResponse) GetHorizontalIcon() string {
	if x != nil {
		return x.HorizontalIcon
	}
	return ""
}

func (x *DramaResponse) GetScore() float32 {
	if x != nil {
		return x.Score
	}
	return 0
}

func (x *DramaResponse) GetEpisodeCount() int64 {
	if x != nil {
		return x.EpisodeCount
	}
	return 0
}

func (x *DramaResponse) GetFavoriteCount() int64 {
	if x != nil {
		return x.FavoriteCount
	}
	return 0
}

func (x *DramaResponse) GetLikeCount() int64 {
	if x != nil {
		return x.LikeCount
	}
	return 0
}

func (x *DramaResponse) GetPlayCount() int64 {
	if x != nil {
		return x.PlayCount
	}
	return 0
}

func (x *DramaResponse) GetBarrageCount() int64 {
	if x != nil {
		return x.BarrageCount
	}
	return 0
}

func (x *DramaResponse) GetIsNew() bool {
	if x != nil {
		return x.IsNew
	}
	return false
}

func (x *DramaResponse) GetIsHot() bool {
	if x != nil {
		return x.IsHot
	}
	return false
}

func (x *DramaResponse) GetIsEnd() bool {
	if x != nil {
		return x.IsEnd
	}
	return false
}

func (x *DramaResponse) GetIsVisible() bool {
	if x != nil {
		return x.IsVisible
	}
	return false
}

type UpdateDramaRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	CategoryId     int64   `protobuf:"varint,2,opt,name=categoryId,proto3" json:"categoryId,omitempty"`
	RegionId       int64   `protobuf:"varint,3,opt,name=regionId,proto3" json:"regionId,omitempty"`
	Name           string  `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Introduction   string  `protobuf:"bytes,5,opt,name=introduction,proto3" json:"introduction,omitempty"`
	Icon           string  `protobuf:"bytes,6,opt,name=icon,proto3" json:"icon,omitempty"`
	HorizontalIcon string  `protobuf:"bytes,7,opt,name=horizontalIcon,proto3" json:"horizontalIcon,omitempty"`
	Score          float32 `protobuf:"fixed32,8,opt,name=score,proto3" json:"score,omitempty"`
	EpisodeCount   int64   `protobuf:"varint,9,opt,name=episodeCount,proto3" json:"episodeCount,omitempty"`
	FavoriteCount  int64   `protobuf:"varint,10,opt,name=favoriteCount,proto3" json:"favoriteCount,omitempty"`
	LikeCount      int64   `protobuf:"varint,11,opt,name=likeCount,proto3" json:"likeCount,omitempty"`
	PlayCount      int64   `protobuf:"varint,12,opt,name=playCount,proto3" json:"playCount,omitempty"`
	BarrageCount   int64   `protobuf:"varint,13,opt,name=barrageCount,proto3" json:"barrageCount,omitempty"`
	IsNew          bool    `protobuf:"varint,14,opt,name=is_new,json=isNew,proto3" json:"is_new,omitempty"`
	IsHot          bool    `protobuf:"varint,15,opt,name=is_hot,json=isHot,proto3" json:"is_hot,omitempty"`
	IsEnd          bool    `protobuf:"varint,16,opt,name=is_end,json=isEnd,proto3" json:"is_end,omitempty"`
	IsVisible      bool    `protobuf:"varint,17,opt,name=isVisible,proto3" json:"isVisible,omitempty"`
	Year           int64   `protobuf:"varint,18,opt,name=year,proto3" json:"year,omitempty"`
	Quarter        int64   `protobuf:"varint,19,opt,name=quarter,proto3" json:"quarter,omitempty"`
}

func (x *UpdateDramaRequest) Reset() {
	*x = UpdateDramaRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_drama_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateDramaRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateDramaRequest) ProtoMessage() {}

func (x *UpdateDramaRequest) ProtoReflect() protoreflect.Message {
	mi := &file_drama_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateDramaRequest.ProtoReflect.Descriptor instead.
func (*UpdateDramaRequest) Descriptor() ([]byte, []int) {
	return file_drama_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateDramaRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateDramaRequest) GetCategoryId() int64 {
	if x != nil {
		return x.CategoryId
	}
	return 0
}

func (x *UpdateDramaRequest) GetRegionId() int64 {
	if x != nil {
		return x.RegionId
	}
	return 0
}

func (x *UpdateDramaRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateDramaRequest) GetIntroduction() string {
	if x != nil {
		return x.Introduction
	}
	return ""
}

func (x *UpdateDramaRequest) GetIcon() string {
	if x != nil {
		return x.Icon
	}
	return ""
}

func (x *UpdateDramaRequest) GetHorizontalIcon() string {
	if x != nil {
		return x.HorizontalIcon
	}
	return ""
}

func (x *UpdateDramaRequest) GetScore() float32 {
	if x != nil {
		return x.Score
	}
	return 0
}

func (x *UpdateDramaRequest) GetEpisodeCount() int64 {
	if x != nil {
		return x.EpisodeCount
	}
	return 0
}

func (x *UpdateDramaRequest) GetFavoriteCount() int64 {
	if x != nil {
		return x.FavoriteCount
	}
	return 0
}

func (x *UpdateDramaRequest) GetLikeCount() int64 {
	if x != nil {
		return x.LikeCount
	}
	return 0
}

func (x *UpdateDramaRequest) GetPlayCount() int64 {
	if x != nil {
		return x.PlayCount
	}
	return 0
}

func (x *UpdateDramaRequest) GetBarrageCount() int64 {
	if x != nil {
		return x.BarrageCount
	}
	return 0
}

func (x *UpdateDramaRequest) GetIsNew() bool {
	if x != nil {
		return x.IsNew
	}
	return false
}

func (x *UpdateDramaRequest) GetIsHot() bool {
	if x != nil {
		return x.IsHot
	}
	return false
}

func (x *UpdateDramaRequest) GetIsEnd() bool {
	if x != nil {
		return x.IsEnd
	}
	return false
}

func (x *UpdateDramaRequest) GetIsVisible() bool {
	if x != nil {
		return x.IsVisible
	}
	return false
}

func (x *UpdateDramaRequest) GetYear() int64 {
	if x != nil {
		return x.Year
	}
	return 0
}

func (x *UpdateDramaRequest) GetQuarter() int64 {
	if x != nil {
		return x.Quarter
	}
	return 0
}

type DramaListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total int64            `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Drama []*DramaResponse `protobuf:"bytes,2,rep,name=drama,proto3" json:"drama,omitempty"`
}

func (x *DramaListResponse) Reset() {
	*x = DramaListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_drama_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DramaListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DramaListResponse) ProtoMessage() {}

func (x *DramaListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_drama_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DramaListResponse.ProtoReflect.Descriptor instead.
func (*DramaListResponse) Descriptor() ([]byte, []int) {
	return file_drama_proto_rawDescGZIP(), []int{2}
}

func (x *DramaListResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *DramaListResponse) GetDrama() []*DramaResponse {
	if x != nil {
		return x.Drama
	}
	return nil
}

type SearchDramaRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        []int64         `protobuf:"varint,1,rep,packed,name=id,proto3" json:"id,omitempty"`
	Keyword   string          `protobuf:"bytes,2,opt,name=keyword,proto3" json:"keyword,omitempty"`
	Sort      string          `protobuf:"bytes,3,opt,name=sort,proto3" json:"sort,omitempty"`
	Year      int64           `protobuf:"varint,5,opt,name=year,proto3" json:"year,omitempty"`
	Quarter   int64           `protobuf:"varint,6,opt,name=quarter,proto3" json:"quarter,omitempty"`
	IsVisible bool            `protobuf:"varint,7,opt,name=isVisible,proto3" json:"isVisible,omitempty"`
	Page      *v1.PageRequest `protobuf:"bytes,8,opt,name=page,proto3" json:"page,omitempty"`
}

func (x *SearchDramaRequest) Reset() {
	*x = SearchDramaRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_drama_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchDramaRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchDramaRequest) ProtoMessage() {}

func (x *SearchDramaRequest) ProtoReflect() protoreflect.Message {
	mi := &file_drama_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchDramaRequest.ProtoReflect.Descriptor instead.
func (*SearchDramaRequest) Descriptor() ([]byte, []int) {
	return file_drama_proto_rawDescGZIP(), []int{3}
}

func (x *SearchDramaRequest) GetId() []int64 {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *SearchDramaRequest) GetKeyword() string {
	if x != nil {
		return x.Keyword
	}
	return ""
}

func (x *SearchDramaRequest) GetSort() string {
	if x != nil {
		return x.Sort
	}
	return ""
}

func (x *SearchDramaRequest) GetYear() int64 {
	if x != nil {
		return x.Year
	}
	return 0
}

func (x *SearchDramaRequest) GetQuarter() int64 {
	if x != nil {
		return x.Quarter
	}
	return 0
}

func (x *SearchDramaRequest) GetIsVisible() bool {
	if x != nil {
		return x.IsVisible
	}
	return false
}

func (x *SearchDramaRequest) GetPage() *v1.PageRequest {
	if x != nil {
		return x.Page
	}
	return nil
}

var File_drama_proto protoreflect.FileDescriptor

var file_drama_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x64, 0x72, 0x61, 0x6d, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x76,
	0x69, 0x64, 0x65, 0x6f, 0x2e, 0x70, 0x62, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0e, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1e, 0x61, 0x70, 0x69, 0x2f, 0x71, 0x76, 0x62, 0x69, 0x6c, 0x61, 0x6d, 0x2f,
	0x70, 0x61, 0x67, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xa0, 0x04, 0x0a, 0x0d, 0x44, 0x72, 0x61, 0x6d, 0x61, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x36, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x70,
	0x62, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x30, 0x0a, 0x06,
	0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x76,
	0x69, 0x64, 0x65, 0x6f, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x69, 0x6e, 0x74, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x69, 0x6e, 0x74, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x63, 0x6f, 0x6e, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x63, 0x6f, 0x6e, 0x12, 0x26, 0x0a, 0x0e, 0x68, 0x6f,
	0x72, 0x69, 0x7a, 0x6f, 0x6e, 0x74, 0x61, 0x6c, 0x49, 0x63, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0e, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x6f, 0x6e, 0x74, 0x61, 0x6c, 0x49, 0x63,
	0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x65, 0x70, 0x69, 0x73,
	0x6f, 0x64, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c,
	0x65, 0x70, 0x69, 0x73, 0x6f, 0x64, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x24, 0x0a, 0x0d,
	0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0d, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x69, 0x6b, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x6c, 0x69, 0x6b, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x1c, 0x0a, 0x09, 0x70, 0x6c, 0x61, 0x79, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x09, 0x70, 0x6c, 0x61, 0x79, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x22,
	0x0a, 0x0c, 0x62, 0x61, 0x72, 0x72, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0d,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x62, 0x61, 0x72, 0x72, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x69, 0x73, 0x5f, 0x6e, 0x65, 0x77, 0x18, 0x0e, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x05, 0x69, 0x73, 0x4e, 0x65, 0x77, 0x12, 0x15, 0x0a, 0x06, 0x69, 0x73, 0x5f,
	0x68, 0x6f, 0x74, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x69, 0x73, 0x48, 0x6f, 0x74,
	0x12, 0x15, 0x0a, 0x06, 0x69, 0x73, 0x5f, 0x65, 0x6e, 0x64, 0x18, 0x10, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x05, 0x69, 0x73, 0x45, 0x6e, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x73, 0x56, 0x69, 0x73,
	0x69, 0x62, 0x6c, 0x65, 0x18, 0x11, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x56, 0x69,
	0x73, 0x69, 0x62, 0x6c, 0x65, 0x22, 0xa5, 0x04, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x44, 0x72, 0x61, 0x6d, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1e, 0x0a, 0x0a,
	0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08,
	0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08,
	0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c,
	0x69, 0x6e, 0x74, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x69, 0x6e, 0x74, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x12, 0x0a, 0x04, 0x69, 0x63, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x69, 0x63, 0x6f, 0x6e, 0x12, 0x26, 0x0a, 0x0e, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x6f, 0x6e, 0x74,
	0x61, 0x6c, 0x49, 0x63, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x68, 0x6f,
	0x72, 0x69, 0x7a, 0x6f, 0x6e, 0x74, 0x61, 0x6c, 0x49, 0x63, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05,
	0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x73, 0x63, 0x6f,
	0x72, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x65, 0x70, 0x69, 0x73, 0x6f, 0x64, 0x65, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x65, 0x70, 0x69, 0x73, 0x6f, 0x64,
	0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69,
	0x74, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x66,
	0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09,
	0x6c, 0x69, 0x6b, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x6c, 0x69, 0x6b, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x6c,
	0x61, 0x79, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x70,
	0x6c, 0x61, 0x79, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x62, 0x61, 0x72, 0x72,
	0x61, 0x67, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c,
	0x62, 0x61, 0x72, 0x72, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x15, 0x0a, 0x06,
	0x69, 0x73, 0x5f, 0x6e, 0x65, 0x77, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x69, 0x73,
	0x4e, 0x65, 0x77, 0x12, 0x15, 0x0a, 0x06, 0x69, 0x73, 0x5f, 0x68, 0x6f, 0x74, 0x18, 0x0f, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x05, 0x69, 0x73, 0x48, 0x6f, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x69, 0x73,
	0x5f, 0x65, 0x6e, 0x64, 0x18, 0x10, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x69, 0x73, 0x45, 0x6e,
	0x64, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x73, 0x56, 0x69, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x18, 0x11,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x56, 0x69, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x79, 0x65, 0x61, 0x72, 0x18, 0x12, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x79,
	0x65, 0x61, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x71, 0x75, 0x61, 0x72, 0x74, 0x65, 0x72, 0x18, 0x13,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x71, 0x75, 0x61, 0x72, 0x74, 0x65, 0x72, 0x22, 0x58, 0x0a,
	0x11, 0x44, 0x72, 0x61, 0x6d, 0x61, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x2d, 0x0a, 0x05, 0x64, 0x72, 0x61, 0x6d,
	0x61, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e,
	0x70, 0x62, 0x2e, 0x44, 0x72, 0x61, 0x6d, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x52, 0x05, 0x64, 0x72, 0x61, 0x6d, 0x61, 0x22, 0xca, 0x01, 0x0a, 0x12, 0x53, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x44, 0x72, 0x61, 0x6d, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18,
	0x0a, 0x07, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x79, 0x65, 0x61, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x79, 0x65, 0x61, 0x72,
	0x12, 0x18, 0x0a, 0x07, 0x71, 0x75, 0x61, 0x72, 0x74, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x07, 0x71, 0x75, 0x61, 0x72, 0x74, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x73,
	0x56, 0x69, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69,
	0x73, 0x56, 0x69, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x12, 0x2a, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70, 0x61, 0x67, 0x65, 0x50, 0x62, 0x2e,
	0x76, 0x31, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x04,
	0x70, 0x61, 0x67, 0x65, 0x32, 0xce, 0x02, 0x0a, 0x05, 0x44, 0x72, 0x61, 0x6d, 0x61, 0x12, 0x3f,
	0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x1c, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f,
	0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x72, 0x61, 0x6d, 0x61, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x70,
	0x62, 0x2e, 0x44, 0x72, 0x61, 0x6d, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x3e, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1c, 0x2e, 0x76, 0x69, 0x64, 0x65,
	0x6f, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x72, 0x61, 0x6d, 0x61,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12,
	0x3e, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x1c, 0x2e, 0x76, 0x69, 0x64, 0x65,
	0x6f, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x72, 0x61, 0x6d, 0x61,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12,
	0x42, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x1c, 0x2e, 0x76,
	0x69, 0x64, 0x65, 0x6f, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x44, 0x72,
	0x61, 0x6d, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x76, 0x69, 0x64,
	0x65, 0x6f, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x72, 0x61, 0x6d, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x1c, 0x2e, 0x76, 0x69, 0x64,
	0x65, 0x6f, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x44, 0x72, 0x61, 0x6d,
	0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f,
	0x2e, 0x70, 0x62, 0x2e, 0x44, 0x72, 0x61, 0x6d, 0x61, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x24, 0x5a, 0x22, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x71, 0x76, 0x62, 0x69, 0x6c, 0x61, 0x6d, 0x2f, 0x76, 0x69, 0x64, 0x65, 0x6f,
	0x2f, 0x76, 0x31, 0x3b, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_drama_proto_rawDescOnce sync.Once
	file_drama_proto_rawDescData = file_drama_proto_rawDesc
)

func file_drama_proto_rawDescGZIP() []byte {
	file_drama_proto_rawDescOnce.Do(func() {
		file_drama_proto_rawDescData = protoimpl.X.CompressGZIP(file_drama_proto_rawDescData)
	})
	return file_drama_proto_rawDescData
}

var file_drama_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_drama_proto_goTypes = []interface{}{
	(*DramaResponse)(nil),      // 0: video.pb.DramaResponse
	(*UpdateDramaRequest)(nil), // 1: video.pb.UpdateDramaRequest
	(*DramaListResponse)(nil),  // 2: video.pb.DramaListResponse
	(*SearchDramaRequest)(nil), // 3: video.pb.SearchDramaRequest
	(*CategoryResponse)(nil),   // 4: video.pb.CategoryResponse
	(*RegionResponse)(nil),     // 5: video.pb.RegionResponse
	(*v1.PageRequest)(nil),     // 6: pagePb.v1.PageRequest
	(*emptypb.Empty)(nil),      // 7: google.protobuf.Empty
}
var file_drama_proto_depIdxs = []int32{
	4, // 0: video.pb.DramaResponse.category:type_name -> video.pb.CategoryResponse
	5, // 1: video.pb.DramaResponse.region:type_name -> video.pb.RegionResponse
	0, // 2: video.pb.DramaListResponse.drama:type_name -> video.pb.DramaResponse
	6, // 3: video.pb.SearchDramaRequest.page:type_name -> pagePb.v1.PageRequest
	1, // 4: video.pb.Drama.Create:input_type -> video.pb.UpdateDramaRequest
	1, // 5: video.pb.Drama.Update:input_type -> video.pb.UpdateDramaRequest
	1, // 6: video.pb.Drama.Delete:input_type -> video.pb.UpdateDramaRequest
	3, // 7: video.pb.Drama.GetDetail:input_type -> video.pb.SearchDramaRequest
	3, // 8: video.pb.Drama.Get:input_type -> video.pb.SearchDramaRequest
	0, // 9: video.pb.Drama.Create:output_type -> video.pb.DramaResponse
	7, // 10: video.pb.Drama.Update:output_type -> google.protobuf.Empty
	7, // 11: video.pb.Drama.Delete:output_type -> google.protobuf.Empty
	0, // 12: video.pb.Drama.GetDetail:output_type -> video.pb.DramaResponse
	2, // 13: video.pb.Drama.Get:output_type -> video.pb.DramaListResponse
	9, // [9:14] is the sub-list for method output_type
	4, // [4:9] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_drama_proto_init() }
func file_drama_proto_init() {
	if File_drama_proto != nil {
		return
	}
	file_category_proto_init()
	file_region_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_drama_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DramaResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_drama_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateDramaRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_drama_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DramaListResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_drama_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchDramaRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_drama_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_drama_proto_goTypes,
		DependencyIndexes: file_drama_proto_depIdxs,
		MessageInfos:      file_drama_proto_msgTypes,
	}.Build()
	File_drama_proto = out.File
	file_drama_proto_rawDesc = nil
	file_drama_proto_goTypes = nil
	file_drama_proto_depIdxs = nil
}
