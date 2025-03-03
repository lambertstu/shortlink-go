// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        v3.19.4
// source: shortlink.proto

package shortlink

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

type ShortLinkCreateRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Domain        string                 `protobuf:"bytes,1,opt,name=domain,proto3" json:"domain,omitempty"`                                       // 域名
	OriginUrl     string                 `protobuf:"bytes,2,opt,name=origin_url,json=originUrl,proto3" json:"origin_url,omitempty"`                // 原始链接
	Gid           string                 `protobuf:"bytes,3,opt,name=gid,proto3" json:"gid,omitempty"`                                             // 分组标识
	ValidDateType int32                  `protobuf:"varint,4,opt,name=valid_date_type,json=validDateType,proto3" json:"valid_date_type,omitempty"` // 有效期类型 0：永久有效 1：自定义
	ValidDate     int64                  `protobuf:"varint,5,opt,name=valid_date,json=validDate,proto3" json:"valid_date,omitempty"`               // 有效期（时间戳格式，单位：秒）
	Describe      string                 `protobuf:"bytes,6,opt,name=describe,proto3" json:"describe,omitempty"`                                   // 描述
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ShortLinkCreateRequest) Reset() {
	*x = ShortLinkCreateRequest{}
	mi := &file_shortlink_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ShortLinkCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShortLinkCreateRequest) ProtoMessage() {}

func (x *ShortLinkCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_shortlink_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShortLinkCreateRequest.ProtoReflect.Descriptor instead.
func (*ShortLinkCreateRequest) Descriptor() ([]byte, []int) {
	return file_shortlink_proto_rawDescGZIP(), []int{0}
}

func (x *ShortLinkCreateRequest) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *ShortLinkCreateRequest) GetOriginUrl() string {
	if x != nil {
		return x.OriginUrl
	}
	return ""
}

func (x *ShortLinkCreateRequest) GetGid() string {
	if x != nil {
		return x.Gid
	}
	return ""
}

func (x *ShortLinkCreateRequest) GetValidDateType() int32 {
	if x != nil {
		return x.ValidDateType
	}
	return 0
}

func (x *ShortLinkCreateRequest) GetValidDate() int64 {
	if x != nil {
		return x.ValidDate
	}
	return 0
}

func (x *ShortLinkCreateRequest) GetDescribe() string {
	if x != nil {
		return x.Describe
	}
	return ""
}

type ShortLinkCreateResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Domain        string                 `protobuf:"bytes,2,opt,name=domain,proto3" json:"domain,omitempty"`                        // 域名
	OriginUrl     string                 `protobuf:"bytes,3,opt,name=origin_url,json=originUrl,proto3" json:"origin_url,omitempty"` // 原始链接
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ShortLinkCreateResponse) Reset() {
	*x = ShortLinkCreateResponse{}
	mi := &file_shortlink_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ShortLinkCreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShortLinkCreateResponse) ProtoMessage() {}

func (x *ShortLinkCreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_shortlink_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShortLinkCreateResponse.ProtoReflect.Descriptor instead.
func (*ShortLinkCreateResponse) Descriptor() ([]byte, []int) {
	return file_shortlink_proto_rawDescGZIP(), []int{1}
}

func (x *ShortLinkCreateResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *ShortLinkCreateResponse) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *ShortLinkCreateResponse) GetOriginUrl() string {
	if x != nil {
		return x.OriginUrl
	}
	return ""
}

type ShortLinkUpdateResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ShortLinkUpdateResponse) Reset() {
	*x = ShortLinkUpdateResponse{}
	mi := &file_shortlink_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ShortLinkUpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShortLinkUpdateResponse) ProtoMessage() {}

func (x *ShortLinkUpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_shortlink_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShortLinkUpdateResponse.ProtoReflect.Descriptor instead.
func (*ShortLinkUpdateResponse) Descriptor() ([]byte, []int) {
	return file_shortlink_proto_rawDescGZIP(), []int{2}
}

func (x *ShortLinkUpdateResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type ShortLinkUpdateRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Domain        string                 `protobuf:"bytes,1,opt,name=domain,proto3" json:"domain,omitempty"`                                       // 域名
	OriginUrl     string                 `protobuf:"bytes,2,opt,name=origin_url,json=originUrl,proto3" json:"origin_url,omitempty"`                // 原始链接
	Gid           string                 `protobuf:"bytes,3,opt,name=gid,proto3" json:"gid,omitempty"`                                             // 分组标识
	ValidDateType int32                  `protobuf:"varint,4,opt,name=valid_date_type,json=validDateType,proto3" json:"valid_date_type,omitempty"` // 有效期类型 0：永久有效 1：自定义
	ValidDate     int64                  `protobuf:"varint,5,opt,name=valid_date,json=validDate,proto3" json:"valid_date,omitempty"`               // 有效期（时间戳格式，单位：秒）
	Describe      string                 `protobuf:"bytes,6,opt,name=describe,proto3" json:"describe,omitempty"`                                   // 描述
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ShortLinkUpdateRequest) Reset() {
	*x = ShortLinkUpdateRequest{}
	mi := &file_shortlink_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ShortLinkUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShortLinkUpdateRequest) ProtoMessage() {}

func (x *ShortLinkUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_shortlink_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShortLinkUpdateRequest.ProtoReflect.Descriptor instead.
func (*ShortLinkUpdateRequest) Descriptor() ([]byte, []int) {
	return file_shortlink_proto_rawDescGZIP(), []int{3}
}

func (x *ShortLinkUpdateRequest) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *ShortLinkUpdateRequest) GetOriginUrl() string {
	if x != nil {
		return x.OriginUrl
	}
	return ""
}

func (x *ShortLinkUpdateRequest) GetGid() string {
	if x != nil {
		return x.Gid
	}
	return ""
}

func (x *ShortLinkUpdateRequest) GetValidDateType() int32 {
	if x != nil {
		return x.ValidDateType
	}
	return 0
}

func (x *ShortLinkUpdateRequest) GetValidDate() int64 {
	if x != nil {
		return x.ValidDate
	}
	return 0
}

func (x *ShortLinkUpdateRequest) GetDescribe() string {
	if x != nil {
		return x.Describe
	}
	return ""
}

type ShortLinkBatchCreateRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	OriginUrls    []string               `protobuf:"bytes,1,rep,name=origin_urls,json=originUrls,proto3" json:"origin_urls,omitempty"`             // 原始链接集合
	Describes     []string               `protobuf:"bytes,2,rep,name=describes,proto3" json:"describes,omitempty"`                                 // 描述集合
	Gid           string                 `protobuf:"bytes,3,opt,name=gid,proto3" json:"gid,omitempty"`                                             // 分组标识
	ValidDateType int32                  `protobuf:"varint,4,opt,name=valid_date_type,json=validDateType,proto3" json:"valid_date_type,omitempty"` // 有效期类型 0：永久有效 1：自定义
	ValidDate     int64                  `protobuf:"varint,5,opt,name=valid_date,json=validDate,proto3" json:"valid_date,omitempty"`               // 有效期（时间戳格式，单位：秒）
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ShortLinkBatchCreateRequest) Reset() {
	*x = ShortLinkBatchCreateRequest{}
	mi := &file_shortlink_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ShortLinkBatchCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShortLinkBatchCreateRequest) ProtoMessage() {}

func (x *ShortLinkBatchCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_shortlink_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShortLinkBatchCreateRequest.ProtoReflect.Descriptor instead.
func (*ShortLinkBatchCreateRequest) Descriptor() ([]byte, []int) {
	return file_shortlink_proto_rawDescGZIP(), []int{4}
}

func (x *ShortLinkBatchCreateRequest) GetOriginUrls() []string {
	if x != nil {
		return x.OriginUrls
	}
	return nil
}

func (x *ShortLinkBatchCreateRequest) GetDescribes() []string {
	if x != nil {
		return x.Describes
	}
	return nil
}

func (x *ShortLinkBatchCreateRequest) GetGid() string {
	if x != nil {
		return x.Gid
	}
	return ""
}

func (x *ShortLinkBatchCreateRequest) GetValidDateType() int32 {
	if x != nil {
		return x.ValidDateType
	}
	return 0
}

func (x *ShortLinkBatchCreateRequest) GetValidDate() int64 {
	if x != nil {
		return x.ValidDate
	}
	return 0
}

type ShortLinkPageRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Pid           string                 `protobuf:"bytes,1,opt,name=pid,proto3" json:"pid,omitempty"`
	OrderTag      string                 `protobuf:"bytes,2,opt,name=orderTag,proto3" json:"orderTag,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ShortLinkPageRequest) Reset() {
	*x = ShortLinkPageRequest{}
	mi := &file_shortlink_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ShortLinkPageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShortLinkPageRequest) ProtoMessage() {}

func (x *ShortLinkPageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_shortlink_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShortLinkPageRequest.ProtoReflect.Descriptor instead.
func (*ShortLinkPageRequest) Descriptor() ([]byte, []int) {
	return file_shortlink_proto_rawDescGZIP(), []int{5}
}

func (x *ShortLinkPageRequest) GetPid() string {
	if x != nil {
		return x.Pid
	}
	return ""
}

func (x *ShortLinkPageRequest) GetOrderTag() string {
	if x != nil {
		return x.OrderTag
	}
	return ""
}

type ShortLinkPageResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Domain        string                 `protobuf:"bytes,2,opt,name=domain,proto3" json:"domain,omitempty"`                                       // 域名
	ShortUri      string                 `protobuf:"bytes,3,opt,name=short_uri,json=shortUri,proto3" json:"short_uri,omitempty"`                   // 短链接
	FullShortUrl  string                 `protobuf:"bytes,4,opt,name=full_short_url,json=fullShortUrl,proto3" json:"full_short_url,omitempty"`     // 完整短链接
	OriginUrl     string                 `protobuf:"bytes,5,opt,name=origin_url,json=originUrl,proto3" json:"origin_url,omitempty"`                // 原始链接
	Gid           string                 `protobuf:"bytes,6,opt,name=gid,proto3" json:"gid,omitempty"`                                             // 分组标识
	ValidDateType int32                  `protobuf:"varint,7,opt,name=valid_date_type,json=validDateType,proto3" json:"valid_date_type,omitempty"` // 有效期类型 0：永久有效 1：自定义
	EnableStatus  int32                  `protobuf:"varint,8,opt,name=enable_status,json=enableStatus,proto3" json:"enable_status,omitempty"`      // 启用标识 0：启用 1：未启用
	ValidDate     int64                  `protobuf:"varint,9,opt,name=valid_date,json=validDate,proto3" json:"valid_date,omitempty"`               // 有效期（时间戳格式，单位：秒）
	CreateTime    int64                  `protobuf:"varint,10,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`           // 创建时间（时间戳格式，单位：秒）
	Describe      string                 `protobuf:"bytes,11,opt,name=describe,proto3" json:"describe,omitempty"`                                  // 描述
	Favicon       string                 `protobuf:"bytes,12,opt,name=favicon,proto3" json:"favicon,omitempty"`                                    // 网站标识
	TotalPv       int32                  `protobuf:"varint,13,opt,name=total_pv,json=totalPv,proto3" json:"total_pv,omitempty"`                    // 历史PV
	TodayPv       int32                  `protobuf:"varint,14,opt,name=today_pv,json=todayPv,proto3" json:"today_pv,omitempty"`                    // 今日PV
	TotalUv       int32                  `protobuf:"varint,15,opt,name=total_uv,json=totalUv,proto3" json:"total_uv,omitempty"`                    // 历史UV
	TodayUv       int32                  `protobuf:"varint,16,opt,name=today_uv,json=todayUv,proto3" json:"today_uv,omitempty"`                    // 今日UV
	TotalUip      int32                  `protobuf:"varint,17,opt,name=total_uip,json=totalUip,proto3" json:"total_uip,omitempty"`                 // 历史UIP
	TodayUip      int32                  `protobuf:"varint,18,opt,name=today_uip,json=todayUip,proto3" json:"today_uip,omitempty"`                 // 今日UIP
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ShortLinkPageResponse) Reset() {
	*x = ShortLinkPageResponse{}
	mi := &file_shortlink_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ShortLinkPageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShortLinkPageResponse) ProtoMessage() {}

func (x *ShortLinkPageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_shortlink_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShortLinkPageResponse.ProtoReflect.Descriptor instead.
func (*ShortLinkPageResponse) Descriptor() ([]byte, []int) {
	return file_shortlink_proto_rawDescGZIP(), []int{6}
}

func (x *ShortLinkPageResponse) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *ShortLinkPageResponse) GetShortUri() string {
	if x != nil {
		return x.ShortUri
	}
	return ""
}

func (x *ShortLinkPageResponse) GetFullShortUrl() string {
	if x != nil {
		return x.FullShortUrl
	}
	return ""
}

func (x *ShortLinkPageResponse) GetOriginUrl() string {
	if x != nil {
		return x.OriginUrl
	}
	return ""
}

func (x *ShortLinkPageResponse) GetGid() string {
	if x != nil {
		return x.Gid
	}
	return ""
}

func (x *ShortLinkPageResponse) GetValidDateType() int32 {
	if x != nil {
		return x.ValidDateType
	}
	return 0
}

func (x *ShortLinkPageResponse) GetEnableStatus() int32 {
	if x != nil {
		return x.EnableStatus
	}
	return 0
}

func (x *ShortLinkPageResponse) GetValidDate() int64 {
	if x != nil {
		return x.ValidDate
	}
	return 0
}

func (x *ShortLinkPageResponse) GetCreateTime() int64 {
	if x != nil {
		return x.CreateTime
	}
	return 0
}

func (x *ShortLinkPageResponse) GetDescribe() string {
	if x != nil {
		return x.Describe
	}
	return ""
}

func (x *ShortLinkPageResponse) GetFavicon() string {
	if x != nil {
		return x.Favicon
	}
	return ""
}

func (x *ShortLinkPageResponse) GetTotalPv() int32 {
	if x != nil {
		return x.TotalPv
	}
	return 0
}

func (x *ShortLinkPageResponse) GetTodayPv() int32 {
	if x != nil {
		return x.TodayPv
	}
	return 0
}

func (x *ShortLinkPageResponse) GetTotalUv() int32 {
	if x != nil {
		return x.TotalUv
	}
	return 0
}

func (x *ShortLinkPageResponse) GetTodayUv() int32 {
	if x != nil {
		return x.TodayUv
	}
	return 0
}

func (x *ShortLinkPageResponse) GetTotalUip() int32 {
	if x != nil {
		return x.TotalUip
	}
	return 0
}

func (x *ShortLinkPageResponse) GetTodayUip() int32 {
	if x != nil {
		return x.TodayUip
	}
	return 0
}

type ListGroupShortLinkCountRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Gid           string                 `protobuf:"bytes,1,opt,name=gid,proto3" json:"gid,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListGroupShortLinkCountRequest) Reset() {
	*x = ListGroupShortLinkCountRequest{}
	mi := &file_shortlink_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListGroupShortLinkCountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListGroupShortLinkCountRequest) ProtoMessage() {}

func (x *ListGroupShortLinkCountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_shortlink_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListGroupShortLinkCountRequest.ProtoReflect.Descriptor instead.
func (*ListGroupShortLinkCountRequest) Descriptor() ([]byte, []int) {
	return file_shortlink_proto_rawDescGZIP(), []int{7}
}

func (x *ListGroupShortLinkCountRequest) GetGid() string {
	if x != nil {
		return x.Gid
	}
	return ""
}

type ListGroupShortLinkCountResponse struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	Gid            string                 `protobuf:"bytes,1,opt,name=gid,proto3" json:"gid,omitempty"`
	ShortlinkCount int32                  `protobuf:"varint,2,opt,name=shortlinkCount,proto3" json:"shortlinkCount,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *ListGroupShortLinkCountResponse) Reset() {
	*x = ListGroupShortLinkCountResponse{}
	mi := &file_shortlink_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListGroupShortLinkCountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListGroupShortLinkCountResponse) ProtoMessage() {}

func (x *ListGroupShortLinkCountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_shortlink_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListGroupShortLinkCountResponse.ProtoReflect.Descriptor instead.
func (*ListGroupShortLinkCountResponse) Descriptor() ([]byte, []int) {
	return file_shortlink_proto_rawDescGZIP(), []int{8}
}

func (x *ListGroupShortLinkCountResponse) GetGid() string {
	if x != nil {
		return x.Gid
	}
	return ""
}

func (x *ListGroupShortLinkCountResponse) GetShortlinkCount() int32 {
	if x != nil {
		return x.ShortlinkCount
	}
	return 0
}

var File_shortlink_proto protoreflect.FileDescriptor

var file_shortlink_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x6c, 0x69, 0x6e, 0x6b, 0x22, 0xc4, 0x01, 0x0a,
	0x16, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12,
	0x1d, 0x0a, 0x0a, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x55, 0x72, 0x6c, 0x12, 0x10,
	0x0a, 0x03, 0x67, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x67, 0x69, 0x64,
	0x12, 0x26, 0x0a, 0x0f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x44, 0x61, 0x74, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x44, 0x61, 0x74, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x62, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x62, 0x65, 0x22, 0x6a, 0x0a, 0x17, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x4c, 0x69, 0x6e, 0x6b,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x12, 0x1d, 0x0a, 0x0a, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x55, 0x72, 0x6c, 0x22,
	0x33, 0x0a, 0x17, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x22, 0xc4, 0x01, 0x0a, 0x16, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x4c, 0x69,
	0x6e, 0x6b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x6f, 0x72, 0x69, 0x67, 0x69,
	0x6e, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6f, 0x72, 0x69,
	0x67, 0x69, 0x6e, 0x55, 0x72, 0x6c, 0x12, 0x10, 0x0a, 0x03, 0x67, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x67, 0x69, 0x64, 0x12, 0x26, 0x0a, 0x0f, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0d, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x44, 0x61, 0x74, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x1d, 0x0a, 0x0a, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x44, 0x61, 0x74, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x22, 0xb5, 0x01, 0x0a, 0x1b,
	0x53, 0x68, 0x6f, 0x72, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x42, 0x61, 0x74, 0x63, 0x68, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x6f,
	0x72, 0x69, 0x67, 0x69, 0x6e, 0x5f, 0x75, 0x72, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x0a, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x55, 0x72, 0x6c, 0x73, 0x12, 0x1c, 0x0a, 0x09,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x09, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x67, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x67, 0x69, 0x64, 0x12, 0x26, 0x0a, 0x0f,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x44, 0x61, 0x74, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x5f, 0x64, 0x61,
	0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x44,
	0x61, 0x74, 0x65, 0x22, 0x44, 0x0a, 0x14, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x4c, 0x69, 0x6e, 0x6b,
	0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x70,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x70, 0x69, 0x64, 0x12, 0x1a, 0x0a,
	0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x54, 0x61, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x54, 0x61, 0x67, 0x22, 0x8c, 0x04, 0x0a, 0x15, 0x53, 0x68,
	0x6f, 0x72, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x73,
	0x68, 0x6f, 0x72, 0x74, 0x5f, 0x75, 0x72, 0x69, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x73, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x72, 0x69, 0x12, 0x24, 0x0a, 0x0e, 0x66, 0x75, 0x6c, 0x6c,
	0x5f, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x66, 0x75, 0x6c, 0x6c, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x72, 0x6c, 0x12, 0x1d,
	0x0a, 0x0a, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x55, 0x72, 0x6c, 0x12, 0x10, 0x0a,
	0x03, 0x67, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x67, 0x69, 0x64, 0x12,
	0x26, 0x0a, 0x0f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x44,
	0x61, 0x74, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x65, 0x6e, 0x61, 0x62, 0x6c,
	0x65, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c,
	0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1d, 0x0a, 0x0a,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x09, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x44, 0x61, 0x74, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x66, 0x61, 0x76, 0x69,
	0x63, 0x6f, 0x6e, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x66, 0x61, 0x76, 0x69, 0x63,
	0x6f, 0x6e, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x70, 0x76, 0x18, 0x0d,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x76, 0x12, 0x19, 0x0a,
	0x08, 0x74, 0x6f, 0x64, 0x61, 0x79, 0x5f, 0x70, 0x76, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x07, 0x74, 0x6f, 0x64, 0x61, 0x79, 0x50, 0x76, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x5f, 0x75, 0x76, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x55, 0x76, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x6f, 0x64, 0x61, 0x79, 0x5f, 0x75, 0x76, 0x18,
	0x10, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x74, 0x6f, 0x64, 0x61, 0x79, 0x55, 0x76, 0x12, 0x1b,
	0x0a, 0x09, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x75, 0x69, 0x70, 0x18, 0x11, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x08, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x55, 0x69, 0x70, 0x12, 0x1b, 0x0a, 0x09, 0x74,
	0x6f, 0x64, 0x61, 0x79, 0x5f, 0x75, 0x69, 0x70, 0x18, 0x12, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08,
	0x74, 0x6f, 0x64, 0x61, 0x79, 0x55, 0x69, 0x70, 0x22, 0x32, 0x0a, 0x1e, 0x4c, 0x69, 0x73, 0x74,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x67, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x67, 0x69, 0x64, 0x22, 0x5b, 0x0a, 0x1f,
	0x4c, 0x69, 0x73, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x4c, 0x69,
	0x6e, 0x6b, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x67, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x67, 0x69,
	0x64, 0x12, 0x26, 0x0a, 0x0e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x6c, 0x69, 0x6e, 0x6b, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x73, 0x68, 0x6f, 0x72, 0x74,
	0x6c, 0x69, 0x6e, 0x6b, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x32, 0x9e, 0x05, 0x0a, 0x09, 0x53, 0x68,
	0x6f, 0x72, 0x74, 0x6c, 0x69, 0x6e, 0x6b, 0x12, 0x53, 0x0a, 0x0a, 0x52, 0x65, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x21, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x6c, 0x69, 0x6e,
	0x6b, 0x2e, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74,
	0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x58, 0x0a, 0x0f,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x12,
	0x21, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x53, 0x68, 0x6f, 0x72,
	0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x22, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x53,
	0x68, 0x6f, 0x72, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5e, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x53, 0x68, 0x6f, 0x72, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x42, 0x79, 0x4c, 0x6f, 0x63, 0x6b, 0x12,
	0x21, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x53, 0x68, 0x6f, 0x72,
	0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x22, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x53,
	0x68, 0x6f, 0x72, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x62, 0x0a, 0x14, 0x42, 0x61, 0x74, 0x63, 0x68, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x26,
	0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x53, 0x68, 0x6f, 0x72, 0x74,
	0x4c, 0x69, 0x6e, 0x6b, 0x42, 0x61, 0x74, 0x63, 0x68, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x6c, 0x69,
	0x6e, 0x6b, 0x2e, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x58, 0x0a, 0x0f, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x21, 0x2e,
	0x73, 0x68, 0x6f, 0x72, 0x74, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x4c,
	0x69, 0x6e, 0x6b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x22, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x53, 0x68, 0x6f,
	0x72, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x52, 0x0a, 0x0d, 0x50, 0x61, 0x67, 0x65, 0x53, 0x68, 0x6f, 0x72,
	0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x1f, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x6c, 0x69, 0x6e,
	0x6b, 0x2e, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x50, 0x61, 0x67, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x6c, 0x69,
	0x6e, 0x6b, 0x2e, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x50, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x70, 0x0a, 0x17, 0x4c, 0x69, 0x73, 0x74,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x29, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x6c, 0x69, 0x6e, 0x6b, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x4c, 0x69,
	0x6e, 0x6b, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a,
	0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0d, 0x5a, 0x0b, 0x2e, 0x2f,
	0x73, 0x68, 0x6f, 0x72, 0x74, 0x6c, 0x69, 0x6e, 0x6b, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_shortlink_proto_rawDescOnce sync.Once
	file_shortlink_proto_rawDescData = file_shortlink_proto_rawDesc
)

func file_shortlink_proto_rawDescGZIP() []byte {
	file_shortlink_proto_rawDescOnce.Do(func() {
		file_shortlink_proto_rawDescData = protoimpl.X.CompressGZIP(file_shortlink_proto_rawDescData)
	})
	return file_shortlink_proto_rawDescData
}

var file_shortlink_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_shortlink_proto_goTypes = []any{
	(*ShortLinkCreateRequest)(nil),          // 0: shortlink.ShortLinkCreateRequest
	(*ShortLinkCreateResponse)(nil),         // 1: shortlink.ShortLinkCreateResponse
	(*ShortLinkUpdateResponse)(nil),         // 2: shortlink.ShortLinkUpdateResponse
	(*ShortLinkUpdateRequest)(nil),          // 3: shortlink.ShortLinkUpdateRequest
	(*ShortLinkBatchCreateRequest)(nil),     // 4: shortlink.ShortLinkBatchCreateRequest
	(*ShortLinkPageRequest)(nil),            // 5: shortlink.ShortLinkPageRequest
	(*ShortLinkPageResponse)(nil),           // 6: shortlink.ShortLinkPageResponse
	(*ListGroupShortLinkCountRequest)(nil),  // 7: shortlink.ListGroupShortLinkCountRequest
	(*ListGroupShortLinkCountResponse)(nil), // 8: shortlink.ListGroupShortLinkCountResponse
}
var file_shortlink_proto_depIdxs = []int32{
	0, // 0: shortlink.Shortlink.RestoreUrl:input_type -> shortlink.ShortLinkCreateRequest
	0, // 1: shortlink.Shortlink.CreateShortLink:input_type -> shortlink.ShortLinkCreateRequest
	0, // 2: shortlink.Shortlink.CreateShortLinkByLock:input_type -> shortlink.ShortLinkCreateRequest
	4, // 3: shortlink.Shortlink.BatchCreateShortLink:input_type -> shortlink.ShortLinkBatchCreateRequest
	3, // 4: shortlink.Shortlink.UpdateShortLink:input_type -> shortlink.ShortLinkUpdateRequest
	5, // 5: shortlink.Shortlink.PageShortLink:input_type -> shortlink.ShortLinkPageRequest
	7, // 6: shortlink.Shortlink.ListGroupShortLinkCount:input_type -> shortlink.ListGroupShortLinkCountRequest
	1, // 7: shortlink.Shortlink.RestoreUrl:output_type -> shortlink.ShortLinkCreateResponse
	1, // 8: shortlink.Shortlink.CreateShortLink:output_type -> shortlink.ShortLinkCreateResponse
	1, // 9: shortlink.Shortlink.CreateShortLinkByLock:output_type -> shortlink.ShortLinkCreateResponse
	1, // 10: shortlink.Shortlink.BatchCreateShortLink:output_type -> shortlink.ShortLinkCreateResponse
	2, // 11: shortlink.Shortlink.UpdateShortLink:output_type -> shortlink.ShortLinkUpdateResponse
	6, // 12: shortlink.Shortlink.PageShortLink:output_type -> shortlink.ShortLinkPageResponse
	8, // 13: shortlink.Shortlink.ListGroupShortLinkCount:output_type -> shortlink.ListGroupShortLinkCountResponse
	7, // [7:14] is the sub-list for method output_type
	0, // [0:7] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_shortlink_proto_init() }
func file_shortlink_proto_init() {
	if File_shortlink_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_shortlink_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_shortlink_proto_goTypes,
		DependencyIndexes: file_shortlink_proto_depIdxs,
		MessageInfos:      file_shortlink_proto_msgTypes,
	}.Build()
	File_shortlink_proto = out.File
	file_shortlink_proto_rawDesc = nil
	file_shortlink_proto_goTypes = nil
	file_shortlink_proto_depIdxs = nil
}
