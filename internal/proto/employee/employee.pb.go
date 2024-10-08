// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.1
// source: internal/proto/employee/employee.proto

package __

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

type Employee struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           uint64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Username     string  `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Password     *string `protobuf:"bytes,3,opt,name=password,proto3,oneof" json:"password,omitempty"` // Hashed
	ProfileImage string  `protobuf:"bytes,4,opt,name=profileImage,proto3" json:"profileImage,omitempty"`
	Salary       float32 `protobuf:"fixed32,5,opt,name=salary,proto3" json:"salary,omitempty"`
}

func (x *Employee) Reset() {
	*x = Employee{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_employee_employee_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Employee) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Employee) ProtoMessage() {}

func (x *Employee) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_employee_employee_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Employee.ProtoReflect.Descriptor instead.
func (*Employee) Descriptor() ([]byte, []int) {
	return file_internal_proto_employee_employee_proto_rawDescGZIP(), []int{0}
}

func (x *Employee) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Employee) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *Employee) GetPassword() string {
	if x != nil && x.Password != nil {
		return *x.Password
	}
	return ""
}

func (x *Employee) GetProfileImage() string {
	if x != nil {
		return x.ProfileImage
	}
	return ""
}

func (x *Employee) GetSalary() float32 {
	if x != nil {
		return x.Salary
	}
	return 0
}

type Employees struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Employees []*Employee `protobuf:"bytes,1,rep,name=employees,proto3" json:"employees,omitempty"`
}

func (x *Employees) Reset() {
	*x = Employees{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_employee_employee_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Employees) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Employees) ProtoMessage() {}

func (x *Employees) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_employee_employee_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Employees.ProtoReflect.Descriptor instead.
func (*Employees) Descriptor() ([]byte, []int) {
	return file_internal_proto_employee_employee_proto_rawDescGZIP(), []int{1}
}

func (x *Employees) GetEmployees() []*Employee {
	if x != nil {
		return x.Employees
	}
	return nil
}

var File_internal_proto_employee_employee_proto protoreflect.FileDescriptor

var file_internal_proto_employee_employee_proto_rawDesc = []byte{
	0x0a, 0x26, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x2f, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79,
	0x65, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa0, 0x01, 0x0a, 0x08, 0x45, 0x6d, 0x70,
	0x6c, 0x6f, 0x79, 0x65, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x1f, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x88,
	0x01, 0x01, 0x12, 0x22, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x6d, 0x61,
	0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x61, 0x6c, 0x61, 0x72, 0x79,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x73, 0x61, 0x6c, 0x61, 0x72, 0x79, 0x42, 0x0b,
	0x0a, 0x09, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x34, 0x0a, 0x09, 0x45,
	0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x73, 0x12, 0x27, 0x0a, 0x09, 0x65, 0x6d, 0x70, 0x6c,
	0x6f, 0x79, 0x65, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x45, 0x6d,
	0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x52, 0x09, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65,
	0x73, 0x42, 0x03, 0x5a, 0x01, 0x2e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_proto_employee_employee_proto_rawDescOnce sync.Once
	file_internal_proto_employee_employee_proto_rawDescData = file_internal_proto_employee_employee_proto_rawDesc
)

func file_internal_proto_employee_employee_proto_rawDescGZIP() []byte {
	file_internal_proto_employee_employee_proto_rawDescOnce.Do(func() {
		file_internal_proto_employee_employee_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_proto_employee_employee_proto_rawDescData)
	})
	return file_internal_proto_employee_employee_proto_rawDescData
}

var file_internal_proto_employee_employee_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_internal_proto_employee_employee_proto_goTypes = []any{
	(*Employee)(nil),  // 0: Employee
	(*Employees)(nil), // 1: Employees
}
var file_internal_proto_employee_employee_proto_depIdxs = []int32{
	0, // 0: Employees.employees:type_name -> Employee
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_internal_proto_employee_employee_proto_init() }
func file_internal_proto_employee_employee_proto_init() {
	if File_internal_proto_employee_employee_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_proto_employee_employee_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Employee); i {
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
		file_internal_proto_employee_employee_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*Employees); i {
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
	file_internal_proto_employee_employee_proto_msgTypes[0].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_internal_proto_employee_employee_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_internal_proto_employee_employee_proto_goTypes,
		DependencyIndexes: file_internal_proto_employee_employee_proto_depIdxs,
		MessageInfos:      file_internal_proto_employee_employee_proto_msgTypes,
	}.Build()
	File_internal_proto_employee_employee_proto = out.File
	file_internal_proto_employee_employee_proto_rawDesc = nil
	file_internal_proto_employee_employee_proto_goTypes = nil
	file_internal_proto_employee_employee_proto_depIdxs = nil
}
