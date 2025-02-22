// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: proto/todo.proto

package todo

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

type NewTodo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Done        bool   `protobuf:"varint,3,opt,name=done,proto3" json:"done,omitempty"`
}

func (x *NewTodo) Reset() {
	*x = NewTodo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_todo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewTodo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewTodo) ProtoMessage() {}

func (x *NewTodo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_todo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewTodo.ProtoReflect.Descriptor instead.
func (*NewTodo) Descriptor() ([]byte, []int) {
	return file_proto_todo_proto_rawDescGZIP(), []int{0}
}

func (x *NewTodo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *NewTodo) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *NewTodo) GetDone() bool {
	if x != nil {
		return x.Done
	}
	return false
}

type Todo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Done        bool   `protobuf:"varint,4,opt,name=done,proto3" json:"done,omitempty"`
}

func (x *Todo) Reset() {
	*x = Todo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_todo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Todo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Todo) ProtoMessage() {}

func (x *Todo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_todo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Todo.ProtoReflect.Descriptor instead.
func (*Todo) Descriptor() ([]byte, []int) {
	return file_proto_todo_proto_rawDescGZIP(), []int{1}
}

func (x *Todo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Todo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Todo) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Todo) GetDone() bool {
	if x != nil {
		return x.Done
	}
	return false
}

var File_proto_todo_proto protoreflect.FileDescriptor

var file_proto_todo_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x04, 0x74, 0x6f, 0x64, 0x6f, 0x22, 0x53, 0x0a, 0x07, 0x4e, 0x65, 0x77, 0x54,
	0x6f, 0x64, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x6f, 0x6e,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x64, 0x6f, 0x6e, 0x65, 0x22, 0x60, 0x0a,
	0x04, 0x54, 0x6f, 0x64, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x64,
	0x6f, 0x6e, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x64, 0x6f, 0x6e, 0x65, 0x32,
	0x36, 0x0a, 0x0b, 0x54, 0x6f, 0x64, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x27,
	0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x64, 0x6f, 0x12, 0x0d, 0x2e, 0x74,
	0x6f, 0x64, 0x6f, 0x2e, 0x4e, 0x65, 0x77, 0x54, 0x6f, 0x64, 0x6f, 0x1a, 0x0a, 0x2e, 0x74, 0x6f,
	0x64, 0x6f, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x42, 0x27, 0x5a, 0x25, 0x6c, 0x6f, 0x63, 0x61, 0x6c,
	0x68, 0x6f, 0x73, 0x74, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x70, 0x6f, 0x73,
	0x74, 0x67, 0x72, 0x65, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x6f, 0x64, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_todo_proto_rawDescOnce sync.Once
	file_proto_todo_proto_rawDescData = file_proto_todo_proto_rawDesc
)

func file_proto_todo_proto_rawDescGZIP() []byte {
	file_proto_todo_proto_rawDescOnce.Do(func() {
		file_proto_todo_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_todo_proto_rawDescData)
	})
	return file_proto_todo_proto_rawDescData
}

var file_proto_todo_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_todo_proto_goTypes = []interface{}{
	(*NewTodo)(nil), // 0: todo.NewTodo
	(*Todo)(nil),    // 1: todo.Todo
}
var file_proto_todo_proto_depIdxs = []int32{
	0, // 0: todo.TodoService.CreateTodo:input_type -> todo.NewTodo
	1, // 1: todo.TodoService.CreateTodo:output_type -> todo.Todo
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_todo_proto_init() }
func file_proto_todo_proto_init() {
	if File_proto_todo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_todo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewTodo); i {
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
		file_proto_todo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Todo); i {
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
			RawDescriptor: file_proto_todo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_todo_proto_goTypes,
		DependencyIndexes: file_proto_todo_proto_depIdxs,
		MessageInfos:      file_proto_todo_proto_msgTypes,
	}.Build()
	File_proto_todo_proto = out.File
	file_proto_todo_proto_rawDesc = nil
	file_proto_todo_proto_goTypes = nil
	file_proto_todo_proto_depIdxs = nil
}
