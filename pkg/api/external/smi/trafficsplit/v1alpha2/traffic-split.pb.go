// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/mesh-projects/api/external/smi/trafficsplit/v1alpha2/traffic-split.proto

package v1alpha2

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	core "github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// TrafficSplit allows users to incrementally direct percentages of traffic
// between various services. It will be used by clients such as ingress
// controllers or service mesh sidecars to split the outgoing traffic to
// different destinations.
type TrafficSplit struct {
	// Metadata contains the object metadata for this resource
	Metadata core.Metadata `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata"`
	// Status indicates the validation status of this resource.
	// Status is read-only by clients, and set by controllers during validation
	Status core.Status `protobuf:"bytes,2,opt,name=status,proto3" json:"status" testdiff:"ignore"`
	// More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#spec-and-status
	// +optional
	Spec                 *TrafficSplitSpec `protobuf:"bytes,3,opt,name=spec,proto3" json:"spec,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *TrafficSplit) Reset()         { *m = TrafficSplit{} }
func (m *TrafficSplit) String() string { return proto.CompactTextString(m) }
func (*TrafficSplit) ProtoMessage()    {}
func (*TrafficSplit) Descriptor() ([]byte, []int) {
	return fileDescriptor_fb32cb495875581b, []int{0}
}
func (m *TrafficSplit) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TrafficSplit.Unmarshal(m, b)
}
func (m *TrafficSplit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TrafficSplit.Marshal(b, m, deterministic)
}
func (m *TrafficSplit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TrafficSplit.Merge(m, src)
}
func (m *TrafficSplit) XXX_Size() int {
	return xxx_messageInfo_TrafficSplit.Size(m)
}
func (m *TrafficSplit) XXX_DiscardUnknown() {
	xxx_messageInfo_TrafficSplit.DiscardUnknown(m)
}

var xxx_messageInfo_TrafficSplit proto.InternalMessageInfo

func (m *TrafficSplit) GetMetadata() core.Metadata {
	if m != nil {
		return m.Metadata
	}
	return core.Metadata{}
}

func (m *TrafficSplit) GetStatus() core.Status {
	if m != nil {
		return m.Status
	}
	return core.Status{}
}

func (m *TrafficSplit) GetSpec() *TrafficSplitSpec {
	if m != nil {
		return m.Spec
	}
	return nil
}

// TrafficSplitSpec is the specification for a TrafficSplit
type TrafficSplitSpec struct {
	Service              string                 `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	Backends             []*TrafficSplitBackend `protobuf:"bytes,2,rep,name=backends,proto3" json:"backends,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *TrafficSplitSpec) Reset()         { *m = TrafficSplitSpec{} }
func (m *TrafficSplitSpec) String() string { return proto.CompactTextString(m) }
func (*TrafficSplitSpec) ProtoMessage()    {}
func (*TrafficSplitSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_fb32cb495875581b, []int{1}
}
func (m *TrafficSplitSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TrafficSplitSpec.Unmarshal(m, b)
}
func (m *TrafficSplitSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TrafficSplitSpec.Marshal(b, m, deterministic)
}
func (m *TrafficSplitSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TrafficSplitSpec.Merge(m, src)
}
func (m *TrafficSplitSpec) XXX_Size() int {
	return xxx_messageInfo_TrafficSplitSpec.Size(m)
}
func (m *TrafficSplitSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_TrafficSplitSpec.DiscardUnknown(m)
}

var xxx_messageInfo_TrafficSplitSpec proto.InternalMessageInfo

func (m *TrafficSplitSpec) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

func (m *TrafficSplitSpec) GetBackends() []*TrafficSplitBackend {
	if m != nil {
		return m.Backends
	}
	return nil
}

// TrafficSplitBackend defines a backend
type TrafficSplitBackend struct {
	Service              string   `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	Weight               int32    `protobuf:"varint,2,opt,name=weight,proto3" json:"weight,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TrafficSplitBackend) Reset()         { *m = TrafficSplitBackend{} }
func (m *TrafficSplitBackend) String() string { return proto.CompactTextString(m) }
func (*TrafficSplitBackend) ProtoMessage()    {}
func (*TrafficSplitBackend) Descriptor() ([]byte, []int) {
	return fileDescriptor_fb32cb495875581b, []int{2}
}
func (m *TrafficSplitBackend) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TrafficSplitBackend.Unmarshal(m, b)
}
func (m *TrafficSplitBackend) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TrafficSplitBackend.Marshal(b, m, deterministic)
}
func (m *TrafficSplitBackend) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TrafficSplitBackend.Merge(m, src)
}
func (m *TrafficSplitBackend) XXX_Size() int {
	return xxx_messageInfo_TrafficSplitBackend.Size(m)
}
func (m *TrafficSplitBackend) XXX_DiscardUnknown() {
	xxx_messageInfo_TrafficSplitBackend.DiscardUnknown(m)
}

var xxx_messageInfo_TrafficSplitBackend proto.InternalMessageInfo

func (m *TrafficSplitBackend) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

func (m *TrafficSplitBackend) GetWeight() int32 {
	if m != nil {
		return m.Weight
	}
	return 0
}

func init() {
	proto.RegisterType((*TrafficSplit)(nil), "smi.trafficsplit.v1alpha2.TrafficSplit")
	proto.RegisterType((*TrafficSplitSpec)(nil), "smi.trafficsplit.v1alpha2.TrafficSplitSpec")
	proto.RegisterType((*TrafficSplitBackend)(nil), "smi.trafficsplit.v1alpha2.TrafficSplitBackend")
}

func init() {
	proto.RegisterFile("github.com/solo-io/mesh-projects/api/external/smi/trafficsplit/v1alpha2/traffic-split.proto", fileDescriptor_fb32cb495875581b)
}

var fileDescriptor_fb32cb495875581b = []byte{
	// 406 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0x4f, 0xcf, 0xd2, 0x30,
	0x1c, 0xc7, 0x9d, 0x4c, 0xc4, 0x82, 0xff, 0x26, 0x21, 0x83, 0x83, 0x98, 0x25, 0x26, 0x24, 0x4a,
	0x1b, 0xf0, 0x62, 0xb8, 0x98, 0xec, 0x42, 0xa2, 0xf1, 0xe0, 0xf0, 0xa4, 0xa7, 0x52, 0xba, 0xad,
	0xb2, 0xd1, 0xa6, 0x2d, 0xc8, 0x99, 0x57, 0xe3, 0x4b, 0xf1, 0x55, 0x70, 0xf0, 0x1d, 0xe0, 0xcd,
	0x9b, 0x59, 0xb7, 0x91, 0x69, 0x78, 0x9e, 0x67, 0xcf, 0x69, 0xfb, 0xfd, 0xfa, 0xfd, 0xfc, 0xfa,
	0xfd, 0x6e, 0x2d, 0xf8, 0x1a, 0x31, 0x1d, 0x6f, 0x97, 0x90, 0xf0, 0x14, 0x29, 0x9e, 0xf0, 0x31,
	0xe3, 0x28, 0xa5, 0x2a, 0x1e, 0x0b, 0xc9, 0xbf, 0x51, 0xa2, 0x15, 0xc2, 0x82, 0x21, 0xba, 0xd7,
	0x54, 0x6e, 0x70, 0x82, 0x54, 0xca, 0x90, 0x96, 0x38, 0x0c, 0x19, 0x51, 0x22, 0x61, 0x1a, 0xed,
	0x26, 0x38, 0x11, 0x31, 0x9e, 0x96, 0xdd, 0xb1, 0x69, 0x43, 0x21, 0xb9, 0xe6, 0x4e, 0x5f, 0xa5,
	0x0c, 0x56, 0xe5, 0xb0, 0x94, 0x0f, 0xba, 0x11, 0x8f, 0xb8, 0x51, 0xa1, 0xec, 0x2d, 0x07, 0x06,
	0x93, 0x0b, 0x6e, 0xcc, 0x73, 0xcd, 0xb4, 0x31, 0xb2, 0x9b, 0xa0, 0x94, 0x6a, 0xbc, 0xc2, 0x1a,
	0x17, 0xc8, 0xeb, 0x1a, 0x88, 0xa4, 0xe1, 0x2d, 0x36, 0x28, 0xeb, 0x02, 0x41, 0x75, 0x10, 0x8d,
	0xf5, 0x56, 0xe5, 0x80, 0xf7, 0xc7, 0x02, 0x9d, 0xcf, 0x79, 0xe8, 0x45, 0x16, 0xda, 0x79, 0x0b,
	0x5a, 0xa5, 0x69, 0xd7, 0x7a, 0x61, 0x8d, 0xda, 0xd3, 0x1e, 0x24, 0x5c, 0x52, 0x98, 0x8d, 0x81,
	0x8c, 0xc3, 0x8f, 0xc5, 0xaa, 0x6f, 0xff, 0x3c, 0x0e, 0xef, 0x04, 0x67, 0xb5, 0x33, 0x07, 0xcd,
	0x7c, 0xb4, 0x7b, 0xd7, 0x70, 0xdd, 0x7f, 0xb9, 0x85, 0x59, 0xf3, 0xfb, 0x19, 0xf5, 0xfb, 0x38,
	0x7c, 0xaa, 0xa9, 0xd2, 0x2b, 0x16, 0x86, 0x33, 0x8f, 0x45, 0x1b, 0x2e, 0xa9, 0x17, 0x14, 0xb8,
	0xf3, 0x0e, 0xd8, 0x4a, 0x50, 0xe2, 0x36, 0xcc, 0x98, 0x57, 0xf0, 0xca, 0x1f, 0x03, 0xab, 0xce,
	0x17, 0x82, 0x92, 0xc0, 0x80, 0xb3, 0x97, 0x87, 0x93, 0xfd, 0x08, 0x74, 0xaa, 0xcc, 0xe1, 0x64,
	0x3f, 0x76, 0x1e, 0x56, 0x3b, 0xca, 0xdb, 0x83, 0x27, 0xff, 0x0f, 0x70, 0x5c, 0x70, 0x5f, 0x51,
	0xb9, 0x63, 0x84, 0x9a, 0xf4, 0x0f, 0x82, 0xb2, 0x74, 0xde, 0x83, 0xd6, 0x12, 0x93, 0x35, 0xdd,
	0xac, 0xb2, 0x80, 0x8d, 0x51, 0x7b, 0x0a, 0x6b, 0x3a, 0xf3, 0x73, 0x2c, 0x38, 0xf3, 0xde, 0x1c,
	0x3c, 0xbb, 0x20, 0xb8, 0x66, 0xf3, 0x1e, 0x68, 0x7e, 0xa7, 0x2c, 0x8a, 0xb5, 0xf9, 0xb6, 0xf7,
	0x82, 0xa2, 0xf2, 0x3f, 0xfd, 0xf8, 0xf5, 0xdc, 0xfa, 0xf2, 0xe1, 0xc6, 0x7b, 0x21, 0xd6, 0x51,
	0xcd, 0xbb, 0xb1, 0x6c, 0x9a, 0x83, 0xf1, 0xe6, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x76, 0x27,
	0xa7, 0x0c, 0x6d, 0x03, 0x00, 0x00,
}

func (this *TrafficSplit) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*TrafficSplit)
	if !ok {
		that2, ok := that.(TrafficSplit)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Metadata.Equal(&that1.Metadata) {
		return false
	}
	if !this.Status.Equal(&that1.Status) {
		return false
	}
	if !this.Spec.Equal(that1.Spec) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *TrafficSplitSpec) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*TrafficSplitSpec)
	if !ok {
		that2, ok := that.(TrafficSplitSpec)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Service != that1.Service {
		return false
	}
	if len(this.Backends) != len(that1.Backends) {
		return false
	}
	for i := range this.Backends {
		if !this.Backends[i].Equal(that1.Backends[i]) {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *TrafficSplitBackend) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*TrafficSplitBackend)
	if !ok {
		that2, ok := that.(TrafficSplitBackend)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Service != that1.Service {
		return false
	}
	if this.Weight != that1.Weight {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}