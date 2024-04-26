// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: metrics/metrics.proto

package metrics

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
	Metrics_CpuMemoryInNamespace_FullMethodName = "/metrics.Metrics/CpuMemoryInNamespace"
	Metrics_CpuMemoryInProject_FullMethodName   = "/metrics.Metrics/CpuMemoryInProject"
	Metrics_TopPod_FullMethodName               = "/metrics.Metrics/TopPod"
	Metrics_StreamTopPod_FullMethodName         = "/metrics.Metrics/StreamTopPod"
)

// MetricsClient is the client API for Metrics service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MetricsClient interface {
	// CpuMemoryInNamespace 名称空间总共使用的 cpu memory
	CpuMemoryInNamespace(ctx context.Context, in *CpuMemoryInNamespaceRequest, opts ...grpc.CallOption) (*CpuMemoryInNamespaceResponse, error)
	// CpuMemoryInProject 项目空间总共使用的 cpu memory
	CpuMemoryInProject(ctx context.Context, in *CpuMemoryInProjectRequest, opts ...grpc.CallOption) (*CpuMemoryInProjectResponse, error)
	// TopPod 获取 pod 的 cpu memory 信息
	TopPod(ctx context.Context, in *TopPodRequest, opts ...grpc.CallOption) (*TopPodResponse, error)
	// StreamTopPod stream 的方式获取 pod 的 cpu memory 信息
	StreamTopPod(ctx context.Context, in *TopPodRequest, opts ...grpc.CallOption) (Metrics_StreamTopPodClient, error)
}

type metricsClient struct {
	cc grpc.ClientConnInterface
}

func NewMetricsClient(cc grpc.ClientConnInterface) MetricsClient {
	return &metricsClient{cc}
}

func (c *metricsClient) CpuMemoryInNamespace(ctx context.Context, in *CpuMemoryInNamespaceRequest, opts ...grpc.CallOption) (*CpuMemoryInNamespaceResponse, error) {
	out := new(CpuMemoryInNamespaceResponse)
	err := c.cc.Invoke(ctx, Metrics_CpuMemoryInNamespace_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metricsClient) CpuMemoryInProject(ctx context.Context, in *CpuMemoryInProjectRequest, opts ...grpc.CallOption) (*CpuMemoryInProjectResponse, error) {
	out := new(CpuMemoryInProjectResponse)
	err := c.cc.Invoke(ctx, Metrics_CpuMemoryInProject_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metricsClient) TopPod(ctx context.Context, in *TopPodRequest, opts ...grpc.CallOption) (*TopPodResponse, error) {
	out := new(TopPodResponse)
	err := c.cc.Invoke(ctx, Metrics_TopPod_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metricsClient) StreamTopPod(ctx context.Context, in *TopPodRequest, opts ...grpc.CallOption) (Metrics_StreamTopPodClient, error) {
	stream, err := c.cc.NewStream(ctx, &Metrics_ServiceDesc.Streams[0], Metrics_StreamTopPod_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &metricsStreamTopPodClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Metrics_StreamTopPodClient interface {
	Recv() (*TopPodResponse, error)
	grpc.ClientStream
}

type metricsStreamTopPodClient struct {
	grpc.ClientStream
}

func (x *metricsStreamTopPodClient) Recv() (*TopPodResponse, error) {
	m := new(TopPodResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MetricsServer is the server API for Metrics service.
// All implementations must embed UnimplementedMetricsServer
// for forward compatibility
type MetricsServer interface {
	// CpuMemoryInNamespace 名称空间总共使用的 cpu memory
	CpuMemoryInNamespace(context.Context, *CpuMemoryInNamespaceRequest) (*CpuMemoryInNamespaceResponse, error)
	// CpuMemoryInProject 项目空间总共使用的 cpu memory
	CpuMemoryInProject(context.Context, *CpuMemoryInProjectRequest) (*CpuMemoryInProjectResponse, error)
	// TopPod 获取 pod 的 cpu memory 信息
	TopPod(context.Context, *TopPodRequest) (*TopPodResponse, error)
	// StreamTopPod stream 的方式获取 pod 的 cpu memory 信息
	StreamTopPod(*TopPodRequest, Metrics_StreamTopPodServer) error
	mustEmbedUnimplementedMetricsServer()
}

// UnimplementedMetricsServer must be embedded to have forward compatible implementations.
type UnimplementedMetricsServer struct {
}

func (UnimplementedMetricsServer) CpuMemoryInNamespace(context.Context, *CpuMemoryInNamespaceRequest) (*CpuMemoryInNamespaceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CpuMemoryInNamespace not implemented")
}
func (UnimplementedMetricsServer) CpuMemoryInProject(context.Context, *CpuMemoryInProjectRequest) (*CpuMemoryInProjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CpuMemoryInProject not implemented")
}
func (UnimplementedMetricsServer) TopPod(context.Context, *TopPodRequest) (*TopPodResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TopPod not implemented")
}
func (UnimplementedMetricsServer) StreamTopPod(*TopPodRequest, Metrics_StreamTopPodServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamTopPod not implemented")
}
func (UnimplementedMetricsServer) mustEmbedUnimplementedMetricsServer() {}

// UnsafeMetricsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MetricsServer will
// result in compilation errors.
type UnsafeMetricsServer interface {
	mustEmbedUnimplementedMetricsServer()
}

func RegisterMetricsServer(s grpc.ServiceRegistrar, srv MetricsServer) {
	s.RegisterService(&Metrics_ServiceDesc, srv)
}

func _Metrics_CpuMemoryInNamespace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CpuMemoryInNamespaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricsServer).CpuMemoryInNamespace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Metrics_CpuMemoryInNamespace_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricsServer).CpuMemoryInNamespace(ctx, req.(*CpuMemoryInNamespaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Metrics_CpuMemoryInProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CpuMemoryInProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricsServer).CpuMemoryInProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Metrics_CpuMemoryInProject_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricsServer).CpuMemoryInProject(ctx, req.(*CpuMemoryInProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Metrics_TopPod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TopPodRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricsServer).TopPod(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Metrics_TopPod_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricsServer).TopPod(ctx, req.(*TopPodRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Metrics_StreamTopPod_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(TopPodRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MetricsServer).StreamTopPod(m, &metricsStreamTopPodServer{stream})
}

type Metrics_StreamTopPodServer interface {
	Send(*TopPodResponse) error
	grpc.ServerStream
}

type metricsStreamTopPodServer struct {
	grpc.ServerStream
}

func (x *metricsStreamTopPodServer) Send(m *TopPodResponse) error {
	return x.ServerStream.SendMsg(m)
}

// Metrics_ServiceDesc is the grpc.ServiceDesc for Metrics service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Metrics_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "metrics.Metrics",
	HandlerType: (*MetricsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CpuMemoryInNamespace",
			Handler:    _Metrics_CpuMemoryInNamespace_Handler,
		},
		{
			MethodName: "CpuMemoryInProject",
			Handler:    _Metrics_CpuMemoryInProject_Handler,
		},
		{
			MethodName: "TopPod",
			Handler:    _Metrics_TopPod_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamTopPod",
			Handler:       _Metrics_StreamTopPod_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "metrics/metrics.proto",
}