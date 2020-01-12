// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/jakewright/home-automation/service.scene/proto/scene.proto

package sceneproto

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "github.com/jakewright/home-automation/tools/protoc-gen-jrpc/proto"
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

type Scene struct {
	Id                   uint32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string    `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Actions              []*Action `protobuf:"bytes,3,rep,name=actions,proto3" json:"actions,omitempty"`
	CreatedAt            string    `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            string    `protobuf:"bytes,5,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	DeletedAt            string    `protobuf:"bytes,6,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Scene) Reset()         { *m = Scene{} }
func (m *Scene) String() string { return proto.CompactTextString(m) }
func (*Scene) ProtoMessage()    {}
func (*Scene) Descriptor() ([]byte, []int) {
	return fileDescriptor_e138e95ef801d2a9, []int{0}
}

func (m *Scene) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Scene.Unmarshal(m, b)
}
func (m *Scene) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Scene.Marshal(b, m, deterministic)
}
func (m *Scene) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Scene.Merge(m, src)
}
func (m *Scene) XXX_Size() int {
	return xxx_messageInfo_Scene.Size(m)
}
func (m *Scene) XXX_DiscardUnknown() {
	xxx_messageInfo_Scene.DiscardUnknown(m)
}

var xxx_messageInfo_Scene proto.InternalMessageInfo

func (m *Scene) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Scene) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Scene) GetActions() []*Action {
	if m != nil {
		return m.Actions
	}
	return nil
}

func (m *Scene) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *Scene) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

func (m *Scene) GetDeletedAt() string {
	if m != nil {
		return m.DeletedAt
	}
	return ""
}

type Action struct {
	Id                   uint32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Stage                int32    `protobuf:"varint,2,opt,name=stage,proto3" json:"stage,omitempty"`
	Sequence             int32    `protobuf:"varint,3,opt,name=sequence,proto3" json:"sequence,omitempty"`
	Func                 string   `protobuf:"bytes,4,opt,name=func,proto3" json:"func,omitempty"`
	ControllerName       string   `protobuf:"bytes,5,opt,name=controller_name,json=controllerName,proto3" json:"controller_name,omitempty"`
	Command              string   `protobuf:"bytes,6,opt,name=command,proto3" json:"command,omitempty"`
	Property             string   `protobuf:"bytes,7,opt,name=property,proto3" json:"property,omitempty"`
	PropertyValue        string   `protobuf:"bytes,8,opt,name=property_value,json=propertyValue,proto3" json:"property_value,omitempty"`
	CreatedAt            string   `protobuf:"bytes,9,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            string   `protobuf:"bytes,10,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	DeletedAt            string   `protobuf:"bytes,11,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Action) Reset()         { *m = Action{} }
func (m *Action) String() string { return proto.CompactTextString(m) }
func (*Action) ProtoMessage()    {}
func (*Action) Descriptor() ([]byte, []int) {
	return fileDescriptor_e138e95ef801d2a9, []int{1}
}

func (m *Action) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Action.Unmarshal(m, b)
}
func (m *Action) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Action.Marshal(b, m, deterministic)
}
func (m *Action) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Action.Merge(m, src)
}
func (m *Action) XXX_Size() int {
	return xxx_messageInfo_Action.Size(m)
}
func (m *Action) XXX_DiscardUnknown() {
	xxx_messageInfo_Action.DiscardUnknown(m)
}

var xxx_messageInfo_Action proto.InternalMessageInfo

func (m *Action) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Action) GetStage() int32 {
	if m != nil {
		return m.Stage
	}
	return 0
}

func (m *Action) GetSequence() int32 {
	if m != nil {
		return m.Sequence
	}
	return 0
}

func (m *Action) GetFunc() string {
	if m != nil {
		return m.Func
	}
	return ""
}

func (m *Action) GetControllerName() string {
	if m != nil {
		return m.ControllerName
	}
	return ""
}

func (m *Action) GetCommand() string {
	if m != nil {
		return m.Command
	}
	return ""
}

func (m *Action) GetProperty() string {
	if m != nil {
		return m.Property
	}
	return ""
}

func (m *Action) GetPropertyValue() string {
	if m != nil {
		return m.PropertyValue
	}
	return ""
}

func (m *Action) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *Action) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

func (m *Action) GetDeletedAt() string {
	if m != nil {
		return m.DeletedAt
	}
	return ""
}

type CreateSceneRequest struct {
	Name                 string                       `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Actions              []*CreateSceneRequest_Action `protobuf:"bytes,2,rep,name=actions,proto3" json:"actions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *CreateSceneRequest) Reset()         { *m = CreateSceneRequest{} }
func (m *CreateSceneRequest) String() string { return proto.CompactTextString(m) }
func (*CreateSceneRequest) ProtoMessage()    {}
func (*CreateSceneRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e138e95ef801d2a9, []int{2}
}

func (m *CreateSceneRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateSceneRequest.Unmarshal(m, b)
}
func (m *CreateSceneRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateSceneRequest.Marshal(b, m, deterministic)
}
func (m *CreateSceneRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateSceneRequest.Merge(m, src)
}
func (m *CreateSceneRequest) XXX_Size() int {
	return xxx_messageInfo_CreateSceneRequest.Size(m)
}
func (m *CreateSceneRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateSceneRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateSceneRequest proto.InternalMessageInfo

func (m *CreateSceneRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateSceneRequest) GetActions() []*CreateSceneRequest_Action {
	if m != nil {
		return m.Actions
	}
	return nil
}

type CreateSceneRequest_Action struct {
	Stage                int32    `protobuf:"varint,1,opt,name=stage,proto3" json:"stage,omitempty"`
	Sequence             int32    `protobuf:"varint,2,opt,name=sequence,proto3" json:"sequence,omitempty"`
	Func                 string   `protobuf:"bytes,3,opt,name=func,proto3" json:"func,omitempty"`
	ControllerName       string   `protobuf:"bytes,4,opt,name=controller_name,json=controllerName,proto3" json:"controller_name,omitempty"`
	Command              string   `protobuf:"bytes,5,opt,name=command,proto3" json:"command,omitempty"`
	Property             string   `protobuf:"bytes,6,opt,name=property,proto3" json:"property,omitempty"`
	PropertyValue        string   `protobuf:"bytes,7,opt,name=property_value,json=propertyValue,proto3" json:"property_value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateSceneRequest_Action) Reset()         { *m = CreateSceneRequest_Action{} }
func (m *CreateSceneRequest_Action) String() string { return proto.CompactTextString(m) }
func (*CreateSceneRequest_Action) ProtoMessage()    {}
func (*CreateSceneRequest_Action) Descriptor() ([]byte, []int) {
	return fileDescriptor_e138e95ef801d2a9, []int{2, 0}
}

func (m *CreateSceneRequest_Action) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateSceneRequest_Action.Unmarshal(m, b)
}
func (m *CreateSceneRequest_Action) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateSceneRequest_Action.Marshal(b, m, deterministic)
}
func (m *CreateSceneRequest_Action) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateSceneRequest_Action.Merge(m, src)
}
func (m *CreateSceneRequest_Action) XXX_Size() int {
	return xxx_messageInfo_CreateSceneRequest_Action.Size(m)
}
func (m *CreateSceneRequest_Action) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateSceneRequest_Action.DiscardUnknown(m)
}

var xxx_messageInfo_CreateSceneRequest_Action proto.InternalMessageInfo

func (m *CreateSceneRequest_Action) GetStage() int32 {
	if m != nil {
		return m.Stage
	}
	return 0
}

func (m *CreateSceneRequest_Action) GetSequence() int32 {
	if m != nil {
		return m.Sequence
	}
	return 0
}

func (m *CreateSceneRequest_Action) GetFunc() string {
	if m != nil {
		return m.Func
	}
	return ""
}

func (m *CreateSceneRequest_Action) GetControllerName() string {
	if m != nil {
		return m.ControllerName
	}
	return ""
}

func (m *CreateSceneRequest_Action) GetCommand() string {
	if m != nil {
		return m.Command
	}
	return ""
}

func (m *CreateSceneRequest_Action) GetProperty() string {
	if m != nil {
		return m.Property
	}
	return ""
}

func (m *CreateSceneRequest_Action) GetPropertyValue() string {
	if m != nil {
		return m.PropertyValue
	}
	return ""
}

type CreateSceneResponse struct {
	Scene                *Scene   `protobuf:"bytes,1,opt,name=scene,proto3" json:"scene,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateSceneResponse) Reset()         { *m = CreateSceneResponse{} }
func (m *CreateSceneResponse) String() string { return proto.CompactTextString(m) }
func (*CreateSceneResponse) ProtoMessage()    {}
func (*CreateSceneResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e138e95ef801d2a9, []int{3}
}

func (m *CreateSceneResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateSceneResponse.Unmarshal(m, b)
}
func (m *CreateSceneResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateSceneResponse.Marshal(b, m, deterministic)
}
func (m *CreateSceneResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateSceneResponse.Merge(m, src)
}
func (m *CreateSceneResponse) XXX_Size() int {
	return xxx_messageInfo_CreateSceneResponse.Size(m)
}
func (m *CreateSceneResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateSceneResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateSceneResponse proto.InternalMessageInfo

func (m *CreateSceneResponse) GetScene() *Scene {
	if m != nil {
		return m.Scene
	}
	return nil
}

type ReadSceneRequest struct {
	SceneId              uint32   `protobuf:"varint,1,opt,name=scene_id,json=sceneId,proto3" json:"scene_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadSceneRequest) Reset()         { *m = ReadSceneRequest{} }
func (m *ReadSceneRequest) String() string { return proto.CompactTextString(m) }
func (*ReadSceneRequest) ProtoMessage()    {}
func (*ReadSceneRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e138e95ef801d2a9, []int{4}
}

func (m *ReadSceneRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadSceneRequest.Unmarshal(m, b)
}
func (m *ReadSceneRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadSceneRequest.Marshal(b, m, deterministic)
}
func (m *ReadSceneRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadSceneRequest.Merge(m, src)
}
func (m *ReadSceneRequest) XXX_Size() int {
	return xxx_messageInfo_ReadSceneRequest.Size(m)
}
func (m *ReadSceneRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadSceneRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReadSceneRequest proto.InternalMessageInfo

func (m *ReadSceneRequest) GetSceneId() uint32 {
	if m != nil {
		return m.SceneId
	}
	return 0
}

type ReadSceneResponse struct {
	Scene                *Scene   `protobuf:"bytes,1,opt,name=scene,proto3" json:"scene,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadSceneResponse) Reset()         { *m = ReadSceneResponse{} }
func (m *ReadSceneResponse) String() string { return proto.CompactTextString(m) }
func (*ReadSceneResponse) ProtoMessage()    {}
func (*ReadSceneResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e138e95ef801d2a9, []int{5}
}

func (m *ReadSceneResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadSceneResponse.Unmarshal(m, b)
}
func (m *ReadSceneResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadSceneResponse.Marshal(b, m, deterministic)
}
func (m *ReadSceneResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadSceneResponse.Merge(m, src)
}
func (m *ReadSceneResponse) XXX_Size() int {
	return xxx_messageInfo_ReadSceneResponse.Size(m)
}
func (m *ReadSceneResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadSceneResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ReadSceneResponse proto.InternalMessageInfo

func (m *ReadSceneResponse) GetScene() *Scene {
	if m != nil {
		return m.Scene
	}
	return nil
}

type ListScenesRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListScenesRequest) Reset()         { *m = ListScenesRequest{} }
func (m *ListScenesRequest) String() string { return proto.CompactTextString(m) }
func (*ListScenesRequest) ProtoMessage()    {}
func (*ListScenesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e138e95ef801d2a9, []int{6}
}

func (m *ListScenesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListScenesRequest.Unmarshal(m, b)
}
func (m *ListScenesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListScenesRequest.Marshal(b, m, deterministic)
}
func (m *ListScenesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListScenesRequest.Merge(m, src)
}
func (m *ListScenesRequest) XXX_Size() int {
	return xxx_messageInfo_ListScenesRequest.Size(m)
}
func (m *ListScenesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListScenesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListScenesRequest proto.InternalMessageInfo

type ListScenesResponse struct {
	Scenes               []*Scene `protobuf:"bytes,1,rep,name=scenes,proto3" json:"scenes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListScenesResponse) Reset()         { *m = ListScenesResponse{} }
func (m *ListScenesResponse) String() string { return proto.CompactTextString(m) }
func (*ListScenesResponse) ProtoMessage()    {}
func (*ListScenesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e138e95ef801d2a9, []int{7}
}

func (m *ListScenesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListScenesResponse.Unmarshal(m, b)
}
func (m *ListScenesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListScenesResponse.Marshal(b, m, deterministic)
}
func (m *ListScenesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListScenesResponse.Merge(m, src)
}
func (m *ListScenesResponse) XXX_Size() int {
	return xxx_messageInfo_ListScenesResponse.Size(m)
}
func (m *ListScenesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListScenesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListScenesResponse proto.InternalMessageInfo

func (m *ListScenesResponse) GetScenes() []*Scene {
	if m != nil {
		return m.Scenes
	}
	return nil
}

type DeleteSceneRequest struct {
	SceneId              int32    `protobuf:"varint,1,opt,name=scene_id,json=sceneId,proto3" json:"scene_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteSceneRequest) Reset()         { *m = DeleteSceneRequest{} }
func (m *DeleteSceneRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteSceneRequest) ProtoMessage()    {}
func (*DeleteSceneRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e138e95ef801d2a9, []int{8}
}

func (m *DeleteSceneRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteSceneRequest.Unmarshal(m, b)
}
func (m *DeleteSceneRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteSceneRequest.Marshal(b, m, deterministic)
}
func (m *DeleteSceneRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteSceneRequest.Merge(m, src)
}
func (m *DeleteSceneRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteSceneRequest.Size(m)
}
func (m *DeleteSceneRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteSceneRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteSceneRequest proto.InternalMessageInfo

func (m *DeleteSceneRequest) GetSceneId() int32 {
	if m != nil {
		return m.SceneId
	}
	return 0
}

type DeleteSceneResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteSceneResponse) Reset()         { *m = DeleteSceneResponse{} }
func (m *DeleteSceneResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteSceneResponse) ProtoMessage()    {}
func (*DeleteSceneResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e138e95ef801d2a9, []int{9}
}

func (m *DeleteSceneResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteSceneResponse.Unmarshal(m, b)
}
func (m *DeleteSceneResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteSceneResponse.Marshal(b, m, deterministic)
}
func (m *DeleteSceneResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteSceneResponse.Merge(m, src)
}
func (m *DeleteSceneResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteSceneResponse.Size(m)
}
func (m *DeleteSceneResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteSceneResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteSceneResponse proto.InternalMessageInfo

type SetSceneRequest struct {
	SceneId              uint32   `protobuf:"varint,1,opt,name=scene_id,json=sceneId,proto3" json:"scene_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetSceneRequest) Reset()         { *m = SetSceneRequest{} }
func (m *SetSceneRequest) String() string { return proto.CompactTextString(m) }
func (*SetSceneRequest) ProtoMessage()    {}
func (*SetSceneRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e138e95ef801d2a9, []int{10}
}

func (m *SetSceneRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetSceneRequest.Unmarshal(m, b)
}
func (m *SetSceneRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetSceneRequest.Marshal(b, m, deterministic)
}
func (m *SetSceneRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetSceneRequest.Merge(m, src)
}
func (m *SetSceneRequest) XXX_Size() int {
	return xxx_messageInfo_SetSceneRequest.Size(m)
}
func (m *SetSceneRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetSceneRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetSceneRequest proto.InternalMessageInfo

func (m *SetSceneRequest) GetSceneId() uint32 {
	if m != nil {
		return m.SceneId
	}
	return 0
}

type SetSceneResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetSceneResponse) Reset()         { *m = SetSceneResponse{} }
func (m *SetSceneResponse) String() string { return proto.CompactTextString(m) }
func (*SetSceneResponse) ProtoMessage()    {}
func (*SetSceneResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e138e95ef801d2a9, []int{11}
}

func (m *SetSceneResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetSceneResponse.Unmarshal(m, b)
}
func (m *SetSceneResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetSceneResponse.Marshal(b, m, deterministic)
}
func (m *SetSceneResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetSceneResponse.Merge(m, src)
}
func (m *SetSceneResponse) XXX_Size() int {
	return xxx_messageInfo_SetSceneResponse.Size(m)
}
func (m *SetSceneResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SetSceneResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SetSceneResponse proto.InternalMessageInfo

type SetSceneEvent struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetSceneEvent) Reset()         { *m = SetSceneEvent{} }
func (m *SetSceneEvent) String() string { return proto.CompactTextString(m) }
func (*SetSceneEvent) ProtoMessage()    {}
func (*SetSceneEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_e138e95ef801d2a9, []int{12}
}

func (m *SetSceneEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetSceneEvent.Unmarshal(m, b)
}
func (m *SetSceneEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetSceneEvent.Marshal(b, m, deterministic)
}
func (m *SetSceneEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetSceneEvent.Merge(m, src)
}
func (m *SetSceneEvent) XXX_Size() int {
	return xxx_messageInfo_SetSceneEvent.Size(m)
}
func (m *SetSceneEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_SetSceneEvent.DiscardUnknown(m)
}

var xxx_messageInfo_SetSceneEvent proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Scene)(nil), "sceneproto.Scene")
	proto.RegisterType((*Action)(nil), "sceneproto.Action")
	proto.RegisterType((*CreateSceneRequest)(nil), "sceneproto.CreateSceneRequest")
	proto.RegisterType((*CreateSceneRequest_Action)(nil), "sceneproto.CreateSceneRequest.Action")
	proto.RegisterType((*CreateSceneResponse)(nil), "sceneproto.CreateSceneResponse")
	proto.RegisterType((*ReadSceneRequest)(nil), "sceneproto.ReadSceneRequest")
	proto.RegisterType((*ReadSceneResponse)(nil), "sceneproto.ReadSceneResponse")
	proto.RegisterType((*ListScenesRequest)(nil), "sceneproto.ListScenesRequest")
	proto.RegisterType((*ListScenesResponse)(nil), "sceneproto.ListScenesResponse")
	proto.RegisterType((*DeleteSceneRequest)(nil), "sceneproto.DeleteSceneRequest")
	proto.RegisterType((*DeleteSceneResponse)(nil), "sceneproto.DeleteSceneResponse")
	proto.RegisterType((*SetSceneRequest)(nil), "sceneproto.SetSceneRequest")
	proto.RegisterType((*SetSceneResponse)(nil), "sceneproto.SetSceneResponse")
	proto.RegisterType((*SetSceneEvent)(nil), "sceneproto.SetSceneEvent")
}

func init() {
	proto.RegisterFile("github.com/jakewright/home-automation/service.scene/proto/scene.proto", fileDescriptor_e138e95ef801d2a9)
}

var fileDescriptor_e138e95ef801d2a9 = []byte{
	// 758 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x55, 0xcd, 0x6e, 0xd3, 0x4a,
	0x14, 0x96, 0x9d, 0xd8, 0x49, 0x4e, 0x9a, 0xa6, 0x9d, 0xdc, 0xde, 0x3b, 0xf5, 0xed, 0xed, 0x8d,
	0x8c, 0xaa, 0x16, 0xa9, 0x89, 0xa5, 0xb2, 0x03, 0x44, 0xd5, 0xd2, 0x80, 0x90, 0x2a, 0x40, 0x4e,
	0xc5, 0x82, 0x4d, 0xe4, 0xda, 0x43, 0xe2, 0x92, 0xd8, 0xc1, 0x33, 0x29, 0x62, 0xdb, 0x17, 0xe1,
	0x1d, 0x90, 0xd8, 0x57, 0x5d, 0xb1, 0x87, 0x07, 0x41, 0xea, 0x0b, 0x20, 0xcf, 0xf8, 0xb7, 0x49,
	0xda, 0xc0, 0x2a, 0x33, 0xe7, 0x9c, 0xf9, 0xce, 0x39, 0xdf, 0xf7, 0x29, 0x86, 0x4e, 0xdf, 0x65,
	0x83, 0xc9, 0x69, 0xdb, 0xf6, 0x47, 0xc6, 0x99, 0xf5, 0x9e, 0x7c, 0x0c, 0xdc, 0xfe, 0x80, 0x19,
	0x03, 0x7f, 0x44, 0x5a, 0xd6, 0x84, 0xf9, 0x23, 0x8b, 0xb9, 0xbe, 0x67, 0x50, 0x12, 0x9c, 0xbb,
	0x36, 0x69, 0x53, 0x9b, 0x78, 0xc4, 0x18, 0x07, 0x3e, 0xf3, 0x0d, 0x7e, 0x6e, 0xf3, 0x33, 0x02,
	0x7e, 0xe1, 0x67, 0xed, 0x78, 0x31, 0x48, 0xe6, 0xfb, 0x43, 0x2a, 0xa0, 0xec, 0x56, 0x9f, 0x78,
	0xad, 0xb3, 0x60, 0x6c, 0x47, 0xd0, 0xe1, 0x51, 0x20, 0xeb, 0xdf, 0x24, 0x50, 0xba, 0x21, 0x38,
	0x5a, 0x06, 0xd9, 0x75, 0xb0, 0xd4, 0x94, 0x76, 0x6a, 0xa6, 0xec, 0x3a, 0x08, 0x41, 0xd1, 0xb3,
	0x46, 0x04, 0xcb, 0x4d, 0x69, 0xa7, 0x62, 0xf2, 0x33, 0xda, 0x85, 0x92, 0x65, 0x87, 0xf0, 0x14,
	0x17, 0x9a, 0x85, 0x9d, 0xea, 0x1e, 0x6a, 0xa7, 0x93, 0xb5, 0x0f, 0x78, 0xca, 0x8c, 0x4b, 0xd0,
	0x3d, 0x00, 0x3b, 0x20, 0x16, 0x23, 0x4e, 0xcf, 0x62, 0xb8, 0x18, 0xe2, 0x1c, 0x16, 0x2f, 0x7f,
	0xae, 0x4b, 0x66, 0x25, 0x8a, 0x1f, 0xb0, 0xb0, 0x68, 0x32, 0x76, 0xe2, 0x22, 0x25, 0x5b, 0x14,
	0xc5, 0x45, 0x91, 0x43, 0x86, 0x24, 0x2a, 0x52, 0xb3, 0x45, 0x51, 0xfc, 0x80, 0xe9, 0xdf, 0x65,
	0x50, 0xc5, 0x08, 0x53, 0xbb, 0xfc, 0x05, 0x0a, 0x65, 0x56, 0x5f, 0x2c, 0xa3, 0x98, 0xe2, 0x82,
	0x34, 0x28, 0x53, 0xf2, 0x61, 0x42, 0x3c, 0x9b, 0xe0, 0x02, 0x4f, 0x24, 0xf7, 0x70, 0xfb, 0x77,
	0x13, 0xcf, 0x16, 0x53, 0x9b, 0xfc, 0x8c, 0xb6, 0xa1, 0x6e, 0xfb, 0x1e, 0x0b, 0xfc, 0xe1, 0x90,
	0x04, 0x3d, 0x4e, 0x0e, 0x9f, 0xd7, 0x5c, 0x4e, 0xc3, 0x2f, 0x43, 0x9a, 0x30, 0x94, 0x6c, 0x7f,
	0x34, 0xb2, 0x3c, 0x47, 0xcc, 0x6a, 0xc6, 0xd7, 0xb0, 0xe5, 0x38, 0xf0, 0xc7, 0x24, 0x60, 0x9f,
	0x70, 0x89, 0xa7, 0x92, 0x3b, 0xda, 0x82, 0xe5, 0xf8, 0xdc, 0x3b, 0xb7, 0x86, 0x13, 0x82, 0xcb,
	0xbc, 0xa2, 0x16, 0x47, 0xdf, 0x84, 0xc1, 0x1b, 0xac, 0x56, 0x16, 0x61, 0x15, 0x16, 0x61, 0xb5,
	0x3a, 0x9b, 0xd5, 0x4b, 0x19, 0xd0, 0x53, 0x8e, 0xcb, 0x6d, 0x62, 0x86, 0x04, 0x51, 0x96, 0xb8,
	0x43, 0xca, 0xb8, 0x63, 0x3f, 0x75, 0x87, 0xcc, 0xdd, 0xb1, 0x95, 0x75, 0xc7, 0x34, 0xc8, 0x4d,
	0xc3, 0x68, 0x3f, 0xa4, 0x44, 0xc1, 0x44, 0x31, 0x69, 0x9e, 0x62, 0xf2, 0x1c, 0xc5, 0x0a, 0xb7,
	0x2b, 0x56, 0xbc, 0x4b, 0x31, 0x65, 0xbe, 0x62, 0xea, 0x9d, 0x8a, 0x95, 0x66, 0x28, 0xa6, 0x3f,
	0x81, 0x46, 0x6e, 0x79, 0x3a, 0xf6, 0x3d, 0x4a, 0xd0, 0x36, 0x28, 0x9c, 0x1e, 0xbe, 0x62, 0x75,
	0x6f, 0x35, 0x4b, 0x96, 0xa8, 0x14, 0x79, 0xbd, 0x05, 0x2b, 0x26, 0xb1, 0x9c, 0x1c, 0xff, 0xeb,
	0x50, 0xe6, 0xc9, 0x5e, 0xe2, 0xf3, 0x12, 0xbf, 0xbf, 0x70, 0xf4, 0xc7, 0xb0, 0x9a, 0x29, 0xff,
	0xdd, 0x66, 0x0d, 0x58, 0x3d, 0x76, 0x29, 0xe3, 0x31, 0x1a, 0x75, 0xd3, 0xf7, 0x01, 0x65, 0x83,
	0x11, 0xe6, 0x7d, 0x50, 0xf9, 0x1b, 0x8a, 0x25, 0x2e, 0xf7, 0x0c, 0xd0, 0xa8, 0x40, 0x37, 0x00,
	0x1d, 0x71, 0x4b, 0xdd, 0xba, 0x84, 0x92, 0x2e, 0xb1, 0x06, 0x8d, 0xdc, 0x03, 0xd1, 0x52, 0xdf,
	0x85, 0x7a, 0x97, 0xb0, 0x45, 0x99, 0x40, 0xb0, 0x92, 0x56, 0x47, 0x08, 0x4d, 0xa8, 0xc5, 0xb1,
	0xce, 0x39, 0xf1, 0xd8, 0xc3, 0xfa, 0xc5, 0x97, 0xf5, 0x2a, 0x54, 0x28, 0x61, 0x2d, 0xfe, 0x6e,
	0xef, 0x73, 0x11, 0x96, 0x78, 0xbe, 0x2b, 0xfe, 0x95, 0x51, 0x1f, 0xaa, 0x19, 0xfd, 0xd0, 0xe6,
	0xed, 0xae, 0xd6, 0xfe, 0x9f, 0x9b, 0x8f, 0x46, 0xf8, 0xe7, 0xea, 0x1a, 0xab, 0x50, 0x7c, 0xfd,
	0xaa, 0x7b, 0x72, 0x75, 0x8d, 0x2b, 0xa8, 0x24, 0xfe, 0xeb, 0x29, 0xea, 0x41, 0x25, 0x51, 0x0e,
	0x6d, 0x64, 0x61, 0x6e, 0xea, 0xaf, 0xfd, 0x37, 0x27, 0x1b, 0xb5, 0x58, 0xbb, 0xba, 0xc6, 0x0a,
	0x14, 0x9e, 0x77, 0xc2, 0x0e, 0x65, 0xa4, 0x8a, 0x0e, 0xc8, 0x06, 0x48, 0x75, 0x44, 0x39, 0x8c,
	0x29, 0xd1, 0xb5, 0xcd, 0x79, 0xe9, 0xa8, 0xc7, 0xdf, 0xd9, 0x1e, 0x99, 0x2d, 0x06, 0x50, 0xcd,
	0x48, 0x97, 0xa7, 0x6b, 0xda, 0x04, 0x79, 0xba, 0x66, 0x69, 0x8e, 0xc3, 0x15, 0x40, 0x3d, 0xea,
	0x1c, 0x77, 0x4e, 0x3a, 0xb9, 0x75, 0x2c, 0x28, 0xc7, 0x5a, 0xa2, 0x7f, 0x73, 0xe6, 0xcb, 0x7b,
	0x44, 0xdb, 0x98, 0x9d, 0x8c, 0x1a, 0x68, 0x39, 0x3d, 0x96, 0x10, 0x08, 0x78, 0x83, 0x12, 0xa6,
	0x35, 0x2e, 0xbe, 0xe2, 0x3a, 0xd4, 0x72, 0x1f, 0xe8, 0xc3, 0x67, 0x6f, 0x8f, 0xfe, 0xf8, 0xbb,
	0xfe, 0x28, 0x9d, 0xe4, 0x54, 0xe5, 0x3f, 0x0f, 0x7e, 0x05, 0x00, 0x00, 0xff, 0xff, 0xa7, 0x6f,
	0xad, 0x2e, 0x26, 0x08, 0x00, 0x00,
}