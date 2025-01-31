// Code generated by lark_sdk_gen. DO NOT EDIT.
/**
 * Copyright 2022 chyroc
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package lark

import (
	"context"
)

// GetVCScopeConfig 该接口可以用来查询某个会议层级范围下或者某个会议室的配置
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/vc-v1/scope_config/get
func (r *VCService) GetVCScopeConfig(ctx context.Context, request *GetVCScopeConfigReq, options ...MethodOptionFunc) (*GetVCScopeConfigResp, *Response, error) {
	if r.cli.mock.mockVCGetVCScopeConfig != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] VC#GetVCScopeConfig mock enable")
		return r.cli.mock.mockVCGetVCScopeConfig(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "VC",
		API:                   "GetVCScopeConfig",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/vc/v1/scope_config",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(getVCScopeConfigResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockVCGetVCScopeConfig mock VCGetVCScopeConfig method
func (r *Mock) MockVCGetVCScopeConfig(f func(ctx context.Context, request *GetVCScopeConfigReq, options ...MethodOptionFunc) (*GetVCScopeConfigResp, *Response, error)) {
	r.mockVCGetVCScopeConfig = f
}

// UnMockVCGetVCScopeConfig un-mock VCGetVCScopeConfig method
func (r *Mock) UnMockVCGetVCScopeConfig() {
	r.mockVCGetVCScopeConfig = nil
}

// GetVCScopeConfigReq ...
type GetVCScopeConfigReq struct {
	ScopeType  int64   `query:"scope_type" json:"-"`   // 查询节点范围, 示例值: 1, 可选值有: 1: 会议室层级, 2: 会议室
	ScopeID    string  `query:"scope_id" json:"-"`     // 查询节点ID: 如果scope_type为1, 则为层级ID, 如果scope_type为2, 则为会议室ID, 示例值: "omm_608d34d82d531b27fa993902d350a307"
	UserIDType *IDType `query:"user_id_type" json:"-"` // 用户 ID 类型, 示例值: "open_id", 可选值有: open_id: 用户的 open id, union_id: 用户的 union id, user_id: 用户的 user id, 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
}

// GetVCScopeConfigResp ...
type GetVCScopeConfigResp struct {
	CurrentConfig *GetVCScopeConfigRespCurrentConfig  `json:"current_config,omitempty"` // 当前节点的配置, 根据层级顺序从底向上进行合并计算后的结果；如果当前节点某个值已配置, 则取该节点的值, 否则会从该节点的父层级节点获取, 如果父节点依然未配置, 则继续向上递归获取；若所有节点均未配置, 则该值返回为空
	OriginConfigs []*GetVCScopeConfigRespOriginConfig `json:"origin_configs,omitempty"` // 所有节点的原始配置, 按照层级顺序从底向上返回；如果某节点某个值未配置, 则该值返回为空
}

// GetVCScopeConfigRespCurrentConfig ...
type GetVCScopeConfigRespCurrentConfig struct {
	ScopeType   int64                                         `json:"scope_type,omitempty"`   // 查询节点范围, 可选值有: 1: 会议室层级, 2: 会议室
	ScopeID     string                                        `json:"scope_id,omitempty"`     // 查询节点ID: 如果scope_type为1, 则为层级ID, 如果scope_type为2, 则为会议室ID
	ScopeConfig *GetVCScopeConfigRespCurrentConfigScopeConfig `json:"scope_config,omitempty"` // 节点配置
}

// GetVCScopeConfigRespCurrentConfigScopeConfig ...
type GetVCScopeConfigRespCurrentConfigScopeConfig struct {
	RoomBackground        string                                                             `json:"room_background,omitempty"`          // 飞书会议室背景图
	DisplayBackground     string                                                             `json:"display_background,omitempty"`       // 飞书签到板背景图
	DigitalSignage        *GetVCScopeConfigRespCurrentConfigScopeConfigDigitalSignage        `json:"digital_signage,omitempty"`          // 飞书会议室数字标牌
	RoomBoxDigitalSignage *GetVCScopeConfigRespCurrentConfigScopeConfigRoomBoxDigitalSignage `json:"room_box_digital_signage,omitempty"` // 飞书投屏盒子数字标牌
	RoomStatus            *GetVCScopeConfigRespCurrentConfigScopeConfigRoomStatus            `json:"room_status,omitempty"`              // 会议室状态
}

// GetVCScopeConfigRespCurrentConfigScopeConfigDigitalSignage ...
type GetVCScopeConfigRespCurrentConfigScopeConfigDigitalSignage struct {
	Enable       bool                                                                  `json:"enable,omitempty"`        // 是否开启数字标牌功能
	Mute         bool                                                                  `json:"mute,omitempty"`          // 是否静音播放
	StartDisplay int64                                                                 `json:"start_display,omitempty"` // 日程会议开始前n分钟结束播放
	StopDisplay  int64                                                                 `json:"stop_display,omitempty"`  // 会议结束后n分钟开始播放
	Materials    []*GetVCScopeConfigRespCurrentConfigScopeConfigDigitalSignageMaterial `json:"materials,omitempty"`     // 素材列表
}

// GetVCScopeConfigRespCurrentConfigScopeConfigDigitalSignageMaterial ...
type GetVCScopeConfigRespCurrentConfigScopeConfigDigitalSignageMaterial struct {
	ID           string `json:"id,omitempty"`            // 素材ID
	Name         string `json:"name,omitempty"`          // 素材名称
	MaterialType int64  `json:"material_type,omitempty"` // 素材类型, 可选值有: 1: 图片, 2: 视频, 3: GIF
	URL          string `json:"url,omitempty"`           // 素材url
	Duration     int64  `json:"duration,omitempty"`      // 播放时长（单位sec）
	Cover        string `json:"cover,omitempty"`         // 素材封面url
	Md5          string `json:"md5,omitempty"`           // 素材文件md5
	Vid          string `json:"vid,omitempty"`           // 素材文件vid
	Size         string `json:"size,omitempty"`          // 素材文件大小（单位byte）
}

// GetVCScopeConfigRespCurrentConfigScopeConfigRoomBoxDigitalSignage ...
type GetVCScopeConfigRespCurrentConfigScopeConfigRoomBoxDigitalSignage struct {
	Enable       bool                                                                         `json:"enable,omitempty"`        // 是否开启数字标牌功能
	Mute         bool                                                                         `json:"mute,omitempty"`          // 是否静音播放
	StartDisplay int64                                                                        `json:"start_display,omitempty"` // 日程会议开始前n分钟结束播放
	StopDisplay  int64                                                                        `json:"stop_display,omitempty"`  // 会议结束后n分钟开始播放
	Materials    []*GetVCScopeConfigRespCurrentConfigScopeConfigRoomBoxDigitalSignageMaterial `json:"materials,omitempty"`     // 素材列表
}

// GetVCScopeConfigRespCurrentConfigScopeConfigRoomBoxDigitalSignageMaterial ...
type GetVCScopeConfigRespCurrentConfigScopeConfigRoomBoxDigitalSignageMaterial struct {
	ID           string `json:"id,omitempty"`            // 素材ID
	Name         string `json:"name,omitempty"`          // 素材名称
	MaterialType int64  `json:"material_type,omitempty"` // 素材类型, 可选值有: 1: 图片, 2: 视频, 3: GIF
	URL          string `json:"url,omitempty"`           // 素材url
	Duration     int64  `json:"duration,omitempty"`      // 播放时长（单位sec）
	Cover        string `json:"cover,omitempty"`         // 素材封面url
	Md5          string `json:"md5,omitempty"`           // 素材文件md5
	Vid          string `json:"vid,omitempty"`           // 素材文件vid
	Size         string `json:"size,omitempty"`          // 素材文件大小（单位byte）
}

// GetVCScopeConfigRespCurrentConfigScopeConfigRoomStatus ...
type GetVCScopeConfigRespCurrentConfigScopeConfigRoomStatus struct {
	Status           bool     `json:"status,omitempty"`             // 是否启用会议室
	ScheduleStatus   bool     `json:"schedule_status,omitempty"`    // 会议室未来状态为启用或禁用
	DisableStartTime string   `json:"disable_start_time,omitempty"` // 禁用开始时间（unix时间, 单位sec）
	DisableEndTime   string   `json:"disable_end_time,omitempty"`   // 禁用结束时间（unix时间, 单位sec, 数值0表示永久禁用）
	DisableReason    string   `json:"disable_reason,omitempty"`     // 禁用原因
	ContactIDs       []string `json:"contact_ids,omitempty"`        // 联系人列表, id类型由user_id_type参数决定
	DisableNotice    bool     `json:"disable_notice,omitempty"`     // 是否在禁用时发送通知给预定了该会议室的员工
	ResumeNotice     bool     `json:"resume_notice,omitempty"`      // 是否在恢复启用时发送通知给预定了该会议室的员工
}

// GetVCScopeConfigRespOriginConfig ...
type GetVCScopeConfigRespOriginConfig struct {
	ScopeType   int64                                        `json:"scope_type,omitempty"`   // 查询节点范围, 可选值有: 1: 会议室层级, 2: 会议室
	ScopeID     string                                       `json:"scope_id,omitempty"`     // 查询节点ID: 如果scope_type为1, 则为层级ID, 如果scope_type为2, 则为会议室ID
	ScopeConfig *GetVCScopeConfigRespOriginConfigScopeConfig `json:"scope_config,omitempty"` // 节点配置
}

// GetVCScopeConfigRespOriginConfigScopeConfig ...
type GetVCScopeConfigRespOriginConfigScopeConfig struct {
	RoomBackground        string                                                            `json:"room_background,omitempty"`          // 飞书会议室背景图
	DisplayBackground     string                                                            `json:"display_background,omitempty"`       // 飞书签到板背景图
	DigitalSignage        *GetVCScopeConfigRespOriginConfigScopeConfigDigitalSignage        `json:"digital_signage,omitempty"`          // 飞书会议室数字标牌
	RoomBoxDigitalSignage *GetVCScopeConfigRespOriginConfigScopeConfigRoomBoxDigitalSignage `json:"room_box_digital_signage,omitempty"` // 飞书投屏盒子数字标牌
	RoomStatus            *GetVCScopeConfigRespOriginConfigScopeConfigRoomStatus            `json:"room_status,omitempty"`              // 会议室状态
}

// GetVCScopeConfigRespOriginConfigScopeConfigDigitalSignage ...
type GetVCScopeConfigRespOriginConfigScopeConfigDigitalSignage struct {
	Enable       bool                                                                 `json:"enable,omitempty"`        // 是否开启数字标牌功能
	Mute         bool                                                                 `json:"mute,omitempty"`          // 是否静音播放
	StartDisplay int64                                                                `json:"start_display,omitempty"` // 日程会议开始前n分钟结束播放
	StopDisplay  int64                                                                `json:"stop_display,omitempty"`  // 会议结束后n分钟开始播放
	Materials    []*GetVCScopeConfigRespOriginConfigScopeConfigDigitalSignageMaterial `json:"materials,omitempty"`     // 素材列表
}

// GetVCScopeConfigRespOriginConfigScopeConfigDigitalSignageMaterial ...
type GetVCScopeConfigRespOriginConfigScopeConfigDigitalSignageMaterial struct {
	ID           string `json:"id,omitempty"`            // 素材ID
	Name         string `json:"name,omitempty"`          // 素材名称
	MaterialType int64  `json:"material_type,omitempty"` // 素材类型, 可选值有: 1: 图片, 2: 视频, 3: GIF
	URL          string `json:"url,omitempty"`           // 素材url
	Duration     int64  `json:"duration,omitempty"`      // 播放时长（单位sec）
	Cover        string `json:"cover,omitempty"`         // 素材封面url
	Md5          string `json:"md5,omitempty"`           // 素材文件md5
	Vid          string `json:"vid,omitempty"`           // 素材文件vid
	Size         string `json:"size,omitempty"`          // 素材文件大小（单位byte）
}

// GetVCScopeConfigRespOriginConfigScopeConfigRoomBoxDigitalSignage ...
type GetVCScopeConfigRespOriginConfigScopeConfigRoomBoxDigitalSignage struct {
	Enable       bool                                                                        `json:"enable,omitempty"`        // 是否开启数字标牌功能
	Mute         bool                                                                        `json:"mute,omitempty"`          // 是否静音播放
	StartDisplay int64                                                                       `json:"start_display,omitempty"` // 日程会议开始前n分钟结束播放
	StopDisplay  int64                                                                       `json:"stop_display,omitempty"`  // 会议结束后n分钟开始播放
	Materials    []*GetVCScopeConfigRespOriginConfigScopeConfigRoomBoxDigitalSignageMaterial `json:"materials,omitempty"`     // 素材列表
}

// GetVCScopeConfigRespOriginConfigScopeConfigRoomBoxDigitalSignageMaterial ...
type GetVCScopeConfigRespOriginConfigScopeConfigRoomBoxDigitalSignageMaterial struct {
	ID           string `json:"id,omitempty"`            // 素材ID
	Name         string `json:"name,omitempty"`          // 素材名称
	MaterialType int64  `json:"material_type,omitempty"` // 素材类型, 可选值有: 1: 图片, 2: 视频, 3: GIF
	URL          string `json:"url,omitempty"`           // 素材url
	Duration     int64  `json:"duration,omitempty"`      // 播放时长（单位sec）
	Cover        string `json:"cover,omitempty"`         // 素材封面url
	Md5          string `json:"md5,omitempty"`           // 素材文件md5
	Vid          string `json:"vid,omitempty"`           // 素材文件vid
	Size         string `json:"size,omitempty"`          // 素材文件大小（单位byte）
}

// GetVCScopeConfigRespOriginConfigScopeConfigRoomStatus ...
type GetVCScopeConfigRespOriginConfigScopeConfigRoomStatus struct {
	Status           bool     `json:"status,omitempty"`             // 是否启用会议室
	ScheduleStatus   bool     `json:"schedule_status,omitempty"`    // 会议室未来状态为启用或禁用
	DisableStartTime string   `json:"disable_start_time,omitempty"` // 禁用开始时间（unix时间, 单位sec）
	DisableEndTime   string   `json:"disable_end_time,omitempty"`   // 禁用结束时间（unix时间, 单位sec, 数值0表示永久禁用）
	DisableReason    string   `json:"disable_reason,omitempty"`     // 禁用原因
	ContactIDs       []string `json:"contact_ids,omitempty"`        // 联系人列表, id类型由user_id_type参数决定
	DisableNotice    bool     `json:"disable_notice,omitempty"`     // 是否在禁用时发送通知给预定了该会议室的员工
	ResumeNotice     bool     `json:"resume_notice,omitempty"`      // 是否在恢复启用时发送通知给预定了该会议室的员工
}

// getVCScopeConfigResp ...
type getVCScopeConfigResp struct {
	Code int64                 `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                `json:"msg,omitempty"`  // 错误描述
	Data *GetVCScopeConfigResp `json:"data,omitempty"`
}
