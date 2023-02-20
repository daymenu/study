// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/container/container.proto

package daymenu_shipping_srv_container

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "golang.org/x/net/context"
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

// 集装箱
type Container struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	CustomerId           string   `protobuf:"bytes,2,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	Origin               string   `protobuf:"bytes,3,opt,name=origin,proto3" json:"origin,omitempty"`
	UserId               string   `protobuf:"bytes,4,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Height               int64    `protobuf:"varint,5,opt,name=height,proto3" json:"height,omitempty"`
	Width                int64    `protobuf:"varint,6,opt,name=width,proto3" json:"width,omitempty"`
	Long                 int64    `protobuf:"varint,7,opt,name=long,proto3" json:"long,omitempty"`
	Status               int32    `protobuf:"varint,8,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Container) Reset()         { *m = Container{} }
func (m *Container) String() string { return proto.CompactTextString(m) }
func (*Container) ProtoMessage()    {}
func (*Container) Descriptor() ([]byte, []int) {
	return fileDescriptor_3ec6c08f9f0309d9, []int{0}
}

func (m *Container) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Container.Unmarshal(m, b)
}
func (m *Container) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Container.Marshal(b, m, deterministic)
}
func (m *Container) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Container.Merge(m, src)
}
func (m *Container) XXX_Size() int {
	return xxx_messageInfo_Container.Size(m)
}
func (m *Container) XXX_DiscardUnknown() {
	xxx_messageInfo_Container.DiscardUnknown(m)
}

var xxx_messageInfo_Container proto.InternalMessageInfo

func (m *Container) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Container) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *Container) GetOrigin() string {
	if m != nil {
		return m.Origin
	}
	return ""
}

func (m *Container) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *Container) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *Container) GetWidth() int64 {
	if m != nil {
		return m.Width
	}
	return 0
}

func (m *Container) GetLong() int64 {
	if m != nil {
		return m.Long
	}
	return 0
}

func (m *Container) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

// 请求
type Request struct {
	Height               int64        `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
	Width                int64        `protobuf:"varint,2,opt,name=width,proto3" json:"width,omitempty"`
	Long                 int64        `protobuf:"varint,3,opt,name=long,proto3" json:"long,omitempty"`
	Page                 int64        `protobuf:"varint,4,opt,name=page,proto3" json:"page,omitempty"`
	PageSize             int64        `protobuf:"varint,5,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
	Id                   int64        `protobuf:"varint,6,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string       `protobuf:"bytes,7,opt,name=name,proto3" json:"name,omitempty"`
	Containers           []*Container `protobuf:"bytes,8,rep,name=containers,proto3" json:"containers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_3ec6c08f9f0309d9, []int{1}
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

func (m *Request) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *Request) GetWidth() int64 {
	if m != nil {
		return m.Width
	}
	return 0
}

func (m *Request) GetLong() int64 {
	if m != nil {
		return m.Long
	}
	return 0
}

func (m *Request) GetPage() int64 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *Request) GetPageSize() int64 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *Request) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Request) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Request) GetContainers() []*Container {
	if m != nil {
		return m.Containers
	}
	return nil
}

//响应
type Response struct {
	Code                 int32        `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Containers           []*Container `protobuf:"bytes,2,rep,name=containers,proto3" json:"containers,omitempty"`
	Container            *Container   `protobuf:"bytes,3,opt,name=container,proto3" json:"container,omitempty"`
	Count                int64        `protobuf:"varint,9,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_3ec6c08f9f0309d9, []int{2}
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

func (m *Response) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Response) GetContainers() []*Container {
	if m != nil {
		return m.Containers
	}
	return nil
}

func (m *Response) GetContainer() *Container {
	if m != nil {
		return m.Container
	}
	return nil
}

func (m *Response) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

func init() {
	proto.RegisterType((*Container)(nil), "daymenu.shipping.srv.container.Container")
	proto.RegisterType((*Request)(nil), "daymenu.shipping.srv.container.Request")
	proto.RegisterType((*Response)(nil), "daymenu.shipping.srv.container.Response")
}

func init() { proto.RegisterFile("proto/container/container.proto", fileDescriptor_3ec6c08f9f0309d9) }

var fileDescriptor_3ec6c08f9f0309d9 = []byte{
	// 422 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x94, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xc7, 0xf1, 0x47, 0x9c, 0x78, 0x2a, 0x21, 0x34, 0x42, 0x60, 0xf5, 0x40, 0xa3, 0x5c, 0x30,
	0x17, 0x23, 0x95, 0x37, 0xa0, 0x87, 0xa8, 0x37, 0xb4, 0x55, 0x2f, 0x70, 0x88, 0x8c, 0x77, 0xe4,
	0xac, 0x20, 0xbb, 0xc6, 0xbb, 0x0e, 0x82, 0x17, 0x04, 0xf1, 0x28, 0x3c, 0x05, 0xda, 0xb1, 0xe3,
	0x46, 0xa8, 0xa2, 0x54, 0x95, 0x4f, 0x99, 0xff, 0xce, 0xc7, 0x6f, 0xf2, 0xdf, 0x95, 0xe1, 0xac,
	0x69, 0x8d, 0x33, 0xaf, 0x2b, 0xa3, 0x5d, 0xa9, 0x34, 0xb5, 0x37, 0x51, 0xc1, 0x19, 0x7c, 0x21,
	0xcb, 0x6f, 0x3b, 0xd2, 0x5d, 0x61, 0xb7, 0xaa, 0x69, 0x94, 0xae, 0x0b, 0xdb, 0xee, 0x8b, 0xb1,
	0x6a, 0xf5, 0x23, 0x80, 0xf4, 0xe2, 0xa0, 0xf0, 0x31, 0x84, 0x4a, 0x66, 0xc1, 0x32, 0xc8, 0x23,
	0x11, 0x2a, 0x89, 0x67, 0x70, 0x52, 0x75, 0xd6, 0x99, 0x1d, 0xb5, 0x1b, 0x25, 0xb3, 0x70, 0x19,
	0xe4, 0xa9, 0x80, 0xc3, 0xd1, 0xa5, 0xc4, 0x67, 0x90, 0x98, 0x56, 0xd5, 0x4a, 0x67, 0x11, 0xe7,
	0x06, 0x85, 0xcf, 0x61, 0xde, 0xd9, 0xbe, 0x29, 0xee, 0x13, 0x5e, 0xf6, 0x0d, 0x5b, 0x52, 0xf5,
	0xd6, 0x65, 0x33, 0xa6, 0x0c, 0x0a, 0x9f, 0xc2, 0xec, 0xab, 0x92, 0x6e, 0x9b, 0x25, 0x7c, 0xdc,
	0x0b, 0x44, 0x88, 0x3f, 0x1b, 0x5d, 0x67, 0x73, 0x3e, 0xe4, 0xd8, 0x4f, 0xb0, 0xae, 0x74, 0x9d,
	0xcd, 0x16, 0xcb, 0x20, 0x9f, 0x89, 0x41, 0xad, 0x7e, 0x07, 0x30, 0x17, 0xf4, 0xa5, 0x23, 0xeb,
	0x8e, 0x28, 0xc1, 0xed, 0x94, 0xf0, 0x36, 0x4a, 0x74, 0x44, 0x41, 0x88, 0x9b, 0xb2, 0x26, 0xde,
	0x3e, 0x12, 0x1c, 0xe3, 0x29, 0x2c, 0xfc, 0xef, 0x95, 0xfa, 0x4e, 0xc3, 0xf6, 0xa3, 0x1e, 0x9c,
	0x4b, 0x46, 0xe7, 0x10, 0x62, 0x5d, 0xee, 0x88, 0x37, 0x4f, 0x05, 0xc7, 0x78, 0x09, 0x30, 0x1a,
	0xef, 0xb7, 0x8f, 0xf2, 0x93, 0xf3, 0x57, 0xc5, 0xbf, 0x2f, 0xa8, 0x18, 0x2f, 0x47, 0x1c, 0x35,
	0xaf, 0x7e, 0x06, 0xb0, 0x10, 0x64, 0x1b, 0xa3, 0x2d, 0x79, 0x56, 0x65, 0x24, 0xf1, 0x7f, 0x9d,
	0x09, 0x8e, 0xff, 0x62, 0x85, 0x0f, 0x60, 0xe1, 0x1a, 0xd2, 0x51, 0xb1, 0x47, 0xf7, 0x9a, 0x74,
	0xd3, 0xeb, 0xdd, 0xaf, 0x4c, 0xa7, 0x5d, 0x96, 0xf6, 0xee, 0xb3, 0x38, 0xff, 0x15, 0xc3, 0x93,
	0xb1, 0xfc, 0x8a, 0xda, 0xbd, 0xaa, 0x08, 0x37, 0x90, 0x5c, 0xb4, 0x54, 0x3a, 0xc2, 0xff, 0x47,
	0x9d, 0xe6, 0x77, 0x95, 0x1e, 0x1c, 0x5b, 0x3d, 0xf2, 0x80, 0xeb, 0x46, 0x4e, 0x08, 0x78, 0x0f,
	0xd1, 0x9a, 0x1c, 0xbe, 0xbc, 0xbb, 0x85, 0x9f, 0xec, 0x7d, 0x67, 0x5f, 0x5b, 0x9a, 0x66, 0xf6,
	0x07, 0x88, 0xdf, 0xf9, 0xc7, 0x3e, 0xc9, 0xf0, 0x0d, 0x2c, 0xd6, 0x6a, 0x4f, 0x6f, 0xcb, 0xea,
	0xd3, 0x24, 0x80, 0x8f, 0x09, 0x7f, 0xf5, 0xde, 0xfc, 0x09, 0x00, 0x00, 0xff, 0xff, 0x28, 0x8a,
	0x6f, 0xe8, 0x18, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for ContainerService service

type ContainerServiceClient interface {
	Create(ctx context.Context, in *Container, opts ...client.CallOption) (*Response, error)
	Update(ctx context.Context, in *Container, opts ...client.CallOption) (*Response, error)
	Get(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	Use(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	Page(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	GiveBack(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
}

type containerServiceClient struct {
	c           client.Client
	serviceName string
}

func NewContainerServiceClient(serviceName string, c client.Client) ContainerServiceClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "daymenu.shipping.srv.container"
	}
	return &containerServiceClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *containerServiceClient) Create(ctx context.Context, in *Container, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "ContainerService.Create", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *containerServiceClient) Update(ctx context.Context, in *Container, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "ContainerService.Update", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *containerServiceClient) Get(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "ContainerService.Get", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *containerServiceClient) Use(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "ContainerService.Use", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *containerServiceClient) Page(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "ContainerService.Page", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *containerServiceClient) GiveBack(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "ContainerService.GiveBack", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ContainerService service

type ContainerServiceHandler interface {
	Create(context.Context, *Container, *Response) error
	Update(context.Context, *Container, *Response) error
	Get(context.Context, *Request, *Response) error
	Use(context.Context, *Request, *Response) error
	Page(context.Context, *Request, *Response) error
	GiveBack(context.Context, *Request, *Response) error
}

func RegisterContainerServiceHandler(s server.Server, hdlr ContainerServiceHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&ContainerService{hdlr}, opts...))
}

type ContainerService struct {
	ContainerServiceHandler
}

func (h *ContainerService) Create(ctx context.Context, in *Container, out *Response) error {
	return h.ContainerServiceHandler.Create(ctx, in, out)
}

func (h *ContainerService) Update(ctx context.Context, in *Container, out *Response) error {
	return h.ContainerServiceHandler.Update(ctx, in, out)
}

func (h *ContainerService) Get(ctx context.Context, in *Request, out *Response) error {
	return h.ContainerServiceHandler.Get(ctx, in, out)
}

func (h *ContainerService) Use(ctx context.Context, in *Request, out *Response) error {
	return h.ContainerServiceHandler.Use(ctx, in, out)
}

func (h *ContainerService) Page(ctx context.Context, in *Request, out *Response) error {
	return h.ContainerServiceHandler.Page(ctx, in, out)
}

func (h *ContainerService) GiveBack(ctx context.Context, in *Request, out *Response) error {
	return h.ContainerServiceHandler.GiveBack(ctx, in, out)
}