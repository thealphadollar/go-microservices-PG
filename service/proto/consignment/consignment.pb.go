// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/consignment/consignment.proto

package consignment

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type GetRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5e5ab05dfa973d5, []int{0}
}

func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequest.Unmarshal(m, b)
}
func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
}
func (m *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(m, src)
}
func (m *GetRequest) XXX_Size() int {
	return xxx_messageInfo_GetRequest.Size(m)
}
func (m *GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequest proto.InternalMessageInfo

type Consignment struct {
	Id                   string       `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Description          string       `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Weight               int32        `protobuf:"varint,3,opt,name=weight,proto3" json:"weight,omitempty"`
	Containers           []*Container `protobuf:"bytes,4,rep,name=containers,proto3" json:"containers,omitempty"`
	VesselId             string       `protobuf:"bytes,5,opt,name=vessel_id,json=vesselId,proto3" json:"vessel_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Consignment) Reset()         { *m = Consignment{} }
func (m *Consignment) String() string { return proto.CompactTextString(m) }
func (*Consignment) ProtoMessage()    {}
func (*Consignment) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5e5ab05dfa973d5, []int{1}
}

func (m *Consignment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Consignment.Unmarshal(m, b)
}
func (m *Consignment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Consignment.Marshal(b, m, deterministic)
}
func (m *Consignment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Consignment.Merge(m, src)
}
func (m *Consignment) XXX_Size() int {
	return xxx_messageInfo_Consignment.Size(m)
}
func (m *Consignment) XXX_DiscardUnknown() {
	xxx_messageInfo_Consignment.DiscardUnknown(m)
}

var xxx_messageInfo_Consignment proto.InternalMessageInfo

func (m *Consignment) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Consignment) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Consignment) GetWeight() int32 {
	if m != nil {
		return m.Weight
	}
	return 0
}

func (m *Consignment) GetContainers() []*Container {
	if m != nil {
		return m.Containers
	}
	return nil
}

func (m *Consignment) GetVesselId() string {
	if m != nil {
		return m.VesselId
	}
	return ""
}

type Container struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CustomerId           string   `protobuf:"bytes,2,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	Origin               string   `protobuf:"bytes,3,opt,name=origin,proto3" json:"origin,omitempty"`
	UserId               string   `protobuf:"bytes,4,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Container) Reset()         { *m = Container{} }
func (m *Container) String() string { return proto.CompactTextString(m) }
func (*Container) ProtoMessage()    {}
func (*Container) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5e5ab05dfa973d5, []int{2}
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

func (m *Container) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
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

type Response struct {
	Created              bool           `protobuf:"varint,1,opt,name=created,proto3" json:"created,omitempty"`
	Consignment          *Consignment   `protobuf:"bytes,2,opt,name=consignment,proto3" json:"consignment,omitempty"`
	Consignments         []*Consignment `protobuf:"bytes,3,rep,name=consignments,proto3" json:"consignments,omitempty"`
	TotalConsignments    int32          `protobuf:"varint,4,opt,name=total_consignments,json=totalConsignments,proto3" json:"total_consignments,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5e5ab05dfa973d5, []int{3}
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

func (m *Response) GetCreated() bool {
	if m != nil {
		return m.Created
	}
	return false
}

func (m *Response) GetConsignment() *Consignment {
	if m != nil {
		return m.Consignment
	}
	return nil
}

func (m *Response) GetConsignments() []*Consignment {
	if m != nil {
		return m.Consignments
	}
	return nil
}

func (m *Response) GetTotalConsignments() int32 {
	if m != nil {
		return m.TotalConsignments
	}
	return 0
}

func init() {
	proto.RegisterType((*GetRequest)(nil), "consignment.GetRequest")
	proto.RegisterType((*Consignment)(nil), "consignment.Consignment")
	proto.RegisterType((*Container)(nil), "consignment.Container")
	proto.RegisterType((*Response)(nil), "consignment.Response")
}

func init() {
	proto.RegisterFile("proto/consignment/consignment.proto", fileDescriptor_e5e5ab05dfa973d5)
}

var fileDescriptor_e5e5ab05dfa973d5 = []byte{
	// 359 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x52, 0xcf, 0x4e, 0xfa, 0x40,
	0x10, 0xfe, 0x95, 0xff, 0x9d, 0x12, 0x7e, 0x61, 0x12, 0x61, 0xa3, 0x07, 0x9b, 0x7a, 0xe1, 0x22,
	0x26, 0x98, 0x78, 0x30, 0x9e, 0x24, 0x91, 0x70, 0x5d, 0x1e, 0x80, 0x60, 0x3b, 0x29, 0x9b, 0xc0,
	0x6e, 0xed, 0x2e, 0xf8, 0x34, 0xbe, 0x81, 0xaf, 0xe2, 0x3b, 0x19, 0xb6, 0x54, 0xb6, 0x1a, 0x6e,
	0xfb, 0xcd, 0x37, 0xdf, 0xcc, 0x37, 0xb3, 0x03, 0x37, 0x59, 0xae, 0x8c, 0xba, 0x8b, 0x95, 0xd4,
	0x22, 0x95, 0x5b, 0x92, 0xc6, 0x7d, 0x8f, 0x2d, 0x8b, 0x81, 0x13, 0x8a, 0xba, 0x00, 0x33, 0x32,
	0x9c, 0xde, 0x76, 0xa4, 0x4d, 0xf4, 0xe9, 0x41, 0x30, 0x3d, 0xb1, 0xd8, 0x83, 0x9a, 0x48, 0x98,
	0x17, 0x7a, 0x23, 0x9f, 0xd7, 0x44, 0x82, 0x21, 0x04, 0x09, 0xe9, 0x38, 0x17, 0x99, 0x11, 0x4a,
	0xb2, 0x9a, 0x25, 0xdc, 0x10, 0x0e, 0xa0, 0xf5, 0x4e, 0x22, 0x5d, 0x1b, 0x56, 0x0f, 0xbd, 0x51,
	0x93, 0x1f, 0x11, 0x3e, 0x00, 0xc4, 0x4a, 0x9a, 0x95, 0x90, 0x94, 0x6b, 0xd6, 0x08, 0xeb, 0xa3,
	0x60, 0x32, 0x18, 0xbb, 0xe6, 0xa6, 0x25, 0xcd, 0x9d, 0x4c, 0xbc, 0x02, 0x7f, 0x4f, 0x5a, 0xd3,
	0x66, 0x29, 0x12, 0xd6, 0xb4, 0xfd, 0x3a, 0x45, 0x60, 0x9e, 0x44, 0x5b, 0xf0, 0x7f, 0x54, 0x7f,
	0xbc, 0x5e, 0x43, 0x10, 0xef, 0xb4, 0x51, 0x5b, 0xca, 0x0f, 0xda, 0xc2, 0x2b, 0x94, 0xa1, 0x79,
	0x72, 0xb0, 0xaa, 0x72, 0x91, 0x0a, 0x69, 0xad, 0xfa, 0xfc, 0x88, 0x70, 0x08, 0xed, 0x9d, 0x2e,
	0x44, 0x8d, 0x82, 0x38, 0xc0, 0x79, 0x12, 0x7d, 0x79, 0xd0, 0xe1, 0xa4, 0x33, 0x25, 0x35, 0x21,
	0x83, 0x76, 0x9c, 0xd3, 0xca, 0x50, 0xd1, 0xb3, 0xc3, 0x4b, 0x88, 0x8f, 0xe0, 0x6e, 0xd8, 0x36,
	0x0e, 0x26, 0xec, 0xf7, 0xac, 0xe5, 0x9b, 0xbb, 0xc9, 0xf8, 0x04, 0x5d, 0x07, 0x6a, 0x56, 0xb7,
	0x8b, 0x3a, 0x2f, 0xae, 0x64, 0xe3, 0x2d, 0xa0, 0x51, 0x66, 0xb5, 0x59, 0x56, 0x6a, 0x34, 0xec,
	0x47, 0xf4, 0x2d, 0xe3, 0x68, 0xf5, 0xe4, 0xc3, 0x83, 0xff, 0x8b, 0xb5, 0xc8, 0x32, 0x21, 0xd3,
	0x05, 0xe5, 0x7b, 0x11, 0x13, 0xbe, 0x40, 0x7f, 0x6a, 0xe7, 0x70, 0xcf, 0xe0, 0x6c, 0xff, 0xcb,
	0x8b, 0x0a, 0x53, 0x2e, 0x27, 0xfa, 0x87, 0xcf, 0xd0, 0x9b, 0x91, 0x71, 0x8b, 0x0c, 0x2b, 0xa9,
	0xa7, 0xa3, 0x3b, 0x5b, 0xe3, 0xb5, 0x65, 0xef, 0xf5, 0xfe, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x80,
	0x42, 0xfb, 0xd8, 0xd6, 0x02, 0x00, 0x00,
}
