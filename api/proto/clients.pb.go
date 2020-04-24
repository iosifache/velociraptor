// Code generated by protoc-gen-go. DO NOT EDIT.
// source: clients.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
	_ "www.velocidex.com/golang/velociraptor/proto"
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

type ApiClient_IPAddressClass int32

const (
	ApiClient_UNKNOWN  ApiClient_IPAddressClass = 0
	ApiClient_INTERNAL ApiClient_IPAddressClass = 1
	ApiClient_EXTERNAL ApiClient_IPAddressClass = 2
	ApiClient_VPN      ApiClient_IPAddressClass = 3
)

var ApiClient_IPAddressClass_name = map[int32]string{
	0: "UNKNOWN",
	1: "INTERNAL",
	2: "EXTERNAL",
	3: "VPN",
}

var ApiClient_IPAddressClass_value = map[string]int32{
	"UNKNOWN":  0,
	"INTERNAL": 1,
	"EXTERNAL": 2,
	"VPN":      3,
}

func (x ApiClient_IPAddressClass) String() string {
	return proto.EnumName(ApiClient_IPAddressClass_name, int32(x))
}

func (ApiClient_IPAddressClass) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_6c7b36ecb5ad4a28, []int{2, 0}
}

type SearchClientsRequest_QueryType int32

const (
	SearchClientsRequest_VALUE SearchClientsRequest_QueryType = 0
	SearchClientsRequest_KEY   SearchClientsRequest_QueryType = 1
)

var SearchClientsRequest_QueryType_name = map[int32]string{
	0: "VALUE",
	1: "KEY",
}

var SearchClientsRequest_QueryType_value = map[string]int32{
	"VALUE": 0,
	"KEY":   1,
}

func (x SearchClientsRequest_QueryType) String() string {
	return proto.EnumName(SearchClientsRequest_QueryType_name, int32(x))
}

func (SearchClientsRequest_QueryType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_6c7b36ecb5ad4a28, []int{3, 0}
}

// GRR uses an int for client_version which is difficult to use
// semantic versioning. We use a string instead which represents the
// commit number.
type AgentInformation struct {
	Version              string   `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	BuildTime            string   `protobuf:"bytes,3,opt,name=build_time,json=buildTime,proto3" json:"build_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AgentInformation) Reset()         { *m = AgentInformation{} }
func (m *AgentInformation) String() string { return proto.CompactTextString(m) }
func (*AgentInformation) ProtoMessage()    {}
func (*AgentInformation) Descriptor() ([]byte, []int) {
	return fileDescriptor_6c7b36ecb5ad4a28, []int{0}
}

func (m *AgentInformation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AgentInformation.Unmarshal(m, b)
}
func (m *AgentInformation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AgentInformation.Marshal(b, m, deterministic)
}
func (m *AgentInformation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AgentInformation.Merge(m, src)
}
func (m *AgentInformation) XXX_Size() int {
	return xxx_messageInfo_AgentInformation.Size(m)
}
func (m *AgentInformation) XXX_DiscardUnknown() {
	xxx_messageInfo_AgentInformation.DiscardUnknown(m)
}

var xxx_messageInfo_AgentInformation proto.InternalMessageInfo

func (m *AgentInformation) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *AgentInformation) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AgentInformation) GetBuildTime() string {
	if m != nil {
		return m.BuildTime
	}
	return ""
}

// Message to carry uname information.
type Uname struct {
	System               string   `protobuf:"bytes,1,opt,name=system,proto3" json:"system,omitempty"`
	Node                 string   `protobuf:"bytes,2,opt,name=node,proto3" json:"node,omitempty"`
	Release              string   `protobuf:"bytes,3,opt,name=release,proto3" json:"release,omitempty"`
	Version              string   `protobuf:"bytes,4,opt,name=version,proto3" json:"version,omitempty"`
	Machine              string   `protobuf:"bytes,5,opt,name=machine,proto3" json:"machine,omitempty"`
	Kernel               string   `protobuf:"bytes,6,opt,name=kernel,proto3" json:"kernel,omitempty"`
	Fqdn                 string   `protobuf:"bytes,7,opt,name=fqdn,proto3" json:"fqdn,omitempty"`
	InstallDate          uint64   `protobuf:"varint,8,opt,name=install_date,json=installDate,proto3" json:"install_date,omitempty"`
	LibcVer              string   `protobuf:"bytes,9,opt,name=libc_ver,json=libcVer,proto3" json:"libc_ver,omitempty"`
	Architecture         string   `protobuf:"bytes,10,opt,name=architecture,proto3" json:"architecture,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Uname) Reset()         { *m = Uname{} }
func (m *Uname) String() string { return proto.CompactTextString(m) }
func (*Uname) ProtoMessage()    {}
func (*Uname) Descriptor() ([]byte, []int) {
	return fileDescriptor_6c7b36ecb5ad4a28, []int{1}
}

func (m *Uname) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Uname.Unmarshal(m, b)
}
func (m *Uname) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Uname.Marshal(b, m, deterministic)
}
func (m *Uname) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Uname.Merge(m, src)
}
func (m *Uname) XXX_Size() int {
	return xxx_messageInfo_Uname.Size(m)
}
func (m *Uname) XXX_DiscardUnknown() {
	xxx_messageInfo_Uname.DiscardUnknown(m)
}

var xxx_messageInfo_Uname proto.InternalMessageInfo

func (m *Uname) GetSystem() string {
	if m != nil {
		return m.System
	}
	return ""
}

func (m *Uname) GetNode() string {
	if m != nil {
		return m.Node
	}
	return ""
}

func (m *Uname) GetRelease() string {
	if m != nil {
		return m.Release
	}
	return ""
}

func (m *Uname) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *Uname) GetMachine() string {
	if m != nil {
		return m.Machine
	}
	return ""
}

func (m *Uname) GetKernel() string {
	if m != nil {
		return m.Kernel
	}
	return ""
}

func (m *Uname) GetFqdn() string {
	if m != nil {
		return m.Fqdn
	}
	return ""
}

func (m *Uname) GetInstallDate() uint64 {
	if m != nil {
		return m.InstallDate
	}
	return 0
}

func (m *Uname) GetLibcVer() string {
	if m != nil {
		return m.LibcVer
	}
	return ""
}

func (m *Uname) GetArchitecture() string {
	if m != nil {
		return m.Architecture
	}
	return ""
}

// Describe a client. We fill in some metadata about the client but
// this is by no means exhaustive.
type ApiClient struct {
	ClientId              string                   `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	AgentInformation      *AgentInformation        `protobuf:"bytes,2,opt,name=agent_information,json=agentInformation,proto3" json:"agent_information,omitempty"`
	OsInfo                *Uname                   `protobuf:"bytes,3,opt,name=os_info,json=osInfo,proto3" json:"os_info,omitempty"`
	FirstSeenAt           uint64                   `protobuf:"varint,6,opt,name=first_seen_at,json=firstSeenAt,proto3" json:"first_seen_at,omitempty"`
	LastSeenAt            uint64                   `protobuf:"varint,7,opt,name=last_seen_at,json=lastSeenAt,proto3" json:"last_seen_at,omitempty"`
	LastBootedAt          uint64                   `protobuf:"varint,8,opt,name=last_booted_at,json=lastBootedAt,proto3" json:"last_booted_at,omitempty"`
	LastClock             uint64                   `protobuf:"varint,9,opt,name=last_clock,json=lastClock,proto3" json:"last_clock,omitempty"`
	LastCrashAt           uint64                   `protobuf:"varint,10,opt,name=last_crash_at,json=lastCrashAt,proto3" json:"last_crash_at,omitempty"`
	LastIp                string                   `protobuf:"bytes,16,opt,name=last_ip,json=lastIp,proto3" json:"last_ip,omitempty"`
	LastInterrogateFlowId string                   `protobuf:"bytes,19,opt,name=last_interrogate_flow_id,json=lastInterrogateFlowId,proto3" json:"last_interrogate_flow_id,omitempty"`
	LastIpClass           ApiClient_IPAddressClass `protobuf:"varint,17,opt,name=last_ip_class,json=lastIpClass,proto3,enum=proto.ApiClient_IPAddressClass" json:"last_ip_class,omitempty"`
	Labels                []string                 `protobuf:"bytes,18,rep,name=labels,proto3" json:"labels,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}                 `json:"-"`
	XXX_unrecognized      []byte                   `json:"-"`
	XXX_sizecache         int32                    `json:"-"`
}

func (m *ApiClient) Reset()         { *m = ApiClient{} }
func (m *ApiClient) String() string { return proto.CompactTextString(m) }
func (*ApiClient) ProtoMessage()    {}
func (*ApiClient) Descriptor() ([]byte, []int) {
	return fileDescriptor_6c7b36ecb5ad4a28, []int{2}
}

func (m *ApiClient) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApiClient.Unmarshal(m, b)
}
func (m *ApiClient) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApiClient.Marshal(b, m, deterministic)
}
func (m *ApiClient) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApiClient.Merge(m, src)
}
func (m *ApiClient) XXX_Size() int {
	return xxx_messageInfo_ApiClient.Size(m)
}
func (m *ApiClient) XXX_DiscardUnknown() {
	xxx_messageInfo_ApiClient.DiscardUnknown(m)
}

var xxx_messageInfo_ApiClient proto.InternalMessageInfo

func (m *ApiClient) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

func (m *ApiClient) GetAgentInformation() *AgentInformation {
	if m != nil {
		return m.AgentInformation
	}
	return nil
}

func (m *ApiClient) GetOsInfo() *Uname {
	if m != nil {
		return m.OsInfo
	}
	return nil
}

func (m *ApiClient) GetFirstSeenAt() uint64 {
	if m != nil {
		return m.FirstSeenAt
	}
	return 0
}

func (m *ApiClient) GetLastSeenAt() uint64 {
	if m != nil {
		return m.LastSeenAt
	}
	return 0
}

func (m *ApiClient) GetLastBootedAt() uint64 {
	if m != nil {
		return m.LastBootedAt
	}
	return 0
}

func (m *ApiClient) GetLastClock() uint64 {
	if m != nil {
		return m.LastClock
	}
	return 0
}

func (m *ApiClient) GetLastCrashAt() uint64 {
	if m != nil {
		return m.LastCrashAt
	}
	return 0
}

func (m *ApiClient) GetLastIp() string {
	if m != nil {
		return m.LastIp
	}
	return ""
}

func (m *ApiClient) GetLastInterrogateFlowId() string {
	if m != nil {
		return m.LastInterrogateFlowId
	}
	return ""
}

func (m *ApiClient) GetLastIpClass() ApiClient_IPAddressClass {
	if m != nil {
		return m.LastIpClass
	}
	return ApiClient_UNKNOWN
}

func (m *ApiClient) GetLabels() []string {
	if m != nil {
		return m.Labels
	}
	return nil
}

type SearchClientsRequest struct {
	Offset               uint64                         `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit                uint64                         `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Query                string                         `protobuf:"bytes,3,opt,name=query,proto3" json:"query,omitempty"`
	NameOnly             bool                           `protobuf:"varint,4,opt,name=name_only,json=nameOnly,proto3" json:"name_only,omitempty"`
	Type                 SearchClientsRequest_QueryType `protobuf:"varint,5,opt,name=type,proto3,enum=proto.SearchClientsRequest_QueryType" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *SearchClientsRequest) Reset()         { *m = SearchClientsRequest{} }
func (m *SearchClientsRequest) String() string { return proto.CompactTextString(m) }
func (*SearchClientsRequest) ProtoMessage()    {}
func (*SearchClientsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6c7b36ecb5ad4a28, []int{3}
}

func (m *SearchClientsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchClientsRequest.Unmarshal(m, b)
}
func (m *SearchClientsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchClientsRequest.Marshal(b, m, deterministic)
}
func (m *SearchClientsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchClientsRequest.Merge(m, src)
}
func (m *SearchClientsRequest) XXX_Size() int {
	return xxx_messageInfo_SearchClientsRequest.Size(m)
}
func (m *SearchClientsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchClientsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SearchClientsRequest proto.InternalMessageInfo

func (m *SearchClientsRequest) GetOffset() uint64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *SearchClientsRequest) GetLimit() uint64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *SearchClientsRequest) GetQuery() string {
	if m != nil {
		return m.Query
	}
	return ""
}

func (m *SearchClientsRequest) GetNameOnly() bool {
	if m != nil {
		return m.NameOnly
	}
	return false
}

func (m *SearchClientsRequest) GetType() SearchClientsRequest_QueryType {
	if m != nil {
		return m.Type
	}
	return SearchClientsRequest_VALUE
}

type SearchClientsResponse struct {
	Items                []*ApiClient `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	Names                []string     `protobuf:"bytes,2,rep,name=names,proto3" json:"names,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *SearchClientsResponse) Reset()         { *m = SearchClientsResponse{} }
func (m *SearchClientsResponse) String() string { return proto.CompactTextString(m) }
func (*SearchClientsResponse) ProtoMessage()    {}
func (*SearchClientsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6c7b36ecb5ad4a28, []int{4}
}

func (m *SearchClientsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchClientsResponse.Unmarshal(m, b)
}
func (m *SearchClientsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchClientsResponse.Marshal(b, m, deterministic)
}
func (m *SearchClientsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchClientsResponse.Merge(m, src)
}
func (m *SearchClientsResponse) XXX_Size() int {
	return xxx_messageInfo_SearchClientsResponse.Size(m)
}
func (m *SearchClientsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchClientsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SearchClientsResponse proto.InternalMessageInfo

func (m *SearchClientsResponse) GetItems() []*ApiClient {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *SearchClientsResponse) GetNames() []string {
	if m != nil {
		return m.Names
	}
	return nil
}

type GetClientRequest struct {
	ClientId             string   `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	Lightweight          bool     `protobuf:"varint,2,opt,name=lightweight,proto3" json:"lightweight,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetClientRequest) Reset()         { *m = GetClientRequest{} }
func (m *GetClientRequest) String() string { return proto.CompactTextString(m) }
func (*GetClientRequest) ProtoMessage()    {}
func (*GetClientRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6c7b36ecb5ad4a28, []int{5}
}

func (m *GetClientRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetClientRequest.Unmarshal(m, b)
}
func (m *GetClientRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetClientRequest.Marshal(b, m, deterministic)
}
func (m *GetClientRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetClientRequest.Merge(m, src)
}
func (m *GetClientRequest) XXX_Size() int {
	return xxx_messageInfo_GetClientRequest.Size(m)
}
func (m *GetClientRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetClientRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetClientRequest proto.InternalMessageInfo

func (m *GetClientRequest) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

func (m *GetClientRequest) GetLightweight() bool {
	if m != nil {
		return m.Lightweight
	}
	return false
}

type LabelClientsRequest struct {
	ClientIds            []string `protobuf:"bytes,1,rep,name=client_ids,json=clientIds,proto3" json:"client_ids,omitempty"`
	Labels               []string `protobuf:"bytes,2,rep,name=labels,proto3" json:"labels,omitempty"`
	Operation            string   `protobuf:"bytes,3,opt,name=operation,proto3" json:"operation,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LabelClientsRequest) Reset()         { *m = LabelClientsRequest{} }
func (m *LabelClientsRequest) String() string { return proto.CompactTextString(m) }
func (*LabelClientsRequest) ProtoMessage()    {}
func (*LabelClientsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6c7b36ecb5ad4a28, []int{6}
}

func (m *LabelClientsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LabelClientsRequest.Unmarshal(m, b)
}
func (m *LabelClientsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LabelClientsRequest.Marshal(b, m, deterministic)
}
func (m *LabelClientsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LabelClientsRequest.Merge(m, src)
}
func (m *LabelClientsRequest) XXX_Size() int {
	return xxx_messageInfo_LabelClientsRequest.Size(m)
}
func (m *LabelClientsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LabelClientsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LabelClientsRequest proto.InternalMessageInfo

func (m *LabelClientsRequest) GetClientIds() []string {
	if m != nil {
		return m.ClientIds
	}
	return nil
}

func (m *LabelClientsRequest) GetLabels() []string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *LabelClientsRequest) GetOperation() string {
	if m != nil {
		return m.Operation
	}
	return ""
}

func init() {
	proto.RegisterEnum("proto.ApiClient_IPAddressClass", ApiClient_IPAddressClass_name, ApiClient_IPAddressClass_value)
	proto.RegisterEnum("proto.SearchClientsRequest_QueryType", SearchClientsRequest_QueryType_name, SearchClientsRequest_QueryType_value)
	proto.RegisterType((*AgentInformation)(nil), "proto.AgentInformation")
	proto.RegisterType((*Uname)(nil), "proto.Uname")
	proto.RegisterType((*ApiClient)(nil), "proto.ApiClient")
	proto.RegisterType((*SearchClientsRequest)(nil), "proto.SearchClientsRequest")
	proto.RegisterType((*SearchClientsResponse)(nil), "proto.SearchClientsResponse")
	proto.RegisterType((*GetClientRequest)(nil), "proto.GetClientRequest")
	proto.RegisterType((*LabelClientsRequest)(nil), "proto.LabelClientsRequest")
}

func init() {
	proto.RegisterFile("clients.proto", fileDescriptor_6c7b36ecb5ad4a28)
}

var fileDescriptor_6c7b36ecb5ad4a28 = []byte{
	// 1447 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x96, 0xdd, 0x72, 0x5b, 0xb7,
	0x11, 0xc7, 0x43, 0x8b, 0x22, 0x45, 0x50, 0x76, 0x19, 0xc4, 0x6e, 0xce, 0x24, 0x76, 0x8d, 0x72,
	0xea, 0x86, 0x6a, 0xe4, 0x23, 0x91, 0x52, 0x65, 0x3b, 0xed, 0x4c, 0x86, 0x94, 0x94, 0x0c, 0x63,
	0x45, 0x52, 0x8f, 0xe4, 0x8f, 0xe6, 0xa2, 0x1c, 0xf0, 0x9c, 0x25, 0x89, 0xfa, 0x10, 0xa0, 0x00,
	0x50, 0x34, 0x67, 0x72, 0xd1, 0x97, 0xe8, 0x4c, 0x6f, 0x3a, 0xbd, 0xeb, 0x93, 0xe4, 0x0d, 0xfa,
	0x06, 0xed, 0x6b, 0xb4, 0x33, 0x1d, 0x2c, 0x40, 0x49, 0x54, 0x9b, 0x1b, 0xf1, 0x00, 0xbb, 0xfb,
	0xdb, 0x3f, 0x80, 0xc5, 0x42, 0xe4, 0x6e, 0x9a, 0x0b, 0x90, 0xd6, 0xc4, 0x13, 0xad, 0xac, 0xa2,
	0xab, 0xf8, 0xf3, 0xc9, 0x7d, 0xfc, 0xd9, 0x32, 0x30, 0xe6, 0xd2, 0x8a, 0xd4, 0x1b, 0xeb, 0x3d,
	0x52, 0x6b, 0x0f, 0x41, 0xda, 0xae, 0x1c, 0x28, 0x3d, 0xe6, 0x56, 0x28, 0x49, 0x23, 0x52, 0xbe,
	0x04, 0x6d, 0x84, 0x92, 0x51, 0x81, 0x15, 0x1a, 0x95, 0x64, 0x31, 0xa4, 0x94, 0x14, 0x25, 0x1f,
	0x43, 0x74, 0x07, 0xa7, 0xf1, 0x9b, 0x3e, 0x22, 0xa4, 0x3f, 0x15, 0x79, 0xd6, 0xb3, 0x62, 0x0c,
	0xd1, 0x0a, 0x5a, 0x2a, 0x38, 0x73, 0x2e, 0xc6, 0x50, 0xff, 0x4f, 0x89, 0xac, 0xbe, 0x42, 0xc7,
	0x97, 0xa4, 0x64, 0xe6, 0xc6, 0xc2, 0xd8, 0x53, 0x3b, 0x3b, 0xff, 0xfc, 0xf7, 0xbf, 0x7e, 0x28,
	0x3c, 0xa5, 0x9f, 0x9f, 0x8f, 0x80, 0x79, 0x0b, 0x9b, 0xe4, 0xdc, 0x3a, 0x21, 0xac, 0xf1, 0x46,
	0xc8, 0x4c, 0xcd, 0xcc, 0xf7, 0x07, 0x5c, 0xcf, 0x84, 0xfc, 0xfe, 0x48, 0xc8, 0xe9, 0xfb, 0x8d,
	0x38, 0x09, 0x08, 0xfa, 0x9c, 0x14, 0xa5, 0xca, 0x82, 0x92, 0xce, 0x2f, 0x10, 0xf5, 0x33, 0xfa,
	0xd0, 0xa1, 0x46, 0xca, 0x58, 0x97, 0x90, 0xa9, 0x01, 0xb3, 0x23, 0x61, 0x02, 0x3b, 0x4e, 0x30,
	0x82, 0x9e, 0x92, 0xb2, 0x86, 0x1c, 0xb8, 0x09, 0x62, 0x3b, 0x7b, 0x18, 0xbc, 0x4d, 0x63, 0x17,
	0x7c, 0x72, 0xc6, 0x82, 0x95, 0x89, 0x0c, 0xa4, 0x15, 0x03, 0x01, 0x9a, 0x41, 0x3c, 0x8c, 0xd9,
	0xb3, 0x4d, 0x76, 0x72, 0xf6, 0x76, 0x93, 0x65, 0xd0, 0x17, 0x5c, 0xc6, 0xc9, 0x02, 0x43, 0xcf,
	0xaf, 0xf7, 0xab, 0x88, 0xc4, 0x2f, 0x90, 0xb8, 0x4b, 0x5b, 0x81, 0x18, 0xac, 0xac, 0x7b, 0xe0,
	0x49, 0x7b, 0x71, 0x33, 0x7e, 0xb6, 0xb7, 0xdd, 0x3c, 0x3b, 0x6d, 0x6e, 0xb2, 0xe6, 0x76, 0xfc,
	0x22, 0x6e, 0x6d, 0xb2, 0xe6, 0x6e, 0xbc, 0xbd, 0x1b, 0x5f, 0xef, 0xf5, 0xb7, 0xa4, 0x3c, 0xe6,
	0xe9, 0x48, 0x48, 0x88, 0x56, 0x7f, 0x74, 0xbf, 0xb8, 0x4e, 0x47, 0xc2, 0x42, 0x6a, 0xa7, 0x1a,
	0x3c, 0xbb, 0xfd, 0xed, 0xc1, 0xde, 0xee, 0x26, 0x7b, 0xff, 0x7c, 0xaf, 0xb7, 0xe7, 0x70, 0x81,
	0x41, 0xbf, 0x23, 0xa5, 0x77, 0xa0, 0x25, 0xe4, 0x51, 0x09, 0x69, 0x1d, 0xa4, 0xfd, 0x96, 0x7e,
	0xe1, 0x68, 0xde, 0x72, 0xa5, 0xd3, 0x58, 0x2d, 0xe4, 0x70, 0x59, 0xeb, 0x26, 0x6b, 0xee, 0xc4,
	0xcd, 0x78, 0x7b, 0x93, 0xed, 0xc4, 0xcd, 0x5f, 0x3f, 0xd5, 0x69, 0x2b, 0x4e, 0x02, 0x91, 0x1e,
	0x92, 0xe2, 0xe0, 0x22, 0x93, 0x51, 0x19, 0xc9, 0x4d, 0x24, 0x7f, 0x4e, 0x37, 0xae, 0x75, 0x7e,
	0x66, 0xd8, 0x60, 0x9a, 0xe7, 0x73, 0x76, 0x31, 0xe5, 0xb9, 0xdb, 0xd2, 0x8c, 0x65, 0x6a, 0xcc,
	0x85, 0x64, 0xee, 0xa0, 0xe2, 0x04, 0xc3, 0x69, 0x42, 0xd6, 0x85, 0x34, 0x96, 0xe7, 0x79, 0x2f,
	0xe3, 0x16, 0xa2, 0x35, 0x56, 0x68, 0x14, 0x3b, 0x5b, 0x88, 0xdb, 0x20, 0xd5, 0xe4, 0xe0, 0xab,
	0x03, 0x6e, 0xc1, 0x95, 0x19, 0xfd, 0xe4, 0xcd, 0x08, 0xe4, 0x62, 0x13, 0x66, 0xdc, 0xb0, 0x10,
	0x08, 0x59, 0x9c, 0x54, 0xc3, 0xb7, 0x73, 0xa6, 0xcf, 0xc9, 0x5a, 0x2e, 0xfa, 0x69, 0xef, 0x12,
	0x74, 0x54, 0x41, 0x79, 0x8f, 0x90, 0xf7, 0x31, 0x7d, 0xe0, 0xe4, 0xed, 0xb3, 0x5c, 0xf4, 0x35,
	0xd7, 0xf3, 0xc5, 0xda, 0x93, 0xb2, 0x73, 0x7f, 0x0d, 0x9a, 0xfe, 0x50, 0x20, 0xeb, 0x37, 0xb7,
	0x37, 0x22, 0x18, 0xfe, 0xb7, 0x02, 0xc6, 0xff, 0xa5, 0x40, 0xff, 0x5c, 0x70, 0x84, 0xa5, 0x13,
	0x58, 0x54, 0x5c, 0x5f, 0x48, 0xae, 0xe7, 0x31, 0x6b, 0x1c, 0x2b, 0x0b, 0x7e, 0x2a, 0xe5, 0x92,
	0xf5, 0x81, 0x65, 0x62, 0x30, 0x00, 0x0d, 0xd2, 0xb2, 0x81, 0x56, 0x63, 0x66, 0x47, 0xc0, 0xc2,
	0x09, 0x2d, 0x93, 0x84, 0x44, 0x5b, 0xea, 0x0a, 0x51, 0x0d, 0x18, 0x67, 0x3b, 0x2d, 0xd6, 0x17,
	0x36, 0x90, 0x99, 0x9e, 0x4a, 0xe9, 0x8e, 0x48, 0x49, 0xc6, 0xd9, 0xde, 0x2e, 0x9a, 0xfc, 0x6e,
	0x6c, 0x24, 0x4b, 0xaa, 0xeb, 0x7f, 0x2d, 0x93, 0x4a, 0x7b, 0x22, 0xf6, 0xb1, 0x25, 0xd0, 0x2f,
	0x49, 0xc5, 0x37, 0x87, 0x9e, 0xc8, 0xc2, 0x35, 0xac, 0xe3, 0x7a, 0x1e, 0x92, 0xea, 0x95, 0x57,
	0x37, 0xa3, 0x77, 0xdd, 0xd2, 0xbc, 0x27, 0x13, 0x59, 0xb2, 0x96, 0x2e, 0x0c, 0x07, 0xe4, 0x43,
	0x3e, 0xc4, 0xf8, 0xeb, 0x86, 0x81, 0x97, 0xb0, 0xda, 0xfa, 0xd8, 0xb7, 0x94, 0xf8, 0x76, 0x3f,
	0x49, 0x6a, 0xfc, 0x76, 0x87, 0x79, 0x42, 0xca, 0xca, 0x20, 0x02, 0xef, 0x60, 0xb5, 0xb5, 0x1e,
	0x62, 0xb1, 0x53, 0x24, 0x25, 0x65, 0x9c, 0x37, 0xb5, 0xe4, 0xee, 0x40, 0x68, 0x63, 0x7b, 0x06,
	0x40, 0xf6, 0xb8, 0xc5, 0xd2, 0x2d, 0x76, 0x4e, 0x51, 0xf1, 0x37, 0xcb, 0x15, 0xf1, 0x1b, 0xac,
	0x08, 0x7b, 0x2d, 0xdb, 0x55, 0x05, 0x46, 0x33, 0x17, 0xcd, 0x1a, 0x22, 0x86, 0x98, 0xcd, 0x9c,
	0x93, 0xf0, 0x46, 0x90, 0x5a, 0xb9, 0x8a, 0xd9, 0x88, 0x93, 0x2a, 0x3a, 0x9e, 0x01, 0xc8, 0xb6,
	0xa5, 0x6f, 0xc9, 0x7a, 0xce, 0x6f, 0x24, 0x2d, 0x63, 0xd2, 0xd0, 0x25, 0x96, 0x93, 0xfe, 0xfc,
	0x88, 0x1b, 0xcb, 0xdc, 0xa7, 0x27, 0x87, 0xd4, 0xe9, 0x08, 0xd2, 0x77, 0x90, 0x31, 0x21, 0xe3,
	0x84, 0x38, 0x56, 0x20, 0x7f, 0x43, 0xee, 0x21, 0xb9, 0xaf, 0x94, 0x85, 0xcc, 0xb1, 0x7d, 0x89,
	0x87, 0xf6, 0xb5, 0xcc, 0xfe, 0x09, 0xb2, 0x9d, 0x2b, 0x26, 0x88, 0x13, 0x54, 0xd5, 0xc1, 0xd0,
	0xb6, 0xa5, 0x7f, 0x20, 0x48, 0xee, 0xa5, 0xb9, 0x4a, 0xdf, 0x61, 0x69, 0x17, 0x3b, 0x5f, 0x22,
	0xe7, 0xc5, 0x32, 0xe7, 0x57, 0xfb, 0x41, 0x94, 0x73, 0x34, 0xec, 0x92, 0xe7, 0x53, 0x60, 0xd9,
	0x14, 0x6f, 0x78, 0xce, 0x2d, 0x98, 0xa0, 0xd7, 0x89, 0xad, 0x38, 0xe4, 0xbe, 0x73, 0xa4, 0x5d,
	0x72, 0xd7, 0xf3, 0x35, 0x37, 0x23, 0x27, 0x95, 0x60, 0x8a, 0x27, 0x98, 0xe2, 0xf1, 0x72, 0x8a,
	0x1a, 0x4a, 0x45, 0xcf, 0xa0, 0xb5, 0x8a, 0x20, 0x37, 0xd1, 0xb6, 0xb4, 0x4d, 0xca, 0x88, 0x12,
	0x93, 0xa8, 0x86, 0x25, 0xd7, 0x40, 0x48, 0x9d, 0x32, 0x57, 0x65, 0xce, 0xe4, 0x8f, 0x48, 0xc3,
	0xd8, 0x5d, 0x97, 0xf6, 0x69, 0x97, 0xf1, 0x2c, 0xd3, 0x60, 0x4c, 0x52, 0x72, 0xd6, 0xee, 0x84,
	0x3e, 0x23, 0x91, 0x47, 0x48, 0x0b, 0x5a, 0xab, 0x21, 0xb7, 0xd0, 0x1b, 0xe4, 0x6a, 0xe6, 0xca,
	0xf8, 0x23, 0x7c, 0x72, 0x1e, 0xa0, 0xe7, 0xb5, 0xf9, 0xab, 0x5c, 0xcd, 0xba, 0x19, 0xdd, 0x0f,
	0xcb, 0x10, 0x93, 0x5e, 0x9a, 0x73, 0x63, 0xa2, 0x0f, 0x59, 0xa1, 0x71, 0xaf, 0xf5, 0x78, 0x51,
	0xab, 0x8b, 0x9a, 0x8f, 0xbb, 0xa7, 0x6d, 0x9f, 0x73, 0xdf, 0xb9, 0xf9, 0x05, 0x74, 0x27, 0x38,
	0xa0, 0x3f, 0x25, 0xa5, 0x9c, 0xf7, 0x21, 0x37, 0x11, 0x65, 0x2b, 0x8d, 0x4a, 0x12, 0x46, 0xf5,
	0x0e, 0xb9, 0xb7, 0x1c, 0x46, 0xab, 0xa4, 0xfc, 0xea, 0xf8, 0xe5, 0xf1, 0xc9, 0x9b, 0xe3, 0xda,
	0x07, 0x74, 0x9d, 0xac, 0x75, 0x8f, 0xcf, 0x0f, 0x93, 0xe3, 0xf6, 0x51, 0xad, 0xe0, 0x46, 0x87,
	0x6f, 0xc3, 0xe8, 0x0e, 0x2d, 0x93, 0x95, 0xd7, 0xa7, 0xc7, 0xb5, 0x95, 0xfa, 0x3f, 0x0a, 0xe4,
	0xfe, 0x19, 0xb8, 0x2b, 0xeb, 0x85, 0x98, 0x04, 0x2e, 0xa6, 0x60, 0xac, 0x4b, 0xaa, 0x06, 0x03,
	0x03, 0x16, 0xef, 0x69, 0x31, 0x09, 0x23, 0x7a, 0x9f, 0xac, 0xe6, 0x62, 0x2c, 0x2c, 0xde, 0xba,
	0x62, 0xe2, 0x07, 0x6e, 0xf6, 0x62, 0x0a, 0x7a, 0x1e, 0x1e, 0x60, 0x3f, 0xa0, 0x9f, 0x92, 0x8a,
	0xbb, 0x50, 0x3d, 0x25, 0xf3, 0x39, 0xbe, 0x4d, 0x6b, 0xc9, 0x9a, 0x9b, 0x38, 0x91, 0xf9, 0x9c,
	0xbe, 0x20, 0x45, 0x3b, 0x9f, 0xf8, 0xd7, 0xe5, 0x5e, 0xeb, 0x49, 0xd8, 0x91, 0xff, 0xa7, 0x25,
	0xfe, 0x9d, 0xa3, 0x9d, 0xcf, 0x27, 0x90, 0x60, 0x48, 0xfd, 0x31, 0xa9, 0x5c, 0x4d, 0xd1, 0x0a,
	0x59, 0x7d, 0xdd, 0x3e, 0x7a, 0x75, 0x58, 0xfb, 0xc0, 0xad, 0xea, 0xe5, 0xe1, 0xef, 0x6b, 0x85,
	0xfa, 0xdf, 0x0b, 0xe4, 0xc1, 0x2d, 0x92, 0x99, 0x28, 0x69, 0x80, 0xfe, 0x92, 0xac, 0x0a, 0x0b,
	0x63, 0x13, 0x15, 0xd8, 0x4a, 0xa3, 0xda, 0xaa, 0xdd, 0x3e, 0x88, 0xc4, 0x9b, 0x29, 0x90, 0x55,
	0xa7, 0xd4, 0x44, 0x77, 0xdc, 0x96, 0x77, 0x4e, 0xb0, 0x64, 0xba, 0xf4, 0xeb, 0xee, 0x80, 0x5d,
	0x2d, 0x89, 0xb9, 0xc7, 0x7d, 0x02, 0xa9, 0x7f, 0x51, 0x42, 0xcb, 0xd4, 0x5e, 0x33, 0x9b, 0x01,
	0x43, 0x1f, 0x0d, 0x76, 0xaa, 0xbd, 0x09, 0x81, 0x6c, 0x04, 0x1a, 0xe2, 0xc4, 0xd3, 0xeb, 0x7f,
	0x2a, 0x90, 0xda, 0xd7, 0x60, 0x43, 0xee, 0xb0, 0xf5, 0x9f, 0xfe, 0x4f, 0x97, 0xbc, 0xd1, 0x01,
	0x4f, 0x48, 0x35, 0x17, 0xc3, 0x91, 0x9d, 0x81, 0xfb, 0x8b, 0xa7, 0xb0, 0xd6, 0x79, 0x8a, 0xf2,
	0x3e, 0xa3, 0x4f, 0xba, 0x03, 0x66, 0xc0, 0xfa, 0xcc, 0xa9, 0x6b, 0x2f, 0xa9, 0x65, 0x46, 0x8d,
	0x5d, 0x43, 0xbf, 0x6a, 0x7f, 0xee, 0x7a, 0x5c, 0x13, 0xea, 0x7f, 0x24, 0x1f, 0x1d, 0xb9, 0x7a,
	0xba, 0x75, 0xfe, 0x8f, 0x08, 0xb9, 0x12, 0xe1, 0x77, 0xab, 0x92, 0x54, 0x16, 0x2a, 0x6e, 0xd6,
	0xe4, 0x9d, 0x9b, 0x35, 0x49, 0x1f, 0x92, 0x8a, 0x9a, 0x80, 0xf6, 0x8d, 0x39, 0xfc, 0x37, 0x76,
	0x35, 0xd1, 0x69, 0x7e, 0xb7, 0x35, 0x9b, 0xcd, 0xe2, 0x4b, 0xc8, 0x55, 0x2a, 0x32, 0x78, 0x1f,
	0xa7, 0x6a, 0xbc, 0x35, 0x54, 0x39, 0x97, 0xc3, 0x2d, 0x3f, 0xa9, 0xf9, 0xc4, 0x2a, 0xbd, 0xc5,
	0x27, 0x62, 0x0b, 0x4f, 0xa6, 0x5f, 0xc2, 0x9f, 0x9d, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0x5a,
	0x62, 0x55, 0x8e, 0x56, 0x0a, 0x00, 0x00,
}
