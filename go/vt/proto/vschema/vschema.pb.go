// Code generated by protoc-gen-go.
// source: vschema.proto
// DO NOT EDIT!

/*
Package vschema is a generated protocol buffer package.

It is generated from these files:
	vschema.proto

It has these top-level messages:
	Keyspace
	Vindex
	Table
	ColumnVindex
	AutoIncrement
	SrvVSchema
*/
package vschema

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

// Keyspace is the vschema for a keyspace.
type Keyspace struct {
	// If sharded is false, vindexes and tables are ignored.
	Sharded  bool               `protobuf:"varint,1,opt,name=sharded" json:"sharded,omitempty"`
	Vindexes map[string]*Vindex `protobuf:"bytes,2,rep,name=vindexes" json:"vindexes,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Tables   map[string]*Table  `protobuf:"bytes,3,rep,name=tables" json:"tables,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Keyspace) Reset()                    { *m = Keyspace{} }
func (m *Keyspace) String() string            { return proto.CompactTextString(m) }
func (*Keyspace) ProtoMessage()               {}
func (*Keyspace) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Keyspace) GetVindexes() map[string]*Vindex {
	if m != nil {
		return m.Vindexes
	}
	return nil
}

func (m *Keyspace) GetTables() map[string]*Table {
	if m != nil {
		return m.Tables
	}
	return nil
}

// Vindex is the vindex info for a Keyspace.
type Vindex struct {
	// The type must match one of the predefined
	// (or plugged in) vindex names.
	Type string `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
	// params is a map of attribute value pairs
	// that must be defined as required by the
	// vindex constructors. The values can only
	// be strings.
	Params map[string]string `protobuf:"bytes,2,rep,name=params" json:"params,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// A lookup vindex can have an owner table defined.
	// If so, rows in the lookup table are created or
	// deleted in sync with corresponding rows in the
	// owner table.
	Owner string `protobuf:"bytes,3,opt,name=owner" json:"owner,omitempty"`
}

func (m *Vindex) Reset()                    { *m = Vindex{} }
func (m *Vindex) String() string            { return proto.CompactTextString(m) }
func (*Vindex) ProtoMessage()               {}
func (*Vindex) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Vindex) GetParams() map[string]string {
	if m != nil {
		return m.Params
	}
	return nil
}

// Table is the table info for a Keyspace.
type Table struct {
	// If the table is a sequence, type must be
	// "sequence". Otherwise, it should be empty.
	Type string `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
	// column_vindexes associates columns to vindexes.
	ColumnVindexes []*ColumnVindex `protobuf:"bytes,2,rep,name=column_vindexes,json=columnVindexes" json:"column_vindexes,omitempty"`
	// auto_increment is specified if a column needs
	// to be associated with a sequence.
	AutoIncrement *AutoIncrement `protobuf:"bytes,3,opt,name=auto_increment,json=autoIncrement" json:"auto_increment,omitempty"`
}

func (m *Table) Reset()                    { *m = Table{} }
func (m *Table) String() string            { return proto.CompactTextString(m) }
func (*Table) ProtoMessage()               {}
func (*Table) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Table) GetColumnVindexes() []*ColumnVindex {
	if m != nil {
		return m.ColumnVindexes
	}
	return nil
}

func (m *Table) GetAutoIncrement() *AutoIncrement {
	if m != nil {
		return m.AutoIncrement
	}
	return nil
}

// ColumnVindex is used to associate a column to a vindex.
type ColumnVindex struct {
	Column string `protobuf:"bytes,1,opt,name=column" json:"column,omitempty"`
	// The name must match a vindex defined in Keyspace.
	Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
}

func (m *ColumnVindex) Reset()                    { *m = ColumnVindex{} }
func (m *ColumnVindex) String() string            { return proto.CompactTextString(m) }
func (*ColumnVindex) ProtoMessage()               {}
func (*ColumnVindex) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

// Autoincrement is used to designate a column as auto-inc.
type AutoIncrement struct {
	Column string `protobuf:"bytes,1,opt,name=column" json:"column,omitempty"`
	// The sequence must match a table of type SEQUENCE.
	Sequence string `protobuf:"bytes,2,opt,name=sequence" json:"sequence,omitempty"`
}

func (m *AutoIncrement) Reset()                    { *m = AutoIncrement{} }
func (m *AutoIncrement) String() string            { return proto.CompactTextString(m) }
func (*AutoIncrement) ProtoMessage()               {}
func (*AutoIncrement) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

// SrvVSchema is the roll-up of all the Keyspace schema for a cell.
type SrvVSchema struct {
	// keyspaces is a map of keyspace name -> Keyspace object.
	Keyspaces map[string]*Keyspace `protobuf:"bytes,1,rep,name=keyspaces" json:"keyspaces,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *SrvVSchema) Reset()                    { *m = SrvVSchema{} }
func (m *SrvVSchema) String() string            { return proto.CompactTextString(m) }
func (*SrvVSchema) ProtoMessage()               {}
func (*SrvVSchema) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *SrvVSchema) GetKeyspaces() map[string]*Keyspace {
	if m != nil {
		return m.Keyspaces
	}
	return nil
}

func init() {
	proto.RegisterType((*Keyspace)(nil), "vschema.Keyspace")
	proto.RegisterType((*Vindex)(nil), "vschema.Vindex")
	proto.RegisterType((*Table)(nil), "vschema.Table")
	proto.RegisterType((*ColumnVindex)(nil), "vschema.ColumnVindex")
	proto.RegisterType((*AutoIncrement)(nil), "vschema.AutoIncrement")
	proto.RegisterType((*SrvVSchema)(nil), "vschema.SrvVSchema")
}

func init() { proto.RegisterFile("vschema.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 433 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0x93, 0x5d, 0xeb, 0xd3, 0x30,
	0x14, 0xc6, 0xe9, 0xe6, 0xfa, 0xef, 0x4e, 0x6d, 0xa7, 0x61, 0x8e, 0x51, 0x11, 0x47, 0x51, 0xdc,
	0x55, 0x2f, 0x36, 0x04, 0x9d, 0x28, 0xca, 0xf0, 0x62, 0x28, 0x28, 0x9d, 0xec, 0x76, 0x64, 0x5d,
	0x60, 0x63, 0x6b, 0x5a, 0xfb, 0xa6, 0xfd, 0x2a, 0xde, 0x08, 0x7e, 0x03, 0xbf, 0xa1, 0x6d, 0x9a,
	0x66, 0xe9, 0x56, 0xef, 0x72, 0x78, 0xce, 0xf3, 0xcb, 0x93, 0x9c, 0x04, 0x8c, 0x2c, 0xf6, 0x0e,
	0xc4, 0xc7, 0x4e, 0x18, 0x05, 0x49, 0x80, 0xee, 0x78, 0x69, 0xff, 0xed, 0x80, 0xf6, 0x89, 0xe4,
	0x71, 0x88, 0x3d, 0x82, 0xc6, 0x70, 0x17, 0x1f, 0x70, 0xb4, 0x27, 0xfb, 0xb1, 0x32, 0x51, 0xa6,
	0x9a, 0x5b, 0x97, 0xe8, 0x0d, 0x68, 0xd9, 0x91, 0xee, 0xc9, 0x4f, 0x12, 0x8f, 0x3b, 0x93, 0xee,
	0x54, 0x9f, 0x3d, 0x75, 0x6a, 0x62, 0x6d, 0x77, 0x36, 0xbc, 0xe3, 0x23, 0x4d, 0xa2, 0xdc, 0x15,
	0x06, 0xf4, 0x12, 0xd4, 0x04, 0xef, 0xce, 0x85, 0xb5, 0xcb, 0xac, 0x4f, 0x6e, 0xad, 0xdf, 0x98,
	0x5e, 0x19, 0x79, 0xb3, 0xf5, 0x19, 0x8c, 0x06, 0x11, 0x3d, 0x80, 0xee, 0x89, 0xe4, 0x2c, 0x5a,
	0xdf, 0x2d, 0x97, 0xe8, 0x39, 0xf4, 0x32, 0x7c, 0x4e, 0x49, 0x91, 0x49, 0x29, 0xc0, 0x03, 0x01,
	0xae, 0x8c, 0x6e, 0xa5, 0x2e, 0x3a, 0xaf, 0x14, 0x6b, 0x05, 0xba, 0xb4, 0x49, 0x0b, 0xeb, 0x59,
	0x93, 0x65, 0x0a, 0x16, 0xb3, 0x49, 0x28, 0xfb, 0x8f, 0x02, 0x6a, 0xb5, 0x01, 0x42, 0x70, 0x2f,
	0xc9, 0x43, 0xc2, 0x39, 0x6c, 0x8d, 0xe6, 0xa0, 0x86, 0x38, 0xc2, 0x7e, 0x7d, 0x53, 0x8f, 0xaf,
	0x52, 0x39, 0x5f, 0x99, 0xca, 0x0f, 0x5b, 0xb5, 0xa2, 0x21, 0xf4, 0x82, 0x1f, 0x94, 0x44, 0xc5,
	0x15, 0x95, 0xa4, 0xaa, 0xb0, 0x5e, 0x83, 0x2e, 0x35, 0xb7, 0x84, 0x1e, 0xca, 0xa1, 0xfb, 0x72,
	0xc8, 0x5f, 0x0a, 0xf4, 0x58, 0xf2, 0xd6, 0x8c, 0xef, 0x60, 0xe0, 0x05, 0xe7, 0xd4, 0xa7, 0xdb,
	0xab, 0xb1, 0x3e, 0x12, 0x61, 0x97, 0x4c, 0xe7, 0x17, 0x69, 0x7a, 0x52, 0x55, 0x8c, 0xf4, 0x2d,
	0x98, 0x38, 0x4d, 0x82, 0xed, 0x91, 0x7a, 0x11, 0xf1, 0x09, 0x4d, 0x58, 0x6e, 0x7d, 0x36, 0x12,
	0xf6, 0x0f, 0x85, 0xbc, 0xaa, 0x55, 0xd7, 0xc0, 0x72, 0x69, 0x2f, 0xe0, 0xbe, 0x8c, 0x47, 0x23,
	0x50, 0xab, 0x0d, 0x78, 0x48, 0x5e, 0x95, 0xd1, 0x29, 0xf6, 0xeb, 0xd3, 0xb1, 0xb5, 0xbd, 0x04,
	0xa3, 0xc1, 0xfe, 0xaf, 0xd9, 0x02, 0x2d, 0x26, 0xdf, 0x53, 0x42, 0xbd, 0x1a, 0x20, 0x6a, 0xfb,
	0xb7, 0x02, 0xb0, 0x8e, 0xb2, 0xcd, 0x9a, 0x85, 0x45, 0xef, 0xa1, 0x7f, 0xe2, 0x4f, 0x31, 0x2e,
	0x28, 0xe5, 0x45, 0xd8, 0xe2, 0x24, 0x97, 0x3e, 0xf1, 0x5e, 0xf9, 0xf0, 0x2e, 0x26, 0xeb, 0x0b,
	0x98, 0x4d, 0xb1, 0x65, 0x58, 0x2f, 0x9a, 0x2f, 0xec, 0xe1, 0xcd, 0x37, 0x90, 0xe6, 0xb7, 0x53,
	0xd9, 0x47, 0x9d, 0xff, 0x0b, 0x00, 0x00, 0xff, 0xff, 0xf1, 0x47, 0x7e, 0x27, 0xb9, 0x03, 0x00,
	0x00,
}