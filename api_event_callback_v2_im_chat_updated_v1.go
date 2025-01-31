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

// EventV2IMChatUpdatedV1 群组配置被修改后触发此事件, 包含:
//
// - 群主转移
// - 群基本信息修改(群头像/群名称/群描述/群国际化名称)
// - 群权限修改(加人入群权限/群编辑权限/at所有人权限/群分享权限)。{使用示例}(url=/api/tools/api_explore/api_explore_config?project=im&version=v1&resource=chat&event=updated)
// 注意事项:
// - 需要开启[机器人能力](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-enable-bot-ability)
// - 需要订阅 [消息与群组] 分类下的 [群配置修改] 事件
// - 事件会向群内订阅了该事件的机器人进行推送
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat/events/updated
func (r *EventCallbackService) HandlerEventV2IMChatUpdatedV1(f EventV2IMChatUpdatedV1Handler) {
	r.cli.eventHandler.eventV2IMChatUpdatedV1Handler = f
}

// EventV2IMChatUpdatedV1Handler event EventV2IMChatUpdatedV1 handler
type EventV2IMChatUpdatedV1Handler func(ctx context.Context, cli *Lark, schema string, header *EventHeaderV2, event *EventV2IMChatUpdatedV1) (string, error)

// EventV2IMChatUpdatedV1 ...
type EventV2IMChatUpdatedV1 struct {
	ChatID            string                               `json:"chat_id,omitempty"`             // 群组 ID, 详情参见[群ID 说明](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat-id-description)
	OperatorID        *EventV2IMChatUpdatedV1OperatorID    `json:"operator_id,omitempty"`         // 用户 ID
	External          bool                                 `json:"external,omitempty"`            // 是否是外部群
	OperatorTenantKey string                               `json:"operator_tenant_key,omitempty"` // 操作者租户 Key
	AfterChange       *EventV2IMChatUpdatedV1AfterChange   `json:"after_change,omitempty"`        // 更新后的群信息
	BeforeChange      *EventV2IMChatUpdatedV1BeforeChange  `json:"before_change,omitempty"`       // 更新前的群信息
	ModeratorList     *EventV2IMChatUpdatedV1ModeratorList `json:"moderator_list,omitempty"`      // 群可发言成员名单的变更信息
}

// EventV2IMChatUpdatedV1AfterChange ...
type EventV2IMChatUpdatedV1AfterChange struct {
	Avatar                 string                                    `json:"avatar,omitempty"`                   // 群头像
	Name                   string                                    `json:"name,omitempty"`                     // 群名称
	Description            string                                    `json:"description,omitempty"`              // 群描述
	I18nNames              *I18nNames                                `json:"i18n_names,omitempty"`               // 群国际化名称
	AddMemberPermission    AddMemberPermission                       `json:"add_member_permission,omitempty"`    // 加人入群权限(all_members/only_owner/unknown)
	ShareCardPermission    ShareCardPermission                       `json:"share_card_permission,omitempty"`    // 群分享权限(allowed/not_allowed/unknown)
	AtAllPermission        AtAllPermission                           `json:"at_all_permission,omitempty"`        // at 所有人权限(all_members/only_owner/unknown)
	EditPermission         EditPermission                            `json:"edit_permission,omitempty"`          // 群编辑权限(all_members/only_owner/unknown)
	MembershipApproval     MembershipApproval                        `json:"membership_approval,omitempty"`      // 加群审批(no_approval_required/approval_required)
	JoinMessageVisibility  MessageVisibility                         `json:"join_message_visibility,omitempty"`  // 入群消息可见性(only_owner/all_members/not_anyone)
	LeaveMessageVisibility MessageVisibility                         `json:"leave_message_visibility,omitempty"` // 出群消息可见性(only_owner/all_members/not_anyone)
	ModerationPermission   ModerationPermission                      `json:"moderation_permission,omitempty"`    // 发言权限(all_members/only_owner)
	OwnerID                *EventV2IMChatUpdatedV1AfterChangeOwnerID `json:"owner_id,omitempty"`                 // 用户 ID
}

// EventV2IMChatUpdatedV1AfterChangeOwnerID ...
type EventV2IMChatUpdatedV1AfterChangeOwnerID struct {
	UnionID string `json:"union_id,omitempty"` // 用户的 union id
	UserID  string `json:"user_id,omitempty"`  // 用户的 user id, 字段权限要求: 获取用户 user ID
	OpenID  string `json:"open_id,omitempty"`  // 用户的 open id
}

// EventV2IMChatUpdatedV1BeforeChange ...
type EventV2IMChatUpdatedV1BeforeChange struct {
	Avatar                 string                                     `json:"avatar,omitempty"`                   // 群头像
	Name                   string                                     `json:"name,omitempty"`                     // 群名称
	Description            string                                     `json:"description,omitempty"`              // 群描述
	I18nNames              *I18nNames                                 `json:"i18n_names,omitempty"`               // 群国际化名称
	AddMemberPermission    AddMemberPermission                        `json:"add_member_permission,omitempty"`    // 加人入群权限(all_members/only_owner/unknown)
	ShareCardPermission    ShareCardPermission                        `json:"share_card_permission,omitempty"`    // 群分享权限(allowed/not_allowed/unknown)
	AtAllPermission        AtAllPermission                            `json:"at_all_permission,omitempty"`        // at 所有人权限(all_members/only_owner/unknown)
	EditPermission         EditPermission                             `json:"edit_permission,omitempty"`          // 群编辑权限(all_members/only_owner/unknown)
	MembershipApproval     MembershipApproval                         `json:"membership_approval,omitempty"`      // 加群审批(no_approval_required/approval_required)
	JoinMessageVisibility  MessageVisibility                          `json:"join_message_visibility,omitempty"`  // 入群消息可见性(only_owner/all_members/not_anyone)
	LeaveMessageVisibility MessageVisibility                          `json:"leave_message_visibility,omitempty"` // 出群消息可见性(only_owner/all_members/not_anyone)
	ModerationPermission   ModerationPermission                       `json:"moderation_permission,omitempty"`    // 发言权限(all_members/only_owner)
	OwnerID                *EventV2IMChatUpdatedV1BeforeChangeOwnerID `json:"owner_id,omitempty"`                 // 用户 ID
}

// EventV2IMChatUpdatedV1BeforeChangeOwnerID ...
type EventV2IMChatUpdatedV1BeforeChangeOwnerID struct {
	UnionID string `json:"union_id,omitempty"` // 用户的 union id
	UserID  string `json:"user_id,omitempty"`  // 用户的 user id, 字段权限要求: 获取用户 user ID
	OpenID  string `json:"open_id,omitempty"`  // 用户的 open id
}

// EventV2IMChatUpdatedV1ModeratorList ...
type EventV2IMChatUpdatedV1ModeratorList struct {
	AddedMemberList   []*EventV2IMChatUpdatedV1ModeratorListAddedMember   `json:"added_member_list,omitempty"`   // 被添加进可发言名单的用户列表（列表中一定会有owner）
	RemovedMemberList []*EventV2IMChatUpdatedV1ModeratorListRemovedMember `json:"removed_member_list,omitempty"` // 被移除出可发言名单的用户列表
}

// EventV2IMChatUpdatedV1ModeratorListAddedMember ...
type EventV2IMChatUpdatedV1ModeratorListAddedMember struct {
	TenantKey string                                                `json:"tenant_key,omitempty"` // 租户 Key
	UserID    *EventV2IMChatUpdatedV1ModeratorListAddedMemberUserID `json:"user_id,omitempty"`    // 用户 ID
}

// EventV2IMChatUpdatedV1ModeratorListAddedMemberUserID ...
type EventV2IMChatUpdatedV1ModeratorListAddedMemberUserID struct {
	UnionID string `json:"union_id,omitempty"` // 用户的 union id
	UserID  string `json:"user_id,omitempty"`  // 用户的 user id, 字段权限要求: 获取用户 user ID
	OpenID  string `json:"open_id,omitempty"`  // 用户的 open id
}

// EventV2IMChatUpdatedV1ModeratorListRemovedMember ...
type EventV2IMChatUpdatedV1ModeratorListRemovedMember struct {
	TenantKey string                                                  `json:"tenant_key,omitempty"` // 租户 Key
	UserID    *EventV2IMChatUpdatedV1ModeratorListRemovedMemberUserID `json:"user_id,omitempty"`    // 用户 ID
}

// EventV2IMChatUpdatedV1ModeratorListRemovedMemberUserID ...
type EventV2IMChatUpdatedV1ModeratorListRemovedMemberUserID struct {
	UnionID string `json:"union_id,omitempty"` // 用户的 union id
	UserID  string `json:"user_id,omitempty"`  // 用户的 user id, 字段权限要求: 获取用户 user ID
	OpenID  string `json:"open_id,omitempty"`  // 用户的 open id
}

// EventV2IMChatUpdatedV1OperatorID ...
type EventV2IMChatUpdatedV1OperatorID struct {
	UnionID string `json:"union_id,omitempty"` // 用户的 union id
	UserID  string `json:"user_id,omitempty"`  // 用户的 user id, 字段权限要求: 获取用户 user ID
	OpenID  string `json:"open_id,omitempty"`  // 用户的 open id
}
