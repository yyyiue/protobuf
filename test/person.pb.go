// Code generated by protoc-gen-go. DO NOT EDIT.
// source: person.proto

package example

import (
	context "context"
	fmt "fmt"
	proto "github.com/yyyiue/protobuf/proto"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type SearchRequest struct {
	Request              string   `protobuf:"bytes,1,opt,name=request,proto3" json:"request,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchRequest) Reset()         { *m = SearchRequest{} }
func (m *SearchRequest) String() string { return proto.CompactTextString(m) }
func (*SearchRequest) ProtoMessage()    {}
func (*SearchRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c9e10cf24b1156d, []int{0}
}

func (m *SearchRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchRequest.Unmarshal(m, b)
}
func (m *SearchRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchRequest.Marshal(b, m, deterministic)
}
func (m *SearchRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchRequest.Merge(m, src)
}
func (m *SearchRequest) XXX_Size() int {
	return xxx_messageInfo_SearchRequest.Size(m)
}
func (m *SearchRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SearchRequest proto.InternalMessageInfo

func (m *SearchRequest) GetRequest() string {
	if m != nil {
		return m.Request
	}
	return ""
}

type SearchResponse struct {
	ID                   string   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Response             string   `protobuf:"bytes,2,opt,name=response,proto3" json:"response,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchResponse) Reset()         { *m = SearchResponse{} }
func (m *SearchResponse) String() string { return proto.CompactTextString(m) }
func (*SearchResponse) ProtoMessage()    {}
func (*SearchResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c9e10cf24b1156d, []int{1}
}

func (m *SearchResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchResponse.Unmarshal(m, b)
}
func (m *SearchResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchResponse.Marshal(b, m, deterministic)
}
func (m *SearchResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchResponse.Merge(m, src)
}
func (m *SearchResponse) XXX_Size() int {
	return xxx_messageInfo_SearchResponse.Size(m)
}
func (m *SearchResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SearchResponse proto.InternalMessageInfo

func (m *SearchResponse) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *SearchResponse) GetResponse() string {
	if m != nil {
		return m.Response
	}
	return ""
}

type SearchResponse_A struct {
	ID                   string   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchResponse_A) Reset()         { *m = SearchResponse_A{} }
func (m *SearchResponse_A) String() string { return proto.CompactTextString(m) }
func (*SearchResponse_A) ProtoMessage()    {}
func (*SearchResponse_A) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c9e10cf24b1156d, []int{1, 0}
}

func (m *SearchResponse_A) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchResponse_A.Unmarshal(m, b)
}
func (m *SearchResponse_A) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchResponse_A.Marshal(b, m, deterministic)
}
func (m *SearchResponse_A) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchResponse_A.Merge(m, src)
}
func (m *SearchResponse_A) XXX_Size() int {
	return xxx_messageInfo_SearchResponse_A.Size(m)
}
func (m *SearchResponse_A) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchResponse_A.DiscardUnknown(m)
}

var xxx_messageInfo_SearchResponse_A proto.InternalMessageInfo

func (m *SearchResponse_A) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

type SearchResponse_A_B struct {
	ID                   string   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchResponse_A_B) Reset()         { *m = SearchResponse_A_B{} }
func (m *SearchResponse_A_B) String() string { return proto.CompactTextString(m) }
func (*SearchResponse_A_B) ProtoMessage()    {}
func (*SearchResponse_A_B) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c9e10cf24b1156d, []int{1, 0, 0}
}

func (m *SearchResponse_A_B) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchResponse_A_B.Unmarshal(m, b)
}
func (m *SearchResponse_A_B) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchResponse_A_B.Marshal(b, m, deterministic)
}
func (m *SearchResponse_A_B) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchResponse_A_B.Merge(m, src)
}
func (m *SearchResponse_A_B) XXX_Size() int {
	return xxx_messageInfo_SearchResponse_A_B.Size(m)
}
func (m *SearchResponse_A_B) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchResponse_A_B.DiscardUnknown(m)
}

var xxx_messageInfo_SearchResponse_A_B proto.InternalMessageInfo

func (m *SearchResponse_A_B) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

type Person struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	MyName               string   `protobuf:"bytes,2,opt,name=my_name,json=myName,proto3" json:"my_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Person) Reset()         { *m = Person{} }
func (m *Person) String() string { return proto.CompactTextString(m) }
func (*Person) ProtoMessage()    {}
func (*Person) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c9e10cf24b1156d, []int{2}
}

func (m *Person) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Person.Unmarshal(m, b)
}
func (m *Person) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Person.Marshal(b, m, deterministic)
}
func (m *Person) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Person.Merge(m, src)
}
func (m *Person) XXX_Size() int {
	return xxx_messageInfo_Person.Size(m)
}
func (m *Person) XXX_DiscardUnknown() {
	xxx_messageInfo_Person.DiscardUnknown(m)
}

var xxx_messageInfo_Person proto.InternalMessageInfo

func (m *Person) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Person) GetMyName() string {
	if m != nil {
		return m.MyName
	}
	return ""
}

type AllPerson struct {
	Persons              []*Person `protobuf:"bytes,1,rep,name=persons,proto3" json:"persons,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *AllPerson) Reset()         { *m = AllPerson{} }
func (m *AllPerson) String() string { return proto.CompactTextString(m) }
func (*AllPerson) ProtoMessage()    {}
func (*AllPerson) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c9e10cf24b1156d, []int{3}
}

func (m *AllPerson) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AllPerson.Unmarshal(m, b)
}
func (m *AllPerson) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AllPerson.Marshal(b, m, deterministic)
}
func (m *AllPerson) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AllPerson.Merge(m, src)
}
func (m *AllPerson) XXX_Size() int {
	return xxx_messageInfo_AllPerson.Size(m)
}
func (m *AllPerson) XXX_DiscardUnknown() {
	xxx_messageInfo_AllPerson.DiscardUnknown(m)
}

var xxx_messageInfo_AllPerson proto.InternalMessageInfo

func (m *AllPerson) GetPersons() []*Person {
	if m != nil {
		return m.Persons
	}
	return nil
}

type Staff struct {
	ID                   string       `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name                 string       `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Age                  int64        `protobuf:"varint,3,opt,name=Age,proto3" json:"Age,omitempty"`
	MyClass              *Staff_Class `protobuf:"bytes,4,opt,name=MyClass,proto3" json:"MyClass,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Staff) Reset()         { *m = Staff{} }
func (m *Staff) String() string { return proto.CompactTextString(m) }
func (*Staff) ProtoMessage()    {}
func (*Staff) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c9e10cf24b1156d, []int{4}
}

func (m *Staff) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Staff.Unmarshal(m, b)
}
func (m *Staff) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Staff.Marshal(b, m, deterministic)
}
func (m *Staff) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Staff.Merge(m, src)
}
func (m *Staff) XXX_Size() int {
	return xxx_messageInfo_Staff.Size(m)
}
func (m *Staff) XXX_DiscardUnknown() {
	xxx_messageInfo_Staff.DiscardUnknown(m)
}

var xxx_messageInfo_Staff proto.InternalMessageInfo

func (m *Staff) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *Staff) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Staff) GetAge() int64 {
	if m != nil {
		return m.Age
	}
	return 0
}

func (m *Staff) GetMyClass() *Staff_Class {
	if m != nil {
		return m.MyClass
	}
	return nil
}

type Staff_Class struct {
	ID                   string   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Type                 string   `protobuf:"bytes,2,opt,name=Type,proto3" json:"Type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Staff_Class) Reset()         { *m = Staff_Class{} }
func (m *Staff_Class) String() string { return proto.CompactTextString(m) }
func (*Staff_Class) ProtoMessage()    {}
func (*Staff_Class) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c9e10cf24b1156d, []int{4, 0}
}

func (m *Staff_Class) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Staff_Class.Unmarshal(m, b)
}
func (m *Staff_Class) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Staff_Class.Marshal(b, m, deterministic)
}
func (m *Staff_Class) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Staff_Class.Merge(m, src)
}
func (m *Staff_Class) XXX_Size() int {
	return xxx_messageInfo_Staff_Class.Size(m)
}
func (m *Staff_Class) XXX_DiscardUnknown() {
	xxx_messageInfo_Staff_Class.DiscardUnknown(m)
}

var xxx_messageInfo_Staff_Class proto.InternalMessageInfo

func (m *Staff_Class) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *Staff_Class) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func init() {
	proto.RegisterType((*SearchRequest)(nil), "example.SearchRequest")
	proto.RegisterType((*SearchResponse)(nil), "example.SearchResponse")
	proto.RegisterType((*SearchResponse_A)(nil), "example.SearchResponse.A")
	proto.RegisterType((*SearchResponse_A_B)(nil), "example.SearchResponse.A.B")
	proto.RegisterType((*Person)(nil), "example.Person")
	proto.RegisterType((*AllPerson)(nil), "example.AllPerson")
	proto.RegisterType((*Staff)(nil), "example.Staff")
	proto.RegisterType((*Staff_Class)(nil), "example.Staff.Class")
}

func init() { proto.RegisterFile("person.proto", fileDescriptor_4c9e10cf24b1156d) }

var fileDescriptor_4c9e10cf24b1156d = []byte{
	// 324 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x51, 0x4d, 0x4f, 0x02, 0x31,
	0x14, 0xb4, 0xbb, 0xb0, 0x2b, 0x0f, 0x45, 0x53, 0x8d, 0x6c, 0xf6, 0xb4, 0x69, 0x3c, 0xac, 0x31,
	0xd9, 0x44, 0x8c, 0x5e, 0x3c, 0x81, 0x5c, 0x48, 0x94, 0x98, 0xe2, 0xdd, 0x54, 0x28, 0xb8, 0xc9,
	0x7e, 0xd4, 0x16, 0x0d, 0x7b, 0xf5, 0x47, 0x78, 0xf0, 0xd7, 0x1a, 0x5a, 0x0a, 0x06, 0xbc, 0x4d,
	0x67, 0x3a, 0x6f, 0xf2, 0xde, 0xc0, 0x81, 0xe0, 0x52, 0x95, 0x45, 0x22, 0x64, 0x39, 0x2f, 0xb1,
	0xcf, 0x17, 0x2c, 0x17, 0x19, 0x0f, 0x1b, 0x33, 0xb5, 0x30, 0x1c, 0xb9, 0x81, 0xc3, 0x11, 0x67,
	0x72, 0xfc, 0x46, 0xf9, 0xfb, 0x07, 0x57, 0x73, 0x7c, 0x0e, 0xbe, 0x34, 0x30, 0x40, 0x11, 0x8a,
	0x1b, 0x3d, 0xf8, 0xfa, 0x8e, 0x3c, 0xc1, 0x8a, 0x74, 0x1c, 0x51, 0x2b, 0x91, 0x29, 0xb4, 0xac,
	0x4d, 0x89, 0xb2, 0x50, 0x1c, 0xb7, 0xc0, 0x19, 0xf4, 0x8d, 0x85, 0x3a, 0x83, 0x3e, 0x0e, 0x61,
	0x5f, 0xae, 0xb4, 0xc0, 0xd1, 0xec, 0xfa, 0x1d, 0xc6, 0x80, 0xba, 0xdb, 0x86, 0xf0, 0x04, 0x50,
	0x6f, 0x9b, 0x24, 0x57, 0xe0, 0x3d, 0xe9, 0x15, 0x96, 0x4a, 0x3a, 0xd1, 0x4a, 0x9d, 0x3a, 0xe9,
	0x04, 0xb7, 0xc1, 0xcf, 0xab, 0x97, 0x82, 0xe5, 0x76, 0xbc, 0x97, 0x57, 0x43, 0x96, 0x73, 0x72,
	0x0b, 0x8d, 0x6e, 0x96, 0xad, 0x5c, 0x17, 0xe0, 0x9b, 0x13, 0xa8, 0x00, 0x45, 0x6e, 0xdc, 0xec,
	0x1c, 0x25, 0xab, 0x23, 0x24, 0xe6, 0x07, 0xb5, 0x3a, 0xf9, 0x41, 0x50, 0x1f, 0xcd, 0xd9, 0x74,
	0xba, 0xb3, 0x0a, 0x86, 0xda, 0x70, 0x93, 0xa3, 0x31, 0x3e, 0x06, 0xb7, 0x3b, 0xe3, 0x81, 0x1b,
	0xa1, 0xd8, 0xa5, 0x4b, 0x88, 0x13, 0xf0, 0x1f, 0xab, 0xfb, 0x8c, 0x29, 0x15, 0xd4, 0x22, 0x14,
	0x37, 0x3b, 0xa7, 0xeb, 0x28, 0x3d, 0x36, 0xd1, 0x1a, 0xb5, 0x9f, 0xc2, 0x4b, 0xa8, 0x6b, 0xf0,
	0x5f, 0xdc, 0x73, 0x25, 0xd6, 0x71, 0x4b, 0xdc, 0x79, 0xb0, 0x35, 0x8d, 0xb8, 0xfc, 0x4c, 0xc7,
	0x1c, 0xdf, 0x81, 0x67, 0x08, 0x7c, 0xb6, 0x89, 0xf9, 0x5b, 0x64, 0xd8, 0xde, 0xe1, 0xcd, 0xf5,
	0xc9, 0xde, 0xab, 0xa7, 0xbb, 0xbf, 0xfe, 0x0d, 0x00, 0x00, 0xff, 0xff, 0xe2, 0x95, 0xc4, 0x0b,
	0x1f, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SearchServiceClient is the client API for SearchService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SearchServiceClient interface {
	Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponse, error)
}

type searchServiceClient struct {
	cc *grpc.ClientConn
}

func NewSearchServiceClient(cc *grpc.ClientConn) SearchServiceClient {
	return &searchServiceClient{cc}
}

func (c *searchServiceClient) Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponse, error) {
	out := new(SearchResponse)
	err := c.cc.Invoke(ctx, "/example.SearchService/Search", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SearchServiceServer is the server API for SearchService service.
type SearchServiceServer interface {
	Search(context.Context, *SearchRequest) (*SearchResponse, error)
}

func RegisterSearchServiceServer(s *grpc.Server, srv SearchServiceServer) {
	s.RegisterService(&_SearchService_serviceDesc, srv)
}

func _SearchService_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchServiceServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/example.SearchService/Search",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchServiceServer).Search(ctx, req.(*SearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SearchService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "example.SearchService",
	HandlerType: (*SearchServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Search",
			Handler:    _SearchService_Search_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "person.proto",
}
