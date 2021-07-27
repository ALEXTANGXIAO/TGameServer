package Gameproto

import (
	fmt "fmt"
	math "math"

	proto "github.com/gogo/protobuf/proto"
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

type RequestCode int32

const (
	RequestCode_RequestNone RequestCode = 0
	RequestCode_User        RequestCode = 1
	RequestCode_Room        RequestCode = 2
	RequestCode_Game        RequestCode = 3
	RequestCode_Heart       RequestCode = 4
)

var RequestCode_name = map[int32]string{
	0: "RequestNone",
	1: "User",
	2: "Room",
	3: "Game",
	4: "Heart",
}

var RequestCode_value = map[string]int32{
	"RequestNone": 0,
	"User":        1,
	"Room":        2,
	"Game":        3,
	"Heart":       4,
}

func (x RequestCode) String() string {
	return proto.EnumName(RequestCode_name, int32(x))
}

func (RequestCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8b6e03bb4c8f5ddc, []int{0}
}

type ActionCode int32

const (
	ActionCode_ActionNone      ActionCode = 0
	ActionCode_Register        ActionCode = 1
	ActionCode_Login           ActionCode = 2
	ActionCode_CreateRoom      ActionCode = 3
	ActionCode_FindRoom        ActionCode = 4
	ActionCode_PlayerList      ActionCode = 5
	ActionCode_JoinRoom        ActionCode = 6
	ActionCode_Exit            ActionCode = 7
	ActionCode_Chat            ActionCode = 8
	ActionCode_StartGame       ActionCode = 9
	ActionCode_Starting        ActionCode = 10
	ActionCode_UpdateState     ActionCode = 11
	ActionCode_ExitGame        ActionCode = 12
	ActionCode_UpCharacterList ActionCode = 13
	ActionCode_UpPos           ActionCode = 14
	ActionCode_Fire            ActionCode = 15
	ActionCode_HeartBeat       ActionCode = 16
	ActionCode_AddCharacter    ActionCode = 17
	ActionCode_RemoveCharacter ActionCode = 18
)

var ActionCode_name = map[int32]string{
	0:  "ActionNone",
	1:  "Register",
	2:  "Login",
	3:  "CreateRoom",
	4:  "FindRoom",
	5:  "PlayerList",
	6:  "JoinRoom",
	7:  "Exit",
	8:  "Chat",
	9:  "StartGame",
	10: "Starting",
	11: "UpdateState",
	12: "ExitGame",
	13: "UpCharacterList",
	14: "UpPos",
	15: "Fire",
	16: "HeartBeat",
	17: "AddCharacter",
	18: "RemoveCharacter",
}

var ActionCode_value = map[string]int32{
	"ActionNone":      0,
	"Register":        1,
	"Login":           2,
	"CreateRoom":      3,
	"FindRoom":        4,
	"PlayerList":      5,
	"JoinRoom":        6,
	"Exit":            7,
	"Chat":            8,
	"StartGame":       9,
	"Starting":        10,
	"UpdateState":     11,
	"ExitGame":        12,
	"UpCharacterList": 13,
	"UpPos":           14,
	"Fire":            15,
	"HeartBeat":       16,
	"AddCharacter":    17,
	"RemoveCharacter": 18,
}

func (x ActionCode) String() string {
	return proto.EnumName(ActionCode_name, int32(x))
}

func (ActionCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8b6e03bb4c8f5ddc, []int{1}
}

type ReturnCode int32

const (
	ReturnCode_ReturnNone ReturnCode = 0
	ReturnCode_Success    ReturnCode = 1
	ReturnCode_Fail       ReturnCode = 2
	ReturnCode_NoneRoom   ReturnCode = 3
)

var ReturnCode_name = map[int32]string{
	0: "ReturnNone",
	1: "Success",
	2: "Fail",
	3: "NoneRoom",
}

var ReturnCode_value = map[string]int32{
	"ReturnNone": 0,
	"Success":    1,
	"Fail":       2,
	"NoneRoom":   3,
}

func (x ReturnCode) String() string {
	return proto.EnumName(ReturnCode_name, int32(x))
}

func (ReturnCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8b6e03bb4c8f5ddc, []int{2}
}

type MainPack struct {
	Requestcode          RequestCode   `protobuf:"varint,1,opt,name=requestcode,proto3,enum=Gameproto.RequestCode" json:"requestcode,omitempty"`
	Actioncode           ActionCode    `protobuf:"varint,2,opt,name=actioncode,proto3,enum=Gameproto.ActionCode" json:"actioncode,omitempty"`
	Returncode           ReturnCode    `protobuf:"varint,3,opt,name=returncode,proto3,enum=Gameproto.ReturnCode" json:"returncode,omitempty"`
	LoginPack            *LoginPack    `protobuf:"bytes,4,opt,name=loginPack,proto3" json:"loginPack,omitempty"`
	Str                  string        `protobuf:"bytes,5,opt,name=str,proto3" json:"str,omitempty"`
	Roompack             []*RoomPack   `protobuf:"bytes,6,rep,name=roompack,proto3" json:"roompack,omitempty"`
	Playerpack           []*PlayerPack `protobuf:"bytes,7,rep,name=playerpack,proto3" json:"playerpack,omitempty"`
	User                 string        `protobuf:"bytes,8,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *MainPack) Reset()         { *m = MainPack{} }
func (m *MainPack) String() string { return proto.CompactTextString(m) }
func (*MainPack) ProtoMessage()    {}
func (*MainPack) Descriptor() ([]byte, []int) {
	return fileDescriptor_8b6e03bb4c8f5ddc, []int{0}
}
func (m *MainPack) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MainPack.Unmarshal(m, b)
}
func (m *MainPack) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MainPack.Marshal(b, m, deterministic)
}
func (m *MainPack) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MainPack.Merge(m, src)
}
func (m *MainPack) XXX_Size() int {
	return xxx_messageInfo_MainPack.Size(m)
}
func (m *MainPack) XXX_DiscardUnknown() {
	xxx_messageInfo_MainPack.DiscardUnknown(m)
}

var xxx_messageInfo_MainPack proto.InternalMessageInfo

func (m *MainPack) GetRequestcode() RequestCode {
	if m != nil {
		return m.Requestcode
	}
	return RequestCode_RequestNone
}

func (m *MainPack) GetActioncode() ActionCode {
	if m != nil {
		return m.Actioncode
	}
	return ActionCode_ActionNone
}

func (m *MainPack) GetReturncode() ReturnCode {
	if m != nil {
		return m.Returncode
	}
	return ReturnCode_ReturnNone
}

func (m *MainPack) GetLoginPack() *LoginPack {
	if m != nil {
		return m.LoginPack
	}
	return nil
}

func (m *MainPack) GetStr() string {
	if m != nil {
		return m.Str
	}
	return ""
}

func (m *MainPack) GetRoompack() []*RoomPack {
	if m != nil {
		return m.Roompack
	}
	return nil
}

func (m *MainPack) GetPlayerpack() []*PlayerPack {
	if m != nil {
		return m.Playerpack
	}
	return nil
}

func (m *MainPack) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

type LoginPack struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginPack) Reset()         { *m = LoginPack{} }
func (m *LoginPack) String() string { return proto.CompactTextString(m) }
func (*LoginPack) ProtoMessage()    {}
func (*LoginPack) Descriptor() ([]byte, []int) {
	return fileDescriptor_8b6e03bb4c8f5ddc, []int{1}
}
func (m *LoginPack) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginPack.Unmarshal(m, b)
}
func (m *LoginPack) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginPack.Marshal(b, m, deterministic)
}
func (m *LoginPack) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginPack.Merge(m, src)
}
func (m *LoginPack) XXX_Size() int {
	return xxx_messageInfo_LoginPack.Size(m)
}
func (m *LoginPack) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginPack.DiscardUnknown(m)
}

var xxx_messageInfo_LoginPack proto.InternalMessageInfo

func (m *LoginPack) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *LoginPack) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type RoomPack struct {
	Roomname             string   `protobuf:"bytes,1,opt,name=roomname,proto3" json:"roomname,omitempty"`
	Maxnum               int32    `protobuf:"varint,2,opt,name=maxnum,proto3" json:"maxnum,omitempty"`
	Curnum               int32    `protobuf:"varint,3,opt,name=curnum,proto3" json:"curnum,omitempty"`
	State                int32    `protobuf:"varint,4,opt,name=state,proto3" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RoomPack) Reset()         { *m = RoomPack{} }
func (m *RoomPack) String() string { return proto.CompactTextString(m) }
func (*RoomPack) ProtoMessage()    {}
func (*RoomPack) Descriptor() ([]byte, []int) {
	return fileDescriptor_8b6e03bb4c8f5ddc, []int{2}
}
func (m *RoomPack) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoomPack.Unmarshal(m, b)
}
func (m *RoomPack) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoomPack.Marshal(b, m, deterministic)
}
func (m *RoomPack) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoomPack.Merge(m, src)
}
func (m *RoomPack) XXX_Size() int {
	return xxx_messageInfo_RoomPack.Size(m)
}
func (m *RoomPack) XXX_DiscardUnknown() {
	xxx_messageInfo_RoomPack.DiscardUnknown(m)
}

var xxx_messageInfo_RoomPack proto.InternalMessageInfo

func (m *RoomPack) GetRoomname() string {
	if m != nil {
		return m.Roomname
	}
	return ""
}

func (m *RoomPack) GetMaxnum() int32 {
	if m != nil {
		return m.Maxnum
	}
	return 0
}

func (m *RoomPack) GetCurnum() int32 {
	if m != nil {
		return m.Curnum
	}
	return 0
}

func (m *RoomPack) GetState() int32 {
	if m != nil {
		return m.State
	}
	return 0
}

type PlayerPack struct {
	Playername           string   `protobuf:"bytes,1,opt,name=playername,proto3" json:"playername,omitempty"`
	PlayerID             string   `protobuf:"bytes,2,opt,name=playerID,proto3" json:"playerID,omitempty"`
	Hp                   int32    `protobuf:"varint,3,opt,name=hp,proto3" json:"hp,omitempty"`
	PosPack              *PosPack `protobuf:"bytes,4,opt,name=posPack,proto3" json:"posPack,omitempty"`
	PlayerJob            int32    `protobuf:"varint,5,opt,name=playerJob,proto3" json:"playerJob,omitempty"`
	PlayerAppearance     string   `protobuf:"bytes,6,opt,name=playerAppearance,proto3" json:"playerAppearance,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PlayerPack) Reset()         { *m = PlayerPack{} }
func (m *PlayerPack) String() string { return proto.CompactTextString(m) }
func (*PlayerPack) ProtoMessage()    {}
func (*PlayerPack) Descriptor() ([]byte, []int) {
	return fileDescriptor_8b6e03bb4c8f5ddc, []int{3}
}
func (m *PlayerPack) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlayerPack.Unmarshal(m, b)
}
func (m *PlayerPack) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlayerPack.Marshal(b, m, deterministic)
}
func (m *PlayerPack) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlayerPack.Merge(m, src)
}
func (m *PlayerPack) XXX_Size() int {
	return xxx_messageInfo_PlayerPack.Size(m)
}
func (m *PlayerPack) XXX_DiscardUnknown() {
	xxx_messageInfo_PlayerPack.DiscardUnknown(m)
}

var xxx_messageInfo_PlayerPack proto.InternalMessageInfo

func (m *PlayerPack) GetPlayername() string {
	if m != nil {
		return m.Playername
	}
	return ""
}

func (m *PlayerPack) GetPlayerID() string {
	if m != nil {
		return m.PlayerID
	}
	return ""
}

func (m *PlayerPack) GetHp() int32 {
	if m != nil {
		return m.Hp
	}
	return 0
}

func (m *PlayerPack) GetPosPack() *PosPack {
	if m != nil {
		return m.PosPack
	}
	return nil
}

func (m *PlayerPack) GetPlayerJob() int32 {
	if m != nil {
		return m.PlayerJob
	}
	return 0
}

func (m *PlayerPack) GetPlayerAppearance() string {
	if m != nil {
		return m.PlayerAppearance
	}
	return ""
}

type PosPack struct {
	PosX                 float32  `protobuf:"fixed32,1,opt,name=PosX,proto3" json:"PosX,omitempty"`
	PosY                 float32  `protobuf:"fixed32,2,opt,name=PosY,proto3" json:"PosY,omitempty"`
	PosZ                 float32  `protobuf:"fixed32,3,opt,name=PosZ,proto3" json:"PosZ,omitempty"`
	RotaX                float32  `protobuf:"fixed32,4,opt,name=RotaX,proto3" json:"RotaX,omitempty"`
	RotaY                float32  `protobuf:"fixed32,5,opt,name=RotaY,proto3" json:"RotaY,omitempty"`
	RotaZ                float32  `protobuf:"fixed32,6,opt,name=RotaZ,proto3" json:"RotaZ,omitempty"`
	GunRotZ              float32  `protobuf:"fixed32,7,opt,name=GunRotZ,proto3" json:"GunRotZ,omitempty"`
	Animation            int32    `protobuf:"varint,8,opt,name=Animation,proto3" json:"Animation,omitempty"`
	Dirt                 int32    `protobuf:"varint,9,opt,name=Dirt,proto3" json:"Dirt,omitempty"`
	MoveX                float32  `protobuf:"fixed32,10,opt,name=MoveX,proto3" json:"MoveX,omitempty"`
	MoveY                float32  `protobuf:"fixed32,11,opt,name=MoveY,proto3" json:"MoveY,omitempty"`
	MoveZ                float32  `protobuf:"fixed32,12,opt,name=MoveZ,proto3" json:"MoveZ,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PosPack) Reset()         { *m = PosPack{} }
func (m *PosPack) String() string { return proto.CompactTextString(m) }
func (*PosPack) ProtoMessage()    {}
func (*PosPack) Descriptor() ([]byte, []int) {
	return fileDescriptor_8b6e03bb4c8f5ddc, []int{4}
}
func (m *PosPack) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PosPack.Unmarshal(m, b)
}
func (m *PosPack) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PosPack.Marshal(b, m, deterministic)
}
func (m *PosPack) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PosPack.Merge(m, src)
}
func (m *PosPack) XXX_Size() int {
	return xxx_messageInfo_PosPack.Size(m)
}
func (m *PosPack) XXX_DiscardUnknown() {
	xxx_messageInfo_PosPack.DiscardUnknown(m)
}

var xxx_messageInfo_PosPack proto.InternalMessageInfo

func (m *PosPack) GetPosX() float32 {
	if m != nil {
		return m.PosX
	}
	return 0
}

func (m *PosPack) GetPosY() float32 {
	if m != nil {
		return m.PosY
	}
	return 0
}

func (m *PosPack) GetPosZ() float32 {
	if m != nil {
		return m.PosZ
	}
	return 0
}

func (m *PosPack) GetRotaX() float32 {
	if m != nil {
		return m.RotaX
	}
	return 0
}

func (m *PosPack) GetRotaY() float32 {
	if m != nil {
		return m.RotaY
	}
	return 0
}

func (m *PosPack) GetRotaZ() float32 {
	if m != nil {
		return m.RotaZ
	}
	return 0
}

func (m *PosPack) GetGunRotZ() float32 {
	if m != nil {
		return m.GunRotZ
	}
	return 0
}

func (m *PosPack) GetAnimation() int32 {
	if m != nil {
		return m.Animation
	}
	return 0
}

func (m *PosPack) GetDirt() int32 {
	if m != nil {
		return m.Dirt
	}
	return 0
}

func (m *PosPack) GetMoveX() float32 {
	if m != nil {
		return m.MoveX
	}
	return 0
}

func (m *PosPack) GetMoveY() float32 {
	if m != nil {
		return m.MoveY
	}
	return 0
}

func (m *PosPack) GetMoveZ() float32 {
	if m != nil {
		return m.MoveZ
	}
	return 0
}

func init() {
	proto.RegisterEnum("Gameproto.RequestCode", RequestCode_name, RequestCode_value)
	proto.RegisterEnum("Gameproto.ActionCode", ActionCode_name, ActionCode_value)
	proto.RegisterEnum("Gameproto.ReturnCode", ReturnCode_name, ReturnCode_value)
	proto.RegisterType((*MainPack)(nil), "Gameproto.MainPack")
	proto.RegisterType((*LoginPack)(nil), "Gameproto.LoginPack")
	proto.RegisterType((*RoomPack)(nil), "Gameproto.RoomPack")
	proto.RegisterType((*PlayerPack)(nil), "Gameproto.PlayerPack")
	proto.RegisterType((*PosPack)(nil), "Gameproto.PosPack")
}

func init() { proto.RegisterFile("Goproto/SocketGameProtocol.proto", fileDescriptor_8b6e03bb4c8f5ddc) }

var fileDescriptor_8b6e03bb4c8f5ddc = []byte{
	// 751 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x54, 0xdd, 0x6e, 0xea, 0x46,
	0x10, 0x2e, 0x06, 0x83, 0x3d, 0x70, 0xc8, 0x76, 0xcf, 0xe9, 0x91, 0x55, 0x55, 0x15, 0xe2, 0x0a,
	0xa1, 0x2a, 0x47, 0xa2, 0xaa, 0xd4, 0x5b, 0xca, 0x69, 0x68, 0xa3, 0xa4, 0x42, 0x8b, 0x90, 0x02,
	0x77, 0x1b, 0xb3, 0x0a, 0x56, 0xb0, 0xd7, 0x5d, 0x2f, 0x69, 0xfa, 0x20, 0x7d, 0x80, 0xbe, 0x51,
	0x1f, 0xa6, 0x0f, 0x70, 0x34, 0xb3, 0xfe, 0x93, 0x72, 0x37, 0xdf, 0x37, 0xf3, 0xcd, 0x7c, 0x3b,
	0xde, 0x35, 0x4c, 0xd6, 0x3a, 0x37, 0xda, 0xea, 0x4f, 0x5b, 0x1d, 0x3f, 0x2b, 0xbb, 0x96, 0xa9,
	0xda, 0x20, 0x8e, 0xf5, 0xf9, 0x9a, 0x12, 0x3c, 0x44, 0x8e, 0xc2, 0xe9, 0xff, 0x1e, 0x04, 0xf7,
	0x32, 0xc9, 0x36, 0x32, 0x7e, 0xe6, 0x3f, 0xc3, 0xd0, 0xa8, 0x3f, 0x2f, 0xaa, 0xb0, 0xb1, 0x3e,
	0xaa, 0xa8, 0x33, 0xe9, 0xcc, 0xc6, 0x8b, 0x8f, 0xd7, 0x75, 0xf5, 0xb5, 0x70, 0xd9, 0x95, 0x3e,
	0x2a, 0xd1, 0x2e, 0xe5, 0x3f, 0x01, 0xc8, 0xd8, 0x26, 0x3a, 0x23, 0xa1, 0x47, 0xc2, 0x6f, 0x5a,
	0xc2, 0x25, 0x25, 0x49, 0xd7, 0x2a, 0x44, 0x99, 0x51, 0xf6, 0x62, 0x9c, 0xac, 0xfb, 0x46, 0x26,
	0x28, 0xe9, 0x64, 0x4d, 0x21, 0x5f, 0x40, 0x78, 0xd6, 0x4f, 0xce, 0x74, 0xd4, 0x9b, 0x74, 0x66,
	0xc3, 0xc5, 0x87, 0x96, 0xea, 0xae, 0xca, 0x89, 0xa6, 0x8c, 0x33, 0xe8, 0x16, 0xd6, 0x44, 0xfe,
	0xa4, 0x33, 0x0b, 0x05, 0x86, 0xfc, 0x13, 0x04, 0x46, 0xeb, 0x34, 0xc7, 0x26, 0xfd, 0x49, 0x77,
	0x36, 0x5c, 0xbc, 0x6f, 0x8f, 0xd6, 0x3a, 0xa5, 0x1e, 0x75, 0x11, 0xba, 0xcd, 0xcf, 0xf2, 0x6f,
	0x65, 0x48, 0x32, 0x20, 0x49, 0xdb, 0xed, 0x86, 0x92, 0x24, 0x6a, 0x15, 0x72, 0x0e, 0xbd, 0x4b,
	0xa1, 0x4c, 0x14, 0xd0, 0x68, 0x8a, 0xa7, 0x2b, 0x08, 0x6b, 0x97, 0xfc, 0x5b, 0x08, 0x90, 0xcc,
	0x64, 0xea, 0x76, 0x1e, 0x8a, 0x1a, 0x63, 0x2e, 0x97, 0x45, 0xf1, 0x97, 0x36, 0x47, 0x5a, 0x6b,
	0x28, 0x6a, 0x3c, 0x3d, 0x43, 0x50, 0xb9, 0xc4, 0x3a, 0xf4, 0xd9, 0xee, 0x51, 0x61, 0xfe, 0x11,
	0xfa, 0xa9, 0x7c, 0xcd, 0x2e, 0x29, 0x75, 0xf0, 0x45, 0x89, 0x90, 0x8f, 0x2f, 0x06, 0xf9, 0xae,
	0xe3, 0x1d, 0xe2, 0x1f, 0xc0, 0x2f, 0xac, 0xb4, 0x8a, 0x56, 0xeb, 0x0b, 0x07, 0xa6, 0xff, 0x75,
	0x00, 0x9a, 0x13, 0xf2, 0xef, 0xab, 0x65, 0xb4, 0x46, 0xb6, 0x18, 0x32, 0x4e, 0xe8, 0xf7, 0xcf,
	0xb5, 0xf1, 0x12, 0xf3, 0x31, 0x78, 0xa7, 0xbc, 0x1c, 0xea, 0x9d, 0x72, 0xfe, 0x03, 0x0c, 0x72,
	0x5d, 0xb4, 0xbe, 0x26, 0x6f, 0x6f, 0xd5, 0x65, 0x44, 0x55, 0xc2, 0xbf, 0x83, 0xd0, 0x75, 0xba,
	0xd5, 0x8f, 0xf4, 0x3d, 0x7d, 0xd1, 0x10, 0x7c, 0x0e, 0xcc, 0x81, 0x65, 0x9e, 0x2b, 0x69, 0x64,
	0x16, 0xab, 0xa8, 0x4f, 0xf3, 0xdf, 0xf0, 0xd3, 0x7f, 0x3c, 0x18, 0x94, 0xed, 0xf1, 0x2b, 0x6d,
	0x74, 0xf1, 0x40, 0x27, 0xf1, 0x04, 0xc5, 0x25, 0xb7, 0x27, 0xff, 0x8e, 0xdb, 0x97, 0xdc, 0x81,
	0xdc, 0x3b, 0xee, 0x80, 0x0b, 0x13, 0xda, 0xca, 0x07, 0x72, 0xef, 0x09, 0x07, 0x2a, 0x76, 0x4f,
	0x1e, 0x4b, 0x76, 0x5f, 0xb1, 0x07, 0x32, 0x55, 0xb2, 0x07, 0x1e, 0xc1, 0x60, 0x7d, 0xc9, 0x84,
	0xb6, 0x87, 0x68, 0x40, 0x7c, 0x05, 0xf1, 0xb4, 0xcb, 0x2c, 0x49, 0x25, 0xbe, 0x19, 0xba, 0x42,
	0xbe, 0x68, 0x08, 0x74, 0xf3, 0x39, 0x31, 0x36, 0x0a, 0x29, 0x41, 0x31, 0x4e, 0xb8, 0xd7, 0x2f,
	0xea, 0x21, 0x02, 0x37, 0x81, 0x40, 0xc5, 0xee, 0xa3, 0x61, 0xc3, 0xee, 0x2b, 0xf6, 0x10, 0x8d,
	0x1a, 0xf6, 0x30, 0x5f, 0xc3, 0xb0, 0xf5, 0xd2, 0xf9, 0x55, 0x0d, 0xff, 0xd0, 0x99, 0x62, 0x5f,
	0xf1, 0x00, 0x7a, 0xbb, 0x42, 0x19, 0xd6, 0xc1, 0x08, 0xaf, 0x20, 0xf3, 0x30, 0xc2, 0x6f, 0xc6,
	0xba, 0x3c, 0x04, 0xff, 0x37, 0x25, 0x8d, 0x65, 0xbd, 0xf9, 0xbf, 0x1e, 0x40, 0xf3, 0xf4, 0xf9,
	0xb8, 0x42, 0x65, 0x9f, 0x11, 0x04, 0x42, 0x3d, 0x25, 0x85, 0xa5, 0x5e, 0x21, 0xf8, 0xf4, 0x26,
	0x98, 0x87, 0x85, 0x2b, 0xa3, 0xa4, 0x55, 0xd4, 0xbc, 0x8b, 0x85, 0x37, 0x49, 0x76, 0x24, 0xd4,
	0xc3, 0xac, 0xbb, 0x88, 0x77, 0x49, 0x61, 0x99, 0x8f, 0xd9, 0x5b, 0x9d, 0x64, 0x94, 0xed, 0xa3,
	0x91, 0x5f, 0x5f, 0x13, 0xcb, 0x06, 0x18, 0xad, 0x4e, 0xd2, 0xb2, 0x80, 0xbf, 0x83, 0x70, 0x6b,
	0xa5, 0xa1, 0x7f, 0x21, 0x0b, 0x51, 0x40, 0x30, 0xc9, 0x9e, 0x18, 0xe0, 0xf1, 0x76, 0xf9, 0x51,
	0x5a, 0xb5, 0xc5, 0x7b, 0xce, 0x86, 0x98, 0xc6, 0x0e, 0x54, 0x3c, 0xe2, 0xef, 0xe1, 0x6a, 0x97,
	0xaf, 0x4e, 0xd2, 0xc8, 0xd8, 0x96, 0x23, 0xdf, 0xa1, 0xd7, 0x5d, 0xbe, 0xd1, 0x05, 0x1b, 0xe3,
	0x94, 0x9b, 0xc4, 0x28, 0x76, 0x85, 0x53, 0xe8, 0xe0, 0xbf, 0x28, 0x69, 0x19, 0xe3, 0x0c, 0x46,
	0xcb, 0xe3, 0xb1, 0x56, 0xb2, 0xaf, 0xb1, 0x95, 0x50, 0xa9, 0x7e, 0x51, 0x0d, 0xc9, 0xe7, 0x4b,
	0x80, 0xe6, 0x37, 0x87, 0x67, 0x73, 0xa8, 0x5c, 0xd1, 0x10, 0x06, 0xdb, 0x4b, 0x1c, 0xab, 0xa2,
	0x70, 0xdb, 0xbe, 0x91, 0xc9, 0x99, 0x79, 0x68, 0x11, 0x0b, 0xdc, 0x7a, 0x1e, 0xfb, 0xf4, 0x52,
	0x7e, 0xfc, 0x12, 0x00, 0x00, 0xff, 0xff, 0x79, 0xb3, 0x96, 0x34, 0xfa, 0x05, 0x00, 0x00,
}
