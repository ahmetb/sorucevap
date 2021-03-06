// Code generated by protoc-gen-go. DO NOT EDIT.
// source: sorucevap.proto

package sorucevap

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type VoteOp int32

const (
	VoteOp_UNKNOWN  VoteOp = 0
	VoteOp_UPVOTE   VoteOp = 1
	VoteOp_DOWNVOTE VoteOp = 2
)

var VoteOp_name = map[int32]string{
	0: "UNKNOWN",
	1: "UPVOTE",
	2: "DOWNVOTE",
}

var VoteOp_value = map[string]int32{
	"UNKNOWN":  0,
	"UPVOTE":   1,
	"DOWNVOTE": 2,
}

func (x VoteOp) String() string {
	return proto.EnumName(VoteOp_name, int32(x))
}

func (VoteOp) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9587f7a504948045, []int{0}
}

// UserRecord represents database record of a signed in user.
type UserRecord struct {
	ID                   string   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	FullName             string   `protobuf:"bytes,2,opt,name=FullName,proto3" json:"FullName,omitempty"`
	ProfilePictureURL    string   `protobuf:"bytes,3,opt,name=ProfilePictureURL,proto3" json:"ProfilePictureURL,omitempty"`
	RegistrationDate     int64    `protobuf:"varint,4,opt,name=RegistrationDate,proto3" json:"RegistrationDate,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserRecord) Reset()         { *m = UserRecord{} }
func (m *UserRecord) String() string { return proto.CompactTextString(m) }
func (*UserRecord) ProtoMessage()    {}
func (*UserRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_9587f7a504948045, []int{0}
}

func (m *UserRecord) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserRecord.Unmarshal(m, b)
}
func (m *UserRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserRecord.Marshal(b, m, deterministic)
}
func (m *UserRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserRecord.Merge(m, src)
}
func (m *UserRecord) XXX_Size() int {
	return xxx_messageInfo_UserRecord.Size(m)
}
func (m *UserRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_UserRecord.DiscardUnknown(m)
}

var xxx_messageInfo_UserRecord proto.InternalMessageInfo

func (m *UserRecord) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *UserRecord) GetFullName() string {
	if m != nil {
		return m.FullName
	}
	return ""
}

func (m *UserRecord) GetProfilePictureURL() string {
	if m != nil {
		return m.ProfilePictureURL
	}
	return ""
}

func (m *UserRecord) GetRegistrationDate() int64 {
	if m != nil {
		return m.RegistrationDate
	}
	return 0
}

// UserSummary is to record user details in an embedded document.
type UserSummary struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	FullName             string   `protobuf:"bytes,2,opt,name=fullName,proto3" json:"fullName,omitempty"`
	ProfilePictureURL    string   `protobuf:"bytes,3,opt,name=profilePictureURL,proto3" json:"profilePictureURL,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserSummary) Reset()         { *m = UserSummary{} }
func (m *UserSummary) String() string { return proto.CompactTextString(m) }
func (*UserSummary) ProtoMessage()    {}
func (*UserSummary) Descriptor() ([]byte, []int) {
	return fileDescriptor_9587f7a504948045, []int{1}
}

func (m *UserSummary) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserSummary.Unmarshal(m, b)
}
func (m *UserSummary) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserSummary.Marshal(b, m, deterministic)
}
func (m *UserSummary) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserSummary.Merge(m, src)
}
func (m *UserSummary) XXX_Size() int {
	return xxx_messageInfo_UserSummary.Size(m)
}
func (m *UserSummary) XXX_DiscardUnknown() {
	xxx_messageInfo_UserSummary.DiscardUnknown(m)
}

var xxx_messageInfo_UserSummary proto.InternalMessageInfo

func (m *UserSummary) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *UserSummary) GetFullName() string {
	if m != nil {
		return m.FullName
	}
	return ""
}

func (m *UserSummary) GetProfilePictureURL() string {
	if m != nil {
		return m.ProfilePictureURL
	}
	return ""
}

type Event struct {
	Id                   string       `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string       `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description          string       `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	CreatedAt            int64        `protobuf:"varint,4,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	ExpiresAt            int64        `protobuf:"varint,8,opt,name=expiresAt,proto3" json:"expiresAt,omitempty"`
	Creator              *UserSummary `protobuf:"bytes,20,opt,name=creator,proto3" json:"creator,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_9587f7a504948045, []int{2}
}

func (m *Event) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Event.Unmarshal(m, b)
}
func (m *Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Event.Marshal(b, m, deterministic)
}
func (m *Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event.Merge(m, src)
}
func (m *Event) XXX_Size() int {
	return xxx_messageInfo_Event.Size(m)
}
func (m *Event) XXX_DiscardUnknown() {
	xxx_messageInfo_Event.DiscardUnknown(m)
}

var xxx_messageInfo_Event proto.InternalMessageInfo

func (m *Event) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Event) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Event) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Event) GetCreatedAt() int64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func (m *Event) GetExpiresAt() int64 {
	if m != nil {
		return m.ExpiresAt
	}
	return 0
}

func (m *Event) GetCreator() *UserSummary {
	if m != nil {
		return m.Creator
	}
	return nil
}

// Vote represents an embedded vote record
type Vote struct {
	Author               *UserSummary `protobuf:"bytes,1,opt,name=author,proto3" json:"author,omitempty"`
	Op                   VoteOp       `protobuf:"varint,2,opt,name=op,proto3,enum=VoteOp" json:"op,omitempty"`
	CreatedAt            int64        `protobuf:"varint,3,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Vote) Reset()         { *m = Vote{} }
func (m *Vote) String() string { return proto.CompactTextString(m) }
func (*Vote) ProtoMessage()    {}
func (*Vote) Descriptor() ([]byte, []int) {
	return fileDescriptor_9587f7a504948045, []int{3}
}

func (m *Vote) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Vote.Unmarshal(m, b)
}
func (m *Vote) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Vote.Marshal(b, m, deterministic)
}
func (m *Vote) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Vote.Merge(m, src)
}
func (m *Vote) XXX_Size() int {
	return xxx_messageInfo_Vote.Size(m)
}
func (m *Vote) XXX_DiscardUnknown() {
	xxx_messageInfo_Vote.DiscardUnknown(m)
}

var xxx_messageInfo_Vote proto.InternalMessageInfo

func (m *Vote) GetAuthor() *UserSummary {
	if m != nil {
		return m.Author
	}
	return nil
}

func (m *Vote) GetOp() VoteOp {
	if m != nil {
		return m.Op
	}
	return VoteOp_UNKNOWN
}

func (m *Vote) GetCreatedAt() int64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

// Question represents the db record.
type Question struct {
	Id                   string       `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	EventID              string       `protobuf:"bytes,2,opt,name=eventID,proto3" json:"eventID,omitempty"`
	Author               *UserSummary `protobuf:"bytes,3,opt,name=author,proto3" json:"author,omitempty"`
	Deleted              bool         `protobuf:"varint,4,opt,name=deleted,proto3" json:"deleted,omitempty"`
	CreatedAt            int64        `protobuf:"varint,5,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	Contents             string       `protobuf:"bytes,10,opt,name=contents,proto3" json:"contents,omitempty"`
	ContextLinks         []string     `protobuf:"bytes,11,rep,name=contextLinks,proto3" json:"contextLinks,omitempty"`
	NetVotes             int32        `protobuf:"varint,100,opt,name=netVotes,proto3" json:"netVotes,omitempty"`
	Votes                []*Vote      `protobuf:"bytes,101,rep,name=votes,proto3" json:"votes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Question) Reset()         { *m = Question{} }
func (m *Question) String() string { return proto.CompactTextString(m) }
func (*Question) ProtoMessage()    {}
func (*Question) Descriptor() ([]byte, []int) {
	return fileDescriptor_9587f7a504948045, []int{4}
}

func (m *Question) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Question.Unmarshal(m, b)
}
func (m *Question) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Question.Marshal(b, m, deterministic)
}
func (m *Question) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Question.Merge(m, src)
}
func (m *Question) XXX_Size() int {
	return xxx_messageInfo_Question.Size(m)
}
func (m *Question) XXX_DiscardUnknown() {
	xxx_messageInfo_Question.DiscardUnknown(m)
}

var xxx_messageInfo_Question proto.InternalMessageInfo

func (m *Question) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Question) GetEventID() string {
	if m != nil {
		return m.EventID
	}
	return ""
}

func (m *Question) GetAuthor() *UserSummary {
	if m != nil {
		return m.Author
	}
	return nil
}

func (m *Question) GetDeleted() bool {
	if m != nil {
		return m.Deleted
	}
	return false
}

func (m *Question) GetCreatedAt() int64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func (m *Question) GetContents() string {
	if m != nil {
		return m.Contents
	}
	return ""
}

func (m *Question) GetContextLinks() []string {
	if m != nil {
		return m.ContextLinks
	}
	return nil
}

func (m *Question) GetNetVotes() int32 {
	if m != nil {
		return m.NetVotes
	}
	return 0
}

func (m *Question) GetVotes() []*Vote {
	if m != nil {
		return m.Votes
	}
	return nil
}

type DeleteEventRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteEventRequest) Reset()         { *m = DeleteEventRequest{} }
func (m *DeleteEventRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteEventRequest) ProtoMessage()    {}
func (*DeleteEventRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9587f7a504948045, []int{5}
}

func (m *DeleteEventRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteEventRequest.Unmarshal(m, b)
}
func (m *DeleteEventRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteEventRequest.Marshal(b, m, deterministic)
}
func (m *DeleteEventRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteEventRequest.Merge(m, src)
}
func (m *DeleteEventRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteEventRequest.Size(m)
}
func (m *DeleteEventRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteEventRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteEventRequest proto.InternalMessageInfo

func (m *DeleteEventRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type GetEventRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetEventRequest) Reset()         { *m = GetEventRequest{} }
func (m *GetEventRequest) String() string { return proto.CompactTextString(m) }
func (*GetEventRequest) ProtoMessage()    {}
func (*GetEventRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9587f7a504948045, []int{6}
}

func (m *GetEventRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetEventRequest.Unmarshal(m, b)
}
func (m *GetEventRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetEventRequest.Marshal(b, m, deterministic)
}
func (m *GetEventRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetEventRequest.Merge(m, src)
}
func (m *GetEventRequest) XXX_Size() int {
	return xxx_messageInfo_GetEventRequest.Size(m)
}
func (m *GetEventRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetEventRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetEventRequest proto.InternalMessageInfo

func (m *GetEventRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type GetUserRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserRequest) Reset()         { *m = GetUserRequest{} }
func (m *GetUserRequest) String() string { return proto.CompactTextString(m) }
func (*GetUserRequest) ProtoMessage()    {}
func (*GetUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9587f7a504948045, []int{7}
}

func (m *GetUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserRequest.Unmarshal(m, b)
}
func (m *GetUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserRequest.Marshal(b, m, deterministic)
}
func (m *GetUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserRequest.Merge(m, src)
}
func (m *GetUserRequest) XXX_Size() int {
	return xxx_messageInfo_GetUserRequest.Size(m)
}
func (m *GetUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserRequest proto.InternalMessageInfo

func (m *GetUserRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func init() {
	proto.RegisterEnum("VoteOp", VoteOp_name, VoteOp_value)
	proto.RegisterType((*UserRecord)(nil), "UserRecord")
	proto.RegisterType((*UserSummary)(nil), "UserSummary")
	proto.RegisterType((*Event)(nil), "Event")
	proto.RegisterType((*Vote)(nil), "Vote")
	proto.RegisterType((*Question)(nil), "Question")
	proto.RegisterType((*DeleteEventRequest)(nil), "DeleteEventRequest")
	proto.RegisterType((*GetEventRequest)(nil), "GetEventRequest")
	proto.RegisterType((*GetUserRequest)(nil), "GetUserRequest")
}

func init() { proto.RegisterFile("sorucevap.proto", fileDescriptor_9587f7a504948045) }

var fileDescriptor_9587f7a504948045 = []byte{
	// 608 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x54, 0xd1, 0x6e, 0xd3, 0x30,
	0x14, 0x5d, 0xd2, 0xb5, 0x49, 0x6f, 0xa6, 0xad, 0x18, 0x04, 0x51, 0x87, 0x44, 0x88, 0xa6, 0x51,
	0x4d, 0xc8, 0x93, 0xca, 0x2b, 0x2f, 0x83, 0x8e, 0x69, 0x62, 0x6a, 0x87, 0xa1, 0x9b, 0xc4, 0x13,
	0x59, 0x72, 0x5b, 0x22, 0xda, 0x38, 0x38, 0xce, 0xb4, 0x7d, 0x03, 0x0f, 0x48, 0xfc, 0x09, 0x7f,
	0x88, 0xec, 0x34, 0x5d, 0x9a, 0xb2, 0x3d, 0x25, 0xf7, 0x9c, 0x1b, 0xfb, 0x9c, 0x63, 0xe7, 0xc2,
	0x4e, 0xc6, 0x45, 0x1e, 0xe2, 0x75, 0x90, 0xd2, 0x54, 0x70, 0xc9, 0xbb, 0xbb, 0x53, 0xce, 0xa7,
	0x33, 0x3c, 0xd4, 0xd5, 0x55, 0x3e, 0x39, 0xc4, 0x79, 0x2a, 0x6f, 0x0b, 0xd2, 0xff, 0x63, 0x00,
	0x8c, 0x33, 0x14, 0x0c, 0x43, 0x2e, 0x22, 0xb2, 0x0d, 0xe6, 0xe9, 0xc0, 0x35, 0x3c, 0xa3, 0xd7,
	0x66, 0xe6, 0xe9, 0x80, 0x74, 0xc1, 0xfe, 0x90, 0xcf, 0x66, 0xc3, 0x60, 0x8e, 0xae, 0xa9, 0xd1,
	0x65, 0x4d, 0x5e, 0xc3, 0xa3, 0x73, 0xc1, 0x27, 0xf1, 0x0c, 0xcf, 0xe3, 0x50, 0xe6, 0x02, 0xc7,
	0xec, 0xcc, 0x6d, 0xe8, 0xa6, 0x75, 0x82, 0x1c, 0x40, 0x87, 0xe1, 0x34, 0xce, 0xa4, 0x08, 0x64,
	0xcc, 0x93, 0x41, 0x20, 0xd1, 0xdd, 0xf4, 0x8c, 0x5e, 0x83, 0xad, 0xe1, 0xfe, 0x14, 0x1c, 0xa5,
	0xe9, 0x73, 0x3e, 0x9f, 0x07, 0xe2, 0x56, 0x89, 0x8a, 0xa3, 0x52, 0x54, 0x1c, 0x29, 0x51, 0x93,
	0x9a, 0xa8, 0x49, 0x45, 0x54, 0x7a, 0x9f, 0xa8, 0x35, 0xc2, 0xff, 0x6b, 0x40, 0xf3, 0xf8, 0x1a,
	0x13, 0xb9, 0xb6, 0x07, 0x81, 0xcd, 0xe4, 0x6e, 0x7d, 0xfd, 0x4e, 0x3c, 0x70, 0x22, 0xcc, 0x42,
	0x11, 0xa7, 0x4a, 0xe9, 0x62, 0xd5, 0x2a, 0x44, 0x9e, 0x43, 0x3b, 0x14, 0x18, 0x48, 0x8c, 0x8e,
	0xe4, 0xc2, 0xdd, 0x1d, 0xa0, 0x58, 0xbc, 0x49, 0x63, 0x81, 0xd9, 0x91, 0x74, 0xed, 0x82, 0x5d,
	0x02, 0x64, 0x1f, 0x2c, 0xdd, 0xca, 0x85, 0xfb, 0xc4, 0x33, 0x7a, 0x4e, 0x7f, 0x8b, 0x56, 0x42,
	0x60, 0x25, 0xe9, 0x87, 0xb0, 0x79, 0xc1, 0x25, 0x92, 0x3d, 0x68, 0x05, 0xb9, 0xfc, 0xce, 0x85,
	0x56, 0x5d, 0x6f, 0x5f, 0x70, 0xe4, 0x19, 0x98, 0x3c, 0xd5, 0x2e, 0xb6, 0xfb, 0x16, 0x55, 0x1f,
	0x8e, 0x52, 0x66, 0xf2, 0x74, 0x55, 0x6a, 0xa3, 0x26, 0xd5, 0xff, 0x65, 0x82, 0xfd, 0x29, 0xc7,
	0x4c, 0xbb, 0xaa, 0x67, 0xe3, 0x82, 0x85, 0x2a, 0xb4, 0xd3, 0xc1, 0x22, 0x9e, 0xb2, 0xac, 0x68,
	0x6a, 0x3c, 0xa0, 0xc9, 0x05, 0x2b, 0xc2, 0x19, 0x4a, 0x8c, 0x74, 0x46, 0x36, 0x2b, 0xcb, 0x55,
	0x51, 0xcd, 0x7a, 0x7e, 0x5d, 0xb0, 0x43, 0x9e, 0x48, 0x4c, 0x64, 0xe6, 0x42, 0x71, 0xee, 0x65,
	0x4d, 0x7c, 0xd8, 0xd2, 0xef, 0x37, 0xf2, 0x2c, 0x4e, 0x7e, 0x64, 0xae, 0xe3, 0x35, 0x7a, 0x6d,
	0xb6, 0x82, 0xa9, 0xef, 0x13, 0x94, 0x2a, 0x83, 0xcc, 0x8d, 0x3c, 0xa3, 0xd7, 0x64, 0xcb, 0x9a,
	0xec, 0x42, 0xf3, 0x5a, 0x13, 0xe8, 0x35, 0x7a, 0x4e, 0xbf, 0xa9, 0xa3, 0x62, 0x05, 0xe6, 0xef,
	0x01, 0x19, 0x68, 0x85, 0xfa, 0xae, 0x30, 0xfc, 0xa9, 0x92, 0xa9, 0xc7, 0xe2, 0xbf, 0x84, 0x9d,
	0x13, 0x94, 0x0f, 0xb6, 0x78, 0xb0, 0x7d, 0x82, 0xb2, 0xf8, 0xdf, 0xfe, 0xdb, 0x71, 0x70, 0x08,
	0xad, 0xe2, 0x90, 0x88, 0x03, 0xd6, 0x78, 0xf8, 0x71, 0x38, 0xba, 0x1c, 0x76, 0x36, 0x08, 0x40,
	0x6b, 0x7c, 0x7e, 0x31, 0xfa, 0x72, 0xdc, 0x31, 0xc8, 0x16, 0xd8, 0x83, 0xd1, 0xe5, 0x50, 0x57,
	0x66, 0xff, 0xb7, 0x01, 0x2d, 0xbd, 0x67, 0x46, 0x5e, 0x80, 0xf3, 0x5e, 0x87, 0x55, 0x5c, 0xe9,
	0x16, 0xd5, 0xcf, 0xee, 0xe2, 0xe9, 0x6f, 0x90, 0xb7, 0xe0, 0x54, 0x7c, 0x90, 0xc7, 0x74, 0xdd,
	0x55, 0xf7, 0x29, 0x2d, 0xc6, 0x05, 0x2d, 0xc7, 0x05, 0x3d, 0x56, 0xe3, 0xc2, 0xdf, 0x20, 0xfb,
	0x60, 0x97, 0xfe, 0x48, 0x87, 0xd6, 0xac, 0xde, 0xed, 0xd2, 0xff, 0x06, 0x4d, 0xe5, 0x30, 0x23,
	0xaf, 0xc0, 0x5a, 0xb8, 0x25, 0x3b, 0x74, 0xd5, 0x77, 0xd7, 0xa1, 0x95, 0xa9, 0x43, 0xc1, 0x3a,
	0x8a, 0x22, 0xdd, 0x58, 0xc5, 0xef, 0xd3, 0xf2, 0xce, 0xf9, 0xda, 0x5e, 0x0e, 0xb9, 0xab, 0x96,
	0x26, 0xdf, 0xfc, 0x0b, 0x00, 0x00, 0xff, 0xff, 0xa0, 0xd6, 0x29, 0xf1, 0xf8, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// EventsClient is the client API for Events service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EventsClient interface {
	CreateEvent(ctx context.Context, in *Event, opts ...grpc.CallOption) (*Event, error)
	DeleteEvent(ctx context.Context, in *DeleteEventRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	GetEvent(ctx context.Context, in *GetEventRequest, opts ...grpc.CallOption) (*Event, error)
}

type eventsClient struct {
	cc *grpc.ClientConn
}

func NewEventsClient(cc *grpc.ClientConn) EventsClient {
	return &eventsClient{cc}
}

func (c *eventsClient) CreateEvent(ctx context.Context, in *Event, opts ...grpc.CallOption) (*Event, error) {
	out := new(Event)
	err := c.cc.Invoke(ctx, "/Events/CreateEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventsClient) DeleteEvent(ctx context.Context, in *DeleteEventRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/Events/DeleteEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventsClient) GetEvent(ctx context.Context, in *GetEventRequest, opts ...grpc.CallOption) (*Event, error) {
	out := new(Event)
	err := c.cc.Invoke(ctx, "/Events/GetEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EventsServer is the server API for Events service.
type EventsServer interface {
	CreateEvent(context.Context, *Event) (*Event, error)
	DeleteEvent(context.Context, *DeleteEventRequest) (*empty.Empty, error)
	GetEvent(context.Context, *GetEventRequest) (*Event, error)
}

// UnimplementedEventsServer can be embedded to have forward compatible implementations.
type UnimplementedEventsServer struct {
}

func (*UnimplementedEventsServer) CreateEvent(ctx context.Context, req *Event) (*Event, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateEvent not implemented")
}
func (*UnimplementedEventsServer) DeleteEvent(ctx context.Context, req *DeleteEventRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteEvent not implemented")
}
func (*UnimplementedEventsServer) GetEvent(ctx context.Context, req *GetEventRequest) (*Event, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEvent not implemented")
}

func RegisterEventsServer(s *grpc.Server, srv EventsServer) {
	s.RegisterService(&_Events_serviceDesc, srv)
}

func _Events_CreateEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Event)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventsServer).CreateEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Events/CreateEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventsServer).CreateEvent(ctx, req.(*Event))
	}
	return interceptor(ctx, in, info, handler)
}

func _Events_DeleteEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventsServer).DeleteEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Events/DeleteEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventsServer).DeleteEvent(ctx, req.(*DeleteEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Events_GetEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventsServer).GetEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Events/GetEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventsServer).GetEvent(ctx, req.(*GetEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Events_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Events",
	HandlerType: (*EventsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateEvent",
			Handler:    _Events_CreateEvent_Handler,
		},
		{
			MethodName: "DeleteEvent",
			Handler:    _Events_DeleteEvent_Handler,
		},
		{
			MethodName: "GetEvent",
			Handler:    _Events_GetEvent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sorucevap.proto",
}

// UsersClient is the client API for Users service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UsersClient interface {
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*UserRecord, error)
	AddUser(ctx context.Context, in *UserRecord, opts ...grpc.CallOption) (*empty.Empty, error)
}

type usersClient struct {
	cc *grpc.ClientConn
}

func NewUsersClient(cc *grpc.ClientConn) UsersClient {
	return &usersClient{cc}
}

func (c *usersClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*UserRecord, error) {
	out := new(UserRecord)
	err := c.cc.Invoke(ctx, "/Users/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) AddUser(ctx context.Context, in *UserRecord, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/Users/AddUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UsersServer is the server API for Users service.
type UsersServer interface {
	GetUser(context.Context, *GetUserRequest) (*UserRecord, error)
	AddUser(context.Context, *UserRecord) (*empty.Empty, error)
}

// UnimplementedUsersServer can be embedded to have forward compatible implementations.
type UnimplementedUsersServer struct {
}

func (*UnimplementedUsersServer) GetUser(ctx context.Context, req *GetUserRequest) (*UserRecord, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (*UnimplementedUsersServer) AddUser(ctx context.Context, req *UserRecord) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddUser not implemented")
}

func RegisterUsersServer(s *grpc.Server, srv UsersServer) {
	s.RegisterService(&_Users_serviceDesc, srv)
}

func _Users_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Users/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).GetUser(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_AddUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRecord)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).AddUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Users/AddUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).AddUser(ctx, req.(*UserRecord))
	}
	return interceptor(ctx, in, info, handler)
}

var _Users_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Users",
	HandlerType: (*UsersServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUser",
			Handler:    _Users_GetUser_Handler,
		},
		{
			MethodName: "AddUser",
			Handler:    _Users_AddUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sorucevap.proto",
}
