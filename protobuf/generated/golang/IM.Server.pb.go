// Code generated by protoc-gen-go. DO NOT EDIT.
// source: IM.Server.proto

/*
Package IM_Server is a generated protocol buffer package.

It is generated from these files:
	IM.Server.proto

It has these top-level messages:
	IMStopReceivePacket
	IMValidateReq
	IMValidateRsp
	IMGetDeviceTokenReq
	IMGetDeviceTokenRsp
	IMRoleSet
	IMOnlineUserInfo
	IMMsgServInfo
	IMUserStatusUpdate
	IMUserCntUpdate
	IMServerKickUser
	IMServerPCLoginStatusNotify
	IMPushToUserReq
	IMPushToUserRsp
	IMGroupGetShieldReq
	IMGroupGetShieldRsp
	IMFileTransferReq
	IMFileTransferRsp
	IMFileServerIPReq
	IMFileServerIPRsp
*/
package IM_Server

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import IM_BaseDefine "."

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// db_proxy
type IMStopReceivePacket struct {
	// cmd id:   0x0702
	Result           *uint32 `protobuf:"varint,1,req,name=result" json:"result,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *IMStopReceivePacket) Reset()                    { *m = IMStopReceivePacket{} }
func (m *IMStopReceivePacket) String() string            { return proto.CompactTextString(m) }
func (*IMStopReceivePacket) ProtoMessage()               {}
func (*IMStopReceivePacket) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *IMStopReceivePacket) GetResult() uint32 {
	if m != nil && m.Result != nil {
		return *m.Result
	}
	return 0
}

// db_proxy
type IMValidateReq struct {
	// cmd id:	0x0703
	UserName         *string `protobuf:"bytes,1,req,name=user_name,json=userName" json:"user_name,omitempty"`
	Password         *string `protobuf:"bytes,2,req,name=password" json:"password,omitempty"`
	AttachData       []byte  `protobuf:"bytes,20,opt,name=attach_data,json=attachData" json:"attach_data,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *IMValidateReq) Reset()                    { *m = IMValidateReq{} }
func (m *IMValidateReq) String() string            { return proto.CompactTextString(m) }
func (*IMValidateReq) ProtoMessage()               {}
func (*IMValidateReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *IMValidateReq) GetUserName() string {
	if m != nil && m.UserName != nil {
		return *m.UserName
	}
	return ""
}

func (m *IMValidateReq) GetPassword() string {
	if m != nil && m.Password != nil {
		return *m.Password
	}
	return ""
}

func (m *IMValidateReq) GetAttachData() []byte {
	if m != nil {
		return m.AttachData
	}
	return nil
}

// db_proxy
type IMValidateRsp struct {
	// cmd id:	0x0704
	UserName         *string                 `protobuf:"bytes,1,req,name=user_name,json=userName" json:"user_name,omitempty"`
	ResultCode       *uint32                 `protobuf:"varint,2,req,name=result_code,json=resultCode" json:"result_code,omitempty"`
	ResultString     *string                 `protobuf:"bytes,3,opt,name=result_string,json=resultString" json:"result_string,omitempty"`
	UserInfo         *IM_BaseDefine.UserInfo `protobuf:"bytes,4,opt,name=user_info,json=userInfo" json:"user_info,omitempty"`
	AttachData       []byte                  `protobuf:"bytes,20,opt,name=attach_data,json=attachData" json:"attach_data,omitempty"`
	XXX_unrecognized []byte                  `json:"-"`
}

func (m *IMValidateRsp) Reset()                    { *m = IMValidateRsp{} }
func (m *IMValidateRsp) String() string            { return proto.CompactTextString(m) }
func (*IMValidateRsp) ProtoMessage()               {}
func (*IMValidateRsp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *IMValidateRsp) GetUserName() string {
	if m != nil && m.UserName != nil {
		return *m.UserName
	}
	return ""
}

func (m *IMValidateRsp) GetResultCode() uint32 {
	if m != nil && m.ResultCode != nil {
		return *m.ResultCode
	}
	return 0
}

func (m *IMValidateRsp) GetResultString() string {
	if m != nil && m.ResultString != nil {
		return *m.ResultString
	}
	return ""
}

func (m *IMValidateRsp) GetUserInfo() *IM_BaseDefine.UserInfo {
	if m != nil {
		return m.UserInfo
	}
	return nil
}

func (m *IMValidateRsp) GetAttachData() []byte {
	if m != nil {
		return m.AttachData
	}
	return nil
}

// db_proxy
type IMGetDeviceTokenReq struct {
	// cmd id:	0x0705
	UserId           []uint32 `protobuf:"varint,1,rep,name=user_id,json=userId" json:"user_id,omitempty"`
	AttachData       []byte   `protobuf:"bytes,20,opt,name=attach_data,json=attachData" json:"attach_data,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *IMGetDeviceTokenReq) Reset()                    { *m = IMGetDeviceTokenReq{} }
func (m *IMGetDeviceTokenReq) String() string            { return proto.CompactTextString(m) }
func (*IMGetDeviceTokenReq) ProtoMessage()               {}
func (*IMGetDeviceTokenReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *IMGetDeviceTokenReq) GetUserId() []uint32 {
	if m != nil {
		return m.UserId
	}
	return nil
}

func (m *IMGetDeviceTokenReq) GetAttachData() []byte {
	if m != nil {
		return m.AttachData
	}
	return nil
}

// db_proxy
type IMGetDeviceTokenRsp struct {
	// cmd id:	0x0706
	UserTokenInfo    []*IM_BaseDefine.UserTokenInfo `protobuf:"bytes,1,rep,name=user_token_info,json=userTokenInfo" json:"user_token_info,omitempty"`
	AttachData       []byte                         `protobuf:"bytes,20,opt,name=attach_data,json=attachData" json:"attach_data,omitempty"`
	XXX_unrecognized []byte                         `json:"-"`
}

func (m *IMGetDeviceTokenRsp) Reset()                    { *m = IMGetDeviceTokenRsp{} }
func (m *IMGetDeviceTokenRsp) String() string            { return proto.CompactTextString(m) }
func (*IMGetDeviceTokenRsp) ProtoMessage()               {}
func (*IMGetDeviceTokenRsp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *IMGetDeviceTokenRsp) GetUserTokenInfo() []*IM_BaseDefine.UserTokenInfo {
	if m != nil {
		return m.UserTokenInfo
	}
	return nil
}

func (m *IMGetDeviceTokenRsp) GetAttachData() []byte {
	if m != nil {
		return m.AttachData
	}
	return nil
}

type IMRoleSet struct {
	// cmd id:	0x0707
	Master           *uint32 `protobuf:"varint,1,req,name=master" json:"master,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *IMRoleSet) Reset()                    { *m = IMRoleSet{} }
func (m *IMRoleSet) String() string            { return proto.CompactTextString(m) }
func (*IMRoleSet) ProtoMessage()               {}
func (*IMRoleSet) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *IMRoleSet) GetMaster() uint32 {
	if m != nil && m.Master != nil {
		return *m.Master
	}
	return 0
}

type IMOnlineUserInfo struct {
	// cmd id:	0x0708
	UserStatList     []*IM_BaseDefine.ServerUserStat `protobuf:"bytes,1,rep,name=user_stat_list,json=userStatList" json:"user_stat_list,omitempty"`
	XXX_unrecognized []byte                          `json:"-"`
}

func (m *IMOnlineUserInfo) Reset()                    { *m = IMOnlineUserInfo{} }
func (m *IMOnlineUserInfo) String() string            { return proto.CompactTextString(m) }
func (*IMOnlineUserInfo) ProtoMessage()               {}
func (*IMOnlineUserInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *IMOnlineUserInfo) GetUserStatList() []*IM_BaseDefine.ServerUserStat {
	if m != nil {
		return m.UserStatList
	}
	return nil
}

type IMMsgServInfo struct {
	// cmd id:	0x0709
	Ip1              *string `protobuf:"bytes,1,req,name=ip1" json:"ip1,omitempty"`
	Ip2              *string `protobuf:"bytes,2,req,name=ip2" json:"ip2,omitempty"`
	Port             *uint32 `protobuf:"varint,3,req,name=port" json:"port,omitempty"`
	MaxConnCnt       *uint32 `protobuf:"varint,4,req,name=max_conn_cnt,json=maxConnCnt" json:"max_conn_cnt,omitempty"`
	CurConnCnt       *uint32 `protobuf:"varint,5,req,name=cur_conn_cnt,json=curConnCnt" json:"cur_conn_cnt,omitempty"`
	HostName         *string `protobuf:"bytes,6,req,name=host_name,json=hostName" json:"host_name,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *IMMsgServInfo) Reset()                    { *m = IMMsgServInfo{} }
func (m *IMMsgServInfo) String() string            { return proto.CompactTextString(m) }
func (*IMMsgServInfo) ProtoMessage()               {}
func (*IMMsgServInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *IMMsgServInfo) GetIp1() string {
	if m != nil && m.Ip1 != nil {
		return *m.Ip1
	}
	return ""
}

func (m *IMMsgServInfo) GetIp2() string {
	if m != nil && m.Ip2 != nil {
		return *m.Ip2
	}
	return ""
}

func (m *IMMsgServInfo) GetPort() uint32 {
	if m != nil && m.Port != nil {
		return *m.Port
	}
	return 0
}

func (m *IMMsgServInfo) GetMaxConnCnt() uint32 {
	if m != nil && m.MaxConnCnt != nil {
		return *m.MaxConnCnt
	}
	return 0
}

func (m *IMMsgServInfo) GetCurConnCnt() uint32 {
	if m != nil && m.CurConnCnt != nil {
		return *m.CurConnCnt
	}
	return 0
}

func (m *IMMsgServInfo) GetHostName() string {
	if m != nil && m.HostName != nil {
		return *m.HostName
	}
	return ""
}

type IMUserStatusUpdate struct {
	// cmd id:	0x070a
	UserStatus       *uint32                   `protobuf:"varint,1,req,name=user_status,json=userStatus" json:"user_status,omitempty"`
	UserId           *uint32                   `protobuf:"varint,2,req,name=user_id,json=userId" json:"user_id,omitempty"`
	ClientType       *IM_BaseDefine.ClientType `protobuf:"varint,3,req,name=client_type,json=clientType,enum=IM.BaseDefine.ClientType" json:"client_type,omitempty"`
	XXX_unrecognized []byte                    `json:"-"`
}

func (m *IMUserStatusUpdate) Reset()                    { *m = IMUserStatusUpdate{} }
func (m *IMUserStatusUpdate) String() string            { return proto.CompactTextString(m) }
func (*IMUserStatusUpdate) ProtoMessage()               {}
func (*IMUserStatusUpdate) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *IMUserStatusUpdate) GetUserStatus() uint32 {
	if m != nil && m.UserStatus != nil {
		return *m.UserStatus
	}
	return 0
}

func (m *IMUserStatusUpdate) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *IMUserStatusUpdate) GetClientType() IM_BaseDefine.ClientType {
	if m != nil && m.ClientType != nil {
		return *m.ClientType
	}
	return IM_BaseDefine.ClientType_CLIENT_TYPE_WINDOWS
}

type IMUserCntUpdate struct {
	// cmd id:	0x070b
	UserAction       *uint32 `protobuf:"varint,1,req,name=user_action,json=userAction" json:"user_action,omitempty"`
	UserId           *uint32 `protobuf:"varint,2,req,name=user_id,json=userId" json:"user_id,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *IMUserCntUpdate) Reset()                    { *m = IMUserCntUpdate{} }
func (m *IMUserCntUpdate) String() string            { return proto.CompactTextString(m) }
func (*IMUserCntUpdate) ProtoMessage()               {}
func (*IMUserCntUpdate) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *IMUserCntUpdate) GetUserAction() uint32 {
	if m != nil && m.UserAction != nil {
		return *m.UserAction
	}
	return 0
}

func (m *IMUserCntUpdate) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

type IMServerKickUser struct {
	// cmd id:	0x070d
	UserId           *uint32                   `protobuf:"varint,1,req,name=user_id,json=userId" json:"user_id,omitempty"`
	ClientType       *IM_BaseDefine.ClientType `protobuf:"varint,2,req,name=client_type,json=clientType,enum=IM.BaseDefine.ClientType" json:"client_type,omitempty"`
	Reason           *uint32                   `protobuf:"varint,3,req,name=reason" json:"reason,omitempty"`
	XXX_unrecognized []byte                    `json:"-"`
}

func (m *IMServerKickUser) Reset()                    { *m = IMServerKickUser{} }
func (m *IMServerKickUser) String() string            { return proto.CompactTextString(m) }
func (*IMServerKickUser) ProtoMessage()               {}
func (*IMServerKickUser) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *IMServerKickUser) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *IMServerKickUser) GetClientType() IM_BaseDefine.ClientType {
	if m != nil && m.ClientType != nil {
		return *m.ClientType
	}
	return IM_BaseDefine.ClientType_CLIENT_TYPE_WINDOWS
}

func (m *IMServerKickUser) GetReason() uint32 {
	if m != nil && m.Reason != nil {
		return *m.Reason
	}
	return 0
}

type IMServerPCLoginStatusNotify struct {
	// cmd id:	0x070e
	UserId           *uint32 `protobuf:"varint,1,req,name=user_id,json=userId" json:"user_id,omitempty"`
	LoginStatus      *uint32 `protobuf:"varint,2,req,name=login_status,json=loginStatus" json:"login_status,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *IMServerPCLoginStatusNotify) Reset()                    { *m = IMServerPCLoginStatusNotify{} }
func (m *IMServerPCLoginStatusNotify) String() string            { return proto.CompactTextString(m) }
func (*IMServerPCLoginStatusNotify) ProtoMessage()               {}
func (*IMServerPCLoginStatusNotify) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *IMServerPCLoginStatusNotify) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *IMServerPCLoginStatusNotify) GetLoginStatus() uint32 {
	if m != nil && m.LoginStatus != nil {
		return *m.LoginStatus
	}
	return 0
}

type IMPushToUserReq struct {
	// cmd id:	0x070f
	Flash            *string                        `protobuf:"bytes,1,req,name=flash" json:"flash,omitempty"`
	Data             *string                        `protobuf:"bytes,2,req,name=data" json:"data,omitempty"`
	UserTokenList    []*IM_BaseDefine.UserTokenInfo `protobuf:"bytes,3,rep,name=user_token_list,json=userTokenList" json:"user_token_list,omitempty"`
	XXX_unrecognized []byte                         `json:"-"`
}

func (m *IMPushToUserReq) Reset()                    { *m = IMPushToUserReq{} }
func (m *IMPushToUserReq) String() string            { return proto.CompactTextString(m) }
func (*IMPushToUserReq) ProtoMessage()               {}
func (*IMPushToUserReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *IMPushToUserReq) GetFlash() string {
	if m != nil && m.Flash != nil {
		return *m.Flash
	}
	return ""
}

func (m *IMPushToUserReq) GetData() string {
	if m != nil && m.Data != nil {
		return *m.Data
	}
	return ""
}

func (m *IMPushToUserReq) GetUserTokenList() []*IM_BaseDefine.UserTokenInfo {
	if m != nil {
		return m.UserTokenList
	}
	return nil
}

type IMPushToUserRsp struct {
	// cmd id:	0x0710
	PushResultList   []*IM_BaseDefine.PushResult `protobuf:"bytes,1,rep,name=push_result_list,json=pushResultList" json:"push_result_list,omitempty"`
	XXX_unrecognized []byte                      `json:"-"`
}

func (m *IMPushToUserRsp) Reset()                    { *m = IMPushToUserRsp{} }
func (m *IMPushToUserRsp) String() string            { return proto.CompactTextString(m) }
func (*IMPushToUserRsp) ProtoMessage()               {}
func (*IMPushToUserRsp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *IMPushToUserRsp) GetPushResultList() []*IM_BaseDefine.PushResult {
	if m != nil {
		return m.PushResultList
	}
	return nil
}

type IMGroupGetShieldReq struct {
	// cmd id:			0x0711
	GroupId          *uint32  `protobuf:"varint,1,req,name=group_id,json=groupId" json:"group_id,omitempty"`
	UserId           []uint32 `protobuf:"varint,2,rep,name=user_id,json=userId" json:"user_id,omitempty"`
	AttachData       []byte   `protobuf:"bytes,20,opt,name=attach_data,json=attachData" json:"attach_data,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *IMGroupGetShieldReq) Reset()                    { *m = IMGroupGetShieldReq{} }
func (m *IMGroupGetShieldReq) String() string            { return proto.CompactTextString(m) }
func (*IMGroupGetShieldReq) ProtoMessage()               {}
func (*IMGroupGetShieldReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

func (m *IMGroupGetShieldReq) GetGroupId() uint32 {
	if m != nil && m.GroupId != nil {
		return *m.GroupId
	}
	return 0
}

func (m *IMGroupGetShieldReq) GetUserId() []uint32 {
	if m != nil {
		return m.UserId
	}
	return nil
}

func (m *IMGroupGetShieldReq) GetAttachData() []byte {
	if m != nil {
		return m.AttachData
	}
	return nil
}

type IMGroupGetShieldRsp struct {
	// cmd id: 			0x0712
	GroupId          *uint32                       `protobuf:"varint,1,req,name=group_id,json=groupId" json:"group_id,omitempty"`
	ShieldStatusList []*IM_BaseDefine.ShieldStatus `protobuf:"bytes,2,rep,name=shield_status_list,json=shieldStatusList" json:"shield_status_list,omitempty"`
	AttachData       []byte                        `protobuf:"bytes,20,opt,name=attach_data,json=attachData" json:"attach_data,omitempty"`
	XXX_unrecognized []byte                        `json:"-"`
}

func (m *IMGroupGetShieldRsp) Reset()                    { *m = IMGroupGetShieldRsp{} }
func (m *IMGroupGetShieldRsp) String() string            { return proto.CompactTextString(m) }
func (*IMGroupGetShieldRsp) ProtoMessage()               {}
func (*IMGroupGetShieldRsp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{15} }

func (m *IMGroupGetShieldRsp) GetGroupId() uint32 {
	if m != nil && m.GroupId != nil {
		return *m.GroupId
	}
	return 0
}

func (m *IMGroupGetShieldRsp) GetShieldStatusList() []*IM_BaseDefine.ShieldStatus {
	if m != nil {
		return m.ShieldStatusList
	}
	return nil
}

func (m *IMGroupGetShieldRsp) GetAttachData() []byte {
	if m != nil {
		return m.AttachData
	}
	return nil
}

type IMFileTransferReq struct {
	// cmd id:			0x0715
	FromUserId       *uint32                         `protobuf:"varint,1,req,name=from_user_id,json=fromUserId" json:"from_user_id,omitempty"`
	ToUserId         *uint32                         `protobuf:"varint,2,req,name=to_user_id,json=toUserId" json:"to_user_id,omitempty"`
	FileName         *string                         `protobuf:"bytes,3,req,name=file_name,json=fileName" json:"file_name,omitempty"`
	FileSize         *uint32                         `protobuf:"varint,4,req,name=file_size,json=fileSize" json:"file_size,omitempty"`
	TransMode        *IM_BaseDefine.TransferFileType `protobuf:"varint,5,req,name=trans_mode,json=transMode,enum=IM.BaseDefine.TransferFileType" json:"trans_mode,omitempty"`
	AttachData       []byte                          `protobuf:"bytes,20,opt,name=attach_data,json=attachData" json:"attach_data,omitempty"`
	XXX_unrecognized []byte                          `json:"-"`
}

func (m *IMFileTransferReq) Reset()                    { *m = IMFileTransferReq{} }
func (m *IMFileTransferReq) String() string            { return proto.CompactTextString(m) }
func (*IMFileTransferReq) ProtoMessage()               {}
func (*IMFileTransferReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{16} }

func (m *IMFileTransferReq) GetFromUserId() uint32 {
	if m != nil && m.FromUserId != nil {
		return *m.FromUserId
	}
	return 0
}

func (m *IMFileTransferReq) GetToUserId() uint32 {
	if m != nil && m.ToUserId != nil {
		return *m.ToUserId
	}
	return 0
}

func (m *IMFileTransferReq) GetFileName() string {
	if m != nil && m.FileName != nil {
		return *m.FileName
	}
	return ""
}

func (m *IMFileTransferReq) GetFileSize() uint32 {
	if m != nil && m.FileSize != nil {
		return *m.FileSize
	}
	return 0
}

func (m *IMFileTransferReq) GetTransMode() IM_BaseDefine.TransferFileType {
	if m != nil && m.TransMode != nil {
		return *m.TransMode
	}
	return IM_BaseDefine.TransferFileType_FILE_TYPE_ONLINE
}

func (m *IMFileTransferReq) GetAttachData() []byte {
	if m != nil {
		return m.AttachData
	}
	return nil
}

type IMFileTransferRsp struct {
	// cmd id:			0x0716
	ResultCode       *uint32                         `protobuf:"varint,1,req,name=result_code,json=resultCode" json:"result_code,omitempty"`
	FromUserId       *uint32                         `protobuf:"varint,2,req,name=from_user_id,json=fromUserId" json:"from_user_id,omitempty"`
	ToUserId         *uint32                         `protobuf:"varint,3,req,name=to_user_id,json=toUserId" json:"to_user_id,omitempty"`
	FileName         *string                         `protobuf:"bytes,4,opt,name=file_name,json=fileName" json:"file_name,omitempty"`
	FileSize         *uint32                         `protobuf:"varint,5,opt,name=file_size,json=fileSize" json:"file_size,omitempty"`
	TaskId           *string                         `protobuf:"bytes,6,opt,name=task_id,json=taskId" json:"task_id,omitempty"`
	TransMode        *IM_BaseDefine.TransferFileType `protobuf:"varint,7,opt,name=trans_mode,json=transMode,enum=IM.BaseDefine.TransferFileType" json:"trans_mode,omitempty"`
	AttachData       []byte                          `protobuf:"bytes,20,opt,name=attach_data,json=attachData" json:"attach_data,omitempty"`
	XXX_unrecognized []byte                          `json:"-"`
}

func (m *IMFileTransferRsp) Reset()                    { *m = IMFileTransferRsp{} }
func (m *IMFileTransferRsp) String() string            { return proto.CompactTextString(m) }
func (*IMFileTransferRsp) ProtoMessage()               {}
func (*IMFileTransferRsp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{17} }

func (m *IMFileTransferRsp) GetResultCode() uint32 {
	if m != nil && m.ResultCode != nil {
		return *m.ResultCode
	}
	return 0
}

func (m *IMFileTransferRsp) GetFromUserId() uint32 {
	if m != nil && m.FromUserId != nil {
		return *m.FromUserId
	}
	return 0
}

func (m *IMFileTransferRsp) GetToUserId() uint32 {
	if m != nil && m.ToUserId != nil {
		return *m.ToUserId
	}
	return 0
}

func (m *IMFileTransferRsp) GetFileName() string {
	if m != nil && m.FileName != nil {
		return *m.FileName
	}
	return ""
}

func (m *IMFileTransferRsp) GetFileSize() uint32 {
	if m != nil && m.FileSize != nil {
		return *m.FileSize
	}
	return 0
}

func (m *IMFileTransferRsp) GetTaskId() string {
	if m != nil && m.TaskId != nil {
		return *m.TaskId
	}
	return ""
}

func (m *IMFileTransferRsp) GetTransMode() IM_BaseDefine.TransferFileType {
	if m != nil && m.TransMode != nil {
		return *m.TransMode
	}
	return IM_BaseDefine.TransferFileType_FILE_TYPE_ONLINE
}

func (m *IMFileTransferRsp) GetAttachData() []byte {
	if m != nil {
		return m.AttachData
	}
	return nil
}

type IMFileServerIPReq struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *IMFileServerIPReq) Reset()                    { *m = IMFileServerIPReq{} }
func (m *IMFileServerIPReq) String() string            { return proto.CompactTextString(m) }
func (*IMFileServerIPReq) ProtoMessage()               {}
func (*IMFileServerIPReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{18} }

type IMFileServerIPRsp struct {
	// cmd id:			0x0718
	IpAddrList       []*IM_BaseDefine.IpAddr `protobuf:"bytes,1,rep,name=ip_addr_list,json=ipAddrList" json:"ip_addr_list,omitempty"`
	XXX_unrecognized []byte                  `json:"-"`
}

func (m *IMFileServerIPRsp) Reset()                    { *m = IMFileServerIPRsp{} }
func (m *IMFileServerIPRsp) String() string            { return proto.CompactTextString(m) }
func (*IMFileServerIPRsp) ProtoMessage()               {}
func (*IMFileServerIPRsp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{19} }

func (m *IMFileServerIPRsp) GetIpAddrList() []*IM_BaseDefine.IpAddr {
	if m != nil {
		return m.IpAddrList
	}
	return nil
}

func init() {
	proto.RegisterType((*IMStopReceivePacket)(nil), "IM.Server.IMStopReceivePacket")
	proto.RegisterType((*IMValidateReq)(nil), "IM.Server.IMValidateReq")
	proto.RegisterType((*IMValidateRsp)(nil), "IM.Server.IMValidateRsp")
	proto.RegisterType((*IMGetDeviceTokenReq)(nil), "IM.Server.IMGetDeviceTokenReq")
	proto.RegisterType((*IMGetDeviceTokenRsp)(nil), "IM.Server.IMGetDeviceTokenRsp")
	proto.RegisterType((*IMRoleSet)(nil), "IM.Server.IMRoleSet")
	proto.RegisterType((*IMOnlineUserInfo)(nil), "IM.Server.IMOnlineUserInfo")
	proto.RegisterType((*IMMsgServInfo)(nil), "IM.Server.IMMsgServInfo")
	proto.RegisterType((*IMUserStatusUpdate)(nil), "IM.Server.IMUserStatusUpdate")
	proto.RegisterType((*IMUserCntUpdate)(nil), "IM.Server.IMUserCntUpdate")
	proto.RegisterType((*IMServerKickUser)(nil), "IM.Server.IMServerKickUser")
	proto.RegisterType((*IMServerPCLoginStatusNotify)(nil), "IM.Server.IMServerPCLoginStatusNotify")
	proto.RegisterType((*IMPushToUserReq)(nil), "IM.Server.IMPushToUserReq")
	proto.RegisterType((*IMPushToUserRsp)(nil), "IM.Server.IMPushToUserRsp")
	proto.RegisterType((*IMGroupGetShieldReq)(nil), "IM.Server.IMGroupGetShieldReq")
	proto.RegisterType((*IMGroupGetShieldRsp)(nil), "IM.Server.IMGroupGetShieldRsp")
	proto.RegisterType((*IMFileTransferReq)(nil), "IM.Server.IMFileTransferReq")
	proto.RegisterType((*IMFileTransferRsp)(nil), "IM.Server.IMFileTransferRsp")
	proto.RegisterType((*IMFileServerIPReq)(nil), "IM.Server.IMFileServerIPReq")
	proto.RegisterType((*IMFileServerIPRsp)(nil), "IM.Server.IMFileServerIPRsp")
}

func init() { proto.RegisterFile("IM.Server.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 970 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0xdd, 0x6e, 0xdb, 0x36,
	0x14, 0x86, 0xec, 0xfc, 0xf9, 0xd8, 0x4e, 0x33, 0xa5, 0x5b, 0xdc, 0xba, 0x43, 0x3c, 0xf5, 0x26,
	0x37, 0x0b, 0xb0, 0x60, 0xc0, 0x80, 0x5d, 0x0c, 0x68, 0x1d, 0xac, 0x13, 0x1a, 0xb5, 0x81, 0x9c,
	0x74, 0xd8, 0x95, 0x40, 0x48, 0x74, 0xcc, 0x45, 0x26, 0x59, 0x92, 0xca, 0x9a, 0x62, 0xc0, 0xb0,
	0xdb, 0x3d, 0xc3, 0x1e, 0x20, 0xaf, 0xb2, 0x97, 0xd9, 0x33, 0x0c, 0x87, 0x94, 0x1c, 0x5b, 0xf9,
	0xf1, 0x02, 0xf4, 0x4a, 0xe4, 0xd1, 0xe1, 0x39, 0xdf, 0xf9, 0xf8, 0x9d, 0x43, 0x78, 0x14, 0x46,
	0xfb, 0x23, 0xaa, 0x2e, 0xa8, 0xda, 0x97, 0x4a, 0x18, 0xe1, 0xb7, 0x66, 0x86, 0xa7, 0xdb, 0x61,
	0xb4, 0xff, 0x92, 0x68, 0x7a, 0x48, 0xc7, 0x8c, 0x53, 0xf7, 0x3f, 0xf8, 0x1a, 0xb6, 0xc3, 0x68,
	0x64, 0x84, 0x8c, 0x69, 0x4a, 0xd9, 0x05, 0x3d, 0x26, 0xe9, 0x39, 0x35, 0xfe, 0x17, 0xb0, 0xa6,
	0xa8, 0x2e, 0x72, 0xd3, 0xf3, 0x06, 0x8d, 0xbd, 0x6e, 0x5c, 0xee, 0x02, 0x06, 0xdd, 0x30, 0x7a,
	0x47, 0x72, 0x96, 0x11, 0x43, 0x63, 0xfa, 0xde, 0xef, 0x43, 0xab, 0xd0, 0x54, 0x25, 0x9c, 0x4c,
	0xa9, 0xf5, 0x6d, 0xc5, 0x1b, 0x68, 0x78, 0x43, 0xa6, 0xd4, 0x7f, 0x0a, 0x1b, 0x92, 0x68, 0xfd,
	0x9b, 0x50, 0x59, 0xaf, 0xe1, 0xfe, 0x55, 0x7b, 0x7f, 0x17, 0xda, 0xc4, 0x18, 0x92, 0x4e, 0x92,
	0x8c, 0x18, 0xd2, 0x7b, 0x3c, 0xf0, 0xf6, 0x3a, 0x31, 0x38, 0xd3, 0x21, 0x31, 0x24, 0xf8, 0xc7,
	0x5b, 0xc8, 0xa5, 0xe5, 0xfd, 0xb9, 0x76, 0xa1, 0xed, 0x30, 0x26, 0xa9, 0xc8, 0xa8, 0x4d, 0xd7,
	0x8d, 0xc1, 0x99, 0x86, 0x22, 0xa3, 0xfe, 0x73, 0xe8, 0x96, 0x0e, 0xda, 0x28, 0xc6, 0xcf, 0x7a,
	0xcd, 0x81, 0xb7, 0xd7, 0x8a, 0x3b, 0xce, 0x38, 0xb2, 0x36, 0xff, 0xdb, 0x32, 0x05, 0xe3, 0x63,
	0xd1, 0x5b, 0x19, 0x78, 0x7b, 0xed, 0x83, 0x9d, 0xfd, 0x45, 0xde, 0x4e, 0x35, 0x55, 0x21, 0x1f,
	0x0b, 0x97, 0x1b, 0x57, 0xcb, 0x6b, 0x79, 0x8b, 0x2c, 0xbf, 0xa2, 0xe6, 0x90, 0x5e, 0xb0, 0x94,
	0x9e, 0x88, 0x73, 0xca, 0x91, 0xbc, 0x1d, 0x58, 0x77, 0xd9, 0xb2, 0x9e, 0x37, 0x68, 0x22, 0xcd,
	0x36, 0xe4, 0xff, 0x20, 0xe7, 0xf7, 0x5b, 0x02, 0x6a, 0xe9, 0x1f, 0xc2, 0x23, 0x1b, 0xd0, 0xa0,
	0xc1, 0x15, 0x81, 0x81, 0xdb, 0x07, 0xcf, 0x6e, 0x29, 0xc2, 0x9e, 0xb2, 0x95, 0x74, 0x8b, 0xf9,
	0xed, 0xf2, 0xec, 0xcf, 0xa1, 0x15, 0x46, 0xb1, 0xc8, 0xe9, 0xc8, 0x49, 0x65, 0x4a, 0xb4, 0xa1,
	0xaa, 0x92, 0x8a, 0xdb, 0x05, 0x3f, 0xc3, 0x56, 0x18, 0xbd, 0xe5, 0x39, 0xe3, 0xb4, 0xa2, 0xcc,
	0x1f, 0xc2, 0xa6, 0xc5, 0xa7, 0x0d, 0x31, 0x49, 0xce, 0xb4, 0x29, 0xe1, 0x7d, 0x59, 0x83, 0xe7,
	0x14, 0x8b, 0xc7, 0x46, 0x86, 0x98, 0xb8, 0x53, 0x94, 0xab, 0x23, 0xa6, 0x4d, 0x70, 0x65, 0x85,
	0x11, 0xe9, 0x33, 0xf4, 0xb2, 0x61, 0xb7, 0xa0, 0xc9, 0xe4, 0x37, 0xa5, 0x24, 0x70, 0xe9, 0x2c,
	0x07, 0xa5, 0xe8, 0x70, 0xe9, 0xfb, 0xb0, 0x22, 0x85, 0x32, 0xbd, 0xa6, 0x05, 0x69, 0xd7, 0xfe,
	0x00, 0x3a, 0x53, 0xf2, 0x21, 0x49, 0x05, 0xe7, 0x49, 0xca, 0x4d, 0x6f, 0xc5, 0x89, 0x66, 0x4a,
	0x3e, 0x0c, 0x05, 0xe7, 0x43, 0x6e, 0x3d, 0xd2, 0x42, 0x5d, 0x7b, 0xac, 0x3a, 0x8f, 0xb4, 0x50,
	0x95, 0x47, 0x1f, 0x5a, 0x13, 0xa1, 0x8d, 0x13, 0xe5, 0x9a, 0x13, 0x25, 0x1a, 0x50, 0x94, 0xc1,
	0x5f, 0x1e, 0xf8, 0x61, 0x54, 0xd5, 0x51, 0xe8, 0x53, 0x89, 0x5a, 0x46, 0x82, 0x67, 0x34, 0x14,
	0xba, 0xe4, 0x0d, 0x8a, 0x99, 0xdb, 0xbc, 0x30, 0x9c, 0x90, 0x2b, 0x61, 0x7c, 0x0f, 0xed, 0x34,
	0x67, 0x94, 0x9b, 0xc4, 0x5c, 0x4a, 0x6a, 0x8b, 0xd9, 0x3c, 0x78, 0x52, 0x63, 0x6f, 0x68, 0x3d,
	0x4e, 0x2e, 0x25, 0x8d, 0x21, 0x9d, 0xad, 0x83, 0xd7, 0x38, 0x1d, 0x10, 0xcb, 0x90, 0x9b, 0x1a,
	0x10, 0x92, 0x1a, 0x26, 0xf8, 0x3c, 0x90, 0x17, 0xd6, 0x72, 0x27, 0x90, 0xe0, 0x0f, 0xbc, 0x5d,
	0x77, 0x4d, 0xaf, 0x59, 0x7a, 0x8e, 0x61, 0x17, 0xe5, 0x7c, 0x0f, 0xea, 0xc6, 0x03, 0x50, 0xbb,
	0x49, 0x44, 0xb4, 0xe0, 0xe5, 0xcd, 0x95, 0xbb, 0xe0, 0x17, 0xe8, 0x57, 0x00, 0x8e, 0x87, 0x47,
	0xe2, 0x8c, 0x71, 0xc7, 0xdd, 0x1b, 0x61, 0xd8, 0xf8, 0xf2, 0x6e, 0x2c, 0x5f, 0x41, 0x27, 0x47,
	0xef, 0x8a, 0x7c, 0x57, 0x56, 0x3b, 0xbf, 0x8e, 0x10, 0xfc, 0xe9, 0x21, 0x53, 0xc7, 0x85, 0x9e,
	0x9c, 0x08, 0x2c, 0x0c, 0x5b, 0xf5, 0x31, 0xac, 0x8e, 0x73, 0xa2, 0x27, 0xa5, 0xc8, 0xdc, 0x06,
	0x45, 0x65, 0x5b, 0xc4, 0xe9, 0xcc, 0xae, 0x6b, 0x3d, 0x68, 0x45, 0xde, 0x7c, 0x50, 0x0f, 0x5a,
	0x91, 0xbf, 0xab, 0x41, 0xd0, 0xd2, 0x1f, 0xc2, 0x96, 0x2c, 0xf4, 0x24, 0x29, 0xa7, 0xd8, 0x5c,
	0xfb, 0xd4, 0xa9, 0xc4, 0x73, 0xb1, 0xf5, 0x8a, 0x37, 0xe5, 0x6c, 0x6d, 0xe3, 0xfe, 0x6a, 0x07,
	0x87, 0x12, 0x85, 0x7c, 0x45, 0xcd, 0x68, 0xc2, 0x68, 0x9e, 0x61, 0x79, 0x4f, 0x60, 0xe3, 0x0c,
	0x8d, 0xd7, 0x7c, 0xad, 0xdb, 0x7d, 0x98, 0x2d, 0x4a, 0xe0, 0x41, 0x43, 0xea, 0x6f, 0xef, 0x96,
	0x64, 0x5a, 0xde, 0x97, 0x2c, 0x04, 0x5f, 0x5b, 0xbf, 0xf2, 0x7a, 0x5c, 0x95, 0x0d, 0x5b, 0x65,
	0xbf, 0x3e, 0x24, 0xac, 0xa3, 0xbb, 0xb3, 0x78, 0x4b, 0xcf, 0xed, 0xb0, 0xd2, 0xe5, 0xf0, 0xfe,
	0xf5, 0xe0, 0xb3, 0x30, 0xfa, 0x91, 0xe5, 0xf4, 0x44, 0x11, 0xae, 0xc7, 0xee, 0xa2, 0x07, 0xd0,
	0x19, 0x2b, 0x31, 0x4d, 0x16, 0xd5, 0x03, 0x68, 0x3b, 0x75, 0x75, 0x3f, 0x03, 0x30, 0x22, 0x59,
	0x6c, 0x8b, 0x0d, 0x23, 0xca, 0xbf, 0x7d, 0x68, 0x8d, 0x59, 0x4e, 0xdd, 0x3c, 0x68, 0xba, 0x79,
	0x80, 0x06, 0xfb, 0x48, 0x55, 0x3f, 0x35, 0xfb, 0x48, 0xcb, 0x69, 0x63, 0x7f, 0x8e, 0xd8, 0x47,
	0xea, 0xff, 0x00, 0x60, 0x10, 0x48, 0x32, 0xc5, 0x07, 0x6c, 0xd5, 0x36, 0xc9, 0x6e, 0xad, 0xe6,
	0x0a, 0xa9, 0x45, 0x8d, 0xad, 0xd2, 0xb2, 0x47, 0x22, 0x7c, 0xe0, 0x96, 0x16, 0x7c, 0xd5, 0xb8,
	0x51, 0xb0, 0x96, 0xf5, 0x87, 0xd3, 0xbb, 0xf1, 0x70, 0xd6, 0x19, 0x69, 0x2c, 0x61, 0xa4, 0x79,
	0x1f, 0x23, 0x2b, 0xf6, 0xd1, 0xbd, 0x83, 0x91, 0xd5, 0x81, 0xb7, 0xc0, 0xc8, 0x0e, 0xac, 0x1b,
	0xa2, 0xcf, 0x31, 0xe8, 0x9a, 0x3d, 0xb7, 0x86, 0xdb, 0x30, 0xab, 0x51, 0xb5, 0x3e, 0xf0, 0x3e,
	0x35, 0x55, 0xdb, 0x15, 0x53, 0x6e, 0xc2, 0x84, 0xc7, 0x31, 0x7d, 0x1f, 0x1c, 0xdd, 0x30, 0x6a,
	0xe9, 0x7f, 0x07, 0x1d, 0x26, 0x13, 0x92, 0x65, 0x6a, 0xbe, 0x23, 0x3f, 0xaf, 0x81, 0x09, 0xe5,
	0x8b, 0x2c, 0x53, 0x31, 0x30, 0xfb, 0x45, 0x7d, 0xbe, 0x6c, 0xfc, 0xd4, 0xfc, 0x2f, 0x00, 0x00,
	0xff, 0xff, 0xc0, 0xec, 0xc3, 0x09, 0xaf, 0x09, 0x00, 0x00,
}
