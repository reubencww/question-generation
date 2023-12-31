// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.9
// source: hskw/hskw.proto

package hskw

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

type GenerateCaption struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChallengeId int32  `protobuf:"varint,1,opt,name=challenge_id,json=challengeId,proto3" json:"challenge_id,omitempty"`
	ImagePath   string `protobuf:"bytes,2,opt,name=image_path,json=imagePath,proto3" json:"image_path,omitempty"`
}

func (x *GenerateCaption) Reset() {
	*x = GenerateCaption{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hskw_hskw_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateCaption) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateCaption) ProtoMessage() {}

func (x *GenerateCaption) ProtoReflect() protoreflect.Message {
	mi := &file_hskw_hskw_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateCaption.ProtoReflect.Descriptor instead.
func (*GenerateCaption) Descriptor() ([]byte, []int) {
	return file_hskw_hskw_proto_rawDescGZIP(), []int{0}
}

func (x *GenerateCaption) GetChallengeId() int32 {
	if x != nil {
		return x.ChallengeId
	}
	return 0
}

func (x *GenerateCaption) GetImagePath() string {
	if x != nil {
		return x.ImagePath
	}
	return ""
}

type GenerateQuestionAnswer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChallengeId int32  `protobuf:"varint,1,opt,name=challenge_id,json=challengeId,proto3" json:"challenge_id,omitempty"`
	Caption     string `protobuf:"bytes,2,opt,name=caption,proto3" json:"caption,omitempty"`
}

func (x *GenerateQuestionAnswer) Reset() {
	*x = GenerateQuestionAnswer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hskw_hskw_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateQuestionAnswer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateQuestionAnswer) ProtoMessage() {}

func (x *GenerateQuestionAnswer) ProtoReflect() protoreflect.Message {
	mi := &file_hskw_hskw_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateQuestionAnswer.ProtoReflect.Descriptor instead.
func (*GenerateQuestionAnswer) Descriptor() ([]byte, []int) {
	return file_hskw_hskw_proto_rawDescGZIP(), []int{1}
}

func (x *GenerateQuestionAnswer) GetChallengeId() int32 {
	if x != nil {
		return x.ChallengeId
	}
	return 0
}

func (x *GenerateQuestionAnswer) GetCaption() string {
	if x != nil {
		return x.Caption
	}
	return ""
}

type CreatedCaptionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChallengeId int32  `protobuf:"varint,1,opt,name=challenge_id,json=challengeId,proto3" json:"challenge_id,omitempty"`
	Caption     string `protobuf:"bytes,2,opt,name=caption,proto3" json:"caption,omitempty"`
}

func (x *CreatedCaptionRequest) Reset() {
	*x = CreatedCaptionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hskw_hskw_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatedCaptionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatedCaptionRequest) ProtoMessage() {}

func (x *CreatedCaptionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_hskw_hskw_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatedCaptionRequest.ProtoReflect.Descriptor instead.
func (*CreatedCaptionRequest) Descriptor() ([]byte, []int) {
	return file_hskw_hskw_proto_rawDescGZIP(), []int{2}
}

func (x *CreatedCaptionRequest) GetChallengeId() int32 {
	if x != nil {
		return x.ChallengeId
	}
	return 0
}

func (x *CreatedCaptionRequest) GetCaption() string {
	if x != nil {
		return x.Caption
	}
	return ""
}

type CreatedCaptionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreatedCaptionResponse) Reset() {
	*x = CreatedCaptionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hskw_hskw_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatedCaptionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatedCaptionResponse) ProtoMessage() {}

func (x *CreatedCaptionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_hskw_hskw_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatedCaptionResponse.ProtoReflect.Descriptor instead.
func (*CreatedCaptionResponse) Descriptor() ([]byte, []int) {
	return file_hskw_hskw_proto_rawDescGZIP(), []int{3}
}

type QNA struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Question string `protobuf:"bytes,1,opt,name=question,proto3" json:"question,omitempty"`
	Answer   string `protobuf:"bytes,2,opt,name=answer,proto3" json:"answer,omitempty"`
}

func (x *QNA) Reset() {
	*x = QNA{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hskw_hskw_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QNA) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QNA) ProtoMessage() {}

func (x *QNA) ProtoReflect() protoreflect.Message {
	mi := &file_hskw_hskw_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QNA.ProtoReflect.Descriptor instead.
func (*QNA) Descriptor() ([]byte, []int) {
	return file_hskw_hskw_proto_rawDescGZIP(), []int{4}
}

func (x *QNA) GetQuestion() string {
	if x != nil {
		return x.Question
	}
	return ""
}

func (x *QNA) GetAnswer() string {
	if x != nil {
		return x.Answer
	}
	return ""
}

type CreatedQuestionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChallengeId int32  `protobuf:"varint,1,opt,name=challenge_id,json=challengeId,proto3" json:"challenge_id,omitempty"`
	Qnas        []*QNA `protobuf:"bytes,2,rep,name=qnas,proto3" json:"qnas,omitempty"`
}

func (x *CreatedQuestionRequest) Reset() {
	*x = CreatedQuestionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hskw_hskw_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatedQuestionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatedQuestionRequest) ProtoMessage() {}

func (x *CreatedQuestionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_hskw_hskw_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatedQuestionRequest.ProtoReflect.Descriptor instead.
func (*CreatedQuestionRequest) Descriptor() ([]byte, []int) {
	return file_hskw_hskw_proto_rawDescGZIP(), []int{5}
}

func (x *CreatedQuestionRequest) GetChallengeId() int32 {
	if x != nil {
		return x.ChallengeId
	}
	return 0
}

func (x *CreatedQuestionRequest) GetQnas() []*QNA {
	if x != nil {
		return x.Qnas
	}
	return nil
}

type CreatedQuestionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreatedQuestionResponse) Reset() {
	*x = CreatedQuestionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hskw_hskw_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatedQuestionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatedQuestionResponse) ProtoMessage() {}

func (x *CreatedQuestionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_hskw_hskw_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatedQuestionResponse.ProtoReflect.Descriptor instead.
func (*CreatedQuestionResponse) Descriptor() ([]byte, []int) {
	return file_hskw_hskw_proto_rawDescGZIP(), []int{6}
}

var File_hskw_hskw_proto protoreflect.FileDescriptor

var file_hskw_hskw_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x68, 0x73, 0x6b, 0x77, 0x2f, 0x68, 0x73, 0x6b, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x53, 0x0a, 0x0f, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x43, 0x61, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x63, 0x68, 0x61, 0x6c,
	0x6c, 0x65, 0x6e, 0x67, 0x65, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x50, 0x61, 0x74, 0x68, 0x22, 0x55, 0x0a, 0x16, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x74, 0x65, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72,
	0x12, 0x21, 0x0a, 0x0c, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67,
	0x65, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x61, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x61, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x54, 0x0a,
	0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x43, 0x61, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65,
	0x6e, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x63, 0x68,
	0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x61, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x61, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x22, 0x18, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x43, 0x61,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x39, 0x0a,
	0x03, 0x51, 0x4e, 0x41, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x16, 0x0a, 0x06, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x22, 0x55, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65,
	0x6e, 0x67, 0x65, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x04, 0x71, 0x6e, 0x61, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x04, 0x2e, 0x51, 0x4e, 0x41, 0x52, 0x04, 0x71, 0x6e, 0x61, 0x73, 0x22,
	0x19, 0x0a, 0x17, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x95, 0x01, 0x0a, 0x06, 0x53,
	0x6e, 0x61, 0x70, 0x70, 0x79, 0x12, 0x43, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x43, 0x61, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x43, 0x61, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x17, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x43, 0x61, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x46, 0x0a, 0x0f, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x17, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x1b, 0x5a, 0x19, 0x73, 0x65, 0x6e, 0x6b, 0x61, 0x77, 0x61, 0x2e, 0x6d, 0x6f,
	0x65, 0x2f, 0x68, 0x61, 0x61, 0x2d, 0x63, 0x68, 0x61, 0x6e, 0x2f, 0x68, 0x73, 0x6b, 0x77, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_hskw_hskw_proto_rawDescOnce sync.Once
	file_hskw_hskw_proto_rawDescData = file_hskw_hskw_proto_rawDesc
)

func file_hskw_hskw_proto_rawDescGZIP() []byte {
	file_hskw_hskw_proto_rawDescOnce.Do(func() {
		file_hskw_hskw_proto_rawDescData = protoimpl.X.CompressGZIP(file_hskw_hskw_proto_rawDescData)
	})
	return file_hskw_hskw_proto_rawDescData
}

var file_hskw_hskw_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_hskw_hskw_proto_goTypes = []interface{}{
	(*GenerateCaption)(nil),         // 0: GenerateCaption
	(*GenerateQuestionAnswer)(nil),  // 1: GenerateQuestionAnswer
	(*CreatedCaptionRequest)(nil),   // 2: CreatedCaptionRequest
	(*CreatedCaptionResponse)(nil),  // 3: CreatedCaptionResponse
	(*QNA)(nil),                     // 4: QNA
	(*CreatedQuestionRequest)(nil),  // 5: CreatedQuestionRequest
	(*CreatedQuestionResponse)(nil), // 6: CreatedQuestionResponse
}
var file_hskw_hskw_proto_depIdxs = []int32{
	4, // 0: CreatedQuestionRequest.qnas:type_name -> QNA
	2, // 1: Snappy.CreatedCaption:input_type -> CreatedCaptionRequest
	5, // 2: Snappy.CreatedQuestion:input_type -> CreatedQuestionRequest
	3, // 3: Snappy.CreatedCaption:output_type -> CreatedCaptionResponse
	6, // 4: Snappy.CreatedQuestion:output_type -> CreatedQuestionResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_hskw_hskw_proto_init() }
func file_hskw_hskw_proto_init() {
	if File_hskw_hskw_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_hskw_hskw_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenerateCaption); i {
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
		file_hskw_hskw_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenerateQuestionAnswer); i {
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
		file_hskw_hskw_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatedCaptionRequest); i {
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
		file_hskw_hskw_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatedCaptionResponse); i {
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
		file_hskw_hskw_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QNA); i {
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
		file_hskw_hskw_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatedQuestionRequest); i {
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
		file_hskw_hskw_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatedQuestionResponse); i {
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
			RawDescriptor: file_hskw_hskw_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_hskw_hskw_proto_goTypes,
		DependencyIndexes: file_hskw_hskw_proto_depIdxs,
		MessageInfos:      file_hskw_hskw_proto_msgTypes,
	}.Build()
	File_hskw_hskw_proto = out.File
	file_hskw_hskw_proto_rawDesc = nil
	file_hskw_hskw_proto_goTypes = nil
	file_hskw_hskw_proto_depIdxs = nil
}
