// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: pkg/secrets/protocol.proto

package secrets

import (
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

var File_pkg_secrets_protocol_proto protoreflect.FileDescriptor

var file_pkg_secrets_protocol_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x70, 0x6b, 0x67, 0x2f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x65,
	0x63, 0x72, 0x65, 0x74, 0x73, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x19, 0x70, 0x6b, 0x67, 0x2f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x2f,
	0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xc9, 0x04,
	0x0a, 0x0e, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x5b, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74,
	0x12, 0x23, 0x2e, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x2e, 0x53, 0x65, 0x63, 0x72, 0x65,
	0x74, 0x73, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x2e,
	0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x75, 0x63,
	0x6b, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4d, 0x0a,
	0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x23, 0x2e,
	0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x2e, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x45, 0x0a, 0x0b,
	0x53, 0x74, 0x6f, 0x72, 0x65, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x1c, 0x2e, 0x73, 0x65,
	0x63, 0x72, 0x65, 0x74, 0x73, 0x2e, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x53, 0x74, 0x6f,
	0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x22, 0x00, 0x12, 0x55, 0x0a, 0x0e, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x53,
	0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x1f, 0x2e, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x2e,
	0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73,
	0x2e, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4f, 0x0a, 0x0c, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x1d, 0x2e, 0x73, 0x65, 0x63,
	0x72, 0x65, 0x74, 0x73, 0x2e, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x73, 0x65, 0x63, 0x72,
	0x65, 0x74, 0x73, 0x2e, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x47, 0x0a, 0x0a, 0x47,
	0x65, 0x74, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x12, 0x1a, 0x2e, 0x73, 0x65, 0x63, 0x72,
	0x65, 0x74, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x2e,
	0x47, 0x65, 0x74, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x53, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x53, 0x65, 0x63, 0x72, 0x65,
	0x74, 0x73, 0x57, 0x69, 0x74, 0x68, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1a, 0x2e, 0x73, 0x65, 0x63,
	0x72, 0x65, 0x74, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73,
	0x2e, 0x47, 0x65, 0x74, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x44, 0x61, 0x74, 0x61, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x29, 0x5a, 0x27, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x76, 0x6f, 0x72, 0x74, 0x65, 0x69, 0x6c, 0x2f,
	0x64, 0x69, 0x72, 0x65, 0x6b, 0x74, 0x69, 0x76, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x73, 0x65, 0x63,
	0x72, 0x65, 0x74, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_pkg_secrets_protocol_proto_goTypes = []interface{}{
	(*SecretsCreateBucketRequest)(nil),  // 0: secrets.SecretsCreateBucketRequest
	(*SecretsDeleteBucketRequest)(nil),  // 1: secrets.SecretsDeleteBucketRequest
	(*SecretsStoreRequest)(nil),         // 2: secrets.SecretsStoreRequest
	(*SecretsRetrieveRequest)(nil),      // 3: secrets.SecretsRetrieveRequest
	(*SecretsDeleteRequest)(nil),        // 4: secrets.SecretsDeleteRequest
	(*GetSecretsRequest)(nil),           // 5: secrets.GetSecretsRequest
	(*SecretsCreateBucketResponse)(nil), // 6: secrets.SecretsCreateBucketResponse
	(*empty.Empty)(nil),                 // 7: google.protobuf.Empty
	(*SecretsRetrieveResponse)(nil),     // 8: secrets.SecretsRetrieveResponse
	(*SecretsDeleteResponse)(nil),       // 9: secrets.SecretsDeleteResponse
	(*GetSecretsResponse)(nil),          // 10: secrets.GetSecretsResponse
	(*GetSecretsDataResponse)(nil),      // 11: secrets.GetSecretsDataResponse
}
var file_pkg_secrets_protocol_proto_depIdxs = []int32{
	0,  // 0: secrets.SecretsService.CreateBucket:input_type -> secrets.SecretsCreateBucketRequest
	1,  // 1: secrets.SecretsService.DeleteBucket:input_type -> secrets.SecretsDeleteBucketRequest
	2,  // 2: secrets.SecretsService.StoreSecret:input_type -> secrets.SecretsStoreRequest
	3,  // 3: secrets.SecretsService.RetrieveSecret:input_type -> secrets.SecretsRetrieveRequest
	4,  // 4: secrets.SecretsService.DeleteSecret:input_type -> secrets.SecretsDeleteRequest
	5,  // 5: secrets.SecretsService.GetSecrets:input_type -> secrets.GetSecretsRequest
	5,  // 6: secrets.SecretsService.GetSecretsWithData:input_type -> secrets.GetSecretsRequest
	6,  // 7: secrets.SecretsService.CreateBucket:output_type -> secrets.SecretsCreateBucketResponse
	7,  // 8: secrets.SecretsService.DeleteBucket:output_type -> google.protobuf.Empty
	7,  // 9: secrets.SecretsService.StoreSecret:output_type -> google.protobuf.Empty
	8,  // 10: secrets.SecretsService.RetrieveSecret:output_type -> secrets.SecretsRetrieveResponse
	9,  // 11: secrets.SecretsService.DeleteSecret:output_type -> secrets.SecretsDeleteResponse
	10, // 12: secrets.SecretsService.GetSecrets:output_type -> secrets.GetSecretsResponse
	11, // 13: secrets.SecretsService.GetSecretsWithData:output_type -> secrets.GetSecretsDataResponse
	7,  // [7:14] is the sub-list for method output_type
	0,  // [0:7] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_pkg_secrets_protocol_proto_init() }
func file_pkg_secrets_protocol_proto_init() {
	if File_pkg_secrets_protocol_proto != nil {
		return
	}
	file_pkg_secrets_secrets_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pkg_secrets_protocol_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_secrets_protocol_proto_goTypes,
		DependencyIndexes: file_pkg_secrets_protocol_proto_depIdxs,
	}.Build()
	File_pkg_secrets_protocol_proto = out.File
	file_pkg_secrets_protocol_proto_rawDesc = nil
	file_pkg_secrets_protocol_proto_goTypes = nil
	file_pkg_secrets_protocol_proto_depIdxs = nil
}