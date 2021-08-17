// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// CreateDepartment 该接口用于向通讯录中创建部门。
//
// 只可在应用的通讯录权限范围内的部门下创建部门。若需要在根部门下创建子部门，则应用通讯录权限范围需要设置为“全部成员”。应用商店应用无权限调用此接口。
// 该接口部分返回字段受到 <b>数据权限控制</b> ，应用要获取对应字段数据需要额外申请数据权限。具体的数据权限与字段的关系请参考[应用权限](https://open.feishu.cn/document/ukTMukTMukTM/uQjN3QjL0YzN04CN2cDN)，或查看每个接口响应体参数列表的字段描述。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/department/create
func (r *ContactService) CreateDepartment(ctx context.Context, request *CreateDepartmentReq, options ...MethodOptionFunc) (*CreateDepartmentResp, *Response, error) {
	if r.cli.mock.mockContactCreateDepartment != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Contact#CreateDepartment mock enable")
		return r.cli.mock.mockContactCreateDepartment(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Contact",
		API:                   "CreateDepartment",
		Method:                "POST",
		URL:                   "https://open.feishu.cn/open-apis/contact/v3/departments",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(createDepartmentResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

func (r *Mock) MockContactCreateDepartment(f func(ctx context.Context, request *CreateDepartmentReq, options ...MethodOptionFunc) (*CreateDepartmentResp, *Response, error)) {
	r.mockContactCreateDepartment = f
}

func (r *Mock) UnMockContactCreateDepartment() {
	r.mockContactCreateDepartment = nil
}

type CreateDepartmentReq struct {
	UserIDType         *IDType                      `query:"user_id_type" json:"-"`         // 用户 ID 类型, 示例值："open_id", 可选值有: `open_id`：用户的 open id, `union_id`：用户的 union id, `user_id`：用户的 user id, 默认值: `open_id`,, 当值为 `user_id`, 字段权限要求: 获取用户 userid
	DepartmentIDType   *DepartmentIDType            `query:"department_id_type" json:"-"`   // 此次调用中使用的部门ID的类型, 示例值："open_department_id", 可选值有: `department_id`：以自定义department_id来标识部门, `open_department_id`：以open_department_id来标识部门
	ClientToken        *string                      `query:"client_token" json:"-"`         // 根据client_token是否一致来判断是否为同一请求, 示例值："473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E"
	Name               string                       `json:"name,omitempty"`                 // 部门名称, 示例值："DemoName", 最小长度：`1` 字符,**字段权限要求（满足任一）**：, 获取部门基础信息, 以应用身份访问通讯录
	I18nName           *CreateDepartmentReqI18nName `json:"i18n_name,omitempty"`            // 国际化的部门名称,**字段权限要求（满足任一）**：, 获取部门基础信息, 以应用身份访问通讯录
	ParentDepartmentID string                       `json:"parent_department_id,omitempty"` // 父部门的ID,* 创建根部门，该参数值为 “0”, 示例值："D067"
	DepartmentID       *string                      `json:"department_id,omitempty"`        // 本部门的自定义部门ID, 示例值："D096", 最大长度：`128` 字符, 正则校验：`^0|[^od][A-Za-z0-9]*`
	LeaderUserID       *string                      `json:"leader_user_id,omitempty"`       // 部门主管用户ID, 示例值："ou_7dab8a3d3cdcc9da365777c7ad535d62"
	Order              *string                      `json:"order,omitempty"`                // 部门的排序，即部门在其同级部门的展示顺序, 示例值："100",**字段权限要求（满足任一）**：, 获取部门组织架构信息, 以应用身份访问通讯录
	UnitIDs            []string                     `json:"unit_ids,omitempty"`             // 部门单位自定义ID列表，当前只支持一个, 示例值：custom_unit_id,**字段权限要求（满足任一）**：, 获取部门组织架构信息, 以应用身份访问通讯录
	CreateGroupChat    *bool                        `json:"create_group_chat,omitempty"`    // 是否创建部门群，默认不创建, 示例值：false
}

type CreateDepartmentReqI18nName struct {
	ZhCn *string `json:"zh_cn,omitempty"` // 部门的中文名, 示例值："Demo名称"
	JaJp *string `json:"ja_jp,omitempty"` // 部门的日文名, 示例值："デモ名"
	EnUs *string `json:"en_us,omitempty"` // 部门的英文名, 示例值："Demo Name"
}

type createDepartmentResp struct {
	Code int64                 `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                `json:"msg,omitempty"`  // 错误描述
	Data *CreateDepartmentResp `json:"data,omitempty"`
}

type CreateDepartmentResp struct {
	Department *CreateDepartmentRespDepartment `json:"department,omitempty"` // 部门信息
}

type CreateDepartmentRespDepartment struct {
	Name               string                                  `json:"name,omitempty"`                 // 部门名称, 最小长度：`1` 字符,**字段权限要求（满足任一）**：, 获取部门基础信息, 以应用身份访问通讯录
	I18nName           *CreateDepartmentRespDepartmentI18nName `json:"i18n_name,omitempty"`            // 国际化的部门名称,**字段权限要求（满足任一）**：, 获取部门基础信息, 以应用身份访问通讯录
	ParentDepartmentID string                                  `json:"parent_department_id,omitempty"` // 父部门的ID,* 创建根部门，该参数值为 “0”,**字段权限要求（满足任一）**：, 获取部门组织架构信息, 以应用身份访问通讯录
	DepartmentID       string                                  `json:"department_id,omitempty"`        // 本部门的自定义部门ID, 最大长度：`128` 字符, 正则校验：`^0|[^od][A-Za-z0-9]*`,**字段权限要求（满足任一）**：, 获取部门基础信息, 以应用身份访问通讯录
	OpenDepartmentID   string                                  `json:"open_department_id,omitempty"`   // 部门的open_id
	LeaderUserID       string                                  `json:"leader_user_id,omitempty"`       // 部门主管用户ID,**字段权限要求（满足任一）**：, 获取部门组织架构信息, 以应用身份访问通讯录
	ChatID             string                                  `json:"chat_id,omitempty"`              // 部门群ID,**字段权限要求（满足任一）**：, 获取部门基础信息, 以应用身份访问通讯录
	Order              string                                  `json:"order,omitempty"`                // 部门的排序，即部门在其同级部门的展示顺序,**字段权限要求（满足任一）**：, 获取部门组织架构信息, 以应用身份访问通讯录
	UnitIDs            []string                                `json:"unit_ids,omitempty"`             // 部门单位自定义ID列表，当前只支持一个,**字段权限要求（满足任一）**：, 获取部门组织架构信息, 以应用身份访问通讯录
	MemberCount        int64                                   `json:"member_count,omitempty"`         // 部门下用户的个数,**字段权限要求（满足任一）**：, 获取部门组织架构信息, 以应用身份访问通讯录
	Status             *CreateDepartmentRespDepartmentStatus   `json:"status,omitempty"`               // 部门状态,**字段权限要求（满足任一）**：, 获取部门基础信息, 以应用身份访问通讯录
}

type CreateDepartmentRespDepartmentI18nName struct {
	ZhCn string `json:"zh_cn,omitempty"` // 部门的中文名
	JaJp string `json:"ja_jp,omitempty"` // 部门的日文名
	EnUs string `json:"en_us,omitempty"` // 部门的英文名
}

type CreateDepartmentRespDepartmentStatus struct {
	IsDeleted bool `json:"is_deleted,omitempty"` // 是否被删除
}
