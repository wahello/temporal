// The MIT License
//
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.
//
// Copyright (c) 2020 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: temporal/server/api/persistence/v1/cluster_metadata.proto

package persistence

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"

	proto "github.com/gogo/protobuf/proto"
	v1 "go.temporal.io/api/version/v1"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// data column
type ClusterMetadata struct {
	ClusterName       string          `protobuf:"bytes,1,opt,name=cluster_name,json=clusterName,proto3" json:"cluster_name,omitempty"`
	HistoryShardCount int32           `protobuf:"varint,2,opt,name=history_shard_count,json=historyShardCount,proto3" json:"history_shard_count,omitempty"`
	ClusterId         string          `protobuf:"bytes,3,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	VersionInfo       *v1.VersionInfo `protobuf:"bytes,4,opt,name=version_info,json=versionInfo,proto3" json:"version_info,omitempty"`
}

func (m *ClusterMetadata) Reset()      { *m = ClusterMetadata{} }
func (*ClusterMetadata) ProtoMessage() {}
func (*ClusterMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_1f4771d63f405884, []int{0}
}
func (m *ClusterMetadata) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ClusterMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ClusterMetadata.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ClusterMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterMetadata.Merge(m, src)
}
func (m *ClusterMetadata) XXX_Size() int {
	return m.Size()
}
func (m *ClusterMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterMetadata proto.InternalMessageInfo

func (m *ClusterMetadata) GetClusterName() string {
	if m != nil {
		return m.ClusterName
	}
	return ""
}

func (m *ClusterMetadata) GetHistoryShardCount() int32 {
	if m != nil {
		return m.HistoryShardCount
	}
	return 0
}

func (m *ClusterMetadata) GetClusterId() string {
	if m != nil {
		return m.ClusterId
	}
	return ""
}

func (m *ClusterMetadata) GetVersionInfo() *v1.VersionInfo {
	if m != nil {
		return m.VersionInfo
	}
	return nil
}

func init() {
	proto.RegisterType((*ClusterMetadata)(nil), "temporal.server.api.persistence.v1.ClusterMetadata")
}

func init() {
	proto.RegisterFile("temporal/server/api/persistence/v1/cluster_metadata.proto", fileDescriptor_1f4771d63f405884)
}

var fileDescriptor_1f4771d63f405884 = []byte{
	// 330 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xbf, 0x4e, 0xf3, 0x30,
	0x14, 0xc5, 0xe3, 0xef, 0x03, 0xa4, 0xba, 0x95, 0x10, 0x61, 0x89, 0x90, 0xb8, 0x2a, 0x15, 0x48,
	0x9d, 0x1c, 0x15, 0x58, 0x10, 0x1b, 0x1d, 0x50, 0x07, 0x18, 0x8a, 0xc4, 0xc0, 0x12, 0x99, 0xe6,
	0xb6, 0x35, 0x6a, 0xe2, 0xc8, 0x76, 0x23, 0xb1, 0xf1, 0x08, 0x3c, 0x06, 0x8f, 0xc2, 0xc0, 0xd0,
	0xb1, 0x23, 0x75, 0x17, 0xc6, 0x3e, 0x02, 0x4a, 0xe3, 0xfe, 0x59, 0xd8, 0x7c, 0xef, 0xf1, 0xf9,
	0x1d, 0x5f, 0x5f, 0x7a, 0x65, 0x30, 0xc9, 0xa4, 0xe2, 0xa3, 0x50, 0xa3, 0xca, 0x51, 0x85, 0x3c,
	0x13, 0x61, 0x86, 0x4a, 0x0b, 0x6d, 0x30, 0xed, 0x61, 0x98, 0xb7, 0xc2, 0xde, 0x68, 0xac, 0x0d,
	0xaa, 0x28, 0x41, 0xc3, 0x63, 0x6e, 0x38, 0xcb, 0x94, 0x34, 0xd2, 0x6f, 0xac, 0xac, 0xac, 0xb4,
	0x32, 0x9e, 0x09, 0xb6, 0x65, 0x65, 0x79, 0xeb, 0xe8, 0x6c, 0x8d, 0x2f, 0xb8, 0x79, 0x21, 0xca,
	0xb4, 0x60, 0x26, 0xa8, 0x35, 0x1f, 0x60, 0x89, 0x6a, 0x7c, 0x11, 0xba, 0xdf, 0x2e, 0x53, 0xee,
	0x5c, 0x88, 0x7f, 0x42, 0x6b, 0xab, 0xe0, 0x94, 0x27, 0x18, 0x90, 0x3a, 0x69, 0x56, 0xba, 0x55,
	0xd7, 0xbb, 0xe7, 0x09, 0xfa, 0x8c, 0x1e, 0x0e, 0x85, 0x36, 0x52, 0xbd, 0x46, 0x7a, 0xc8, 0x55,
	0x1c, 0xf5, 0xe4, 0x38, 0x35, 0xc1, 0xbf, 0x3a, 0x69, 0xee, 0x76, 0x0f, 0x9c, 0xf4, 0x50, 0x28,
	0xed, 0x42, 0xf0, 0x8f, 0x29, 0x5d, 0x21, 0x45, 0x1c, 0xfc, 0x5f, 0x02, 0x2b, 0xae, 0xd3, 0x89,
	0xfd, 0x5b, 0x5a, 0x73, 0x2f, 0x8c, 0x44, 0xda, 0x97, 0xc1, 0x4e, 0x9d, 0x34, 0xab, 0xe7, 0xa7,
	0x6c, 0x3d, 0x67, 0x31, 0xa0, 0xbb, 0xc1, 0xf2, 0x16, 0x7b, 0x2c, 0x8f, 0x9d, 0xb4, 0x2f, 0xbb,
	0xd5, 0x7c, 0x53, 0xdc, 0xbc, 0x4c, 0x66, 0xe0, 0x4d, 0x67, 0xe0, 0x2d, 0x66, 0x40, 0xde, 0x2c,
	0x90, 0x0f, 0x0b, 0xe4, 0xd3, 0x02, 0x99, 0x58, 0x20, 0xdf, 0x16, 0xc8, 0x8f, 0x05, 0x6f, 0x61,
	0x81, 0xbc, 0xcf, 0xc1, 0x9b, 0xcc, 0xc1, 0x9b, 0xce, 0xc1, 0x7b, 0xba, 0x1c, 0xc8, 0x4d, 0x94,
	0x90, 0x7f, 0x2f, 0xe4, 0x7a, 0xab, 0x7c, 0xde, 0x5b, 0xfe, 0xe0, 0xc5, 0x6f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xeb, 0xca, 0x77, 0xc3, 0xc9, 0x01, 0x00, 0x00,
}

func (this *ClusterMetadata) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ClusterMetadata)
	if !ok {
		that2, ok := that.(ClusterMetadata)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.ClusterName != that1.ClusterName {
		return false
	}
	if this.HistoryShardCount != that1.HistoryShardCount {
		return false
	}
	if this.ClusterId != that1.ClusterId {
		return false
	}
	if !this.VersionInfo.Equal(that1.VersionInfo) {
		return false
	}
	return true
}
func (this *ClusterMetadata) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 8)
	s = append(s, "&persistence.ClusterMetadata{")
	s = append(s, "ClusterName: "+fmt.Sprintf("%#v", this.ClusterName)+",\n")
	s = append(s, "HistoryShardCount: "+fmt.Sprintf("%#v", this.HistoryShardCount)+",\n")
	s = append(s, "ClusterId: "+fmt.Sprintf("%#v", this.ClusterId)+",\n")
	if this.VersionInfo != nil {
		s = append(s, "VersionInfo: "+fmt.Sprintf("%#v", this.VersionInfo)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringClusterMetadata(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *ClusterMetadata) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ClusterMetadata) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ClusterMetadata) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.VersionInfo != nil {
		{
			size, err := m.VersionInfo.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintClusterMetadata(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if len(m.ClusterId) > 0 {
		i -= len(m.ClusterId)
		copy(dAtA[i:], m.ClusterId)
		i = encodeVarintClusterMetadata(dAtA, i, uint64(len(m.ClusterId)))
		i--
		dAtA[i] = 0x1a
	}
	if m.HistoryShardCount != 0 {
		i = encodeVarintClusterMetadata(dAtA, i, uint64(m.HistoryShardCount))
		i--
		dAtA[i] = 0x10
	}
	if len(m.ClusterName) > 0 {
		i -= len(m.ClusterName)
		copy(dAtA[i:], m.ClusterName)
		i = encodeVarintClusterMetadata(dAtA, i, uint64(len(m.ClusterName)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintClusterMetadata(dAtA []byte, offset int, v uint64) int {
	offset -= sovClusterMetadata(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ClusterMetadata) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ClusterName)
	if l > 0 {
		n += 1 + l + sovClusterMetadata(uint64(l))
	}
	if m.HistoryShardCount != 0 {
		n += 1 + sovClusterMetadata(uint64(m.HistoryShardCount))
	}
	l = len(m.ClusterId)
	if l > 0 {
		n += 1 + l + sovClusterMetadata(uint64(l))
	}
	if m.VersionInfo != nil {
		l = m.VersionInfo.Size()
		n += 1 + l + sovClusterMetadata(uint64(l))
	}
	return n
}

func sovClusterMetadata(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozClusterMetadata(x uint64) (n int) {
	return sovClusterMetadata(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *ClusterMetadata) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ClusterMetadata{`,
		`ClusterName:` + fmt.Sprintf("%v", this.ClusterName) + `,`,
		`HistoryShardCount:` + fmt.Sprintf("%v", this.HistoryShardCount) + `,`,
		`ClusterId:` + fmt.Sprintf("%v", this.ClusterId) + `,`,
		`VersionInfo:` + strings.Replace(fmt.Sprintf("%v", this.VersionInfo), "VersionInfo", "v1.VersionInfo", 1) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringClusterMetadata(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *ClusterMetadata) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowClusterMetadata
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ClusterMetadata: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClusterMetadata: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClusterName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClusterMetadata
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthClusterMetadata
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthClusterMetadata
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClusterName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field HistoryShardCount", wireType)
			}
			m.HistoryShardCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClusterMetadata
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.HistoryShardCount |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClusterId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClusterMetadata
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthClusterMetadata
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthClusterMetadata
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClusterId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VersionInfo", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClusterMetadata
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthClusterMetadata
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthClusterMetadata
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.VersionInfo == nil {
				m.VersionInfo = &v1.VersionInfo{}
			}
			if err := m.VersionInfo.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipClusterMetadata(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthClusterMetadata
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthClusterMetadata
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipClusterMetadata(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowClusterMetadata
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowClusterMetadata
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowClusterMetadata
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthClusterMetadata
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupClusterMetadata
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthClusterMetadata
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthClusterMetadata        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowClusterMetadata          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupClusterMetadata = fmt.Errorf("proto: unexpected end of group")
)
