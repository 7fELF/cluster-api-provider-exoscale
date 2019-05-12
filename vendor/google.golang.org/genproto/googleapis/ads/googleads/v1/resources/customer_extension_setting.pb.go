// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/resources/customer_extension_setting.proto

package resources // import "google.golang.org/genproto/googleapis/ads/googleads/v1/resources"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import wrappers "github.com/golang/protobuf/ptypes/wrappers"
import enums "google.golang.org/genproto/googleapis/ads/googleads/v1/enums"
import _ "google.golang.org/genproto/googleapis/api/annotations"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// A customer extension setting.
type CustomerExtensionSetting struct {
	// The resource name of the customer extension setting.
	// CustomerExtensionSetting resource names have the form:
	//
	// `customers/{customer_id}/customerExtensionSettings/{extension_type}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The extension type of the customer extension setting.
	ExtensionType enums.ExtensionTypeEnum_ExtensionType `protobuf:"varint,2,opt,name=extension_type,json=extensionType,proto3,enum=google.ads.googleads.v1.enums.ExtensionTypeEnum_ExtensionType" json:"extension_type,omitempty"`
	// The resource names of the extension feed items to serve under the customer.
	// ExtensionFeedItem resource names have the form:
	//
	// `customers/{customer_id}/extensionFeedItems/{feed_item_id}`
	ExtensionFeedItems []*wrappers.StringValue `protobuf:"bytes,3,rep,name=extension_feed_items,json=extensionFeedItems,proto3" json:"extension_feed_items,omitempty"`
	// The device for which the extensions will serve. Optional.
	Device               enums.ExtensionSettingDeviceEnum_ExtensionSettingDevice `protobuf:"varint,4,opt,name=device,proto3,enum=google.ads.googleads.v1.enums.ExtensionSettingDeviceEnum_ExtensionSettingDevice" json:"device,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                                `json:"-"`
	XXX_unrecognized     []byte                                                  `json:"-"`
	XXX_sizecache        int32                                                   `json:"-"`
}

func (m *CustomerExtensionSetting) Reset()         { *m = CustomerExtensionSetting{} }
func (m *CustomerExtensionSetting) String() string { return proto.CompactTextString(m) }
func (*CustomerExtensionSetting) ProtoMessage()    {}
func (*CustomerExtensionSetting) Descriptor() ([]byte, []int) {
	return fileDescriptor_customer_extension_setting_36805005c826cc7e, []int{0}
}
func (m *CustomerExtensionSetting) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CustomerExtensionSetting.Unmarshal(m, b)
}
func (m *CustomerExtensionSetting) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CustomerExtensionSetting.Marshal(b, m, deterministic)
}
func (dst *CustomerExtensionSetting) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CustomerExtensionSetting.Merge(dst, src)
}
func (m *CustomerExtensionSetting) XXX_Size() int {
	return xxx_messageInfo_CustomerExtensionSetting.Size(m)
}
func (m *CustomerExtensionSetting) XXX_DiscardUnknown() {
	xxx_messageInfo_CustomerExtensionSetting.DiscardUnknown(m)
}

var xxx_messageInfo_CustomerExtensionSetting proto.InternalMessageInfo

func (m *CustomerExtensionSetting) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *CustomerExtensionSetting) GetExtensionType() enums.ExtensionTypeEnum_ExtensionType {
	if m != nil {
		return m.ExtensionType
	}
	return enums.ExtensionTypeEnum_UNSPECIFIED
}

func (m *CustomerExtensionSetting) GetExtensionFeedItems() []*wrappers.StringValue {
	if m != nil {
		return m.ExtensionFeedItems
	}
	return nil
}

func (m *CustomerExtensionSetting) GetDevice() enums.ExtensionSettingDeviceEnum_ExtensionSettingDevice {
	if m != nil {
		return m.Device
	}
	return enums.ExtensionSettingDeviceEnum_UNSPECIFIED
}

func init() {
	proto.RegisterType((*CustomerExtensionSetting)(nil), "google.ads.googleads.v1.resources.CustomerExtensionSetting")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/resources/customer_extension_setting.proto", fileDescriptor_customer_extension_setting_36805005c826cc7e)
}

var fileDescriptor_customer_extension_setting_36805005c826cc7e = []byte{
	// 431 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0xc1, 0x6e, 0xd4, 0x30,
	0x10, 0xd5, 0x66, 0x51, 0x25, 0x02, 0xed, 0x21, 0xe2, 0x10, 0x55, 0x05, 0x6d, 0x41, 0x95, 0xf6,
	0x64, 0x2b, 0xcb, 0xcd, 0x20, 0xa4, 0x2c, 0x94, 0x0a, 0x0e, 0xd5, 0x2a, 0x45, 0x7b, 0x40, 0x2b,
	0x45, 0xee, 0x7a, 0x6a, 0x2c, 0x6d, 0xec, 0xc8, 0x76, 0x16, 0xfa, 0x0b, 0x7c, 0x06, 0x47, 0x3e,
	0x85, 0xaf, 0xe0, 0xcc, 0x57, 0xa0, 0xc4, 0xb1, 0x21, 0x42, 0x4b, 0xf7, 0x36, 0xf1, 0xbc, 0xf7,
	0xe6, 0xbd, 0xc9, 0xc4, 0x73, 0xae, 0x14, 0xdf, 0x00, 0xa6, 0xcc, 0x60, 0x57, 0xb6, 0xd5, 0x36,
	0xc3, 0x1a, 0x8c, 0x6a, 0xf4, 0x1a, 0x0c, 0x5e, 0x37, 0xc6, 0xaa, 0x0a, 0x74, 0x09, 0x5f, 0x2c,
	0x48, 0x23, 0x94, 0x2c, 0x0d, 0x58, 0x2b, 0x24, 0x47, 0xb5, 0x56, 0x56, 0x25, 0xa7, 0x8e, 0x88,
	0x28, 0x33, 0x28, 0x68, 0xa0, 0x6d, 0x86, 0x82, 0xc6, 0xf1, 0xcb, 0x5d, 0x63, 0x40, 0x36, 0x95,
	0xc1, 0xff, 0x28, 0x97, 0x0c, 0xb6, 0x62, 0x0d, 0x6e, 0xc0, 0xf1, 0x6c, 0x5f, 0xb6, 0xbd, 0xad,
	0x3d, 0xe7, 0x49, 0xcf, 0xe9, 0xbe, 0xae, 0x9b, 0x1b, 0xfc, 0x59, 0xd3, 0xba, 0x06, 0x6d, 0xfa,
	0xfe, 0x89, 0xd7, 0xac, 0x05, 0xa6, 0x52, 0x2a, 0x4b, 0xad, 0x50, 0xb2, 0xef, 0x3e, 0xfd, 0x19,
	0xc5, 0xe9, 0xeb, 0x3e, 0xf7, 0xb9, 0x97, 0xbf, 0x72, 0xde, 0x92, 0x67, 0xf1, 0xa1, 0x4f, 0x56,
	0x4a, 0x5a, 0x41, 0x3a, 0x9a, 0x8c, 0xa6, 0xf7, 0x8b, 0x87, 0xfe, 0xf1, 0x92, 0x56, 0x90, 0x40,
	0x7c, 0x34, 0xf4, 0x95, 0x46, 0x93, 0xd1, 0xf4, 0x68, 0xf6, 0x0a, 0xed, 0xda, 0x56, 0x17, 0x06,
	0x85, 0x69, 0x1f, 0x6e, 0x6b, 0x38, 0x97, 0x4d, 0x35, 0x7c, 0x29, 0x0e, 0xe1, 0xef, 0xcf, 0xe4,
	0x32, 0x7e, 0xf4, 0x67, 0xcc, 0x0d, 0x00, 0x2b, 0x85, 0x85, 0xca, 0xa4, 0xe3, 0xc9, 0x78, 0xfa,
	0x60, 0x76, 0xe2, 0x87, 0xf9, 0x2d, 0xa0, 0x2b, 0xab, 0x85, 0xe4, 0x4b, 0xba, 0x69, 0xa0, 0x48,
	0x02, 0xf3, 0x2d, 0x00, 0x7b, 0xd7, 0xf2, 0x92, 0x4f, 0xf1, 0x81, 0x5b, 0x7d, 0x7a, 0xaf, 0xb3,
	0xbb, 0xd8, 0xd7, 0x6e, 0xbf, 0x9c, 0x37, 0x1d, 0x79, 0xe8, 0x7b, 0xd0, 0x2a, 0x7a, 0xfd, 0xf9,
	0xd7, 0x28, 0x3e, 0x5b, 0xab, 0x0a, 0xdd, 0x79, 0x3c, 0xf3, 0xc7, 0xbb, 0xfe, 0xc4, 0xa2, 0x4d,
	0xb5, 0x18, 0x7d, 0x7c, 0xdf, 0x6b, 0x70, 0xb5, 0xa1, 0x92, 0x23, 0xa5, 0x39, 0xe6, 0x20, 0xbb,
	0xcc, 0xfe, 0x5e, 0x6a, 0x61, 0xfe, 0x73, 0xe3, 0x2f, 0x42, 0xf5, 0x2d, 0x1a, 0x5f, 0xe4, 0xf9,
	0xf7, 0xe8, 0xf4, 0xc2, 0x49, 0xe6, 0xcc, 0x20, 0x57, 0xb6, 0xd5, 0x32, 0x43, 0x85, 0x47, 0xfe,
	0xf0, 0x98, 0x55, 0xce, 0xcc, 0x2a, 0x60, 0x56, 0xcb, 0x6c, 0x15, 0x30, 0xbf, 0xa2, 0x33, 0xd7,
	0x20, 0x24, 0x67, 0x86, 0x90, 0x80, 0x22, 0x64, 0x99, 0x11, 0x12, 0x70, 0xd7, 0x07, 0x9d, 0xd9,
	0xe7, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x13, 0xdf, 0xd9, 0x40, 0x8f, 0x03, 0x00, 0x00,
}
