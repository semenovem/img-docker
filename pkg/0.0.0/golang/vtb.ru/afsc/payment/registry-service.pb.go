// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.11.4
// source: registry-service.proto

package payment

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type AccountReply_Result int32

const (
	AccountReply_OK            AccountReply_Result = 0
	AccountReply_ORG_NOT_FOUND AccountReply_Result = 1 // .....
)

// Enum value maps for AccountReply_Result.
var (
	AccountReply_Result_name = map[int32]string{
		0: "OK",
		1: "ORG_NOT_FOUND",
	}
	AccountReply_Result_value = map[string]int32{
		"OK":            0,
		"ORG_NOT_FOUND": 1,
	}
)

func (x AccountReply_Result) Enum() *AccountReply_Result {
	p := new(AccountReply_Result)
	*p = x
	return p
}

func (x AccountReply_Result) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AccountReply_Result) Descriptor() protoreflect.EnumDescriptor {
	return file_registry_service_proto_enumTypes[0].Descriptor()
}

func (AccountReply_Result) Type() protoreflect.EnumType {
	return &file_registry_service_proto_enumTypes[0]
}

func (x AccountReply_Result) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AccountReply_Result.Descriptor instead.
func (AccountReply_Result) EnumDescriptor() ([]byte, []int) {
	return file_registry_service_proto_rawDescGZIP(), []int{2, 0}
}

type Account struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bic                  string `protobuf:"bytes,1,opt,name=bic,proto3" json:"bic,omitempty"`
	Kpp                  string `protobuf:"bytes,2,opt,name=kpp,proto3" json:"kpp,omitempty"`
	SettlementAccount    string `protobuf:"bytes,3,opt,name=settlement_account,json=settlementAccount,proto3" json:"settlement_account,omitempty"`
	CorrespondentAccount string `protobuf:"bytes,4,opt,name=correspondent_account,json=correspondentAccount,proto3" json:"correspondent_account,omitempty"`
}

func (x *Account) Reset() {
	*x = Account{}
	if protoimpl.UnsafeEnabled {
		mi := &file_registry_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Account) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Account) ProtoMessage() {}

func (x *Account) ProtoReflect() protoreflect.Message {
	mi := &file_registry_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Account.ProtoReflect.Descriptor instead.
func (*Account) Descriptor() ([]byte, []int) {
	return file_registry_service_proto_rawDescGZIP(), []int{0}
}

func (x *Account) GetBic() string {
	if x != nil {
		return x.Bic
	}
	return ""
}

func (x *Account) GetKpp() string {
	if x != nil {
		return x.Kpp
	}
	return ""
}

func (x *Account) GetSettlementAccount() string {
	if x != nil {
		return x.SettlementAccount
	}
	return ""
}

func (x *Account) GetCorrespondentAccount() string {
	if x != nil {
		return x.CorrespondentAccount
	}
	return ""
}

type AccountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Inn     string   `protobuf:"bytes,2,opt,name=inn,proto3" json:"inn,omitempty"`
	Kpp     string   `protobuf:"bytes,3,opt,name=kpp,proto3" json:"kpp,omitempty"`
	Name    string   `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Account *Account `protobuf:"bytes,5,opt,name=account,proto3" json:"account,omitempty"`
}

func (x *AccountRequest) Reset() {
	*x = AccountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_registry_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountRequest) ProtoMessage() {}

func (x *AccountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_registry_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountRequest.ProtoReflect.Descriptor instead.
func (*AccountRequest) Descriptor() ([]byte, []int) {
	return file_registry_service_proto_rawDescGZIP(), []int{1}
}

func (x *AccountRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AccountRequest) GetInn() string {
	if x != nil {
		return x.Inn
	}
	return ""
}

func (x *AccountRequest) GetKpp() string {
	if x != nil {
		return x.Kpp
	}
	return ""
}

func (x *AccountRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AccountRequest) GetAccount() *Account {
	if x != nil {
		return x.Account
	}
	return nil
}

type AccountReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result  AccountReply_Result `protobuf:"varint,1,opt,name=result,proto3,enum=AccountReply_Result" json:"result,omitempty"`
	Message string              `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *AccountReply) Reset() {
	*x = AccountReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_registry_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountReply) ProtoMessage() {}

func (x *AccountReply) ProtoReflect() protoreflect.Message {
	mi := &file_registry_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountReply.ProtoReflect.Descriptor instead.
func (*AccountReply) Descriptor() ([]byte, []int) {
	return file_registry_service_proto_rawDescGZIP(), []int{2}
}

func (x *AccountReply) GetResult() AccountReply_Result {
	if x != nil {
		return x.Result
	}
	return AccountReply_OK
}

func (x *AccountReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_registry_service_proto protoreflect.FileDescriptor

var file_registry_service_proto_rawDesc = []byte{
	0x0a, 0x16, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x91, 0x01, 0x0a, 0x07, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x62, 0x69, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x62, 0x69, 0x63, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x70, 0x70, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x70, 0x70, 0x12, 0x2d, 0x0a, 0x12, 0x73, 0x65, 0x74, 0x74,
	0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x73, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x33, 0x0a, 0x15, 0x63, 0x6f, 0x72, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x64, 0x65, 0x6e, 0x74, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x64, 0x65, 0x6e, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x7c, 0x0a, 0x0e,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10,
	0x0a, 0x03, 0x69, 0x6e, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x69, 0x6e, 0x6e,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x70, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x70, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x7b, 0x0a, 0x0c, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x2c, 0x0a, 0x06, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x22, 0x23, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x06, 0x0a, 0x02,
	0x4f, 0x4b, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d, 0x4f, 0x52, 0x47, 0x5f, 0x4e, 0x4f, 0x54, 0x5f,
	0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x01, 0x32, 0x69, 0x0a, 0x08, 0x52, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x72, 0x79, 0x12, 0x2c, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x0f, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x2f, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x0f, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x42, 0x35, 0x0a, 0x19, 0x72, 0x75, 0x2e, 0x76, 0x74, 0x62, 0x2e, 0x67, 0x72, 0x70,
	0x63, 0x2e, 0x6d, 0x62, 0x61, 0x6e, 0x6b, 0x2e, 0x61, 0x64, 0x61, 0x70, 0x74, 0x65, 0x72, 0x50,
	0x01, 0x5a, 0x13, 0x76, 0x74, 0x62, 0x2e, 0x72, 0x75, 0x2f, 0x61, 0x66, 0x73, 0x63, 0x2f, 0x70,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x90, 0x01, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_registry_service_proto_rawDescOnce sync.Once
	file_registry_service_proto_rawDescData = file_registry_service_proto_rawDesc
)

func file_registry_service_proto_rawDescGZIP() []byte {
	file_registry_service_proto_rawDescOnce.Do(func() {
		file_registry_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_registry_service_proto_rawDescData)
	})
	return file_registry_service_proto_rawDescData
}

var file_registry_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_registry_service_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_registry_service_proto_goTypes = []interface{}{
	(AccountReply_Result)(0), // 0: AccountReply.Result
	(*Account)(nil),          // 1: Account
	(*AccountRequest)(nil),   // 2: AccountRequest
	(*AccountReply)(nil),     // 3: AccountReply
}
var file_registry_service_proto_depIdxs = []int32{
	1, // 0: AccountRequest.account:type_name -> Account
	0, // 1: AccountReply.result:type_name -> AccountReply.Result
	2, // 2: Registry.GetAccount:input_type -> AccountRequest
	2, // 3: Registry.DeleteAccount:input_type -> AccountRequest
	3, // 4: Registry.GetAccount:output_type -> AccountReply
	3, // 5: Registry.DeleteAccount:output_type -> AccountReply
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_registry_service_proto_init() }
func file_registry_service_proto_init() {
	if File_registry_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_registry_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Account); i {
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
		file_registry_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountRequest); i {
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
		file_registry_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountReply); i {
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
			RawDescriptor: file_registry_service_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_registry_service_proto_goTypes,
		DependencyIndexes: file_registry_service_proto_depIdxs,
		EnumInfos:         file_registry_service_proto_enumTypes,
		MessageInfos:      file_registry_service_proto_msgTypes,
	}.Build()
	File_registry_service_proto = out.File
	file_registry_service_proto_rawDesc = nil
	file_registry_service_proto_goTypes = nil
	file_registry_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// RegistryClient is the client API for Registry service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RegistryClient interface {
	GetAccount(ctx context.Context, in *AccountRequest, opts ...grpc.CallOption) (*AccountReply, error)
	DeleteAccount(ctx context.Context, in *AccountRequest, opts ...grpc.CallOption) (*AccountReply, error)
}

type registryClient struct {
	cc grpc.ClientConnInterface
}

func NewRegistryClient(cc grpc.ClientConnInterface) RegistryClient {
	return &registryClient{cc}
}

func (c *registryClient) GetAccount(ctx context.Context, in *AccountRequest, opts ...grpc.CallOption) (*AccountReply, error) {
	out := new(AccountReply)
	err := c.cc.Invoke(ctx, "/Registry/GetAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registryClient) DeleteAccount(ctx context.Context, in *AccountRequest, opts ...grpc.CallOption) (*AccountReply, error) {
	out := new(AccountReply)
	err := c.cc.Invoke(ctx, "/Registry/DeleteAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RegistryServer is the server API for Registry service.
type RegistryServer interface {
	GetAccount(context.Context, *AccountRequest) (*AccountReply, error)
	DeleteAccount(context.Context, *AccountRequest) (*AccountReply, error)
}

// UnimplementedRegistryServer can be embedded to have forward compatible implementations.
type UnimplementedRegistryServer struct {
}

func (*UnimplementedRegistryServer) GetAccount(context.Context, *AccountRequest) (*AccountReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccount not implemented")
}
func (*UnimplementedRegistryServer) DeleteAccount(context.Context, *AccountRequest) (*AccountReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAccount not implemented")
}

func RegisterRegistryServer(s *grpc.Server, srv RegistryServer) {
	s.RegisterService(&_Registry_serviceDesc, srv)
}

func _Registry_GetAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryServer).GetAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Registry/GetAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryServer).GetAccount(ctx, req.(*AccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Registry_DeleteAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryServer).DeleteAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Registry/DeleteAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryServer).DeleteAccount(ctx, req.(*AccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Registry_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Registry",
	HandlerType: (*RegistryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAccount",
			Handler:    _Registry_GetAccount_Handler,
		},
		{
			MethodName: "DeleteAccount",
			Handler:    _Registry_DeleteAccount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "registry-service.proto",
}
