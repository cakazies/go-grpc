// Code generated by protoc-gen-go. DO NOT EDIT.
// source: student.proto

package student

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Student struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Nis                  string   `protobuf:"bytes,3,opt,name=nis,proto3" json:"nis,omitempty"`
	Jk                   string   `protobuf:"bytes,4,opt,name=jk,proto3" json:"jk,omitempty"`
	Nilai                int32    `protobuf:"varint,5,opt,name=nilai,proto3" json:"nilai,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Student) Reset()         { *m = Student{} }
func (m *Student) String() string { return proto.CompactTextString(m) }
func (*Student) ProtoMessage()    {}
func (*Student) Descriptor() ([]byte, []int) {
	return fileDescriptor_94a1c1b032ad0c00, []int{0}
}

func (m *Student) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Student.Unmarshal(m, b)
}
func (m *Student) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Student.Marshal(b, m, deterministic)
}
func (m *Student) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Student.Merge(m, src)
}
func (m *Student) XXX_Size() int {
	return xxx_messageInfo_Student.Size(m)
}
func (m *Student) XXX_DiscardUnknown() {
	xxx_messageInfo_Student.DiscardUnknown(m)
}

var xxx_messageInfo_Student proto.InternalMessageInfo

func (m *Student) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Student) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Student) GetNis() string {
	if m != nil {
		return m.Nis
	}
	return ""
}

func (m *Student) GetJk() string {
	if m != nil {
		return m.Jk
	}
	return ""
}

func (m *Student) GetNilai() int32 {
	if m != nil {
		return m.Nilai
	}
	return 0
}

type Reponse struct {
	Data                 string   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Reponse) Reset()         { *m = Reponse{} }
func (m *Reponse) String() string { return proto.CompactTextString(m) }
func (*Reponse) ProtoMessage()    {}
func (*Reponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_94a1c1b032ad0c00, []int{1}
}

func (m *Reponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Reponse.Unmarshal(m, b)
}
func (m *Reponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Reponse.Marshal(b, m, deterministic)
}
func (m *Reponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Reponse.Merge(m, src)
}
func (m *Reponse) XXX_Size() int {
	return xxx_messageInfo_Reponse.Size(m)
}
func (m *Reponse) XXX_DiscardUnknown() {
	xxx_messageInfo_Reponse.DiscardUnknown(m)
}

var xxx_messageInfo_Reponse proto.InternalMessageInfo

func (m *Reponse) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_94a1c1b032ad0c00, []int{2}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Student)(nil), "student.Student")
	proto.RegisterType((*Reponse)(nil), "student.Reponse")
	proto.RegisterType((*Empty)(nil), "student.Empty")
}

func init() { proto.RegisterFile("student.proto", fileDescriptor_94a1c1b032ad0c00) }

var fileDescriptor_94a1c1b032ad0c00 = []byte{
	// 187 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x2e, 0x29, 0x4d,
	0x49, 0xcd, 0x2b, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x87, 0x72, 0x95, 0x52, 0xb9,
	0xd8, 0x83, 0x21, 0x4c, 0x21, 0x3e, 0x2e, 0xa6, 0xcc, 0x14, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xd6,
	0x20, 0xa6, 0xcc, 0x14, 0x21, 0x21, 0x2e, 0x96, 0xbc, 0xc4, 0xdc, 0x54, 0x09, 0x26, 0x05, 0x46,
	0x0d, 0xce, 0x20, 0x30, 0x5b, 0x48, 0x80, 0x8b, 0x39, 0x2f, 0xb3, 0x58, 0x82, 0x19, 0x2c, 0x04,
	0x62, 0x82, 0x74, 0x65, 0x65, 0x4b, 0xb0, 0x80, 0x05, 0x98, 0xb2, 0xb2, 0x85, 0x44, 0xb8, 0x58,
	0xf3, 0x32, 0x73, 0x12, 0x33, 0x25, 0x58, 0xc1, 0x06, 0x41, 0x38, 0x4a, 0xb2, 0x5c, 0xec, 0x41,
	0xa9, 0x05, 0xf9, 0x79, 0xc5, 0xa9, 0x20, 0x63, 0x53, 0x12, 0x4b, 0x12, 0xc1, 0x16, 0x71, 0x06,
	0x81, 0xd9, 0x4a, 0xec, 0x5c, 0xac, 0xae, 0xb9, 0x05, 0x25, 0x95, 0x46, 0xa6, 0x5c, 0x1c, 0x50,
	0xe7, 0x14, 0x0b, 0x69, 0x72, 0x31, 0xbb, 0xa7, 0x96, 0x08, 0xf1, 0xe9, 0xc1, 0x9c, 0x0e, 0x56,
	0x22, 0x25, 0x00, 0xe7, 0x43, 0x4d, 0x54, 0x62, 0x48, 0x62, 0x03, 0xfb, 0xca, 0x18, 0x10, 0x00,
	0x00, 0xff, 0xff, 0x2a, 0x3b, 0x94, 0x32, 0xe6, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// StudentsClient is the client API for Students service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StudentsClient interface {
	Get(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Reponse, error)
}

type studentsClient struct {
	cc *grpc.ClientConn
}

func NewStudentsClient(cc *grpc.ClientConn) StudentsClient {
	return &studentsClient{cc}
}

func (c *studentsClient) Get(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Reponse, error) {
	out := new(Reponse)
	err := c.cc.Invoke(ctx, "/student.Students/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StudentsServer is the server API for Students service.
type StudentsServer interface {
	Get(context.Context, *Empty) (*Reponse, error)
}

// UnimplementedStudentsServer can be embedded to have forward compatible implementations.
type UnimplementedStudentsServer struct {
}

func (*UnimplementedStudentsServer) Get(ctx context.Context, req *Empty) (*Reponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}

func RegisterStudentsServer(s *grpc.Server, srv StudentsServer) {
	s.RegisterService(&_Students_serviceDesc, srv)
}

func _Students_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudentsServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/student.Students/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudentsServer).Get(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _Students_serviceDesc = grpc.ServiceDesc{
	ServiceName: "student.Students",
	HandlerType: (*StudentsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _Students_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "student.proto",
}
