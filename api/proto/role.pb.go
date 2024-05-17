// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.1
// source: role.proto

package alpha_pos_system_company_service

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// PosRole
type PosRole struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoleId    string                 `protobuf:"bytes,1,opt,name=role_id,json=roleId,proto3" json:"role_id,omitempty"`
	RoleName  string                 `protobuf:"bytes,2,opt,name=role_name,json=roleName,proto3" json:"role_name,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	CreatedBy string                 `protobuf:"bytes,4,opt,name=created_by,json=createdBy,proto3" json:"created_by,omitempty"`
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	UpdatedBy string                 `protobuf:"bytes,6,opt,name=updated_by,json=updatedBy,proto3" json:"updated_by,omitempty"`
}

func (x *PosRole) Reset() {
	*x = PosRole{}
	if protoimpl.UnsafeEnabled {
		mi := &file_role_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PosRole) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PosRole) ProtoMessage() {}

func (x *PosRole) ProtoReflect() protoreflect.Message {
	mi := &file_role_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PosRole.ProtoReflect.Descriptor instead.
func (*PosRole) Descriptor() ([]byte, []int) {
	return file_role_proto_rawDescGZIP(), []int{0}
}

func (x *PosRole) GetRoleId() string {
	if x != nil {
		return x.RoleId
	}
	return ""
}

func (x *PosRole) GetRoleName() string {
	if x != nil {
		return x.RoleName
	}
	return ""
}

func (x *PosRole) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *PosRole) GetCreatedBy() string {
	if x != nil {
		return x.CreatedBy
	}
	return ""
}

func (x *PosRole) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *PosRole) GetUpdatedBy() string {
	if x != nil {
		return x.UpdatedBy
	}
	return ""
}

// Request and Response messages
type CreatePosRoleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PosRole    *PosRole    `protobuf:"bytes,1,opt,name=pos_role,json=posRole,proto3" json:"pos_role,omitempty"`
	JwtPayload *JWTPayload `protobuf:"bytes,2,opt,name=jwt_payload,json=jwtPayload,proto3" json:"jwt_payload,omitempty"`
}

func (x *CreatePosRoleRequest) Reset() {
	*x = CreatePosRoleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_role_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePosRoleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePosRoleRequest) ProtoMessage() {}

func (x *CreatePosRoleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_role_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePosRoleRequest.ProtoReflect.Descriptor instead.
func (*CreatePosRoleRequest) Descriptor() ([]byte, []int) {
	return file_role_proto_rawDescGZIP(), []int{1}
}

func (x *CreatePosRoleRequest) GetPosRole() *PosRole {
	if x != nil {
		return x.PosRole
	}
	return nil
}

func (x *CreatePosRoleRequest) GetJwtPayload() *JWTPayload {
	if x != nil {
		return x.JwtPayload
	}
	return nil
}

type CreatePosRoleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PosRole *PosRole `protobuf:"bytes,1,opt,name=pos_role,json=posRole,proto3" json:"pos_role,omitempty"`
}

func (x *CreatePosRoleResponse) Reset() {
	*x = CreatePosRoleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_role_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePosRoleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePosRoleResponse) ProtoMessage() {}

func (x *CreatePosRoleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_role_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePosRoleResponse.ProtoReflect.Descriptor instead.
func (*CreatePosRoleResponse) Descriptor() ([]byte, []int) {
	return file_role_proto_rawDescGZIP(), []int{2}
}

func (x *CreatePosRoleResponse) GetPosRole() *PosRole {
	if x != nil {
		return x.PosRole
	}
	return nil
}

type ReadPosRoleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoleId     string      `protobuf:"bytes,1,opt,name=role_id,json=roleId,proto3" json:"role_id,omitempty"`
	JwtPayload *JWTPayload `protobuf:"bytes,2,opt,name=jwt_payload,json=jwtPayload,proto3" json:"jwt_payload,omitempty"`
}

func (x *ReadPosRoleRequest) Reset() {
	*x = ReadPosRoleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_role_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadPosRoleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadPosRoleRequest) ProtoMessage() {}

func (x *ReadPosRoleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_role_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadPosRoleRequest.ProtoReflect.Descriptor instead.
func (*ReadPosRoleRequest) Descriptor() ([]byte, []int) {
	return file_role_proto_rawDescGZIP(), []int{3}
}

func (x *ReadPosRoleRequest) GetRoleId() string {
	if x != nil {
		return x.RoleId
	}
	return ""
}

func (x *ReadPosRoleRequest) GetJwtPayload() *JWTPayload {
	if x != nil {
		return x.JwtPayload
	}
	return nil
}

type ReadPosRoleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PosRole *PosRole `protobuf:"bytes,1,opt,name=pos_role,json=posRole,proto3" json:"pos_role,omitempty"`
}

func (x *ReadPosRoleResponse) Reset() {
	*x = ReadPosRoleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_role_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadPosRoleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadPosRoleResponse) ProtoMessage() {}

func (x *ReadPosRoleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_role_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadPosRoleResponse.ProtoReflect.Descriptor instead.
func (*ReadPosRoleResponse) Descriptor() ([]byte, []int) {
	return file_role_proto_rawDescGZIP(), []int{4}
}

func (x *ReadPosRoleResponse) GetPosRole() *PosRole {
	if x != nil {
		return x.PosRole
	}
	return nil
}

type UpdatePosRoleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PosRole    *PosRole    `protobuf:"bytes,1,opt,name=pos_role,json=posRole,proto3" json:"pos_role,omitempty"`
	JwtPayload *JWTPayload `protobuf:"bytes,2,opt,name=jwt_payload,json=jwtPayload,proto3" json:"jwt_payload,omitempty"`
}

func (x *UpdatePosRoleRequest) Reset() {
	*x = UpdatePosRoleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_role_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePosRoleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePosRoleRequest) ProtoMessage() {}

func (x *UpdatePosRoleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_role_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePosRoleRequest.ProtoReflect.Descriptor instead.
func (*UpdatePosRoleRequest) Descriptor() ([]byte, []int) {
	return file_role_proto_rawDescGZIP(), []int{5}
}

func (x *UpdatePosRoleRequest) GetPosRole() *PosRole {
	if x != nil {
		return x.PosRole
	}
	return nil
}

func (x *UpdatePosRoleRequest) GetJwtPayload() *JWTPayload {
	if x != nil {
		return x.JwtPayload
	}
	return nil
}

type UpdatePosRoleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PosRole *PosRole `protobuf:"bytes,1,opt,name=pos_role,json=posRole,proto3" json:"pos_role,omitempty"`
}

func (x *UpdatePosRoleResponse) Reset() {
	*x = UpdatePosRoleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_role_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePosRoleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePosRoleResponse) ProtoMessage() {}

func (x *UpdatePosRoleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_role_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePosRoleResponse.ProtoReflect.Descriptor instead.
func (*UpdatePosRoleResponse) Descriptor() ([]byte, []int) {
	return file_role_proto_rawDescGZIP(), []int{6}
}

func (x *UpdatePosRoleResponse) GetPosRole() *PosRole {
	if x != nil {
		return x.PosRole
	}
	return nil
}

type DeletePosRoleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoleId     string      `protobuf:"bytes,1,opt,name=role_id,json=roleId,proto3" json:"role_id,omitempty"`
	JwtPayload *JWTPayload `protobuf:"bytes,2,opt,name=jwt_payload,json=jwtPayload,proto3" json:"jwt_payload,omitempty"`
}

func (x *DeletePosRoleRequest) Reset() {
	*x = DeletePosRoleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_role_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePosRoleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePosRoleRequest) ProtoMessage() {}

func (x *DeletePosRoleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_role_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePosRoleRequest.ProtoReflect.Descriptor instead.
func (*DeletePosRoleRequest) Descriptor() ([]byte, []int) {
	return file_role_proto_rawDescGZIP(), []int{7}
}

func (x *DeletePosRoleRequest) GetRoleId() string {
	if x != nil {
		return x.RoleId
	}
	return ""
}

func (x *DeletePosRoleRequest) GetJwtPayload() *JWTPayload {
	if x != nil {
		return x.JwtPayload
	}
	return nil
}

type DeletePosRoleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *DeletePosRoleResponse) Reset() {
	*x = DeletePosRoleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_role_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePosRoleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePosRoleResponse) ProtoMessage() {}

func (x *DeletePosRoleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_role_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePosRoleResponse.ProtoReflect.Descriptor instead.
func (*DeletePosRoleResponse) Descriptor() ([]byte, []int) {
	return file_role_proto_rawDescGZIP(), []int{8}
}

func (x *DeletePosRoleResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type ReadAllPosRolesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit      int32       `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Page       int32       `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	JwtPayload *JWTPayload `protobuf:"bytes,3,opt,name=jwt_payload,json=jwtPayload,proto3" json:"jwt_payload,omitempty"`
}

func (x *ReadAllPosRolesRequest) Reset() {
	*x = ReadAllPosRolesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_role_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadAllPosRolesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadAllPosRolesRequest) ProtoMessage() {}

func (x *ReadAllPosRolesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_role_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadAllPosRolesRequest.ProtoReflect.Descriptor instead.
func (*ReadAllPosRolesRequest) Descriptor() ([]byte, []int) {
	return file_role_proto_rawDescGZIP(), []int{9}
}

func (x *ReadAllPosRolesRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ReadAllPosRolesRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ReadAllPosRolesRequest) GetJwtPayload() *JWTPayload {
	if x != nil {
		return x.JwtPayload
	}
	return nil
}

type ReadAllPosRolesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PosRoles []*PosRole `protobuf:"bytes,1,rep,name=pos_roles,json=posRoles,proto3" json:"pos_roles,omitempty"`
	Limit    int32      `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Page     int32      `protobuf:"varint,3,opt,name=page,proto3" json:"page,omitempty"`
	MaxPage  int32      `protobuf:"varint,4,opt,name=max_page,json=maxPage,proto3" json:"max_page,omitempty"`
	Count    int64      `protobuf:"varint,5,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *ReadAllPosRolesResponse) Reset() {
	*x = ReadAllPosRolesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_role_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadAllPosRolesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadAllPosRolesResponse) ProtoMessage() {}

func (x *ReadAllPosRolesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_role_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadAllPosRolesResponse.ProtoReflect.Descriptor instead.
func (*ReadAllPosRolesResponse) Descriptor() ([]byte, []int) {
	return file_role_proto_rawDescGZIP(), []int{10}
}

func (x *ReadAllPosRolesResponse) GetPosRoles() []*PosRole {
	if x != nil {
		return x.PosRoles
	}
	return nil
}

func (x *ReadAllPosRolesResponse) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ReadAllPosRolesResponse) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ReadAllPosRolesResponse) GetMaxPage() int32 {
	if x != nil {
		return x.MaxPage
	}
	return 0
}

func (x *ReadAllPosRolesResponse) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

var File_role_proto protoreflect.FileDescriptor

var file_role_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x70, 0x6f,
	0x73, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xf3, 0x01, 0x0a, 0x07, 0x50, 0x6f, 0x73, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x17, 0x0a, 0x07,
	0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72,
	0x6f, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x6f, 0x6c, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x12, 0x39, 0x0a, 0x0a,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x62, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x22, 0x71, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x50, 0x6f, 0x73, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x27,
	0x0a, 0x08, 0x70, 0x6f, 0x73, 0x5f, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0c, 0x2e, 0x70, 0x6f, 0x73, 0x2e, 0x50, 0x6f, 0x73, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x07,
	0x70, 0x6f, 0x73, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x30, 0x0a, 0x0b, 0x6a, 0x77, 0x74, 0x5f, 0x70,
	0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70,
	0x6f, 0x73, 0x2e, 0x4a, 0x57, 0x54, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x0a, 0x6a,
	0x77, 0x74, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x40, 0x0a, 0x15, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x27, 0x0a, 0x08, 0x70, 0x6f, 0x73, 0x5f, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x6f, 0x73, 0x2e, 0x50, 0x6f, 0x73, 0x52, 0x6f,
	0x6c, 0x65, 0x52, 0x07, 0x70, 0x6f, 0x73, 0x52, 0x6f, 0x6c, 0x65, 0x22, 0x5f, 0x0a, 0x12, 0x52,
	0x65, 0x61, 0x64, 0x50, 0x6f, 0x73, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x72, 0x6f, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x30, 0x0a, 0x0b, 0x6a, 0x77,
	0x74, 0x5f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0f, 0x2e, 0x70, 0x6f, 0x73, 0x2e, 0x4a, 0x57, 0x54, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64,
	0x52, 0x0a, 0x6a, 0x77, 0x74, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x3e, 0x0a, 0x13,
	0x52, 0x65, 0x61, 0x64, 0x50, 0x6f, 0x73, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x08, 0x70, 0x6f, 0x73, 0x5f, 0x72, 0x6f, 0x6c, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x6f, 0x73, 0x2e, 0x50, 0x6f, 0x73, 0x52,
	0x6f, 0x6c, 0x65, 0x52, 0x07, 0x70, 0x6f, 0x73, 0x52, 0x6f, 0x6c, 0x65, 0x22, 0x71, 0x0a, 0x14,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x08, 0x70, 0x6f, 0x73, 0x5f, 0x72, 0x6f, 0x6c, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x6f, 0x73, 0x2e, 0x50, 0x6f, 0x73,
	0x52, 0x6f, 0x6c, 0x65, 0x52, 0x07, 0x70, 0x6f, 0x73, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x30, 0x0a,
	0x0b, 0x6a, 0x77, 0x74, 0x5f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x6f, 0x73, 0x2e, 0x4a, 0x57, 0x54, 0x50, 0x61, 0x79, 0x6c,
	0x6f, 0x61, 0x64, 0x52, 0x0a, 0x6a, 0x77, 0x74, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22,
	0x40, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x52, 0x6f, 0x6c, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x08, 0x70, 0x6f, 0x73, 0x5f,
	0x72, 0x6f, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x6f, 0x73,
	0x2e, 0x50, 0x6f, 0x73, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x07, 0x70, 0x6f, 0x73, 0x52, 0x6f, 0x6c,
	0x65, 0x22, 0x61, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x52, 0x6f,
	0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x6f, 0x6c,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x6f, 0x6c, 0x65,
	0x49, 0x64, 0x12, 0x30, 0x0a, 0x0b, 0x6a, 0x77, 0x74, 0x5f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x6f, 0x73, 0x2e, 0x4a, 0x57,
	0x54, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x0a, 0x6a, 0x77, 0x74, 0x50, 0x61, 0x79,
	0x6c, 0x6f, 0x61, 0x64, 0x22, 0x31, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6f,
	0x73, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07,
	0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x74, 0x0a, 0x16, 0x52, 0x65, 0x61, 0x64, 0x41,
	0x6c, 0x6c, 0x50, 0x6f, 0x73, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x30, 0x0a, 0x0b, 0x6a,
	0x77, 0x74, 0x5f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0f, 0x2e, 0x70, 0x6f, 0x73, 0x2e, 0x4a, 0x57, 0x54, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x52, 0x0a, 0x6a, 0x77, 0x74, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x9f, 0x01,
	0x0a, 0x17, 0x52, 0x65, 0x61, 0x64, 0x41, 0x6c, 0x6c, 0x50, 0x6f, 0x73, 0x52, 0x6f, 0x6c, 0x65,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x09, 0x70, 0x6f, 0x73,
	0x5f, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70,
	0x6f, 0x73, 0x2e, 0x50, 0x6f, 0x73, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x52,
	0x6f, 0x6c, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61,
	0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x19,
	0x0a, 0x08, 0x6d, 0x61, 0x78, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x07, 0x6d, 0x61, 0x78, 0x50, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x32,
	0xf8, 0x02, 0x0a, 0x0e, 0x50, 0x6f, 0x73, 0x52, 0x6f, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x46, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x52,
	0x6f, 0x6c, 0x65, 0x12, 0x19, 0x2e, 0x70, 0x6f, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x50, 0x6f, 0x73, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a,
	0x2e, 0x70, 0x6f, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x52, 0x6f,
	0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a, 0x0b, 0x52, 0x65,
	0x61, 0x64, 0x50, 0x6f, 0x73, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x17, 0x2e, 0x70, 0x6f, 0x73, 0x2e,
	0x52, 0x65, 0x61, 0x64, 0x50, 0x6f, 0x73, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x18, 0x2e, 0x70, 0x6f, 0x73, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x50, 0x6f, 0x73,
	0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x46, 0x0a, 0x0d,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x19, 0x2e,
	0x70, 0x6f, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x52, 0x6f, 0x6c,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x70, 0x6f, 0x73, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x46, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6f,
	0x73, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x19, 0x2e, 0x70, 0x6f, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x50, 0x6f, 0x73, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1a, 0x2e, 0x70, 0x6f, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6f, 0x73,
	0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4c, 0x0a, 0x0f,
	0x52, 0x65, 0x61, 0x64, 0x41, 0x6c, 0x6c, 0x50, 0x6f, 0x73, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x12,
	0x1b, 0x2e, 0x70, 0x6f, 0x73, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x41, 0x6c, 0x6c, 0x50, 0x6f, 0x73,
	0x52, 0x6f, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x70,
	0x6f, 0x73, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x41, 0x6c, 0x6c, 0x50, 0x6f, 0x73, 0x52, 0x6f, 0x6c,
	0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x39, 0x5a, 0x37, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x6e, 0x64, 0x72, 0x65, 0x77, 0x61,
	0x6c, 0x69, 0x66, 0x62, 0x2f, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2d, 0x70, 0x6f, 0x73, 0x2d, 0x73,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x2d, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2d, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_role_proto_rawDescOnce sync.Once
	file_role_proto_rawDescData = file_role_proto_rawDesc
)

func file_role_proto_rawDescGZIP() []byte {
	file_role_proto_rawDescOnce.Do(func() {
		file_role_proto_rawDescData = protoimpl.X.CompressGZIP(file_role_proto_rawDescData)
	})
	return file_role_proto_rawDescData
}

var file_role_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_role_proto_goTypes = []interface{}{
	(*PosRole)(nil),                 // 0: pos.PosRole
	(*CreatePosRoleRequest)(nil),    // 1: pos.CreatePosRoleRequest
	(*CreatePosRoleResponse)(nil),   // 2: pos.CreatePosRoleResponse
	(*ReadPosRoleRequest)(nil),      // 3: pos.ReadPosRoleRequest
	(*ReadPosRoleResponse)(nil),     // 4: pos.ReadPosRoleResponse
	(*UpdatePosRoleRequest)(nil),    // 5: pos.UpdatePosRoleRequest
	(*UpdatePosRoleResponse)(nil),   // 6: pos.UpdatePosRoleResponse
	(*DeletePosRoleRequest)(nil),    // 7: pos.DeletePosRoleRequest
	(*DeletePosRoleResponse)(nil),   // 8: pos.DeletePosRoleResponse
	(*ReadAllPosRolesRequest)(nil),  // 9: pos.ReadAllPosRolesRequest
	(*ReadAllPosRolesResponse)(nil), // 10: pos.ReadAllPosRolesResponse
	(*timestamppb.Timestamp)(nil),   // 11: google.protobuf.Timestamp
	(*JWTPayload)(nil),              // 12: pos.JWTPayload
}
var file_role_proto_depIdxs = []int32{
	11, // 0: pos.PosRole.created_at:type_name -> google.protobuf.Timestamp
	11, // 1: pos.PosRole.updated_at:type_name -> google.protobuf.Timestamp
	0,  // 2: pos.CreatePosRoleRequest.pos_role:type_name -> pos.PosRole
	12, // 3: pos.CreatePosRoleRequest.jwt_payload:type_name -> pos.JWTPayload
	0,  // 4: pos.CreatePosRoleResponse.pos_role:type_name -> pos.PosRole
	12, // 5: pos.ReadPosRoleRequest.jwt_payload:type_name -> pos.JWTPayload
	0,  // 6: pos.ReadPosRoleResponse.pos_role:type_name -> pos.PosRole
	0,  // 7: pos.UpdatePosRoleRequest.pos_role:type_name -> pos.PosRole
	12, // 8: pos.UpdatePosRoleRequest.jwt_payload:type_name -> pos.JWTPayload
	0,  // 9: pos.UpdatePosRoleResponse.pos_role:type_name -> pos.PosRole
	12, // 10: pos.DeletePosRoleRequest.jwt_payload:type_name -> pos.JWTPayload
	12, // 11: pos.ReadAllPosRolesRequest.jwt_payload:type_name -> pos.JWTPayload
	0,  // 12: pos.ReadAllPosRolesResponse.pos_roles:type_name -> pos.PosRole
	1,  // 13: pos.PosRoleService.CreatePosRole:input_type -> pos.CreatePosRoleRequest
	3,  // 14: pos.PosRoleService.ReadPosRole:input_type -> pos.ReadPosRoleRequest
	5,  // 15: pos.PosRoleService.UpdatePosRole:input_type -> pos.UpdatePosRoleRequest
	7,  // 16: pos.PosRoleService.DeletePosRole:input_type -> pos.DeletePosRoleRequest
	9,  // 17: pos.PosRoleService.ReadAllPosRoles:input_type -> pos.ReadAllPosRolesRequest
	2,  // 18: pos.PosRoleService.CreatePosRole:output_type -> pos.CreatePosRoleResponse
	4,  // 19: pos.PosRoleService.ReadPosRole:output_type -> pos.ReadPosRoleResponse
	6,  // 20: pos.PosRoleService.UpdatePosRole:output_type -> pos.UpdatePosRoleResponse
	8,  // 21: pos.PosRoleService.DeletePosRole:output_type -> pos.DeletePosRoleResponse
	10, // 22: pos.PosRoleService.ReadAllPosRoles:output_type -> pos.ReadAllPosRolesResponse
	18, // [18:23] is the sub-list for method output_type
	13, // [13:18] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_role_proto_init() }
func file_role_proto_init() {
	if File_role_proto != nil {
		return
	}
	file_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_role_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PosRole); i {
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
		file_role_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePosRoleRequest); i {
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
		file_role_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePosRoleResponse); i {
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
		file_role_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadPosRoleRequest); i {
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
		file_role_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadPosRoleResponse); i {
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
		file_role_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePosRoleRequest); i {
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
		file_role_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePosRoleResponse); i {
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
		file_role_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePosRoleRequest); i {
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
		file_role_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePosRoleResponse); i {
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
		file_role_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadAllPosRolesRequest); i {
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
		file_role_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadAllPosRolesResponse); i {
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
			RawDescriptor: file_role_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_role_proto_goTypes,
		DependencyIndexes: file_role_proto_depIdxs,
		MessageInfos:      file_role_proto_msgTypes,
	}.Build()
	File_role_proto = out.File
	file_role_proto_rawDesc = nil
	file_role_proto_goTypes = nil
	file_role_proto_depIdxs = nil
}
