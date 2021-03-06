// Code generated by protoc-gen-gogo.
// source: cockroach/config/config.proto
// DO NOT EDIT!

/*
	Package config is a generated protocol buffer package.

	It is generated from these files:
		cockroach/config/config.proto

	It has these top-level messages:
		GCPolicy
		ZoneConfig
		PrefixConfigMap
		PrefixConfig
		ConfigUnion
		SystemConfig
*/
package config

import proto "github.com/gogo/protobuf/proto"
import math "math"
import cockroach_proto "github.com/cockroachdb/cockroach/proto"
import cockroach_proto1 "github.com/cockroachdb/cockroach/proto"

// discarding unused import gogoproto "github.com/cockroachdb/gogoproto"

import github_com_cockroachdb_cockroach_proto "github.com/cockroachdb/cockroach/proto"

import io "io"
import fmt "fmt"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

// GCPolicy defines garbage collection policies which apply to MVCC
// values within a zone.
//
// TODO(spencer): flesh this out to include maximum number of values
//   as well as whether there's an intersection between max values
//   and TTL or a union.
type GCPolicy struct {
	// TTLSeconds specifies the maximum age of a value before it's
	// garbage collected. Only older versions of values are garbage
	// collected. Specifying <=0 mean older versions are never GC'd.
	TTLSeconds int32 `protobuf:"varint,1,opt,name=ttl_seconds" json:"ttl_seconds"`
}

func (m *GCPolicy) Reset()         { *m = GCPolicy{} }
func (m *GCPolicy) String() string { return proto.CompactTextString(m) }
func (*GCPolicy) ProtoMessage()    {}

func (m *GCPolicy) GetTTLSeconds() int32 {
	if m != nil {
		return m.TTLSeconds
	}
	return 0
}

// ZoneConfig holds configuration that is needed for a range of KV pairs.
type ZoneConfig struct {
	// ReplicaAttrs is a slice of Attributes, each describing required attributes
	// for each replica in the zone. The order in which the attributes are stored
	// in ReplicaAttrs is arbitrary and may change.
	ReplicaAttrs  []cockroach_proto.Attributes `protobuf:"bytes,1,rep,name=replica_attrs" json:"replica_attrs" yaml:"replicas,omitempty"`
	RangeMinBytes int64                        `protobuf:"varint,2,opt,name=range_min_bytes" json:"range_min_bytes" yaml:"range_min_bytes,omitempty"`
	RangeMaxBytes int64                        `protobuf:"varint,3,opt,name=range_max_bytes" json:"range_max_bytes" yaml:"range_max_bytes,omitempty"`
	// If GC policy is not set, uses the next highest, non-null policy
	// in the zone config hierarchy, up to the default policy if necessary.
	GC *GCPolicy `protobuf:"bytes,4,opt,name=gc" json:"gc,omitempty" yaml:"gc,omitempty"`
}

func (m *ZoneConfig) Reset()         { *m = ZoneConfig{} }
func (m *ZoneConfig) String() string { return proto.CompactTextString(m) }
func (*ZoneConfig) ProtoMessage()    {}

func (m *ZoneConfig) GetReplicaAttrs() []cockroach_proto.Attributes {
	if m != nil {
		return m.ReplicaAttrs
	}
	return nil
}

func (m *ZoneConfig) GetRangeMinBytes() int64 {
	if m != nil {
		return m.RangeMinBytes
	}
	return 0
}

func (m *ZoneConfig) GetRangeMaxBytes() int64 {
	if m != nil {
		return m.RangeMaxBytes
	}
	return 0
}

func (m *ZoneConfig) GetGC() *GCPolicy {
	if m != nil {
		return m.GC
	}
	return nil
}

// PrefixConfigMap contains a slice of prefix configs, sorted by
// prefix. Along with various accessor methods, the config map
// also contains additional prefix configs in the slice to
// account for the ends of prefix ranges.
type PrefixConfigMap struct {
	Configs []PrefixConfig `protobuf:"bytes,1,rep,name=configs" json:"configs"`
}

func (m *PrefixConfigMap) Reset()         { *m = PrefixConfigMap{} }
func (m *PrefixConfigMap) String() string { return proto.CompactTextString(m) }
func (*PrefixConfigMap) ProtoMessage()    {}

func (m *PrefixConfigMap) GetConfigs() []PrefixConfig {
	if m != nil {
		return m.Configs
	}
	return nil
}

// PrefixConfig relates a prefix key to a config object. PrefixConfig
// objects are the constituents of PrefixConfigMap objects. In order to
// support binary searches of hierarchical PrefixConfig objects,
// end-of-prefix PrefixConfig objects are added to a PrefixConfigMap to
// demarcate the end of a prefix range. Such end-of-range sentinels
// need to refer back to their parent prefix.
// The Canonical key refers to this parent PrefixConfig by its prefix.
type PrefixConfig struct {
	Prefix    github_com_cockroachdb_cockroach_proto.Key `protobuf:"bytes,1,opt,name=prefix,casttype=github.com/cockroachdb/cockroach/proto.Key" json:"prefix,omitempty"`
	Canonical github_com_cockroachdb_cockroach_proto.Key `protobuf:"bytes,2,opt,name=canonical,casttype=github.com/cockroachdb/cockroach/proto.Key" json:"canonical,omitempty"`
	Config    ConfigUnion                                `protobuf:"bytes,3,opt,name=config" json:"config"`
}

func (m *PrefixConfig) Reset()      { *m = PrefixConfig{} }
func (*PrefixConfig) ProtoMessage() {}

func (m *PrefixConfig) GetPrefix() github_com_cockroachdb_cockroach_proto.Key {
	if m != nil {
		return m.Prefix
	}
	return nil
}

func (m *PrefixConfig) GetCanonical() github_com_cockroachdb_cockroach_proto.Key {
	if m != nil {
		return m.Canonical
	}
	return nil
}

func (m *PrefixConfig) GetConfig() ConfigUnion {
	if m != nil {
		return m.Config
	}
	return ConfigUnion{}
}

type ConfigUnion struct {
	Zone *ZoneConfig `protobuf:"bytes,1,opt,name=zone" json:"zone,omitempty"`
}

func (m *ConfigUnion) Reset()         { *m = ConfigUnion{} }
func (m *ConfigUnion) String() string { return proto.CompactTextString(m) }
func (*ConfigUnion) ProtoMessage()    {}

func (m *ConfigUnion) GetZone() *ZoneConfig {
	if m != nil {
		return m.Zone
	}
	return nil
}

type SystemConfig struct {
	Values []cockroach_proto1.KeyValue `protobuf:"bytes,1,rep,name=values" json:"values"`
}

func (m *SystemConfig) Reset()         { *m = SystemConfig{} }
func (m *SystemConfig) String() string { return proto.CompactTextString(m) }
func (*SystemConfig) ProtoMessage()    {}

func (m *SystemConfig) GetValues() []cockroach_proto1.KeyValue {
	if m != nil {
		return m.Values
	}
	return nil
}

func (m *GCPolicy) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *GCPolicy) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0x8
	i++
	i = encodeVarintConfig(data, i, uint64(m.TTLSeconds))
	return i, nil
}

func (m *ZoneConfig) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *ZoneConfig) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.ReplicaAttrs) > 0 {
		for _, msg := range m.ReplicaAttrs {
			data[i] = 0xa
			i++
			i = encodeVarintConfig(data, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(data[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	data[i] = 0x10
	i++
	i = encodeVarintConfig(data, i, uint64(m.RangeMinBytes))
	data[i] = 0x18
	i++
	i = encodeVarintConfig(data, i, uint64(m.RangeMaxBytes))
	if m.GC != nil {
		data[i] = 0x22
		i++
		i = encodeVarintConfig(data, i, uint64(m.GC.Size()))
		n1, err := m.GC.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	return i, nil
}

func (m *PrefixConfigMap) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *PrefixConfigMap) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Configs) > 0 {
		for _, msg := range m.Configs {
			data[i] = 0xa
			i++
			i = encodeVarintConfig(data, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(data[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *PrefixConfig) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *PrefixConfig) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Prefix != nil {
		data[i] = 0xa
		i++
		i = encodeVarintConfig(data, i, uint64(len(m.Prefix)))
		i += copy(data[i:], m.Prefix)
	}
	if m.Canonical != nil {
		data[i] = 0x12
		i++
		i = encodeVarintConfig(data, i, uint64(len(m.Canonical)))
		i += copy(data[i:], m.Canonical)
	}
	data[i] = 0x1a
	i++
	i = encodeVarintConfig(data, i, uint64(m.Config.Size()))
	n2, err := m.Config.MarshalTo(data[i:])
	if err != nil {
		return 0, err
	}
	i += n2
	return i, nil
}

func (m *ConfigUnion) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *ConfigUnion) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Zone != nil {
		data[i] = 0xa
		i++
		i = encodeVarintConfig(data, i, uint64(m.Zone.Size()))
		n3, err := m.Zone.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	return i, nil
}

func (m *SystemConfig) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *SystemConfig) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Values) > 0 {
		for _, msg := range m.Values {
			data[i] = 0xa
			i++
			i = encodeVarintConfig(data, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(data[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func encodeFixed64Config(data []byte, offset int, v uint64) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	data[offset+4] = uint8(v >> 32)
	data[offset+5] = uint8(v >> 40)
	data[offset+6] = uint8(v >> 48)
	data[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Config(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintConfig(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (m *GCPolicy) Size() (n int) {
	var l int
	_ = l
	n += 1 + sovConfig(uint64(m.TTLSeconds))
	return n
}

func (m *ZoneConfig) Size() (n int) {
	var l int
	_ = l
	if len(m.ReplicaAttrs) > 0 {
		for _, e := range m.ReplicaAttrs {
			l = e.Size()
			n += 1 + l + sovConfig(uint64(l))
		}
	}
	n += 1 + sovConfig(uint64(m.RangeMinBytes))
	n += 1 + sovConfig(uint64(m.RangeMaxBytes))
	if m.GC != nil {
		l = m.GC.Size()
		n += 1 + l + sovConfig(uint64(l))
	}
	return n
}

func (m *PrefixConfigMap) Size() (n int) {
	var l int
	_ = l
	if len(m.Configs) > 0 {
		for _, e := range m.Configs {
			l = e.Size()
			n += 1 + l + sovConfig(uint64(l))
		}
	}
	return n
}

func (m *PrefixConfig) Size() (n int) {
	var l int
	_ = l
	if m.Prefix != nil {
		l = len(m.Prefix)
		n += 1 + l + sovConfig(uint64(l))
	}
	if m.Canonical != nil {
		l = len(m.Canonical)
		n += 1 + l + sovConfig(uint64(l))
	}
	l = m.Config.Size()
	n += 1 + l + sovConfig(uint64(l))
	return n
}

func (m *ConfigUnion) Size() (n int) {
	var l int
	_ = l
	if m.Zone != nil {
		l = m.Zone.Size()
		n += 1 + l + sovConfig(uint64(l))
	}
	return n
}

func (m *SystemConfig) Size() (n int) {
	var l int
	_ = l
	if len(m.Values) > 0 {
		for _, e := range m.Values {
			l = e.Size()
			n += 1 + l + sovConfig(uint64(l))
		}
	}
	return n
}

func sovConfig(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozConfig(x uint64) (n int) {
	return sovConfig(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *ConfigUnion) GetValue() interface{} {
	if this.Zone != nil {
		return this.Zone
	}
	return nil
}

func (this *ConfigUnion) SetValue(value interface{}) bool {
	switch vt := value.(type) {
	case *ZoneConfig:
		this.Zone = vt
	default:
		return false
	}
	return true
}
func (m *GCPolicy) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TTLSeconds", wireType)
			}
			m.TTLSeconds = 0
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.TTLSeconds |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			iNdEx -= sizeOfWire
			skippy, err := skipConfig(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthConfig
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	return nil
}
func (m *ZoneConfig) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReplicaAttrs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ReplicaAttrs = append(m.ReplicaAttrs, cockroach_proto.Attributes{})
			if err := m.ReplicaAttrs[len(m.ReplicaAttrs)-1].Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RangeMinBytes", wireType)
			}
			m.RangeMinBytes = 0
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.RangeMinBytes |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RangeMaxBytes", wireType)
			}
			m.RangeMaxBytes = 0
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.RangeMaxBytes |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GC", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.GC == nil {
				m.GC = &GCPolicy{}
			}
			if err := m.GC.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			iNdEx -= sizeOfWire
			skippy, err := skipConfig(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthConfig
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	return nil
}
func (m *PrefixConfigMap) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Configs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Configs = append(m.Configs, PrefixConfig{})
			if err := m.Configs[len(m.Configs)-1].Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			iNdEx -= sizeOfWire
			skippy, err := skipConfig(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthConfig
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	return nil
}
func (m *PrefixConfig) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Prefix", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Prefix = append([]byte{}, data[iNdEx:postIndex]...)
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Canonical", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Canonical = append([]byte{}, data[iNdEx:postIndex]...)
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Config", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Config.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			iNdEx -= sizeOfWire
			skippy, err := skipConfig(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthConfig
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	return nil
}
func (m *ConfigUnion) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Zone", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Zone == nil {
				m.Zone = &ZoneConfig{}
			}
			if err := m.Zone.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			iNdEx -= sizeOfWire
			skippy, err := skipConfig(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthConfig
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	return nil
}
func (m *SystemConfig) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Values", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Values = append(m.Values, cockroach_proto1.KeyValue{})
			if err := m.Values[len(m.Values)-1].Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			iNdEx -= sizeOfWire
			skippy, err := skipConfig(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthConfig
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	return nil
}
func skipConfig(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for {
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if data[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthConfig
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := data[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipConfig(data[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthConfig = fmt.Errorf("proto: negative length found during unmarshaling")
)
