// Code generated by protoc-gen-go.
// source: statsdbag_datarate.proto
// DO NOT EDIT!

/*
Package cisco_ios_xr_infra_statsd_oper_infra_statistics_interfaces_interface_data_rate is a generated protocol buffer package.

It is generated from these files:
	statsdbag_datarate.proto

It has these top-level messages:
	StatsdbagDatarate_KEYS
	StatsdbagDatarate
*/
package cisco_ios_xr_infra_statsd_oper_infra_statistics_interfaces_interface_data_rate

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

// Datarate information
type StatsdbagDatarate_KEYS struct {
	InterfaceName string `protobuf:"bytes,1,opt,name=interface_name,json=interfaceName" json:"interface_name,omitempty"`
}

func (m *StatsdbagDatarate_KEYS) Reset()                    { *m = StatsdbagDatarate_KEYS{} }
func (m *StatsdbagDatarate_KEYS) String() string            { return proto.CompactTextString(m) }
func (*StatsdbagDatarate_KEYS) ProtoMessage()               {}
func (*StatsdbagDatarate_KEYS) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *StatsdbagDatarate_KEYS) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

type StatsdbagDatarate struct {
	// Input data rate in 1000's of bps
	InputDataRate uint64 `protobuf:"varint,50,opt,name=input_data_rate,json=inputDataRate" json:"input_data_rate,omitempty"`
	// Input packets per second
	InputPacketRate uint64 `protobuf:"varint,51,opt,name=input_packet_rate,json=inputPacketRate" json:"input_packet_rate,omitempty"`
	// Output data rate in 1000's of bps
	OutputDataRate uint64 `protobuf:"varint,52,opt,name=output_data_rate,json=outputDataRate" json:"output_data_rate,omitempty"`
	// Output packets per second
	OutputPacketRate uint64 `protobuf:"varint,53,opt,name=output_packet_rate,json=outputPacketRate" json:"output_packet_rate,omitempty"`
	// Peak input data rate
	PeakInputDataRate uint64 `protobuf:"varint,54,opt,name=peak_input_data_rate,json=peakInputDataRate" json:"peak_input_data_rate,omitempty"`
	// Peak input packet rate
	PeakInputPacketRate uint64 `protobuf:"varint,55,opt,name=peak_input_packet_rate,json=peakInputPacketRate" json:"peak_input_packet_rate,omitempty"`
	// Peak output data rate
	PeakOutputDataRate uint64 `protobuf:"varint,56,opt,name=peak_output_data_rate,json=peakOutputDataRate" json:"peak_output_data_rate,omitempty"`
	// Peak output packet rate
	PeakOutputPacketRate uint64 `protobuf:"varint,57,opt,name=peak_output_packet_rate,json=peakOutputPacketRate" json:"peak_output_packet_rate,omitempty"`
	// Bandwidth (in kbps)
	Bandwidth uint32 `protobuf:"varint,58,opt,name=bandwidth" json:"bandwidth,omitempty"`
	// Number of 30-sec intervals less one
	LoadInterval uint32 `protobuf:"varint,59,opt,name=load_interval,json=loadInterval" json:"load_interval,omitempty"`
	// Output load as fraction of 255
	OutputLoad uint32 `protobuf:"varint,60,opt,name=output_load,json=outputLoad" json:"output_load,omitempty"`
	// Input load as fraction of 255
	InputLoad uint32 `protobuf:"varint,61,opt,name=input_load,json=inputLoad" json:"input_load,omitempty"`
	// Reliability coefficient
	Reliability uint32 `protobuf:"varint,62,opt,name=reliability" json:"reliability,omitempty"`
}

func (m *StatsdbagDatarate) Reset()                    { *m = StatsdbagDatarate{} }
func (m *StatsdbagDatarate) String() string            { return proto.CompactTextString(m) }
func (*StatsdbagDatarate) ProtoMessage()               {}
func (*StatsdbagDatarate) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *StatsdbagDatarate) GetInputDataRate() uint64 {
	if m != nil {
		return m.InputDataRate
	}
	return 0
}

func (m *StatsdbagDatarate) GetInputPacketRate() uint64 {
	if m != nil {
		return m.InputPacketRate
	}
	return 0
}

func (m *StatsdbagDatarate) GetOutputDataRate() uint64 {
	if m != nil {
		return m.OutputDataRate
	}
	return 0
}

func (m *StatsdbagDatarate) GetOutputPacketRate() uint64 {
	if m != nil {
		return m.OutputPacketRate
	}
	return 0
}

func (m *StatsdbagDatarate) GetPeakInputDataRate() uint64 {
	if m != nil {
		return m.PeakInputDataRate
	}
	return 0
}

func (m *StatsdbagDatarate) GetPeakInputPacketRate() uint64 {
	if m != nil {
		return m.PeakInputPacketRate
	}
	return 0
}

func (m *StatsdbagDatarate) GetPeakOutputDataRate() uint64 {
	if m != nil {
		return m.PeakOutputDataRate
	}
	return 0
}

func (m *StatsdbagDatarate) GetPeakOutputPacketRate() uint64 {
	if m != nil {
		return m.PeakOutputPacketRate
	}
	return 0
}

func (m *StatsdbagDatarate) GetBandwidth() uint32 {
	if m != nil {
		return m.Bandwidth
	}
	return 0
}

func (m *StatsdbagDatarate) GetLoadInterval() uint32 {
	if m != nil {
		return m.LoadInterval
	}
	return 0
}

func (m *StatsdbagDatarate) GetOutputLoad() uint32 {
	if m != nil {
		return m.OutputLoad
	}
	return 0
}

func (m *StatsdbagDatarate) GetInputLoad() uint32 {
	if m != nil {
		return m.InputLoad
	}
	return 0
}

func (m *StatsdbagDatarate) GetReliability() uint32 {
	if m != nil {
		return m.Reliability
	}
	return 0
}

func init() {
	proto.RegisterType((*StatsdbagDatarate_KEYS)(nil), "cisco_ios_xr_infra_statsd_oper.infra_statistics.interfaces.interface.data_rate.statsdbag_datarate_KEYS")
	proto.RegisterType((*StatsdbagDatarate)(nil), "cisco_ios_xr_infra_statsd_oper.infra_statistics.interfaces.interface.data_rate.statsdbag_datarate")
}

func init() { proto.RegisterFile("statsdbag_datarate.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 379 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x92, 0x4b, 0x6b, 0xdb, 0x40,
	0x14, 0x85, 0x11, 0x98, 0x82, 0xaf, 0x6b, 0xb7, 0x9e, 0xba, 0xb5, 0x16, 0x2d, 0x15, 0x2e, 0x2d,
	0xa2, 0x14, 0x95, 0xd6, 0x75, 0x5f, 0x79, 0x90, 0x45, 0xb2, 0x30, 0x09, 0x4e, 0x70, 0x56, 0x59,
	0x0d, 0x57, 0xd2, 0x38, 0x19, 0x2c, 0x6b, 0x84, 0x34, 0xce, 0xe3, 0x87, 0xe5, 0xff, 0x85, 0xb9,
	0x23, 0xe4, 0x71, 0xbc, 0x33, 0xe7, 0x7c, 0xe7, 0x9c, 0xb9, 0x58, 0xe0, 0x57, 0x1a, 0x75, 0x95,
	0xc6, 0x78, 0xcd, 0x53, 0xd4, 0x58, 0xa2, 0x16, 0x51, 0x51, 0x2a, 0xad, 0xd8, 0x2c, 0x91, 0x55,
	0xa2, 0xb8, 0x54, 0x15, 0xbf, 0x2f, 0xb9, 0xcc, 0x17, 0x25, 0x72, 0x0b, 0x73, 0x55, 0x88, 0x32,
	0xda, 0x28, 0xb2, 0xd2, 0x32, 0xa9, 0x22, 0x99, 0x6b, 0x51, 0x2e, 0x30, 0x11, 0xce, 0xcf, 0xc8,
	0x94, 0x72, 0xd3, 0x3a, 0x3a, 0x82, 0xe1, 0xee, 0x16, 0x3f, 0x3d, 0xb9, 0xba, 0x64, 0x9f, 0xa1,
	0xd7, 0x24, 0x78, 0x8e, 0x2b, 0xe1, 0x7b, 0x81, 0x17, 0xb6, 0xe7, 0xdd, 0x46, 0x9d, 0xe1, 0x4a,
	0x8c, 0x1e, 0x5b, 0xc0, 0x76, 0x2b, 0xd8, 0x17, 0x78, 0x25, 0xf3, 0x62, 0xad, 0x79, 0xb3, 0xe5,
	0xff, 0x0c, 0xbc, 0xb0, 0x65, 0xe2, 0xc5, 0x5a, 0x1f, 0xa3, 0xc6, 0xb9, 0xe1, 0xbe, 0x42, 0xdf,
	0x72, 0x05, 0x26, 0x4b, 0xa1, 0x2d, 0x39, 0x26, 0xd2, 0x16, 0x5c, 0x90, 0x4e, 0x6c, 0x08, 0xaf,
	0xd5, 0x5a, 0x6f, 0x97, 0xfe, 0x22, 0xb4, 0x67, 0xf5, 0xa6, 0xf5, 0x1b, 0xb0, 0x9a, 0x74, 0x6b,
	0x27, 0xc4, 0xd6, 0x1d, 0x4e, 0xef, 0x77, 0x18, 0x14, 0x02, 0x97, 0xfc, 0xf9, 0x83, 0x7f, 0x13,
	0xdf, 0x37, 0xde, 0x74, 0xeb, 0xd1, 0x63, 0x78, 0xe7, 0x04, 0xdc, 0x89, 0x3f, 0x14, 0x79, 0xd3,
	0x44, 0x9c, 0x95, 0x1f, 0xf0, 0x96, 0x42, 0x3b, 0x27, 0xfc, 0xa5, 0x0c, 0x33, 0xe6, 0xf9, 0xf6,
	0x19, 0x13, 0x18, 0xba, 0x11, 0x77, 0xe8, 0x1f, 0x85, 0x06, 0x9b, 0x90, 0xb3, 0xf4, 0x1e, 0xda,
	0x31, 0xe6, 0xe9, 0x9d, 0x4c, 0xf5, 0x8d, 0xff, 0x3f, 0xf0, 0xc2, 0xee, 0x7c, 0x23, 0xb0, 0x4f,
	0xd0, 0xcd, 0x14, 0xa6, 0x9c, 0xfe, 0xc6, 0x5b, 0xcc, 0xfc, 0x3d, 0x22, 0x5e, 0x1a, 0x71, 0x5a,
	0x6b, 0xec, 0x23, 0x74, 0xea, 0x51, 0x23, 0xfb, 0xfb, 0x84, 0x80, 0x95, 0xce, 0x14, 0xa6, 0xec,
	0x03, 0x80, 0xbd, 0x9e, 0xfc, 0x03, 0x3b, 0x42, 0x0a, 0xd9, 0x01, 0x74, 0x4a, 0x91, 0x49, 0x8c,
	0x65, 0x26, 0xf5, 0x83, 0x7f, 0x48, 0xbe, 0x2b, 0xc5, 0x2f, 0xe8, 0x83, 0x1e, 0x3f, 0x05, 0x00,
	0x00, 0xff, 0xff, 0xeb, 0x3c, 0x8a, 0x81, 0xec, 0x02, 0x00, 0x00,
}
