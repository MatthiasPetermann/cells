// Code generated by protoc-gen-go. DO NOT EDIT.
// source: broker.proto

/*
Package rest is a generated protocol buffer package.

It is generated from these files:
	broker.proto
	common.proto
	config.proto
	data.proto
	frontend.proto
	graph.proto
	idm.proto
	rest.proto
	scheduler.proto
	share.proto
	templates.proto

It has these top-level messages:
	ActivitiesCollection
	SubscriptionsCollection
	LogCollection
	LogMessageCollection
	TimeRangeResultCollection
	DeleteResponse
	Error
	Configuration
	ListDataSourceRequest
	DataSourceCollection
	DeleteDataSourceResponse
	ListPeersAddressesRequest
	ListPeersAddressesResponse
	ListPeerFoldersRequest
	CreatePeerFolderRequest
	CreatePeerFolderResponse
	ListStorageBucketsRequest
	Process
	ListProcessesRequest
	ListProcessesResponse
	ListVersioningPolicyRequest
	VersioningPolicyCollection
	ListVirtualNodesRequest
	ListServiceRequest
	ServiceCollection
	ControlServiceRequest
	DiscoveryRequest
	DiscoveryResponse
	ConfigFormRequest
	OpenApiResponse
	ActionDescription
	SchedulerActionsRequest
	SchedulerActionsResponse
	SchedulerActionFormRequest
	SchedulerActionFormResponse
	ListSitesRequest
	ListSitesResponse
	SearchResults
	Pagination
	Metadata
	MetaCollection
	MetaNamespaceRequest
	GetBulkMetaRequest
	BulkMetaResponse
	HeadNodeRequest
	HeadNodeResponse
	CreateNodesRequest
	CreateSelectionRequest
	CreateSelectionResponse
	NodesCollection
	DeleteNodesRequest
	BackgroundJobResult
	DeleteNodesResponse
	RestoreNodesRequest
	RestoreNodesResponse
	ListDocstoreRequest
	DocstoreCollection
	SettingsMenuRequest
	SettingsEntryMeta
	SettingsEntry
	SettingsAccess
	SettingsAccessRestPolicy
	SettingsSection
	SettingsMenuResponse
	FrontStateRequest
	FrontStateResponse
	FrontPluginsRequest
	FrontPluginsResponse
	FrontMessagesRequest
	FrontMessagesResponse
	FrontSessionGetRequest
	FrontSessionGetResponse
	FrontSessionRequest
	FrontSessionResponse
	FrontSessionDelRequest
	FrontSessionDelResponse
	FrontAuthRequest
	FrontAuthResponse
	FrontEnrollAuthRequest
	FrontEnrollAuthResponse
	FrontBinaryRequest
	FrontBinaryResponse
	FrontBootConfRequest
	FrontBootConfResponse
	UserStateRequest
	UserStateResponse
	RelationRequest
	RelationResponse
	ResourcePolicyQuery
	SearchRoleRequest
	RolesCollection
	SearchUserRequest
	UsersCollection
	BindResponse
	SearchACLRequest
	ACLCollection
	SearchWorkspaceRequest
	WorkspaceCollection
	UserMetaCollection
	UserMetaNamespaceCollection
	ListUserMetaTagsRequest
	ListUserMetaTagsResponse
	PutUserMetaTagRequest
	PutUserMetaTagResponse
	DeleteUserMetaTagsRequest
	DeleteUserMetaTagsResponse
	UserBookmarksRequest
	RevokeRequest
	RevokeResponse
	ResetPasswordTokenRequest
	ResetPasswordTokenResponse
	ResetPasswordRequest
	ResetPasswordResponse
	UserJobRequest
	UserJobResponse
	UserJobsCollection
	CellAcl
	Cell
	ShareLinkTargetUser
	ShareLink
	PutCellRequest
	GetCellRequest
	DeleteCellRequest
	DeleteCellResponse
	GetShareLinkRequest
	PutShareLinkRequest
	DeleteShareLinkRequest
	DeleteShareLinkResponse
	ListSharedResourcesRequest
	ListSharedResourcesResponse
	UpdateSharePoliciesRequest
	UpdateSharePoliciesResponse
	TemplateNode
	Template
	ListTemplatesRequest
	ListTemplatesResponse
	CreateFromTemplateRequest
	CreateFromTemplateResponse
*/
package rest

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import activity "github.com/pydio/cells/common/proto/activity"
import log "github.com/pydio/cells/common/proto/log"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Collection of Activities
type ActivitiesCollection struct {
	Activities []*activity.Object `protobuf:"bytes,1,rep,name=activities" json:"activities,omitempty"`
}

func (m *ActivitiesCollection) Reset()                    { *m = ActivitiesCollection{} }
func (m *ActivitiesCollection) String() string            { return proto.CompactTextString(m) }
func (*ActivitiesCollection) ProtoMessage()               {}
func (*ActivitiesCollection) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ActivitiesCollection) GetActivities() []*activity.Object {
	if m != nil {
		return m.Activities
	}
	return nil
}

type SubscriptionsCollection struct {
	Subscriptions []*activity.Subscription `protobuf:"bytes,1,rep,name=subscriptions" json:"subscriptions,omitempty"`
}

func (m *SubscriptionsCollection) Reset()                    { *m = SubscriptionsCollection{} }
func (m *SubscriptionsCollection) String() string            { return proto.CompactTextString(m) }
func (*SubscriptionsCollection) ProtoMessage()               {}
func (*SubscriptionsCollection) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *SubscriptionsCollection) GetSubscriptions() []*activity.Subscription {
	if m != nil {
		return m.Subscriptions
	}
	return nil
}

// Collection of serialized log messages
type LogCollection struct {
	Lines []*log.Log `protobuf:"bytes,1,rep,name=lines" json:"lines,omitempty"`
}

func (m *LogCollection) Reset()                    { *m = LogCollection{} }
func (m *LogCollection) String() string            { return proto.CompactTextString(m) }
func (*LogCollection) ProtoMessage()               {}
func (*LogCollection) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *LogCollection) GetLines() []*log.Log {
	if m != nil {
		return m.Lines
	}
	return nil
}

// Collection of serialized log messages
type LogMessageCollection struct {
	Logs []*log.LogMessage `protobuf:"bytes,1,rep,name=Logs" json:"Logs,omitempty"`
}

func (m *LogMessageCollection) Reset()                    { *m = LogMessageCollection{} }
func (m *LogMessageCollection) String() string            { return proto.CompactTextString(m) }
func (*LogMessageCollection) ProtoMessage()               {}
func (*LogMessageCollection) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *LogMessageCollection) GetLogs() []*log.LogMessage {
	if m != nil {
		return m.Logs
	}
	return nil
}

// Collection of serialized aggregated result of time range request
// with a cursor to ease navigation implementation
type TimeRangeResultCollection struct {
	Results []*log.TimeRangeResult `protobuf:"bytes,1,rep,name=Results" json:"Results,omitempty"`
	Links   []*log.TimeRangeCursor `protobuf:"bytes,2,rep,name=Links" json:"Links,omitempty"`
}

func (m *TimeRangeResultCollection) Reset()                    { *m = TimeRangeResultCollection{} }
func (m *TimeRangeResultCollection) String() string            { return proto.CompactTextString(m) }
func (*TimeRangeResultCollection) ProtoMessage()               {}
func (*TimeRangeResultCollection) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *TimeRangeResultCollection) GetResults() []*log.TimeRangeResult {
	if m != nil {
		return m.Results
	}
	return nil
}

func (m *TimeRangeResultCollection) GetLinks() []*log.TimeRangeCursor {
	if m != nil {
		return m.Links
	}
	return nil
}

func init() {
	proto.RegisterType((*ActivitiesCollection)(nil), "rest.ActivitiesCollection")
	proto.RegisterType((*SubscriptionsCollection)(nil), "rest.SubscriptionsCollection")
	proto.RegisterType((*LogCollection)(nil), "rest.LogCollection")
	proto.RegisterType((*LogMessageCollection)(nil), "rest.LogMessageCollection")
	proto.RegisterType((*TimeRangeResultCollection)(nil), "rest.TimeRangeResultCollection")
}

func init() { proto.RegisterFile("broker.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 295 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xcf, 0x4b, 0xf3, 0x30,
	0x18, 0xc7, 0x79, 0x5f, 0x37, 0x95, 0x47, 0x87, 0x32, 0x86, 0xce, 0x1d, 0x44, 0xea, 0x45, 0x3c,
	0x24, 0xfe, 0x38, 0xea, 0x65, 0xec, 0xe2, 0xa1, 0x22, 0x44, 0xc1, 0x73, 0x1b, 0x1f, 0x62, 0x5c,
	0xda, 0xa7, 0x24, 0xa9, 0xb2, 0xff, 0x5e, 0xda, 0xac, 0x35, 0x8a, 0x07, 0x0f, 0x81, 0xf0, 0x7d,
	0x3e, 0xdf, 0x4f, 0x12, 0x02, 0xbb, 0xb9, 0xa5, 0x25, 0x5a, 0x56, 0x59, 0xf2, 0x34, 0x1e, 0x58,
	0x74, 0x7e, 0x36, 0x57, 0xda, 0xbf, 0xd6, 0x39, 0x93, 0x54, 0xf0, 0x6a, 0xf5, 0xa2, 0x89, 0x4b,
	0x34, 0xc6, 0x71, 0x49, 0x45, 0x41, 0x25, 0x6f, 0x51, 0x9e, 0x49, 0xaf, 0xdf, 0xb5, 0x5f, 0xf5,
	0x1b, 0xe7, 0x2d, 0x66, 0x45, 0x10, 0xcd, 0x2e, 0xff, 0xa2, 0x30, 0xa4, 0x9a, 0x15, 0x2a, 0xc9,
	0x1d, 0x4c, 0xe6, 0x41, 0xa5, 0xd1, 0x2d, 0xc8, 0x18, 0x94, 0x5e, 0x53, 0x39, 0xbe, 0x00, 0xc8,
	0xfa, 0x7c, 0xfa, 0xef, 0x64, 0xe3, 0x6c, 0xe7, 0x6a, 0x9f, 0x75, 0xa7, 0xb2, 0x87, 0xfc, 0x0d,
	0xa5, 0x17, 0x11, 0x93, 0x3c, 0xc3, 0xe1, 0x63, 0x9d, 0x3b, 0x69, 0x75, 0xd5, 0x18, 0x62, 0xd9,
	0x2d, 0x8c, 0x5c, 0x3c, 0x5a, 0xfb, 0x0e, 0xbe, 0x7c, 0x71, 0x53, 0x7c, 0x87, 0x13, 0x0e, 0xa3,
	0x94, 0x54, 0xa4, 0x3b, 0x86, 0xa1, 0xd1, 0x65, 0x7f, 0xad, 0x6d, 0xd6, 0x3c, 0x27, 0x25, 0x25,
	0x42, 0x9c, 0xdc, 0xc0, 0x24, 0x25, 0x75, 0x8f, 0xce, 0x65, 0x0a, 0xa3, 0xde, 0x29, 0x0c, 0x52,
	0x52, 0x5d, 0x6d, 0xaf, 0xab, 0xad, 0x41, 0xd1, 0x0e, 0x93, 0x0f, 0x38, 0x7a, 0xd2, 0x05, 0x8a,
	0xac, 0x54, 0x28, 0xd0, 0xd5, 0xc6, 0x47, 0x06, 0x06, 0x5b, 0x21, 0xeb, 0x24, 0x93, 0x56, 0xf2,
	0xa3, 0x20, 0x3a, 0x68, 0x7c, 0x0e, 0xc3, 0x54, 0x97, 0x4b, 0x37, 0xfd, 0xff, 0x1b, 0xbd, 0xa8,
	0xad, 0x23, 0x2b, 0x02, 0x92, 0x6f, 0xb6, 0x1f, 0x72, 0xfd, 0x19, 0x00, 0x00, 0xff, 0xff, 0x9c,
	0xda, 0xdc, 0x76, 0x1c, 0x02, 0x00, 0x00,
}
