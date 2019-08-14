// Code generated by protoc-gen-go. DO NOT EDIT.
// source: dockerator.proto

package dockerator

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Request struct {
	Node                 string   `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	Service              string   `protobuf:"bytes,2,opt,name=service,proto3" json:"service,omitempty"`
	State                string   `protobuf:"bytes,3,opt,name=state,proto3" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_51773407af17b204, []int{0}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetNode() string {
	if m != nil {
		return m.Node
	}
	return ""
}

func (m *Request) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

func (m *Request) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

type Response struct {
	Command              string   `protobuf:"bytes,1,opt,name=command,proto3" json:"command,omitempty"`
	Params               string   `protobuf:"bytes,2,opt,name=params,proto3" json:"params,omitempty"`
	Status               bool     `protobuf:"varint,3,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_51773407af17b204, []int{1}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetCommand() string {
	if m != nil {
		return m.Command
	}
	return ""
}

func (m *Response) GetParams() string {
	if m != nil {
		return m.Params
	}
	return ""
}

func (m *Response) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

type TaskRequest struct {
	Node                 string   `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TaskRequest) Reset()         { *m = TaskRequest{} }
func (m *TaskRequest) String() string { return proto.CompactTextString(m) }
func (*TaskRequest) ProtoMessage()    {}
func (*TaskRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_51773407af17b204, []int{2}
}

func (m *TaskRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskRequest.Unmarshal(m, b)
}
func (m *TaskRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskRequest.Marshal(b, m, deterministic)
}
func (m *TaskRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskRequest.Merge(m, src)
}
func (m *TaskRequest) XXX_Size() int {
	return xxx_messageInfo_TaskRequest.Size(m)
}
func (m *TaskRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TaskRequest proto.InternalMessageInfo

func (m *TaskRequest) GetNode() string {
	if m != nil {
		return m.Node
	}
	return ""
}

type TaskResponse struct {
	Job                  string   `protobuf:"bytes,1,opt,name=job,proto3" json:"job,omitempty"`
	Params               string   `protobuf:"bytes,2,opt,name=params,proto3" json:"params,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TaskResponse) Reset()         { *m = TaskResponse{} }
func (m *TaskResponse) String() string { return proto.CompactTextString(m) }
func (*TaskResponse) ProtoMessage()    {}
func (*TaskResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_51773407af17b204, []int{3}
}

func (m *TaskResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskResponse.Unmarshal(m, b)
}
func (m *TaskResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskResponse.Marshal(b, m, deterministic)
}
func (m *TaskResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskResponse.Merge(m, src)
}
func (m *TaskResponse) XXX_Size() int {
	return xxx_messageInfo_TaskResponse.Size(m)
}
func (m *TaskResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TaskResponse proto.InternalMessageInfo

func (m *TaskResponse) GetJob() string {
	if m != nil {
		return m.Job
	}
	return ""
}

func (m *TaskResponse) GetParams() string {
	if m != nil {
		return m.Params
	}
	return ""
}

func init() {
	proto.RegisterType((*Request)(nil), "dockerator.Request")
	proto.RegisterType((*Response)(nil), "dockerator.Response")
	proto.RegisterType((*TaskRequest)(nil), "dockerator.TaskRequest")
	proto.RegisterType((*TaskResponse)(nil), "dockerator.TaskResponse")
}

func init() { proto.RegisterFile("dockerator.proto", fileDescriptor_51773407af17b204) }

var fileDescriptor_51773407af17b204 = []byte{
	// 247 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x51, 0xb1, 0x4e, 0xc3, 0x40,
	0x0c, 0x25, 0x14, 0xda, 0xe2, 0x76, 0xa8, 0x4c, 0x05, 0xa7, 0x4e, 0x70, 0x13, 0x53, 0x07, 0x58,
	0x10, 0x6b, 0x11, 0x1b, 0x4b, 0x54, 0x89, 0xf9, 0x9a, 0x58, 0x02, 0xa2, 0xc4, 0xc1, 0xbe, 0xf2,
	0x17, 0xfc, 0x33, 0xba, 0xcb, 0x05, 0x82, 0x10, 0x6c, 0x7e, 0x2f, 0xcf, 0xef, 0xbd, 0xf8, 0x60,
	0x51, 0x72, 0x51, 0x91, 0x38, 0xcf, 0xb2, 0x6e, 0x85, 0x3d, 0x23, 0x7c, 0x33, 0xf6, 0x11, 0x26,
	0x39, 0xbd, 0xed, 0x49, 0x3d, 0x22, 0x1c, 0x35, 0x5c, 0x92, 0xc9, 0x2e, 0xb2, 0xab, 0x93, 0x3c,
	0xce, 0x68, 0x60, 0xa2, 0x24, 0xef, 0x2f, 0x05, 0x99, 0xc3, 0x48, 0xf7, 0x10, 0x97, 0x70, 0xac,
	0xde, 0x79, 0x32, 0xa3, 0xc8, 0x77, 0xc0, 0x6e, 0x61, 0x9a, 0x93, 0xb6, 0xdc, 0x68, 0xdc, 0x2d,
	0xb8, 0xae, 0x5d, 0x53, 0x26, 0xcb, 0x1e, 0xe2, 0x19, 0x8c, 0x5b, 0x27, 0xae, 0xd6, 0x64, 0x9a,
	0x50, 0xe0, 0x83, 0xcd, 0x5e, 0xa3, 0xe9, 0x34, 0x4f, 0xc8, 0x5e, 0xc2, 0x6c, 0xeb, 0xb4, 0xfa,
	0xa7, 0xa8, 0xbd, 0x85, 0x79, 0x27, 0x49, 0xe1, 0x0b, 0x18, 0xbd, 0xf2, 0x2e, 0x49, 0xc2, 0xf8,
	0x57, 0xe8, 0xf5, 0x47, 0x06, 0x70, 0xff, 0x75, 0x10, 0xbc, 0x83, 0xd9, 0xe6, 0x99, 0x8a, 0xea,
	0x89, 0xa5, 0x22, 0xc1, 0xd3, 0xf5, 0xe0, 0x7c, 0xa9, 0xc0, 0x6a, 0xf9, 0x93, 0xec, 0x22, 0xed,
	0x01, 0x6e, 0x60, 0x1e, 0x77, 0x1f, 0x58, 0x42, 0x19, 0x3c, 0x1f, 0xea, 0x06, 0x7f, 0xb0, 0x32,
	0xbf, 0x3f, 0xf4, 0x26, 0xbb, 0x71, 0x7c, 0xa4, 0x9b, 0xcf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa0,
	0x5d, 0x4b, 0xda, 0xb8, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DockeratorClient is the client API for Dockerator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DockeratorClient interface {
	CheckWorker(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	CheckForTask(ctx context.Context, in *TaskRequest, opts ...grpc.CallOption) (*TaskResponse, error)
}

type dockeratorClient struct {
	cc *grpc.ClientConn
}

func NewDockeratorClient(cc *grpc.ClientConn) DockeratorClient {
	return &dockeratorClient{cc}
}

func (c *dockeratorClient) CheckWorker(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/dockerator.Dockerator/CheckWorker", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dockeratorClient) CheckForTask(ctx context.Context, in *TaskRequest, opts ...grpc.CallOption) (*TaskResponse, error) {
	out := new(TaskResponse)
	err := c.cc.Invoke(ctx, "/dockerator.Dockerator/CheckForTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DockeratorServer is the server API for Dockerator service.
type DockeratorServer interface {
	CheckWorker(context.Context, *Request) (*Response, error)
	CheckForTask(context.Context, *TaskRequest) (*TaskResponse, error)
}

// UnimplementedDockeratorServer can be embedded to have forward compatible implementations.
type UnimplementedDockeratorServer struct {
}

func (*UnimplementedDockeratorServer) CheckWorker(ctx context.Context, req *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckWorker not implemented")
}
func (*UnimplementedDockeratorServer) CheckForTask(ctx context.Context, req *TaskRequest) (*TaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckForTask not implemented")
}

func RegisterDockeratorServer(s *grpc.Server, srv DockeratorServer) {
	s.RegisterService(&_Dockerator_serviceDesc, srv)
}

func _Dockerator_CheckWorker_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DockeratorServer).CheckWorker(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dockerator.Dockerator/CheckWorker",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DockeratorServer).CheckWorker(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dockerator_CheckForTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DockeratorServer).CheckForTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dockerator.Dockerator/CheckForTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DockeratorServer).CheckForTask(ctx, req.(*TaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Dockerator_serviceDesc = grpc.ServiceDesc{
	ServiceName: "dockerator.Dockerator",
	HandlerType: (*DockeratorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CheckWorker",
			Handler:    _Dockerator_CheckWorker_Handler,
		},
		{
			MethodName: "CheckForTask",
			Handler:    _Dockerator_CheckForTask_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dockerator.proto",
}
