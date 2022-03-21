// Copyright (c) 2022 yeqown
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.3
// source: fieldmask/option.proto

package fieldmask

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type MaskMode int32

const (
	// FiledMaskMode_Filter indicates that the only fields in fieldmask should be kept.
	MaskMode_Filter MaskMode = 0
	// FiledMaskMode_Prune indicates that the field in fieldmask should be removed.
	MaskMode_Prune MaskMode = 1
)

// Enum value maps for MaskMode.
var (
	MaskMode_name = map[int32]string{
		0: "Filter",
		1: "Prune",
	}
	MaskMode_value = map[string]int32{
		"Filter": 0,
		"Prune":  1,
	}
)

func (x MaskMode) Enum() *MaskMode {
	p := new(MaskMode)
	*p = x
	return p
}

func (x MaskMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MaskMode) Descriptor() protoreflect.EnumDescriptor {
	return file_fieldmask_option_proto_enumTypes[0].Descriptor()
}

func (MaskMode) Type() protoreflect.EnumType {
	return &file_fieldmask_option_proto_enumTypes[0]
}

func (x MaskMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MaskMode.Descriptor instead.
func (MaskMode) EnumDescriptor() ([]byte, []int) {
	return file_fieldmask_option_proto_rawDescGZIP(), []int{0}
}

// FieldMask rules applied at the field level
type FieldMask struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	In  *InMessageOption  `protobuf:"bytes,1,opt,name=in,proto3" json:"in,omitempty"`
	Out *OutMessageOption `protobuf:"bytes,2,opt,name=out,proto3" json:"out,omitempty"`
}

func (x *FieldMask) Reset() {
	*x = FieldMask{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fieldmask_option_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FieldMask) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FieldMask) ProtoMessage() {}

func (x *FieldMask) ProtoReflect() protoreflect.Message {
	mi := &file_fieldmask_option_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FieldMask.ProtoReflect.Descriptor instead.
func (*FieldMask) Descriptor() ([]byte, []int) {
	return file_fieldmask_option_proto_rawDescGZIP(), []int{0}
}

func (x *FieldMask) GetIn() *InMessageOption {
	if x != nil {
		return x.In
	}
	return nil
}

func (x *FieldMask) GetOut() *OutMessageOption {
	if x != nil {
		return x.Out
	}
	return nil
}

// FieldMaskMode indicates the options help generate the field mask extension
// for in message.
type InMessageOption struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// gen indicates the Message type that was used to apply FieldMask rules. If gen is
	// set, the fieldmask plugin will generates `MaskIn_FieldName` and `MaskIn_FieldName`
	// for in message.
	Gen bool `protobuf:"varint,1,opt,name=gen,proto3" json:"gen,omitempty"`
}

func (x *InMessageOption) Reset() {
	*x = InMessageOption{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fieldmask_option_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InMessageOption) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InMessageOption) ProtoMessage() {}

func (x *InMessageOption) ProtoReflect() protoreflect.Message {
	mi := &file_fieldmask_option_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InMessageOption.ProtoReflect.Descriptor instead.
func (*InMessageOption) Descriptor() ([]byte, []int) {
	return file_fieldmask_option_proto_rawDescGZIP(), []int{1}
}

func (x *InMessageOption) GetGen() bool {
	if x != nil {
		return x.Gen
	}
	return false
}

// OutMessageOption indicates the options help generating the field mask extension
// for out message.
type OutMessageOption struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// gen indicates the Message type that was used to apply FieldMask rules, if gen is
	// set, the fieldmask plugin will generates `$outMessage_FieldMask` helper object
	// to help deal with $outMessage with $inMessage.FieldMask.
	Gen bool `protobuf:"varint,1,opt,name=gen,proto3" json:"gen,omitempty"`
	// message indicates the Message type that was used to apply FieldMask rules.
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *OutMessageOption) Reset() {
	*x = OutMessageOption{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fieldmask_option_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OutMessageOption) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OutMessageOption) ProtoMessage() {}

func (x *OutMessageOption) ProtoReflect() protoreflect.Message {
	mi := &file_fieldmask_option_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OutMessageOption.ProtoReflect.Descriptor instead.
func (*OutMessageOption) Descriptor() ([]byte, []int) {
	return file_fieldmask_option_proto_rawDescGZIP(), []int{2}
}

func (x *OutMessageOption) GetGen() bool {
	if x != nil {
		return x.Gen
	}
	return false
}

func (x *OutMessageOption) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var file_fieldmask_option_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*FieldMask)(nil),
		Field:         1142,
		Name:          "fieldmask.option.Option",
		Tag:           "bytes,1142,opt,name=Option",
		Filename:      "fieldmask/option.proto",
	},
}

// Extension fields to descriptorpb.FieldOptions.
var (
	// The extension number was applied from https://github.com/protocolbuffers/protobuf/pull/9646.
	//
	// optional fieldmask.option.FieldMask Option = 1142;
	E_Option = &file_fieldmask_option_proto_extTypes[0]
)

var File_fieldmask_option_proto protoreflect.FileDescriptor

var file_fieldmask_option_proto_rawDesc = []byte{
	0x0a, 0x16, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x6d, 0x61, 0x73, 0x6b, 0x2f, 0x6f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x6d,
	0x61, 0x73, 0x6b, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x74, 0x0a, 0x09,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x12, 0x31, 0x0a, 0x02, 0x69, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x6d, 0x61, 0x73,
	0x6b, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x49, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x02, 0x69, 0x6e, 0x12, 0x34, 0x0a, 0x03,
	0x6f, 0x75, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4f, 0x75, 0x74,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x03, 0x6f,
	0x75, 0x74, 0x22, 0x23, 0x0a, 0x0f, 0x49, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x67, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x03, 0x67, 0x65, 0x6e, 0x22, 0x3e, 0x0a, 0x10, 0x4f, 0x75, 0x74, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x67,
	0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x67, 0x65, 0x6e, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2a, 0x21, 0x0a, 0x08, 0x4d, 0x61, 0x73, 0x6b, 0x4d,
	0x6f, 0x64, 0x65, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x10, 0x00, 0x12,
	0x09, 0x0a, 0x05, 0x50, 0x72, 0x75, 0x6e, 0x65, 0x10, 0x01, 0x3a, 0x56, 0x0a, 0x06, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0xf6, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x52, 0x06, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x88,
	0x01, 0x01, 0x42, 0x42, 0x5a, 0x40, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x79, 0x65, 0x71, 0x6f, 0x77, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67,
	0x65, 0x6e, 0x2d, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x6d, 0x61, 0x73, 0x6b, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x6d, 0x61, 0x73, 0x6b, 0x3b, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x6d, 0x61, 0x73, 0x6b, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_fieldmask_option_proto_rawDescOnce sync.Once
	file_fieldmask_option_proto_rawDescData = file_fieldmask_option_proto_rawDesc
)

func file_fieldmask_option_proto_rawDescGZIP() []byte {
	file_fieldmask_option_proto_rawDescOnce.Do(func() {
		file_fieldmask_option_proto_rawDescData = protoimpl.X.CompressGZIP(file_fieldmask_option_proto_rawDescData)
	})
	return file_fieldmask_option_proto_rawDescData
}

var file_fieldmask_option_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_fieldmask_option_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_fieldmask_option_proto_goTypes = []interface{}{
	(MaskMode)(0),                     // 0: fieldmask.option.MaskMode
	(*FieldMask)(nil),                 // 1: fieldmask.option.FieldMask
	(*InMessageOption)(nil),           // 2: fieldmask.option.InMessageOption
	(*OutMessageOption)(nil),          // 3: fieldmask.option.OutMessageOption
	(*descriptorpb.FieldOptions)(nil), // 4: google.protobuf.FieldOptions
}
var file_fieldmask_option_proto_depIdxs = []int32{
	2, // 0: fieldmask.option.FieldMask.in:type_name -> fieldmask.option.InMessageOption
	3, // 1: fieldmask.option.FieldMask.out:type_name -> fieldmask.option.OutMessageOption
	4, // 2: fieldmask.option.Option:extendee -> google.protobuf.FieldOptions
	1, // 3: fieldmask.option.Option:type_name -> fieldmask.option.FieldMask
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	3, // [3:4] is the sub-list for extension type_name
	2, // [2:3] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_fieldmask_option_proto_init() }
func file_fieldmask_option_proto_init() {
	if File_fieldmask_option_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_fieldmask_option_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FieldMask); i {
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
		file_fieldmask_option_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InMessageOption); i {
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
		file_fieldmask_option_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OutMessageOption); i {
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
			RawDescriptor: file_fieldmask_option_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 1,
			NumServices:   0,
		},
		GoTypes:           file_fieldmask_option_proto_goTypes,
		DependencyIndexes: file_fieldmask_option_proto_depIdxs,
		EnumInfos:         file_fieldmask_option_proto_enumTypes,
		MessageInfos:      file_fieldmask_option_proto_msgTypes,
		ExtensionInfos:    file_fieldmask_option_proto_extTypes,
	}.Build()
	File_fieldmask_option_proto = out.File
	file_fieldmask_option_proto_rawDesc = nil
	file_fieldmask_option_proto_goTypes = nil
	file_fieldmask_option_proto_depIdxs = nil
}
