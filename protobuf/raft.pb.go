// Code generated by protoc-gen-go.
// source: raft.proto
// DO NOT EDIT!

/*
Package protobuf is a generated protocol buffer package.

It is generated from these files:
	raft.proto

It has these top-level messages:
	LogEntry
	AppendEntriesRequest
	AppendEntriesResponse
	RequestVoteRequest
	RequestVoteResponse
	SnapshotRecoveryRequest
	SnapshotRecoveryResponse
	SnapshotRequest
	SnapshotResponse
*/
package protobuf

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type LogEntry struct {
	Index       uint64 `protobuf:"varint,1,opt,name=Index" json:"Index,omitempty"`
	Term        uint64 `protobuf:"varint,2,opt,name=Term" json:"Term,omitempty"`
	CommandName string `protobuf:"bytes,3,opt,name=CommandName" json:"CommandName,omitempty"`
	Command     []byte `protobuf:"bytes,4,opt,name=Command,proto3" json:"Command,omitempty"`
}

func (m *LogEntry) Reset()                    { *m = LogEntry{} }
func (m *LogEntry) String() string            { return proto.CompactTextString(m) }
func (*LogEntry) ProtoMessage()               {}
func (*LogEntry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *LogEntry) GetIndex() uint64 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *LogEntry) GetTerm() uint64 {
	if m != nil {
		return m.Term
	}
	return 0
}

func (m *LogEntry) GetCommandName() string {
	if m != nil {
		return m.CommandName
	}
	return ""
}

func (m *LogEntry) GetCommand() []byte {
	if m != nil {
		return m.Command
	}
	return nil
}

type AppendEntriesRequest struct {
	Term         uint64      `protobuf:"varint,1,opt,name=Term" json:"Term,omitempty"`
	PrevLogIndex uint64      `protobuf:"varint,2,opt,name=PrevLogIndex" json:"PrevLogIndex,omitempty"`
	PrevLogTerm  uint64      `protobuf:"varint,3,opt,name=PrevLogTerm" json:"PrevLogTerm,omitempty"`
	CommitIndex  uint64      `protobuf:"varint,4,opt,name=CommitIndex" json:"CommitIndex,omitempty"`
	LeaderName   string      `protobuf:"bytes,5,opt,name=LeaderName" json:"LeaderName,omitempty"`
	Entries      []*LogEntry `protobuf:"bytes,6,rep,name=Entries" json:"Entries,omitempty"`
}

func (m *AppendEntriesRequest) Reset()                    { *m = AppendEntriesRequest{} }
func (m *AppendEntriesRequest) String() string            { return proto.CompactTextString(m) }
func (*AppendEntriesRequest) ProtoMessage()               {}
func (*AppendEntriesRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *AppendEntriesRequest) GetTerm() uint64 {
	if m != nil {
		return m.Term
	}
	return 0
}

func (m *AppendEntriesRequest) GetPrevLogIndex() uint64 {
	if m != nil {
		return m.PrevLogIndex
	}
	return 0
}

func (m *AppendEntriesRequest) GetPrevLogTerm() uint64 {
	if m != nil {
		return m.PrevLogTerm
	}
	return 0
}

func (m *AppendEntriesRequest) GetCommitIndex() uint64 {
	if m != nil {
		return m.CommitIndex
	}
	return 0
}

func (m *AppendEntriesRequest) GetLeaderName() string {
	if m != nil {
		return m.LeaderName
	}
	return ""
}

func (m *AppendEntriesRequest) GetEntries() []*LogEntry {
	if m != nil {
		return m.Entries
	}
	return nil
}

type AppendEntriesResponse struct {
	Term        uint64 `protobuf:"varint,1,opt,name=Term" json:"Term,omitempty"`
	Index       uint64 `protobuf:"varint,2,opt,name=Index" json:"Index,omitempty"`
	CommitIndex uint64 `protobuf:"varint,3,opt,name=CommitIndex" json:"CommitIndex,omitempty"`
	Success     bool   `protobuf:"varint,4,opt,name=Success" json:"Success,omitempty"`
}

func (m *AppendEntriesResponse) Reset()                    { *m = AppendEntriesResponse{} }
func (m *AppendEntriesResponse) String() string            { return proto.CompactTextString(m) }
func (*AppendEntriesResponse) ProtoMessage()               {}
func (*AppendEntriesResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *AppendEntriesResponse) GetTerm() uint64 {
	if m != nil {
		return m.Term
	}
	return 0
}

func (m *AppendEntriesResponse) GetIndex() uint64 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *AppendEntriesResponse) GetCommitIndex() uint64 {
	if m != nil {
		return m.CommitIndex
	}
	return 0
}

func (m *AppendEntriesResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

type RequestVoteRequest struct {
	Term          uint64 `protobuf:"varint,1,opt,name=Term" json:"Term,omitempty"`
	LastLogIndex  uint64 `protobuf:"varint,2,opt,name=LastLogIndex" json:"LastLogIndex,omitempty"`
	LastLogTerm   uint64 `protobuf:"varint,3,opt,name=LastLogTerm" json:"LastLogTerm,omitempty"`
	CandidateName string `protobuf:"bytes,4,opt,name=CandidateName" json:"CandidateName,omitempty"`
}

func (m *RequestVoteRequest) Reset()                    { *m = RequestVoteRequest{} }
func (m *RequestVoteRequest) String() string            { return proto.CompactTextString(m) }
func (*RequestVoteRequest) ProtoMessage()               {}
func (*RequestVoteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *RequestVoteRequest) GetTerm() uint64 {
	if m != nil {
		return m.Term
	}
	return 0
}

func (m *RequestVoteRequest) GetLastLogIndex() uint64 {
	if m != nil {
		return m.LastLogIndex
	}
	return 0
}

func (m *RequestVoteRequest) GetLastLogTerm() uint64 {
	if m != nil {
		return m.LastLogTerm
	}
	return 0
}

func (m *RequestVoteRequest) GetCandidateName() string {
	if m != nil {
		return m.CandidateName
	}
	return ""
}

type RequestVoteResponse struct {
	Term        uint64 `protobuf:"varint,1,opt,name=Term" json:"Term,omitempty"`
	VoteGranted bool   `protobuf:"varint,2,opt,name=VoteGranted" json:"VoteGranted,omitempty"`
}

func (m *RequestVoteResponse) Reset()                    { *m = RequestVoteResponse{} }
func (m *RequestVoteResponse) String() string            { return proto.CompactTextString(m) }
func (*RequestVoteResponse) ProtoMessage()               {}
func (*RequestVoteResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *RequestVoteResponse) GetTerm() uint64 {
	if m != nil {
		return m.Term
	}
	return 0
}

func (m *RequestVoteResponse) GetVoteGranted() bool {
	if m != nil {
		return m.VoteGranted
	}
	return false
}

type SnapshotRecoveryRequest struct {
	LeaderName string                          `protobuf:"bytes,1,opt,name=LeaderName" json:"LeaderName,omitempty"`
	LastIndex  uint64                          `protobuf:"varint,2,opt,name=LastIndex" json:"LastIndex,omitempty"`
	LastTerm   uint64                          `protobuf:"varint,3,opt,name=LastTerm" json:"LastTerm,omitempty"`
	Peers      []*SnapshotRecoveryRequest_Peer `protobuf:"bytes,4,rep,name=Peers" json:"Peers,omitempty"`
	State      []byte                          `protobuf:"bytes,5,opt,name=State,proto3" json:"State,omitempty"`
}

func (m *SnapshotRecoveryRequest) Reset()                    { *m = SnapshotRecoveryRequest{} }
func (m *SnapshotRecoveryRequest) String() string            { return proto.CompactTextString(m) }
func (*SnapshotRecoveryRequest) ProtoMessage()               {}
func (*SnapshotRecoveryRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *SnapshotRecoveryRequest) GetLeaderName() string {
	if m != nil {
		return m.LeaderName
	}
	return ""
}

func (m *SnapshotRecoveryRequest) GetLastIndex() uint64 {
	if m != nil {
		return m.LastIndex
	}
	return 0
}

func (m *SnapshotRecoveryRequest) GetLastTerm() uint64 {
	if m != nil {
		return m.LastTerm
	}
	return 0
}

func (m *SnapshotRecoveryRequest) GetPeers() []*SnapshotRecoveryRequest_Peer {
	if m != nil {
		return m.Peers
	}
	return nil
}

func (m *SnapshotRecoveryRequest) GetState() []byte {
	if m != nil {
		return m.State
	}
	return nil
}

type SnapshotRecoveryRequest_Peer struct {
	Name             string `protobuf:"bytes,1,opt,name=Name" json:"Name,omitempty"`
	ConnectionString string `protobuf:"bytes,2,opt,name=ConnectionString" json:"ConnectionString,omitempty"`
}

func (m *SnapshotRecoveryRequest_Peer) Reset()                    { *m = SnapshotRecoveryRequest_Peer{} }
func (m *SnapshotRecoveryRequest_Peer) String() string            { return proto.CompactTextString(m) }
func (*SnapshotRecoveryRequest_Peer) ProtoMessage()               {}
func (*SnapshotRecoveryRequest_Peer) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5, 0} }

func (m *SnapshotRecoveryRequest_Peer) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SnapshotRecoveryRequest_Peer) GetConnectionString() string {
	if m != nil {
		return m.ConnectionString
	}
	return ""
}

type SnapshotRecoveryResponse struct {
	Term        uint64 `protobuf:"varint,1,opt,name=Term" json:"Term,omitempty"`
	Success     bool   `protobuf:"varint,2,opt,name=Success" json:"Success,omitempty"`
	CommitIndex uint64 `protobuf:"varint,3,opt,name=CommitIndex" json:"CommitIndex,omitempty"`
}

func (m *SnapshotRecoveryResponse) Reset()                    { *m = SnapshotRecoveryResponse{} }
func (m *SnapshotRecoveryResponse) String() string            { return proto.CompactTextString(m) }
func (*SnapshotRecoveryResponse) ProtoMessage()               {}
func (*SnapshotRecoveryResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *SnapshotRecoveryResponse) GetTerm() uint64 {
	if m != nil {
		return m.Term
	}
	return 0
}

func (m *SnapshotRecoveryResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *SnapshotRecoveryResponse) GetCommitIndex() uint64 {
	if m != nil {
		return m.CommitIndex
	}
	return 0
}

type SnapshotRequest struct {
	LeaderName string `protobuf:"bytes,1,opt,name=LeaderName" json:"LeaderName,omitempty"`
	LastIndex  uint64 `protobuf:"varint,2,opt,name=LastIndex" json:"LastIndex,omitempty"`
	LastTerm   uint64 `protobuf:"varint,3,opt,name=LastTerm" json:"LastTerm,omitempty"`
}

func (m *SnapshotRequest) Reset()                    { *m = SnapshotRequest{} }
func (m *SnapshotRequest) String() string            { return proto.CompactTextString(m) }
func (*SnapshotRequest) ProtoMessage()               {}
func (*SnapshotRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *SnapshotRequest) GetLeaderName() string {
	if m != nil {
		return m.LeaderName
	}
	return ""
}

func (m *SnapshotRequest) GetLastIndex() uint64 {
	if m != nil {
		return m.LastIndex
	}
	return 0
}

func (m *SnapshotRequest) GetLastTerm() uint64 {
	if m != nil {
		return m.LastTerm
	}
	return 0
}

type SnapshotResponse struct {
	Success bool `protobuf:"varint,1,opt,name=Success" json:"Success,omitempty"`
}

func (m *SnapshotResponse) Reset()                    { *m = SnapshotResponse{} }
func (m *SnapshotResponse) String() string            { return proto.CompactTextString(m) }
func (*SnapshotResponse) ProtoMessage()               {}
func (*SnapshotResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *SnapshotResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func init() {
	proto.RegisterType((*LogEntry)(nil), "protobuf.LogEntry")
	proto.RegisterType((*AppendEntriesRequest)(nil), "protobuf.AppendEntriesRequest")
	proto.RegisterType((*AppendEntriesResponse)(nil), "protobuf.AppendEntriesResponse")
	proto.RegisterType((*RequestVoteRequest)(nil), "protobuf.RequestVoteRequest")
	proto.RegisterType((*RequestVoteResponse)(nil), "protobuf.RequestVoteResponse")
	proto.RegisterType((*SnapshotRecoveryRequest)(nil), "protobuf.SnapshotRecoveryRequest")
	proto.RegisterType((*SnapshotRecoveryRequest_Peer)(nil), "protobuf.SnapshotRecoveryRequest.Peer")
	proto.RegisterType((*SnapshotRecoveryResponse)(nil), "protobuf.SnapshotRecoveryResponse")
	proto.RegisterType((*SnapshotRequest)(nil), "protobuf.SnapshotRequest")
	proto.RegisterType((*SnapshotResponse)(nil), "protobuf.SnapshotResponse")
}

func init() { proto.RegisterFile("raft.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 485 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xb4, 0x92, 0xc1, 0x6f, 0xd3, 0x30,
	0x14, 0xc6, 0xe5, 0x36, 0x5d, 0xd3, 0xd7, 0x22, 0x26, 0x33, 0x44, 0x34, 0x21, 0x14, 0x59, 0x08,
	0x55, 0x68, 0xea, 0x01, 0xae, 0x5c, 0x50, 0x05, 0x08, 0x51, 0xa1, 0xc9, 0x45, 0xdc, 0xbd, 0xe6,
	0xad, 0x04, 0x54, 0x3b, 0xd8, 0xee, 0xc4, 0x0e, 0xfc, 0x0d, 0x5c, 0xf8, 0x03, 0xf9, 0x53, 0x90,
	0xed, 0xa4, 0x75, 0x1a, 0xd8, 0x4e, 0x3b, 0x35, 0xdf, 0x57, 0x27, 0xef, 0xf7, 0x7d, 0x7e, 0x00,
	0x5a, 0x5c, 0xda, 0x59, 0xa5, 0x95, 0x55, 0x34, 0xf5, 0x3f, 0x17, 0xdb, 0x4b, 0x56, 0x41, 0xba,
	0x50, 0xeb, 0x37, 0xd2, 0xea, 0x6b, 0x7a, 0x02, 0x83, 0xf7, 0xb2, 0xc0, 0x1f, 0x19, 0xc9, 0xc9,
	0x34, 0xe1, 0x41, 0x50, 0x0a, 0xc9, 0x27, 0xd4, 0x9b, 0xac, 0xe7, 0x4d, 0xff, 0x4c, 0x73, 0x18,
	0xcf, 0xd5, 0x66, 0x23, 0x64, 0xf1, 0x51, 0x6c, 0x30, 0xeb, 0xe7, 0x64, 0x3a, 0xe2, 0xb1, 0x45,
	0x33, 0x18, 0xd6, 0x32, 0x4b, 0x72, 0x32, 0x9d, 0xf0, 0x46, 0xb2, 0x3f, 0x04, 0x4e, 0x5e, 0x57,
	0x15, 0xca, 0xc2, 0x4d, 0x2d, 0xd1, 0x70, 0xfc, 0xbe, 0x45, 0x63, 0x77, 0x83, 0x48, 0x34, 0x88,
	0xc1, 0xe4, 0x5c, 0xe3, 0xd5, 0x42, 0xad, 0x03, 0x59, 0x80, 0x68, 0x79, 0x0e, 0xa6, 0xd6, 0xfe,
	0xf5, 0xbe, 0x3f, 0x12, 0x5b, 0x0d, 0x6e, 0x69, 0xc3, 0x47, 0x92, 0x70, 0x22, 0xb2, 0xe8, 0x13,
	0x80, 0x05, 0x8a, 0x02, 0xb5, 0xcf, 0x33, 0xf0, 0x79, 0x22, 0x87, 0x9e, 0xc1, 0xb0, 0xa6, 0xcd,
	0x8e, 0xf2, 0xfe, 0x74, 0xfc, 0x82, 0xce, 0x9a, 0x0a, 0x67, 0x4d, 0x7f, 0xbc, 0x39, 0xc2, 0x7e,
	0xc2, 0xc3, 0x83, 0x84, 0xa6, 0x52, 0xd2, 0xe0, 0x3f, 0x23, 0xee, 0x5a, 0xef, 0xc5, 0xad, 0x1f,
	0x20, 0xf7, 0xbb, 0xc8, 0x19, 0x0c, 0x97, 0xdb, 0xd5, 0x0a, 0x8d, 0xf1, 0x81, 0x52, 0xde, 0x48,
	0xf6, 0x9b, 0x00, 0xad, 0x4b, 0xfd, 0xac, 0x2c, 0xde, 0xd2, 0xef, 0x42, 0x18, 0x7b, 0xd8, 0x6f,
	0xec, 0x39, 0x94, 0x5a, 0xc7, 0xfd, 0x46, 0x16, 0x7d, 0x0a, 0xf7, 0xe6, 0x42, 0x16, 0x65, 0x21,
	0x2c, 0xfa, 0x02, 0x13, 0x5f, 0x60, 0xdb, 0x64, 0x1f, 0xe0, 0x41, 0x8b, 0xea, 0x86, 0x4e, 0x72,
	0x18, 0xbb, 0x33, 0xef, 0xb4, 0x90, 0x16, 0x0b, 0x4f, 0x95, 0xf2, 0xd8, 0x62, 0xbf, 0x7a, 0xf0,
	0x68, 0x29, 0x45, 0x65, 0xbe, 0x28, 0xcb, 0x71, 0xa5, 0xae, 0x50, 0x5f, 0x37, 0x41, 0xdb, 0x97,
	0x49, 0x3a, 0x97, 0xf9, 0x18, 0x46, 0x8e, 0x3e, 0x4e, 0xbc, 0x37, 0xe8, 0x29, 0xa4, 0x4e, 0x44,
	0x59, 0x77, 0x9a, 0xbe, 0x82, 0xc1, 0x39, 0xa2, 0x76, 0x8d, 0xbb, 0x25, 0x78, 0xb6, 0x5f, 0x82,
	0xff, 0xb0, 0xcc, 0xdc, 0x71, 0x1e, 0x5e, 0x72, 0x37, 0xbd, 0xb4, 0xc2, 0x86, 0xfd, 0x9a, 0xf0,
	0x20, 0x4e, 0xdf, 0x42, 0xe2, 0xfe, 0x76, 0x3d, 0x44, 0xbc, 0xfe, 0x99, 0x3e, 0x87, 0xe3, 0xb9,
	0x92, 0x12, 0x57, 0xb6, 0x54, 0x72, 0x69, 0x75, 0x29, 0xd7, 0x1e, 0x78, 0xc4, 0x3b, 0x3e, 0xfb,
	0x0a, 0x59, 0x17, 0xe2, 0x86, 0x8e, 0xa3, 0xfd, 0xe9, 0xb5, 0xf6, 0xe7, 0xf6, 0xdd, 0x63, 0xdf,
	0xe0, 0xfe, 0x7e, 0xd6, 0x1d, 0x97, 0xce, 0xce, 0xe0, 0x78, 0x3f, 0xac, 0x0e, 0x14, 0xc1, 0x93,
	0x16, 0xfc, 0xc5, 0x91, 0xbf, 0x92, 0x97, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x9f, 0xc8, 0x40,
	0x46, 0xef, 0x04, 0x00, 0x00,
}