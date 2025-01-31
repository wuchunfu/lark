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

// BatchCreateAttendanceUserDailyShift 班表是用来描述考勤组内人员每天按哪个班次进行上班。目前班表支持按一个整月对一位或多位人员进行排班。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/attendance-v1/user_daily_shift/batch_create
func (r *AttendanceService) BatchCreateAttendanceUserDailyShift(ctx context.Context, request *BatchCreateAttendanceUserDailyShiftReq, options ...MethodOptionFunc) (*BatchCreateAttendanceUserDailyShiftResp, *Response, error) {
	if r.cli.mock.mockAttendanceBatchCreateAttendanceUserDailyShift != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Attendance#BatchCreateAttendanceUserDailyShift mock enable")
		return r.cli.mock.mockAttendanceBatchCreateAttendanceUserDailyShift(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Attendance",
		API:                   "BatchCreateAttendanceUserDailyShift",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/attendance/v1/user_daily_shifts/batch_create",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(batchCreateAttendanceUserDailyShiftResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockAttendanceBatchCreateAttendanceUserDailyShift mock AttendanceBatchCreateAttendanceUserDailyShift method
func (r *Mock) MockAttendanceBatchCreateAttendanceUserDailyShift(f func(ctx context.Context, request *BatchCreateAttendanceUserDailyShiftReq, options ...MethodOptionFunc) (*BatchCreateAttendanceUserDailyShiftResp, *Response, error)) {
	r.mockAttendanceBatchCreateAttendanceUserDailyShift = f
}

// UnMockAttendanceBatchCreateAttendanceUserDailyShift un-mock AttendanceBatchCreateAttendanceUserDailyShift method
func (r *Mock) UnMockAttendanceBatchCreateAttendanceUserDailyShift() {
	r.mockAttendanceBatchCreateAttendanceUserDailyShift = nil
}

// BatchCreateAttendanceUserDailyShiftReq ...
type BatchCreateAttendanceUserDailyShiftReq struct {
	EmployeeType    EmployeeType                                            `query:"employee_type" json:"-"`     // 请求体和响应体中的 user_id 的员工工号类型, 示例值: "employee_id", 可选值有: employee_id: 员工 employee ID, 即[飞书管理后台](https://bytedance.feishu.cn/admin/contacts/departmentanduser) > 组织架构 > 成员与部门 > 成员详情中的用户 ID, employee_no: 员工工号, 即[飞书管理后台](https://bytedance.feishu.cn/admin/contacts/departmentanduser) > 组织架构 > 成员与部门 > 成员详情中的工号
	UserDailyShifts []*BatchCreateAttendanceUserDailyShiftReqUserDailyShift `json:"user_daily_shifts,omitempty"` // 班表信息列表（数量限制50以内）
	OperatorID      *string                                                 `json:"operator_id,omitempty"`       // 操作人uid, 如果您未操作[考勤管理后台“API 接入”流程](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/attendance-v1/attendance-development-guidelines), 则此字段为必填字段, 示例值: "dd31248a"
}

// BatchCreateAttendanceUserDailyShiftReqUserDailyShift ...
type BatchCreateAttendanceUserDailyShiftReqUserDailyShift struct {
	GroupID string `json:"group_id,omitempty"` // 考勤组 ID, 获取方式: 1）[创建或修改考勤组](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/attendance-v1/group/create) 2）[按名称查询考勤组](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/attendance-v1/group/search) 3）[获取打卡结果](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/attendance-v1/user_task/query), 示例值: "6737202939523236110"
	ShiftID string `json:"shift_id,omitempty"` // 班次 ID, 获取方式: 1）[按名称查询班次](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/attendance-v1/shift/query) 2）[创建班次](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/attendance-v1/shift/create), 示例值: "6753520403404030215"
	Month   int64  `json:"month,omitempty"`    // 月份, 示例值: 202101
	UserID  string `json:"user_id,omitempty"`  // 用户 ID, 示例值: "abd754f7"
	DayNo   int64  `json:"day_no,omitempty"`   // 日期, 示例值: 21
}

// BatchCreateAttendanceUserDailyShiftResp ...
type BatchCreateAttendanceUserDailyShiftResp struct {
	UserDailyShifts []*BatchCreateAttendanceUserDailyShiftRespUserDailyShift `json:"user_daily_shifts,omitempty"` // 班表信息列表
}

// BatchCreateAttendanceUserDailyShiftRespUserDailyShift ...
type BatchCreateAttendanceUserDailyShiftRespUserDailyShift struct {
	GroupID string `json:"group_id,omitempty"` // 考勤组 ID, 获取方式: 1）[创建或修改考勤组](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/attendance-v1/group/create) 2）[按名称查询考勤组](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/attendance-v1/group/search) 3）[获取打卡结果](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/attendance-v1/user_task/query)
	ShiftID string `json:"shift_id,omitempty"` // 班次 ID, 获取方式: 1）[按名称查询班次](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/attendance-v1/shift/query) 2）[创建班次](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/attendance-v1/shift/create)
	Month   int64  `json:"month,omitempty"`    // 月份
	UserID  string `json:"user_id,omitempty"`  // 用户 ID
	DayNo   int64  `json:"day_no,omitempty"`   // 日期
}

// batchCreateAttendanceUserDailyShiftResp ...
type batchCreateAttendanceUserDailyShiftResp struct {
	Code int64                                    `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                                   `json:"msg,omitempty"`  // 错误描述
	Data *BatchCreateAttendanceUserDailyShiftResp `json:"data,omitempty"`
}
