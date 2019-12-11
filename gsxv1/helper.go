package gsxv1

import (
	proto "github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
)

//func IsEmbed(field *descriptor.FieldDescriptorProto) bool {
//	return proto.GetBoolExtension(field.Options, E_Embed, false)
//}
//
//func IsNullable(field *descriptor.FieldDescriptorProto) bool {
//	return proto.GetBoolExtension(field.Options, E_Nullable, true)
//}
//
//func IsStdTime(field *descriptor.FieldDescriptorProto) bool {
//	return proto.GetBoolExtension(field.Options, E_Stdtime, false)
//}
//
//func IsStdDuration(field *descriptor.FieldDescriptorProto) bool {
//	return proto.GetBoolExtension(field.Options, E_Stdduration, false)
//}
//
//func IsStdDouble(field *descriptor.FieldDescriptorProto) bool {
//	return proto.GetBoolExtension(field.Options, E_Wktpointer, false) && *field.TypeName == ".google.protobuf.DoubleValue"
//}
//
//func IsStdFloat(field *descriptor.FieldDescriptorProto) bool {
//	return proto.GetBoolExtension(field.Options, E_Wktpointer, false) && *field.TypeName == ".google.protobuf.FloatValue"
//}
//
//func IsStdInt64(field *descriptor.FieldDescriptorProto) bool {
//	return proto.GetBoolExtension(field.Options, E_Wktpointer, false) && *field.TypeName == ".google.protobuf.Int64Value"
//}
//
//func IsStdUInt64(field *descriptor.FieldDescriptorProto) bool {
//	return proto.GetBoolExtension(field.Options, E_Wktpointer, false) && *field.TypeName == ".google.protobuf.UInt64Value"
//}
//
//func IsStdInt32(field *descriptor.FieldDescriptorProto) bool {
//	return proto.GetBoolExtension(field.Options, E_Wktpointer, false) && *field.TypeName == ".google.protobuf.Int32Value"
//}
//
//func IsStdUInt32(field *descriptor.FieldDescriptorProto) bool {
//	return proto.GetBoolExtension(field.Options, E_Wktpointer, false) && *field.TypeName == ".google.protobuf.UInt32Value"
//}
//
//func IsStdBool(field *descriptor.FieldDescriptorProto) bool {
//	return proto.GetBoolExtension(field.Options, E_Wktpointer, false) && *field.TypeName == ".google.protobuf.BoolValue"
//}
//
//func IsStdString(field *descriptor.FieldDescriptorProto) bool {
//	return proto.GetBoolExtension(field.Options, E_Wktpointer, false) && *field.TypeName == ".google.protobuf.StringValue"
//}
//
//func IsStdBytes(field *descriptor.FieldDescriptorProto) bool {
//	return proto.GetBoolExtension(field.Options, E_Wktpointer, false) && *field.TypeName == ".google.protobuf.BytesValue"
//}
//
//func IsStdType(field *descriptor.FieldDescriptorProto) bool {
//	return (IsStdTime(field) || IsStdDuration(field) ||
//		IsStdDouble(field) || IsStdFloat(field) ||
//		IsStdInt64(field) || IsStdUInt64(field) ||
//		IsStdInt32(field) || IsStdUInt32(field) ||
//		IsStdBool(field) ||
//		IsStdString(field) || IsStdBytes(field))
//}
//
//func IsWktPtr(field *descriptor.FieldDescriptorProto) bool {
//	return proto.GetBoolExtension(field.Options, E_Wktpointer, false)
//}
//
//func NeedsNilCheck(proto3 bool, field *descriptor.FieldDescriptorProto) bool {
//	nullable := IsNullable(field)
//	if field.IsMessage() || IsCustomType(field) {
//		return nullable
//	}
//	if proto3 {
//		return false
//	}
//	return nullable || *field.Type == descriptor.FieldDescriptorProto_TYPE_BYTES
//}
//
//func IsCustomType(field *descriptor.FieldDescriptorProto) bool {
//	typ := GetCustomType(field)
//	if len(typ) > 0 {
//		return true
//	}
//	return false
//}
//
//func IsCastType(field *descriptor.FieldDescriptorProto) bool {
//	typ := GetCastType(field)
//	if len(typ) > 0 {
//		return true
//	}
//	return false
//}
//
//func IsCastKey(field *descriptor.FieldDescriptorProto) bool {
//	typ := GetCastKey(field)
//	if len(typ) > 0 {
//		return true
//	}
//	return false
//}
//
//func IsCastValue(field *descriptor.FieldDescriptorProto) bool {
//	typ := GetCastValue(field)
//	if len(typ) > 0 {
//		return true
//	}
//	return false
//}
//
//func HasEnumDecl(file *descriptor.FileDescriptorProto, enum *descriptor.EnumDescriptorProto) bool {
//	return proto.GetBoolExtension(enum.Options, E_Enumdecl, proto.GetBoolExtension(file.Options, E_EnumdeclAll, true))
//}
//
//func HasTypeDecl(file *descriptor.FileDescriptorProto, message *descriptor.DescriptorProto) bool {
//	return proto.GetBoolExtension(message.Options, E_Typedecl, proto.GetBoolExtension(file.Options, E_TypedeclAll, true))
//}

//func GetCustomType(field *descriptor.FieldDescriptorProto) string {
//	if field == nil {
//		return ""
//	}
//	if field.Options != nil {
//		v, err := proto.GetExtension(field.Options, E_Customtype)
//		if err == nil && v.(*string) != nil {
//			return *(v.(*string))
//		}
//	}
//	return ""
//}
//
//func GetCastType(field *descriptor.FieldDescriptorProto) string {
//	if field == nil {
//		return ""
//	}
//	if field.Options != nil {
//		v, err := proto.GetExtension(field.Options, E_Casttype)
//		if err == nil && v.(*string) != nil {
//			return *(v.(*string))
//		}
//	}
//	return ""
//}
//
//func GetCastKey(field *descriptor.FieldDescriptorProto) string {
//	if field == nil {
//		return ""
//	}
//	if field.Options != nil {
//		v, err := proto.GetExtension(field.Options, E_Castkey)
//		if err == nil && v.(*string) != nil {
//			return *(v.(*string))
//		}
//	}
//	return ""
//}
//
//func GetCastValue(field *descriptor.FieldDescriptorProto) string {
//	if field == nil {
//		return ""
//	}
//	if field.Options != nil {
//		v, err := proto.GetExtension(field.Options, E_Castvalue)
//		if err == nil && v.(*string) != nil {
//			return *(v.(*string))
//		}
//	}
//	return ""
//}
//
//func IsCustomName(field *descriptor.FieldDescriptorProto) bool {
//	name := GetCustomName(field)
//	if len(name) > 0 {
//		return true
//	}
//	return false
//}
//
//func IsEnumCustomName(field *descriptor.EnumDescriptorProto) bool {
//	name := GetEnumCustomName(field)
//	if len(name) > 0 {
//		return true
//	}
//	return false
//}
//
//func IsEnumValueCustomName(field *descriptor.EnumValueDescriptorProto) bool {
//	name := GetEnumValueCustomName(field)
//	if len(name) > 0 {
//		return true
//	}
//	return false
//}
//
//func GetCustomName(field *descriptor.FieldDescriptorProto) string {
//	if field == nil {
//		return ""
//	}
//	if field.Options != nil {
//		v, err := proto.GetExtension(field.Options, E_Customname)
//		if err == nil && v.(*string) != nil {
//			return *(v.(*string))
//		}
//	}
//	return ""
//}
//
//func GetEnumCustomName(field *descriptor.EnumDescriptorProto) string {
//	if field == nil {
//		return ""
//	}
//	if field.Options != nil {
//		v, err := proto.GetExtension(field.Options, E_EnumCustomname)
//		if err == nil && v.(*string) != nil {
//			return *(v.(*string))
//		}
//	}
//	return ""
//}
//
//func GetEnumValueCustomName(field *descriptor.EnumValueDescriptorProto) string {
//	if field == nil {
//		return ""
//	}
//	if field.Options != nil {
//		v, err := proto.GetExtension(field.Options, E_EnumvalueCustomname)
//		if err == nil && v.(*string) != nil {
//			return *(v.(*string))
//		}
//	}
//	return ""
//}

func GetJsonTag(field *descriptor.FieldDescriptorProto) *string {
	if field == nil {
		return nil
	}
	if field.Options != nil {
		v, err := proto.GetExtension(field.Options, E_Jsontag)
		if err == nil && v.(*string) != nil {
			return v.(*string)
		}
	}
	return nil
}

//func GetMoreTags(field *descriptor.FieldDescriptorProto) *string {
//	if field == nil {
//		return nil
//	}
//	if field.Options != nil {
//		v, err := proto.GetExtension(field.Options, E_Moretags)
//		if err == nil && v.(*string) != nil {
//			return (v.(*string))
//		}
//	}
//	return nil
//}
//
//type EnableFunc func(file *descriptor.FileDescriptorProto, message *descriptor.DescriptorProto) bool
//
//func EnabledGoEnumPrefix(file *descriptor.FileDescriptorProto, enum *descriptor.EnumDescriptorProto) bool {
//	return proto.GetBoolExtension(enum.Options, E_GoprotoEnumPrefix, proto.GetBoolExtension(file.Options, E_GoprotoEnumPrefixAll, true))
//}
//
//func EnabledGoStringer(file *descriptor.FileDescriptorProto, message *descriptor.DescriptorProto) bool {
//	return proto.GetBoolExtension(message.Options, E_GoprotoStringer, proto.GetBoolExtension(file.Options, E_GoprotoStringerAll, true))
//}
//
//func HasGoGetters(file *descriptor.FileDescriptorProto, message *descriptor.DescriptorProto) bool {
//	return proto.GetBoolExtension(message.Options, E_GoprotoGetters, proto.GetBoolExtension(file.Options, E_GoprotoGettersAll, true))
//}
//
//func IsUnion(file *descriptor.FileDescriptorProto, message *descriptor.DescriptorProto) bool {
//	return proto.GetBoolExtension(message.Options, E_Onlyone, proto.GetBoolExtension(file.Options, E_OnlyoneAll, false))
//}
//
//func HasGoString(file *descriptor.FileDescriptorProto, message *descriptor.DescriptorProto) bool {
//	return proto.GetBoolExtension(message.Options, E_Gostring, proto.GetBoolExtension(file.Options, E_GostringAll, false))
//}
//
//func HasEqual(file *descriptor.FileDescriptorProto, message *descriptor.DescriptorProto) bool {
//	return proto.GetBoolExtension(message.Options, E_Equal, proto.GetBoolExtension(file.Options, E_EqualAll, false))
//}
//
//func HasVerboseEqual(file *descriptor.FileDescriptorProto, message *descriptor.DescriptorProto) bool {
//	return proto.GetBoolExtension(message.Options, E_VerboseEqual, proto.GetBoolExtension(file.Options, E_VerboseEqualAll, false))
//}
//
//func IsStringer(file *descriptor.FileDescriptorProto, message *descriptor.DescriptorProto) bool {
//	return proto.GetBoolExtension(message.Options, E_Stringer, proto.GetBoolExtension(file.Options, E_StringerAll, false))
//}
//
//func IsFace(file *descriptor.FileDescriptorProto, message *descriptor.DescriptorProto) bool {
//	return proto.GetBoolExtension(message.Options, E_Face, proto.GetBoolExtension(file.Options, E_FaceAll, false))
//}
//
//func HasDescription(file *descriptor.FileDescriptorProto, message *descriptor.DescriptorProto) bool {
//	return proto.GetBoolExtension(message.Options, E_Description, proto.GetBoolExtension(file.Options, E_DescriptionAll, false))
//}
//
//func HasPopulate(file *descriptor.FileDescriptorProto, message *descriptor.DescriptorProto) bool {
//	return proto.GetBoolExtension(message.Options, E_Populate, proto.GetBoolExtension(file.Options, E_PopulateAll, false))
//}
//
//func HasTestGen(file *descriptor.FileDescriptorProto, message *descriptor.DescriptorProto) bool {
//	return proto.GetBoolExtension(message.Options, E_Testgen, proto.GetBoolExtension(file.Options, E_TestgenAll, false))
//}
//
//func HasBenchGen(file *descriptor.FileDescriptorProto, message *descriptor.DescriptorProto) bool {
//	return proto.GetBoolExtension(message.Options, E_Benchgen, proto.GetBoolExtension(file.Options, E_BenchgenAll, false))
//}
//
//func IsMarshaler(file *descriptor.FileDescriptorProto, message *descriptor.DescriptorProto) bool {
//	return proto.GetBoolExtension(message.Options, E_Marshaler, proto.GetBoolExtension(file.Options, E_MarshalerAll, false))
//}
//
//func IsUnmarshaler(file *descriptor.FileDescriptorProto, message *descriptor.DescriptorProto) bool {
//	return proto.GetBoolExtension(message.Options, E_Unmarshaler, proto.GetBoolExtension(file.Options, E_UnmarshalerAll, false))
//}
//
//func IsStableMarshaler(file *descriptor.FileDescriptorProto, message *descriptor.DescriptorProto) bool {
//	return proto.GetBoolExtension(message.Options, E_StableMarshaler, proto.GetBoolExtension(file.Options, E_StableMarshalerAll, false))
//}
//
//func IsSizer(file *descriptor.FileDescriptorProto, message *descriptor.DescriptorProto) bool {
//	return proto.GetBoolExtension(message.Options, E_Sizer, proto.GetBoolExtension(file.Options, E_SizerAll, false))
//}
//
//func IsProtoSizer(file *descriptor.FileDescriptorProto, message *descriptor.DescriptorProto) bool {
//	return proto.GetBoolExtension(message.Options, E_Protosizer, proto.GetBoolExtension(file.Options, E_ProtosizerAll, false))
//}
//
//func IsGoEnumStringer(file *descriptor.FileDescriptorProto, enum *descriptor.EnumDescriptorProto) bool {
//	return proto.GetBoolExtension(enum.Options, E_GoprotoEnumStringer, proto.GetBoolExtension(file.Options, E_GoprotoEnumStringerAll, true))
//}
//
//func IsEnumStringer(file *descriptor.FileDescriptorProto, enum *descriptor.EnumDescriptorProto) bool {
//	return proto.GetBoolExtension(enum.Options, E_EnumStringer, proto.GetBoolExtension(file.Options, E_EnumStringerAll, false))
//}
//
//func IsUnsafeMarshaler(file *descriptor.FileDescriptorProto, message *descriptor.DescriptorProto) bool {
//	return proto.GetBoolExtension(message.Options, E_UnsafeMarshaler, proto.GetBoolExtension(file.Options, E_UnsafeMarshalerAll, false))
//}
//
//func IsUnsafeUnmarshaler(file *descriptor.FileDescriptorProto, message *descriptor.DescriptorProto) bool {
//	return proto.GetBoolExtension(message.Options, E_UnsafeUnmarshaler, proto.GetBoolExtension(file.Options, E_UnsafeUnmarshalerAll, false))
//}
//
//func HasExtensionsMap(file *descriptor.FileDescriptorProto, message *descriptor.DescriptorProto) bool {
//	return proto.GetBoolExtension(message.Options, E_GoprotoExtensionsMap, proto.GetBoolExtension(file.Options, E_GoprotoExtensionsMapAll, true))
//}
//
//func HasUnrecognized(file *descriptor.FileDescriptorProto, message *descriptor.DescriptorProto) bool {
//	return proto.GetBoolExtension(message.Options, E_GoprotoUnrecognized, proto.GetBoolExtension(file.Options, E_GoprotoUnrecognizedAll, true))
//}
//
//func IsProto3(file *descriptor.FileDescriptorProto) bool {
//	return file.GetSyntax() == "proto3"
//}
//
//func ImportsGoGoProto(file *descriptor.FileDescriptorProto) bool {
//	return proto.GetBoolExtension(file.Options, E_GogoprotoImport, true)
//}
//
//func HasCompare(file *descriptor.FileDescriptorProto, message *descriptor.DescriptorProto) bool {
//	return proto.GetBoolExtension(message.Options, E_Compare, proto.GetBoolExtension(file.Options, E_CompareAll, false))
//}
//
//func RegistersGolangProto(file *descriptor.FileDescriptorProto) bool {
//	return proto.GetBoolExtension(file.Options, E_GoprotoRegistration, false)
//}
//
//func HasMessageName(file *descriptor.FileDescriptorProto, message *descriptor.DescriptorProto) bool {
//	return proto.GetBoolExtension(message.Options, E_Messagename, proto.GetBoolExtension(file.Options, E_MessagenameAll, false))
//}
//
//func HasSizecache(file *descriptor.FileDescriptorProto, message *descriptor.DescriptorProto) bool {
//	return proto.GetBoolExtension(message.Options, E_GoprotoSizecache, proto.GetBoolExtension(file.Options, E_GoprotoSizecacheAll, true))
//}
//
//func HasUnkeyed(file *descriptor.FileDescriptorProto, message *descriptor.DescriptorProto) bool {
//	return proto.GetBoolExtension(message.Options, E_GoprotoUnkeyed, proto.GetBoolExtension(file.Options, E_GoprotoUnkeyedAll, true))
//}
