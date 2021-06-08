// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1alpha1

import (
	context "context"
	v1alpha1 "github.com/kubeapps/kubeapps/cmd/kubeapps-apis/gen/core/packages/v1alpha1"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// FluxV2PackagesServiceClient is the client API for FluxV2PackagesService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FluxV2PackagesServiceClient interface {
	// GetAvailablePackages returns the available packages managed by the 'fluxv2' plugin
	GetAvailablePackages(ctx context.Context, in *v1alpha1.GetAvailablePackagesRequest, opts ...grpc.CallOption) (*v1alpha1.GetAvailablePackagesResponse, error)
	// GetPackageRepositories returns the repositories managed by the 'fluxv2' plugin
	GetPackageRepositories(ctx context.Context, in *v1alpha1.GetPackageRepositoriesRequest, opts ...grpc.CallOption) (*v1alpha1.GetPackageRepositoriesResponse, error)
	// GetPackageMeta returns the package metadata managed by the 'fluxv2' plugin
	GetPackageMeta(ctx context.Context, in *v1alpha1.GetPackageMetaRequest, opts ...grpc.CallOption) (*v1alpha1.GetPackageMetaResponse, error)
}

type fluxV2PackagesServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFluxV2PackagesServiceClient(cc grpc.ClientConnInterface) FluxV2PackagesServiceClient {
	return &fluxV2PackagesServiceClient{cc}
}

func (c *fluxV2PackagesServiceClient) GetAvailablePackages(ctx context.Context, in *v1alpha1.GetAvailablePackagesRequest, opts ...grpc.CallOption) (*v1alpha1.GetAvailablePackagesResponse, error) {
	out := new(v1alpha1.GetAvailablePackagesResponse)
	err := c.cc.Invoke(ctx, "/kubeappsapis.plugins.fluxv2.packages.v1alpha1.FluxV2PackagesService/GetAvailablePackages", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fluxV2PackagesServiceClient) GetPackageRepositories(ctx context.Context, in *v1alpha1.GetPackageRepositoriesRequest, opts ...grpc.CallOption) (*v1alpha1.GetPackageRepositoriesResponse, error) {
	out := new(v1alpha1.GetPackageRepositoriesResponse)
	err := c.cc.Invoke(ctx, "/kubeappsapis.plugins.fluxv2.packages.v1alpha1.FluxV2PackagesService/GetPackageRepositories", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fluxV2PackagesServiceClient) GetPackageMeta(ctx context.Context, in *v1alpha1.GetPackageMetaRequest, opts ...grpc.CallOption) (*v1alpha1.GetPackageMetaResponse, error) {
	out := new(v1alpha1.GetPackageMetaResponse)
	err := c.cc.Invoke(ctx, "/kubeappsapis.plugins.fluxv2.packages.v1alpha1.FluxV2PackagesService/GetPackageMeta", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FluxV2PackagesServiceServer is the server API for FluxV2PackagesService service.
// All implementations should embed UnimplementedFluxV2PackagesServiceServer
// for forward compatibility
type FluxV2PackagesServiceServer interface {
	// GetAvailablePackages returns the available packages managed by the 'fluxv2' plugin
	GetAvailablePackages(context.Context, *v1alpha1.GetAvailablePackagesRequest) (*v1alpha1.GetAvailablePackagesResponse, error)
	// GetPackageRepositories returns the repositories managed by the 'fluxv2' plugin
	GetPackageRepositories(context.Context, *v1alpha1.GetPackageRepositoriesRequest) (*v1alpha1.GetPackageRepositoriesResponse, error)
	// GetPackageMeta returns the package metadata managed by the 'fluxv2' plugin
	GetPackageMeta(context.Context, *v1alpha1.GetPackageMetaRequest) (*v1alpha1.GetPackageMetaResponse, error)
}

// UnimplementedFluxV2PackagesServiceServer should be embedded to have forward compatible implementations.
type UnimplementedFluxV2PackagesServiceServer struct {
}

func (UnimplementedFluxV2PackagesServiceServer) GetAvailablePackages(context.Context, *v1alpha1.GetAvailablePackagesRequest) (*v1alpha1.GetAvailablePackagesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAvailablePackages not implemented")
}
func (UnimplementedFluxV2PackagesServiceServer) GetPackageRepositories(context.Context, *v1alpha1.GetPackageRepositoriesRequest) (*v1alpha1.GetPackageRepositoriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPackageRepositories not implemented")
}
func (UnimplementedFluxV2PackagesServiceServer) GetPackageMeta(context.Context, *v1alpha1.GetPackageMetaRequest) (*v1alpha1.GetPackageMetaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPackageMeta not implemented")
}

// UnsafeFluxV2PackagesServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FluxV2PackagesServiceServer will
// result in compilation errors.
type UnsafeFluxV2PackagesServiceServer interface {
	mustEmbedUnimplementedFluxV2PackagesServiceServer()
}

func RegisterFluxV2PackagesServiceServer(s grpc.ServiceRegistrar, srv FluxV2PackagesServiceServer) {
	s.RegisterService(&FluxV2PackagesService_ServiceDesc, srv)
}

func _FluxV2PackagesService_GetAvailablePackages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1alpha1.GetAvailablePackagesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FluxV2PackagesServiceServer).GetAvailablePackages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kubeappsapis.plugins.fluxv2.packages.v1alpha1.FluxV2PackagesService/GetAvailablePackages",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FluxV2PackagesServiceServer).GetAvailablePackages(ctx, req.(*v1alpha1.GetAvailablePackagesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FluxV2PackagesService_GetPackageRepositories_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1alpha1.GetPackageRepositoriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FluxV2PackagesServiceServer).GetPackageRepositories(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kubeappsapis.plugins.fluxv2.packages.v1alpha1.FluxV2PackagesService/GetPackageRepositories",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FluxV2PackagesServiceServer).GetPackageRepositories(ctx, req.(*v1alpha1.GetPackageRepositoriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FluxV2PackagesService_GetPackageMeta_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1alpha1.GetPackageMetaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FluxV2PackagesServiceServer).GetPackageMeta(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kubeappsapis.plugins.fluxv2.packages.v1alpha1.FluxV2PackagesService/GetPackageMeta",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FluxV2PackagesServiceServer).GetPackageMeta(ctx, req.(*v1alpha1.GetPackageMetaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FluxV2PackagesService_ServiceDesc is the grpc.ServiceDesc for FluxV2PackagesService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FluxV2PackagesService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "kubeappsapis.plugins.fluxv2.packages.v1alpha1.FluxV2PackagesService",
	HandlerType: (*FluxV2PackagesServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAvailablePackages",
			Handler:    _FluxV2PackagesService_GetAvailablePackages_Handler,
		},
		{
			MethodName: "GetPackageRepositories",
			Handler:    _FluxV2PackagesService_GetPackageRepositories_Handler,
		},
		{
			MethodName: "GetPackageMeta",
			Handler:    _FluxV2PackagesService_GetPackageMeta_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "kubeappsapis/plugins/fluxv2/packages/v1alpha1/fluxv2.proto",
}
