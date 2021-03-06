// Code generated by protoc-gen-go.
// source: file.proto
// DO NOT EDIT!

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	file.proto

It has these top-level messages:
	PathError
	SyscallError
	Error
	WriteRequest
	WriteReply
	ReadRequest
	ReadReply
*/
package proto

import proto1 "github.com/golang/protobuf/proto"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal

// PathError records an error and the operation and file path that caused it.
type PathError struct {
	Op    string `protobuf:"bytes,1,opt,name=op" json:"op,omitempty"`
	Path  string `protobuf:"bytes,2,opt,name=path" json:"path,omitempty"`
	Error string `protobuf:"bytes,3,opt,name=error" json:"error,omitempty"`
}

func (m *PathError) Reset()         { *m = PathError{} }
func (m *PathError) String() string { return proto1.CompactTextString(m) }
func (*PathError) ProtoMessage()    {}

// SyscallError records an error from a specific system call.
type SyscallError struct {
	Syscall string `protobuf:"bytes,1,opt,name=syscall" json:"syscall,omitempty"`
	Error   string `protobuf:"bytes,2,opt,name=error" json:"error,omitempty"`
}

func (m *SyscallError) Reset()         { *m = SyscallError{} }
func (m *SyscallError) String() string { return proto1.CompactTextString(m) }
func (*SyscallError) ProtoMessage()    {}

type Error struct {
	PathErr *PathError    `protobuf:"bytes,1,opt,name=pathErr" json:"pathErr,omitempty"`
	SysErr  *SyscallError `protobuf:"bytes,2,opt,name=sysErr" json:"sysErr,omitempty"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto1.CompactTextString(m) }
func (*Error) ProtoMessage()    {}

func (m *Error) GetPathErr() *PathError {
	if m != nil {
		return m.PathErr
	}
	return nil
}

func (m *Error) GetSysErr() *SyscallError {
	if m != nil {
		return m.SysErr
	}
	return nil
}

// Write writes len(b) bytes from the given offset. It returns the number
// of bytes written and an error, if any.
// Write returns an error when n != len(b).
type WriteRequest struct {
	Name   string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Offset int64  `protobuf:"varint,2,opt,name=offset" json:"offset,omitempty"`
	Data   []byte `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	Append bool   `protobuf:"varint,4,opt,name=append" json:"append,omitempty"`
}

func (m *WriteRequest) Reset()         { *m = WriteRequest{} }
func (m *WriteRequest) String() string { return proto1.CompactTextString(m) }
func (*WriteRequest) ProtoMessage()    {}

type WriteReply struct {
	Error        *Error `protobuf:"bytes,1,opt,name=error" json:"error,omitempty"`
	BytesWritten int64  `protobuf:"varint,2,opt,name=bytes_written" json:"bytes_written,omitempty"`
}

func (m *WriteReply) Reset()         { *m = WriteReply{} }
func (m *WriteReply) String() string { return proto1.CompactTextString(m) }
func (*WriteReply) ProtoMessage()    {}

func (m *WriteReply) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

// Read reads up to length bytes. The checksum of the data must match the exp_checksum if given, or an error is returned.
type ReadRequest struct {
	Name        string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Offset      int64  `protobuf:"varint,2,opt,name=offset" json:"offset,omitempty"`
	Length      int64  `protobuf:"varint,3,opt,name=length" json:"length,omitempty"`
	ExpChecksum uint32 `protobuf:"fixed32,4,opt,name=exp_checksum" json:"exp_checksum,omitempty"`
}

func (m *ReadRequest) Reset()         { *m = ReadRequest{} }
func (m *ReadRequest) String() string { return proto1.CompactTextString(m) }
func (*ReadRequest) ProtoMessage()    {}

type ReadReply struct {
	Error     *Error `protobuf:"bytes,1,opt,name=error" json:"error,omitempty"`
	BytesRead int64  `protobuf:"varint,2,opt,name=bytes_read" json:"bytes_read,omitempty"`
	Data      []byte `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	Checksum  uint32 `protobuf:"fixed32,4,opt,name=checksum" json:"checksum,omitempty"`
}

func (m *ReadReply) Reset()         { *m = ReadReply{} }
func (m *ReadReply) String() string { return proto1.CompactTextString(m) }
func (*ReadReply) ProtoMessage()    {}

func (m *ReadReply) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func init() {
}

// Client API for Cfs service

type CfsClient interface {
	Write(ctx context.Context, in *WriteRequest, opts ...grpc.CallOption) (*WriteReply, error)
	Read(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (*ReadReply, error)
}

type cfsClient struct {
	cc *grpc.ClientConn
}

func NewCfsClient(cc *grpc.ClientConn) CfsClient {
	return &cfsClient{cc}
}

func (c *cfsClient) Write(ctx context.Context, in *WriteRequest, opts ...grpc.CallOption) (*WriteReply, error) {
	out := new(WriteReply)
	err := grpc.Invoke(ctx, "/proto.cfs/Write", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cfsClient) Read(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (*ReadReply, error) {
	out := new(ReadReply)
	err := grpc.Invoke(ctx, "/proto.cfs/Read", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Cfs service

type CfsServer interface {
	Write(context.Context, *WriteRequest) (*WriteReply, error)
	Read(context.Context, *ReadRequest) (*ReadReply, error)
}

func RegisterCfsServer(s *grpc.Server, srv CfsServer) {
	s.RegisterService(&_Cfs_serviceDesc, srv)
}

func _Cfs_Write_Handler(srv interface{}, ctx context.Context, codec grpc.Codec, buf []byte) (interface{}, error) {
	in := new(WriteRequest)
	if err := codec.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(CfsServer).Write(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Cfs_Read_Handler(srv interface{}, ctx context.Context, codec grpc.Codec, buf []byte) (interface{}, error) {
	in := new(ReadRequest)
	if err := codec.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(CfsServer).Read(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _Cfs_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.cfs",
	HandlerType: (*CfsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Write",
			Handler:    _Cfs_Write_Handler,
		},
		{
			MethodName: "Read",
			Handler:    _Cfs_Read_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}
