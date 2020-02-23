// Code generated by protoc-gen-go. DO NOT EDIT.
// source: components/automate-gateway/api/iam/v2/request/tokens.proto

package request

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
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

type CreateTokenReq struct {
	// Unique, user-specified ID. Cannot be changed.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// User-specified name.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Active state. Defaults to true.
	// If set to false, token will not be authenticated or authorized.
	Active *wrappers.BoolValue `protobuf:"bytes,3,opt,name=active,proto3" json:"active,omitempty"`
	// Unique, optionally user-specified value.
	Value string `protobuf:"bytes,4,opt,name=value,proto3" json:"value,omitempty"`
	// Array of projects the token is in. Empty by default.
	Projects             []string `protobuf:"bytes,5,rep,name=projects,proto3" json:"projects,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateTokenReq) Reset()         { *m = CreateTokenReq{} }
func (m *CreateTokenReq) String() string { return proto.CompactTextString(m) }
func (*CreateTokenReq) ProtoMessage()    {}
func (*CreateTokenReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_928a1e5050fcda6c, []int{0}
}

func (m *CreateTokenReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateTokenReq.Unmarshal(m, b)
}
func (m *CreateTokenReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateTokenReq.Marshal(b, m, deterministic)
}
func (m *CreateTokenReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateTokenReq.Merge(m, src)
}
func (m *CreateTokenReq) XXX_Size() int {
	return xxx_messageInfo_CreateTokenReq.Size(m)
}
func (m *CreateTokenReq) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateTokenReq.DiscardUnknown(m)
}

var xxx_messageInfo_CreateTokenReq proto.InternalMessageInfo

func (m *CreateTokenReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *CreateTokenReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateTokenReq) GetActive() *wrappers.BoolValue {
	if m != nil {
		return m.Active
	}
	return nil
}

func (m *CreateTokenReq) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *CreateTokenReq) GetProjects() []string {
	if m != nil {
		return m.Projects
	}
	return nil
}

type GetTokenReq struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTokenReq) Reset()         { *m = GetTokenReq{} }
func (m *GetTokenReq) String() string { return proto.CompactTextString(m) }
func (*GetTokenReq) ProtoMessage()    {}
func (*GetTokenReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_928a1e5050fcda6c, []int{1}
}

func (m *GetTokenReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTokenReq.Unmarshal(m, b)
}
func (m *GetTokenReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTokenReq.Marshal(b, m, deterministic)
}
func (m *GetTokenReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTokenReq.Merge(m, src)
}
func (m *GetTokenReq) XXX_Size() int {
	return xxx_messageInfo_GetTokenReq.Size(m)
}
func (m *GetTokenReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTokenReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetTokenReq proto.InternalMessageInfo

func (m *GetTokenReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type UpdateTokenReq struct {
	// Unique, user-specified ID. Cannot be changed.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// User-specified name.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Active state. Defaults to true.
	// If set to false, token will not authenticate.
	Active *wrappers.BoolValue `protobuf:"bytes,3,opt,name=active,proto3" json:"active,omitempty"`
	// Array of projects token is in. Empty by default.
	Projects             []string `protobuf:"bytes,5,rep,name=projects,proto3" json:"projects,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateTokenReq) Reset()         { *m = UpdateTokenReq{} }
func (m *UpdateTokenReq) String() string { return proto.CompactTextString(m) }
func (*UpdateTokenReq) ProtoMessage()    {}
func (*UpdateTokenReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_928a1e5050fcda6c, []int{2}
}

func (m *UpdateTokenReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateTokenReq.Unmarshal(m, b)
}
func (m *UpdateTokenReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateTokenReq.Marshal(b, m, deterministic)
}
func (m *UpdateTokenReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateTokenReq.Merge(m, src)
}
func (m *UpdateTokenReq) XXX_Size() int {
	return xxx_messageInfo_UpdateTokenReq.Size(m)
}
func (m *UpdateTokenReq) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateTokenReq.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateTokenReq proto.InternalMessageInfo

func (m *UpdateTokenReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *UpdateTokenReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UpdateTokenReq) GetActive() *wrappers.BoolValue {
	if m != nil {
		return m.Active
	}
	return nil
}

func (m *UpdateTokenReq) GetProjects() []string {
	if m != nil {
		return m.Projects
	}
	return nil
}

type DeleteTokenReq struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteTokenReq) Reset()         { *m = DeleteTokenReq{} }
func (m *DeleteTokenReq) String() string { return proto.CompactTextString(m) }
func (*DeleteTokenReq) ProtoMessage()    {}
func (*DeleteTokenReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_928a1e5050fcda6c, []int{3}
}

func (m *DeleteTokenReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteTokenReq.Unmarshal(m, b)
}
func (m *DeleteTokenReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteTokenReq.Marshal(b, m, deterministic)
}
func (m *DeleteTokenReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteTokenReq.Merge(m, src)
}
func (m *DeleteTokenReq) XXX_Size() int {
	return xxx_messageInfo_DeleteTokenReq.Size(m)
}
func (m *DeleteTokenReq) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteTokenReq.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteTokenReq proto.InternalMessageInfo

func (m *DeleteTokenReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type ListTokensReq struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListTokensReq) Reset()         { *m = ListTokensReq{} }
func (m *ListTokensReq) String() string { return proto.CompactTextString(m) }
func (*ListTokensReq) ProtoMessage()    {}
func (*ListTokensReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_928a1e5050fcda6c, []int{4}
}

func (m *ListTokensReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListTokensReq.Unmarshal(m, b)
}
func (m *ListTokensReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListTokensReq.Marshal(b, m, deterministic)
}
func (m *ListTokensReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListTokensReq.Merge(m, src)
}
func (m *ListTokensReq) XXX_Size() int {
	return xxx_messageInfo_ListTokensReq.Size(m)
}
func (m *ListTokensReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ListTokensReq.DiscardUnknown(m)
}

var xxx_messageInfo_ListTokensReq proto.InternalMessageInfo

type ResetAllTokenProjectsReq struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResetAllTokenProjectsReq) Reset()         { *m = ResetAllTokenProjectsReq{} }
func (m *ResetAllTokenProjectsReq) String() string { return proto.CompactTextString(m) }
func (*ResetAllTokenProjectsReq) ProtoMessage()    {}
func (*ResetAllTokenProjectsReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_928a1e5050fcda6c, []int{5}
}

func (m *ResetAllTokenProjectsReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResetAllTokenProjectsReq.Unmarshal(m, b)
}
func (m *ResetAllTokenProjectsReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResetAllTokenProjectsReq.Marshal(b, m, deterministic)
}
func (m *ResetAllTokenProjectsReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResetAllTokenProjectsReq.Merge(m, src)
}
func (m *ResetAllTokenProjectsReq) XXX_Size() int {
	return xxx_messageInfo_ResetAllTokenProjectsReq.Size(m)
}
func (m *ResetAllTokenProjectsReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ResetAllTokenProjectsReq.DiscardUnknown(m)
}

var xxx_messageInfo_ResetAllTokenProjectsReq proto.InternalMessageInfo

func init() {
	proto.RegisterType((*CreateTokenReq)(nil), "chef.automate.api.iam.v2.CreateTokenReq")
	proto.RegisterType((*GetTokenReq)(nil), "chef.automate.api.iam.v2.GetTokenReq")
	proto.RegisterType((*UpdateTokenReq)(nil), "chef.automate.api.iam.v2.UpdateTokenReq")
	proto.RegisterType((*DeleteTokenReq)(nil), "chef.automate.api.iam.v2.DeleteTokenReq")
	proto.RegisterType((*ListTokensReq)(nil), "chef.automate.api.iam.v2.ListTokensReq")
	proto.RegisterType((*ResetAllTokenProjectsReq)(nil), "chef.automate.api.iam.v2.ResetAllTokenProjectsReq")
}

func init() {
	proto.RegisterFile("components/automate-gateway/api/iam/v2/request/tokens.proto", fileDescriptor_928a1e5050fcda6c)
}

var fileDescriptor_928a1e5050fcda6c = []byte{
	// 365 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x92, 0xcf, 0x6a, 0xdb, 0x40,
	0x10, 0xc6, 0x91, 0xfc, 0xa7, 0xf5, 0x9a, 0xaa, 0x20, 0x7a, 0x10, 0x82, 0x16, 0xa1, 0x93, 0x0f,
	0xf5, 0x2e, 0xa8, 0xb7, 0xf6, 0x64, 0xb7, 0x60, 0x0a, 0x39, 0x04, 0x91, 0xe4, 0x90, 0xdb, 0x58,
	0x1a, 0xcb, 0x9b, 0x48, 0xbb, 0x2b, 0xed, 0x4a, 0x26, 0xf7, 0x3c, 0x41, 0x1e, 0x22, 0x0f, 0x94,
	0x27, 0x0a, 0x5e, 0xd9, 0xbe, 0x39, 0x90, 0x4b, 0x4e, 0x3b, 0x33, 0x7c, 0x33, 0xfb, 0xfb, 0xe0,
	0x23, 0x7f, 0x32, 0x59, 0x29, 0x29, 0x50, 0x18, 0xcd, 0xa0, 0x35, 0xb2, 0x02, 0x83, 0xf3, 0x02,
	0x0c, 0xee, 0xe0, 0x81, 0x81, 0xe2, 0x8c, 0x43, 0xc5, 0xba, 0x84, 0x35, 0x58, 0xb7, 0xa8, 0x0d,
	0x33, 0xf2, 0x1e, 0x85, 0xa6, 0xaa, 0x91, 0x46, 0xfa, 0x41, 0xb6, 0xc5, 0x0d, 0x3d, 0xae, 0x51,
	0x50, 0x9c, 0x72, 0xa8, 0x68, 0x97, 0x84, 0x3f, 0x0a, 0x29, 0x8b, 0x12, 0x99, 0xd5, 0xad, 0xdb,
	0x0d, 0xdb, 0x35, 0xa0, 0x14, 0x36, 0x87, 0xcd, 0xf0, 0xa7, 0x7d, 0xb2, 0x79, 0x81, 0x62, 0xae,
	0x77, 0x50, 0x14, 0xd8, 0x30, 0xa9, 0x0c, 0x97, 0x42, 0x33, 0x10, 0x42, 0x1a, 0xb0, 0x75, 0xaf,
	0x8e, 0x9f, 0x1d, 0xe2, 0xfd, 0x6d, 0x10, 0x0c, 0x5e, 0xed, 0xbf, 0x4f, 0xb1, 0xf6, 0x3d, 0xe2,
	0xf2, 0x3c, 0x70, 0x22, 0x67, 0x36, 0x49, 0x5d, 0x9e, 0xfb, 0x3e, 0x19, 0x0a, 0xa8, 0x30, 0x70,
	0xed, 0xc4, 0xd6, 0x7e, 0x42, 0xc6, 0x90, 0x19, 0xde, 0x61, 0x30, 0x88, 0x9c, 0xd9, 0x34, 0x09,
	0x69, 0x4f, 0x45, 0x8f, 0x54, 0x74, 0x29, 0x65, 0x79, 0x03, 0x65, 0x8b, 0xe9, 0x41, 0xe9, 0x7f,
	0x23, 0xa3, 0x6e, 0x3f, 0x08, 0x86, 0xf6, 0x50, 0xdf, 0xf8, 0x21, 0xf9, 0xac, 0x1a, 0x79, 0x87,
	0x99, 0xd1, 0xc1, 0x28, 0x1a, 0xcc, 0x26, 0xe9, 0xa9, 0xff, 0x4d, 0x9e, 0x16, 0x9f, 0xc8, 0xe8,
	0xc5, 0x71, 0x79, 0x1e, 0x7f, 0x27, 0xd3, 0x15, 0x9a, 0x73, 0x90, 0xf1, 0xa3, 0x43, 0xbc, 0x6b,
	0x95, 0x7f, 0x84, 0x8f, 0x37, 0x88, 0xe3, 0x88, 0x78, 0xff, 0xb0, 0xc4, 0xf3, 0x14, 0xf1, 0x57,
	0xf2, 0xe5, 0x82, 0xeb, 0xde, 0x88, 0x4e, 0xb1, 0x8e, 0x43, 0x12, 0xa4, 0xa8, 0xd1, 0x2c, 0xca,
	0xd2, 0x0e, 0x2f, 0x0f, 0xb7, 0x52, 0xac, 0x97, 0xff, 0x6f, 0x57, 0x05, 0x37, 0xdb, 0x76, 0x4d,
	0x33, 0x59, 0xb1, 0x7d, 0x24, 0x4e, 0x49, 0x62, 0xef, 0x4b, 0xd7, 0x7a, 0x6c, 0x1d, 0xfd, 0x7a,
	0x0d, 0x00, 0x00, 0xff, 0xff, 0x0c, 0xd1, 0xec, 0xa9, 0x96, 0x02, 0x00, 0x00,
}
