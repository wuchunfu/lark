// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// CreateUser 使用该接口向通讯录创建一个用户，可以理解为员工入职。创建部门后只返回有数据权限的数据。具体的数据权限的与字段的对应关系请参照[应用权限](https://open.feishu.cn/document/ukTMukTMukTM/uQjN3QjL0YzN04CN2cDN)
//
// 新增用户的所有部门必须都在当前应用的通讯录授权范围内才允许新增用户，如果想要在根部门下新增用户，必须要有全员权限。 应用商店应用无权限调用此接口。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/user/create
func (r *ContactService) CreateUser(ctx context.Context, request *CreateUserReq, options ...MethodOptionFunc) (*CreateUserResp, *Response, error) {
	if r.cli.mock.mockContactCreateUser != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Contact#CreateUser mock enable")
		return r.cli.mock.mockContactCreateUser(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Contact",
		API:                   "CreateUser",
		Method:                "POST",
		URL:                   "https://open.feishu.cn/open-apis/contact/v3/users",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(createUserResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

func (r *Mock) MockContactCreateUser(f func(ctx context.Context, request *CreateUserReq, options ...MethodOptionFunc) (*CreateUserResp, *Response, error)) {
	r.mockContactCreateUser = f
}

func (r *Mock) UnMockContactCreateUser() {
	r.mockContactCreateUser = nil
}

type CreateUserReq struct {
	UserIDType           *IDType                          `query:"user_id_type" json:"-"`           // 用户 ID 类型, 示例值："open_id", 可选值有: `open_id`：用户的 open id, `union_id`：用户的 union id, `user_id`：用户的 user id, 默认值: `open_id`,, 当值为 `user_id`, 字段权限要求: 获取用户 userid
	DepartmentIDType     *DepartmentIDType                `query:"department_id_type" json:"-"`     // 此次调用中使用的部门ID的类型, 示例值："open_department_id", 可选值有: `department_id`：以自定义department_id来标识部门, `open_department_id`：以open_department_id来标识部门, 默认值: `open_department_id`
	ClientToken          *string                          `query:"client_token" json:"-"`           // 根据client_token是否一致来判断是否为同一请求, 示例值："xxxx-xxxxx-xxx"
	UserID               *string                          `json:"user_id,omitempty"`                // 租户内用户的唯一标识, 示例值："ou_7dab8a3d3cdcc9da365777c7ad535d62", 字段权限要求:  获取用户 user ID
	Name                 string                           `json:"name,omitempty"`                   // 用户名, 示例值："张三", 最小长度：`1` 字符,**字段权限要求（满足任一）**：, 获取用户基本信息, 以应用身份访问通讯录（历史版本）
	EnName               *string                          `json:"en_name,omitempty"`                // 英文名, 示例值："San Zhang",**字段权限要求（满足任一）**：, 获取用户基本信息, 以应用身份访问通讯录（历史版本）
	Email                *string                          `json:"email,omitempty"`                  // 邮箱, 示例值："zhangsan@gmail.com", 字段权限要求:  获取用户邮箱信息
	Mobile               string                           `json:"mobile,omitempty"`                 // 手机号, 示例值："13011111111", 字段权限要求:  获取用户手机号
	MobileVisible        *bool                            `json:"mobile_visible,omitempty"`         // 手机号码可见性，true 为可见，false 为不可见，目前默认为 true。不可见时，组织员工将无法查看该员工的手机号码, 示例值：false
	Gender               *int64                           `json:"gender,omitempty"`                 // 性别, 示例值：1, 可选值有: `0`：保密, `1`：男, `2`：女,**字段权限要求（满足任一）**：, 获取用户性别, 以应用身份访问通讯录（历史版本）
	AvatarKey            *string                          `json:"avatar_key,omitempty"`             // 头像的文件Key，可通过“消息与群组/消息/图片信息”中的“上传图片”接口上传并获取头像文件 Key, 示例值："2500c7a9-5fff-4d9a-a2de-3d59614ae28g"
	DepartmentIDs        []string                         `json:"department_ids,omitempty"`         // 用户所属部门的ID列表, 示例值：od-4e6ac4d14bcd5071a37a39de902c7141
	LeaderUserID         *string                          `json:"leader_user_id,omitempty"`         // 用户的直接主管的用户ID, 示例值："ou_7dab8a3d3cdcc9da365777c7ad535d62"
	City                 *string                          `json:"city,omitempty"`                   // 城市, 示例值："杭州",**字段权限要求（满足任一）**：, 获取用户雇佣信息, 以应用身份访问通讯录（历史版本）
	Country              *string                          `json:"country,omitempty"`                // 国家或地区, 示例值："中国",**字段权限要求（满足任一）**：, 获取用户雇佣信息, 以应用身份访问通讯录（历史版本）
	WorkStation          *string                          `json:"work_station,omitempty"`           // 工位, 示例值："北楼-H34",**字段权限要求（满足任一）**：, 获取用户雇佣信息, 以应用身份访问通讯录（历史版本）
	JoinTime             *int64                           `json:"join_time,omitempty"`              // 入职时间, 示例值：2147483647,**字段权限要求（满足任一）**：, 获取用户雇佣信息, 以应用身份访问通讯录（历史版本）
	EmployeeNo           *string                          `json:"employee_no,omitempty"`            // 工号, 示例值："1",**字段权限要求（满足任一）**：, 获取用户雇佣信息, 以应用身份访问通讯录（历史版本）
	EmployeeType         int64                            `json:"employee_type,omitempty"`          // 员工类型，可选值有：, 1：正式员工, 2：实习生, 3：外包, 4：劳务, 5：顾问   ,同时可读取到自定义员工类型的 int 值，可通过下方接口获取到该租户的自定义员工类型的名称   ,[获取人员类型](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/employee_type_enum/list), 示例值：1
	Orders               []*CreateUserReqOrder            `json:"orders,omitempty"`                 // 用户排序信息
	CustomAttrs          []*CreateUserReqCustomAttr       `json:"custom_attrs,omitempty"`           // 自定义字段，请确保你的组织管理员已在管理后台/组织架构/成员字段管理/自定义字段管理/全局设置中开启了“允许开放平台 API 调用“，否则该字段不会生效/返回。,**字段权限要求（满足任一）**：, 获取用户雇佣信息, 以应用身份访问通讯录（历史版本）
	EnterpriseEmail      *string                          `json:"enterprise_email,omitempty"`       // 企业邮箱，请先确保已在管理后台启用飞书邮箱服务, 示例值："demo@mail.com",**字段权限要求（满足任一）**：, 获取用户雇佣信息, 以应用身份访问通讯录（历史版本）
	JobTitle             *string                          `json:"job_title,omitempty"`              // 职务, 示例值："xxxxx",**字段权限要求（满足任一）**：, 获取用户雇佣信息, 以应用身份访问通讯录（历史版本）
	NeedSendNotification *bool                            `json:"need_send_notification,omitempty"` // 是否发送提示消息, 示例值：false
	NotificationOption   *CreateUserReqNotificationOption `json:"notification_option,omitempty"`    // 创建用户的邀请方式
}

type CreateUserReqOrder struct {
	DepartmentID    *string `json:"department_id,omitempty"`    // 排序信息对应的部门ID, 示例值："od-4e6ac4d14bcd5071a37a39de902c7141"
	UserOrder       *int64  `json:"user_order,omitempty"`       // 用户在其直属部门内的排序，数值越大，排序越靠前, 示例值：100
	DepartmentOrder *int64  `json:"department_order,omitempty"` // 用户所属的多个部门间的排序，数值越大，排序越靠前, 示例值：100
}

type CreateUserReqCustomAttr struct {
	Type  *string                       `json:"type,omitempty"`  // 自定义字段类型   , `TEXT`, `HREF`, `ENUMERATION`, `PICTURE_ENUM`, `GENERIC_USER`, 示例值："TEXT"
	ID    *string                       `json:"id,omitempty"`    // 自定义字段ID, 示例值："DemoId"
	Value *CreateUserReqCustomAttrValue `json:"value,omitempty"` // 自定义字段取值
}

type CreateUserReqCustomAttrValue struct {
	Text        *string                                  `json:"text,omitempty"`         // 字段类型为 TEXT 时该参数定义字段值，字段类型为 HREF 时该参数定义网页标题, 示例值："DemoText"
	URL         *string                                  `json:"url,omitempty"`          // 字段类型为 HREF 时，该参数定义默认 URL, 示例值："http://www.feishu.cn"
	PcURL       *string                                  `json:"pc_url,omitempty"`       // 字段类型为 HREF 时，该参数定义PC端 URL, 示例值："http://www.feishu.cn"
	OptionID    *string                                  `json:"option_id,omitempty"`    // 字段类型为 ENUMERATION 或 PICTURE_ENUM 时，该参数定义选项值, 示例值："edcvfrtg"
	GenericUser *CreateUserReqCustomAttrValueGenericUser `json:"generic_user,omitempty"` // 字段类型为 GENERIC_USER 时，该参数定义引用人员
}

type CreateUserReqCustomAttrValueGenericUser struct {
	ID   string `json:"id,omitempty"`   // 用户的user_id, 示例值："9b2fabg5"
	Type int64  `json:"type,omitempty"` // 用户类型    1：用户, 示例值：1
}

type CreateUserReqNotificationOption struct {
	Channels []string `json:"channels,omitempty"` // 通道列表，枚举值：,sms（短信邀请），email（邮件邀请）, 示例值：["sms", "email"]
	Language *string  `json:"language,omitempty"` // 语言类型, 示例值："zh-CN", 可选值有: `zh-CN`：中文, `en-US`：英文, `ja-JP`：日文
}

type createUserResp struct {
	Code int64           `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string          `json:"msg,omitempty"`  // 错误描述
	Data *CreateUserResp `json:"data,omitempty"`
}

type CreateUserResp struct {
	User *CreateUserRespUser `json:"user,omitempty"` // 用户信息
}

type CreateUserRespUser struct {
	UnionID              string                                `json:"union_id,omitempty"`               // 用户的union_id
	UserID               string                                `json:"user_id,omitempty"`                // 租户内用户的唯一标识, 字段权限要求:  获取用户 user ID
	OpenID               string                                `json:"open_id,omitempty"`                // 用户的open_id
	Name                 string                                `json:"name,omitempty"`                   // 用户名, 最小长度：`1` 字符,**字段权限要求（满足任一）**：, 获取用户基本信息, 以应用身份访问通讯录（历史版本）
	EnName               string                                `json:"en_name,omitempty"`                // 英文名,**字段权限要求（满足任一）**：, 获取用户基本信息, 以应用身份访问通讯录（历史版本）
	Email                string                                `json:"email,omitempty"`                  // 邮箱, 字段权限要求:  获取用户邮箱信息
	Mobile               string                                `json:"mobile,omitempty"`                 // 手机号, 字段权限要求:  获取用户手机号
	MobileVisible        bool                                  `json:"mobile_visible,omitempty"`         // 手机号码可见性，true 为可见，false 为不可见，目前默认为 true。不可见时，组织员工将无法查看该员工的手机号码
	Gender               int64                                 `json:"gender,omitempty"`                 // 性别, 可选值有: `0`：保密, `1`：男, `2`：女,**字段权限要求（满足任一）**：, 获取用户性别, 以应用身份访问通讯录（历史版本）
	AvatarKey            string                                `json:"avatar_key,omitempty"`             // 头像的文件Key，可通过“消息与群组/消息/图片信息”中的“上传图片”接口上传并获取头像文件 Key
	Avatar               *CreateUserRespUserAvatar             `json:"avatar,omitempty"`                 // 用户头像信息,**字段权限要求（满足任一）**：, 获取用户基本信息, 以应用身份访问通讯录（历史版本）
	Status               *CreateUserRespUserStatus             `json:"status,omitempty"`                 // 用户状态,**字段权限要求（满足任一）**：, 获取用户雇佣信息, 以应用身份访问通讯录（历史版本）
	DepartmentIDs        []string                              `json:"department_ids,omitempty"`         // 用户所属部门的ID列表,**字段权限要求（满足任一）**：, 获取用户组织架构信息, 以应用身份访问通讯录（历史版本）
	LeaderUserID         string                                `json:"leader_user_id,omitempty"`         // 用户的直接主管的用户ID,**字段权限要求（满足任一）**：, 获取用户组织架构信息, 以应用身份访问通讯录（历史版本）
	City                 string                                `json:"city,omitempty"`                   // 城市,**字段权限要求（满足任一）**：, 获取用户雇佣信息, 以应用身份访问通讯录（历史版本）
	Country              string                                `json:"country,omitempty"`                // 国家或地区,**字段权限要求（满足任一）**：, 获取用户雇佣信息, 以应用身份访问通讯录（历史版本）
	WorkStation          string                                `json:"work_station,omitempty"`           // 工位,**字段权限要求（满足任一）**：, 获取用户雇佣信息, 以应用身份访问通讯录（历史版本）
	JoinTime             int64                                 `json:"join_time,omitempty"`              // 入职时间,**字段权限要求（满足任一）**：, 获取用户雇佣信息, 以应用身份访问通讯录（历史版本）
	IsTenantManager      bool                                  `json:"is_tenant_manager,omitempty"`      // 是否是租户管理员,**字段权限要求（满足任一）**：, 获取用户雇佣信息, 以应用身份访问通讯录（历史版本）
	EmployeeNo           string                                `json:"employee_no,omitempty"`            // 工号,**字段权限要求（满足任一）**：, 获取用户雇佣信息, 以应用身份访问通讯录（历史版本）
	EmployeeType         int64                                 `json:"employee_type,omitempty"`          // 员工类型，可选值有：, 1：正式员工, 2：实习生, 3：外包, 4：劳务, 5：顾问   ,同时可读取到自定义员工类型的 int 值，可通过下方接口获取到该租户的自定义员工类型的名称   ,[获取人员类型](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/employee_type_enum/list),**字段权限要求（满足任一）**：, 获取用户雇佣信息, 以应用身份访问通讯录（历史版本）
	Orders               []*CreateUserRespUserOrder            `json:"orders,omitempty"`                 // 用户排序信息,**字段权限要求（满足任一）**：, 获取用户组织架构信息, 以应用身份访问通讯录（历史版本）
	CustomAttrs          []*CreateUserRespUserCustomAttr       `json:"custom_attrs,omitempty"`           // 自定义字段，请确保你的组织管理员已在管理后台/组织架构/成员字段管理/自定义字段管理/全局设置中开启了“允许开放平台 API 调用“，否则该字段不会生效/返回。,**字段权限要求（满足任一）**：, 获取用户雇佣信息, 以应用身份访问通讯录（历史版本）
	EnterpriseEmail      string                                `json:"enterprise_email,omitempty"`       // 企业邮箱，请先确保已在管理后台启用飞书邮箱服务,**字段权限要求（满足任一）**：, 获取用户雇佣信息, 以应用身份访问通讯录（历史版本）
	JobTitle             string                                `json:"job_title,omitempty"`              // 职务,**字段权限要求（满足任一）**：, 获取用户雇佣信息, 以应用身份访问通讯录（历史版本）
	NeedSendNotification bool                                  `json:"need_send_notification,omitempty"` // 是否发送提示消息
	NotificationOption   *CreateUserRespUserNotificationOption `json:"notification_option,omitempty"`    // 创建用户的邀请方式
	IsFrozen             bool                                  `json:"is_frozen,omitempty"`              // 是否暂停用户
}

type CreateUserRespUserAvatar struct {
	Avatar72     string `json:"avatar_72,omitempty"`     // 72*72像素头像链接
	Avatar240    string `json:"avatar_240,omitempty"`    // 240*240像素头像链接
	Avatar640    string `json:"avatar_640,omitempty"`    // 640*640像素头像链接
	AvatarOrigin string `json:"avatar_origin,omitempty"` // 原始头像链接
}

type CreateUserRespUserStatus struct {
	IsFrozen    bool `json:"is_frozen,omitempty"`    // 是否暂停
	IsResigned  bool `json:"is_resigned,omitempty"`  // 是否离职
	IsActivated bool `json:"is_activated,omitempty"` // 是否激活
}

type CreateUserRespUserOrder struct {
	DepartmentID    string `json:"department_id,omitempty"`    // 排序信息对应的部门ID
	UserOrder       int64  `json:"user_order,omitempty"`       // 用户在其直属部门内的排序，数值越大，排序越靠前
	DepartmentOrder int64  `json:"department_order,omitempty"` // 用户所属的多个部门间的排序，数值越大，排序越靠前
}

type CreateUserRespUserCustomAttr struct {
	Type  string                             `json:"type,omitempty"`  // 自定义字段类型   , `TEXT`, `HREF`, `ENUMERATION`, `PICTURE_ENUM`, `GENERIC_USER`
	ID    string                             `json:"id,omitempty"`    // 自定义字段ID
	Value *CreateUserRespUserCustomAttrValue `json:"value,omitempty"` // 自定义字段取值
}

type CreateUserRespUserCustomAttrValue struct {
	Text        string                                        `json:"text,omitempty"`         // 字段类型为 TEXT 时该参数定义字段值，字段类型为 HREF 时该参数定义网页标题
	URL         string                                        `json:"url,omitempty"`          // 字段类型为 HREF 时，该参数定义默认 URL
	PcURL       string                                        `json:"pc_url,omitempty"`       // 字段类型为 HREF 时，该参数定义PC端 URL
	OptionID    string                                        `json:"option_id,omitempty"`    // 字段类型为 ENUMERATION 或 PICTURE_ENUM 时，该参数定义选项值
	OptionValue string                                        `json:"option_value,omitempty"` // 选项值
	Name        string                                        `json:"name,omitempty"`         // 名称
	PictureURL  string                                        `json:"picture_url,omitempty"`  // 图片链接
	GenericUser *CreateUserRespUserCustomAttrValueGenericUser `json:"generic_user,omitempty"` // 字段类型为 GENERIC_USER 时，该参数定义引用人员
}

type CreateUserRespUserCustomAttrValueGenericUser struct {
	ID   string `json:"id,omitempty"`   // 用户的user_id
	Type int64  `json:"type,omitempty"` // 用户类型    1：用户
}

type CreateUserRespUserNotificationOption struct {
	Channels []string `json:"channels,omitempty"` // 通道列表，枚举值：,sms（短信邀请），email（邮件邀请）
	Language string   `json:"language,omitempty"` // 语言类型, 可选值有: `zh-CN`：中文, `en-US`：英文, `ja-JP`：日文
}
