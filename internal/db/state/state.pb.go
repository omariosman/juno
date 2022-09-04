// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.6.1
// source: state.proto

package state

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

type Code struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code [][]byte `protobuf:"bytes,1,rep,name=code,proto3" json:"code,omitempty"`
}

func (x *Code) Reset() {
	*x = Code{}
	if protoimpl.UnsafeEnabled {
		mi := &file_state_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Code) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Code) ProtoMessage() {}

func (x *Code) ProtoReflect() protoreflect.Message {
	mi := &file_state_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Code.ProtoReflect.Descriptor instead.
func (*Code) Descriptor() ([]byte, []int) {
	return file_state_proto_rawDescGZIP(), []int{0}
}

func (x *Code) GetCode() [][]byte {
	if x != nil {
		return x.Code
	}
	return nil
}

type CodeDefinition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Definition string `protobuf:"bytes,1,opt,name=definition,proto3" json:"definition,omitempty"`
}

func (x *CodeDefinition) Reset() {
	*x = CodeDefinition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_state_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CodeDefinition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CodeDefinition) ProtoMessage() {}

func (x *CodeDefinition) ProtoReflect() protoreflect.Message {
	mi := &file_state_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CodeDefinition.ProtoReflect.Descriptor instead.
func (*CodeDefinition) Descriptor() ([]byte, []int) {
	return file_state_proto_rawDescGZIP(), []int{1}
}

func (x *CodeDefinition) GetDefinition() string {
	if x != nil {
		return x.Definition
	}
	return ""
}

type ContractState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ContractHash []byte `protobuf:"bytes,1,opt,name=contractHash,proto3" json:"contractHash,omitempty"`
	StorageRoot  []byte `protobuf:"bytes,2,opt,name=storageRoot,proto3" json:"storageRoot,omitempty"`
	Nonce        []byte `protobuf:"bytes,3,opt,name=nonce,proto3" json:"nonce,omitempty"`
}

func (x *ContractState) Reset() {
	*x = ContractState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_state_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContractState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContractState) ProtoMessage() {}

func (x *ContractState) ProtoReflect() protoreflect.Message {
	mi := &file_state_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContractState.ProtoReflect.Descriptor instead.
func (*ContractState) Descriptor() ([]byte, []int) {
	return file_state_proto_rawDescGZIP(), []int{2}
}

func (x *ContractState) GetContractHash() []byte {
	if x != nil {
		return x.ContractHash
	}
	return nil
}

func (x *ContractState) GetStorageRoot() []byte {
	if x != nil {
		return x.StorageRoot
	}
	return nil
}

func (x *ContractState) GetNonce() []byte {
	if x != nil {
		return x.Nonce
	}
	return nil
}

type TrieNode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Node:
	//
	//	*TrieNode_BinaryNode
	//	*TrieNode_EdgeNode
	Node isTrieNode_Node `protobuf_oneof:"node"`
}

func (x *TrieNode) Reset() {
	*x = TrieNode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_state_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TrieNode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TrieNode) ProtoMessage() {}

func (x *TrieNode) ProtoReflect() protoreflect.Message {
	mi := &file_state_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TrieNode.ProtoReflect.Descriptor instead.
func (*TrieNode) Descriptor() ([]byte, []int) {
	return file_state_proto_rawDescGZIP(), []int{3}
}

func (m *TrieNode) GetNode() isTrieNode_Node {
	if m != nil {
		return m.Node
	}
	return nil
}

func (x *TrieNode) GetBinaryNode() *BinaryNode {
	if x, ok := x.GetNode().(*TrieNode_BinaryNode); ok {
		return x.BinaryNode
	}
	return nil
}

func (x *TrieNode) GetEdgeNode() *EdgeNode {
	if x, ok := x.GetNode().(*TrieNode_EdgeNode); ok {
		return x.EdgeNode
	}
	return nil
}

type isTrieNode_Node interface {
	isTrieNode_Node()
}

type TrieNode_BinaryNode struct {
	BinaryNode *BinaryNode `protobuf:"bytes,1,opt,name=binary_node,json=binaryNode,proto3,oneof"`
}

type TrieNode_EdgeNode struct {
	EdgeNode *EdgeNode `protobuf:"bytes,2,opt,name=edge_node,json=edgeNode,proto3,oneof"`
}

func (*TrieNode_BinaryNode) isTrieNode_Node() {}

func (*TrieNode_EdgeNode) isTrieNode_Node() {}

type BinaryNode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LeftH  []byte `protobuf:"bytes,1,opt,name=leftH,proto3" json:"leftH,omitempty"`
	RightH []byte `protobuf:"bytes,2,opt,name=rightH,proto3" json:"rightH,omitempty"`
}

func (x *BinaryNode) Reset() {
	*x = BinaryNode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_state_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BinaryNode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BinaryNode) ProtoMessage() {}

func (x *BinaryNode) ProtoReflect() protoreflect.Message {
	mi := &file_state_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BinaryNode.ProtoReflect.Descriptor instead.
func (*BinaryNode) Descriptor() ([]byte, []int) {
	return file_state_proto_rawDescGZIP(), []int{4}
}

func (x *BinaryNode) GetLeftH() []byte {
	if x != nil {
		return x.LeftH
	}
	return nil
}

func (x *BinaryNode) GetRightH() []byte {
	if x != nil {
		return x.RightH
	}
	return nil
}

type EdgeNode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Length uint32 `protobuf:"varint,1,opt,name=length,proto3" json:"length,omitempty"`
	Path   []byte `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	Bottom []byte `protobuf:"bytes,3,opt,name=bottom,proto3" json:"bottom,omitempty"`
}

func (x *EdgeNode) Reset() {
	*x = EdgeNode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_state_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EdgeNode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EdgeNode) ProtoMessage() {}

func (x *EdgeNode) ProtoReflect() protoreflect.Message {
	mi := &file_state_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EdgeNode.ProtoReflect.Descriptor instead.
func (*EdgeNode) Descriptor() ([]byte, []int) {
	return file_state_proto_rawDescGZIP(), []int{5}
}

func (x *EdgeNode) GetLength() uint32 {
	if x != nil {
		return x.Length
	}
	return 0
}

func (x *EdgeNode) GetPath() []byte {
	if x != nil {
		return x.Path
	}
	return nil
}

func (x *EdgeNode) GetBottom() []byte {
	if x != nil {
		return x.Bottom
	}
	return nil
}

var File_state_proto protoreflect.FileDescriptor

var file_state_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1a, 0x0a,
	0x04, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0c, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x30, 0x0a, 0x0e, 0x43, 0x6f, 0x64,
	0x65, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x64,
	0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x6b, 0x0a, 0x0d, 0x43,
	0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x22, 0x0a, 0x0c,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x48, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x48, 0x61, 0x73, 0x68,
	0x12, 0x20, 0x0a, 0x0b, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x52, 0x6f, 0x6f, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x52, 0x6f,
	0x6f, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x22, 0x6c, 0x0a, 0x08, 0x54, 0x72, 0x69, 0x65,
	0x4e, 0x6f, 0x64, 0x65, 0x12, 0x2e, 0x0a, 0x0b, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x5f, 0x6e,
	0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x42, 0x69, 0x6e, 0x61,
	0x72, 0x79, 0x4e, 0x6f, 0x64, 0x65, 0x48, 0x00, 0x52, 0x0a, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x79,
	0x4e, 0x6f, 0x64, 0x65, 0x12, 0x28, 0x0a, 0x09, 0x65, 0x64, 0x67, 0x65, 0x5f, 0x6e, 0x6f, 0x64,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x45, 0x64, 0x67, 0x65, 0x4e, 0x6f,
	0x64, 0x65, 0x48, 0x00, 0x52, 0x08, 0x65, 0x64, 0x67, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x42, 0x06,
	0x0a, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x22, 0x3a, 0x0a, 0x0a, 0x42, 0x69, 0x6e, 0x61, 0x72, 0x79,
	0x4e, 0x6f, 0x64, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65, 0x66, 0x74, 0x48, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x05, 0x6c, 0x65, 0x66, 0x74, 0x48, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x69,
	0x67, 0x68, 0x74, 0x48, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x72, 0x69, 0x67, 0x68,
	0x74, 0x48, 0x22, 0x4e, 0x0a, 0x08, 0x45, 0x64, 0x67, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06,
	0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x6f,
	0x74, 0x74, 0x6f, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x62, 0x6f, 0x74, 0x74,
	0x6f, 0x6d, 0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x4e, 0x65, 0x74, 0x68, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x64, 0x45, 0x74, 0x68, 0x2f, 0x6a,
	0x75, 0x6e, 0x6f, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_state_proto_rawDescOnce sync.Once
	file_state_proto_rawDescData = file_state_proto_rawDesc
)

func file_state_proto_rawDescGZIP() []byte {
	file_state_proto_rawDescOnce.Do(func() {
		file_state_proto_rawDescData = protoimpl.X.CompressGZIP(file_state_proto_rawDescData)
	})
	return file_state_proto_rawDescData
}

var file_state_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_state_proto_goTypes = []interface{}{
	(*Code)(nil),           // 0: Code
	(*CodeDefinition)(nil), // 1: CodeDefinition
	(*ContractState)(nil),  // 2: ContractState
	(*TrieNode)(nil),       // 3: TrieNode
	(*BinaryNode)(nil),     // 4: BinaryNode
	(*EdgeNode)(nil),       // 5: EdgeNode
}
var file_state_proto_depIdxs = []int32{
	4, // 0: TrieNode.binary_node:type_name -> BinaryNode
	5, // 1: TrieNode.edge_node:type_name -> EdgeNode
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_state_proto_init() }
func file_state_proto_init() {
	if File_state_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_state_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Code); i {
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
		file_state_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CodeDefinition); i {
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
		file_state_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContractState); i {
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
		file_state_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TrieNode); i {
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
		file_state_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BinaryNode); i {
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
		file_state_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EdgeNode); i {
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
	file_state_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*TrieNode_BinaryNode)(nil),
		(*TrieNode_EdgeNode)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_state_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_state_proto_goTypes,
		DependencyIndexes: file_state_proto_depIdxs,
		MessageInfos:      file_state_proto_msgTypes,
	}.Build()
	File_state_proto = out.File
	file_state_proto_rawDesc = nil
	file_state_proto_goTypes = nil
	file_state_proto_depIdxs = nil
}
