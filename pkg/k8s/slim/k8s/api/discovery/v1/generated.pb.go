// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/cilium/cilium/pkg/k8s/slim/k8s/api/discovery/v1/generated.proto

package v1

import (
	fmt "fmt"

	v1 "github.com/cilium/cilium/pkg/k8s/slim/k8s/apis/meta/v1"
	proto "github.com/gogo/protobuf/proto"

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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

func (m *Endpoint) Reset()         { *m = Endpoint{} }
func (m *Endpoint) String() string { return proto.CompactTextString(m) }
func (*Endpoint) ProtoMessage()    {}
func (*Endpoint) Descriptor() ([]byte, []int) {
	return fileDescriptor_824daf76e2aebd1d, []int{0}
}
func (m *Endpoint) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Endpoint.Unmarshal(m, b)
}
func (m *Endpoint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Endpoint.Marshal(b, m, deterministic)
}
func (m *Endpoint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Endpoint.Merge(m, src)
}
func (m *Endpoint) XXX_Size() int {
	return xxx_messageInfo_Endpoint.Size(m)
}
func (m *Endpoint) XXX_DiscardUnknown() {
	xxx_messageInfo_Endpoint.DiscardUnknown(m)
}

var xxx_messageInfo_Endpoint proto.InternalMessageInfo

func (m *Endpoint) GetAddresses() []string {
	if m != nil {
		return m.Addresses
	}
	return nil
}

func (m *Endpoint) GetConditions() *EndpointConditions {
	if m != nil {
		return m.Conditions
	}
	return nil
}

func (m *Endpoint) GetDeprecatedTopology() map[string]string {
	if m != nil {
		return m.DeprecatedTopology
	}
	return nil
}

func (m *Endpoint) GetNodeName() string {
	if m != nil && m.NodeName != nil {
		return *m.NodeName
	}
	return ""
}

func (m *Endpoint) GetZone() string {
	if m != nil && m.Zone != nil {
		return *m.Zone
	}
	return ""
}

func (m *Endpoint) GetHints() *EndpointHints {
	if m != nil {
		return m.Hints
	}
	return nil
}

func (m *EndpointConditions) Reset()         { *m = EndpointConditions{} }
func (m *EndpointConditions) String() string { return proto.CompactTextString(m) }
func (*EndpointConditions) ProtoMessage()    {}
func (*EndpointConditions) Descriptor() ([]byte, []int) {
	return fileDescriptor_824daf76e2aebd1d, []int{1}
}
func (m *EndpointConditions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EndpointConditions.Unmarshal(m, b)
}
func (m *EndpointConditions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EndpointConditions.Marshal(b, m, deterministic)
}
func (m *EndpointConditions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EndpointConditions.Merge(m, src)
}
func (m *EndpointConditions) XXX_Size() int {
	return xxx_messageInfo_EndpointConditions.Size(m)
}
func (m *EndpointConditions) XXX_DiscardUnknown() {
	xxx_messageInfo_EndpointConditions.DiscardUnknown(m)
}

var xxx_messageInfo_EndpointConditions proto.InternalMessageInfo

func (m *EndpointConditions) GetReady() bool {
	if m != nil && m.Ready != nil {
		return *m.Ready
	}
	return false
}

func (m *EndpointConditions) GetServing() bool {
	if m != nil && m.Serving != nil {
		return *m.Serving
	}
	return false
}

func (m *EndpointConditions) GetTerminating() bool {
	if m != nil && m.Terminating != nil {
		return *m.Terminating
	}
	return false
}

func (m *EndpointHints) Reset()         { *m = EndpointHints{} }
func (m *EndpointHints) String() string { return proto.CompactTextString(m) }
func (*EndpointHints) ProtoMessage()    {}
func (*EndpointHints) Descriptor() ([]byte, []int) {
	return fileDescriptor_824daf76e2aebd1d, []int{2}
}
func (m *EndpointHints) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EndpointHints.Unmarshal(m, b)
}
func (m *EndpointHints) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EndpointHints.Marshal(b, m, deterministic)
}
func (m *EndpointHints) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EndpointHints.Merge(m, src)
}
func (m *EndpointHints) XXX_Size() int {
	return xxx_messageInfo_EndpointHints.Size(m)
}
func (m *EndpointHints) XXX_DiscardUnknown() {
	xxx_messageInfo_EndpointHints.DiscardUnknown(m)
}

var xxx_messageInfo_EndpointHints proto.InternalMessageInfo

func (m *EndpointHints) GetForZones() []*ForZone {
	if m != nil {
		return m.ForZones
	}
	return nil
}

func (m *EndpointPort) Reset()         { *m = EndpointPort{} }
func (m *EndpointPort) String() string { return proto.CompactTextString(m) }
func (*EndpointPort) ProtoMessage()    {}
func (*EndpointPort) Descriptor() ([]byte, []int) {
	return fileDescriptor_824daf76e2aebd1d, []int{3}
}
func (m *EndpointPort) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EndpointPort.Unmarshal(m, b)
}
func (m *EndpointPort) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EndpointPort.Marshal(b, m, deterministic)
}
func (m *EndpointPort) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EndpointPort.Merge(m, src)
}
func (m *EndpointPort) XXX_Size() int {
	return xxx_messageInfo_EndpointPort.Size(m)
}
func (m *EndpointPort) XXX_DiscardUnknown() {
	xxx_messageInfo_EndpointPort.DiscardUnknown(m)
}

var xxx_messageInfo_EndpointPort proto.InternalMessageInfo

func (m *EndpointPort) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *EndpointPort) GetProtocol() string {
	if m != nil && m.Protocol != nil {
		return *m.Protocol
	}
	return ""
}

func (m *EndpointPort) GetPort() int32 {
	if m != nil && m.Port != nil {
		return *m.Port
	}
	return 0
}

func (m *EndpointSlice) Reset()         { *m = EndpointSlice{} }
func (m *EndpointSlice) String() string { return proto.CompactTextString(m) }
func (*EndpointSlice) ProtoMessage()    {}
func (*EndpointSlice) Descriptor() ([]byte, []int) {
	return fileDescriptor_824daf76e2aebd1d, []int{4}
}
func (m *EndpointSlice) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EndpointSlice.Unmarshal(m, b)
}
func (m *EndpointSlice) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EndpointSlice.Marshal(b, m, deterministic)
}
func (m *EndpointSlice) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EndpointSlice.Merge(m, src)
}
func (m *EndpointSlice) XXX_Size() int {
	return xxx_messageInfo_EndpointSlice.Size(m)
}
func (m *EndpointSlice) XXX_DiscardUnknown() {
	xxx_messageInfo_EndpointSlice.DiscardUnknown(m)
}

var xxx_messageInfo_EndpointSlice proto.InternalMessageInfo

func (m *EndpointSlice) GetMetadata() *v1.ObjectMeta {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *EndpointSlice) GetAddressType() string {
	if m != nil && m.AddressType != nil {
		return *m.AddressType
	}
	return ""
}

func (m *EndpointSlice) GetEndpoints() []*Endpoint {
	if m != nil {
		return m.Endpoints
	}
	return nil
}

func (m *EndpointSlice) GetPorts() []*EndpointPort {
	if m != nil {
		return m.Ports
	}
	return nil
}

func (m *EndpointSliceList) Reset()         { *m = EndpointSliceList{} }
func (m *EndpointSliceList) String() string { return proto.CompactTextString(m) }
func (*EndpointSliceList) ProtoMessage()    {}
func (*EndpointSliceList) Descriptor() ([]byte, []int) {
	return fileDescriptor_824daf76e2aebd1d, []int{5}
}
func (m *EndpointSliceList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EndpointSliceList.Unmarshal(m, b)
}
func (m *EndpointSliceList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EndpointSliceList.Marshal(b, m, deterministic)
}
func (m *EndpointSliceList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EndpointSliceList.Merge(m, src)
}
func (m *EndpointSliceList) XXX_Size() int {
	return xxx_messageInfo_EndpointSliceList.Size(m)
}
func (m *EndpointSliceList) XXX_DiscardUnknown() {
	xxx_messageInfo_EndpointSliceList.DiscardUnknown(m)
}

var xxx_messageInfo_EndpointSliceList proto.InternalMessageInfo

func (m *EndpointSliceList) GetMetadata() *v1.ListMeta {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *EndpointSliceList) GetItems() []*EndpointSlice {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *ForZone) Reset()         { *m = ForZone{} }
func (m *ForZone) String() string { return proto.CompactTextString(m) }
func (*ForZone) ProtoMessage()    {}
func (*ForZone) Descriptor() ([]byte, []int) {
	return fileDescriptor_824daf76e2aebd1d, []int{6}
}
func (m *ForZone) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ForZone.Unmarshal(m, b)
}
func (m *ForZone) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ForZone.Marshal(b, m, deterministic)
}
func (m *ForZone) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ForZone.Merge(m, src)
}
func (m *ForZone) XXX_Size() int {
	return xxx_messageInfo_ForZone.Size(m)
}
func (m *ForZone) XXX_DiscardUnknown() {
	xxx_messageInfo_ForZone.DiscardUnknown(m)
}

var xxx_messageInfo_ForZone proto.InternalMessageInfo

func (m *ForZone) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*Endpoint)(nil), "github.com.cilium.cilium.pkg.k8s.slim.k8s.api.discovery.v1.Endpoint")
	proto.RegisterMapType((map[string]string)(nil), "github.com.cilium.cilium.pkg.k8s.slim.k8s.api.discovery.v1.Endpoint.DeprecatedTopologyEntry")
	proto.RegisterType((*EndpointConditions)(nil), "github.com.cilium.cilium.pkg.k8s.slim.k8s.api.discovery.v1.EndpointConditions")
	proto.RegisterType((*EndpointHints)(nil), "github.com.cilium.cilium.pkg.k8s.slim.k8s.api.discovery.v1.EndpointHints")
	proto.RegisterType((*EndpointPort)(nil), "github.com.cilium.cilium.pkg.k8s.slim.k8s.api.discovery.v1.EndpointPort")
	proto.RegisterType((*EndpointSlice)(nil), "github.com.cilium.cilium.pkg.k8s.slim.k8s.api.discovery.v1.EndpointSlice")
	proto.RegisterType((*EndpointSliceList)(nil), "github.com.cilium.cilium.pkg.k8s.slim.k8s.api.discovery.v1.EndpointSliceList")
	proto.RegisterType((*ForZone)(nil), "github.com.cilium.cilium.pkg.k8s.slim.k8s.api.discovery.v1.ForZone")
}

func init() {
	proto.RegisterFile("github.com/cilium/cilium/pkg/k8s/slim/k8s/api/discovery/v1/generated.proto", fileDescriptor_824daf76e2aebd1d)
}

var fileDescriptor_824daf76e2aebd1d = []byte{
	// 618 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x95, 0xcf, 0x6a, 0xdb, 0x4e,
	0x10, 0xc7, 0x51, 0x1c, 0xff, 0x22, 0x8f, 0x7f, 0x85, 0x76, 0x29, 0x54, 0x98, 0x16, 0x8c, 0x4e,
	0x3e, 0xad, 0x88, 0x0f, 0x25, 0x84, 0x1e, 0x4a, 0x12, 0x87, 0xb4, 0xb4, 0x69, 0xd9, 0xe6, 0x14,
	0x42, 0xc2, 0x46, 0x9a, 0xd8, 0x5b, 0x4b, 0xbb, 0x62, 0x77, 0x6d, 0x70, 0x9f, 0xa1, 0x4f, 0x57,
	0xe8, 0xf3, 0xb4, 0xec, 0xca, 0x96, 0x93, 0xd8, 0xa5, 0x24, 0xf5, 0x49, 0x33, 0xfb, 0xe7, 0x33,
	0x33, 0xdf, 0xd9, 0x41, 0xf0, 0x7e, 0x28, 0xec, 0x68, 0x72, 0x4d, 0x53, 0x55, 0x24, 0xa9, 0xc8,
	0xc5, 0xa4, 0xfe, 0x94, 0xe3, 0x61, 0x32, 0xde, 0x33, 0x89, 0xc9, 0x45, 0xe1, 0x0d, 0x5e, 0x8a,
	0x24, 0x13, 0x26, 0x55, 0x53, 0xd4, 0xb3, 0x64, 0xba, 0x9b, 0x0c, 0x51, 0xa2, 0xe6, 0x16, 0x33,
	0x5a, 0x6a, 0x65, 0x15, 0xd9, 0x5f, 0xb2, 0x68, 0x05, 0x59, 0x7c, 0xca, 0xf1, 0x90, 0x8e, 0xf7,
	0x0c, 0x75, 0x2c, 0x6f, 0xf0, 0x52, 0xd0, 0x9a, 0x45, 0xa7, 0xbb, 0x9d, 0xc1, 0xc3, 0xf2, 0x48,
	0x95, 0xc6, 0x35, 0x29, 0x74, 0x8e, 0x1f, 0x84, 0x31, 0x49, 0x81, 0x96, 0xaf, 0xe3, 0xb8, 0x13,
	0x54, 0x28, 0x77, 0xa8, 0xe0, 0xe9, 0x48, 0x48, 0x57, 0xb2, 0x43, 0xe8, 0x89, 0xb4, 0xa2, 0xc0,
	0x95, 0x0b, 0xaf, 0xff, 0x76, 0xc1, 0xa4, 0x23, 0x2c, 0xf8, 0xfd, 0x7b, 0xf1, 0xaf, 0x06, 0x84,
	0x03, 0x99, 0x95, 0x4a, 0x48, 0x4b, 0x5e, 0x42, 0x8b, 0x67, 0x99, 0x46, 0x63, 0xd0, 0x44, 0x41,
	0xb7, 0xd1, 0x6b, 0xb1, 0xe5, 0x02, 0x91, 0x00, 0xa9, 0x92, 0x99, 0xb0, 0x42, 0x49, 0x13, 0x6d,
	0x75, 0x83, 0x5e, 0xbb, 0x7f, 0x4a, 0x1f, 0xaf, 0x39, 0x5d, 0xc4, 0x3d, 0xac, 0xa9, 0xec, 0x56,
	0x04, 0xf2, 0x3d, 0x00, 0x92, 0x61, 0xa9, 0x31, 0x75, 0xf9, 0x9e, 0xa9, 0x52, 0xe5, 0x6a, 0x38,
	0x8b, 0x9a, 0xdd, 0x46, 0xaf, 0xdd, 0xbf, 0xd8, 0x44, 0x60, 0x7a, 0xb4, 0x82, 0x1f, 0x48, 0xab,
	0x67, 0x6c, 0x4d, 0x5c, 0xd2, 0x81, 0x50, 0xaa, 0x0c, 0x4f, 0x79, 0x81, 0xd1, 0x7f, 0xdd, 0xa0,
	0xd7, 0x62, 0xb5, 0x4f, 0x08, 0x6c, 0x7f, 0x53, 0x12, 0xa3, 0x1d, 0xbf, 0xee, 0x6d, 0x72, 0x05,
	0xcd, 0x91, 0x90, 0xd6, 0x44, 0xa1, 0x57, 0xea, 0xdd, 0x26, 0x12, 0x3e, 0x71, 0x40, 0x56, 0x71,
	0x3b, 0x03, 0x78, 0xf1, 0x87, 0xfc, 0xc9, 0x53, 0x68, 0x8c, 0x71, 0x16, 0x05, 0x3e, 0x1d, 0x67,
	0x92, 0xe7, 0xd0, 0x9c, 0xf2, 0x7c, 0x82, 0xbe, 0x6f, 0x2d, 0x56, 0x39, 0xfb, 0x5b, 0x7b, 0x41,
	0x7c, 0x03, 0x64, 0xb5, 0x11, 0xee, 0xbc, 0x46, 0x9e, 0x55, 0x8c, 0x90, 0x55, 0x0e, 0x89, 0x60,
	0xc7, 0xa0, 0x9e, 0x0a, 0x39, 0xf4, 0x9c, 0x90, 0x2d, 0x5c, 0xd2, 0x85, 0xb6, 0x45, 0x5d, 0x08,
	0xc9, 0xad, 0xdb, 0x6d, 0xf8, 0xdd, 0xdb, 0x4b, 0x71, 0x09, 0x4f, 0xee, 0x94, 0x41, 0xae, 0x20,
	0xbc, 0x51, 0xfa, 0x5c, 0xc9, 0xf9, 0x63, 0x6b, 0xf7, 0x0f, 0xff, 0x45, 0xa3, 0xe3, 0x8a, 0xc5,
	0x6a, 0x68, 0xcc, 0xe0, 0xff, 0x45, 0xc4, 0xcf, 0x4a, 0x5b, 0xd7, 0x25, 0xe9, 0xba, 0x57, 0xc9,
	0xe2, 0x6d, 0xd7, 0x55, 0x3f, 0x08, 0xa9, 0xca, 0xe7, 0xd2, 0xd4, 0xbe, 0x3b, 0x5f, 0x2a, 0x6d,
	0x7d, 0x31, 0x4d, 0xe6, 0xed, 0xf8, 0xe7, 0xd6, 0xb2, 0x8c, 0x2f, 0xb9, 0x48, 0x91, 0x5c, 0x42,
	0xe8, 0xa6, 0x38, 0xe3, 0x96, 0x7b, 0x72, 0xbb, 0x7f, 0xf0, 0xb0, 0x32, 0x0c, 0x75, 0xf7, 0x5d,
	0x09, 0x9f, 0xae, 0xbf, 0x62, 0x6a, 0x3f, 0xa2, 0xe5, 0xac, 0x66, 0x3a, 0x65, 0xe7, 0x33, 0x78,
	0x36, 0x2b, 0x31, 0xda, 0xf6, 0x49, 0xde, 0x5e, 0x22, 0xd7, 0xd0, 0xc2, 0x79, 0x4a, 0x6e, 0x2e,
	0x9d, 0x92, 0x47, 0x9b, 0x78, 0x6d, 0x6c, 0x89, 0x25, 0x97, 0xd0, 0x74, 0xf5, 0x9b, 0xa8, 0xe1,
	0xf9, 0x27, 0x9b, 0xe0, 0xbb, 0xa6, 0xb0, 0x0a, 0x1b, 0xff, 0x08, 0xe0, 0xd9, 0x1d, 0x5d, 0x3f,
	0x08, 0x63, 0xc9, 0xc5, 0x8a, 0xb6, 0x6f, 0x1f, 0xab, 0xad, 0xe3, 0xdd, 0x53, 0xf6, 0x0a, 0x9a,
	0xc2, 0x62, 0xb1, 0xd0, 0x6c, 0x23, 0x13, 0xea, 0x73, 0x67, 0x15, 0x37, 0x7e, 0x05, 0x3b, 0xf3,
	0x57, 0xb9, 0xee, 0xed, 0x1d, 0xbc, 0x39, 0xdf, 0x7f, 0xfc, 0xdf, 0xef, 0x77, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xe2, 0xa2, 0xd5, 0x27, 0x3a, 0x07, 0x00, 0x00,
}
