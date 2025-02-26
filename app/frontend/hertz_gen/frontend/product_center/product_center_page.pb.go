// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.3
// source: product_center_page.proto

package product_center

import (
	_ "gomall_demo/app/frontend/hertz_gen/api"
	common "gomall_demo/app/frontend/hertz_gen/frontend/common"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ProductCenterReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" query:"id"`
}

func (x *ProductCenterReq) Reset() {
	*x = ProductCenterReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_center_page_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductCenterReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductCenterReq) ProtoMessage() {}

func (x *ProductCenterReq) ProtoReflect() protoreflect.Message {
	mi := &file_product_center_page_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductCenterReq.ProtoReflect.Descriptor instead.
func (*ProductCenterReq) Descriptor() ([]byte, []int) {
	return file_product_center_page_proto_rawDescGZIP(), []int{0}
}

func (x *ProductCenterReq) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type SearchproductCentersReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Q    string `protobuf:"bytes,1,opt,name=q,proto3" json:"q,omitempty" query:"q"`
	Next string `protobuf:"bytes,2,opt,name=next,proto3" json:"next,omitempty" query:"next"`
}

func (x *SearchproductCentersReq) Reset() {
	*x = SearchproductCentersReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_center_page_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchproductCentersReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchproductCentersReq) ProtoMessage() {}

func (x *SearchproductCentersReq) ProtoReflect() protoreflect.Message {
	mi := &file_product_center_page_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchproductCentersReq.ProtoReflect.Descriptor instead.
func (*SearchproductCentersReq) Descriptor() ([]byte, []int) {
	return file_product_center_page_proto_rawDescGZIP(), []int{1}
}

func (x *SearchproductCentersReq) GetQ() string {
	if x != nil {
		return x.Q
	}
	return ""
}

func (x *SearchproductCentersReq) GetNext() string {
	if x != nil {
		return x.Next
	}
	return ""
}

type CreateproductCenterReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name            string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty" form:"name" query:"name"`
	Description     string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty" form:"description" query:"description"`
	Picture         string   `protobuf:"bytes,3,opt,name=picture,proto3" json:"picture,omitempty" form:"picture" query:"picture"`
	Price           float32  `protobuf:"fixed32,4,opt,name=price,proto3" json:"price,omitempty" form:"price" query:"price"`
	Categories      []string `protobuf:"bytes,5,rep,name=categories,proto3" json:"categories,omitempty" form:"categories" query:"categories"`
	Token           string   `protobuf:"bytes,6,opt,name=token,proto3" json:"token,omitempty" form:"token" query:"token"`
	Password        string   `protobuf:"bytes,7,opt,name=password,proto3" json:"password,omitempty" form:"password" query:"password"`
	PasswordConfirm string   `protobuf:"bytes,8,opt,name=password_confirm,json=passwordConfirm,proto3" json:"password_confirm,omitempty" form:"password_confirm" query:"password_confirm"`
	UserId          int32    `protobuf:"varint,9,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty" form:"user_id" query:"user_id"`
	Next            string   `protobuf:"bytes,10,opt,name=next,proto3" json:"next,omitempty" query:"next"`
}

func (x *CreateproductCenterReq) Reset() {
	*x = CreateproductCenterReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_center_page_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateproductCenterReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateproductCenterReq) ProtoMessage() {}

func (x *CreateproductCenterReq) ProtoReflect() protoreflect.Message {
	mi := &file_product_center_page_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateproductCenterReq.ProtoReflect.Descriptor instead.
func (*CreateproductCenterReq) Descriptor() ([]byte, []int) {
	return file_product_center_page_proto_rawDescGZIP(), []int{2}
}

func (x *CreateproductCenterReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateproductCenterReq) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateproductCenterReq) GetPicture() string {
	if x != nil {
		return x.Picture
	}
	return ""
}

func (x *CreateproductCenterReq) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *CreateproductCenterReq) GetCategories() []string {
	if x != nil {
		return x.Categories
	}
	return nil
}

func (x *CreateproductCenterReq) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *CreateproductCenterReq) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *CreateproductCenterReq) GetPasswordConfirm() string {
	if x != nil {
		return x.PasswordConfirm
	}
	return ""
}

func (x *CreateproductCenterReq) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *CreateproductCenterReq) GetNext() string {
	if x != nil {
		return x.Next
	}
	return ""
}

type UpdateproductCenterReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          uint32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" form:"id" query:"id"`
	Name        string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty" form:"name" query:"name"`
	Description string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty" form:"description" query:"description"`
	Picture     string   `protobuf:"bytes,4,opt,name=picture,proto3" json:"picture,omitempty" form:"picture" query:"picture"`
	Price       float32  `protobuf:"fixed32,5,opt,name=price,proto3" json:"price,omitempty" form:"price" query:"price"`
	Categories  []string `protobuf:"bytes,6,rep,name=categories,proto3" json:"categories,omitempty" form:"categories" query:"categories"`
	Token       string   `protobuf:"bytes,7,opt,name=token,proto3" json:"token,omitempty" form:"token" query:"token"`
	Next        string   `protobuf:"bytes,8,opt,name=next,proto3" json:"next,omitempty" query:"next"`
}

func (x *UpdateproductCenterReq) Reset() {
	*x = UpdateproductCenterReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_center_page_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateproductCenterReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateproductCenterReq) ProtoMessage() {}

func (x *UpdateproductCenterReq) ProtoReflect() protoreflect.Message {
	mi := &file_product_center_page_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateproductCenterReq.ProtoReflect.Descriptor instead.
func (*UpdateproductCenterReq) Descriptor() ([]byte, []int) {
	return file_product_center_page_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateproductCenterReq) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateproductCenterReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateproductCenterReq) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *UpdateproductCenterReq) GetPicture() string {
	if x != nil {
		return x.Picture
	}
	return ""
}

func (x *UpdateproductCenterReq) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *UpdateproductCenterReq) GetCategories() []string {
	if x != nil {
		return x.Categories
	}
	return nil
}

func (x *UpdateproductCenterReq) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *UpdateproductCenterReq) GetNext() string {
	if x != nil {
		return x.Next
	}
	return ""
}

type DeleteproductCenterReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" form:"id" query:"id"`
	Token string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty" form:"token" query:"token"`
	Next  string `protobuf:"bytes,3,opt,name=next,proto3" json:"next,omitempty" query:"next"`
}

func (x *DeleteproductCenterReq) Reset() {
	*x = DeleteproductCenterReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_center_page_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteproductCenterReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteproductCenterReq) ProtoMessage() {}

func (x *DeleteproductCenterReq) ProtoReflect() protoreflect.Message {
	mi := &file_product_center_page_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteproductCenterReq.ProtoReflect.Descriptor instead.
func (*DeleteproductCenterReq) Descriptor() ([]byte, []int) {
	return file_product_center_page_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteproductCenterReq) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *DeleteproductCenterReq) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *DeleteproductCenterReq) GetNext() string {
	if x != nil {
		return x.Next
	}
	return ""
}

var File_product_center_page_proto protoreflect.FileDescriptor

var file_product_center_page_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72,
	0x5f, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x66, 0x72, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x63, 0x65,
	0x6e, 0x74, 0x65, 0x72, 0x1a, 0x09, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x15, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2b, 0x0a, 0x11, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x06, 0xb2, 0xbb, 0x18, 0x02, 0x69, 0x64, 0x52,
	0x02, 0x69, 0x64, 0x22, 0x4d, 0x0a, 0x18, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x12,
	0x13, 0x0a, 0x01, 0x71, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x05, 0xb2, 0xbb, 0x18, 0x01,
	0x71, 0x52, 0x01, 0x71, 0x12, 0x1c, 0x0a, 0x04, 0x6e, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x08, 0xb2, 0xbb, 0x18, 0x04, 0x6e, 0x65, 0x78, 0x74, 0x52, 0x04, 0x6e, 0x65,
	0x78, 0x74, 0x22, 0xb3, 0x02, 0x0a, 0x17, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x70,
	0x72, 0x69, 0x63, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69,
	0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x69, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x29, 0x0a, 0x10, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72,
	0x6d, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x04, 0x6e, 0x65,
	0x78, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xb2, 0xbb, 0x18, 0x04, 0x6e, 0x65,
	0x78, 0x74, 0x52, 0x04, 0x6e, 0x65, 0x78, 0x74, 0x22, 0xe3, 0x01, 0x0a, 0x17, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x69,
	0x63, 0x74, 0x75, 0x72, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x69, 0x63,
	0x74, 0x75, 0x72, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a,
	0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x12, 0x1c, 0x0a, 0x04, 0x6e, 0x65, 0x78, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08,
	0xb2, 0xbb, 0x18, 0x04, 0x6e, 0x65, 0x78, 0x74, 0x52, 0x04, 0x6e, 0x65, 0x78, 0x74, 0x22, 0x5d,
	0x0a, 0x17, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f,
	0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12,
	0x1c, 0x0a, 0x04, 0x6e, 0x65, 0x78, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xb2,
	0xbb, 0x18, 0x04, 0x6e, 0x65, 0x78, 0x74, 0x52, 0x04, 0x6e, 0x65, 0x78, 0x74, 0x32, 0xde, 0x04,
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x57, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x16, 0x2e, 0x66,
	0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x12, 0xca, 0xc1,
	0x18, 0x0e, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72,
	0x12, 0x7c, 0x0a, 0x15, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x12, 0x31, 0x2e, 0x66, 0x72, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x63, 0x65, 0x6e,
	0x74, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x66,
	0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x22, 0x18, 0xca, 0xc1, 0x18, 0x14, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x74,
	0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f,
	0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x30, 0x2e, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x64, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x63,
	0x65, 0x6e, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x66, 0x72, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x12, 0xd2, 0xc1, 0x18, 0x0e, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x63, 0x65,
	0x6e, 0x74, 0x65, 0x72, 0x12, 0x7b, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x30, 0x2e, 0x66,
	0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f,
	0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x16,
	0x2e, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x19, 0xd2, 0xc1, 0x18, 0x15, 0x2f, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x12, 0x7b, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x30, 0x2e, 0x66, 0x72, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x63, 0x65, 0x6e,
	0x74, 0x65, 0x72, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x66, 0x72,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x22, 0x19, 0xd2, 0xc1, 0x18, 0x15, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x3c,
	0x5a, 0x3a, 0x67, 0x6f, 0x6d, 0x61, 0x6c, 0x6c, 0x5f, 0x64, 0x65, 0x6d, 0x6f, 0x2f, 0x61, 0x70,
	0x70, 0x2f, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x2f, 0x68, 0x65, 0x72, 0x74, 0x7a,
	0x5f, 0x67, 0x65, 0x6e, 0x2f, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x2f, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_product_center_page_proto_rawDescOnce sync.Once
	file_product_center_page_proto_rawDescData = file_product_center_page_proto_rawDesc
)

func file_product_center_page_proto_rawDescGZIP() []byte {
	file_product_center_page_proto_rawDescOnce.Do(func() {
		file_product_center_page_proto_rawDescData = protoimpl.X.CompressGZIP(file_product_center_page_proto_rawDescData)
	})
	return file_product_center_page_proto_rawDescData
}

var file_product_center_page_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_product_center_page_proto_goTypes = []interface{}{
	(*ProductCenterReq)(nil),        // 0: frontend.product_center.product_centerReq
	(*SearchproductCentersReq)(nil), // 1: frontend.product_center.Searchproduct_centersReq
	(*CreateproductCenterReq)(nil),  // 2: frontend.product_center.Createproduct_centerReq
	(*UpdateproductCenterReq)(nil),  // 3: frontend.product_center.Updateproduct_centerReq
	(*DeleteproductCenterReq)(nil),  // 4: frontend.product_center.Deleteproduct_centerReq
	(*common.Empty)(nil),            // 5: frontend.common.Empty
}
var file_product_center_page_proto_depIdxs = []int32{
	5, // 0: frontend.product_center.product_centerService.Getproduct_center:input_type -> frontend.common.Empty
	1, // 1: frontend.product_center.product_centerService.Searchproduct_centers:input_type -> frontend.product_center.Searchproduct_centersReq
	2, // 2: frontend.product_center.product_centerService.Createproduct_center:input_type -> frontend.product_center.Createproduct_centerReq
	3, // 3: frontend.product_center.product_centerService.Updateproduct_center:input_type -> frontend.product_center.Updateproduct_centerReq
	4, // 4: frontend.product_center.product_centerService.Deleteproduct_center:input_type -> frontend.product_center.Deleteproduct_centerReq
	5, // 5: frontend.product_center.product_centerService.Getproduct_center:output_type -> frontend.common.Empty
	5, // 6: frontend.product_center.product_centerService.Searchproduct_centers:output_type -> frontend.common.Empty
	5, // 7: frontend.product_center.product_centerService.Createproduct_center:output_type -> frontend.common.Empty
	5, // 8: frontend.product_center.product_centerService.Updateproduct_center:output_type -> frontend.common.Empty
	5, // 9: frontend.product_center.product_centerService.Deleteproduct_center:output_type -> frontend.common.Empty
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_product_center_page_proto_init() }
func file_product_center_page_proto_init() {
	if File_product_center_page_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_product_center_page_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductCenterReq); i {
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
		file_product_center_page_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchproductCentersReq); i {
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
		file_product_center_page_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateproductCenterReq); i {
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
		file_product_center_page_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateproductCenterReq); i {
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
		file_product_center_page_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteproductCenterReq); i {
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
			RawDescriptor: file_product_center_page_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_product_center_page_proto_goTypes,
		DependencyIndexes: file_product_center_page_proto_depIdxs,
		MessageInfos:      file_product_center_page_proto_msgTypes,
	}.Build()
	File_product_center_page_proto = out.File
	file_product_center_page_proto_rawDesc = nil
	file_product_center_page_proto_goTypes = nil
	file_product_center_page_proto_depIdxs = nil
}
