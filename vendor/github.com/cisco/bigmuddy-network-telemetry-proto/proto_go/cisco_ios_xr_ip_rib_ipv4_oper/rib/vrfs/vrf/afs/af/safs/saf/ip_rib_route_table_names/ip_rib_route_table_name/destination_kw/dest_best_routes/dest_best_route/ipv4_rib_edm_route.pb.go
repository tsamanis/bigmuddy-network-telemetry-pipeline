// Code generated by protoc-gen-go.
// source: ipv4_rib_edm_route.proto
// DO NOT EDIT!

/*
Package cisco_ios_xr_ip_rib_ipv4_oper_rib_vrfs_vrf_afs_af_safs_saf_ip_rib_route_table_names_ip_rib_route_table_name_destination_kw_dest_best_routes_dest_best_route is a generated protocol buffer package.

It is generated from these files:
	ipv4_rib_edm_route.proto

It has these top-level messages:
	Ipv4RibEdmRoute_KEYS
	Ipv4RibEdmRoute
	Ipv4RibEdmPath
	Ipv4RibEdmPathItem
	RibEdmLocalLabel
*/
package cisco_ios_xr_ip_rib_ipv4_oper_rib_vrfs_vrf_afs_af_safs_saf_ip_rib_route_table_names_ip_rib_route_table_name_destination_kw_dest_best_routes_dest_best_route

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

// Information of a rib route head and rib proto route
type Ipv4RibEdmRoute_KEYS struct {
	VrfName        string `protobuf:"bytes,1,opt,name=vrf_name,json=vrfName" json:"vrf_name,omitempty"`
	AfName         string `protobuf:"bytes,2,opt,name=af_name,json=afName" json:"af_name,omitempty"`
	SafName        string `protobuf:"bytes,3,opt,name=saf_name,json=safName" json:"saf_name,omitempty"`
	RouteTableName string `protobuf:"bytes,4,opt,name=route_table_name,json=routeTableName" json:"route_table_name,omitempty"`
	Address        string `protobuf:"bytes,5,opt,name=address" json:"address,omitempty"`
}

func (m *Ipv4RibEdmRoute_KEYS) Reset()                    { *m = Ipv4RibEdmRoute_KEYS{} }
func (m *Ipv4RibEdmRoute_KEYS) String() string            { return proto.CompactTextString(m) }
func (*Ipv4RibEdmRoute_KEYS) ProtoMessage()               {}
func (*Ipv4RibEdmRoute_KEYS) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Ipv4RibEdmRoute_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

func (m *Ipv4RibEdmRoute_KEYS) GetAfName() string {
	if m != nil {
		return m.AfName
	}
	return ""
}

func (m *Ipv4RibEdmRoute_KEYS) GetSafName() string {
	if m != nil {
		return m.SafName
	}
	return ""
}

func (m *Ipv4RibEdmRoute_KEYS) GetRouteTableName() string {
	if m != nil {
		return m.RouteTableName
	}
	return ""
}

func (m *Ipv4RibEdmRoute_KEYS) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type Ipv4RibEdmRoute struct {
	// Route prefix
	Prefix string `protobuf:"bytes,50,opt,name=prefix" json:"prefix,omitempty"`
	// Length of prefix
	PrefixLength uint32 `protobuf:"varint,51,opt,name=prefix_length,json=prefixLength" json:"prefix_length,omitempty"`
	// Route version
	RouteVersion uint32 `protobuf:"varint,52,opt,name=route_version,json=routeVersion" json:"route_version,omitempty"`
	// Protocol advertising the route
	ProtocolId uint32 `protobuf:"varint,53,opt,name=protocol_id,json=protocolId" json:"protocol_id,omitempty"`
	//  Name of Protocol
	ProtocolName string `protobuf:"bytes,54,opt,name=protocol_name,json=protocolName" json:"protocol_name,omitempty"`
	// Instance name
	Instance string `protobuf:"bytes,55,opt,name=instance" json:"instance,omitempty"`
	// Client adding the route to RIB
	ClientId uint32 `protobuf:"varint,56,opt,name=client_id,json=clientId" json:"client_id,omitempty"`
	// Route type
	RouteType uint32 `protobuf:"varint,57,opt,name=route_type,json=routeType" json:"route_type,omitempty"`
	// Route priority
	Priority uint32 `protobuf:"varint,58,opt,name=priority" json:"priority,omitempty"`
	// SVD Type of route
	SvdType uint32 `protobuf:"varint,59,opt,name=svd_type,json=svdType" json:"svd_type,omitempty"`
	// Route flags
	Flags uint32 `protobuf:"varint,60,opt,name=flags" json:"flags,omitempty"`
	// Extended Route flags
	ExtendedFlags uint64 `protobuf:"varint,61,opt,name=extended_flags,json=extendedFlags" json:"extended_flags,omitempty"`
	// Opaque proto specific info
	Tag uint32 `protobuf:"varint,62,opt,name=tag" json:"tag,omitempty"`
	// Distance of the route
	Distance uint32 `protobuf:"varint,63,opt,name=distance" json:"distance,omitempty"`
	// Diversion distance of the route
	DiversionDistance uint32 `protobuf:"varint,64,opt,name=diversion_distance,json=diversionDistance" json:"diversion_distance,omitempty"`
	// Route metric
	Metric uint32 `protobuf:"varint,65,opt,name=metric" json:"metric,omitempty"`
	// Number of paths
	PathsCount uint32 `protobuf:"varint,66,opt,name=paths_count,json=pathsCount" json:"paths_count,omitempty"`
	// BGP Attribute ID
	AttributeIdentity uint32 `protobuf:"varint,67,opt,name=attribute_identity,json=attributeIdentity" json:"attribute_identity,omitempty"`
	// BGP Traffic Index
	TrafficIndex uint32 `protobuf:"varint,68,opt,name=traffic_index,json=trafficIndex" json:"traffic_index,omitempty"`
	// Route ip precedence
	RoutePrecedence uint32 `protobuf:"varint,69,opt,name=route_precedence,json=routePrecedence" json:"route_precedence,omitempty"`
	// Route qos group
	QosGroup uint32 `protobuf:"varint,70,opt,name=qos_group,json=qosGroup" json:"qos_group,omitempty"`
	// Flow tag
	FlowTag uint32 `protobuf:"varint,71,opt,name=flow_tag,json=flowTag" json:"flow_tag,omitempty"`
	// Forward Class
	FwdClass uint32 `protobuf:"varint,72,opt,name=fwd_class,json=fwdClass" json:"fwd_class,omitempty"`
	// Number of pic paths in this route
	PicCount uint32 `protobuf:"varint,73,opt,name=pic_count,json=picCount" json:"pic_count,omitempty"`
	// Is the route active or backup
	Active bool `protobuf:"varint,74,opt,name=active" json:"active,omitempty"`
	// Route has a diversion path
	Diversion bool `protobuf:"varint,75,opt,name=diversion" json:"diversion,omitempty"`
	// Diversion route protocol name
	DiversionProtoName string `protobuf:"bytes,76,opt,name=diversion_proto_name,json=diversionProtoName" json:"diversion_proto_name,omitempty"`
	// Age of route (seconds)
	RouteAge uint32 `protobuf:"varint,77,opt,name=route_age,json=routeAge" json:"route_age,omitempty"`
	// Local label of the route
	RouteLabel uint32 `protobuf:"varint,78,opt,name=route_label,json=routeLabel" json:"route_label,omitempty"`
	// Route Version
	Version uint32 `protobuf:"varint,79,opt,name=version" json:"version,omitempty"`
	// Table Version
	TblVersion uint64 `protobuf:"varint,80,opt,name=tbl_version,json=tblVersion" json:"tbl_version,omitempty"`
	// Route modification time(nanoseconds)
	RouteModifyTime uint64 `protobuf:"varint,81,opt,name=route_modify_time,json=routeModifyTime" json:"route_modify_time,omitempty"`
	// Path(s) of the route
	RoutePath *Ipv4RibEdmPath `protobuf:"bytes,82,opt,name=route_path,json=routePath" json:"route_path,omitempty"`
}

func (m *Ipv4RibEdmRoute) Reset()                    { *m = Ipv4RibEdmRoute{} }
func (m *Ipv4RibEdmRoute) String() string            { return proto.CompactTextString(m) }
func (*Ipv4RibEdmRoute) ProtoMessage()               {}
func (*Ipv4RibEdmRoute) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Ipv4RibEdmRoute) GetPrefix() string {
	if m != nil {
		return m.Prefix
	}
	return ""
}

func (m *Ipv4RibEdmRoute) GetPrefixLength() uint32 {
	if m != nil {
		return m.PrefixLength
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetRouteVersion() uint32 {
	if m != nil {
		return m.RouteVersion
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetProtocolId() uint32 {
	if m != nil {
		return m.ProtocolId
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetProtocolName() string {
	if m != nil {
		return m.ProtocolName
	}
	return ""
}

func (m *Ipv4RibEdmRoute) GetInstance() string {
	if m != nil {
		return m.Instance
	}
	return ""
}

func (m *Ipv4RibEdmRoute) GetClientId() uint32 {
	if m != nil {
		return m.ClientId
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetRouteType() uint32 {
	if m != nil {
		return m.RouteType
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetPriority() uint32 {
	if m != nil {
		return m.Priority
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetSvdType() uint32 {
	if m != nil {
		return m.SvdType
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetFlags() uint32 {
	if m != nil {
		return m.Flags
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetExtendedFlags() uint64 {
	if m != nil {
		return m.ExtendedFlags
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetTag() uint32 {
	if m != nil {
		return m.Tag
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetDistance() uint32 {
	if m != nil {
		return m.Distance
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetDiversionDistance() uint32 {
	if m != nil {
		return m.DiversionDistance
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetMetric() uint32 {
	if m != nil {
		return m.Metric
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetPathsCount() uint32 {
	if m != nil {
		return m.PathsCount
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetAttributeIdentity() uint32 {
	if m != nil {
		return m.AttributeIdentity
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetTrafficIndex() uint32 {
	if m != nil {
		return m.TrafficIndex
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetRoutePrecedence() uint32 {
	if m != nil {
		return m.RoutePrecedence
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetQosGroup() uint32 {
	if m != nil {
		return m.QosGroup
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetFlowTag() uint32 {
	if m != nil {
		return m.FlowTag
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetFwdClass() uint32 {
	if m != nil {
		return m.FwdClass
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetPicCount() uint32 {
	if m != nil {
		return m.PicCount
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetActive() bool {
	if m != nil {
		return m.Active
	}
	return false
}

func (m *Ipv4RibEdmRoute) GetDiversion() bool {
	if m != nil {
		return m.Diversion
	}
	return false
}

func (m *Ipv4RibEdmRoute) GetDiversionProtoName() string {
	if m != nil {
		return m.DiversionProtoName
	}
	return ""
}

func (m *Ipv4RibEdmRoute) GetRouteAge() uint32 {
	if m != nil {
		return m.RouteAge
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetRouteLabel() uint32 {
	if m != nil {
		return m.RouteLabel
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetTblVersion() uint64 {
	if m != nil {
		return m.TblVersion
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetRouteModifyTime() uint64 {
	if m != nil {
		return m.RouteModifyTime
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetRoutePath() *Ipv4RibEdmPath {
	if m != nil {
		return m.RoutePath
	}
	return nil
}

// Information of a rib path
type Ipv4RibEdmPath struct {
	// Next path
	Ipv4RibEdmPath []*Ipv4RibEdmPathItem `protobuf:"bytes,1,rep,name=ipv4_rib_edm_path,json=ipv4RibEdmPath" json:"ipv4_rib_edm_path,omitempty"`
}

func (m *Ipv4RibEdmPath) Reset()                    { *m = Ipv4RibEdmPath{} }
func (m *Ipv4RibEdmPath) String() string            { return proto.CompactTextString(m) }
func (*Ipv4RibEdmPath) ProtoMessage()               {}
func (*Ipv4RibEdmPath) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Ipv4RibEdmPath) GetIpv4RibEdmPath() []*Ipv4RibEdmPathItem {
	if m != nil {
		return m.Ipv4RibEdmPath
	}
	return nil
}

type Ipv4RibEdmPathItem struct {
	// Nexthop
	Address string `protobuf:"bytes,1,opt,name=address" json:"address,omitempty"`
	// Infosource
	InformationSource string `protobuf:"bytes,2,opt,name=information_source,json=informationSource" json:"information_source,omitempty"`
	// V6 nexthop
	V6Nexthop string `protobuf:"bytes,3,opt,name=v6_nexthop,json=v6Nexthop" json:"v6_nexthop,omitempty"`
	// Interface Name
	InterfaceName string `protobuf:"bytes,4,opt,name=interface_name,json=interfaceName" json:"interface_name,omitempty"`
	// Metrics
	Metric uint32 `protobuf:"varint,5,opt,name=metric" json:"metric,omitempty"`
	// Load Metrics
	LoadMetric uint32 `protobuf:"varint,6,opt,name=load_metric,json=loadMetric" json:"load_metric,omitempty"`
	// Flags extended to 64 bits
	Flags64 uint64 `protobuf:"varint,7,opt,name=flags64" json:"flags64,omitempty"`
	// Flags
	Flags uint32 `protobuf:"varint,8,opt,name=flags" json:"flags,omitempty"`
	// Private Flags
	PrivateFlags uint32 `protobuf:"varint,9,opt,name=private_flags,json=privateFlags" json:"private_flags,omitempty"`
	// Looping path
	Looped bool `protobuf:"varint,10,opt,name=looped" json:"looped,omitempty"`
	// Nexthop tableid
	NextHopTableId uint32 `protobuf:"varint,11,opt,name=next_hop_table_id,json=nextHopTableId" json:"next_hop_table_id,omitempty"`
	// VRF Name of the nh table
	NextHopVrfName string `protobuf:"bytes,12,opt,name=next_hop_vrf_name,json=nextHopVrfName" json:"next_hop_vrf_name,omitempty"`
	// NH table name
	NextHopTableName string `protobuf:"bytes,13,opt,name=next_hop_table_name,json=nextHopTableName" json:"next_hop_table_name,omitempty"`
	// NH afi
	NextHopAfi uint32 `protobuf:"varint,14,opt,name=next_hop_afi,json=nextHopAfi" json:"next_hop_afi,omitempty"`
	// NH safi
	NextHopSafi uint32 `protobuf:"varint,15,opt,name=next_hop_safi,json=nextHopSafi" json:"next_hop_safi,omitempty"`
	// Label associated with this path
	RouteLabel uint32 `protobuf:"varint,16,opt,name=route_label,json=routeLabel" json:"route_label,omitempty"`
	// Tunnel ID associated with this path
	TunnelId uint32 `protobuf:"varint,17,opt,name=tunnel_id,json=tunnelId" json:"tunnel_id,omitempty"`
	// Path id of this path
	Pathid uint32 `protobuf:"varint,18,opt,name=pathid" json:"pathid,omitempty"`
	// Path id of this path's backup
	BackupPathid uint32 `protobuf:"varint,19,opt,name=backup_pathid,json=backupPathid" json:"backup_pathid,omitempty"`
	// Refcnt of backup
	RefCntOfBackup uint32 `protobuf:"varint,20,opt,name=ref_cnt_of_backup,json=refCntOfBackup" json:"ref_cnt_of_backup,omitempty"`
	// Number of extended communities
	NumberOfExtendedCommunities uint32 `protobuf:"varint,21,opt,name=number_of_extended_communities,json=numberOfExtendedCommunities" json:"number_of_extended_communities,omitempty"`
	// MVPN attribute present
	MvpnPresent bool `protobuf:"varint,22,opt,name=mvpn_present,json=mvpnPresent" json:"mvpn_present,omitempty"`
	// Path RT present
	PathrtPresent           bool `protobuf:"varint,23,opt,name=pathrt_present,json=pathrtPresent" json:"pathrt_present,omitempty"`
	VrfimportrtPresent      bool `protobuf:"varint,24,opt,name=vrfimportrt_present,json=vrfimportrtPresent" json:"vrfimportrt_present,omitempty"`
	SourceasrtPresent       bool `protobuf:"varint,25,opt,name=sourceasrt_present,json=sourceasrtPresent" json:"sourceasrt_present,omitempty"`
	SourcerdPresent         bool `protobuf:"varint,26,opt,name=sourcerd_present,json=sourcerdPresent" json:"sourcerd_present,omitempty"`
	SegmentedNexthopPresent bool `protobuf:"varint,27,opt,name=segmented_nexthop_present,json=segmentedNexthopPresent" json:"segmented_nexthop_present,omitempty"`
	// NHID associated with this path
	NextHopId uint32 `protobuf:"varint,28,opt,name=next_hop_id,json=nextHopId" json:"next_hop_id,omitempty"`
	// NHID references
	NextHopIdRefcount uint32 `protobuf:"varint,29,opt,name=next_hop_id_refcount,json=nextHopIdRefcount" json:"next_hop_id_refcount,omitempty"`
	// OSPF area associated with the path
	OspfAreaId string `protobuf:"bytes,30,opt,name=ospf_area_id,json=ospfAreaId" json:"ospf_area_id,omitempty"`
	// Remote backup node address
	RemoteBackupAddr [][]byte `protobuf:"bytes,31,rep,name=remote_backup_addr,json=remoteBackupAddr,proto3" json:"remote_backup_addr,omitempty"`
	// Path has a label stack
	HasLabelstk bool `protobuf:"varint,32,opt,name=has_labelstk,json=hasLabelstk" json:"has_labelstk,omitempty"`
	// Number of labels in stack
	NumLabels uint32 `protobuf:"varint,33,opt,name=num_labels,json=numLabels" json:"num_labels,omitempty"`
	// Labels for this path
	Labelstk []uint32 `protobuf:"varint,34,rep,packed,name=labelstk" json:"labelstk,omitempty"`
	// binding Label for this path
	BindingLabel uint32 `protobuf:"varint,35,opt,name=binding_label,json=bindingLabel" json:"binding_label,omitempty"`
	// Fib nhid encap id
	NhidFeid uint64 `protobuf:"varint,36,opt,name=nhid_feid,json=nhidFeid" json:"nhid_feid,omitempty"`
	// Fib mpls encap id
	MplsFeid uint64 `protobuf:"varint,37,opt,name=mpls_feid,json=mplsFeid" json:"mpls_feid,omitempty"`
}

func (m *Ipv4RibEdmPathItem) Reset()                    { *m = Ipv4RibEdmPathItem{} }
func (m *Ipv4RibEdmPathItem) String() string            { return proto.CompactTextString(m) }
func (*Ipv4RibEdmPathItem) ProtoMessage()               {}
func (*Ipv4RibEdmPathItem) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Ipv4RibEdmPathItem) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Ipv4RibEdmPathItem) GetInformationSource() string {
	if m != nil {
		return m.InformationSource
	}
	return ""
}

func (m *Ipv4RibEdmPathItem) GetV6Nexthop() string {
	if m != nil {
		return m.V6Nexthop
	}
	return ""
}

func (m *Ipv4RibEdmPathItem) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

func (m *Ipv4RibEdmPathItem) GetMetric() uint32 {
	if m != nil {
		return m.Metric
	}
	return 0
}

func (m *Ipv4RibEdmPathItem) GetLoadMetric() uint32 {
	if m != nil {
		return m.LoadMetric
	}
	return 0
}

func (m *Ipv4RibEdmPathItem) GetFlags64() uint64 {
	if m != nil {
		return m.Flags64
	}
	return 0
}

func (m *Ipv4RibEdmPathItem) GetFlags() uint32 {
	if m != nil {
		return m.Flags
	}
	return 0
}

func (m *Ipv4RibEdmPathItem) GetPrivateFlags() uint32 {
	if m != nil {
		return m.PrivateFlags
	}
	return 0
}

func (m *Ipv4RibEdmPathItem) GetLooped() bool {
	if m != nil {
		return m.Looped
	}
	return false
}

func (m *Ipv4RibEdmPathItem) GetNextHopTableId() uint32 {
	if m != nil {
		return m.NextHopTableId
	}
	return 0
}

func (m *Ipv4RibEdmPathItem) GetNextHopVrfName() string {
	if m != nil {
		return m.NextHopVrfName
	}
	return ""
}

func (m *Ipv4RibEdmPathItem) GetNextHopTableName() string {
	if m != nil {
		return m.NextHopTableName
	}
	return ""
}

func (m *Ipv4RibEdmPathItem) GetNextHopAfi() uint32 {
	if m != nil {
		return m.NextHopAfi
	}
	return 0
}

func (m *Ipv4RibEdmPathItem) GetNextHopSafi() uint32 {
	if m != nil {
		return m.NextHopSafi
	}
	return 0
}

func (m *Ipv4RibEdmPathItem) GetRouteLabel() uint32 {
	if m != nil {
		return m.RouteLabel
	}
	return 0
}

func (m *Ipv4RibEdmPathItem) GetTunnelId() uint32 {
	if m != nil {
		return m.TunnelId
	}
	return 0
}

func (m *Ipv4RibEdmPathItem) GetPathid() uint32 {
	if m != nil {
		return m.Pathid
	}
	return 0
}

func (m *Ipv4RibEdmPathItem) GetBackupPathid() uint32 {
	if m != nil {
		return m.BackupPathid
	}
	return 0
}

func (m *Ipv4RibEdmPathItem) GetRefCntOfBackup() uint32 {
	if m != nil {
		return m.RefCntOfBackup
	}
	return 0
}

func (m *Ipv4RibEdmPathItem) GetNumberOfExtendedCommunities() uint32 {
	if m != nil {
		return m.NumberOfExtendedCommunities
	}
	return 0
}

func (m *Ipv4RibEdmPathItem) GetMvpnPresent() bool {
	if m != nil {
		return m.MvpnPresent
	}
	return false
}

func (m *Ipv4RibEdmPathItem) GetPathrtPresent() bool {
	if m != nil {
		return m.PathrtPresent
	}
	return false
}

func (m *Ipv4RibEdmPathItem) GetVrfimportrtPresent() bool {
	if m != nil {
		return m.VrfimportrtPresent
	}
	return false
}

func (m *Ipv4RibEdmPathItem) GetSourceasrtPresent() bool {
	if m != nil {
		return m.SourceasrtPresent
	}
	return false
}

func (m *Ipv4RibEdmPathItem) GetSourcerdPresent() bool {
	if m != nil {
		return m.SourcerdPresent
	}
	return false
}

func (m *Ipv4RibEdmPathItem) GetSegmentedNexthopPresent() bool {
	if m != nil {
		return m.SegmentedNexthopPresent
	}
	return false
}

func (m *Ipv4RibEdmPathItem) GetNextHopId() uint32 {
	if m != nil {
		return m.NextHopId
	}
	return 0
}

func (m *Ipv4RibEdmPathItem) GetNextHopIdRefcount() uint32 {
	if m != nil {
		return m.NextHopIdRefcount
	}
	return 0
}

func (m *Ipv4RibEdmPathItem) GetOspfAreaId() string {
	if m != nil {
		return m.OspfAreaId
	}
	return ""
}

func (m *Ipv4RibEdmPathItem) GetRemoteBackupAddr() [][]byte {
	if m != nil {
		return m.RemoteBackupAddr
	}
	return nil
}

func (m *Ipv4RibEdmPathItem) GetHasLabelstk() bool {
	if m != nil {
		return m.HasLabelstk
	}
	return false
}

func (m *Ipv4RibEdmPathItem) GetNumLabels() uint32 {
	if m != nil {
		return m.NumLabels
	}
	return 0
}

func (m *Ipv4RibEdmPathItem) GetLabelstk() []uint32 {
	if m != nil {
		return m.Labelstk
	}
	return nil
}

func (m *Ipv4RibEdmPathItem) GetBindingLabel() uint32 {
	if m != nil {
		return m.BindingLabel
	}
	return 0
}

func (m *Ipv4RibEdmPathItem) GetNhidFeid() uint64 {
	if m != nil {
		return m.NhidFeid
	}
	return 0
}

func (m *Ipv4RibEdmPathItem) GetMplsFeid() uint64 {
	if m != nil {
		return m.MplsFeid
	}
	return 0
}

// Information of local label for route head
type RibEdmLocalLabel struct {
	// Protocol Name
	ProtocolName string `protobuf:"bytes,1,opt,name=protocol_name,json=protocolName" json:"protocol_name,omitempty"`
	// Client ID
	ClientId uint32 `protobuf:"varint,2,opt,name=client_id,json=clientId" json:"client_id,omitempty"`
	// Stale
	Stale uint32 `protobuf:"varint,3,opt,name=stale" json:"stale,omitempty"`
	// Mirrored
	Mirrored uint32 `protobuf:"varint,4,opt,name=mirrored" json:"mirrored,omitempty"`
	// Merge disable
	MergeDisable uint32 `protobuf:"varint,5,opt,name=merge_disable,json=mergeDisable" json:"merge_disable,omitempty"`
	// Redist only
	RedistOnly uint32 `protobuf:"varint,6,opt,name=redist_only,json=redistOnly" json:"redist_only,omitempty"`
	// Local label
	Label uint32 `protobuf:"varint,7,opt,name=label" json:"label,omitempty"`
	// Distance
	Distance uint32 `protobuf:"varint,8,opt,name=distance" json:"distance,omitempty"`
}

func (m *RibEdmLocalLabel) Reset()                    { *m = RibEdmLocalLabel{} }
func (m *RibEdmLocalLabel) String() string            { return proto.CompactTextString(m) }
func (*RibEdmLocalLabel) ProtoMessage()               {}
func (*RibEdmLocalLabel) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *RibEdmLocalLabel) GetProtocolName() string {
	if m != nil {
		return m.ProtocolName
	}
	return ""
}

func (m *RibEdmLocalLabel) GetClientId() uint32 {
	if m != nil {
		return m.ClientId
	}
	return 0
}

func (m *RibEdmLocalLabel) GetStale() uint32 {
	if m != nil {
		return m.Stale
	}
	return 0
}

func (m *RibEdmLocalLabel) GetMirrored() uint32 {
	if m != nil {
		return m.Mirrored
	}
	return 0
}

func (m *RibEdmLocalLabel) GetMergeDisable() uint32 {
	if m != nil {
		return m.MergeDisable
	}
	return 0
}

func (m *RibEdmLocalLabel) GetRedistOnly() uint32 {
	if m != nil {
		return m.RedistOnly
	}
	return 0
}

func (m *RibEdmLocalLabel) GetLabel() uint32 {
	if m != nil {
		return m.Label
	}
	return 0
}

func (m *RibEdmLocalLabel) GetDistance() uint32 {
	if m != nil {
		return m.Distance
	}
	return 0
}

func init() {
	proto.RegisterType((*Ipv4RibEdmRoute_KEYS)(nil), "cisco_ios_xr_ip_rib_ipv4_oper.rib.vrfs.vrf.afs.af.safs.saf.ip_rib_route_table_names.ip_rib_route_table_name.destination_kw.dest_best_routes.dest_best_route.ipv4_rib_edm_route_KEYS")
	proto.RegisterType((*Ipv4RibEdmRoute)(nil), "cisco_ios_xr_ip_rib_ipv4_oper.rib.vrfs.vrf.afs.af.safs.saf.ip_rib_route_table_names.ip_rib_route_table_name.destination_kw.dest_best_routes.dest_best_route.ipv4_rib_edm_route")
	proto.RegisterType((*Ipv4RibEdmPath)(nil), "cisco_ios_xr_ip_rib_ipv4_oper.rib.vrfs.vrf.afs.af.safs.saf.ip_rib_route_table_names.ip_rib_route_table_name.destination_kw.dest_best_routes.dest_best_route.ipv4_rib_edm_path")
	proto.RegisterType((*Ipv4RibEdmPathItem)(nil), "cisco_ios_xr_ip_rib_ipv4_oper.rib.vrfs.vrf.afs.af.safs.saf.ip_rib_route_table_names.ip_rib_route_table_name.destination_kw.dest_best_routes.dest_best_route.ipv4_rib_edm_path_item")
	proto.RegisterType((*RibEdmLocalLabel)(nil), "cisco_ios_xr_ip_rib_ipv4_oper.rib.vrfs.vrf.afs.af.safs.saf.ip_rib_route_table_names.ip_rib_route_table_name.destination_kw.dest_best_routes.dest_best_route.rib_edm_local_label")
}

func init() { proto.RegisterFile("ipv4_rib_edm_route.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 1494 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x57, 0xcd, 0x72, 0x1b, 0xb9,
	0x11, 0x2e, 0x5a, 0x96, 0x44, 0x42, 0xa2, 0x2c, 0x8e, 0x14, 0x09, 0xb2, 0xfc, 0x43, 0xcb, 0x71,
	0x15, 0x95, 0x8a, 0xe9, 0x94, 0xed, 0x28, 0x89, 0xf3, 0x2b, 0xcb, 0xb2, 0xcd, 0x58, 0xb6, 0x14,
	0xda, 0xe5, 0xaa, 0x9c, 0x50, 0xe0, 0x00, 0x43, 0xa2, 0x3c, 0x33, 0x18, 0x03, 0x20, 0x25, 0x1d,
	0xf3, 0x1c, 0x79, 0x85, 0xad, 0xbd, 0xec, 0x7b, 0xec, 0x2b, 0xec, 0x13, 0xec, 0x65, 0xab, 0xf6,
	0xbe, 0x85, 0x6e, 0xcc, 0x90, 0x92, 0xbc, 0xf7, 0xdd, 0x8b, 0xcc, 0xfe, 0xbe, 0x0f, 0x98, 0x46,
	0x77, 0xa3, 0x1b, 0x26, 0x54, 0x15, 0x93, 0xa7, 0xcc, 0xa8, 0x01, 0x93, 0x22, 0x63, 0x46, 0x8f,
	0x9d, 0xec, 0x16, 0x46, 0x3b, 0x1d, 0xfd, 0xbf, 0x16, 0x2b, 0x1b, 0x6b, 0xa6, 0xb4, 0x65, 0x67,
	0x86, 0xa9, 0x02, 0x54, 0x20, 0xd7, 0x85, 0x34, 0x5d, 0xa3, 0x06, 0xdd, 0x89, 0x49, 0xac, 0xff,
	0xd3, 0xe5, 0x89, 0xed, 0xf2, 0xa4, 0x6b, 0xfd, 0xbf, 0x96, 0x27, 0xdd, 0xa0, 0x86, 0xfd, 0x98,
	0xe3, 0x83, 0x54, 0xb2, 0x9c, 0x67, 0xd2, 0xfe, 0x1c, 0xd1, 0x15, 0xd2, 0x3a, 0x95, 0x73, 0xa7,
	0x74, 0xce, 0x3e, 0x9d, 0x82, 0xc9, 0x06, 0xfe, 0x0f, 0x28, 0xed, 0x65, 0x60, 0xe7, 0xab, 0x1a,
	0xd9, 0xbc, 0xea, 0x3a, 0x7b, 0x73, 0xf8, 0xdf, 0xf7, 0xd1, 0x16, 0xa9, 0x4f, 0x4c, 0x02, 0xfb,
	0xd2, 0x5a, 0xbb, 0xd6, 0x69, 0xf4, 0x17, 0x27, 0x26, 0x79, 0xc7, 0x33, 0x19, 0x6d, 0x92, 0x45,
	0x1e, 0x98, 0x6b, 0xc0, 0x2c, 0x70, 0x24, 0xb6, 0x48, 0xdd, 0x96, 0xcc, 0x1c, 0xae, 0xb1, 0x81,
	0xea, 0x90, 0xd5, 0xcb, 0xee, 0xd2, 0xeb, 0x20, 0x59, 0x01, 0xfc, 0x83, 0x87, 0x41, 0x49, 0xc9,
	0x22, 0x17, 0xc2, 0x48, 0x6b, 0xe9, 0x3c, 0xee, 0x11, 0xcc, 0x9d, 0xef, 0x1b, 0x24, 0xba, 0xea,
	0x6e, 0xb4, 0x41, 0x16, 0x0a, 0x23, 0x13, 0x75, 0x46, 0x1f, 0xa3, 0x37, 0x68, 0x45, 0xf7, 0x49,
	0x13, 0x7f, 0xb1, 0x54, 0xe6, 0x43, 0x37, 0xa2, 0x4f, 0xda, 0xb5, 0x4e, 0xb3, 0xbf, 0x8c, 0xe0,
	0x11, 0x60, 0x5e, 0x84, 0x7e, 0x4d, 0xa4, 0xb1, 0x4a, 0xe7, 0xf4, 0x29, 0x8a, 0x00, 0xfc, 0x88,
	0x58, 0x74, 0x97, 0x2c, 0x41, 0x3a, 0x63, 0x9d, 0x32, 0x25, 0xe8, 0x1f, 0x41, 0x42, 0x4a, 0xa8,
	0x27, 0xf0, 0x53, 0x41, 0x00, 0x47, 0xdb, 0x03, 0x4f, 0x96, 0x4b, 0x10, 0x0e, 0x76, 0x93, 0xd4,
	0x55, 0x6e, 0x1d, 0xcf, 0x63, 0x49, 0xff, 0x04, 0x7c, 0x65, 0x47, 0xdb, 0xa4, 0x11, 0xa7, 0x4a,
	0xe6, 0xce, 0xef, 0xff, 0x67, 0xd8, 0xbf, 0x8e, 0x40, 0x4f, 0x44, 0xb7, 0x09, 0x09, 0xb1, 0x3b,
	0x2f, 0x24, 0xfd, 0x0b, 0xb0, 0x0d, 0x8c, 0xda, 0x79, 0x01, 0xfb, 0x16, 0x46, 0x69, 0xa3, 0xdc,
	0x39, 0x7d, 0x86, 0x4b, 0x4b, 0x1b, 0x32, 0x32, 0x11, 0xb8, 0xf0, 0xaf, 0xc0, 0x2d, 0xda, 0x89,
	0x80, 0x65, 0xeb, 0x64, 0x3e, 0x49, 0xf9, 0xd0, 0xd2, 0xbf, 0x01, 0x8e, 0x46, 0xf4, 0x80, 0xac,
	0xc8, 0x33, 0x27, 0x73, 0x21, 0x05, 0x43, 0xfa, 0xef, 0xed, 0x5a, 0xe7, 0x7a, 0xbf, 0x59, 0xa2,
	0x2f, 0x41, 0xb6, 0x4a, 0xe6, 0x1c, 0x1f, 0xd2, 0x7f, 0xc0, 0x52, 0xff, 0xd3, 0x7b, 0x21, 0x54,
	0x38, 0xdd, 0x3f, 0xd1, 0x8b, 0xd2, 0x8e, 0x1e, 0x92, 0x48, 0xa8, 0x10, 0x60, 0x56, 0xa9, 0xfe,
	0x05, 0xaa, 0x56, 0xc5, 0xbc, 0x28, 0xe5, 0x1b, 0x64, 0x21, 0x93, 0xce, 0xa8, 0x98, 0xee, 0x83,
	0x24, 0x58, 0x90, 0x06, 0xee, 0x46, 0x96, 0xc5, 0x7a, 0x9c, 0x3b, 0xfa, 0x3c, 0xa4, 0xc1, 0x43,
	0x07, 0x1e, 0xf1, 0xdf, 0xe1, 0xce, 0x19, 0x35, 0xf0, 0xc1, 0x52, 0x42, 0xe6, 0xce, 0xc7, 0xe4,
	0x00, 0xbf, 0x53, 0x31, 0xbd, 0x40, 0xf8, 0xac, 0x39, 0xc3, 0x93, 0x44, 0xc5, 0x4c, 0xe5, 0x42,
	0x9e, 0xd1, 0x17, 0x98, 0xfb, 0x00, 0xf6, 0x3c, 0x16, 0xed, 0x96, 0x85, 0x5b, 0x18, 0x19, 0x4b,
	0x21, 0xbd, 0xe7, 0x87, 0xa0, 0xbb, 0x01, 0xf8, 0x49, 0x05, 0xfb, 0x24, 0x7e, 0xd6, 0x96, 0x0d,
	0x8d, 0x1e, 0x17, 0xf4, 0x25, 0xc6, 0xe0, 0xb3, 0xb6, 0xaf, 0xbc, 0xed, 0x33, 0x91, 0xa4, 0xfa,
	0x94, 0xf9, 0xb0, 0xbd, 0xc2, 0x4c, 0x78, 0xfb, 0x03, 0x1f, 0xfa, 0x75, 0xc9, 0xa9, 0x60, 0x71,
	0xca, 0xad, 0xa5, 0xaf, 0x71, 0x5d, 0x72, 0x2a, 0x0e, 0xbc, 0xed, 0xc9, 0x42, 0xc5, 0xe1, 0xc8,
	0xbd, 0x90, 0x5e, 0x15, 0xe3, 0x81, 0x37, 0xc8, 0x02, 0x8f, 0x9d, 0x9a, 0x48, 0xfa, 0xef, 0x76,
	0xad, 0x53, 0xef, 0x07, 0x2b, 0xba, 0x45, 0x1a, 0x55, 0x58, 0xe9, 0x1b, 0xa0, 0xa6, 0x40, 0xf4,
	0x07, 0xb2, 0x3e, 0x4d, 0x07, 0x94, 0x28, 0x16, 0xed, 0x11, 0x14, 0xe5, 0x34, 0x55, 0x27, 0x9e,
	0x82, 0xd2, 0xdd, 0x26, 0x58, 0x6f, 0x8c, 0x0f, 0x25, 0x7d, 0x8b, 0x4e, 0x00, 0xb0, 0x3f, 0x94,
	0x3e, 0x2d, 0x48, 0xa6, 0x7c, 0x20, 0x53, 0xfa, 0x0e, 0xd3, 0x02, 0xd0, 0x91, 0x47, 0xfc, 0x8d,
	0x2e, 0x7d, 0x39, 0xc6, 0x93, 0x4f, 0xa6, 0x17, 0xcb, 0x0d, 0xd2, 0xea, 0xee, 0x9d, 0x40, 0xa9,
	0x11, 0x37, 0x48, 0xcb, 0x9b, 0xf7, 0x3b, 0xd2, 0xc2, 0xbd, 0x33, 0x2d, 0x54, 0x72, 0xce, 0x9c,
	0xca, 0x24, 0xfd, 0x0f, 0xc8, 0x30, 0xfc, 0x6f, 0x01, 0xff, 0xa0, 0x32, 0x19, 0x7d, 0x5b, 0x2b,
	0xef, 0x89, 0x2f, 0x09, 0xda, 0x6f, 0xd7, 0x3a, 0x4b, 0x8f, 0xbf, 0xae, 0x75, 0x7f, 0xc1, 0x1d,
	0xb8, 0x7b, 0xa1, 0x9d, 0x79, 0xb7, 0xc3, 0xc5, 0x3e, 0xe1, 0x6e, 0xb4, 0xf3, 0x63, 0x8d, 0xb4,
	0xae, 0x08, 0xa2, 0xef, 0xbe, 0x84, 0xd2, 0x5a, 0x7b, 0xae, 0xb3, 0xf4, 0xf8, 0x9b, 0x5f, 0xd9,
	0x69, 0x99, 0x72, 0x32, 0xeb, 0xaf, 0x78, 0xbc, 0xaf, 0x06, 0x87, 0x22, 0x83, 0x73, 0xff, 0x40,
	0xc8, 0xc6, 0x97, 0xa5, 0xb3, 0xc3, 0xa1, 0x76, 0x61, 0x38, 0xf8, 0xbb, 0xaf, 0xf2, 0x44, 0x9b,
	0x0c, 0xdd, 0xb1, 0x7a, 0x6c, 0xe2, 0x72, 0x3e, 0xb5, 0x66, 0x98, 0xf7, 0x40, 0xf8, 0x9e, 0x3a,
	0xd9, 0x63, 0xb9, 0x3c, 0x73, 0x23, 0x5d, 0x84, 0x61, 0xd5, 0x98, 0xec, 0xbd, 0x43, 0xc0, 0xb7,
	0x41, 0x95, 0x3b, 0x69, 0x12, 0x1e, 0x5f, 0x18, 0x56, 0xcd, 0x0a, 0x85, 0x7b, 0x31, 0xed, 0x54,
	0xf3, 0x97, 0x3b, 0x55, 0xaa, 0xb9, 0x60, 0x81, 0x5c, 0xc0, 0x2b, 0xe1, 0xa1, 0xb7, 0x28, 0xa0,
	0x64, 0x11, 0xba, 0xeb, 0xde, 0x53, 0xba, 0x08, 0xd5, 0x5c, 0x9a, 0xd3, 0xb6, 0x5c, 0x9f, 0x6d,
	0xcb, 0x30, 0x60, 0xd4, 0x84, 0x3b, 0x19, 0xba, 0x72, 0xa3, 0x9c, 0x65, 0x00, 0x62, 0x53, 0xde,
	0x20, 0x0b, 0xa9, 0xd6, 0x85, 0x14, 0x94, 0x60, 0x37, 0x40, 0x2b, 0xda, 0x25, 0x2d, 0x7f, 0x50,
	0x36, 0xd2, 0x45, 0xc8, 0x9e, 0x12, 0x74, 0x09, 0x36, 0x58, 0xf1, 0xc4, 0x6b, 0x5d, 0xc0, 0xf8,
	0xed, 0x5d, 0x94, 0x56, 0xe3, 0x7f, 0x19, 0xe7, 0x74, 0x90, 0x7e, 0x0c, 0xaf, 0x80, 0x87, 0x64,
	0xed, 0xd2, 0xae, 0x20, 0x6e, 0x82, 0x78, 0x75, 0x76, 0x5f, 0x90, 0xb7, 0xc9, 0x72, 0x25, 0xe7,
	0x89, 0xa2, 0x2b, 0x18, 0x93, 0xa0, 0xdb, 0x4f, 0x54, 0xb4, 0x43, 0x9a, 0x95, 0xc2, 0x7a, 0xc9,
	0x0d, 0x90, 0x2c, 0x05, 0xc9, 0x7b, 0x9e, 0xa8, 0xcb, 0xbd, 0x66, 0xf5, 0x4a, 0xaf, 0xd9, 0x26,
	0x0d, 0x37, 0xce, 0x73, 0x09, 0x83, 0xba, 0x85, 0x9d, 0x0a, 0x81, 0x9e, 0x80, 0x97, 0x02, 0x77,
	0x23, 0x25, 0x68, 0x84, 0xe9, 0x42, 0xcb, 0x47, 0x77, 0xc0, 0xe3, 0x4f, 0xe3, 0x82, 0x05, 0x7a,
	0x0d, 0xa3, 0x8b, 0xe0, 0x09, 0x8a, 0x76, 0x49, 0xcb, 0xc8, 0x84, 0xc5, 0xb9, 0x63, 0x3a, 0x61,
	0x48, 0xd1, 0x75, 0x8c, 0xa2, 0x91, 0xc9, 0x41, 0xee, 0x8e, 0x93, 0xe7, 0x80, 0x46, 0x07, 0xe4,
	0x4e, 0x3e, 0xce, 0x06, 0xd2, 0x78, 0x65, 0x35, 0x4e, 0x63, 0x9d, 0x65, 0xe3, 0x5c, 0x39, 0x25,
	0x2d, 0xfd, 0x0d, 0xac, 0xdb, 0x46, 0xd5, 0x71, 0x72, 0x18, 0x34, 0x07, 0x53, 0x49, 0x74, 0x8f,
	0x2c, 0x67, 0x93, 0xc2, 0x37, 0x68, 0x69, 0x65, 0xee, 0xe8, 0x06, 0xe4, 0x74, 0xc9, 0x63, 0x27,
	0x08, 0xf9, 0x2a, 0xf5, 0x0e, 0x1b, 0x57, 0x89, 0x36, 0x41, 0xd4, 0x44, 0xb4, 0x94, 0x3d, 0x22,
	0x6b, 0x13, 0x93, 0xa8, 0xac, 0xd0, 0xc6, 0xcd, 0x68, 0x29, 0x68, 0xa3, 0x19, 0xaa, 0x5c, 0xf0,
	0x90, 0x44, 0x78, 0x7f, 0xb8, 0x9d, 0xd1, 0x6f, 0x81, 0xbe, 0x35, 0x65, 0x4a, 0xf9, 0x2e, 0x59,
	0x45, 0xd0, 0x88, 0x4a, 0x7c, 0x13, 0xc4, 0x37, 0x4a, 0xbc, 0x94, 0x3e, 0x23, 0x5b, 0x56, 0x0e,
	0x33, 0x99, 0x3b, 0x29, 0xca, 0xdb, 0x57, 0xad, 0xd9, 0x86, 0x35, 0x9b, 0x95, 0x20, 0x5c, 0xc6,
	0x72, 0xed, 0x1d, 0xb2, 0x54, 0xd5, 0x87, 0x12, 0xf4, 0x16, 0xbe, 0x83, 0x42, 0x75, 0xf4, 0x44,
	0xf4, 0x88, 0xac, 0xcf, 0xf0, 0xcc, 0xc8, 0x04, 0x87, 0xe6, 0x6d, 0x9c, 0xff, 0x95, 0xb0, 0x1f,
	0x08, 0x5f, 0x92, 0xda, 0x16, 0x09, 0xe3, 0x46, 0x72, 0xbf, 0xe3, 0x1d, 0x28, 0x5d, 0xe2, 0xb1,
	0x7d, 0x23, 0x79, 0x4f, 0x44, 0xbf, 0x27, 0x91, 0x91, 0x99, 0x76, 0x32, 0xe4, 0x9b, 0xf9, 0x6e,
	0x43, 0xef, 0xb6, 0xe7, 0x3a, 0xcb, 0xfd, 0x55, 0x64, 0x30, 0xe5, 0xfb, 0x42, 0x18, 0x9f, 0xb1,
	0x11, 0xb7, 0x58, 0x9a, 0xd6, 0x7d, 0xa2, 0x6d, 0xcc, 0xd8, 0x88, 0xdb, 0xa3, 0x00, 0xf9, 0xb6,
	0x93, 0x8f, 0xb3, 0x20, 0xa1, 0xf7, 0xc2, 0x11, 0xc6, 0x19, 0x0a, 0xfc, 0x23, 0xaa, 0x5a, 0xbd,
	0xd3, 0x9e, 0xf3, 0xc5, 0x5b, 0xda, 0x50, 0xa4, 0x2a, 0x17, 0x2a, 0x1f, 0x86, 0xe2, 0xbf, 0x1f,
	0x8a, 0x14, 0xc1, 0xaa, 0xfc, 0xf3, 0x91, 0x12, 0x2c, 0x91, 0x4a, 0xd0, 0xdf, 0x42, 0x67, 0xa9,
	0x7b, 0xe0, 0xa5, 0x54, 0xc2, 0x93, 0x59, 0x91, 0x5a, 0x24, 0x1f, 0x20, 0xe9, 0x01, 0x4f, 0xee,
	0xfc, 0xef, 0x1a, 0x59, 0x2b, 0xfb, 0x6d, 0xaa, 0x63, 0x9e, 0xe2, 0x57, 0xae, 0x3e, 0x6d, 0x6b,
	0x5f, 0x78, 0xda, 0x5e, 0x78, 0xbe, 0x5e, 0xbb, 0xf4, 0x7c, 0x5d, 0x27, 0xf3, 0xd6, 0xf1, 0x14,
	0xff, 0x4b, 0xd0, 0xec, 0xa3, 0xe1, 0x8f, 0x9a, 0x29, 0x63, 0xb4, 0x91, 0x02, 0x7a, 0x6b, 0xb3,
	0x5f, 0xd9, 0xfe, 0x9b, 0x99, 0x34, 0x43, 0xe9, 0xdf, 0x8a, 0xbe, 0x81, 0x84, 0xee, 0xba, 0x0c,
	0xe0, 0x0b, 0xc4, 0xa0, 0x15, 0x48, 0xff, 0x98, 0x64, 0x3a, 0x4f, 0xcf, 0xcb, 0x1e, 0x8b, 0xd0,
	0x71, 0x9e, 0x9e, 0xfb, 0xef, 0x62, 0xa0, 0x16, 0xf1, 0xbb, 0x78, 0x9e, 0xd9, 0x77, 0x6a, 0xfd,
	0xe2, 0x3b, 0x75, 0xb0, 0x00, 0x87, 0x7a, 0xf2, 0x53, 0x00, 0x00, 0x00, 0xff, 0xff, 0x0d, 0x39,
	0xe8, 0x99, 0xd0, 0x0d, 0x00, 0x00,
}