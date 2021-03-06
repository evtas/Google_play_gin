// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: google_play_grpc/google_play/google_play.protos

package google_play

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GameId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GameId) Reset() {
	*x = GameId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_play_grpc_google_play_google_play_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GameId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GameId) ProtoMessage() {}

func (x *GameId) ProtoReflect() protoreflect.Message {
	mi := &file_google_play_grpc_google_play_google_play_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GameId.ProtoReflect.Descriptor instead.
func (*GameId) Descriptor() ([]byte, []int) {
	return file_google_play_grpc_google_play_google_play_proto_rawDescGZIP(), []int{0}
}

func (x *GameId) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GameDetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Author        string `protobuf:"bytes,3,opt,name=author,proto3" json:"author,omitempty"`
	StarRating    string `protobuf:"bytes,4,opt,name=starRating,proto3" json:"starRating,omitempty"`
	DownloadTimes string `protobuf:"bytes,5,opt,name=downloadTimes,proto3" json:"downloadTimes,omitempty"`
	ContentRating string `protobuf:"bytes,6,opt,name=contentRating,proto3" json:"contentRating,omitempty"`
	Introduction  string `protobuf:"bytes,7,opt,name=introduction,proto3" json:"introduction,omitempty"`
	UpdateTime    string `protobuf:"bytes,8,opt,name=updateTime,proto3" json:"updateTime,omitempty"`
	Genre         string `protobuf:"bytes,9,opt,name=genre,proto3" json:"genre,omitempty"`
}

func (x *GameDetail) Reset() {
	*x = GameDetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_play_grpc_google_play_google_play_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GameDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GameDetail) ProtoMessage() {}

func (x *GameDetail) ProtoReflect() protoreflect.Message {
	mi := &file_google_play_grpc_google_play_google_play_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GameDetail.ProtoReflect.Descriptor instead.
func (*GameDetail) Descriptor() ([]byte, []int) {
	return file_google_play_grpc_google_play_google_play_proto_rawDescGZIP(), []int{1}
}

func (x *GameDetail) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GameDetail) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GameDetail) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

func (x *GameDetail) GetStarRating() string {
	if x != nil {
		return x.StarRating
	}
	return ""
}

func (x *GameDetail) GetDownloadTimes() string {
	if x != nil {
		return x.DownloadTimes
	}
	return ""
}

func (x *GameDetail) GetContentRating() string {
	if x != nil {
		return x.ContentRating
	}
	return ""
}

func (x *GameDetail) GetIntroduction() string {
	if x != nil {
		return x.Introduction
	}
	return ""
}

func (x *GameDetail) GetUpdateTime() string {
	if x != nil {
		return x.UpdateTime
	}
	return ""
}

func (x *GameDetail) GetGenre() string {
	if x != nil {
		return x.Genre
	}
	return ""
}

var File_google_play_grpc_google_play_google_play_proto protoreflect.FileDescriptor

var file_google_play_grpc_google_play_google_play_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5f, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x67, 0x72,
	0x70, 0x63, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5f, 0x70, 0x6c, 0x61, 0x79, 0x2f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5f, 0x70, 0x6c, 0x61, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5f, 0x70, 0x6c, 0x61, 0x79, 0x22, 0x18, 0x0a,
	0x06, 0x47, 0x61, 0x6d, 0x65, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x8e, 0x02, 0x0a, 0x0a, 0x47, 0x61, 0x6d, 0x65,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x52, 0x61, 0x74, 0x69,
	0x6e, 0x67, 0x12, 0x24, 0x0a, 0x0d, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x64, 0x6f, 0x77, 0x6e, 0x6c,
	0x6f, 0x61, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x12, 0x24, 0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x22,
	0x0a, 0x0c, 0x69, 0x6e, 0x74, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x69, 0x6e, 0x74, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x32, 0x4d, 0x0a, 0x0a, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x12, 0x3f, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x47, 0x61, 0x6d,
	0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x13, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x5f, 0x70, 0x6c, 0x61, 0x79, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0x49, 0x64, 0x1a, 0x17, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5f, 0x70, 0x6c, 0x61, 0x79, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x22, 0x00, 0x42, 0x1e, 0x5a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x5f, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x5f, 0x70, 0x6c, 0x61, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_play_grpc_google_play_google_play_proto_rawDescOnce sync.Once
	file_google_play_grpc_google_play_google_play_proto_rawDescData = file_google_play_grpc_google_play_google_play_proto_rawDesc
)

func file_google_play_grpc_google_play_google_play_proto_rawDescGZIP() []byte {
	file_google_play_grpc_google_play_google_play_proto_rawDescOnce.Do(func() {
		file_google_play_grpc_google_play_google_play_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_play_grpc_google_play_google_play_proto_rawDescData)
	})
	return file_google_play_grpc_google_play_google_play_proto_rawDescData
}

var file_google_play_grpc_google_play_google_play_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_google_play_grpc_google_play_google_play_proto_goTypes = []interface{}{
	(*GameId)(nil),     // 0: google_play.GameId
	(*GameDetail)(nil), // 1: google_play.GameDetail
}
var file_google_play_grpc_google_play_google_play_proto_depIdxs = []int32{
	0, // 0: google_play.GooglePlay.GetGameDetail:input_type -> google_play.GameId
	1, // 1: google_play.GooglePlay.GetGameDetail:output_type -> google_play.GameDetail
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_play_grpc_google_play_google_play_proto_init() }
func file_google_play_grpc_google_play_google_play_proto_init() {
	if File_google_play_grpc_google_play_google_play_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_play_grpc_google_play_google_play_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GameId); i {
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
		file_google_play_grpc_google_play_google_play_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GameDetail); i {
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
			RawDescriptor: file_google_play_grpc_google_play_google_play_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_google_play_grpc_google_play_google_play_proto_goTypes,
		DependencyIndexes: file_google_play_grpc_google_play_google_play_proto_depIdxs,
		MessageInfos:      file_google_play_grpc_google_play_google_play_proto_msgTypes,
	}.Build()
	File_google_play_grpc_google_play_google_play_proto = out.File
	file_google_play_grpc_google_play_google_play_proto_rawDesc = nil
	file_google_play_grpc_google_play_google_play_proto_goTypes = nil
	file_google_play_grpc_google_play_google_play_proto_depIdxs = nil
}
