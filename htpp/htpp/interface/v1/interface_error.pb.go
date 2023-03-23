// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: apis/htpp/htpp/interface/v1/interface_error.proto

package v1

import (
	_ "github.com/go-kratos/kratos/v2/errors"
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

type HtppInterfaceErrorReason int32

const (
	HtppInterfaceErrorReason_UNKNOWN_ERROR     HtppInterfaceErrorReason = 0
	HtppInterfaceErrorReason_NOT_FOUND_ERROR   HtppInterfaceErrorReason = 1
	HtppInterfaceErrorReason_LOGIN_FAILED      HtppInterfaceErrorReason = 2
	HtppInterfaceErrorReason_REGISTER_FAILED   HtppInterfaceErrorReason = 3
	HtppInterfaceErrorReason_USERNAME_CONFLICT HtppInterfaceErrorReason = 4
	HtppInterfaceErrorReason_READ_DEVICE_ERROR HtppInterfaceErrorReason = 5
	HtppInterfaceErrorReason_CREATE_FAILED     HtppInterfaceErrorReason = 6
	HtppInterfaceErrorReason_ADDRESS_CONFLICT  HtppInterfaceErrorReason = 7
)

// Enum value maps for HtppInterfaceErrorReason.
var (
	HtppInterfaceErrorReason_name = map[int32]string{
		0: "UNKNOWN_ERROR",
		1: "NOT_FOUND_ERROR",
		2: "LOGIN_FAILED",
		3: "REGISTER_FAILED",
		4: "USERNAME_CONFLICT",
		5: "READ_DEVICE_ERROR",
		6: "CREATE_FAILED",
		7: "ADDRESS_CONFLICT",
	}
	HtppInterfaceErrorReason_value = map[string]int32{
		"UNKNOWN_ERROR":     0,
		"NOT_FOUND_ERROR":   1,
		"LOGIN_FAILED":      2,
		"REGISTER_FAILED":   3,
		"USERNAME_CONFLICT": 4,
		"READ_DEVICE_ERROR": 5,
		"CREATE_FAILED":     6,
		"ADDRESS_CONFLICT":  7,
	}
)

func (x HtppInterfaceErrorReason) Enum() *HtppInterfaceErrorReason {
	p := new(HtppInterfaceErrorReason)
	*p = x
	return p
}

func (x HtppInterfaceErrorReason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (HtppInterfaceErrorReason) Descriptor() protoreflect.EnumDescriptor {
	return file_apis_htpp_htpp_interface_v1_interface_error_proto_enumTypes[0].Descriptor()
}

func (HtppInterfaceErrorReason) Type() protoreflect.EnumType {
	return &file_apis_htpp_htpp_interface_v1_interface_error_proto_enumTypes[0]
}

func (x HtppInterfaceErrorReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use HtppInterfaceErrorReason.Descriptor instead.
func (HtppInterfaceErrorReason) EnumDescriptor() ([]byte, []int) {
	return file_apis_htpp_htpp_interface_v1_interface_error_proto_rawDescGZIP(), []int{0}
}

var File_apis_htpp_htpp_interface_v1_interface_error_proto protoreflect.FileDescriptor

var file_apis_htpp_htpp_interface_v1_interface_error_proto_rawDesc = []byte{
	0x0a, 0x31, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x68, 0x74, 0x70, 0x70, 0x2f, 0x68, 0x74, 0x70, 0x70,
	0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x16, 0x68, 0x74, 0x70, 0x70, 0x2e, 0x68, 0x74, 0x70, 0x70, 0x2e, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x13, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x73, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2a, 0xc6, 0x01, 0x0a, 0x18, 0x48, 0x74, 0x70, 0x70, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61,
	0x63, 0x65, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x11, 0x0a,
	0x0d, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x00,
	0x12, 0x13, 0x0a, 0x0f, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x5f, 0x45, 0x52,
	0x52, 0x4f, 0x52, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x4c, 0x4f, 0x47, 0x49, 0x4e, 0x5f, 0x46,
	0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x02, 0x12, 0x13, 0x0a, 0x0f, 0x52, 0x45, 0x47, 0x49, 0x53,
	0x54, 0x45, 0x52, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x03, 0x12, 0x15, 0x0a, 0x11,
	0x55, 0x53, 0x45, 0x52, 0x4e, 0x41, 0x4d, 0x45, 0x5f, 0x43, 0x4f, 0x4e, 0x46, 0x4c, 0x49, 0x43,
	0x54, 0x10, 0x04, 0x12, 0x15, 0x0a, 0x11, 0x52, 0x45, 0x41, 0x44, 0x5f, 0x44, 0x45, 0x56, 0x49,
	0x43, 0x45, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x05, 0x12, 0x11, 0x0a, 0x0d, 0x43, 0x52,
	0x45, 0x41, 0x54, 0x45, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x06, 0x12, 0x14, 0x0a,
	0x10, 0x41, 0x44, 0x44, 0x52, 0x45, 0x53, 0x53, 0x5f, 0x43, 0x4f, 0x4e, 0x46, 0x4c, 0x49, 0x43,
	0x54, 0x10, 0x07, 0x1a, 0x04, 0xa0, 0x45, 0xf4, 0x03, 0x42, 0x1d, 0x50, 0x01, 0x5a, 0x19, 0x68,
	0x74, 0x70, 0x70, 0x2f, 0x68, 0x74, 0x70, 0x70, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61,
	0x63, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_apis_htpp_htpp_interface_v1_interface_error_proto_rawDescOnce sync.Once
	file_apis_htpp_htpp_interface_v1_interface_error_proto_rawDescData = file_apis_htpp_htpp_interface_v1_interface_error_proto_rawDesc
)

func file_apis_htpp_htpp_interface_v1_interface_error_proto_rawDescGZIP() []byte {
	file_apis_htpp_htpp_interface_v1_interface_error_proto_rawDescOnce.Do(func() {
		file_apis_htpp_htpp_interface_v1_interface_error_proto_rawDescData = protoimpl.X.CompressGZIP(file_apis_htpp_htpp_interface_v1_interface_error_proto_rawDescData)
	})
	return file_apis_htpp_htpp_interface_v1_interface_error_proto_rawDescData
}

var file_apis_htpp_htpp_interface_v1_interface_error_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_apis_htpp_htpp_interface_v1_interface_error_proto_goTypes = []interface{}{
	(HtppInterfaceErrorReason)(0), // 0: htpp.htpp.interface.v1.HtppInterfaceErrorReason
}
var file_apis_htpp_htpp_interface_v1_interface_error_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_apis_htpp_htpp_interface_v1_interface_error_proto_init() }
func file_apis_htpp_htpp_interface_v1_interface_error_proto_init() {
	if File_apis_htpp_htpp_interface_v1_interface_error_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_apis_htpp_htpp_interface_v1_interface_error_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_apis_htpp_htpp_interface_v1_interface_error_proto_goTypes,
		DependencyIndexes: file_apis_htpp_htpp_interface_v1_interface_error_proto_depIdxs,
		EnumInfos:         file_apis_htpp_htpp_interface_v1_interface_error_proto_enumTypes,
	}.Build()
	File_apis_htpp_htpp_interface_v1_interface_error_proto = out.File
	file_apis_htpp_htpp_interface_v1_interface_error_proto_rawDesc = nil
	file_apis_htpp_htpp_interface_v1_interface_error_proto_goTypes = nil
	file_apis_htpp_htpp_interface_v1_interface_error_proto_depIdxs = nil
}
