// Code generated by protoc-gen-go. DO NOT EDIT.
// source: task.proto

package events

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	codec "github.com/kcarretto/paragon/api/codec"
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

type TaskQueued struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Content              string   `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	SessionID            string   `protobuf:"bytes,3,opt,name=sessionID,proto3" json:"sessionID,omitempty"`
	MachineUUID          string   `protobuf:"bytes,4,opt,name=machineUUID,proto3" json:"machineUUID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TaskQueued) Reset()         { *m = TaskQueued{} }
func (m *TaskQueued) String() string { return proto.CompactTextString(m) }
func (*TaskQueued) ProtoMessage()    {}
func (*TaskQueued) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce5d8dd45b4a91ff, []int{0}
}

func (m *TaskQueued) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskQueued.Unmarshal(m, b)
}
func (m *TaskQueued) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskQueued.Marshal(b, m, deterministic)
}
func (m *TaskQueued) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskQueued.Merge(m, src)
}
func (m *TaskQueued) XXX_Size() int {
	return xxx_messageInfo_TaskQueued.Size(m)
}
func (m *TaskQueued) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskQueued.DiscardUnknown(m)
}

var xxx_messageInfo_TaskQueued proto.InternalMessageInfo

func (m *TaskQueued) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *TaskQueued) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *TaskQueued) GetSessionID() string {
	if m != nil {
		return m.SessionID
	}
	return ""
}

func (m *TaskQueued) GetMachineUUID() string {
	if m != nil {
		return m.MachineUUID
	}
	return ""
}

type TaskClaimed struct {
	Id                   string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Agent                *codec.AgentMetadata `protobuf:"bytes,2,opt,name=agent,proto3" json:"agent,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *TaskClaimed) Reset()         { *m = TaskClaimed{} }
func (m *TaskClaimed) String() string { return proto.CompactTextString(m) }
func (*TaskClaimed) ProtoMessage()    {}
func (*TaskClaimed) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce5d8dd45b4a91ff, []int{1}
}

func (m *TaskClaimed) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskClaimed.Unmarshal(m, b)
}
func (m *TaskClaimed) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskClaimed.Marshal(b, m, deterministic)
}
func (m *TaskClaimed) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskClaimed.Merge(m, src)
}
func (m *TaskClaimed) XXX_Size() int {
	return xxx_messageInfo_TaskClaimed.Size(m)
}
func (m *TaskClaimed) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskClaimed.DiscardUnknown(m)
}

var xxx_messageInfo_TaskClaimed proto.InternalMessageInfo

func (m *TaskClaimed) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *TaskClaimed) GetAgent() *codec.AgentMetadata {
	if m != nil {
		return m.Agent
	}
	return nil
}

type TaskExecuted struct {
	Id                   string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Output               []string             `protobuf:"bytes,2,rep,name=output,proto3" json:"output,omitempty"`
	Error                string               `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
	ExecStartTime        int64                `protobuf:"varint,4,opt,name=execStartTime,proto3" json:"execStartTime,omitempty"`
	ExecStopTime         int64                `protobuf:"varint,5,opt,name=execStopTime,proto3" json:"execStopTime,omitempty"`
	RecvTime             int64                `protobuf:"varint,6,opt,name=recvTime,proto3" json:"recvTime,omitempty"`
	Agent                *codec.AgentMetadata `protobuf:"bytes,7,opt,name=agent,proto3" json:"agent,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *TaskExecuted) Reset()         { *m = TaskExecuted{} }
func (m *TaskExecuted) String() string { return proto.CompactTextString(m) }
func (*TaskExecuted) ProtoMessage()    {}
func (*TaskExecuted) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce5d8dd45b4a91ff, []int{2}
}

func (m *TaskExecuted) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskExecuted.Unmarshal(m, b)
}
func (m *TaskExecuted) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskExecuted.Marshal(b, m, deterministic)
}
func (m *TaskExecuted) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskExecuted.Merge(m, src)
}
func (m *TaskExecuted) XXX_Size() int {
	return xxx_messageInfo_TaskExecuted.Size(m)
}
func (m *TaskExecuted) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskExecuted.DiscardUnknown(m)
}

var xxx_messageInfo_TaskExecuted proto.InternalMessageInfo

func (m *TaskExecuted) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *TaskExecuted) GetOutput() []string {
	if m != nil {
		return m.Output
	}
	return nil
}

func (m *TaskExecuted) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *TaskExecuted) GetExecStartTime() int64 {
	if m != nil {
		return m.ExecStartTime
	}
	return 0
}

func (m *TaskExecuted) GetExecStopTime() int64 {
	if m != nil {
		return m.ExecStopTime
	}
	return 0
}

func (m *TaskExecuted) GetRecvTime() int64 {
	if m != nil {
		return m.RecvTime
	}
	return 0
}

func (m *TaskExecuted) GetAgent() *codec.AgentMetadata {
	if m != nil {
		return m.Agent
	}
	return nil
}

func init() {
	proto.RegisterType((*TaskQueued)(nil), "events.TaskQueued")
	proto.RegisterType((*TaskClaimed)(nil), "events.TaskClaimed")
	proto.RegisterType((*TaskExecuted)(nil), "events.TaskExecuted")
}

func init() { proto.RegisterFile("task.proto", fileDescriptor_ce5d8dd45b4a91ff) }

var fileDescriptor_ce5d8dd45b4a91ff = []byte{
	// 313 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0x41, 0x4f, 0xc2, 0x40,
	0x10, 0x85, 0x03, 0x48, 0x91, 0x01, 0x3d, 0x6c, 0xd0, 0x34, 0xc4, 0x03, 0x69, 0x4c, 0x44, 0x0f,
	0x6d, 0xa2, 0xbf, 0x40, 0xc5, 0x03, 0x07, 0x0f, 0x56, 0xb8, 0x78, 0x1b, 0xb6, 0x13, 0xd8, 0x60,
	0x77, 0x9b, 0xed, 0x2c, 0xe1, 0xd7, 0xfa, 0x5b, 0x4c, 0xb7, 0x28, 0x92, 0x90, 0x78, 0x7c, 0xdf,
	0x9b, 0xcc, 0xbc, 0xc9, 0x03, 0x60, 0x2c, 0xd7, 0x71, 0x61, 0x0d, 0x1b, 0x11, 0xd0, 0x86, 0x34,
	0x97, 0xc3, 0x0b, 0x2c, 0x54, 0x22, 0x4d, 0x46, 0x32, 0xc1, 0x25, 0x69, 0xae, 0xed, 0x68, 0x03,
	0x30, 0xc3, 0x72, 0xfd, 0xe6, 0xc8, 0x51, 0x26, 0xce, 0xa1, 0xa9, 0xb2, 0xb0, 0x31, 0x6a, 0x8c,
	0xbb, 0x69, 0x53, 0x65, 0x22, 0x84, 0x8e, 0x34, 0x9a, 0x49, 0x73, 0xd8, 0xf4, 0xf0, 0x47, 0x8a,
	0x2b, 0xe8, 0x96, 0x54, 0x96, 0xca, 0xe8, 0xe9, 0x24, 0x6c, 0x79, 0x6f, 0x0f, 0xc4, 0x08, 0x7a,
	0x39, 0xca, 0x95, 0xd2, 0x34, 0x9f, 0x4f, 0x27, 0xe1, 0x89, 0xf7, 0xff, 0xa2, 0x68, 0x0a, 0xbd,
	0xea, 0xee, 0xf3, 0x27, 0xaa, 0xfc, 0xc8, 0xe1, 0x3b, 0x68, 0xfb, 0x94, 0xfe, 0x6c, 0xef, 0x7e,
	0x10, 0xfb, 0xe4, 0xf1, 0x63, 0xc5, 0x5e, 0x89, 0x31, 0x43, 0xc6, 0xb4, 0x1e, 0x89, 0xbe, 0x1a,
	0xd0, 0xaf, 0x76, 0xbd, 0x6c, 0x49, 0x3a, 0x3e, 0xb2, 0xec, 0x12, 0x02, 0xe3, 0xb8, 0x70, 0xd5,
	0xb6, 0xd6, 0xb8, 0x9b, 0xee, 0x94, 0x18, 0x40, 0x9b, 0xac, 0x35, 0x76, 0x97, 0xbf, 0x16, 0xe2,
	0x1a, 0xce, 0x68, 0x4b, 0xf2, 0x9d, 0xd1, 0xf2, 0x4c, 0xe5, 0xe4, 0xd3, 0xb7, 0xd2, 0x43, 0x28,
	0x22, 0xe8, 0xd7, 0xc0, 0x14, 0x7e, 0xa8, 0xed, 0x87, 0x0e, 0x98, 0x18, 0xc2, 0xa9, 0x25, 0xb9,
	0xf1, 0x7e, 0xe0, 0xfd, 0x5f, 0xbd, 0x7f, 0xb0, 0xf3, 0xef, 0x83, 0x4f, 0xb7, 0x1f, 0x37, 0x4b,
	0xc5, 0x2b, 0xb7, 0x88, 0xa5, 0xc9, 0x93, 0xb5, 0x44, 0x6b, 0x89, 0xd9, 0x24, 0x05, 0x5a, 0x5c,
	0x1a, 0x9d, 0x54, 0xcd, 0xd6, 0x2d, 0x2f, 0x02, 0xdf, 0xea, 0xc3, 0x77, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xb7, 0x7d, 0x2a, 0x3d, 0x02, 0x02, 0x00, 0x00,
}
