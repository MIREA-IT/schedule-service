// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: api/proto/scheduleService.proto

package api

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

type GroupScheduleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GroupCode     string `protobuf:"bytes,1,opt,name=groupCode,proto3" json:"groupCode,omitempty"`
	InitiatorUser string `protobuf:"bytes,2,opt,name=initiatorUser,proto3" json:"initiatorUser,omitempty"`
}

func (x *GroupScheduleRequest) Reset() {
	*x = GroupScheduleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_scheduleService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupScheduleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupScheduleRequest) ProtoMessage() {}

func (x *GroupScheduleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_scheduleService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupScheduleRequest.ProtoReflect.Descriptor instead.
func (*GroupScheduleRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_scheduleService_proto_rawDescGZIP(), []int{0}
}

func (x *GroupScheduleRequest) GetGroupCode() string {
	if x != nil {
		return x.GroupCode
	}
	return ""
}

func (x *GroupScheduleRequest) GetInitiatorUser() string {
	if x != nil {
		return x.InitiatorUser
	}
	return ""
}

type GroupSchedule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GroupCode string `protobuf:"bytes,1,opt,name=groupCode,proto3" json:"groupCode,omitempty"`
	OddWeek   *Week  `protobuf:"bytes,2,opt,name=oddWeek,proto3" json:"oddWeek,omitempty"`
	EvenWeek  *Week  `protobuf:"bytes,3,opt,name=evenWeek,proto3" json:"evenWeek,omitempty"`
}

func (x *GroupSchedule) Reset() {
	*x = GroupSchedule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_scheduleService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupSchedule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupSchedule) ProtoMessage() {}

func (x *GroupSchedule) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_scheduleService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupSchedule.ProtoReflect.Descriptor instead.
func (*GroupSchedule) Descriptor() ([]byte, []int) {
	return file_api_proto_scheduleService_proto_rawDescGZIP(), []int{1}
}

func (x *GroupSchedule) GetGroupCode() string {
	if x != nil {
		return x.GroupCode
	}
	return ""
}

func (x *GroupSchedule) GetOddWeek() *Week {
	if x != nil {
		return x.OddWeek
	}
	return nil
}

func (x *GroupSchedule) GetEvenWeek() *Week {
	if x != nil {
		return x.EvenWeek
	}
	return nil
}

type Week struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Monday    *WeekDay `protobuf:"bytes,1,opt,name=monday,proto3" json:"monday,omitempty"`
	Tuesday   *WeekDay `protobuf:"bytes,2,opt,name=tuesday,proto3" json:"tuesday,omitempty"`
	Wednesday *WeekDay `protobuf:"bytes,3,opt,name=wednesday,proto3" json:"wednesday,omitempty"`
	Thursday  *WeekDay `protobuf:"bytes,4,opt,name=thursday,proto3" json:"thursday,omitempty"`
	Friday    *WeekDay `protobuf:"bytes,5,opt,name=friday,proto3" json:"friday,omitempty"`
	Saturday  *WeekDay `protobuf:"bytes,6,opt,name=saturday,proto3" json:"saturday,omitempty"`
}

func (x *Week) Reset() {
	*x = Week{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_scheduleService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Week) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Week) ProtoMessage() {}

func (x *Week) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_scheduleService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Week.ProtoReflect.Descriptor instead.
func (*Week) Descriptor() ([]byte, []int) {
	return file_api_proto_scheduleService_proto_rawDescGZIP(), []int{2}
}

func (x *Week) GetMonday() *WeekDay {
	if x != nil {
		return x.Monday
	}
	return nil
}

func (x *Week) GetTuesday() *WeekDay {
	if x != nil {
		return x.Tuesday
	}
	return nil
}

func (x *Week) GetWednesday() *WeekDay {
	if x != nil {
		return x.Wednesday
	}
	return nil
}

func (x *Week) GetThursday() *WeekDay {
	if x != nil {
		return x.Thursday
	}
	return nil
}

func (x *Week) GetFriday() *WeekDay {
	if x != nil {
		return x.Friday
	}
	return nil
}

func (x *Week) GetSaturday() *WeekDay {
	if x != nil {
		return x.Saturday
	}
	return nil
}

type WeekDay struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Lessons []*Lesson `protobuf:"bytes,1,rep,name=lessons,proto3" json:"lessons,omitempty"`
}

func (x *WeekDay) Reset() {
	*x = WeekDay{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_scheduleService_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WeekDay) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WeekDay) ProtoMessage() {}

func (x *WeekDay) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_scheduleService_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WeekDay.ProtoReflect.Descriptor instead.
func (*WeekDay) Descriptor() ([]byte, []int) {
	return file_api_proto_scheduleService_proto_rawDescGZIP(), []int{3}
}

func (x *WeekDay) GetLessons() []*Lesson {
	if x != nil {
		return x.Lessons
	}
	return nil
}

type Lesson struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number   uint32 `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Type     string `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	Teacher  string `protobuf:"bytes,4,opt,name=teacher,proto3" json:"teacher,omitempty"`
	Audience string `protobuf:"bytes,5,opt,name=audience,proto3" json:"audience,omitempty"`
}

func (x *Lesson) Reset() {
	*x = Lesson{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_scheduleService_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Lesson) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Lesson) ProtoMessage() {}

func (x *Lesson) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_scheduleService_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Lesson.ProtoReflect.Descriptor instead.
func (*Lesson) Descriptor() ([]byte, []int) {
	return file_api_proto_scheduleService_proto_rawDescGZIP(), []int{4}
}

func (x *Lesson) GetNumber() uint32 {
	if x != nil {
		return x.Number
	}
	return 0
}

func (x *Lesson) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Lesson) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Lesson) GetTeacher() string {
	if x != nil {
		return x.Teacher
	}
	return ""
}

func (x *Lesson) GetAudience() string {
	if x != nil {
		return x.Audience
	}
	return ""
}

var File_api_proto_scheduleService_proto protoreflect.FileDescriptor

var file_api_proto_scheduleService_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x63, 0x68, 0x65,
	0x64, 0x75, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x03, 0x61, 0x70, 0x69, 0x22, 0x5a, 0x0a, 0x14, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x53,
	0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c,
	0x0a, 0x09, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x24, 0x0a, 0x0d,
	0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x74, 0x6f, 0x72, 0x55, 0x73, 0x65, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x74, 0x6f, 0x72, 0x55, 0x73,
	0x65, 0x72, 0x22, 0x79, 0x0a, 0x0d, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x63, 0x68, 0x65, 0x64,
	0x75, 0x6c, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x43, 0x6f, 0x64, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x43, 0x6f, 0x64,
	0x65, 0x12, 0x23, 0x0a, 0x07, 0x6f, 0x64, 0x64, 0x57, 0x65, 0x65, 0x6b, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x09, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x57, 0x65, 0x65, 0x6b, 0x52, 0x07, 0x6f,
	0x64, 0x64, 0x57, 0x65, 0x65, 0x6b, 0x12, 0x25, 0x0a, 0x08, 0x65, 0x76, 0x65, 0x6e, 0x57, 0x65,
	0x65, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x57,
	0x65, 0x65, 0x6b, 0x52, 0x08, 0x65, 0x76, 0x65, 0x6e, 0x57, 0x65, 0x65, 0x6b, 0x22, 0xfa, 0x01,
	0x0a, 0x04, 0x57, 0x65, 0x65, 0x6b, 0x12, 0x24, 0x0a, 0x06, 0x6d, 0x6f, 0x6e, 0x64, 0x61, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x57, 0x65, 0x65,
	0x6b, 0x44, 0x61, 0x79, 0x52, 0x06, 0x6d, 0x6f, 0x6e, 0x64, 0x61, 0x79, 0x12, 0x26, 0x0a, 0x07,
	0x74, 0x75, 0x65, 0x73, 0x64, 0x61, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x57, 0x65, 0x65, 0x6b, 0x44, 0x61, 0x79, 0x52, 0x07, 0x74, 0x75, 0x65,
	0x73, 0x64, 0x61, 0x79, 0x12, 0x2a, 0x0a, 0x09, 0x77, 0x65, 0x64, 0x6e, 0x65, 0x73, 0x64, 0x61,
	0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x57, 0x65,
	0x65, 0x6b, 0x44, 0x61, 0x79, 0x52, 0x09, 0x77, 0x65, 0x64, 0x6e, 0x65, 0x73, 0x64, 0x61, 0x79,
	0x12, 0x28, 0x0a, 0x08, 0x74, 0x68, 0x75, 0x72, 0x73, 0x64, 0x61, 0x79, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x57, 0x65, 0x65, 0x6b, 0x44, 0x61, 0x79,
	0x52, 0x08, 0x74, 0x68, 0x75, 0x72, 0x73, 0x64, 0x61, 0x79, 0x12, 0x24, 0x0a, 0x06, 0x66, 0x72,
	0x69, 0x64, 0x61, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x57, 0x65, 0x65, 0x6b, 0x44, 0x61, 0x79, 0x52, 0x06, 0x66, 0x72, 0x69, 0x64, 0x61, 0x79,
	0x12, 0x28, 0x0a, 0x08, 0x73, 0x61, 0x74, 0x75, 0x72, 0x64, 0x61, 0x79, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x57, 0x65, 0x65, 0x6b, 0x44, 0x61, 0x79,
	0x52, 0x08, 0x73, 0x61, 0x74, 0x75, 0x72, 0x64, 0x61, 0x79, 0x22, 0x30, 0x0a, 0x07, 0x57, 0x65,
	0x65, 0x6b, 0x44, 0x61, 0x79, 0x12, 0x25, 0x0a, 0x07, 0x6c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x65, 0x73,
	0x73, 0x6f, 0x6e, 0x52, 0x07, 0x6c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x73, 0x22, 0x7e, 0x0a, 0x06,
	0x4c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65,
	0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72,
	0x12, 0x1a, 0x0a, 0x08, 0x61, 0x75, 0x64, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x61, 0x75, 0x64, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x32, 0x56, 0x0a, 0x0f,
	0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x43, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x63, 0x68, 0x65, 0x64,
	0x75, 0x6c, 0x65, 0x12, 0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x53,
	0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75,
	0x6c, 0x65, 0x22, 0x00, 0x42, 0x09, 0x5a, 0x07, 0x73, 0x72, 0x63, 0x2f, 0x61, 0x70, 0x69, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_proto_scheduleService_proto_rawDescOnce sync.Once
	file_api_proto_scheduleService_proto_rawDescData = file_api_proto_scheduleService_proto_rawDesc
)

func file_api_proto_scheduleService_proto_rawDescGZIP() []byte {
	file_api_proto_scheduleService_proto_rawDescOnce.Do(func() {
		file_api_proto_scheduleService_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_scheduleService_proto_rawDescData)
	})
	return file_api_proto_scheduleService_proto_rawDescData
}

var file_api_proto_scheduleService_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_api_proto_scheduleService_proto_goTypes = []interface{}{
	(*GroupScheduleRequest)(nil), // 0: api.GroupScheduleRequest
	(*GroupSchedule)(nil),        // 1: api.GroupSchedule
	(*Week)(nil),                 // 2: api.Week
	(*WeekDay)(nil),              // 3: api.WeekDay
	(*Lesson)(nil),               // 4: api.Lesson
}
var file_api_proto_scheduleService_proto_depIdxs = []int32{
	2,  // 0: api.GroupSchedule.oddWeek:type_name -> api.Week
	2,  // 1: api.GroupSchedule.evenWeek:type_name -> api.Week
	3,  // 2: api.Week.monday:type_name -> api.WeekDay
	3,  // 3: api.Week.tuesday:type_name -> api.WeekDay
	3,  // 4: api.Week.wednesday:type_name -> api.WeekDay
	3,  // 5: api.Week.thursday:type_name -> api.WeekDay
	3,  // 6: api.Week.friday:type_name -> api.WeekDay
	3,  // 7: api.Week.saturday:type_name -> api.WeekDay
	4,  // 8: api.WeekDay.lessons:type_name -> api.Lesson
	0,  // 9: api.ScheduleService.GetGroupSchedule:input_type -> api.GroupScheduleRequest
	1,  // 10: api.ScheduleService.GetGroupSchedule:output_type -> api.GroupSchedule
	10, // [10:11] is the sub-list for method output_type
	9,  // [9:10] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_api_proto_scheduleService_proto_init() }
func file_api_proto_scheduleService_proto_init() {
	if File_api_proto_scheduleService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_proto_scheduleService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupScheduleRequest); i {
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
		file_api_proto_scheduleService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupSchedule); i {
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
		file_api_proto_scheduleService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Week); i {
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
		file_api_proto_scheduleService_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WeekDay); i {
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
		file_api_proto_scheduleService_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Lesson); i {
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
			RawDescriptor: file_api_proto_scheduleService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_proto_scheduleService_proto_goTypes,
		DependencyIndexes: file_api_proto_scheduleService_proto_depIdxs,
		MessageInfos:      file_api_proto_scheduleService_proto_msgTypes,
	}.Build()
	File_api_proto_scheduleService_proto = out.File
	file_api_proto_scheduleService_proto_rawDesc = nil
	file_api_proto_scheduleService_proto_goTypes = nil
	file_api_proto_scheduleService_proto_depIdxs = nil
}
