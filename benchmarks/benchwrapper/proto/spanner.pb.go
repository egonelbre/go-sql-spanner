// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: spanner.proto

package proto

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

type Singer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	FirstName  string `protobuf:"bytes,2,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName   string `protobuf:"bytes,3,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	SingerInfo string `protobuf:"bytes,4,opt,name=singer_info,json=singerInfo,proto3" json:"singer_info,omitempty"`
}

func (x *Singer) Reset() {
	*x = Singer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spanner_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Singer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Singer) ProtoMessage() {}

func (x *Singer) ProtoReflect() protoreflect.Message {
	mi := &file_spanner_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Singer.ProtoReflect.Descriptor instead.
func (*Singer) Descriptor() ([]byte, []int) {
	return file_spanner_proto_rawDescGZIP(), []int{0}
}

func (x *Singer) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Singer) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *Singer) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *Singer) GetSingerInfo() string {
	if x != nil {
		return x.SingerInfo
	}
	return ""
}

type Album struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	SingerId   int64  `protobuf:"varint,2,opt,name=singer_id,json=singerId,proto3" json:"singer_id,omitempty"`
	AlbumTitle string `protobuf:"bytes,3,opt,name=album_title,json=albumTitle,proto3" json:"album_title,omitempty"`
}

func (x *Album) Reset() {
	*x = Album{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spanner_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Album) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Album) ProtoMessage() {}

func (x *Album) ProtoReflect() protoreflect.Message {
	mi := &file_spanner_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Album.ProtoReflect.Descriptor instead.
func (*Album) Descriptor() ([]byte, []int) {
	return file_spanner_proto_rawDescGZIP(), []int{1}
}

func (x *Album) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Album) GetSingerId() int64 {
	if x != nil {
		return x.SingerId
	}
	return 0
}

func (x *Album) GetAlbumTitle() string {
	if x != nil {
		return x.AlbumTitle
	}
	return ""
}

type ReadQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The query to use in the read call.
	Query string `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
}

func (x *ReadQuery) Reset() {
	*x = ReadQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spanner_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadQuery) ProtoMessage() {}

func (x *ReadQuery) ProtoReflect() protoreflect.Message {
	mi := &file_spanner_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadQuery.ProtoReflect.Descriptor instead.
func (*ReadQuery) Descriptor() ([]byte, []int) {
	return file_spanner_proto_rawDescGZIP(), []int{2}
}

func (x *ReadQuery) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

type InsertQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The query to use in the insert call.
	Singers []*Singer `protobuf:"bytes,1,rep,name=singers,proto3" json:"singers,omitempty"`
	Albums  []*Album  `protobuf:"bytes,2,rep,name=albums,proto3" json:"albums,omitempty"`
}

func (x *InsertQuery) Reset() {
	*x = InsertQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spanner_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InsertQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InsertQuery) ProtoMessage() {}

func (x *InsertQuery) ProtoReflect() protoreflect.Message {
	mi := &file_spanner_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InsertQuery.ProtoReflect.Descriptor instead.
func (*InsertQuery) Descriptor() ([]byte, []int) {
	return file_spanner_proto_rawDescGZIP(), []int{3}
}

func (x *InsertQuery) GetSingers() []*Singer {
	if x != nil {
		return x.Singers
	}
	return nil
}

func (x *InsertQuery) GetAlbums() []*Album {
	if x != nil {
		return x.Albums
	}
	return nil
}

type UpdateQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The queries to use in the update call.
	Queries []string `protobuf:"bytes,1,rep,name=queries,proto3" json:"queries,omitempty"`
}

func (x *UpdateQuery) Reset() {
	*x = UpdateQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spanner_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateQuery) ProtoMessage() {}

func (x *UpdateQuery) ProtoReflect() protoreflect.Message {
	mi := &file_spanner_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateQuery.ProtoReflect.Descriptor instead.
func (*UpdateQuery) Descriptor() ([]byte, []int) {
	return file_spanner_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateQuery) GetQueries() []string {
	if x != nil {
		return x.Queries
	}
	return nil
}

type EmptyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyResponse) Reset() {
	*x = EmptyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spanner_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyResponse) ProtoMessage() {}

func (x *EmptyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_spanner_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyResponse.ProtoReflect.Descriptor instead.
func (*EmptyResponse) Descriptor() ([]byte, []int) {
	return file_spanner_proto_rawDescGZIP(), []int{5}
}

var File_spanner_proto protoreflect.FileDescriptor

var file_spanner_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x73, 0x70, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0d, 0x73, 0x70, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x5f, 0x62, 0x65, 0x6e, 0x63, 0x68, 0x22, 0x75,
	0x0a, 0x06, 0x53, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73,
	0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69,
	0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x5f, 0x69,
	0x6e, 0x66, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x69, 0x6e, 0x67, 0x65,
	0x72, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x55, 0x0a, 0x05, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b,
	0x0a, 0x09, 0x73, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x08, 0x73, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x61,
	0x6c, 0x62, 0x75, 0x6d, 0x5f, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x22, 0x21, 0x0a, 0x09,
	0x52, 0x65, 0x61, 0x64, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x71, 0x75, 0x65,
	0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x22,
	0x6c, 0x0a, 0x0b, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x2f,
	0x0a, 0x07, 0x73, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x15, 0x2e, 0x73, 0x70, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x5f, 0x62, 0x65, 0x6e, 0x63, 0x68, 0x2e,
	0x53, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x52, 0x07, 0x73, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x73, 0x12,
	0x2c, 0x0a, 0x06, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x73, 0x70, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x5f, 0x62, 0x65, 0x6e, 0x63, 0x68, 0x2e,
	0x41, 0x6c, 0x62, 0x75, 0x6d, 0x52, 0x06, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x73, 0x22, 0x27, 0x0a,
	0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x18, 0x0a, 0x07,
	0x71, 0x75, 0x65, 0x72, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x71,
	0x75, 0x65, 0x72, 0x69, 0x65, 0x73, 0x22, 0x0f, 0x0a, 0x0d, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xe3, 0x01, 0x0a, 0x13, 0x53, 0x70, 0x61, 0x6e,
	0x6e, 0x65, 0x72, 0x42, 0x65, 0x6e, 0x63, 0x68, 0x57, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x12,
	0x40, 0x0a, 0x04, 0x52, 0x65, 0x61, 0x64, 0x12, 0x18, 0x2e, 0x73, 0x70, 0x61, 0x6e, 0x6e, 0x65,
	0x72, 0x5f, 0x62, 0x65, 0x6e, 0x63, 0x68, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x1a, 0x1c, 0x2e, 0x73, 0x70, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x5f, 0x62, 0x65, 0x6e, 0x63,
	0x68, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x44, 0x0a, 0x06, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x12, 0x1a, 0x2e, 0x73, 0x70,
	0x61, 0x6e, 0x6e, 0x65, 0x72, 0x5f, 0x62, 0x65, 0x6e, 0x63, 0x68, 0x2e, 0x49, 0x6e, 0x73, 0x65,
	0x72, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x1a, 0x1c, 0x2e, 0x73, 0x70, 0x61, 0x6e, 0x6e, 0x65,
	0x72, 0x5f, 0x62, 0x65, 0x6e, 0x63, 0x68, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x44, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x12, 0x1a, 0x2e, 0x73, 0x70, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x5f, 0x62, 0x65, 0x6e, 0x63,
	0x68, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x1a, 0x1c, 0x2e,
	0x73, 0x70, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x5f, 0x62, 0x65, 0x6e, 0x63, 0x68, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x11, 0x5a,
	0x0f, 0x2e, 0x2f, 0x73, 0x70, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x5f, 0x62, 0x65, 0x6e, 0x63, 0x68,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_spanner_proto_rawDescOnce sync.Once
	file_spanner_proto_rawDescData = file_spanner_proto_rawDesc
)

func file_spanner_proto_rawDescGZIP() []byte {
	file_spanner_proto_rawDescOnce.Do(func() {
		file_spanner_proto_rawDescData = protoimpl.X.CompressGZIP(file_spanner_proto_rawDescData)
	})
	return file_spanner_proto_rawDescData
}

var file_spanner_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_spanner_proto_goTypes = []interface{}{
	(*Singer)(nil),        // 0: spanner_bench.Singer
	(*Album)(nil),         // 1: spanner_bench.Album
	(*ReadQuery)(nil),     // 2: spanner_bench.ReadQuery
	(*InsertQuery)(nil),   // 3: spanner_bench.InsertQuery
	(*UpdateQuery)(nil),   // 4: spanner_bench.UpdateQuery
	(*EmptyResponse)(nil), // 5: spanner_bench.EmptyResponse
}
var file_spanner_proto_depIdxs = []int32{
	0, // 0: spanner_bench.InsertQuery.singers:type_name -> spanner_bench.Singer
	1, // 1: spanner_bench.InsertQuery.albums:type_name -> spanner_bench.Album
	2, // 2: spanner_bench.SpannerBenchWrapper.Read:input_type -> spanner_bench.ReadQuery
	3, // 3: spanner_bench.SpannerBenchWrapper.Insert:input_type -> spanner_bench.InsertQuery
	4, // 4: spanner_bench.SpannerBenchWrapper.Update:input_type -> spanner_bench.UpdateQuery
	5, // 5: spanner_bench.SpannerBenchWrapper.Read:output_type -> spanner_bench.EmptyResponse
	5, // 6: spanner_bench.SpannerBenchWrapper.Insert:output_type -> spanner_bench.EmptyResponse
	5, // 7: spanner_bench.SpannerBenchWrapper.Update:output_type -> spanner_bench.EmptyResponse
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_spanner_proto_init() }
func file_spanner_proto_init() {
	if File_spanner_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_spanner_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Singer); i {
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
		file_spanner_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Album); i {
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
		file_spanner_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadQuery); i {
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
		file_spanner_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InsertQuery); i {
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
		file_spanner_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateQuery); i {
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
		file_spanner_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmptyResponse); i {
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
			RawDescriptor: file_spanner_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_spanner_proto_goTypes,
		DependencyIndexes: file_spanner_proto_depIdxs,
		MessageInfos:      file_spanner_proto_msgTypes,
	}.Build()
	File_spanner_proto = out.File
	file_spanner_proto_rawDesc = nil
	file_spanner_proto_goTypes = nil
	file_spanner_proto_depIdxs = nil
}