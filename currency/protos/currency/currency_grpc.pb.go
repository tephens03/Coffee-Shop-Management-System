// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.3
// source: protos/currency.proto

package currency

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// CurrencyClient is the client API for Currency service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CurrencyClient interface {
	GetRate(ctx context.Context, in *RateRequest, opts ...grpc.CallOption) (*RateResponse, error)
	SubscribeRate(ctx context.Context, opts ...grpc.CallOption) (Currency_SubscribeRateClient, error)
}

type currencyClient struct {
	cc grpc.ClientConnInterface
}

func NewCurrencyClient(cc grpc.ClientConnInterface) CurrencyClient {
	return &currencyClient{cc}
}

func (c *currencyClient) GetRate(ctx context.Context, in *RateRequest, opts ...grpc.CallOption) (*RateResponse, error) {
	out := new(RateResponse)
	err := c.cc.Invoke(ctx, "/Currency/GetRate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *currencyClient) SubscribeRate(ctx context.Context, opts ...grpc.CallOption) (Currency_SubscribeRateClient, error) {
	stream, err := c.cc.NewStream(ctx, &Currency_ServiceDesc.Streams[0], "/Currency/SubscribeRate", opts...)
	if err != nil {
		return nil, err
	}
	x := &currencySubscribeRateClient{stream}
	return x, nil
}

type Currency_SubscribeRateClient interface {
	Send(*RateRequest) error
	Recv() (*RateResponse, error)
	grpc.ClientStream
}

type currencySubscribeRateClient struct {
	grpc.ClientStream
}

func (x *currencySubscribeRateClient) Send(m *RateRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *currencySubscribeRateClient) Recv() (*RateResponse, error) {
	m := new(RateResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CurrencyServer is the server API for Currency service.
// All implementations should embed UnimplementedCurrencyServer
// for forward compatibility
type CurrencyServer interface {
	GetRate(context.Context, *RateRequest) (*RateResponse, error)
	SubscribeRate(Currency_SubscribeRateServer) error
}

// UnimplementedCurrencyServer should be embedded to have forward compatible implementations.
type UnimplementedCurrencyServer struct {
}

func (UnimplementedCurrencyServer) GetRate(context.Context, *RateRequest) (*RateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRate not implemented")
}
func (UnimplementedCurrencyServer) SubscribeRate(Currency_SubscribeRateServer) error {
	return status.Errorf(codes.Unimplemented, "method SubscribeRate not implemented")
}

// UnsafeCurrencyServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CurrencyServer will
// result in compilation errors.
type UnsafeCurrencyServer interface {
	mustEmbedUnimplementedCurrencyServer()
}

func RegisterCurrencyServer(s grpc.ServiceRegistrar, srv CurrencyServer) {
	s.RegisterService(&Currency_ServiceDesc, srv)
}

func _Currency_GetRate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CurrencyServer).GetRate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Currency/GetRate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CurrencyServer).GetRate(ctx, req.(*RateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Currency_SubscribeRate_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CurrencyServer).SubscribeRate(&currencySubscribeRateServer{stream})
}

type Currency_SubscribeRateServer interface {
	Send(*RateResponse) error
	Recv() (*RateRequest, error)
	grpc.ServerStream
}

type currencySubscribeRateServer struct {
	grpc.ServerStream
}

func (x *currencySubscribeRateServer) Send(m *RateResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *currencySubscribeRateServer) Recv() (*RateRequest, error) {
	m := new(RateRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Currency_ServiceDesc is the grpc.ServiceDesc for Currency service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Currency_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Currency",
	HandlerType: (*CurrencyServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRate",
			Handler:    _Currency_GetRate_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SubscribeRate",
			Handler:       _Currency_SubscribeRate_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "protos/currency.proto",
}