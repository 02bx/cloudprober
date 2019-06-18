// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/google/cloudprober/targets/rds/proto/rds.proto

package proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type IPConfig_IPType int32

const (
	// Default IP of the resource.
	//  - Private IP for instance resource
	//  - Forwarding rule IP for forwarding rule.
	IPConfig_DEFAULT IPConfig_IPType = 0
	// Instance's external IP.
	IPConfig_PUBLIC IPConfig_IPType = 1
	// First IP address from the first Alias IP range. For example, for
	// alias IP range "192.168.12.0/24", 192.168.12.0 will be returned.
	// Supported only on GCE.
	IPConfig_ALIAS IPConfig_IPType = 2
)

var IPConfig_IPType_name = map[int32]string{
	0: "DEFAULT",
	1: "PUBLIC",
	2: "ALIAS",
}
var IPConfig_IPType_value = map[string]int32{
	"DEFAULT": 0,
	"PUBLIC":  1,
	"ALIAS":   2,
}

func (x IPConfig_IPType) Enum() *IPConfig_IPType {
	p := new(IPConfig_IPType)
	*p = x
	return p
}
func (x IPConfig_IPType) String() string {
	return proto.EnumName(IPConfig_IPType_name, int32(x))
}
func (x *IPConfig_IPType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(IPConfig_IPType_value, data, "IPConfig_IPType")
	if err != nil {
		return err
	}
	*x = IPConfig_IPType(value)
	return nil
}
func (IPConfig_IPType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_rds_0538639bb1883142, []int{2, 0}
}

type ListResourcesRequest struct {
	// Provider is the resource list provider, for example: "gcp", "aws", etc.
	Provider *string `protobuf:"bytes,1,req,name=provider" json:"provider,omitempty"`
	// Provider specific resource path. For example: for GCP, it could be
	// "gce_instances/<project>", "regional_forwarding_rules/<project>", etc.
	ResourcePath *string `protobuf:"bytes,2,opt,name=resource_path,json=resourcePath" json:"resource_path,omitempty"`
	// Filters for the resources list. Filters are ANDed: all filters should
	// succeed for an item to included in the result list.
	Filter []*Filter `protobuf:"bytes,3,rep,name=filter" json:"filter,omitempty"`
	// Optional. If resource has an IP (and a NIC) address, following
	// fields determine which IP address will be included in the results.
	IpConfig             *IPConfig `protobuf:"bytes,4,opt,name=ip_config,json=ipConfig" json:"ip_config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ListResourcesRequest) Reset()         { *m = ListResourcesRequest{} }
func (m *ListResourcesRequest) String() string { return proto.CompactTextString(m) }
func (*ListResourcesRequest) ProtoMessage()    {}
func (*ListResourcesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_rds_0538639bb1883142, []int{0}
}
func (m *ListResourcesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListResourcesRequest.Unmarshal(m, b)
}
func (m *ListResourcesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListResourcesRequest.Marshal(b, m, deterministic)
}
func (dst *ListResourcesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListResourcesRequest.Merge(dst, src)
}
func (m *ListResourcesRequest) XXX_Size() int {
	return xxx_messageInfo_ListResourcesRequest.Size(m)
}
func (m *ListResourcesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListResourcesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListResourcesRequest proto.InternalMessageInfo

func (m *ListResourcesRequest) GetProvider() string {
	if m != nil && m.Provider != nil {
		return *m.Provider
	}
	return ""
}

func (m *ListResourcesRequest) GetResourcePath() string {
	if m != nil && m.ResourcePath != nil {
		return *m.ResourcePath
	}
	return ""
}

func (m *ListResourcesRequest) GetFilter() []*Filter {
	if m != nil {
		return m.Filter
	}
	return nil
}

func (m *ListResourcesRequest) GetIpConfig() *IPConfig {
	if m != nil {
		return m.IpConfig
	}
	return nil
}

type Filter struct {
	Key                  *string  `protobuf:"bytes,1,req,name=key" json:"key,omitempty"`
	Value                *string  `protobuf:"bytes,2,req,name=value" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Filter) Reset()         { *m = Filter{} }
func (m *Filter) String() string { return proto.CompactTextString(m) }
func (*Filter) ProtoMessage()    {}
func (*Filter) Descriptor() ([]byte, []int) {
	return fileDescriptor_rds_0538639bb1883142, []int{1}
}
func (m *Filter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Filter.Unmarshal(m, b)
}
func (m *Filter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Filter.Marshal(b, m, deterministic)
}
func (dst *Filter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Filter.Merge(dst, src)
}
func (m *Filter) XXX_Size() int {
	return xxx_messageInfo_Filter.Size(m)
}
func (m *Filter) XXX_DiscardUnknown() {
	xxx_messageInfo_Filter.DiscardUnknown(m)
}

var xxx_messageInfo_Filter proto.InternalMessageInfo

func (m *Filter) GetKey() string {
	if m != nil && m.Key != nil {
		return *m.Key
	}
	return ""
}

func (m *Filter) GetValue() string {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return ""
}

type IPConfig struct {
	// NIC index
	NicIndex             *int32           `protobuf:"varint,1,opt,name=nic_index,json=nicIndex,def=0" json:"nic_index,omitempty"`
	IpType               *IPConfig_IPType `protobuf:"varint,3,opt,name=ip_type,json=ipType,enum=cloudprober.targets.rds.IPConfig_IPType" json:"ip_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *IPConfig) Reset()         { *m = IPConfig{} }
func (m *IPConfig) String() string { return proto.CompactTextString(m) }
func (*IPConfig) ProtoMessage()    {}
func (*IPConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_rds_0538639bb1883142, []int{2}
}
func (m *IPConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IPConfig.Unmarshal(m, b)
}
func (m *IPConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IPConfig.Marshal(b, m, deterministic)
}
func (dst *IPConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IPConfig.Merge(dst, src)
}
func (m *IPConfig) XXX_Size() int {
	return xxx_messageInfo_IPConfig.Size(m)
}
func (m *IPConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_IPConfig.DiscardUnknown(m)
}

var xxx_messageInfo_IPConfig proto.InternalMessageInfo

const Default_IPConfig_NicIndex int32 = 0

func (m *IPConfig) GetNicIndex() int32 {
	if m != nil && m.NicIndex != nil {
		return *m.NicIndex
	}
	return Default_IPConfig_NicIndex
}

func (m *IPConfig) GetIpType() IPConfig_IPType {
	if m != nil && m.IpType != nil {
		return *m.IpType
	}
	return IPConfig_DEFAULT
}

type Resource struct {
	// Resource name.
	Name *string `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	// Resource's IP address, selected based on the request's ip_config.
	Ip *string `protobuf:"bytes,2,opt,name=ip" json:"ip,omitempty"`
	// Id associated with the resource, if any.
	Id *string `protobuf:"bytes,3,opt,name=id" json:"id,omitempty"`
	// Optional info associated with the resource. Some resource type may make use
	// of it.
	Info                 []byte   `protobuf:"bytes,4,opt,name=info" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Resource) Reset()         { *m = Resource{} }
func (m *Resource) String() string { return proto.CompactTextString(m) }
func (*Resource) ProtoMessage()    {}
func (*Resource) Descriptor() ([]byte, []int) {
	return fileDescriptor_rds_0538639bb1883142, []int{3}
}
func (m *Resource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Resource.Unmarshal(m, b)
}
func (m *Resource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Resource.Marshal(b, m, deterministic)
}
func (dst *Resource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Resource.Merge(dst, src)
}
func (m *Resource) XXX_Size() int {
	return xxx_messageInfo_Resource.Size(m)
}
func (m *Resource) XXX_DiscardUnknown() {
	xxx_messageInfo_Resource.DiscardUnknown(m)
}

var xxx_messageInfo_Resource proto.InternalMessageInfo

func (m *Resource) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Resource) GetIp() string {
	if m != nil && m.Ip != nil {
		return *m.Ip
	}
	return ""
}

func (m *Resource) GetId() string {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return ""
}

func (m *Resource) GetInfo() []byte {
	if m != nil {
		return m.Info
	}
	return nil
}

type ListResourcesResponse struct {
	Resources            []*Resource `protobuf:"bytes,1,rep,name=resources" json:"resources,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ListResourcesResponse) Reset()         { *m = ListResourcesResponse{} }
func (m *ListResourcesResponse) String() string { return proto.CompactTextString(m) }
func (*ListResourcesResponse) ProtoMessage()    {}
func (*ListResourcesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_rds_0538639bb1883142, []int{4}
}
func (m *ListResourcesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListResourcesResponse.Unmarshal(m, b)
}
func (m *ListResourcesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListResourcesResponse.Marshal(b, m, deterministic)
}
func (dst *ListResourcesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListResourcesResponse.Merge(dst, src)
}
func (m *ListResourcesResponse) XXX_Size() int {
	return xxx_messageInfo_ListResourcesResponse.Size(m)
}
func (m *ListResourcesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListResourcesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListResourcesResponse proto.InternalMessageInfo

func (m *ListResourcesResponse) GetResources() []*Resource {
	if m != nil {
		return m.Resources
	}
	return nil
}

func init() {
	proto.RegisterType((*ListResourcesRequest)(nil), "cloudprober.targets.rds.ListResourcesRequest")
	proto.RegisterType((*Filter)(nil), "cloudprober.targets.rds.Filter")
	proto.RegisterType((*IPConfig)(nil), "cloudprober.targets.rds.IPConfig")
	proto.RegisterType((*Resource)(nil), "cloudprober.targets.rds.Resource")
	proto.RegisterType((*ListResourcesResponse)(nil), "cloudprober.targets.rds.ListResourcesResponse")
	proto.RegisterEnum("cloudprober.targets.rds.IPConfig_IPType", IPConfig_IPType_name, IPConfig_IPType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ResourceDiscoveryClient is the client API for ResourceDiscovery service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ResourceDiscoveryClient interface {
	// ListResources returns the list of resources matching the URI provided in
	// the request.
	ListResources(ctx context.Context, in *ListResourcesRequest, opts ...grpc.CallOption) (*ListResourcesResponse, error)
}

type resourceDiscoveryClient struct {
	cc *grpc.ClientConn
}

func NewResourceDiscoveryClient(cc *grpc.ClientConn) ResourceDiscoveryClient {
	return &resourceDiscoveryClient{cc}
}

func (c *resourceDiscoveryClient) ListResources(ctx context.Context, in *ListResourcesRequest, opts ...grpc.CallOption) (*ListResourcesResponse, error) {
	out := new(ListResourcesResponse)
	err := c.cc.Invoke(ctx, "/cloudprober.targets.rds.ResourceDiscovery/ListResources", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ResourceDiscoveryServer is the server API for ResourceDiscovery service.
type ResourceDiscoveryServer interface {
	// ListResources returns the list of resources matching the URI provided in
	// the request.
	ListResources(context.Context, *ListResourcesRequest) (*ListResourcesResponse, error)
}

func RegisterResourceDiscoveryServer(s *grpc.Server, srv ResourceDiscoveryServer) {
	s.RegisterService(&_ResourceDiscovery_serviceDesc, srv)
}

func _ResourceDiscovery_ListResources_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListResourcesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceDiscoveryServer).ListResources(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloudprober.targets.rds.ResourceDiscovery/ListResources",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceDiscoveryServer).ListResources(ctx, req.(*ListResourcesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ResourceDiscovery_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cloudprober.targets.rds.ResourceDiscovery",
	HandlerType: (*ResourceDiscoveryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListResources",
			Handler:    _ResourceDiscovery_ListResources_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/google/cloudprober/targets/rds/proto/rds.proto",
}

func init() {
	proto.RegisterFile("github.com/google/cloudprober/targets/rds/proto/rds.proto", fileDescriptor_rds_0538639bb1883142)
}

var fileDescriptor_rds_0538639bb1883142 = []byte{
	// 453 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0xcd, 0x6e, 0xd3, 0x4c,
	0x14, 0xed, 0x38, 0x4d, 0x6a, 0xdf, 0xb4, 0x55, 0xbe, 0xab, 0x7e, 0xc2, 0xea, 0x02, 0x8c, 0xd9,
	0x78, 0x01, 0x76, 0x95, 0x0d, 0x82, 0x05, 0x28, 0xb4, 0x54, 0xb2, 0x94, 0x45, 0x34, 0xb4, 0x12,
	0xbb, 0xc8, 0xb5, 0x27, 0xce, 0x88, 0xd4, 0x33, 0xcc, 0x8c, 0x23, 0xf2, 0x00, 0x3c, 0x05, 0x2f,
	0xc5, 0x23, 0x21, 0x8f, 0x6d, 0xfe, 0x44, 0x04, 0x2b, 0x9f, 0x7b, 0x74, 0xce, 0xf5, 0x99, 0x39,
	0x03, 0x2f, 0x4a, 0x6e, 0xd6, 0xf5, 0x5d, 0x9c, 0x8b, 0xfb, 0xa4, 0x14, 0xa2, 0xdc, 0xb0, 0x24,
	0xdf, 0x88, 0xba, 0x90, 0x4a, 0xdc, 0x31, 0x95, 0x98, 0x4c, 0x95, 0xcc, 0xe8, 0x44, 0x15, 0x3a,
	0x91, 0x4a, 0x18, 0xd1, 0xa0, 0xd8, 0x22, 0x7c, 0xf0, 0x93, 0x30, 0xee, 0x84, 0xb1, 0x2a, 0x74,
	0xf8, 0x95, 0xc0, 0xd9, 0x9c, 0x6b, 0x43, 0x99, 0x16, 0xb5, 0xca, 0x99, 0xa6, 0xec, 0x63, 0xcd,
	0xb4, 0xc1, 0x73, 0x70, 0xa5, 0x12, 0x5b, 0x5e, 0x30, 0xe5, 0x93, 0xc0, 0x89, 0x3c, 0xfa, 0x7d,
	0xc6, 0x27, 0x70, 0xa2, 0x3a, 0xfd, 0x52, 0x66, 0x66, 0xed, 0x3b, 0x01, 0x89, 0x3c, 0x7a, 0xdc,
	0x93, 0x8b, 0xcc, 0xac, 0xf1, 0x39, 0x8c, 0x56, 0x7c, 0x63, 0x98, 0xf2, 0x07, 0xc1, 0x20, 0x1a,
	0x4f, 0x1f, 0xc5, 0x7b, 0x32, 0xc4, 0xd7, 0x56, 0x46, 0x3b, 0x39, 0xbe, 0x02, 0x8f, 0xcb, 0x65,
	0x2e, 0xaa, 0x15, 0x2f, 0xfd, 0xc3, 0x80, 0x44, 0xe3, 0xe9, 0xe3, 0xbd, 0xde, 0x74, 0x71, 0x69,
	0x85, 0xd4, 0xe5, 0xb2, 0x45, 0xe1, 0x05, 0x8c, 0xda, 0x8d, 0x38, 0x81, 0xc1, 0x07, 0xb6, 0xeb,
	0xe2, 0x37, 0x10, 0xcf, 0x60, 0xb8, 0xcd, 0x36, 0x35, 0xf3, 0x1d, 0xcb, 0xb5, 0x43, 0xf8, 0x85,
	0x80, 0xdb, 0x2f, 0xc2, 0x87, 0xe0, 0x55, 0x3c, 0x5f, 0xf2, 0xaa, 0x60, 0x9f, 0x7c, 0x12, 0x90,
	0x68, 0xf8, 0x92, 0x5c, 0x50, 0xb7, 0xe2, 0x79, 0xda, 0x50, 0x38, 0x83, 0x23, 0x2e, 0x97, 0x66,
	0x27, 0x99, 0x3f, 0x08, 0x48, 0x74, 0x3a, 0x8d, 0xfe, 0x1a, 0x2e, 0x4e, 0x17, 0x37, 0x3b, 0xc9,
	0xe8, 0x88, 0xcb, 0xe6, 0x1b, 0x3e, 0x85, 0x51, 0xcb, 0xe0, 0x18, 0x8e, 0xae, 0xde, 0x5e, 0xcf,
	0x6e, 0xe7, 0x37, 0x93, 0x03, 0x04, 0x18, 0x2d, 0x6e, 0xdf, 0xcc, 0xd3, 0xcb, 0x09, 0x41, 0x0f,
	0x86, 0xb3, 0x79, 0x3a, 0x7b, 0x37, 0x71, 0x42, 0x0a, 0x6e, 0xdf, 0x0e, 0x22, 0x1c, 0x56, 0xd9,
	0x3d, 0xeb, 0x8e, 0x64, 0x31, 0x9e, 0x82, 0xc3, 0x65, 0x57, 0x81, 0xc3, 0xa5, 0x9d, 0x0b, 0x9b,
	0xad, 0x99, 0x8b, 0xc6, 0xc3, 0xab, 0x95, 0xb0, 0x57, 0x79, 0x4c, 0x2d, 0x0e, 0xdf, 0xc3, 0xff,
	0xbf, 0xb5, 0xae, 0xa5, 0xa8, 0x34, 0xc3, 0xd7, 0xe0, 0xf5, 0x2d, 0x6a, 0x9f, 0xd8, 0xe2, 0xf6,
	0x5f, 0x7e, 0x6f, 0xa7, 0x3f, 0x3c, 0xd3, 0xcf, 0x04, 0xfe, 0xeb, 0xf9, 0x2b, 0xae, 0x73, 0xb1,
	0x65, 0x6a, 0x87, 0x12, 0x4e, 0x7e, 0xf9, 0x1f, 0x3e, 0xdb, 0xbb, 0xf4, 0x4f, 0xaf, 0xf1, 0x3c,
	0xfe, 0x57, 0x79, 0x7b, 0x8c, 0xf0, 0xe0, 0x5b, 0x00, 0x00, 0x00, 0xff, 0xff, 0x95, 0xbd, 0xf0,
	0xbb, 0x2d, 0x03, 0x00, 0x00,
}