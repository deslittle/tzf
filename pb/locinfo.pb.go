// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: pb/locinfo.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CompressMethod int32

const (
	CompressMethod_Unknown  CompressMethod = 0
	CompressMethod_Polyline CompressMethod = 1 // https://developers.google.com/maps/documentation/utilities/polylinealgorithm
)

// Enum value maps for CompressMethod.
var (
	CompressMethod_name = map[int32]string{
		0: "Unknown",
		1: "Polyline",
	}
	CompressMethod_value = map[string]int32{
		"Unknown":  0,
		"Polyline": 1,
	}
)

func (x CompressMethod) Enum() *CompressMethod {
	p := new(CompressMethod)
	*p = x
	return p
}

func (x CompressMethod) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CompressMethod) Descriptor() protoreflect.EnumDescriptor {
	return file_pb_locinfo_proto_enumTypes[0].Descriptor()
}

func (CompressMethod) Type() protoreflect.EnumType {
	return &file_pb_locinfo_proto_enumTypes[0]
}

func (x CompressMethod) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CompressMethod.Descriptor instead.
func (CompressMethod) EnumDescriptor() ([]byte, []int) {
	return file_pb_locinfo_proto_rawDescGZIP(), []int{0}
}

// Basic Point data define.
type Point struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Lng float32 `protobuf:"fixed32,1,opt,name=lng,proto3" json:"lng,omitempty"`
	Lat float32 `protobuf:"fixed32,2,opt,name=lat,proto3" json:"lat,omitempty"`
}

func (x *Point) Reset() {
	*x = Point{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_locinfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Point) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Point) ProtoMessage() {}

func (x *Point) ProtoReflect() protoreflect.Message {
	mi := &file_pb_locinfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Point.ProtoReflect.Descriptor instead.
func (*Point) Descriptor() ([]byte, []int) {
	return file_pb_locinfo_proto_rawDescGZIP(), []int{0}
}

func (x *Point) GetLng() float32 {
	if x != nil {
		return x.Lng
	}
	return 0
}

func (x *Point) GetLat() float32 {
	if x != nil {
		return x.Lat
	}
	return 0
}

// Define a polygon, mostly based on GeoJSON's Polygon define.
//
// Excerpt from RFC-9476 section 'Polygon'
//
//   - A linear ring is a closed LineString with four or more positions.
//
//   - The first and last positions are equivalent, and they MUST contain
//     identical values; their representation SHOULD also be identical.
//
//   - A linear ring is the boundary of a surface or the boundary of a
//     hole in a surface.
//
//   - A linear ring MUST follow the right-hand rule with respect to the
//     area it bounds, i.e., exterior rings are counterclockwise, and
//     holes are clockwise.
//
//     Note: the [GJ2008] specification did not discuss linear ring winding
//     order.  For backwards compatibility, parsers SHOULD NOT reject
//     Polygons that do not follow the right-hand rule.
//
//     Though a linear ring is not explicitly represented as a GeoJSON
//     geometry type, it leads to a canonical formulation of the Polygon
//     geometry type definition as follows:
//
//   - For type "Polygon", the "coordinates" member MUST be an array of
//     linear ring coordinate arrays.
//
//   - For Polygons with more than one of these rings, the first MUST be
//     the exterior ring, and any others MUST be interior rings.  The
//     exterior ring bounds the surface, and the interior rings (if
//     present) bound holes within the surface.
//
// [GJ2008]: https://geojson.org/geojson-spec
type Polygon struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Points []*Point   `protobuf:"bytes,1,rep,name=points,proto3" json:"points,omitempty"` // define the "exterior ring"
	Holes  []*Polygon `protobuf:"bytes,2,rep,name=holes,proto3" json:"holes,omitempty"`   // define the "interior rings" as holes
}

func (x *Polygon) Reset() {
	*x = Polygon{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_locinfo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Polygon) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Polygon) ProtoMessage() {}

func (x *Polygon) ProtoReflect() protoreflect.Message {
	mi := &file_pb_locinfo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Polygon.ProtoReflect.Descriptor instead.
func (*Polygon) Descriptor() ([]byte, []int) {
	return file_pb_locinfo_proto_rawDescGZIP(), []int{1}
}

func (x *Polygon) GetPoints() []*Point {
	if x != nil {
		return x.Points
	}
	return nil
}

func (x *Polygon) GetHoles() []*Polygon {
	if x != nil {
		return x.Holes
	}
	return nil
}

// Location is a locations's all data.
type Location struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Polygons []*Polygon `protobuf:"bytes,1,rep,name=polygons,proto3" json:"polygons,omitempty"`
	Name     string     `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Location) Reset() {
	*x = Location{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_locinfo_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Location) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Location) ProtoMessage() {}

func (x *Location) ProtoReflect() protoreflect.Message {
	mi := &file_pb_locinfo_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Location.ProtoReflect.Descriptor instead.
func (*Location) Descriptor() ([]byte, []int) {
	return file_pb_locinfo_proto_rawDescGZIP(), []int{2}
}

func (x *Location) GetPolygons() []*Polygon {
	if x != nil {
		return x.Polygons
	}
	return nil
}

func (x *Location) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Locations struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Locations []*Location `protobuf:"bytes,1,rep,name=locations,proto3" json:"locations,omitempty"`
	Reduced   bool        `protobuf:"varint,2,opt,name=reduced,proto3" json:"reduced,omitempty"` // Reduced data will toggle neighbor search as plan b
}

func (x *Locations) Reset() {
	*x = Locations{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_locinfo_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Locations) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Locations) ProtoMessage() {}

func (x *Locations) ProtoReflect() protoreflect.Message {
	mi := &file_pb_locinfo_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Locations.ProtoReflect.Descriptor instead.
func (*Locations) Descriptor() ([]byte, []int) {
	return file_pb_locinfo_proto_rawDescGZIP(), []int{3}
}

func (x *Locations) GetLocations() []*Location {
	if x != nil {
		return x.Locations
	}
	return nil
}

func (x *Locations) GetReduced() bool {
	if x != nil {
		return x.Reduced
	}
	return false
}

type CompressedPolygon struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Points []byte               `protobuf:"bytes,1,opt,name=points,proto3" json:"points,omitempty"`
	Holes  []*CompressedPolygon `protobuf:"bytes,2,rep,name=holes,proto3" json:"holes,omitempty"`
}

func (x *CompressedPolygon) Reset() {
	*x = CompressedPolygon{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_locinfo_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompressedPolygon) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompressedPolygon) ProtoMessage() {}

func (x *CompressedPolygon) ProtoReflect() protoreflect.Message {
	mi := &file_pb_locinfo_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompressedPolygon.ProtoReflect.Descriptor instead.
func (*CompressedPolygon) Descriptor() ([]byte, []int) {
	return file_pb_locinfo_proto_rawDescGZIP(), []int{4}
}

func (x *CompressedPolygon) GetPoints() []byte {
	if x != nil {
		return x.Points
	}
	return nil
}

func (x *CompressedPolygon) GetHoles() []*CompressedPolygon {
	if x != nil {
		return x.Holes
	}
	return nil
}

// CompressedLocationsItem designed for binary file as small as possible.
type CompressedLocation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*CompressedPolygon `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	Name string               `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CompressedLocation) Reset() {
	*x = CompressedLocation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_locinfo_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompressedLocation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompressedLocation) ProtoMessage() {}

func (x *CompressedLocation) ProtoReflect() protoreflect.Message {
	mi := &file_pb_locinfo_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompressedLocation.ProtoReflect.Descriptor instead.
func (*CompressedLocation) Descriptor() ([]byte, []int) {
	return file_pb_locinfo_proto_rawDescGZIP(), []int{5}
}

func (x *CompressedLocation) GetData() []*CompressedPolygon {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *CompressedLocation) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type CompressedLocations struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Method    CompressMethod        `protobuf:"varint,1,opt,name=method,proto3,enum=pinpoint.pb.v1.CompressMethod" json:"method,omitempty"`
	Locations []*CompressedLocation `protobuf:"bytes,2,rep,name=locations,proto3" json:"locations,omitempty"`
}

func (x *CompressedLocations) Reset() {
	*x = CompressedLocations{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_locinfo_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompressedLocations) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompressedLocations) ProtoMessage() {}

func (x *CompressedLocations) ProtoReflect() protoreflect.Message {
	mi := &file_pb_locinfo_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompressedLocations.ProtoReflect.Descriptor instead.
func (*CompressedLocations) Descriptor() ([]byte, []int) {
	return file_pb_locinfo_proto_rawDescGZIP(), []int{6}
}

func (x *CompressedLocations) GetMethod() CompressMethod {
	if x != nil {
		return x.Method
	}
	return CompressMethod_Unknown
}

func (x *CompressedLocations) GetLocations() []*CompressedLocation {
	if x != nil {
		return x.Locations
	}
	return nil
}

// PreindexLocation tile item.
//
// The X/Y/Z are OSM style like map tile index values.
type PreindexLocation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	X    int32  `protobuf:"varint,2,opt,name=x,proto3" json:"x,omitempty"`
	Y    int32  `protobuf:"varint,3,opt,name=y,proto3" json:"y,omitempty"`
	Z    int32  `protobuf:"varint,4,opt,name=z,proto3" json:"z,omitempty"`
}

func (x *PreindexLocation) Reset() {
	*x = PreindexLocation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_locinfo_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PreindexLocation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PreindexLocation) ProtoMessage() {}

func (x *PreindexLocation) ProtoReflect() protoreflect.Message {
	mi := &file_pb_locinfo_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PreindexLocation.ProtoReflect.Descriptor instead.
func (*PreindexLocation) Descriptor() ([]byte, []int) {
	return file_pb_locinfo_proto_rawDescGZIP(), []int{7}
}

func (x *PreindexLocation) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PreindexLocation) GetX() int32 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *PreindexLocation) GetY() int32 {
	if x != nil {
		return x.Y
	}
	return 0
}

func (x *PreindexLocation) GetZ() int32 {
	if x != nil {
		return x.Z
	}
	return 0
}

// PreindexLocations is all preindex location's dumps.
type PreindexLocations struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IdxZoom int32               `protobuf:"varint,1,opt,name=idxZoom,proto3" json:"idxZoom,omitempty"` // which zoom value the tiles generated
	AggZoom int32               `protobuf:"varint,2,opt,name=aggZoom,proto3" json:"aggZoom,omitempty"` // which zoom value the tiles merge up with.
	Keys    []*PreindexLocation `protobuf:"bytes,3,rep,name=keys,proto3" json:"keys,omitempty"`
}

func (x *PreindexLocations) Reset() {
	*x = PreindexLocations{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_locinfo_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PreindexLocations) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PreindexLocations) ProtoMessage() {}

func (x *PreindexLocations) ProtoReflect() protoreflect.Message {
	mi := &file_pb_locinfo_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PreindexLocations.ProtoReflect.Descriptor instead.
func (*PreindexLocations) Descriptor() ([]byte, []int) {
	return file_pb_locinfo_proto_rawDescGZIP(), []int{8}
}

func (x *PreindexLocations) GetIdxZoom() int32 {
	if x != nil {
		return x.IdxZoom
	}
	return 0
}

func (x *PreindexLocations) GetAggZoom() int32 {
	if x != nil {
		return x.AggZoom
	}
	return 0
}

func (x *PreindexLocations) GetKeys() []*PreindexLocation {
	if x != nil {
		return x.Keys
	}
	return nil
}

var File_pb_locinfo_proto protoreflect.FileDescriptor

var file_pb_locinfo_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x62, 0x2f, 0x6c, 0x6f, 0x63, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0e, 0x70, 0x69, 0x6e, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x70, 0x62, 0x2e,
	0x76, 0x31, 0x22, 0x2b, 0x0a, 0x05, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6c,
	0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x03, 0x6c, 0x6e, 0x67, 0x12, 0x10, 0x0a,
	0x03, 0x6c, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x03, 0x6c, 0x61, 0x74, 0x22,
	0x67, 0x0a, 0x07, 0x50, 0x6f, 0x6c, 0x79, 0x67, 0x6f, 0x6e, 0x12, 0x2d, 0x0a, 0x06, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x69, 0x6e,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x69, 0x6e,
	0x74, 0x52, 0x06, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x2d, 0x0a, 0x05, 0x68, 0x6f, 0x6c,
	0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70, 0x69, 0x6e, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x2e, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x6c, 0x79, 0x67, 0x6f,
	0x6e, 0x52, 0x05, 0x68, 0x6f, 0x6c, 0x65, 0x73, 0x22, 0x53, 0x0a, 0x08, 0x4c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x33, 0x0a, 0x08, 0x70, 0x6f, 0x6c, 0x79, 0x67, 0x6f, 0x6e, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70, 0x69, 0x6e, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x2e, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x6c, 0x79, 0x67, 0x6f, 0x6e, 0x52,
	0x08, 0x70, 0x6f, 0x6c, 0x79, 0x67, 0x6f, 0x6e, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x5d, 0x0a,
	0x09, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x36, 0x0a, 0x09, 0x6c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e,
	0x70, 0x69, 0x6e, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x4c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x64, 0x75, 0x63, 0x65, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x72, 0x65, 0x64, 0x75, 0x63, 0x65, 0x64, 0x22, 0x64, 0x0a, 0x11,
	0x43, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x65, 0x64, 0x50, 0x6f, 0x6c, 0x79, 0x67, 0x6f,
	0x6e, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x06, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x37, 0x0a, 0x05, 0x68, 0x6f, 0x6c,
	0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x70, 0x69, 0x6e, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x2e, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x72, 0x65,
	0x73, 0x73, 0x65, 0x64, 0x50, 0x6f, 0x6c, 0x79, 0x67, 0x6f, 0x6e, 0x52, 0x05, 0x68, 0x6f, 0x6c,
	0x65, 0x73, 0x22, 0x5f, 0x0a, 0x12, 0x43, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x65, 0x64,
	0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x35, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x70, 0x69, 0x6e, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x2e, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73,
	0x65, 0x64, 0x50, 0x6f, 0x6c, 0x79, 0x67, 0x6f, 0x6e, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x22, 0x8f, 0x01, 0x0a, 0x13, 0x43, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73,
	0x65, 0x64, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x36, 0x0a, 0x06, 0x6d,
	0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x70, 0x69,
	0x6e, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d,
	0x70, 0x72, 0x65, 0x73, 0x73, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x52, 0x06, 0x6d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x12, 0x40, 0x0a, 0x09, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x70, 0x69, 0x6e, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x2e, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73,
	0x65, 0x64, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x6c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x50, 0x0a, 0x10, 0x50, 0x72, 0x65, 0x69, 0x6e, 0x64, 0x65,
	0x78, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x0c, 0x0a,
	0x01, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x78, 0x12, 0x0c, 0x0a, 0x01, 0x79,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x79, 0x12, 0x0c, 0x0a, 0x01, 0x7a, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x7a, 0x22, 0x7d, 0x0a, 0x11, 0x50, 0x72, 0x65, 0x69, 0x6e,
	0x64, 0x65, 0x78, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x18, 0x0a, 0x07,
	0x69, 0x64, 0x78, 0x5a, 0x6f, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x69,
	0x64, 0x78, 0x5a, 0x6f, 0x6f, 0x6d, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x67, 0x67, 0x5a, 0x6f, 0x6f,
	0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x61, 0x67, 0x67, 0x5a, 0x6f, 0x6f, 0x6d,
	0x12, 0x34, 0x0a, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20,
	0x2e, 0x70, 0x69, 0x6e, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x2e,
	0x50, 0x72, 0x65, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x2a, 0x2b, 0x0a, 0x0e, 0x43, 0x6f, 0x6d, 0x70, 0x72, 0x65,
	0x73, 0x73, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x6e, 0x6b, 0x6e,
	0x6f, 0x77, 0x6e, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x50, 0x6f, 0x6c, 0x79, 0x6c, 0x69, 0x6e,
	0x65, 0x10, 0x01, 0x42, 0x25, 0x5a, 0x23, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x64, 0x65, 0x73, 0x6c, 0x69, 0x74, 0x74, 0x6c, 0x65, 0x2f, 0x70, 0x69, 0x6e, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x2f, 0x70, 0x62, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_pb_locinfo_proto_rawDescOnce sync.Once
	file_pb_locinfo_proto_rawDescData = file_pb_locinfo_proto_rawDesc
)

func file_pb_locinfo_proto_rawDescGZIP() []byte {
	file_pb_locinfo_proto_rawDescOnce.Do(func() {
		file_pb_locinfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_locinfo_proto_rawDescData)
	})
	return file_pb_locinfo_proto_rawDescData
}

var file_pb_locinfo_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_pb_locinfo_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_pb_locinfo_proto_goTypes = []interface{}{
	(CompressMethod)(0),         // 0: pinpoint.pb.v1.CompressMethod
	(*Point)(nil),               // 1: pinpoint.pb.v1.Point
	(*Polygon)(nil),             // 2: pinpoint.pb.v1.Polygon
	(*Location)(nil),            // 3: pinpoint.pb.v1.Location
	(*Locations)(nil),           // 4: pinpoint.pb.v1.Locations
	(*CompressedPolygon)(nil),   // 5: pinpoint.pb.v1.CompressedPolygon
	(*CompressedLocation)(nil),  // 6: pinpoint.pb.v1.CompressedLocation
	(*CompressedLocations)(nil), // 7: pinpoint.pb.v1.CompressedLocations
	(*PreindexLocation)(nil),    // 8: pinpoint.pb.v1.PreindexLocation
	(*PreindexLocations)(nil),   // 9: pinpoint.pb.v1.PreindexLocations
}
var file_pb_locinfo_proto_depIdxs = []int32{
	1, // 0: pinpoint.pb.v1.Polygon.points:type_name -> pinpoint.pb.v1.Point
	2, // 1: pinpoint.pb.v1.Polygon.holes:type_name -> pinpoint.pb.v1.Polygon
	2, // 2: pinpoint.pb.v1.Location.polygons:type_name -> pinpoint.pb.v1.Polygon
	3, // 3: pinpoint.pb.v1.Locations.locations:type_name -> pinpoint.pb.v1.Location
	5, // 4: pinpoint.pb.v1.CompressedPolygon.holes:type_name -> pinpoint.pb.v1.CompressedPolygon
	5, // 5: pinpoint.pb.v1.CompressedLocation.data:type_name -> pinpoint.pb.v1.CompressedPolygon
	0, // 6: pinpoint.pb.v1.CompressedLocations.method:type_name -> pinpoint.pb.v1.CompressMethod
	6, // 7: pinpoint.pb.v1.CompressedLocations.locations:type_name -> pinpoint.pb.v1.CompressedLocation
	8, // 8: pinpoint.pb.v1.PreindexLocations.keys:type_name -> pinpoint.pb.v1.PreindexLocation
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_pb_locinfo_proto_init() }
func file_pb_locinfo_proto_init() {
	if File_pb_locinfo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_locinfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Point); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pb_locinfo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Polygon); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pb_locinfo_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Location); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pb_locinfo_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Locations); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pb_locinfo_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompressedPolygon); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pb_locinfo_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompressedLocation); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pb_locinfo_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompressedLocations); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pb_locinfo_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PreindexLocation); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pb_locinfo_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PreindexLocations); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pb_locinfo_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pb_locinfo_proto_goTypes,
		DependencyIndexes: file_pb_locinfo_proto_depIdxs,
		EnumInfos:         file_pb_locinfo_proto_enumTypes,
		MessageInfos:      file_pb_locinfo_proto_msgTypes,
	}.Build()
	File_pb_locinfo_proto = out.File
	file_pb_locinfo_proto_rawDesc = nil
	file_pb_locinfo_proto_goTypes = nil
	file_pb_locinfo_proto_depIdxs = nil
}
