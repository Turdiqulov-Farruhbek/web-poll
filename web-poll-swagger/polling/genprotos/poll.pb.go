// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.12.4
// source: protos/poll.proto

package genprotos

import (
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

type PollCreateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title     *string     `protobuf:"bytes,1,opt,name=title,proto3,oneof" json:"title,omitempty"`
	Subtitle  *string     `protobuf:"bytes,2,opt,name=subtitle,proto3,oneof" json:"subtitle,omitempty"`
	Options   []*Option   `protobuf:"bytes,3,rep,name=options,proto3" json:"options,omitempty"`
	Feedbacks []*Feedback `protobuf:"bytes,4,rep,name=feedbacks,proto3" json:"feedbacks,omitempty"`
}

func (x *PollCreateReq) Reset() {
	*x = PollCreateReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_poll_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PollCreateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PollCreateReq) ProtoMessage() {}

func (x *PollCreateReq) ProtoReflect() protoreflect.Message {
	mi := &file_protos_poll_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PollCreateReq.ProtoReflect.Descriptor instead.
func (*PollCreateReq) Descriptor() ([]byte, []int) {
	return file_protos_poll_proto_rawDescGZIP(), []int{0}
}

func (x *PollCreateReq) GetTitle() string {
	if x != nil && x.Title != nil {
		return *x.Title
	}
	return ""
}

func (x *PollCreateReq) GetSubtitle() string {
	if x != nil && x.Subtitle != nil {
		return *x.Subtitle
	}
	return ""
}

func (x *PollCreateReq) GetOptions() []*Option {
	if x != nil {
		return x.Options
	}
	return nil
}

func (x *PollCreateReq) GetFeedbacks() []*Feedback {
	if x != nil {
		return x.Feedbacks
	}
	return nil
}

type Feedback struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	From *int32  `protobuf:"varint,1,opt,name=from,proto3,oneof" json:"from,omitempty"`
	To   *int32  `protobuf:"varint,2,opt,name=to,proto3,oneof" json:"to,omitempty"`
	Text *string `protobuf:"bytes,3,opt,name=text,proto3,oneof" json:"text,omitempty"`
}

func (x *Feedback) Reset() {
	*x = Feedback{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_poll_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Feedback) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Feedback) ProtoMessage() {}

func (x *Feedback) ProtoReflect() protoreflect.Message {
	mi := &file_protos_poll_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Feedback.ProtoReflect.Descriptor instead.
func (*Feedback) Descriptor() ([]byte, []int) {
	return file_protos_poll_proto_rawDescGZIP(), []int{1}
}

func (x *Feedback) GetFrom() int32 {
	if x != nil && x.From != nil {
		return *x.From
	}
	return 0
}

func (x *Feedback) GetTo() int32 {
	if x != nil && x.To != nil {
		return *x.To
	}
	return 0
}

func (x *Feedback) GetText() string {
	if x != nil && x.Text != nil {
		return *x.Text
	}
	return ""
}

type Option struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Variant *string `protobuf:"bytes,1,opt,name=variant,proto3,oneof" json:"variant,omitempty"`
	Ball    *int32  `protobuf:"varint,2,opt,name=ball,proto3,oneof" json:"ball,omitempty"`
}

func (x *Option) Reset() {
	*x = Option{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_poll_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Option) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Option) ProtoMessage() {}

func (x *Option) ProtoReflect() protoreflect.Message {
	mi := &file_protos_poll_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Option.ProtoReflect.Descriptor instead.
func (*Option) Descriptor() ([]byte, []int) {
	return file_protos_poll_proto_rawDescGZIP(), []int{2}
}

func (x *Option) GetVariant() string {
	if x != nil && x.Variant != nil {
		return *x.Variant
	}
	return ""
}

func (x *Option) GetBall() int32 {
	if x != nil && x.Ball != nil {
		return *x.Ball
	}
	return 0
}

type PollUpdateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        *string     `protobuf:"bytes,1,opt,name=id,proto3,oneof" json:"id,omitempty"`
	Title     *string     `protobuf:"bytes,2,opt,name=title,proto3,oneof" json:"title,omitempty"`
	Subtitle  *string     `protobuf:"bytes,3,opt,name=subtitle,proto3,oneof" json:"subtitle,omitempty"`
	Options   []*Option   `protobuf:"bytes,4,rep,name=options,proto3" json:"options,omitempty"`
	Feedbacks []*Feedback `protobuf:"bytes,5,rep,name=feedbacks,proto3" json:"feedbacks,omitempty"`
}

func (x *PollUpdateReq) Reset() {
	*x = PollUpdateReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_poll_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PollUpdateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PollUpdateReq) ProtoMessage() {}

func (x *PollUpdateReq) ProtoReflect() protoreflect.Message {
	mi := &file_protos_poll_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PollUpdateReq.ProtoReflect.Descriptor instead.
func (*PollUpdateReq) Descriptor() ([]byte, []int) {
	return file_protos_poll_proto_rawDescGZIP(), []int{3}
}

func (x *PollUpdateReq) GetId() string {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return ""
}

func (x *PollUpdateReq) GetTitle() string {
	if x != nil && x.Title != nil {
		return *x.Title
	}
	return ""
}

func (x *PollUpdateReq) GetSubtitle() string {
	if x != nil && x.Subtitle != nil {
		return *x.Subtitle
	}
	return ""
}

func (x *PollUpdateReq) GetOptions() []*Option {
	if x != nil {
		return x.Options
	}
	return nil
}

func (x *PollUpdateReq) GetFeedbacks() []*Feedback {
	if x != nil {
		return x.Feedbacks
	}
	return nil
}

type PollUpdateReqSwag struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title     *string     `protobuf:"bytes,1,opt,name=title,proto3,oneof" json:"title,omitempty"`
	Subtitle  *string     `protobuf:"bytes,2,opt,name=subtitle,proto3,oneof" json:"subtitle,omitempty"`
	Options   []*Option   `protobuf:"bytes,3,rep,name=options,proto3" json:"options,omitempty"`
	Feedbacks []*Feedback `protobuf:"bytes,4,rep,name=feedbacks,proto3" json:"feedbacks,omitempty"`
}

func (x *PollUpdateReqSwag) Reset() {
	*x = PollUpdateReqSwag{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_poll_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PollUpdateReqSwag) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PollUpdateReqSwag) ProtoMessage() {}

func (x *PollUpdateReqSwag) ProtoReflect() protoreflect.Message {
	mi := &file_protos_poll_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PollUpdateReqSwag.ProtoReflect.Descriptor instead.
func (*PollUpdateReqSwag) Descriptor() ([]byte, []int) {
	return file_protos_poll_proto_rawDescGZIP(), []int{4}
}

func (x *PollUpdateReqSwag) GetTitle() string {
	if x != nil && x.Title != nil {
		return *x.Title
	}
	return ""
}

func (x *PollUpdateReqSwag) GetSubtitle() string {
	if x != nil && x.Subtitle != nil {
		return *x.Subtitle
	}
	return ""
}

func (x *PollUpdateReqSwag) GetOptions() []*Option {
	if x != nil {
		return x.Options
	}
	return nil
}

func (x *PollUpdateReqSwag) GetFeedbacks() []*Feedback {
	if x != nil {
		return x.Feedbacks
	}
	return nil
}

type PollGetByIDRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       *string     `protobuf:"bytes,1,opt,name=id,proto3,oneof" json:"id,omitempty"`
	PollNum  *int32      `protobuf:"varint,2,opt,name=poll_num,json=pollNum,proto3,oneof" json:"poll_num,omitempty"`
	Title    *string     `protobuf:"bytes,3,opt,name=title,proto3,oneof" json:"title,omitempty"`
	Subtitle *string     `protobuf:"bytes,4,opt,name=subtitle,proto3,oneof" json:"subtitle,omitempty"`
	Options  []*Option   `protobuf:"bytes,5,rep,name=options,proto3" json:"options,omitempty"`
	Feedback []*Feedback `protobuf:"bytes,6,rep,name=feedback,proto3" json:"feedback,omitempty"`
}

func (x *PollGetByIDRes) Reset() {
	*x = PollGetByIDRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_poll_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PollGetByIDRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PollGetByIDRes) ProtoMessage() {}

func (x *PollGetByIDRes) ProtoReflect() protoreflect.Message {
	mi := &file_protos_poll_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PollGetByIDRes.ProtoReflect.Descriptor instead.
func (*PollGetByIDRes) Descriptor() ([]byte, []int) {
	return file_protos_poll_proto_rawDescGZIP(), []int{5}
}

func (x *PollGetByIDRes) GetId() string {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return ""
}

func (x *PollGetByIDRes) GetPollNum() int32 {
	if x != nil && x.PollNum != nil {
		return *x.PollNum
	}
	return 0
}

func (x *PollGetByIDRes) GetTitle() string {
	if x != nil && x.Title != nil {
		return *x.Title
	}
	return ""
}

func (x *PollGetByIDRes) GetSubtitle() string {
	if x != nil && x.Subtitle != nil {
		return *x.Subtitle
	}
	return ""
}

func (x *PollGetByIDRes) GetOptions() []*Option {
	if x != nil {
		return x.Options
	}
	return nil
}

func (x *PollGetByIDRes) GetFeedback() []*Feedback {
	if x != nil {
		return x.Feedback
	}
	return nil
}

type PollGetAllReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId *string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3,oneof" json:"user_id,omitempty"`
}

func (x *PollGetAllReq) Reset() {
	*x = PollGetAllReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_poll_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PollGetAllReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PollGetAllReq) ProtoMessage() {}

func (x *PollGetAllReq) ProtoReflect() protoreflect.Message {
	mi := &file_protos_poll_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PollGetAllReq.ProtoReflect.Descriptor instead.
func (*PollGetAllReq) Descriptor() ([]byte, []int) {
	return file_protos_poll_proto_rawDescGZIP(), []int{6}
}

func (x *PollGetAllReq) GetUserId() string {
	if x != nil && x.UserId != nil {
		return *x.UserId
	}
	return ""
}

type PollGetAllRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Poll []*PollGetByIDRes `protobuf:"bytes,1,rep,name=poll,proto3" json:"poll,omitempty"`
}

func (x *PollGetAllRes) Reset() {
	*x = PollGetAllRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_poll_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PollGetAllRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PollGetAllRes) ProtoMessage() {}

func (x *PollGetAllRes) ProtoReflect() protoreflect.Message {
	mi := &file_protos_poll_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PollGetAllRes.ProtoReflect.Descriptor instead.
func (*PollGetAllRes) Descriptor() ([]byte, []int) {
	return file_protos_poll_proto_rawDescGZIP(), []int{7}
}

func (x *PollGetAllRes) GetPoll() []*PollGetByIDRes {
	if x != nil {
		return x.Poll
	}
	return nil
}

var File_protos_poll_proto protoreflect.FileDescriptor

var file_protos_poll_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x70, 0x6f, 0x6c, 0x6c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x07, 0x70, 0x6f, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x1a, 0x11, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x76, 0x6f, 0x69, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xbe, 0x01, 0x0a, 0x0d, 0x50, 0x6f, 0x6c, 0x6c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x12, 0x19, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x00, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a, 0x08,
	0x73, 0x75, 0x62, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01,
	0x52, 0x08, 0x73, 0x75, 0x62, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x88, 0x01, 0x01, 0x12, 0x29, 0x0a,
	0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f,
	0x2e, 0x70, 0x6f, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x2f, 0x0a, 0x09, 0x66, 0x65, 0x65, 0x64,
	0x62, 0x61, 0x63, 0x6b, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x6f,
	0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x46, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x52, 0x09,
	0x66, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x73, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x73, 0x75, 0x62, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x22, 0x6a, 0x0a, 0x08, 0x46, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x12, 0x17, 0x0a, 0x04,
	0x66, 0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x04, 0x66, 0x72,
	0x6f, 0x6d, 0x88, 0x01, 0x01, 0x12, 0x13, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x48, 0x01, 0x52, 0x02, 0x74, 0x6f, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x04, 0x74, 0x65,
	0x78, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74,
	0x88, 0x01, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x66, 0x72, 0x6f, 0x6d, 0x42, 0x05, 0x0a, 0x03,
	0x5f, 0x74, 0x6f, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x74, 0x65, 0x78, 0x74, 0x22, 0x55, 0x0a, 0x06,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x07, 0x76, 0x61, 0x72, 0x69, 0x61, 0x6e,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x07, 0x76, 0x61, 0x72, 0x69, 0x61,
	0x6e, 0x74, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x04, 0x62, 0x61, 0x6c, 0x6c, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x48, 0x01, 0x52, 0x04, 0x62, 0x61, 0x6c, 0x6c, 0x88, 0x01, 0x01, 0x42, 0x0a,
	0x0a, 0x08, 0x5f, 0x76, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x62,
	0x61, 0x6c, 0x6c, 0x22, 0xda, 0x01, 0x0a, 0x0d, 0x50, 0x6f, 0x6c, 0x6c, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x12, 0x13, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x00, 0x52, 0x02, 0x69, 0x64, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a, 0x08, 0x73, 0x75, 0x62, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x08, 0x73, 0x75, 0x62, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x88, 0x01, 0x01, 0x12, 0x29, 0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x6f, 0x6c, 0x6c, 0x69, 0x6e,
	0x67, 0x2e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x12, 0x2f, 0x0a, 0x09, 0x66, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x73, 0x18, 0x05,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x6f, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x46,
	0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x52, 0x09, 0x66, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63,
	0x6b, 0x73, 0x42, 0x05, 0x0a, 0x03, 0x5f, 0x69, 0x64, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x73, 0x75, 0x62, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x22, 0xc2, 0x01, 0x0a, 0x11, 0x50, 0x6f, 0x6c, 0x6c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x53, 0x77, 0x61, 0x67, 0x12, 0x19, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x88, 0x01,
	0x01, 0x12, 0x1f, 0x0a, 0x08, 0x73, 0x75, 0x62, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x08, 0x73, 0x75, 0x62, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x88,
	0x01, 0x01, 0x12, 0x29, 0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x6f, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x2f, 0x0a,
	0x09, 0x66, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x11, 0x2e, 0x70, 0x6f, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x46, 0x65, 0x65, 0x64, 0x62,
	0x61, 0x63, 0x6b, 0x52, 0x09, 0x66, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x73, 0x42, 0x08,
	0x0a, 0x06, 0x5f, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x73, 0x75, 0x62,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x22, 0x86, 0x02, 0x0a, 0x0e, 0x50, 0x6f, 0x6c, 0x6c, 0x47, 0x65,
	0x74, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x73, 0x12, 0x13, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x02, 0x69, 0x64, 0x88, 0x01, 0x01, 0x12, 0x1e, 0x0a,
	0x08, 0x70, 0x6f, 0x6c, 0x6c, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x48,
	0x01, 0x52, 0x07, 0x70, 0x6f, 0x6c, 0x6c, 0x4e, 0x75, 0x6d, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a, 0x08, 0x73, 0x75, 0x62, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x03, 0x52, 0x08, 0x73, 0x75,
	0x62, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x88, 0x01, 0x01, 0x12, 0x29, 0x0a, 0x07, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x6f, 0x6c,
	0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x12, 0x2d, 0x0a, 0x08, 0x66, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b,
	0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x6f, 0x6c, 0x6c, 0x69, 0x6e, 0x67,
	0x2e, 0x46, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x52, 0x08, 0x66, 0x65, 0x65, 0x64, 0x62,
	0x61, 0x63, 0x6b, 0x42, 0x05, 0x0a, 0x03, 0x5f, 0x69, 0x64, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x70,
	0x6f, 0x6c, 0x6c, 0x5f, 0x6e, 0x75, 0x6d, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x73, 0x75, 0x62, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x22, 0x39,
	0x0a, 0x0d, 0x50, 0x6f, 0x6c, 0x6c, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x12,
	0x1c, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x00, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x88, 0x01, 0x01, 0x42, 0x0a, 0x0a,
	0x08, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x22, 0x3c, 0x0a, 0x0d, 0x50, 0x6f, 0x6c,
	0x6c, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x12, 0x2b, 0x0a, 0x04, 0x70, 0x6f,
	0x6c, 0x6c, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70, 0x6f, 0x6c, 0x6c, 0x69,
	0x6e, 0x67, 0x2e, 0x50, 0x6f, 0x6c, 0x6c, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65,
	0x73, 0x52, 0x04, 0x70, 0x6f, 0x6c, 0x6c, 0x32, 0x84, 0x02, 0x0a, 0x0b, 0x50, 0x6f, 0x6c, 0x6c,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2f, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x12, 0x16, 0x2e, 0x70, 0x6f, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x50, 0x6f, 0x6c, 0x6c,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x0d, 0x2e, 0x70, 0x6f, 0x6c, 0x6c,
	0x69, 0x6e, 0x67, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x12, 0x2f, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x12, 0x16, 0x2e, 0x70, 0x6f, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x50, 0x6f, 0x6c,
	0x6c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x0d, 0x2e, 0x70, 0x6f, 0x6c,
	0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x12, 0x26, 0x0a, 0x06, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x12, 0x0d, 0x2e, 0x70, 0x6f, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x42, 0x79,
	0x49, 0x44, 0x1a, 0x0d, 0x2e, 0x70, 0x6f, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x56, 0x6f, 0x69,
	0x64, 0x12, 0x31, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x44, 0x12, 0x0d, 0x2e, 0x70,
	0x6f, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x42, 0x79, 0x49, 0x44, 0x1a, 0x17, 0x2e, 0x70, 0x6f,
	0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x50, 0x6f, 0x6c, 0x6c, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49,
	0x44, 0x52, 0x65, 0x73, 0x12, 0x38, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x12, 0x16,
	0x2e, 0x70, 0x6f, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x50, 0x6f, 0x6c, 0x6c, 0x47, 0x65, 0x74,
	0x41, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x70, 0x6f, 0x6c, 0x6c, 0x69, 0x6e, 0x67,
	0x2e, 0x50, 0x6f, 0x6c, 0x6c, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x42, 0x0c,
	0x5a, 0x0a, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_poll_proto_rawDescOnce sync.Once
	file_protos_poll_proto_rawDescData = file_protos_poll_proto_rawDesc
)

func file_protos_poll_proto_rawDescGZIP() []byte {
	file_protos_poll_proto_rawDescOnce.Do(func() {
		file_protos_poll_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_poll_proto_rawDescData)
	})
	return file_protos_poll_proto_rawDescData
}

var file_protos_poll_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_protos_poll_proto_goTypes = []any{
	(*PollCreateReq)(nil),     // 0: polling.PollCreateReq
	(*Feedback)(nil),          // 1: polling.Feedback
	(*Option)(nil),            // 2: polling.Option
	(*PollUpdateReq)(nil),     // 3: polling.PollUpdateReq
	(*PollUpdateReqSwag)(nil), // 4: polling.PollUpdateReqSwag
	(*PollGetByIDRes)(nil),    // 5: polling.PollGetByIDRes
	(*PollGetAllReq)(nil),     // 6: polling.PollGetAllReq
	(*PollGetAllRes)(nil),     // 7: polling.PollGetAllRes
	(*ByID)(nil),              // 8: polling.ByID
	(*Void)(nil),              // 9: polling.Void
}
var file_protos_poll_proto_depIdxs = []int32{
	2,  // 0: polling.PollCreateReq.options:type_name -> polling.Option
	1,  // 1: polling.PollCreateReq.feedbacks:type_name -> polling.Feedback
	2,  // 2: polling.PollUpdateReq.options:type_name -> polling.Option
	1,  // 3: polling.PollUpdateReq.feedbacks:type_name -> polling.Feedback
	2,  // 4: polling.PollUpdateReqSwag.options:type_name -> polling.Option
	1,  // 5: polling.PollUpdateReqSwag.feedbacks:type_name -> polling.Feedback
	2,  // 6: polling.PollGetByIDRes.options:type_name -> polling.Option
	1,  // 7: polling.PollGetByIDRes.feedback:type_name -> polling.Feedback
	5,  // 8: polling.PollGetAllRes.poll:type_name -> polling.PollGetByIDRes
	0,  // 9: polling.PollService.Create:input_type -> polling.PollCreateReq
	3,  // 10: polling.PollService.Update:input_type -> polling.PollUpdateReq
	8,  // 11: polling.PollService.Delete:input_type -> polling.ByID
	8,  // 12: polling.PollService.GetByID:input_type -> polling.ByID
	6,  // 13: polling.PollService.GetAll:input_type -> polling.PollGetAllReq
	9,  // 14: polling.PollService.Create:output_type -> polling.Void
	9,  // 15: polling.PollService.Update:output_type -> polling.Void
	9,  // 16: polling.PollService.Delete:output_type -> polling.Void
	5,  // 17: polling.PollService.GetByID:output_type -> polling.PollGetByIDRes
	7,  // 18: polling.PollService.GetAll:output_type -> polling.PollGetAllRes
	14, // [14:19] is the sub-list for method output_type
	9,  // [9:14] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_protos_poll_proto_init() }
func file_protos_poll_proto_init() {
	if File_protos_poll_proto != nil {
		return
	}
	file_protos_void_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_protos_poll_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*PollCreateReq); i {
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
		file_protos_poll_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*Feedback); i {
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
		file_protos_poll_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*Option); i {
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
		file_protos_poll_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*PollUpdateReq); i {
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
		file_protos_poll_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*PollUpdateReqSwag); i {
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
		file_protos_poll_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*PollGetByIDRes); i {
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
		file_protos_poll_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*PollGetAllReq); i {
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
		file_protos_poll_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*PollGetAllRes); i {
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
	file_protos_poll_proto_msgTypes[0].OneofWrappers = []any{}
	file_protos_poll_proto_msgTypes[1].OneofWrappers = []any{}
	file_protos_poll_proto_msgTypes[2].OneofWrappers = []any{}
	file_protos_poll_proto_msgTypes[3].OneofWrappers = []any{}
	file_protos_poll_proto_msgTypes[4].OneofWrappers = []any{}
	file_protos_poll_proto_msgTypes[5].OneofWrappers = []any{}
	file_protos_poll_proto_msgTypes[6].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protos_poll_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_poll_proto_goTypes,
		DependencyIndexes: file_protos_poll_proto_depIdxs,
		MessageInfos:      file_protos_poll_proto_msgTypes,
	}.Build()
	File_protos_poll_proto = out.File
	file_protos_poll_proto_rawDesc = nil
	file_protos_poll_proto_goTypes = nil
	file_protos_poll_proto_depIdxs = nil
}
