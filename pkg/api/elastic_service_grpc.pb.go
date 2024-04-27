// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.26.1
// source: proto/elastic_service.proto

package api

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

// CitySearchClient is the client API for CitySearch service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CitySearchClient interface {
	SearchByName(ctx context.Context, in *CityNameRequest, opts ...grpc.CallOption) (*CitySearchNameResponse, error)
	SearchByCoords(ctx context.Context, in *CoordsRequest, opts ...grpc.CallOption) (*CitySearchCoordsResponse, error)
}

type citySearchClient struct {
	cc grpc.ClientConnInterface
}

func NewCitySearchClient(cc grpc.ClientConnInterface) CitySearchClient {
	return &citySearchClient{cc}
}

func (c *citySearchClient) SearchByName(ctx context.Context, in *CityNameRequest, opts ...grpc.CallOption) (*CitySearchNameResponse, error) {
	out := new(CitySearchNameResponse)
	err := c.cc.Invoke(ctx, "/api.CitySearch/SearchByName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *citySearchClient) SearchByCoords(ctx context.Context, in *CoordsRequest, opts ...grpc.CallOption) (*CitySearchCoordsResponse, error) {
	out := new(CitySearchCoordsResponse)
	err := c.cc.Invoke(ctx, "/api.CitySearch/SearchByCoords", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CitySearchServer is the server API for CitySearch service.
// All implementations must embed UnimplementedCitySearchServer
// for forward compatibility
type CitySearchServer interface {
	SearchByName(context.Context, *CityNameRequest) (*CitySearchNameResponse, error)
	SearchByCoords(context.Context, *CoordsRequest) (*CitySearchCoordsResponse, error)
	mustEmbedUnimplementedCitySearchServer()
}

// UnimplementedCitySearchServer must be embedded to have forward compatible implementations.
type UnimplementedCitySearchServer struct {
}

func (UnimplementedCitySearchServer) SearchByName(context.Context, *CityNameRequest) (*CitySearchNameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchByName not implemented")
}
func (UnimplementedCitySearchServer) SearchByCoords(context.Context, *CoordsRequest) (*CitySearchCoordsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchByCoords not implemented")
}
func (UnimplementedCitySearchServer) mustEmbedUnimplementedCitySearchServer() {}

// UnsafeCitySearchServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CitySearchServer will
// result in compilation errors.
type UnsafeCitySearchServer interface {
	mustEmbedUnimplementedCitySearchServer()
}

func RegisterCitySearchServer(s grpc.ServiceRegistrar, srv CitySearchServer) {
	s.RegisterService(&CitySearch_ServiceDesc, srv)
}

func _CitySearch_SearchByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CityNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CitySearchServer).SearchByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.CitySearch/SearchByName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CitySearchServer).SearchByName(ctx, req.(*CityNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CitySearch_SearchByCoords_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CoordsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CitySearchServer).SearchByCoords(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.CitySearch/SearchByCoords",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CitySearchServer).SearchByCoords(ctx, req.(*CoordsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CitySearch_ServiceDesc is the grpc.ServiceDesc for CitySearch service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CitySearch_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.CitySearch",
	HandlerType: (*CitySearchServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SearchByName",
			Handler:    _CitySearch_SearchByName_Handler,
		},
		{
			MethodName: "SearchByCoords",
			Handler:    _CitySearch_SearchByCoords_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/elastic_service.proto",
}
