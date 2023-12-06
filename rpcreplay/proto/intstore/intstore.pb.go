// Code generated by protoc-gen-go. DO NOT EDIT.
// source: intstore.proto

package intstore

import (
	fmt "fmt"

	proto "github.com/golang/protobuf/proto"

	math "math"

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

type Item struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Value                int32    `protobuf:"varint,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Item) Reset()         { *m = Item{} }
func (m *Item) String() string { return proto.CompactTextString(m) }
func (*Item) ProtoMessage()    {}
func (*Item) Descriptor() ([]byte, []int) {
	return fileDescriptor_intstore_10c1c94979d47ae6, []int{0}
}
func (m *Item) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Item.Unmarshal(m, b)
}
func (m *Item) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Item.Marshal(b, m, deterministic)
}
func (dst *Item) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Item.Merge(dst, src)
}
func (m *Item) XXX_Size() int {
	return xxx_messageInfo_Item.Size(m)
}
func (m *Item) XXX_DiscardUnknown() {
	xxx_messageInfo_Item.DiscardUnknown(m)
}

var xxx_messageInfo_Item proto.InternalMessageInfo

func (m *Item) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Item) GetValue() int32 {
	if m != nil {
		return m.Value
	}
	return 0
}

type SetResponse struct {
	PrevValue            int32    `protobuf:"varint,1,opt,name=prev_value,json=prevValue,proto3" json:"prev_value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetResponse) Reset()         { *m = SetResponse{} }
func (m *SetResponse) String() string { return proto.CompactTextString(m) }
func (*SetResponse) ProtoMessage()    {}
func (*SetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_intstore_10c1c94979d47ae6, []int{1}
}
func (m *SetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetResponse.Unmarshal(m, b)
}
func (m *SetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetResponse.Marshal(b, m, deterministic)
}
func (dst *SetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetResponse.Merge(dst, src)
}
func (m *SetResponse) XXX_Size() int {
	return xxx_messageInfo_SetResponse.Size(m)
}
func (m *SetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SetResponse proto.InternalMessageInfo

func (m *SetResponse) GetPrevValue() int32 {
	if m != nil {
		return m.PrevValue
	}
	return 0
}

type GetRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_intstore_10c1c94979d47ae6, []int{2}
}
func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequest.Unmarshal(m, b)
}
func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
}
func (dst *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(dst, src)
}
func (m *GetRequest) XXX_Size() int {
	return xxx_messageInfo_GetRequest.Size(m)
}
func (m *GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequest proto.InternalMessageInfo

func (m *GetRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Summary struct {
	Count                int32    `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Summary) Reset()         { *m = Summary{} }
func (m *Summary) String() string { return proto.CompactTextString(m) }
func (*Summary) ProtoMessage()    {}
func (*Summary) Descriptor() ([]byte, []int) {
	return fileDescriptor_intstore_10c1c94979d47ae6, []int{3}
}
func (m *Summary) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Summary.Unmarshal(m, b)
}
func (m *Summary) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Summary.Marshal(b, m, deterministic)
}
func (dst *Summary) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Summary.Merge(dst, src)
}
func (m *Summary) XXX_Size() int {
	return xxx_messageInfo_Summary.Size(m)
}
func (m *Summary) XXX_DiscardUnknown() {
	xxx_messageInfo_Summary.DiscardUnknown(m)
}

var xxx_messageInfo_Summary proto.InternalMessageInfo

func (m *Summary) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

type ListItemsRequest struct {
	GreaterThan          int32    `protobuf:"varint,1,opt,name=greaterThan,proto3" json:"greaterThan,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListItemsRequest) Reset()         { *m = ListItemsRequest{} }
func (m *ListItemsRequest) String() string { return proto.CompactTextString(m) }
func (*ListItemsRequest) ProtoMessage()    {}
func (*ListItemsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_intstore_10c1c94979d47ae6, []int{4}
}
func (m *ListItemsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListItemsRequest.Unmarshal(m, b)
}
func (m *ListItemsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListItemsRequest.Marshal(b, m, deterministic)
}
func (dst *ListItemsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListItemsRequest.Merge(dst, src)
}
func (m *ListItemsRequest) XXX_Size() int {
	return xxx_messageInfo_ListItemsRequest.Size(m)
}
func (m *ListItemsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListItemsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListItemsRequest proto.InternalMessageInfo

func (m *ListItemsRequest) GetGreaterThan() int32 {
	if m != nil {
		return m.GreaterThan
	}
	return 0
}

func init() {
	proto.RegisterType((*Item)(nil), "intstore.Item")
	proto.RegisterType((*SetResponse)(nil), "intstore.SetResponse")
	proto.RegisterType((*GetRequest)(nil), "intstore.GetRequest")
	proto.RegisterType((*Summary)(nil), "intstore.Summary")
	proto.RegisterType((*ListItemsRequest)(nil), "intstore.ListItemsRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// IntStoreClient is the client API for IntStore service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type IntStoreClient interface {
	Set(ctx context.Context, in *Item, opts ...grpc.CallOption) (*SetResponse, error)
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*Item, error)
	// A server-to-client streaming RPC.
	ListItems(ctx context.Context, in *ListItemsRequest, opts ...grpc.CallOption) (IntStore_ListItemsClient, error)
	// A client-to-server streaming RPC.
	SetStream(ctx context.Context, opts ...grpc.CallOption) (IntStore_SetStreamClient, error)
	// A Bidirectional streaming RPC.
	StreamChat(ctx context.Context, opts ...grpc.CallOption) (IntStore_StreamChatClient, error)
}

type intStoreClient struct {
	cc *grpc.ClientConn
}

func NewIntStoreClient(cc *grpc.ClientConn) IntStoreClient {
	return &intStoreClient{cc}
}

func (c *intStoreClient) Set(ctx context.Context, in *Item, opts ...grpc.CallOption) (*SetResponse, error) {
	out := new(SetResponse)
	err := c.cc.Invoke(ctx, "/intstore.IntStore/Set", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *intStoreClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*Item, error) {
	out := new(Item)
	err := c.cc.Invoke(ctx, "/intstore.IntStore/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *intStoreClient) ListItems(ctx context.Context, in *ListItemsRequest, opts ...grpc.CallOption) (IntStore_ListItemsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_IntStore_serviceDesc.Streams[0], "/intstore.IntStore/ListItems", opts...)
	if err != nil {
		return nil, err
	}
	x := &intStoreListItemsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type IntStore_ListItemsClient interface {
	Recv() (*Item, error)
	grpc.ClientStream
}

type intStoreListItemsClient struct {
	grpc.ClientStream
}

func (x *intStoreListItemsClient) Recv() (*Item, error) {
	m := new(Item)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *intStoreClient) SetStream(ctx context.Context, opts ...grpc.CallOption) (IntStore_SetStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_IntStore_serviceDesc.Streams[1], "/intstore.IntStore/SetStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &intStoreSetStreamClient{stream}
	return x, nil
}

type IntStore_SetStreamClient interface {
	Send(*Item) error
	CloseAndRecv() (*Summary, error)
	grpc.ClientStream
}

type intStoreSetStreamClient struct {
	grpc.ClientStream
}

func (x *intStoreSetStreamClient) Send(m *Item) error {
	return x.ClientStream.SendMsg(m)
}

func (x *intStoreSetStreamClient) CloseAndRecv() (*Summary, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Summary)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *intStoreClient) StreamChat(ctx context.Context, opts ...grpc.CallOption) (IntStore_StreamChatClient, error) {
	stream, err := c.cc.NewStream(ctx, &_IntStore_serviceDesc.Streams[2], "/intstore.IntStore/StreamChat", opts...)
	if err != nil {
		return nil, err
	}
	x := &intStoreStreamChatClient{stream}
	return x, nil
}

type IntStore_StreamChatClient interface {
	Send(*Item) error
	Recv() (*Item, error)
	grpc.ClientStream
}

type intStoreStreamChatClient struct {
	grpc.ClientStream
}

func (x *intStoreStreamChatClient) Send(m *Item) error {
	return x.ClientStream.SendMsg(m)
}

func (x *intStoreStreamChatClient) Recv() (*Item, error) {
	m := new(Item)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// IntStoreServer is the server API for IntStore service.
type IntStoreServer interface {
	Set(context.Context, *Item) (*SetResponse, error)
	Get(context.Context, *GetRequest) (*Item, error)
	// A server-to-client streaming RPC.
	ListItems(*ListItemsRequest, IntStore_ListItemsServer) error
	// A client-to-server streaming RPC.
	SetStream(IntStore_SetStreamServer) error
	// A Bidirectional streaming RPC.
	StreamChat(IntStore_StreamChatServer) error
}

func RegisterIntStoreServer(s *grpc.Server, srv IntStoreServer) {
	s.RegisterService(&_IntStore_serviceDesc, srv)
}

func _IntStore_Set_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Item)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IntStoreServer).Set(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/intstore.IntStore/Set",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IntStoreServer).Set(ctx, req.(*Item))
	}
	return interceptor(ctx, in, info, handler)
}

func _IntStore_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IntStoreServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/intstore.IntStore/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IntStoreServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IntStore_ListItems_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListItemsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(IntStoreServer).ListItems(m, &intStoreListItemsServer{stream})
}

type IntStore_ListItemsServer interface {
	Send(*Item) error
	grpc.ServerStream
}

type intStoreListItemsServer struct {
	grpc.ServerStream
}

func (x *intStoreListItemsServer) Send(m *Item) error {
	return x.ServerStream.SendMsg(m)
}

func _IntStore_SetStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(IntStoreServer).SetStream(&intStoreSetStreamServer{stream})
}

type IntStore_SetStreamServer interface {
	SendAndClose(*Summary) error
	Recv() (*Item, error)
	grpc.ServerStream
}

type intStoreSetStreamServer struct {
	grpc.ServerStream
}

func (x *intStoreSetStreamServer) SendAndClose(m *Summary) error {
	return x.ServerStream.SendMsg(m)
}

func (x *intStoreSetStreamServer) Recv() (*Item, error) {
	m := new(Item)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _IntStore_StreamChat_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(IntStoreServer).StreamChat(&intStoreStreamChatServer{stream})
}

type IntStore_StreamChatServer interface {
	Send(*Item) error
	Recv() (*Item, error)
	grpc.ServerStream
}

type intStoreStreamChatServer struct {
	grpc.ServerStream
}

func (x *intStoreStreamChatServer) Send(m *Item) error {
	return x.ServerStream.SendMsg(m)
}

func (x *intStoreStreamChatServer) Recv() (*Item, error) {
	m := new(Item)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _IntStore_serviceDesc = grpc.ServiceDesc{
	ServiceName: "intstore.IntStore",
	HandlerType: (*IntStoreServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Set",
			Handler:    _IntStore_Set_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _IntStore_Get_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListItems",
			Handler:       _IntStore_ListItems_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SetStream",
			Handler:       _IntStore_SetStream_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "StreamChat",
			Handler:       _IntStore_StreamChat_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "intstore.proto",
}

func init() { proto.RegisterFile("intstore.proto", fileDescriptor_intstore_10c1c94979d47ae6) }

var fileDescriptor_intstore_10c1c94979d47ae6 = []byte{
	// 287 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x51, 0x4d, 0x4b, 0xc3, 0x40,
	0x10, 0xed, 0xf6, 0x43, 0x9b, 0x29, 0x14, 0x1d, 0x2a, 0x94, 0x82, 0x18, 0xf6, 0x94, 0x83, 0x86,
	0x50, 0xbd, 0x79, 0xf4, 0x50, 0x0a, 0x9e, 0xb2, 0xe2, 0x55, 0x56, 0x19, 0x6c, 0xc1, 0x6c, 0xe2,
	0xee, 0xa4, 0xe0, 0x9f, 0xf0, 0x37, 0xcb, 0x26, 0x6d, 0x13, 0x1a, 0x6f, 0xfb, 0x66, 0xde, 0xbc,
	0x79, 0x6f, 0x16, 0xa6, 0x5b, 0xc3, 0x8e, 0x73, 0x4b, 0x71, 0x61, 0x73, 0xce, 0x71, 0x7c, 0xc0,
	0x32, 0x81, 0xe1, 0x9a, 0x29, 0x43, 0x84, 0xa1, 0xd1, 0x19, 0xcd, 0x45, 0x28, 0xa2, 0x20, 0xad,
	0xde, 0x38, 0x83, 0xd1, 0x4e, 0x7f, 0x95, 0x34, 0xef, 0x87, 0x22, 0x1a, 0xa5, 0x35, 0x90, 0xb7,
	0x30, 0x51, 0xc4, 0x29, 0xb9, 0x22, 0x37, 0x8e, 0xf0, 0x1a, 0xa0, 0xb0, 0xb4, 0x7b, 0xab, 0x99,
	0xa2, 0x62, 0x06, 0xbe, 0xf2, 0x5a, 0xb1, 0x43, 0x80, 0x95, 0x67, 0x7f, 0x97, 0xe4, 0xf8, 0xbf,
	0x2d, 0xf2, 0x06, 0xce, 0x55, 0x99, 0x65, 0xda, 0xfe, 0xf8, 0x85, 0x1f, 0x79, 0x69, 0x78, 0x2f,
	0x53, 0x03, 0xf9, 0x00, 0x17, 0xcf, 0x5b, 0xc7, 0xde, 0xa6, 0x3b, 0x08, 0x85, 0x30, 0xf9, 0xb4,
	0xa4, 0x99, 0xec, 0xcb, 0x46, 0x9b, 0x3d, 0xbf, 0x5d, 0x5a, 0xfe, 0xf6, 0x61, 0xbc, 0x36, 0xac,
	0x7c, 0x4a, 0x8c, 0x61, 0xa0, 0x88, 0x71, 0x1a, 0x1f, 0xef, 0xe0, 0xd5, 0x16, 0x57, 0x0d, 0x6e,
	0x45, 0x92, 0x3d, 0xbc, 0x83, 0xc1, 0x8a, 0x18, 0x67, 0x4d, 0xbf, 0x09, 0xb1, 0x38, 0x51, 0x91,
	0x3d, 0x7c, 0x84, 0xe0, 0xe8, 0x10, 0x17, 0x4d, 0xfb, 0xd4, 0x76, 0x77, 0x34, 0x11, 0xb8, 0x84,
	0x40, 0x11, 0x2b, 0xb6, 0xa4, 0xb3, 0x8e, 0xc3, 0xcb, 0x96, 0xc3, 0xfa, 0x48, 0xb2, 0x17, 0xf9,
	0x19, 0xa8, 0x07, 0x9e, 0x36, 0xba, 0x1b, 0xab, 0xb3, 0x25, 0x12, 0x89, 0x78, 0x3f, 0xab, 0xbe,
	0xfe, 0xfe, 0x2f, 0x00, 0x00, 0xff, 0xff, 0x8f, 0x47, 0xc4, 0x86, 0x0c, 0x02, 0x00, 0x00,
}
