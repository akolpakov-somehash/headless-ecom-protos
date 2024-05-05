// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.1
// source: sale/order.proto

package salev1

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

const (
	OrderService_GetOrders_FullMethodName  = "/sale.OrderService/GetOrders"
	OrderService_GetOrder_FullMethodName   = "/sale.OrderService/GetOrder"
	OrderService_PlaceOrder_FullMethodName = "/sale.OrderService/PlaceOrder"
)

// OrderServiceClient is the client API for OrderService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrderServiceClient interface {
	GetOrders(ctx context.Context, in *CustomerId, opts ...grpc.CallOption) (*OrderList, error)
	GetOrder(ctx context.Context, in *OrderId, opts ...grpc.CallOption) (*Order, error)
	PlaceOrder(ctx context.Context, in *CustomerId, opts ...grpc.CallOption) (OrderService_PlaceOrderClient, error)
}

type orderServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOrderServiceClient(cc grpc.ClientConnInterface) OrderServiceClient {
	return &orderServiceClient{cc}
}

func (c *orderServiceClient) GetOrders(ctx context.Context, in *CustomerId, opts ...grpc.CallOption) (*OrderList, error) {
	out := new(OrderList)
	err := c.cc.Invoke(ctx, OrderService_GetOrders_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) GetOrder(ctx context.Context, in *OrderId, opts ...grpc.CallOption) (*Order, error) {
	out := new(Order)
	err := c.cc.Invoke(ctx, OrderService_GetOrder_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) PlaceOrder(ctx context.Context, in *CustomerId, opts ...grpc.CallOption) (OrderService_PlaceOrderClient, error) {
	stream, err := c.cc.NewStream(ctx, &OrderService_ServiceDesc.Streams[0], OrderService_PlaceOrder_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &orderServicePlaceOrderClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type OrderService_PlaceOrderClient interface {
	Recv() (*ProcessStatus, error)
	grpc.ClientStream
}

type orderServicePlaceOrderClient struct {
	grpc.ClientStream
}

func (x *orderServicePlaceOrderClient) Recv() (*ProcessStatus, error) {
	m := new(ProcessStatus)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// OrderServiceServer is the server API for OrderService service.
// All implementations must embed UnimplementedOrderServiceServer
// for forward compatibility
type OrderServiceServer interface {
	GetOrders(context.Context, *CustomerId) (*OrderList, error)
	GetOrder(context.Context, *OrderId) (*Order, error)
	PlaceOrder(*CustomerId, OrderService_PlaceOrderServer) error
	mustEmbedUnimplementedOrderServiceServer()
}

// UnimplementedOrderServiceServer must be embedded to have forward compatible implementations.
type UnimplementedOrderServiceServer struct {
}

func (UnimplementedOrderServiceServer) GetOrders(context.Context, *CustomerId) (*OrderList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrders not implemented")
}
func (UnimplementedOrderServiceServer) GetOrder(context.Context, *OrderId) (*Order, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrder not implemented")
}
func (UnimplementedOrderServiceServer) PlaceOrder(*CustomerId, OrderService_PlaceOrderServer) error {
	return status.Errorf(codes.Unimplemented, "method PlaceOrder not implemented")
}
func (UnimplementedOrderServiceServer) mustEmbedUnimplementedOrderServiceServer() {}

// UnsafeOrderServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrderServiceServer will
// result in compilation errors.
type UnsafeOrderServiceServer interface {
	mustEmbedUnimplementedOrderServiceServer()
}

func RegisterOrderServiceServer(s grpc.ServiceRegistrar, srv OrderServiceServer) {
	s.RegisterService(&OrderService_ServiceDesc, srv)
}

func _OrderService_GetOrders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CustomerId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).GetOrders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderService_GetOrders_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).GetOrders(ctx, req.(*CustomerId))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_GetOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).GetOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderService_GetOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).GetOrder(ctx, req.(*OrderId))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_PlaceOrder_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(CustomerId)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(OrderServiceServer).PlaceOrder(m, &orderServicePlaceOrderServer{stream})
}

type OrderService_PlaceOrderServer interface {
	Send(*ProcessStatus) error
	grpc.ServerStream
}

type orderServicePlaceOrderServer struct {
	grpc.ServerStream
}

func (x *orderServicePlaceOrderServer) Send(m *ProcessStatus) error {
	return x.ServerStream.SendMsg(m)
}

// OrderService_ServiceDesc is the grpc.ServiceDesc for OrderService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OrderService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sale.OrderService",
	HandlerType: (*OrderServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetOrders",
			Handler:    _OrderService_GetOrders_Handler,
		},
		{
			MethodName: "GetOrder",
			Handler:    _OrderService_GetOrder_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PlaceOrder",
			Handler:       _OrderService_PlaceOrder_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "sale/order.proto",
}
