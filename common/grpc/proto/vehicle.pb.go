// Code generated by protoc-gen-go. DO NOT EDIT.
// source: vehicle.proto

package grpc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

type Vehicle struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	PlatNumber           string   `protobuf:"bytes,2,opt,name=plat_number,json=platNumber,proto3" json:"plat_number,omitempty"`
	VehicleType          string   `protobuf:"bytes,3,opt,name=vehicle_type,json=vehicleType,proto3" json:"vehicle_type,omitempty"`
	TimeCheckIn          string   `protobuf:"bytes,4,opt,name=time_check_in,json=timeCheckIn,proto3" json:"time_check_in,omitempty"`
	TimeCheckOut         string   `protobuf:"bytes,5,opt,name=time_check_out,json=timeCheckOut,proto3" json:"time_check_out,omitempty"`
	DateCheckIn          string   `protobuf:"bytes,6,opt,name=date_check_in,json=dateCheckIn,proto3" json:"date_check_in,omitempty"`
	DateCheckOut         string   `protobuf:"bytes,7,opt,name=date_check_out,json=dateCheckOut,proto3" json:"date_check_out,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Vehicle) Reset()         { *m = Vehicle{} }
func (m *Vehicle) String() string { return proto.CompactTextString(m) }
func (*Vehicle) ProtoMessage()    {}
func (*Vehicle) Descriptor() ([]byte, []int) {
	return fileDescriptor_vehicle_09112e6c9d602995, []int{0}
}
func (m *Vehicle) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Vehicle.Unmarshal(m, b)
}
func (m *Vehicle) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Vehicle.Marshal(b, m, deterministic)
}
func (dst *Vehicle) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Vehicle.Merge(dst, src)
}
func (m *Vehicle) XXX_Size() int {
	return xxx_messageInfo_Vehicle.Size(m)
}
func (m *Vehicle) XXX_DiscardUnknown() {
	xxx_messageInfo_Vehicle.DiscardUnknown(m)
}

var xxx_messageInfo_Vehicle proto.InternalMessageInfo

func (m *Vehicle) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Vehicle) GetPlatNumber() string {
	if m != nil {
		return m.PlatNumber
	}
	return ""
}

func (m *Vehicle) GetVehicleType() string {
	if m != nil {
		return m.VehicleType
	}
	return ""
}

func (m *Vehicle) GetTimeCheckIn() string {
	if m != nil {
		return m.TimeCheckIn
	}
	return ""
}

func (m *Vehicle) GetTimeCheckOut() string {
	if m != nil {
		return m.TimeCheckOut
	}
	return ""
}

func (m *Vehicle) GetDateCheckIn() string {
	if m != nil {
		return m.DateCheckIn
	}
	return ""
}

func (m *Vehicle) GetDateCheckOut() string {
	if m != nil {
		return m.DateCheckOut
	}
	return ""
}

type Filter struct {
	PlatNumber           string   `protobuf:"bytes,1,opt,name=plat_number,json=platNumber,proto3" json:"plat_number,omitempty"`
	VehicleType          string   `protobuf:"bytes,2,opt,name=vehicle_type,json=vehicleType,proto3" json:"vehicle_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Filter) Reset()         { *m = Filter{} }
func (m *Filter) String() string { return proto.CompactTextString(m) }
func (*Filter) ProtoMessage()    {}
func (*Filter) Descriptor() ([]byte, []int) {
	return fileDescriptor_vehicle_09112e6c9d602995, []int{1}
}
func (m *Filter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Filter.Unmarshal(m, b)
}
func (m *Filter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Filter.Marshal(b, m, deterministic)
}
func (dst *Filter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Filter.Merge(dst, src)
}
func (m *Filter) XXX_Size() int {
	return xxx_messageInfo_Filter.Size(m)
}
func (m *Filter) XXX_DiscardUnknown() {
	xxx_messageInfo_Filter.DiscardUnknown(m)
}

var xxx_messageInfo_Filter proto.InternalMessageInfo

func (m *Filter) GetPlatNumber() string {
	if m != nil {
		return m.PlatNumber
	}
	return ""
}

func (m *Filter) GetVehicleType() string {
	if m != nil {
		return m.VehicleType
	}
	return ""
}

type VehicleInTownRequest struct {
	TownCode             string   `protobuf:"bytes,1,opt,name=town_code,json=townCode,proto3" json:"town_code,omitempty"`
	Filter               *Filter  `protobuf:"bytes,2,opt,name=filter,proto3" json:"filter,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VehicleInTownRequest) Reset()         { *m = VehicleInTownRequest{} }
func (m *VehicleInTownRequest) String() string { return proto.CompactTextString(m) }
func (*VehicleInTownRequest) ProtoMessage()    {}
func (*VehicleInTownRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_vehicle_09112e6c9d602995, []int{2}
}
func (m *VehicleInTownRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VehicleInTownRequest.Unmarshal(m, b)
}
func (m *VehicleInTownRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VehicleInTownRequest.Marshal(b, m, deterministic)
}
func (dst *VehicleInTownRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VehicleInTownRequest.Merge(dst, src)
}
func (m *VehicleInTownRequest) XXX_Size() int {
	return xxx_messageInfo_VehicleInTownRequest.Size(m)
}
func (m *VehicleInTownRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_VehicleInTownRequest.DiscardUnknown(m)
}

var xxx_messageInfo_VehicleInTownRequest proto.InternalMessageInfo

func (m *VehicleInTownRequest) GetTownCode() string {
	if m != nil {
		return m.TownCode
	}
	return ""
}

func (m *VehicleInTownRequest) GetFilter() *Filter {
	if m != nil {
		return m.Filter
	}
	return nil
}

type VehicleList struct {
	List                 []*Vehicle `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *VehicleList) Reset()         { *m = VehicleList{} }
func (m *VehicleList) String() string { return proto.CompactTextString(m) }
func (*VehicleList) ProtoMessage()    {}
func (*VehicleList) Descriptor() ([]byte, []int) {
	return fileDescriptor_vehicle_09112e6c9d602995, []int{3}
}
func (m *VehicleList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VehicleList.Unmarshal(m, b)
}
func (m *VehicleList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VehicleList.Marshal(b, m, deterministic)
}
func (dst *VehicleList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VehicleList.Merge(dst, src)
}
func (m *VehicleList) XXX_Size() int {
	return xxx_messageInfo_VehicleList.Size(m)
}
func (m *VehicleList) XXX_DiscardUnknown() {
	xxx_messageInfo_VehicleList.DiscardUnknown(m)
}

var xxx_messageInfo_VehicleList proto.InternalMessageInfo

func (m *VehicleList) GetList() []*Vehicle {
	if m != nil {
		return m.List
	}
	return nil
}

func init() {
	proto.RegisterType((*Vehicle)(nil), "grpc.Vehicle")
	proto.RegisterType((*Filter)(nil), "grpc.Filter")
	proto.RegisterType((*VehicleInTownRequest)(nil), "grpc.VehicleInTownRequest")
	proto.RegisterType((*VehicleList)(nil), "grpc.VehicleList")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// VehiclesClient is the client API for Vehicles service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type VehiclesClient interface {
	GetVehicleInTown(ctx context.Context, in *VehicleInTownRequest, opts ...grpc.CallOption) (*VehicleList, error)
}

type vehiclesClient struct {
	cc *grpc.ClientConn
}

func NewVehiclesClient(cc *grpc.ClientConn) VehiclesClient {
	return &vehiclesClient{cc}
}

func (c *vehiclesClient) GetVehicleInTown(ctx context.Context, in *VehicleInTownRequest, opts ...grpc.CallOption) (*VehicleList, error) {
	out := new(VehicleList)
	err := c.cc.Invoke(ctx, "/grpc.Vehicles/GetVehicleInTown", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VehiclesServer is the server API for Vehicles service.
type VehiclesServer interface {
	GetVehicleInTown(context.Context, *VehicleInTownRequest) (*VehicleList, error)
}

func RegisterVehiclesServer(s *grpc.Server, srv VehiclesServer) {
	s.RegisterService(&_Vehicles_serviceDesc, srv)
}

func _Vehicles_GetVehicleInTown_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VehicleInTownRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VehiclesServer).GetVehicleInTown(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.Vehicles/GetVehicleInTown",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VehiclesServer).GetVehicleInTown(ctx, req.(*VehicleInTownRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Vehicles_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.Vehicles",
	HandlerType: (*VehiclesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetVehicleInTown",
			Handler:    _Vehicles_GetVehicleInTown_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "vehicle.proto",
}

func init() { proto.RegisterFile("vehicle.proto", fileDescriptor_vehicle_09112e6c9d602995) }

var fileDescriptor_vehicle_09112e6c9d602995 = []byte{
	// 322 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xcd, 0x4e, 0xc2, 0x50,
	0x10, 0x85, 0x6d, 0xc1, 0x02, 0x53, 0x20, 0x7a, 0xe3, 0xa2, 0xc1, 0x85, 0xd0, 0xb0, 0x60, 0x45,
	0x0c, 0x3e, 0x02, 0x89, 0x86, 0x84, 0xd8, 0xa4, 0x21, 0x26, 0xae, 0x1a, 0x68, 0x47, 0xb9, 0xb1,
	0xf4, 0xd6, 0x76, 0x2a, 0xe1, 0xb1, 0x7d, 0x03, 0x73, 0x7f, 0x42, 0x5a, 0x5d, 0xb8, 0xfd, 0xe6,
	0xcc, 0xd7, 0x33, 0x6d, 0x61, 0xf0, 0x85, 0x7b, 0x1e, 0xa7, 0x38, 0xcf, 0x0b, 0x41, 0x82, 0xb5,
	0xdf, 0x8b, 0x3c, 0xf6, 0xbf, 0x2d, 0xe8, 0xbc, 0x68, 0xce, 0x86, 0x60, 0xf3, 0xc4, 0xb3, 0xc6,
	0xd6, 0xac, 0x17, 0xda, 0x3c, 0x61, 0x77, 0xe0, 0xe6, 0xe9, 0x96, 0xa2, 0xac, 0x3a, 0xec, 0xb0,
	0xf0, 0x6c, 0x35, 0x00, 0x89, 0x9e, 0x15, 0x61, 0x13, 0xe8, 0x1b, 0x67, 0x44, 0xa7, 0x1c, 0xbd,
	0x96, 0x4a, 0xb8, 0x86, 0x6d, 0x4e, 0x39, 0x32, 0x1f, 0x06, 0xc4, 0x0f, 0x18, 0xc5, 0x7b, 0x8c,
	0x3f, 0x22, 0x9e, 0x79, 0x6d, 0x9d, 0x91, 0x70, 0x29, 0xd9, 0x2a, 0x63, 0x53, 0x18, 0xd6, 0x32,
	0xa2, 0x22, 0xef, 0x52, 0x85, 0xfa, 0xe7, 0x50, 0x50, 0x91, 0x34, 0x25, 0x5b, 0xaa, 0x99, 0x1c,
	0x6d, 0x92, 0xb0, 0x66, 0xaa, 0x65, 0xa4, 0xa9, 0xa3, 0x4d, 0xe7, 0x50, 0x50, 0x91, 0xbf, 0x06,
	0xe7, 0x91, 0xa7, 0x84, 0xc5, 0xef, 0x0b, 0xad, 0x7f, 0x2f, 0xb4, 0xff, 0x5c, 0xe8, 0xbf, 0xc2,
	0x8d, 0x79, 0x81, 0xab, 0x6c, 0x23, 0x8e, 0x59, 0x88, 0x9f, 0x15, 0x96, 0xc4, 0x6e, 0xa1, 0x47,
	0xe2, 0x98, 0x45, 0xb1, 0x48, 0xd0, 0x98, 0xbb, 0x12, 0x2c, 0x45, 0x82, 0x6c, 0x0a, 0xce, 0x9b,
	0xaa, 0xa0, 0x8c, 0xee, 0xa2, 0x3f, 0x97, 0x5f, 0x63, 0xae, 0x6b, 0x85, 0x66, 0xe6, 0xdf, 0x83,
	0x6b, 0xd4, 0x6b, 0x5e, 0x12, 0x9b, 0x40, 0x3b, 0xe5, 0x25, 0x79, 0xd6, 0xb8, 0x35, 0x73, 0x17,
	0x03, 0xbd, 0x62, 0x02, 0xa1, 0x1a, 0x2d, 0x02, 0xe8, 0x1a, 0x50, 0xb2, 0x25, 0x5c, 0x3d, 0x21,
	0x35, 0xba, 0xb1, 0x51, 0x63, 0xa9, 0x51, 0x78, 0x74, 0xdd, 0x98, 0xc9, 0x27, 0xfa, 0x17, 0x3b,
	0x47, 0xfd, 0x2c, 0x0f, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xb0, 0xf4, 0x8b, 0xd3, 0x3d, 0x02,
	0x00, 0x00,
}
