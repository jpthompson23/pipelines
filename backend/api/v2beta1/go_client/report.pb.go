// Copyright 2018 The Kubeflow Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.20.3
// source: backend/api/v2beta1/report.proto

package go_client

import (
	context "context"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ReportWorkflowRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Workflow is a workflow custom resource marshalled into a json string.
	Workflow string `protobuf:"bytes,1,opt,name=workflow,proto3" json:"workflow,omitempty"`
}

func (x *ReportWorkflowRequest) Reset() {
	*x = ReportWorkflowRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_backend_api_v2beta1_report_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReportWorkflowRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportWorkflowRequest) ProtoMessage() {}

func (x *ReportWorkflowRequest) ProtoReflect() protoreflect.Message {
	mi := &file_backend_api_v2beta1_report_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportWorkflowRequest.ProtoReflect.Descriptor instead.
func (*ReportWorkflowRequest) Descriptor() ([]byte, []int) {
	return file_backend_api_v2beta1_report_proto_rawDescGZIP(), []int{0}
}

func (x *ReportWorkflowRequest) GetWorkflow() string {
	if x != nil {
		return x.Workflow
	}
	return ""
}

type ReportScheduledWorkflowRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ScheduledWorkflow a ScheduledWorkflow resource marshalled into a json string.
	ScheduledWorkflow string `protobuf:"bytes,1,opt,name=scheduled_workflow,json=scheduledWorkflow,proto3" json:"scheduled_workflow,omitempty"`
}

func (x *ReportScheduledWorkflowRequest) Reset() {
	*x = ReportScheduledWorkflowRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_backend_api_v2beta1_report_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReportScheduledWorkflowRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportScheduledWorkflowRequest) ProtoMessage() {}

func (x *ReportScheduledWorkflowRequest) ProtoReflect() protoreflect.Message {
	mi := &file_backend_api_v2beta1_report_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportScheduledWorkflowRequest.ProtoReflect.Descriptor instead.
func (*ReportScheduledWorkflowRequest) Descriptor() ([]byte, []int) {
	return file_backend_api_v2beta1_report_proto_rawDescGZIP(), []int{1}
}

func (x *ReportScheduledWorkflowRequest) GetScheduledWorkflow() string {
	if x != nil {
		return x.ScheduledWorkflow
	}
	return ""
}

var File_backend_api_v2beta1_report_proto protoreflect.FileDescriptor

var file_backend_api_v2beta1_report_proto_rawDesc = []byte{
	0x0a, 0x20, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x32,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x26, 0x6b, 0x75, 0x62, 0x65, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x70, 0x69, 0x70,
	0x65, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x33, 0x0a, 0x15, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x57,
	0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a,
	0x0a, 0x08, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x22, 0x4f, 0x0a, 0x1e, 0x52, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x64, 0x57, 0x6f, 0x72,
	0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x12,
	0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x64, 0x5f, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c,
	0x6f, 0x77, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75,
	0x6c, 0x65, 0x64, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x32, 0xde, 0x02, 0x0a, 0x0d,
	0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x92, 0x01,
	0x0a, 0x0e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77,
	0x12, 0x3d, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x70, 0x69, 0x70, 0x65,
	0x6c, 0x69, 0x6e, 0x65, 0x73, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x29, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x23, 0x3a,
	0x08, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x22, 0x17, 0x2f, 0x61, 0x70, 0x69, 0x73,
	0x2f, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f,
	0x77, 0x73, 0x12, 0xb7, 0x01, 0x0a, 0x17, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x53, 0x63, 0x68,
	0x65, 0x64, 0x75, 0x6c, 0x65, 0x64, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x12, 0x46,
	0x2e, 0x6b, 0x75, 0x62, 0x65, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69,
	0x6e, 0x65, 0x73, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x53, 0x63,
	0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x64, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x3c,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x36, 0x3a, 0x12, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65,
	0x64, 0x5f, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x22, 0x20, 0x2f, 0x61, 0x70, 0x69,
	0x73, 0x2f, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75,
	0x6c, 0x65, 0x64, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x73, 0x42, 0x3d, 0x5a, 0x3b,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x75, 0x62, 0x65, 0x66,
	0x6c, 0x6f, 0x77, 0x2f, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x2f, 0x62, 0x61,
	0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2f, 0x67, 0x6f, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_backend_api_v2beta1_report_proto_rawDescOnce sync.Once
	file_backend_api_v2beta1_report_proto_rawDescData = file_backend_api_v2beta1_report_proto_rawDesc
)

func file_backend_api_v2beta1_report_proto_rawDescGZIP() []byte {
	file_backend_api_v2beta1_report_proto_rawDescOnce.Do(func() {
		file_backend_api_v2beta1_report_proto_rawDescData = protoimpl.X.CompressGZIP(file_backend_api_v2beta1_report_proto_rawDescData)
	})
	return file_backend_api_v2beta1_report_proto_rawDescData
}

var file_backend_api_v2beta1_report_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_backend_api_v2beta1_report_proto_goTypes = []interface{}{
	(*ReportWorkflowRequest)(nil),          // 0: kubeflow.pipelines.backend.api.v2beta1.ReportWorkflowRequest
	(*ReportScheduledWorkflowRequest)(nil), // 1: kubeflow.pipelines.backend.api.v2beta1.ReportScheduledWorkflowRequest
	(*emptypb.Empty)(nil),                  // 2: google.protobuf.Empty
}
var file_backend_api_v2beta1_report_proto_depIdxs = []int32{
	0, // 0: kubeflow.pipelines.backend.api.v2beta1.ReportService.ReportWorkflow:input_type -> kubeflow.pipelines.backend.api.v2beta1.ReportWorkflowRequest
	1, // 1: kubeflow.pipelines.backend.api.v2beta1.ReportService.ReportScheduledWorkflow:input_type -> kubeflow.pipelines.backend.api.v2beta1.ReportScheduledWorkflowRequest
	2, // 2: kubeflow.pipelines.backend.api.v2beta1.ReportService.ReportWorkflow:output_type -> google.protobuf.Empty
	2, // 3: kubeflow.pipelines.backend.api.v2beta1.ReportService.ReportScheduledWorkflow:output_type -> google.protobuf.Empty
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_backend_api_v2beta1_report_proto_init() }
func file_backend_api_v2beta1_report_proto_init() {
	if File_backend_api_v2beta1_report_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_backend_api_v2beta1_report_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReportWorkflowRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_backend_api_v2beta1_report_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReportScheduledWorkflowRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_backend_api_v2beta1_report_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_backend_api_v2beta1_report_proto_goTypes,
		DependencyIndexes: file_backend_api_v2beta1_report_proto_depIdxs,
		MessageInfos:      file_backend_api_v2beta1_report_proto_msgTypes,
	}.Build()
	File_backend_api_v2beta1_report_proto = out.File
	file_backend_api_v2beta1_report_proto_rawDesc = nil
	file_backend_api_v2beta1_report_proto_goTypes = nil
	file_backend_api_v2beta1_report_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ReportServiceClient is the client API for ReportService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ReportServiceClient interface {
	ReportWorkflow(ctx context.Context, in *ReportWorkflowRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ReportScheduledWorkflow(ctx context.Context, in *ReportScheduledWorkflowRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type reportServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewReportServiceClient(cc grpc.ClientConnInterface) ReportServiceClient {
	return &reportServiceClient{cc}
}

func (c *reportServiceClient) ReportWorkflow(ctx context.Context, in *ReportWorkflowRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/kubeflow.pipelines.backend.api.v2beta1.ReportService/ReportWorkflow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reportServiceClient) ReportScheduledWorkflow(ctx context.Context, in *ReportScheduledWorkflowRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/kubeflow.pipelines.backend.api.v2beta1.ReportService/ReportScheduledWorkflow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReportServiceServer is the server API for ReportService service.
type ReportServiceServer interface {
	ReportWorkflow(context.Context, *ReportWorkflowRequest) (*emptypb.Empty, error)
	ReportScheduledWorkflow(context.Context, *ReportScheduledWorkflowRequest) (*emptypb.Empty, error)
}

// UnimplementedReportServiceServer can be embedded to have forward compatible implementations.
type UnimplementedReportServiceServer struct {
}

func (*UnimplementedReportServiceServer) ReportWorkflow(context.Context, *ReportWorkflowRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReportWorkflow not implemented")
}
func (*UnimplementedReportServiceServer) ReportScheduledWorkflow(context.Context, *ReportScheduledWorkflowRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReportScheduledWorkflow not implemented")
}

func RegisterReportServiceServer(s *grpc.Server, srv ReportServiceServer) {
	s.RegisterService(&_ReportService_serviceDesc, srv)
}

func _ReportService_ReportWorkflow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReportWorkflowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReportServiceServer).ReportWorkflow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kubeflow.pipelines.backend.api.v2beta1.ReportService/ReportWorkflow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReportServiceServer).ReportWorkflow(ctx, req.(*ReportWorkflowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReportService_ReportScheduledWorkflow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReportScheduledWorkflowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReportServiceServer).ReportScheduledWorkflow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kubeflow.pipelines.backend.api.v2beta1.ReportService/ReportScheduledWorkflow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReportServiceServer).ReportScheduledWorkflow(ctx, req.(*ReportScheduledWorkflowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ReportService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "kubeflow.pipelines.backend.api.v2beta1.ReportService",
	HandlerType: (*ReportServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ReportWorkflow",
			Handler:    _ReportService_ReportWorkflow_Handler,
		},
		{
			MethodName: "ReportScheduledWorkflow",
			Handler:    _ReportService_ReportScheduledWorkflow_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "backend/api/v2beta1/report.proto",
}
