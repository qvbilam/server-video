// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: category.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CategoryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Icon string `protobuf:"bytes,3,opt,name=icon,proto3" json:"icon,omitempty"`
}

func (x *CategoryResponse) Reset() {
	*x = CategoryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_category_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CategoryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CategoryResponse) ProtoMessage() {}

func (x *CategoryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_category_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CategoryResponse.ProtoReflect.Descriptor instead.
func (*CategoryResponse) Descriptor() ([]byte, []int) {
	return file_category_proto_rawDescGZIP(), []int{0}
}

func (x *CategoryResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CategoryResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CategoryResponse) GetIcon() string {
	if x != nil {
		return x.Icon
	}
	return ""
}

type GetCategoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsVisible bool `protobuf:"varint,1,opt,name=isVisible,proto3" json:"isVisible,omitempty"`
}

func (x *GetCategoryRequest) Reset() {
	*x = GetCategoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_category_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCategoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCategoryRequest) ProtoMessage() {}

func (x *GetCategoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_category_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCategoryRequest.ProtoReflect.Descriptor instead.
func (*GetCategoryRequest) Descriptor() ([]byte, []int) {
	return file_category_proto_rawDescGZIP(), []int{1}
}

func (x *GetCategoryRequest) GetIsVisible() bool {
	if x != nil {
		return x.IsVisible
	}
	return false
}

type GetCategoryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Region []*CategoryResponse `protobuf:"bytes,1,rep,name=region,proto3" json:"region,omitempty"`
}

func (x *GetCategoryResponse) Reset() {
	*x = GetCategoryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_category_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCategoryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCategoryResponse) ProtoMessage() {}

func (x *GetCategoryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_category_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCategoryResponse.ProtoReflect.Descriptor instead.
func (*GetCategoryResponse) Descriptor() ([]byte, []int) {
	return file_category_proto_rawDescGZIP(), []int{2}
}

func (x *GetCategoryResponse) GetRegion() []*CategoryResponse {
	if x != nil {
		return x.Region
	}
	return nil
}

type UpdateCategoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Icon      string `protobuf:"bytes,3,opt,name=icon,proto3" json:"icon,omitempty"`
	IsVisible bool   `protobuf:"varint,4,opt,name=isVisible,proto3" json:"isVisible,omitempty"`
}

func (x *UpdateCategoryRequest) Reset() {
	*x = UpdateCategoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_category_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCategoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCategoryRequest) ProtoMessage() {}

func (x *UpdateCategoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_category_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCategoryRequest.ProtoReflect.Descriptor instead.
func (*UpdateCategoryRequest) Descriptor() ([]byte, []int) {
	return file_category_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateCategoryRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateCategoryRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateCategoryRequest) GetIcon() string {
	if x != nil {
		return x.Icon
	}
	return ""
}

func (x *UpdateCategoryRequest) GetIsVisible() bool {
	if x != nil {
		return x.IsVisible
	}
	return false
}

var File_category_proto protoreflect.FileDescriptor

var file_category_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x70, 0x62, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4a, 0x0a, 0x10, 0x43, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x69, 0x63, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69,
	0x63, 0x6f, 0x6e, 0x22, 0x32, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x73, 0x56,
	0x69, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73,
	0x56, 0x69, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x22, 0x49, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x43, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32,
	0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x06, 0x72, 0x65, 0x67, 0x69,
	0x6f, 0x6e, 0x22, 0x6d, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x69, 0x63, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69,
	0x63, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x73, 0x56, 0x69, 0x73, 0x69, 0x62, 0x6c, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x56, 0x69, 0x73, 0x69, 0x62, 0x6c,
	0x65, 0x32, 0x9b, 0x02, 0x0a, 0x08, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x45,
	0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x1f, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f,
	0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x76, 0x69, 0x64, 0x65,
	0x6f, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12,
	0x1f, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x41, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x12, 0x1f, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x42, 0x0a, 0x03, 0x47,
	0x65, 0x74, 0x12, 0x1c, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65,
	0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1d, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x43,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x14, 0x5a, 0x12, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x62, 0x3b,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_category_proto_rawDescOnce sync.Once
	file_category_proto_rawDescData = file_category_proto_rawDesc
)

func file_category_proto_rawDescGZIP() []byte {
	file_category_proto_rawDescOnce.Do(func() {
		file_category_proto_rawDescData = protoimpl.X.CompressGZIP(file_category_proto_rawDescData)
	})
	return file_category_proto_rawDescData
}

var file_category_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_category_proto_goTypes = []interface{}{
	(*CategoryResponse)(nil),      // 0: video.pb.CategoryResponse
	(*GetCategoryRequest)(nil),    // 1: video.pb.GetCategoryRequest
	(*GetCategoryResponse)(nil),   // 2: video.pb.GetCategoryResponse
	(*UpdateCategoryRequest)(nil), // 3: video.pb.UpdateCategoryRequest
	(*emptypb.Empty)(nil),         // 4: google.protobuf.Empty
}
var file_category_proto_depIdxs = []int32{
	0, // 0: video.pb.GetCategoryResponse.region:type_name -> video.pb.CategoryResponse
	3, // 1: video.pb.Category.Create:input_type -> video.pb.UpdateCategoryRequest
	3, // 2: video.pb.Category.Update:input_type -> video.pb.UpdateCategoryRequest
	3, // 3: video.pb.Category.Delete:input_type -> video.pb.UpdateCategoryRequest
	1, // 4: video.pb.Category.Get:input_type -> video.pb.GetCategoryRequest
	0, // 5: video.pb.Category.Create:output_type -> video.pb.CategoryResponse
	4, // 6: video.pb.Category.Update:output_type -> google.protobuf.Empty
	4, // 7: video.pb.Category.Delete:output_type -> google.protobuf.Empty
	2, // 8: video.pb.Category.Get:output_type -> video.pb.GetCategoryResponse
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_category_proto_init() }
func file_category_proto_init() {
	if File_category_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_category_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CategoryResponse); i {
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
		file_category_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCategoryRequest); i {
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
		file_category_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCategoryResponse); i {
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
		file_category_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCategoryRequest); i {
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
			RawDescriptor: file_category_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_category_proto_goTypes,
		DependencyIndexes: file_category_proto_depIdxs,
		MessageInfos:      file_category_proto_msgTypes,
	}.Build()
	File_category_proto = out.File
	file_category_proto_rawDesc = nil
	file_category_proto_goTypes = nil
	file_category_proto_depIdxs = nil
}