// Code generated by protoc-gen-go. DO NOT EDIT.
// source: workload_attestor.proto

/*
Package sri_proto is a generated protocol buffer package.

It is generated from these files:
	workload_attestor.proto

It has these top-level messages:
	AttestRequest
	AttestResponse
*/
package sri_proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import sri_proto1 "github.com/spiffe/sri/common/plugins/common/proto"

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

// ConfigureRequest from public import github.com/spiffe/sri/common/plugins/common/proto/common.proto
type ConfigureRequest sri_proto1.ConfigureRequest

func (m *ConfigureRequest) Reset()         { (*sri_proto1.ConfigureRequest)(m).Reset() }
func (m *ConfigureRequest) String() string { return (*sri_proto1.ConfigureRequest)(m).String() }
func (*ConfigureRequest) ProtoMessage()    {}
func (m *ConfigureRequest) GetConfiguration() string {
	return (*sri_proto1.ConfigureRequest)(m).GetConfiguration()
}

// ConfigureResponse from public import github.com/spiffe/sri/common/plugins/common/proto/common.proto
type ConfigureResponse sri_proto1.ConfigureResponse

func (m *ConfigureResponse) Reset()         { (*sri_proto1.ConfigureResponse)(m).Reset() }
func (m *ConfigureResponse) String() string { return (*sri_proto1.ConfigureResponse)(m).String() }
func (*ConfigureResponse) ProtoMessage()    {}
func (m *ConfigureResponse) GetErrorList() []string {
	return (*sri_proto1.ConfigureResponse)(m).GetErrorList()
}

// GetPluginInfoRequest from public import github.com/spiffe/sri/common/plugins/common/proto/common.proto
type GetPluginInfoRequest sri_proto1.GetPluginInfoRequest

func (m *GetPluginInfoRequest) Reset()         { (*sri_proto1.GetPluginInfoRequest)(m).Reset() }
func (m *GetPluginInfoRequest) String() string { return (*sri_proto1.GetPluginInfoRequest)(m).String() }
func (*GetPluginInfoRequest) ProtoMessage()    {}

// GetPluginInfoResponse from public import github.com/spiffe/sri/common/plugins/common/proto/common.proto
type GetPluginInfoResponse sri_proto1.GetPluginInfoResponse

func (m *GetPluginInfoResponse) Reset() { (*sri_proto1.GetPluginInfoResponse)(m).Reset() }
func (m *GetPluginInfoResponse) String() string {
	return (*sri_proto1.GetPluginInfoResponse)(m).String()
}
func (*GetPluginInfoResponse) ProtoMessage() {}
func (m *GetPluginInfoResponse) GetName() string {
	return (*sri_proto1.GetPluginInfoResponse)(m).GetName()
}
func (m *GetPluginInfoResponse) GetCategory() string {
	return (*sri_proto1.GetPluginInfoResponse)(m).GetCategory()
}
func (m *GetPluginInfoResponse) GetType() string {
	return (*sri_proto1.GetPluginInfoResponse)(m).GetType()
}
func (m *GetPluginInfoResponse) GetDescription() string {
	return (*sri_proto1.GetPluginInfoResponse)(m).GetDescription()
}
func (m *GetPluginInfoResponse) GetDateCreated() string {
	return (*sri_proto1.GetPluginInfoResponse)(m).GetDateCreated()
}
func (m *GetPluginInfoResponse) GetLocation() string {
	return (*sri_proto1.GetPluginInfoResponse)(m).GetLocation()
}
func (m *GetPluginInfoResponse) GetVersion() string {
	return (*sri_proto1.GetPluginInfoResponse)(m).GetVersion()
}
func (m *GetPluginInfoResponse) GetAuthor() string {
	return (*sri_proto1.GetPluginInfoResponse)(m).GetAuthor()
}
func (m *GetPluginInfoResponse) GetCompany() string {
	return (*sri_proto1.GetPluginInfoResponse)(m).GetCompany()
}

// * Represents the workload PID.
type AttestRequest struct {
	Pid int32 `protobuf:"varint,1,opt,name=pid" json:"pid,omitempty"`
}

func (m *AttestRequest) Reset()                    { *m = AttestRequest{} }
func (m *AttestRequest) String() string            { return proto.CompactTextString(m) }
func (*AttestRequest) ProtoMessage()               {}
func (*AttestRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *AttestRequest) GetPid() int32 {
	if m != nil {
		return m.Pid
	}
	return 0
}

// * Represents a list of selectors resolved for a given PID.
type AttestResponse struct {
	Selectors []string `protobuf:"bytes,1,rep,name=selectors" json:"selectors,omitempty"`
}

func (m *AttestResponse) Reset()                    { *m = AttestResponse{} }
func (m *AttestResponse) String() string            { return proto.CompactTextString(m) }
func (*AttestResponse) ProtoMessage()               {}
func (*AttestResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *AttestResponse) GetSelectors() []string {
	if m != nil {
		return m.Selectors
	}
	return nil
}

func init() {
	proto.RegisterType((*AttestRequest)(nil), "sri_proto.AttestRequest")
	proto.RegisterType((*AttestResponse)(nil), "sri_proto.AttestResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for WorkloadAttestor service

type WorkloadAttestorClient interface {
	// / Returns a list of selectors resolved for a given PID
	Attest(ctx context.Context, in *AttestRequest, opts ...grpc.CallOption) (*AttestResponse, error)
	// / Applies the plugin configuration and returns configuration errors
	Configure(ctx context.Context, in *sri_proto1.ConfigureRequest, opts ...grpc.CallOption) (*sri_proto1.ConfigureResponse, error)
	// / Returns the version and related metadata of the plugin
	GetPluginInfo(ctx context.Context, in *sri_proto1.GetPluginInfoRequest, opts ...grpc.CallOption) (*sri_proto1.GetPluginInfoResponse, error)
}

type workloadAttestorClient struct {
	cc *grpc.ClientConn
}

func NewWorkloadAttestorClient(cc *grpc.ClientConn) WorkloadAttestorClient {
	return &workloadAttestorClient{cc}
}

func (c *workloadAttestorClient) Attest(ctx context.Context, in *AttestRequest, opts ...grpc.CallOption) (*AttestResponse, error) {
	out := new(AttestResponse)
	err := grpc.Invoke(ctx, "/sri_proto.WorkloadAttestor/Attest", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workloadAttestorClient) Configure(ctx context.Context, in *sri_proto1.ConfigureRequest, opts ...grpc.CallOption) (*sri_proto1.ConfigureResponse, error) {
	out := new(sri_proto1.ConfigureResponse)
	err := grpc.Invoke(ctx, "/sri_proto.WorkloadAttestor/Configure", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workloadAttestorClient) GetPluginInfo(ctx context.Context, in *sri_proto1.GetPluginInfoRequest, opts ...grpc.CallOption) (*sri_proto1.GetPluginInfoResponse, error) {
	out := new(sri_proto1.GetPluginInfoResponse)
	err := grpc.Invoke(ctx, "/sri_proto.WorkloadAttestor/GetPluginInfo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for WorkloadAttestor service

type WorkloadAttestorServer interface {
	// / Returns a list of selectors resolved for a given PID
	Attest(context.Context, *AttestRequest) (*AttestResponse, error)
	// / Applies the plugin configuration and returns configuration errors
	Configure(context.Context, *sri_proto1.ConfigureRequest) (*sri_proto1.ConfigureResponse, error)
	// / Returns the version and related metadata of the plugin
	GetPluginInfo(context.Context, *sri_proto1.GetPluginInfoRequest) (*sri_proto1.GetPluginInfoResponse, error)
}

func RegisterWorkloadAttestorServer(s *grpc.Server, srv WorkloadAttestorServer) {
	s.RegisterService(&_WorkloadAttestor_serviceDesc, srv)
}

func _WorkloadAttestor_Attest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AttestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkloadAttestorServer).Attest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sri_proto.WorkloadAttestor/Attest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkloadAttestorServer).Attest(ctx, req.(*AttestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkloadAttestor_Configure_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(sri_proto1.ConfigureRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkloadAttestorServer).Configure(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sri_proto.WorkloadAttestor/Configure",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkloadAttestorServer).Configure(ctx, req.(*sri_proto1.ConfigureRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkloadAttestor_GetPluginInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(sri_proto1.GetPluginInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkloadAttestorServer).GetPluginInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sri_proto.WorkloadAttestor/GetPluginInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkloadAttestorServer).GetPluginInfo(ctx, req.(*sri_proto1.GetPluginInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _WorkloadAttestor_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sri_proto.WorkloadAttestor",
	HandlerType: (*WorkloadAttestorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Attest",
			Handler:    _WorkloadAttestor_Attest_Handler,
		},
		{
			MethodName: "Configure",
			Handler:    _WorkloadAttestor_Configure_Handler,
		},
		{
			MethodName: "GetPluginInfo",
			Handler:    _WorkloadAttestor_GetPluginInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "workload_attestor.proto",
}

func init() { proto.RegisterFile("workload_attestor.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 253 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0x41, 0x4b, 0xc3, 0x40,
	0x10, 0x85, 0x0d, 0xc5, 0x42, 0x06, 0x2a, 0x65, 0x2f, 0xc6, 0x5a, 0x30, 0xf6, 0xd4, 0xd3, 0x06,
	0xf4, 0xac, 0x50, 0x04, 0xc5, 0x5b, 0xc9, 0xc5, 0x63, 0x69, 0xd3, 0x49, 0x5c, 0x4c, 0x76, 0xd6,
	0x9d, 0x0d, 0xfe, 0x6a, 0xff, 0x83, 0x34, 0xd9, 0xc4, 0x16, 0xd2, 0xdb, 0x9b, 0x79, 0x6f, 0xbf,
	0x79, 0x2c, 0x5c, 0xff, 0x90, 0xfd, 0x2a, 0x69, 0xbb, 0xdf, 0x6c, 0x9d, 0x43, 0x76, 0x64, 0xa5,
	0xb1, 0xe4, 0x48, 0x84, 0x6c, 0xd5, 0xa6, 0x91, 0xb3, 0xe7, 0x42, 0xb9, 0xcf, 0x7a, 0x27, 0x33,
	0xaa, 0x12, 0x36, 0x2a, 0xcf, 0x31, 0x61, 0xab, 0x92, 0x8c, 0xaa, 0x8a, 0x74, 0x62, 0xca, 0xba,
	0x50, 0x9a, 0xfb, 0xf1, 0xf0, 0xc4, 0x0f, 0x2d, 0x6a, 0x71, 0x0f, 0x93, 0x55, 0x03, 0x4f, 0xf1,
	0xbb, 0x46, 0x76, 0x62, 0x0a, 0x23, 0xa3, 0xf6, 0x51, 0x10, 0x07, 0xcb, 0xcb, 0xf4, 0x20, 0x17,
	0x12, 0xae, 0xba, 0x08, 0x1b, 0xd2, 0x8c, 0x62, 0x0e, 0x21, 0x63, 0x89, 0x99, 0x23, 0xcb, 0x51,
	0x10, 0x8f, 0x96, 0x61, 0xfa, 0xbf, 0x78, 0xf8, 0x0d, 0x60, 0xfa, 0xe1, 0x9b, 0xaf, 0x7c, 0x71,
	0xf1, 0x04, 0xe3, 0x56, 0x8b, 0x48, 0xf6, 0xed, 0xe5, 0xc9, 0xe9, 0xd9, 0xcd, 0x80, 0xe3, 0x2f,
	0xbe, 0x42, 0xf8, 0x42, 0x3a, 0x57, 0x45, 0x6d, 0x51, 0xdc, 0x1e, 0xe5, 0xfa, 0x6d, 0x07, 0x99,
	0x0f, 0x9b, 0x9e, 0x93, 0xc2, 0xe4, 0x0d, 0xdd, 0xba, 0xf9, 0x96, 0x77, 0x9d, 0x93, 0xb8, 0x3b,
	0x8a, 0x9f, 0x38, 0x1d, 0x2f, 0x3e, 0x1f, 0x68, 0x99, 0xeb, 0x8b, 0xdd, 0xb8, 0xb1, 0x1f, 0xff,
	0x02, 0x00, 0x00, 0xff, 0xff, 0x4e, 0xc9, 0xbd, 0x6b, 0xb1, 0x01, 0x00, 0x00,
}
