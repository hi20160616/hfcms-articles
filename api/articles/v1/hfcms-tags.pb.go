// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.14.0
// source: api/articles/v1/hfcms-tags.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ListTagsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The parent resource name
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// The maximum number of items to return.
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// The next_page_token value returned from a previous List request, if any.
	PageToken string `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
}

func (x *ListTagsRequest) Reset() {
	*x = ListTagsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_articles_v1_hfcms_tags_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTagsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTagsRequest) ProtoMessage() {}

func (x *ListTagsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_articles_v1_hfcms_tags_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTagsRequest.ProtoReflect.Descriptor instead.
func (*ListTagsRequest) Descriptor() ([]byte, []int) {
	return file_api_articles_v1_hfcms_tags_proto_rawDescGZIP(), []int{0}
}

func (x *ListTagsRequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *ListTagsRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListTagsRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

type ListTagsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tags          []*Tag `protobuf:"bytes,1,rep,name=tags,proto3" json:"tags,omitempty"`
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *ListTagsResponse) Reset() {
	*x = ListTagsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_articles_v1_hfcms_tags_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTagsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTagsResponse) ProtoMessage() {}

func (x *ListTagsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_articles_v1_hfcms_tags_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTagsResponse.ProtoReflect.Descriptor instead.
func (*ListTagsResponse) Descriptor() ([]byte, []int) {
	return file_api_articles_v1_hfcms_tags_proto_rawDescGZIP(), []int{1}
}

func (x *ListTagsResponse) GetTags() []*Tag {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *ListTagsResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

type GetTagRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetTagRequest) Reset() {
	*x = GetTagRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_articles_v1_hfcms_tags_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTagRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTagRequest) ProtoMessage() {}

func (x *GetTagRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_articles_v1_hfcms_tags_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTagRequest.ProtoReflect.Descriptor instead.
func (*GetTagRequest) Descriptor() ([]byte, []int) {
	return file_api_articles_v1_hfcms_tags_proto_rawDescGZIP(), []int{2}
}

func (x *GetTagRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type CreateTagRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Tag  *Tag   `protobuf:"bytes,2,opt,name=tag,proto3" json:"tag,omitempty"`
}

func (x *CreateTagRequest) Reset() {
	*x = CreateTagRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_articles_v1_hfcms_tags_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTagRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTagRequest) ProtoMessage() {}

func (x *CreateTagRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_articles_v1_hfcms_tags_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTagRequest.ProtoReflect.Descriptor instead.
func (*CreateTagRequest) Descriptor() ([]byte, []int) {
	return file_api_articles_v1_hfcms_tags_proto_rawDescGZIP(), []int{3}
}

func (x *CreateTagRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateTagRequest) GetTag() *Tag {
	if x != nil {
		return x.Tag
	}
	return nil
}

type UpdateTagRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tag *Tag `protobuf:"bytes,1,opt,name=tag,proto3" json:"tag,omitempty"`
}

func (x *UpdateTagRequest) Reset() {
	*x = UpdateTagRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_articles_v1_hfcms_tags_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateTagRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTagRequest) ProtoMessage() {}

func (x *UpdateTagRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_articles_v1_hfcms_tags_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTagRequest.ProtoReflect.Descriptor instead.
func (*UpdateTagRequest) Descriptor() ([]byte, []int) {
	return file_api_articles_v1_hfcms_tags_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateTagRequest) GetTag() *Tag {
	if x != nil {
		return x.Tag
	}
	return nil
}

type DeleteTagRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *DeleteTagRequest) Reset() {
	*x = DeleteTagRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_articles_v1_hfcms_tags_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTagRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTagRequest) ProtoMessage() {}

func (x *DeleteTagRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_articles_v1_hfcms_tags_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTagRequest.ProtoReflect.Descriptor instead.
func (*DeleteTagRequest) Descriptor() ([]byte, []int) {
	return file_api_articles_v1_hfcms_tags_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteTagRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Tag struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name       string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	TagId      string                 `protobuf:"bytes,2,opt,name=tag_id,json=tagId,proto3" json:"tag_id,omitempty"`
	TagName    string                 `protobuf:"bytes,3,opt,name=tag_name,json=tagName,proto3" json:"tag_name,omitempty"`
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
}

func (x *Tag) Reset() {
	*x = Tag{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_articles_v1_hfcms_tags_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Tag) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tag) ProtoMessage() {}

func (x *Tag) ProtoReflect() protoreflect.Message {
	mi := &file_api_articles_v1_hfcms_tags_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tag.ProtoReflect.Descriptor instead.
func (*Tag) Descriptor() ([]byte, []int) {
	return file_api_articles_v1_hfcms_tags_proto_rawDescGZIP(), []int{6}
}

func (x *Tag) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Tag) GetTagId() string {
	if x != nil {
		return x.TagId
	}
	return ""
}

func (x *Tag) GetTagName() string {
	if x != nil {
		return x.TagName
	}
	return ""
}

func (x *Tag) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

var File_api_articles_v1_hfcms_tags_proto protoreflect.FileDescriptor

var file_api_articles_v1_hfcms_tags_proto_rawDesc = []byte{
	0x0a, 0x20, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x2f, 0x76,
	0x31, 0x2f, 0x68, 0x66, 0x63, 0x6d, 0x73, 0x2d, 0x74, 0x61, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x11, 0x68, 0x66, 0x63, 0x6d, 0x73, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c,
	0x65, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x65, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x61, 0x67, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09,
	0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x67,
	0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70,
	0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x66, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74,
	0x54, 0x61, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x04,
	0x74, 0x61, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x68, 0x66, 0x63,
	0x6d, 0x73, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x54,
	0x61, 0x67, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74,
	0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x22, 0x23, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x54, 0x61, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x50, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54,
	0x61, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x28, 0x0a,
	0x03, 0x74, 0x61, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x68, 0x66, 0x63,
	0x6d, 0x73, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x54,
	0x61, 0x67, 0x52, 0x03, 0x74, 0x61, 0x67, 0x22, 0x3c, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x54, 0x61, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x03, 0x74,
	0x61, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x68, 0x66, 0x63, 0x6d, 0x73,
	0x2e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x67,
	0x52, 0x03, 0x74, 0x61, 0x67, 0x22, 0x26, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54,
	0x61, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x88, 0x01,
	0x0a, 0x03, 0x54, 0x61, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x15, 0x0a, 0x06, 0x74, 0x61, 0x67,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x61, 0x67, 0x49, 0x64,
	0x12, 0x19, 0x0a, 0x08, 0x74, 0x61, 0x67, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x74, 0x61, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x3b, 0x0a, 0x0b, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x32, 0x91, 0x04, 0x0a, 0x07, 0x54, 0x61, 0x67,
	0x73, 0x41, 0x50, 0x49, 0x12, 0x65, 0x0a, 0x08, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x61, 0x67, 0x73,
	0x12, 0x22, 0x2e, 0x68, 0x66, 0x63, 0x6d, 0x73, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x61, 0x67, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x68, 0x66, 0x63, 0x6d, 0x73, 0x2e, 0x61, 0x72, 0x74,
	0x69, 0x63, 0x6c, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x61, 0x67,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x10, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x0a, 0x12, 0x08, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x61, 0x67, 0x73, 0x12, 0x5d, 0x0a, 0x06, 0x47,
	0x65, 0x74, 0x54, 0x61, 0x67, 0x12, 0x20, 0x2e, 0x68, 0x66, 0x63, 0x6d, 0x73, 0x2e, 0x61, 0x72,
	0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x61, 0x67,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x68, 0x66, 0x63, 0x6d, 0x73, 0x2e,
	0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x67, 0x22,
	0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x12, 0x11, 0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x6e, 0x61,
	0x6d, 0x65, 0x3d, 0x74, 0x61, 0x67, 0x73, 0x2f, 0x2a, 0x7d, 0x12, 0x5f, 0x0a, 0x09, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x67, 0x12, 0x23, 0x2e, 0x68, 0x66, 0x63, 0x6d, 0x73, 0x2e,
	0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x54, 0x61, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x68,
	0x66, 0x63, 0x6d, 0x73, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x54, 0x61, 0x67, 0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x22, 0x08, 0x2f, 0x76,
	0x31, 0x2f, 0x74, 0x61, 0x67, 0x73, 0x3a, 0x03, 0x74, 0x61, 0x67, 0x12, 0x73, 0x0a, 0x09, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x61, 0x67, 0x12, 0x23, 0x2e, 0x68, 0x66, 0x63, 0x6d, 0x73,
	0x2e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x54, 0x61, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x68, 0x66, 0x63, 0x6d, 0x73, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x54, 0x61, 0x67, 0x22, 0x29, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x23, 0x32, 0x1c, 0x2f,
	0x76, 0x31, 0x2f, 0x7b, 0x74, 0x61, 0x67, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x74, 0x61, 0x67,
	0x73, 0x2f, 0x2a, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x7d, 0x3a, 0x03, 0x74, 0x61, 0x67,
	0x12, 0x6a, 0x0a, 0x09, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x61, 0x67, 0x12, 0x23, 0x2e,
	0x68, 0x66, 0x63, 0x6d, 0x73, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x61, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x20, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x1a, 0x2a, 0x18, 0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x74, 0x61,
	0x67, 0x73, 0x2f, 0x2a, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x7d, 0x42, 0x39, 0x5a, 0x37,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x69, 0x32, 0x30, 0x31,
	0x36, 0x30, 0x36, 0x31, 0x36, 0x2f, 0x68, 0x66, 0x63, 0x6d, 0x73, 0x2d, 0x61, 0x72, 0x74, 0x69,
	0x63, 0x6c, 0x65, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65,
	0x73, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_articles_v1_hfcms_tags_proto_rawDescOnce sync.Once
	file_api_articles_v1_hfcms_tags_proto_rawDescData = file_api_articles_v1_hfcms_tags_proto_rawDesc
)

func file_api_articles_v1_hfcms_tags_proto_rawDescGZIP() []byte {
	file_api_articles_v1_hfcms_tags_proto_rawDescOnce.Do(func() {
		file_api_articles_v1_hfcms_tags_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_articles_v1_hfcms_tags_proto_rawDescData)
	})
	return file_api_articles_v1_hfcms_tags_proto_rawDescData
}

var file_api_articles_v1_hfcms_tags_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_api_articles_v1_hfcms_tags_proto_goTypes = []interface{}{
	(*ListTagsRequest)(nil),       // 0: hfcms.articles.v1.ListTagsRequest
	(*ListTagsResponse)(nil),      // 1: hfcms.articles.v1.ListTagsResponse
	(*GetTagRequest)(nil),         // 2: hfcms.articles.v1.GetTagRequest
	(*CreateTagRequest)(nil),      // 3: hfcms.articles.v1.CreateTagRequest
	(*UpdateTagRequest)(nil),      // 4: hfcms.articles.v1.UpdateTagRequest
	(*DeleteTagRequest)(nil),      // 5: hfcms.articles.v1.DeleteTagRequest
	(*Tag)(nil),                   // 6: hfcms.articles.v1.Tag
	(*timestamppb.Timestamp)(nil), // 7: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),         // 8: google.protobuf.Empty
}
var file_api_articles_v1_hfcms_tags_proto_depIdxs = []int32{
	6, // 0: hfcms.articles.v1.ListTagsResponse.tags:type_name -> hfcms.articles.v1.Tag
	6, // 1: hfcms.articles.v1.CreateTagRequest.tag:type_name -> hfcms.articles.v1.Tag
	6, // 2: hfcms.articles.v1.UpdateTagRequest.tag:type_name -> hfcms.articles.v1.Tag
	7, // 3: hfcms.articles.v1.Tag.update_time:type_name -> google.protobuf.Timestamp
	0, // 4: hfcms.articles.v1.TagsAPI.ListTags:input_type -> hfcms.articles.v1.ListTagsRequest
	2, // 5: hfcms.articles.v1.TagsAPI.GetTag:input_type -> hfcms.articles.v1.GetTagRequest
	3, // 6: hfcms.articles.v1.TagsAPI.CreateTag:input_type -> hfcms.articles.v1.CreateTagRequest
	4, // 7: hfcms.articles.v1.TagsAPI.UpdateTag:input_type -> hfcms.articles.v1.UpdateTagRequest
	5, // 8: hfcms.articles.v1.TagsAPI.DeleteTag:input_type -> hfcms.articles.v1.DeleteTagRequest
	1, // 9: hfcms.articles.v1.TagsAPI.ListTags:output_type -> hfcms.articles.v1.ListTagsResponse
	6, // 10: hfcms.articles.v1.TagsAPI.GetTag:output_type -> hfcms.articles.v1.Tag
	6, // 11: hfcms.articles.v1.TagsAPI.CreateTag:output_type -> hfcms.articles.v1.Tag
	6, // 12: hfcms.articles.v1.TagsAPI.UpdateTag:output_type -> hfcms.articles.v1.Tag
	8, // 13: hfcms.articles.v1.TagsAPI.DeleteTag:output_type -> google.protobuf.Empty
	9, // [9:14] is the sub-list for method output_type
	4, // [4:9] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_api_articles_v1_hfcms_tags_proto_init() }
func file_api_articles_v1_hfcms_tags_proto_init() {
	if File_api_articles_v1_hfcms_tags_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_articles_v1_hfcms_tags_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListTagsRequest); i {
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
		file_api_articles_v1_hfcms_tags_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListTagsResponse); i {
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
		file_api_articles_v1_hfcms_tags_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTagRequest); i {
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
		file_api_articles_v1_hfcms_tags_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTagRequest); i {
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
		file_api_articles_v1_hfcms_tags_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateTagRequest); i {
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
		file_api_articles_v1_hfcms_tags_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteTagRequest); i {
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
		file_api_articles_v1_hfcms_tags_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Tag); i {
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
			RawDescriptor: file_api_articles_v1_hfcms_tags_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_articles_v1_hfcms_tags_proto_goTypes,
		DependencyIndexes: file_api_articles_v1_hfcms_tags_proto_depIdxs,
		MessageInfos:      file_api_articles_v1_hfcms_tags_proto_msgTypes,
	}.Build()
	File_api_articles_v1_hfcms_tags_proto = out.File
	file_api_articles_v1_hfcms_tags_proto_rawDesc = nil
	file_api_articles_v1_hfcms_tags_proto_goTypes = nil
	file_api_articles_v1_hfcms_tags_proto_depIdxs = nil
}
