// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

package users

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type User struct {
	// pmongo.ObjectId id = 1 [(tagger.tags) = "bson:\"_id,omitempty\"" ];
	Id                   string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Username             string               `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Password             string               `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	Email                string               `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	Address              string               `protobuf:"bytes,5,opt,name=address,proto3" json:"address,omitempty"`
	Phone                string               `protobuf:"bytes,6,opt,name=phone,proto3" json:"phone,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,7,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt            *timestamp.Timestamp `protobuf:"bytes,8,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{0}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *User) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *User) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *User) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *User) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *User) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

type FindAllRequest struct {
	Page                 string   `protobuf:"bytes,1,opt,name=page,proto3" json:"page,omitempty"`
	PerPage              string   `protobuf:"bytes,2,opt,name=per_page,json=perPage,proto3" json:"per_page,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindAllRequest) Reset()         { *m = FindAllRequest{} }
func (m *FindAllRequest) String() string { return proto.CompactTextString(m) }
func (*FindAllRequest) ProtoMessage()    {}
func (*FindAllRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{1}
}

func (m *FindAllRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindAllRequest.Unmarshal(m, b)
}
func (m *FindAllRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindAllRequest.Marshal(b, m, deterministic)
}
func (m *FindAllRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindAllRequest.Merge(m, src)
}
func (m *FindAllRequest) XXX_Size() int {
	return xxx_messageInfo_FindAllRequest.Size(m)
}
func (m *FindAllRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FindAllRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FindAllRequest proto.InternalMessageInfo

func (m *FindAllRequest) GetPage() string {
	if m != nil {
		return m.Page
	}
	return ""
}

func (m *FindAllRequest) GetPerPage() string {
	if m != nil {
		return m.PerPage
	}
	return ""
}

type FindAllResponse struct {
	Status               bool     `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Data                 []*User  `protobuf:"bytes,3,rep,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindAllResponse) Reset()         { *m = FindAllResponse{} }
func (m *FindAllResponse) String() string { return proto.CompactTextString(m) }
func (*FindAllResponse) ProtoMessage()    {}
func (*FindAllResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{2}
}

func (m *FindAllResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindAllResponse.Unmarshal(m, b)
}
func (m *FindAllResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindAllResponse.Marshal(b, m, deterministic)
}
func (m *FindAllResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindAllResponse.Merge(m, src)
}
func (m *FindAllResponse) XXX_Size() int {
	return xxx_messageInfo_FindAllResponse.Size(m)
}
func (m *FindAllResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FindAllResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FindAllResponse proto.InternalMessageInfo

func (m *FindAllResponse) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

func (m *FindAllResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *FindAllResponse) GetData() []*User {
	if m != nil {
		return m.Data
	}
	return nil
}

type CreateRequest struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateRequest) Reset()         { *m = CreateRequest{} }
func (m *CreateRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()    {}
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{3}
}

func (m *CreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRequest.Unmarshal(m, b)
}
func (m *CreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRequest.Marshal(b, m, deterministic)
}
func (m *CreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRequest.Merge(m, src)
}
func (m *CreateRequest) XXX_Size() int {
	return xxx_messageInfo_CreateRequest.Size(m)
}
func (m *CreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRequest proto.InternalMessageInfo

func (m *CreateRequest) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type CreateResponse struct {
	Status               bool     `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	User                 *User    `protobuf:"bytes,3,opt,name=user,proto3" json:"user,omitempty"`
	Token                string   `protobuf:"bytes,4,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateResponse) Reset()         { *m = CreateResponse{} }
func (m *CreateResponse) String() string { return proto.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()    {}
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{4}
}

func (m *CreateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateResponse.Unmarshal(m, b)
}
func (m *CreateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateResponse.Marshal(b, m, deterministic)
}
func (m *CreateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateResponse.Merge(m, src)
}
func (m *CreateResponse) XXX_Size() int {
	return xxx_messageInfo_CreateResponse.Size(m)
}
func (m *CreateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateResponse proto.InternalMessageInfo

func (m *CreateResponse) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

func (m *CreateResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *CreateResponse) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *CreateResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type LoginRequest struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginRequest) Reset()         { *m = LoginRequest{} }
func (m *LoginRequest) String() string { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()    {}
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{5}
}

func (m *LoginRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginRequest.Unmarshal(m, b)
}
func (m *LoginRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginRequest.Marshal(b, m, deterministic)
}
func (m *LoginRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginRequest.Merge(m, src)
}
func (m *LoginRequest) XXX_Size() int {
	return xxx_messageInfo_LoginRequest.Size(m)
}
func (m *LoginRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LoginRequest proto.InternalMessageInfo

func (m *LoginRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *LoginRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type GetOneRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Email                string   `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Username             string   `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetOneRequest) Reset()         { *m = GetOneRequest{} }
func (m *GetOneRequest) String() string { return proto.CompactTextString(m) }
func (*GetOneRequest) ProtoMessage()    {}
func (*GetOneRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{6}
}

func (m *GetOneRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetOneRequest.Unmarshal(m, b)
}
func (m *GetOneRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetOneRequest.Marshal(b, m, deterministic)
}
func (m *GetOneRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetOneRequest.Merge(m, src)
}
func (m *GetOneRequest) XXX_Size() int {
	return xxx_messageInfo_GetOneRequest.Size(m)
}
func (m *GetOneRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetOneRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetOneRequest proto.InternalMessageInfo

func (m *GetOneRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *GetOneRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *GetOneRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

type GetOneResponse struct {
	Status               bool     `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Data                 *User    `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetOneResponse) Reset()         { *m = GetOneResponse{} }
func (m *GetOneResponse) String() string { return proto.CompactTextString(m) }
func (*GetOneResponse) ProtoMessage()    {}
func (*GetOneResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{7}
}

func (m *GetOneResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetOneResponse.Unmarshal(m, b)
}
func (m *GetOneResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetOneResponse.Marshal(b, m, deterministic)
}
func (m *GetOneResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetOneResponse.Merge(m, src)
}
func (m *GetOneResponse) XXX_Size() int {
	return xxx_messageInfo_GetOneResponse.Size(m)
}
func (m *GetOneResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetOneResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetOneResponse proto.InternalMessageInfo

func (m *GetOneResponse) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

func (m *GetOneResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *GetOneResponse) GetData() *User {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*User)(nil), "users.User")
	proto.RegisterType((*FindAllRequest)(nil), "users.FindAllRequest")
	proto.RegisterType((*FindAllResponse)(nil), "users.FindAllResponse")
	proto.RegisterType((*CreateRequest)(nil), "users.CreateRequest")
	proto.RegisterType((*CreateResponse)(nil), "users.CreateResponse")
	proto.RegisterType((*LoginRequest)(nil), "users.LoginRequest")
	proto.RegisterType((*GetOneRequest)(nil), "users.GetOneRequest")
	proto.RegisterType((*GetOneResponse)(nil), "users.GetOneResponse")
}

func init() { proto.RegisterFile("user.proto", fileDescriptor_116e343673f7ffaf) }

var fileDescriptor_116e343673f7ffaf = []byte{
	// 476 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x53, 0x4d, 0x6f, 0xd4, 0x30,
	0x10, 0x55, 0xb2, 0xd9, 0xaf, 0x59, 0xba, 0x48, 0xa6, 0xad, 0x42, 0x2e, 0xbb, 0xca, 0x69, 0x4f,
	0x29, 0x5a, 0x24, 0x28, 0x27, 0xa8, 0x90, 0xe0, 0x82, 0x54, 0x88, 0xe0, 0x8c, 0xdc, 0x7a, 0x08,
	0x11, 0x9b, 0xd8, 0xc4, 0x8e, 0x10, 0xff, 0x80, 0x3f, 0xcb, 0x7f, 0x40, 0x9e, 0xd8, 0x61, 0x13,
	0x51, 0x71, 0x80, 0x9b, 0xdf, 0xf8, 0xbd, 0xb1, 0xc7, 0xef, 0x19, 0xa0, 0xd5, 0xd8, 0x64, 0xaa,
	0x91, 0x46, 0xb2, 0xa9, 0x5d, 0xeb, 0x64, 0x53, 0x48, 0x59, 0x1c, 0xf0, 0x82, 0x8a, 0x37, 0xed,
	0xa7, 0x0b, 0x53, 0x56, 0xa8, 0x0d, 0xaf, 0x54, 0xc7, 0x4b, 0x7f, 0x84, 0x10, 0x7d, 0xd0, 0xd8,
	0xb0, 0x35, 0x84, 0xa5, 0x88, 0x83, 0x6d, 0xb0, 0x5b, 0xe6, 0x61, 0x29, 0x58, 0x02, 0x0b, 0xdb,
	0xa2, 0xe6, 0x15, 0xc6, 0x21, 0x55, 0x7b, 0x6c, 0xf7, 0x14, 0xd7, 0xfa, 0x9b, 0x6c, 0x44, 0x3c,
	0xe9, 0xf6, 0x3c, 0x66, 0xa7, 0x30, 0xc5, 0x8a, 0x97, 0x87, 0x38, 0xa2, 0x8d, 0x0e, 0xb0, 0x18,
	0xe6, 0x5c, 0x88, 0x06, 0xb5, 0x8e, 0xa7, 0x54, 0xf7, 0xd0, 0xf2, 0xd5, 0x67, 0x59, 0x63, 0x3c,
	0xeb, 0xf8, 0x04, 0xd8, 0x25, 0x2c, 0x6f, 0x1b, 0xe4, 0x06, 0xc5, 0x95, 0x89, 0xe7, 0xdb, 0x60,
	0xb7, 0xda, 0x27, 0x59, 0x37, 0x4b, 0xe6, 0x67, 0xc9, 0xde, 0xfb, 0x59, 0xf2, 0xdf, 0x64, 0xab,
	0x6c, 0x95, 0x70, 0xca, 0xc5, 0xdf, 0x95, 0x3d, 0x39, 0x7d, 0x0e, 0xeb, 0x57, 0x65, 0x2d, 0xae,
	0x0e, 0x87, 0x1c, 0xbf, 0xb6, 0xa8, 0x0d, 0x63, 0x10, 0x29, 0x5e, 0xa0, 0x7b, 0x15, 0x5a, 0xb3,
	0x87, 0xb0, 0x50, 0xd8, 0x7c, 0xa4, 0x7a, 0xf7, 0x2e, 0x73, 0x85, 0xcd, 0x5b, 0x5e, 0x60, 0x2a,
	0xe0, 0x7e, 0xdf, 0x40, 0x2b, 0x59, 0x6b, 0x64, 0xe7, 0x30, 0xd3, 0x86, 0x9b, 0x56, 0x53, 0x8f,
	0x45, 0xee, 0x90, 0x7d, 0x8f, 0x0a, 0xb5, 0x3e, 0x6a, 0xe2, 0x20, 0xdb, 0x40, 0x24, 0xb8, 0xe1,
	0xf1, 0x64, 0x3b, 0xd9, 0xad, 0xf6, 0xab, 0x8c, 0x7c, 0xcc, 0xac, 0x45, 0x39, 0x6d, 0xa4, 0x8f,
	0xe0, 0xe4, 0x25, 0x4d, 0xeb, 0x6f, 0xb9, 0x81, 0xc8, 0x92, 0xe8, 0x84, 0xb1, 0xc2, 0xae, 0xd3,
	0xef, 0xb0, 0xf6, 0x8a, 0x7f, 0xb9, 0x16, 0x1d, 0x32, 0xb9, 0xe3, 0x10, 0xeb, 0xa3, 0x91, 0x5f,
	0xb0, 0xf6, 0xbe, 0x13, 0x48, 0x5f, 0xc0, 0xbd, 0x37, 0xb2, 0x28, 0x6b, 0x7f, 0xd7, 0x3e, 0x1d,
	0xc1, 0x71, 0x3a, 0x8e, 0xf3, 0x14, 0x0e, 0xf3, 0x94, 0xbe, 0x83, 0x93, 0xd7, 0x68, 0xae, 0xeb,
	0x7e, 0xdc, 0x71, 0x50, 0xfb, 0x96, 0xe1, 0xa8, 0x65, 0x1f, 0xdf, 0xc9, 0x30, 0xbe, 0xe9, 0x2d,
	0xac, 0x7d, 0xcb, 0xff, 0x60, 0x53, 0xf0, 0x47, 0x9b, 0xf6, 0x3f, 0x03, 0x98, 0x5a, 0xa8, 0xd9,
	0x25, 0xcc, 0x5d, 0x2c, 0xd8, 0x99, 0xe3, 0x0d, 0x73, 0x96, 0x9c, 0x8f, 0xcb, 0xee, 0x5a, 0x4f,
	0x48, 0x89, 0xd7, 0x35, 0xb2, 0x53, 0x47, 0x19, 0xbc, 0x45, 0x72, 0x36, 0xaa, 0x3a, 0xdd, 0x33,
	0x80, 0xce, 0x70, 0xfa, 0xd9, 0x5e, 0x3a, 0x48, 0x4d, 0x2f, 0x1d, 0x25, 0xe3, 0x29, 0x2c, 0xc9,
	0x30, 0x52, 0x3e, 0x70, 0x9c, 0x63, 0x0b, 0xef, 0x10, 0xde, 0xcc, 0xe8, 0x73, 0x3d, 0xfe, 0x15,
	0x00, 0x00, 0xff, 0xff, 0x31, 0xca, 0x8b, 0xe8, 0x85, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UsersClient is the client API for Users service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UsersClient interface {
	FindAll(ctx context.Context, in *FindAllRequest, opts ...grpc.CallOption) (*FindAllResponse, error)
	FineOne(ctx context.Context, in *GetOneRequest, opts ...grpc.CallOption) (*GetOneResponse, error)
	CreateUser(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	LoginUser(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*CreateResponse, error)
}

type usersClient struct {
	cc *grpc.ClientConn
}

func NewUsersClient(cc *grpc.ClientConn) UsersClient {
	return &usersClient{cc}
}

func (c *usersClient) FindAll(ctx context.Context, in *FindAllRequest, opts ...grpc.CallOption) (*FindAllResponse, error) {
	out := new(FindAllResponse)
	err := c.cc.Invoke(ctx, "/users.Users/FindAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) FineOne(ctx context.Context, in *GetOneRequest, opts ...grpc.CallOption) (*GetOneResponse, error) {
	out := new(GetOneResponse)
	err := c.cc.Invoke(ctx, "/users.Users/FineOne", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) CreateUser(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, "/users.Users/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) LoginUser(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, "/users.Users/LoginUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UsersServer is the server API for Users service.
type UsersServer interface {
	FindAll(context.Context, *FindAllRequest) (*FindAllResponse, error)
	FineOne(context.Context, *GetOneRequest) (*GetOneResponse, error)
	CreateUser(context.Context, *CreateRequest) (*CreateResponse, error)
	LoginUser(context.Context, *LoginRequest) (*CreateResponse, error)
}

// UnimplementedUsersServer can be embedded to have forward compatible implementations.
type UnimplementedUsersServer struct {
}

func (*UnimplementedUsersServer) FindAll(ctx context.Context, req *FindAllRequest) (*FindAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindAll not implemented")
}
func (*UnimplementedUsersServer) FineOne(ctx context.Context, req *GetOneRequest) (*GetOneResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FineOne not implemented")
}
func (*UnimplementedUsersServer) CreateUser(ctx context.Context, req *CreateRequest) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (*UnimplementedUsersServer) LoginUser(ctx context.Context, req *LoginRequest) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginUser not implemented")
}

func RegisterUsersServer(s *grpc.Server, srv UsersServer) {
	s.RegisterService(&_Users_serviceDesc, srv)
}

func _Users_FindAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).FindAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/users.Users/FindAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).FindAll(ctx, req.(*FindAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_FineOne_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOneRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).FineOne(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/users.Users/FineOne",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).FineOne(ctx, req.(*GetOneRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/users.Users/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).CreateUser(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_LoginUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).LoginUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/users.Users/LoginUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).LoginUser(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Users_serviceDesc = grpc.ServiceDesc{
	ServiceName: "users.Users",
	HandlerType: (*UsersServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FindAll",
			Handler:    _Users_FindAll_Handler,
		},
		{
			MethodName: "FineOne",
			Handler:    _Users_FineOne_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _Users_CreateUser_Handler,
		},
		{
			MethodName: "LoginUser",
			Handler:    _Users_LoginUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}
