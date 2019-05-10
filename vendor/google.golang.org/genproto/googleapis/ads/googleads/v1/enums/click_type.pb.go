// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/enums/click_type.proto

package enums // import "google.golang.org/genproto/googleapis/ads/googleads/v1/enums"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
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

// Enumerates Google Ads click types.
type ClickTypeEnum_ClickType int32

const (
	// Not specified.
	ClickTypeEnum_UNSPECIFIED ClickTypeEnum_ClickType = 0
	// The value is unknown in this version.
	ClickTypeEnum_UNKNOWN ClickTypeEnum_ClickType = 1
	// App engagement ad deep link.
	ClickTypeEnum_APP_DEEPLINK ClickTypeEnum_ClickType = 2
	// Breadcrumbs.
	ClickTypeEnum_BREADCRUMBS ClickTypeEnum_ClickType = 3
	// Broadband Plan.
	ClickTypeEnum_BROADBAND_PLAN ClickTypeEnum_ClickType = 4
	// Manually dialed phone calls.
	ClickTypeEnum_CALL_TRACKING ClickTypeEnum_ClickType = 5
	// Phone calls.
	ClickTypeEnum_CALLS ClickTypeEnum_ClickType = 6
	// Click on engagement ad.
	ClickTypeEnum_CLICK_ON_ENGAGEMENT_AD ClickTypeEnum_ClickType = 7
	// Driving direction.
	ClickTypeEnum_GET_DIRECTIONS ClickTypeEnum_ClickType = 8
	// Get location details.
	ClickTypeEnum_LOCATION_EXPANSION ClickTypeEnum_ClickType = 9
	// Call.
	ClickTypeEnum_LOCATION_FORMAT_CALL ClickTypeEnum_ClickType = 10
	// Directions.
	ClickTypeEnum_LOCATION_FORMAT_DIRECTIONS ClickTypeEnum_ClickType = 11
	// Image(s).
	ClickTypeEnum_LOCATION_FORMAT_IMAGE ClickTypeEnum_ClickType = 12
	// Go to landing page.
	ClickTypeEnum_LOCATION_FORMAT_LANDING_PAGE ClickTypeEnum_ClickType = 13
	// Map.
	ClickTypeEnum_LOCATION_FORMAT_MAP ClickTypeEnum_ClickType = 14
	// Go to store info.
	ClickTypeEnum_LOCATION_FORMAT_STORE_INFO ClickTypeEnum_ClickType = 15
	// Text.
	ClickTypeEnum_LOCATION_FORMAT_TEXT ClickTypeEnum_ClickType = 16
	// Mobile phone calls.
	ClickTypeEnum_MOBILE_CALL_TRACKING ClickTypeEnum_ClickType = 17
	// Print offer.
	ClickTypeEnum_OFFER_PRINTS ClickTypeEnum_ClickType = 18
	// Other.
	ClickTypeEnum_OTHER ClickTypeEnum_ClickType = 19
	// Product plusbox offer.
	ClickTypeEnum_PRODUCT_EXTENSION_CLICKS ClickTypeEnum_ClickType = 20
	// Shopping - Product - Online.
	ClickTypeEnum_PRODUCT_LISTING_AD_CLICKS ClickTypeEnum_ClickType = 21
	// Sitelink.
	ClickTypeEnum_SITELINKS ClickTypeEnum_ClickType = 22
	// Show nearby locations.
	ClickTypeEnum_STORE_LOCATOR ClickTypeEnum_ClickType = 23
	// Headline.
	ClickTypeEnum_URL_CLICKS ClickTypeEnum_ClickType = 25
	// App store.
	ClickTypeEnum_VIDEO_APP_STORE_CLICKS ClickTypeEnum_ClickType = 26
	// Call-to-Action overlay.
	ClickTypeEnum_VIDEO_CALL_TO_ACTION_CLICKS ClickTypeEnum_ClickType = 27
	// Cards.
	ClickTypeEnum_VIDEO_CARD_ACTION_HEADLINE_CLICKS ClickTypeEnum_ClickType = 28
	// End cap.
	ClickTypeEnum_VIDEO_END_CAP_CLICKS ClickTypeEnum_ClickType = 29
	// Website.
	ClickTypeEnum_VIDEO_WEBSITE_CLICKS ClickTypeEnum_ClickType = 30
	// Visual Sitelinks.
	ClickTypeEnum_VISUAL_SITELINKS ClickTypeEnum_ClickType = 31
	// Wireless Plan.
	ClickTypeEnum_WIRELESS_PLAN ClickTypeEnum_ClickType = 32
	// Shopping - Product - Local.
	ClickTypeEnum_PRODUCT_LISTING_AD_LOCAL ClickTypeEnum_ClickType = 33
	// Shopping - Product - MultiChannel Local.
	ClickTypeEnum_PRODUCT_LISTING_AD_MULTICHANNEL_LOCAL ClickTypeEnum_ClickType = 34
	// Shopping - Product - MultiChannel Online.
	ClickTypeEnum_PRODUCT_LISTING_AD_MULTICHANNEL_ONLINE ClickTypeEnum_ClickType = 35
	// Shopping - Product - Coupon.
	ClickTypeEnum_PRODUCT_LISTING_ADS_COUPON ClickTypeEnum_ClickType = 36
	// Shopping - Product - Sell on Google.
	ClickTypeEnum_PRODUCT_LISTING_AD_TRANSACTABLE ClickTypeEnum_ClickType = 37
	// Shopping - Product - App engagement ad deep link.
	ClickTypeEnum_PRODUCT_AD_APP_DEEPLINK ClickTypeEnum_ClickType = 38
	// Shopping - Showcase - Category.
	ClickTypeEnum_SHOWCASE_AD_CATEGORY_LINK ClickTypeEnum_ClickType = 39
	// Shopping - Showcase - Local storefront.
	ClickTypeEnum_SHOWCASE_AD_LOCAL_STOREFRONT_LINK ClickTypeEnum_ClickType = 40
	// Shopping - Showcase - Online product.
	ClickTypeEnum_SHOWCASE_AD_ONLINE_PRODUCT_LINK ClickTypeEnum_ClickType = 42
	// Shopping - Showcase - Local product.
	ClickTypeEnum_SHOWCASE_AD_LOCAL_PRODUCT_LINK ClickTypeEnum_ClickType = 43
	// Promotion Extension.
	ClickTypeEnum_PROMOTION_EXTENSION ClickTypeEnum_ClickType = 44
	// Ad Headline.
	ClickTypeEnum_SWIPEABLE_GALLERY_AD_HEADLINE ClickTypeEnum_ClickType = 45
	// Swipes.
	ClickTypeEnum_SWIPEABLE_GALLERY_AD_SWIPES ClickTypeEnum_ClickType = 46
	// See More.
	ClickTypeEnum_SWIPEABLE_GALLERY_AD_SEE_MORE ClickTypeEnum_ClickType = 47
	// Sitelink 1.
	ClickTypeEnum_SWIPEABLE_GALLERY_AD_SITELINK_ONE ClickTypeEnum_ClickType = 48
	// Sitelink 2.
	ClickTypeEnum_SWIPEABLE_GALLERY_AD_SITELINK_TWO ClickTypeEnum_ClickType = 49
	// Sitelink 3.
	ClickTypeEnum_SWIPEABLE_GALLERY_AD_SITELINK_THREE ClickTypeEnum_ClickType = 50
	// Sitelink 4.
	ClickTypeEnum_SWIPEABLE_GALLERY_AD_SITELINK_FOUR ClickTypeEnum_ClickType = 51
	// Sitelink 5.
	ClickTypeEnum_SWIPEABLE_GALLERY_AD_SITELINK_FIVE ClickTypeEnum_ClickType = 52
	// Hotel price.
	ClickTypeEnum_HOTEL_PRICE ClickTypeEnum_ClickType = 53
	// Price Extension.
	ClickTypeEnum_PRICE_EXTENSION ClickTypeEnum_ClickType = 54
	// Book on Google hotel room selection.
	ClickTypeEnum_HOTEL_BOOK_ON_GOOGLE_ROOM_SELECTION ClickTypeEnum_ClickType = 55
)

var ClickTypeEnum_ClickType_name = map[int32]string{
	0:  "UNSPECIFIED",
	1:  "UNKNOWN",
	2:  "APP_DEEPLINK",
	3:  "BREADCRUMBS",
	4:  "BROADBAND_PLAN",
	5:  "CALL_TRACKING",
	6:  "CALLS",
	7:  "CLICK_ON_ENGAGEMENT_AD",
	8:  "GET_DIRECTIONS",
	9:  "LOCATION_EXPANSION",
	10: "LOCATION_FORMAT_CALL",
	11: "LOCATION_FORMAT_DIRECTIONS",
	12: "LOCATION_FORMAT_IMAGE",
	13: "LOCATION_FORMAT_LANDING_PAGE",
	14: "LOCATION_FORMAT_MAP",
	15: "LOCATION_FORMAT_STORE_INFO",
	16: "LOCATION_FORMAT_TEXT",
	17: "MOBILE_CALL_TRACKING",
	18: "OFFER_PRINTS",
	19: "OTHER",
	20: "PRODUCT_EXTENSION_CLICKS",
	21: "PRODUCT_LISTING_AD_CLICKS",
	22: "SITELINKS",
	23: "STORE_LOCATOR",
	25: "URL_CLICKS",
	26: "VIDEO_APP_STORE_CLICKS",
	27: "VIDEO_CALL_TO_ACTION_CLICKS",
	28: "VIDEO_CARD_ACTION_HEADLINE_CLICKS",
	29: "VIDEO_END_CAP_CLICKS",
	30: "VIDEO_WEBSITE_CLICKS",
	31: "VISUAL_SITELINKS",
	32: "WIRELESS_PLAN",
	33: "PRODUCT_LISTING_AD_LOCAL",
	34: "PRODUCT_LISTING_AD_MULTICHANNEL_LOCAL",
	35: "PRODUCT_LISTING_AD_MULTICHANNEL_ONLINE",
	36: "PRODUCT_LISTING_ADS_COUPON",
	37: "PRODUCT_LISTING_AD_TRANSACTABLE",
	38: "PRODUCT_AD_APP_DEEPLINK",
	39: "SHOWCASE_AD_CATEGORY_LINK",
	40: "SHOWCASE_AD_LOCAL_STOREFRONT_LINK",
	42: "SHOWCASE_AD_ONLINE_PRODUCT_LINK",
	43: "SHOWCASE_AD_LOCAL_PRODUCT_LINK",
	44: "PROMOTION_EXTENSION",
	45: "SWIPEABLE_GALLERY_AD_HEADLINE",
	46: "SWIPEABLE_GALLERY_AD_SWIPES",
	47: "SWIPEABLE_GALLERY_AD_SEE_MORE",
	48: "SWIPEABLE_GALLERY_AD_SITELINK_ONE",
	49: "SWIPEABLE_GALLERY_AD_SITELINK_TWO",
	50: "SWIPEABLE_GALLERY_AD_SITELINK_THREE",
	51: "SWIPEABLE_GALLERY_AD_SITELINK_FOUR",
	52: "SWIPEABLE_GALLERY_AD_SITELINK_FIVE",
	53: "HOTEL_PRICE",
	54: "PRICE_EXTENSION",
	55: "HOTEL_BOOK_ON_GOOGLE_ROOM_SELECTION",
}
var ClickTypeEnum_ClickType_value = map[string]int32{
	"UNSPECIFIED":                            0,
	"UNKNOWN":                                1,
	"APP_DEEPLINK":                           2,
	"BREADCRUMBS":                            3,
	"BROADBAND_PLAN":                         4,
	"CALL_TRACKING":                          5,
	"CALLS":                                  6,
	"CLICK_ON_ENGAGEMENT_AD":                 7,
	"GET_DIRECTIONS":                         8,
	"LOCATION_EXPANSION":                     9,
	"LOCATION_FORMAT_CALL":                   10,
	"LOCATION_FORMAT_DIRECTIONS":             11,
	"LOCATION_FORMAT_IMAGE":                  12,
	"LOCATION_FORMAT_LANDING_PAGE":           13,
	"LOCATION_FORMAT_MAP":                    14,
	"LOCATION_FORMAT_STORE_INFO":             15,
	"LOCATION_FORMAT_TEXT":                   16,
	"MOBILE_CALL_TRACKING":                   17,
	"OFFER_PRINTS":                           18,
	"OTHER":                                  19,
	"PRODUCT_EXTENSION_CLICKS":               20,
	"PRODUCT_LISTING_AD_CLICKS":              21,
	"SITELINKS":                              22,
	"STORE_LOCATOR":                          23,
	"URL_CLICKS":                             25,
	"VIDEO_APP_STORE_CLICKS":                 26,
	"VIDEO_CALL_TO_ACTION_CLICKS":            27,
	"VIDEO_CARD_ACTION_HEADLINE_CLICKS":      28,
	"VIDEO_END_CAP_CLICKS":                   29,
	"VIDEO_WEBSITE_CLICKS":                   30,
	"VISUAL_SITELINKS":                       31,
	"WIRELESS_PLAN":                          32,
	"PRODUCT_LISTING_AD_LOCAL":               33,
	"PRODUCT_LISTING_AD_MULTICHANNEL_LOCAL":  34,
	"PRODUCT_LISTING_AD_MULTICHANNEL_ONLINE": 35,
	"PRODUCT_LISTING_ADS_COUPON":             36,
	"PRODUCT_LISTING_AD_TRANSACTABLE":        37,
	"PRODUCT_AD_APP_DEEPLINK":                38,
	"SHOWCASE_AD_CATEGORY_LINK":              39,
	"SHOWCASE_AD_LOCAL_STOREFRONT_LINK":      40,
	"SHOWCASE_AD_ONLINE_PRODUCT_LINK":        42,
	"SHOWCASE_AD_LOCAL_PRODUCT_LINK":         43,
	"PROMOTION_EXTENSION":                    44,
	"SWIPEABLE_GALLERY_AD_HEADLINE":          45,
	"SWIPEABLE_GALLERY_AD_SWIPES":            46,
	"SWIPEABLE_GALLERY_AD_SEE_MORE":          47,
	"SWIPEABLE_GALLERY_AD_SITELINK_ONE":      48,
	"SWIPEABLE_GALLERY_AD_SITELINK_TWO":      49,
	"SWIPEABLE_GALLERY_AD_SITELINK_THREE":    50,
	"SWIPEABLE_GALLERY_AD_SITELINK_FOUR":     51,
	"SWIPEABLE_GALLERY_AD_SITELINK_FIVE":     52,
	"HOTEL_PRICE":                            53,
	"PRICE_EXTENSION":                        54,
	"HOTEL_BOOK_ON_GOOGLE_ROOM_SELECTION":    55,
}

func (x ClickTypeEnum_ClickType) String() string {
	return proto.EnumName(ClickTypeEnum_ClickType_name, int32(x))
}
func (ClickTypeEnum_ClickType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_click_type_094021787dbc0905, []int{0, 0}
}

// Container for enumeration of Google Ads click types.
type ClickTypeEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClickTypeEnum) Reset()         { *m = ClickTypeEnum{} }
func (m *ClickTypeEnum) String() string { return proto.CompactTextString(m) }
func (*ClickTypeEnum) ProtoMessage()    {}
func (*ClickTypeEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_click_type_094021787dbc0905, []int{0}
}
func (m *ClickTypeEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClickTypeEnum.Unmarshal(m, b)
}
func (m *ClickTypeEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClickTypeEnum.Marshal(b, m, deterministic)
}
func (dst *ClickTypeEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClickTypeEnum.Merge(dst, src)
}
func (m *ClickTypeEnum) XXX_Size() int {
	return xxx_messageInfo_ClickTypeEnum.Size(m)
}
func (m *ClickTypeEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_ClickTypeEnum.DiscardUnknown(m)
}

var xxx_messageInfo_ClickTypeEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ClickTypeEnum)(nil), "google.ads.googleads.v1.enums.ClickTypeEnum")
	proto.RegisterEnum("google.ads.googleads.v1.enums.ClickTypeEnum_ClickType", ClickTypeEnum_ClickType_name, ClickTypeEnum_ClickType_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/enums/click_type.proto", fileDescriptor_click_type_094021787dbc0905)
}

var fileDescriptor_click_type_094021787dbc0905 = []byte{
	// 930 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x95, 0xdb, 0x6e, 0xdb, 0x36,
	0x18, 0xc7, 0xd7, 0x74, 0x6d, 0x17, 0xe6, 0xc4, 0x32, 0x69, 0xd2, 0x1c, 0x9c, 0xd4, 0xce, 0x92,
	0x6e, 0xdd, 0x26, 0xcf, 0xeb, 0x0e, 0x80, 0x77, 0x45, 0x49, 0x9f, 0x65, 0xc2, 0x14, 0x29, 0x90,
	0xb4, 0x9d, 0x0e, 0x01, 0x08, 0x2f, 0x36, 0x8c, 0x60, 0x89, 0x6d, 0xd4, 0x69, 0x81, 0x3e, 0xc0,
	0x5e, 0x64, 0x97, 0x7b, 0x94, 0xbd, 0xc6, 0xee, 0x76, 0xbb, 0x17, 0x18, 0x28, 0x59, 0x8e, 0x93,
	0xba, 0xcd, 0x6e, 0x0c, 0x9a, 0xff, 0x1f, 0x3f, 0x7d, 0x27, 0xf2, 0x43, 0x5e, 0x7f, 0x38, 0xec,
	0x5f, 0xf4, 0xca, 0x9d, 0xee, 0xb8, 0x9c, 0x2d, 0xdd, 0xea, 0x6d, 0xa5, 0xdc, 0x1b, 0xbc, 0xb9,
	0x1c, 0x97, 0xcf, 0x2e, 0xce, 0xcf, 0x7e, 0xb3, 0x57, 0xef, 0x46, 0x3d, 0x6f, 0xf4, 0x7a, 0x78,
	0x35, 0x24, 0x85, 0x0c, 0xf2, 0x3a, 0xdd, 0xb1, 0x37, 0xe5, 0xbd, 0xb7, 0x15, 0x2f, 0xe5, 0x77,
	0xf6, 0x72, 0x73, 0xa3, 0xf3, 0x72, 0x67, 0x30, 0x18, 0x5e, 0x75, 0xae, 0xce, 0x87, 0x83, 0x71,
	0x76, 0xb8, 0xf4, 0xfb, 0x32, 0x5a, 0x09, 0x9c, 0x45, 0xf3, 0x6e, 0xd4, 0x83, 0xc1, 0x9b, 0xcb,
	0xd2, 0xbf, 0x4b, 0x68, 0x71, 0xba, 0x43, 0xd6, 0xd0, 0x52, 0x53, 0xe8, 0x04, 0x02, 0x56, 0x63,
	0x10, 0xe2, 0x4f, 0xc8, 0x12, 0x7a, 0xd4, 0x14, 0x0d, 0x21, 0xdb, 0x02, 0xdf, 0x23, 0x18, 0x2d,
	0xd3, 0x24, 0xb1, 0x21, 0x40, 0xc2, 0x99, 0x68, 0xe0, 0x05, 0xc7, 0xfb, 0x0a, 0x68, 0x18, 0xa8,
	0x66, 0xec, 0x6b, 0x7c, 0x9f, 0x10, 0xb4, 0xea, 0x2b, 0x49, 0x43, 0x9f, 0x8a, 0xd0, 0x26, 0x9c,
	0x0a, 0xfc, 0x29, 0x79, 0x8c, 0x56, 0x02, 0xca, 0xb9, 0x35, 0x8a, 0x06, 0x0d, 0x26, 0x22, 0xfc,
	0x80, 0x2c, 0xa2, 0x07, 0x6e, 0x4b, 0xe3, 0x87, 0x64, 0x07, 0x6d, 0x06, 0x9c, 0x05, 0x0d, 0x2b,
	0x85, 0x05, 0x11, 0xd1, 0x08, 0x62, 0x10, 0xc6, 0xd2, 0x10, 0x3f, 0x72, 0xd6, 0x22, 0x30, 0x36,
	0x64, 0x0a, 0x02, 0xc3, 0xa4, 0xd0, 0xf8, 0x33, 0xb2, 0x89, 0x08, 0x97, 0x01, 0x75, 0x7f, 0x2d,
	0x9c, 0x24, 0x54, 0x68, 0x26, 0x05, 0x5e, 0x24, 0x4f, 0xd1, 0xc6, 0x74, 0xbf, 0x26, 0x55, 0x4c,
	0x8d, 0x75, 0x9f, 0xc0, 0x88, 0xec, 0xa3, 0x9d, 0xdb, 0xca, 0x8c, 0xc5, 0x25, 0xb2, 0x8d, 0x9e,
	0xdc, 0xd6, 0x59, 0x4c, 0x23, 0xc0, 0xcb, 0xe4, 0x19, 0xda, 0xbb, 0x2d, 0x71, 0x2a, 0x42, 0x26,
	0x22, 0x9b, 0x38, 0x62, 0x85, 0x6c, 0xa1, 0xf5, 0xdb, 0x44, 0x4c, 0x13, 0xbc, 0x3a, 0xef, 0xab,
	0xda, 0x48, 0x05, 0x96, 0x89, 0x9a, 0xc4, 0x6b, 0xf3, 0xfc, 0x35, 0x70, 0x62, 0x30, 0x76, 0x4a,
	0x2c, 0x7d, 0xc6, 0xc1, 0xde, 0x4c, 0xdb, 0x63, 0x57, 0x00, 0x59, 0xab, 0x81, 0xb2, 0x89, 0x62,
	0xc2, 0x68, 0x4c, 0x5c, 0x22, 0xa5, 0xa9, 0x83, 0xc2, 0xeb, 0x64, 0x0f, 0x3d, 0x4d, 0x94, 0x0c,
	0x9b, 0x81, 0xb1, 0x70, 0x62, 0x20, 0xcd, 0x8b, 0x4d, 0x53, 0xab, 0xf1, 0x06, 0x29, 0xa0, 0xed,
	0x5c, 0xe5, 0x4c, 0x1b, 0x17, 0x01, 0x0d, 0x73, 0xf9, 0x09, 0x59, 0x41, 0x8b, 0x9a, 0x19, 0x70,
	0x65, 0xd5, 0x78, 0xd3, 0x95, 0x2c, 0x73, 0x36, 0x75, 0x51, 0x2a, 0xbc, 0x45, 0x56, 0x11, 0x6a,
	0x2a, 0x9e, 0x9f, 0xd8, 0x76, 0x75, 0x6b, 0xb1, 0x10, 0xa4, 0x75, 0x2d, 0x91, 0xc1, 0x13, 0x6d,
	0x87, 0x1c, 0xa0, 0xdd, 0x4c, 0xcb, 0x02, 0x90, 0x96, 0xa6, 0xc9, 0xce, 0x81, 0x5d, 0x72, 0x84,
	0x8a, 0x39, 0xa0, 0xc2, 0x5c, 0xad, 0x03, 0x0d, 0x39, 0x13, 0x53, 0x3b, 0x7b, 0x2e, 0x13, 0x19,
	0x06, 0x22, 0xb4, 0x01, 0x4d, 0x72, 0xa5, 0x70, 0xad, 0xb4, 0xc1, 0x77, 0x8e, 0xe7, 0xca, 0x3e,
	0xd9, 0x40, 0xb8, 0xc5, 0x74, 0x93, 0x72, 0x7b, 0x1d, 0xd0, 0x81, 0x0b, 0xa8, 0xcd, 0x14, 0x70,
	0xd0, 0x3a, 0x6b, 0xcb, 0x67, 0xb3, 0xf9, 0x9a, 0xc9, 0x88, 0x0b, 0x98, 0xe3, 0x22, 0xf9, 0x12,
	0x1d, 0xcd, 0x51, 0xe3, 0x26, 0x37, 0x2c, 0xa8, 0x53, 0x21, 0x80, 0x4f, 0xd0, 0x12, 0x79, 0x81,
	0x8e, 0xef, 0x42, 0xa5, 0x70, 0x81, 0xe1, 0x43, 0xd7, 0x15, 0xef, 0xb3, 0xda, 0x06, 0xb2, 0x99,
	0x48, 0x81, 0x3f, 0x27, 0x87, 0xe8, 0x60, 0x8e, 0x2d, 0xa3, 0xa8, 0xd0, 0x34, 0x30, 0xd4, 0xe7,
	0x80, 0x8f, 0xc8, 0x2e, 0xda, 0xca, 0x21, 0x1a, 0xda, 0x1b, 0x57, 0xf2, 0xd8, 0x15, 0x5a, 0xd7,
	0x65, 0x3b, 0xa0, 0x1a, 0xd2, 0x0a, 0x53, 0x03, 0x91, 0x54, 0xaf, 0x6c, 0x2a, 0x3f, 0x77, 0x99,
	0x9f, 0x95, 0xd3, 0x18, 0xb2, 0xf2, 0xd5, 0x94, 0x14, 0x26, 0xc3, 0xbe, 0x70, 0x7e, 0xcc, 0x62,
	0x99, 0xff, 0xf6, 0xda, 0x35, 0xd1, 0xc0, 0x2f, 0x48, 0x09, 0xed, 0xbf, 0x6f, 0xeb, 0x06, 0xf3,
	0x95, 0xbb, 0x1f, 0x89, 0x92, 0xb1, 0x9c, 0xdc, 0xd7, 0x49, 0x5f, 0xe2, 0xaf, 0x49, 0x11, 0x15,
	0x74, 0x9b, 0x25, 0xe0, 0x62, 0xb2, 0x11, 0xe5, 0x1c, 0xd4, 0x2b, 0x67, 0x25, 0xef, 0x02, 0xfc,
	0x8d, 0x6b, 0xa3, 0xb9, 0x48, 0xba, 0xa9, 0xb1, 0xf7, 0x41, 0x1b, 0x1a, 0xc0, 0xc6, 0x52, 0x01,
	0x2e, 0xa7, 0xf1, 0xce, 0x45, 0x26, 0xcd, 0x61, 0xa5, 0x00, 0xfc, 0xed, 0xdd, 0x98, 0x69, 0x4b,
	0x5c, 0x21, 0xcf, 0xd1, 0xe1, 0x1d, 0x58, 0x5d, 0x01, 0xe0, 0xef, 0xc8, 0x31, 0x2a, 0x7d, 0x1c,
	0xac, 0xc9, 0xa6, 0xc2, 0x2f, 0xff, 0x07, 0xc7, 0x5a, 0x80, 0xbf, 0x77, 0x0f, 0x6d, 0x5d, 0x1a,
	0x70, 0xe9, 0x65, 0x01, 0xe0, 0x1f, 0xc8, 0x3a, 0x5a, 0x4b, 0x97, 0x33, 0x39, 0xfd, 0xd1, 0xb9,
	0x97, 0x51, 0xbe, 0x94, 0xe9, 0x83, 0x1a, 0x49, 0x19, 0x71, 0xb0, 0x4a, 0xca, 0xd8, 0x6a, 0xe0,
	0xd9, 0x9b, 0x87, 0x7f, 0xf2, 0xff, 0xbe, 0x87, 0x8a, 0x67, 0xc3, 0x4b, 0xef, 0xa3, 0xb3, 0xc4,
	0x5f, 0x9d, 0x0e, 0x86, 0xc4, 0x4d, 0x8f, 0xe4, 0xde, 0x2f, 0xfe, 0xe4, 0x40, 0x7f, 0x78, 0xd1,
	0x19, 0xf4, 0xbd, 0xe1, 0xeb, 0x7e, 0xb9, 0xdf, 0x1b, 0xa4, 0xb3, 0x25, 0x1f, 0x5e, 0xa3, 0xf3,
	0xf1, 0x07, 0x66, 0xd9, 0xcf, 0xe9, 0xef, 0x1f, 0x0b, 0xf7, 0x23, 0x4a, 0xff, 0x5c, 0x28, 0x44,
	0x99, 0x29, 0xda, 0x1d, 0x7b, 0xd9, 0xd2, 0xad, 0x5a, 0x15, 0xcf, 0x8d, 0xa5, 0xf1, 0x5f, 0xb9,
	0x7e, 0x4a, 0xbb, 0xe3, 0xd3, 0xa9, 0x7e, 0xda, 0xaa, 0x9c, 0xa6, 0xfa, 0x3f, 0x0b, 0xc5, 0x6c,
	0xb3, 0x5a, 0xa5, 0xdd, 0x71, 0xb5, 0x3a, 0x25, 0xaa, 0xd5, 0x56, 0xa5, 0x5a, 0x4d, 0x99, 0x5f,
	0x1f, 0xa6, 0x8e, 0xbd, 0xfc, 0x2f, 0x00, 0x00, 0xff, 0xff, 0x6a, 0x7d, 0x69, 0x8a, 0x63, 0x07,
	0x00, 0x00,
}
