// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.1
// source: submission.proto

package client

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

type FindStudentFormSubmissionsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FormId string `protobuf:"bytes,1,opt,name=form_id,json=formId,proto3" json:"form_id,omitempty"`
}

func (x *FindStudentFormSubmissionsRequest) Reset() {
	*x = FindStudentFormSubmissionsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_submission_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindStudentFormSubmissionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindStudentFormSubmissionsRequest) ProtoMessage() {}

func (x *FindStudentFormSubmissionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_submission_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindStudentFormSubmissionsRequest.ProtoReflect.Descriptor instead.
func (*FindStudentFormSubmissionsRequest) Descriptor() ([]byte, []int) {
	return file_submission_proto_rawDescGZIP(), []int{0}
}

func (x *FindStudentFormSubmissionsRequest) GetFormId() string {
	if x != nil {
		return x.FormId
	}
	return ""
}

type FindStudentFormSubmissionsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StudentUsername     string                `protobuf:"bytes,1,opt,name=student_username,json=studentUsername,proto3" json:"student_username,omitempty"`
	FormId              string                `protobuf:"bytes,2,opt,name=form_id,json=formId,proto3" json:"form_id,omitempty"`
	QuestionSubmissions []*QuestionSubmission `protobuf:"bytes,3,rep,name=question_submissions,json=questionSubmissions,proto3" json:"question_submissions,omitempty"`
}

func (x *FindStudentFormSubmissionsResponse) Reset() {
	*x = FindStudentFormSubmissionsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_submission_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindStudentFormSubmissionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindStudentFormSubmissionsResponse) ProtoMessage() {}

func (x *FindStudentFormSubmissionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_submission_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindStudentFormSubmissionsResponse.ProtoReflect.Descriptor instead.
func (*FindStudentFormSubmissionsResponse) Descriptor() ([]byte, []int) {
	return file_submission_proto_rawDescGZIP(), []int{1}
}

func (x *FindStudentFormSubmissionsResponse) GetStudentUsername() string {
	if x != nil {
		return x.StudentUsername
	}
	return ""
}

func (x *FindStudentFormSubmissionsResponse) GetFormId() string {
	if x != nil {
		return x.FormId
	}
	return ""
}

func (x *FindStudentFormSubmissionsResponse) GetQuestionSubmissions() []*QuestionSubmission {
	if x != nil {
		return x.QuestionSubmissions
	}
	return nil
}

type QuestionSubmission struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	QuestionId string   `protobuf:"bytes,1,opt,name=question_id,json=questionId,proto3" json:"question_id,omitempty"`
	AnswerIds  []string `protobuf:"bytes,2,rep,name=answer_ids,json=answerIds,proto3" json:"answer_ids,omitempty"`
}

func (x *QuestionSubmission) Reset() {
	*x = QuestionSubmission{}
	if protoimpl.UnsafeEnabled {
		mi := &file_submission_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuestionSubmission) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuestionSubmission) ProtoMessage() {}

func (x *QuestionSubmission) ProtoReflect() protoreflect.Message {
	mi := &file_submission_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuestionSubmission.ProtoReflect.Descriptor instead.
func (*QuestionSubmission) Descriptor() ([]byte, []int) {
	return file_submission_proto_rawDescGZIP(), []int{2}
}

func (x *QuestionSubmission) GetQuestionId() string {
	if x != nil {
		return x.QuestionId
	}
	return ""
}

func (x *QuestionSubmission) GetAnswerIds() []string {
	if x != nil {
		return x.AnswerIds
	}
	return nil
}

var File_submission_proto protoreflect.FileDescriptor

var file_submission_proto_rawDesc = []byte{
	0x0a, 0x10, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x03, 0x61, 0x70, 0x69, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x46, 0x0a, 0x21, 0x46, 0x69, 0x6e, 0x64, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x46,
	0x6f, 0x72, 0x6d, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x07, 0x66, 0x6f, 0x72, 0x6d, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x72, 0x03, 0xb0, 0x01, 0x01,
	0x52, 0x06, 0x66, 0x6f, 0x72, 0x6d, 0x49, 0x64, 0x22, 0xb4, 0x01, 0x0a, 0x22, 0x46, 0x69, 0x6e,
	0x64, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x46, 0x6f, 0x72, 0x6d, 0x53, 0x75, 0x62, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x29, 0x0a, 0x10, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x73, 0x74, 0x75, 0x64, 0x65,
	0x6e, 0x74, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x66, 0x6f,
	0x72, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x6f, 0x72,
	0x6d, 0x49, 0x64, 0x12, 0x4a, 0x0a, 0x14, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x73, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x17, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x13, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x69, 0x6f, 0x6e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x22,
	0x54, 0x0a, 0x12, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x75, 0x62, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x61, 0x6e, 0x73, 0x77,
	0x65, 0x72, 0x49, 0x64, 0x73, 0x32, 0x7b, 0x0a, 0x0a, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x6d, 0x0a, 0x1a, 0x46, 0x69, 0x6e, 0x64, 0x53, 0x74, 0x75, 0x64, 0x65,
	0x6e, 0x74, 0x46, 0x6f, 0x72, 0x6d, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x12, 0x26, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x53, 0x74, 0x75, 0x64,
	0x65, 0x6e, 0x74, 0x46, 0x6f, 0x72, 0x6d, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x46, 0x69, 0x6e, 0x64, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x46, 0x6f, 0x72, 0x6d, 0x53,
	0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x1e, 0x5a, 0x1c, 0x75, 0x70, 0x61, 0x73, 0x73, 0x65, 0x64, 0x2e, 0x61, 0x73,
	0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x3b, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_submission_proto_rawDescOnce sync.Once
	file_submission_proto_rawDescData = file_submission_proto_rawDesc
)

func file_submission_proto_rawDescGZIP() []byte {
	file_submission_proto_rawDescOnce.Do(func() {
		file_submission_proto_rawDescData = protoimpl.X.CompressGZIP(file_submission_proto_rawDescData)
	})
	return file_submission_proto_rawDescData
}

var file_submission_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_submission_proto_goTypes = []any{
	(*FindStudentFormSubmissionsRequest)(nil),  // 0: api.FindStudentFormSubmissionsRequest
	(*FindStudentFormSubmissionsResponse)(nil), // 1: api.FindStudentFormSubmissionsResponse
	(*QuestionSubmission)(nil),                 // 2: api.QuestionSubmission
}
var file_submission_proto_depIdxs = []int32{
	2, // 0: api.FindStudentFormSubmissionsResponse.question_submissions:type_name -> api.QuestionSubmission
	0, // 1: api.Submission.FindStudentFormSubmissions:input_type -> api.FindStudentFormSubmissionsRequest
	1, // 2: api.Submission.FindStudentFormSubmissions:output_type -> api.FindStudentFormSubmissionsResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_submission_proto_init() }
func file_submission_proto_init() {
	if File_submission_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_submission_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*FindStudentFormSubmissionsRequest); i {
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
		file_submission_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*FindStudentFormSubmissionsResponse); i {
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
		file_submission_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*QuestionSubmission); i {
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
			RawDescriptor: file_submission_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_submission_proto_goTypes,
		DependencyIndexes: file_submission_proto_depIdxs,
		MessageInfos:      file_submission_proto_msgTypes,
	}.Build()
	File_submission_proto = out.File
	file_submission_proto_rawDesc = nil
	file_submission_proto_goTypes = nil
	file_submission_proto_depIdxs = nil
}