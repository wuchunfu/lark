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

// CreateWikiSpace 此接口用于创建知识空间
//
// 此接口不支持tenant access token（应用身份访问）
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/wiki-v2/space/create
func (r *DriveService) CreateWikiSpace(ctx context.Context, request *CreateWikiSpaceReq, options ...MethodOptionFunc) (*CreateWikiSpaceResp, *Response, error) {
	if r.cli.mock.mockDriveCreateWikiSpace != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Drive#CreateWikiSpace mock enable")
		return r.cli.mock.mockDriveCreateWikiSpace(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:               "Drive",
		API:                 "CreateWikiSpace",
		Method:              "POST",
		URL:                 r.cli.openBaseURL + "/open-apis/wiki/v2/spaces",
		Body:                request,
		MethodOption:        newMethodOption(options),
		NeedUserAccessToken: true,
	}
	resp := new(createWikiSpaceResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockDriveCreateWikiSpace mock DriveCreateWikiSpace method
func (r *Mock) MockDriveCreateWikiSpace(f func(ctx context.Context, request *CreateWikiSpaceReq, options ...MethodOptionFunc) (*CreateWikiSpaceResp, *Response, error)) {
	r.mockDriveCreateWikiSpace = f
}

// UnMockDriveCreateWikiSpace un-mock DriveCreateWikiSpace method
func (r *Mock) UnMockDriveCreateWikiSpace() {
	r.mockDriveCreateWikiSpace = nil
}

// CreateWikiSpaceReq ...
type CreateWikiSpaceReq struct {
	Name        *string `json:"name,omitempty"`        // 知识空间名称, 示例值: "知识空间"
	Description *string `json:"description,omitempty"` // 知识空间描述, 示例值: "知识空间描述"
}

// CreateWikiSpaceResp ...
type CreateWikiSpaceResp struct {
	Space *CreateWikiSpaceRespSpace `json:"space,omitempty"` // 知识空间
}

// CreateWikiSpaceRespSpace ...
type CreateWikiSpaceRespSpace struct {
	Name        string `json:"name,omitempty"`        // 知识空间名称
	Description string `json:"description,omitempty"` // 知识空间描述
	SpaceID     string `json:"space_id,omitempty"`    // 知识空间id
	SpaceType   string `json:"space_type,omitempty"`  // 表示知识空间类型（团队空间 或 个人空间）, 可选值有: team: 团队空间, person: 个人空间
	Visibility  string `json:"visibility,omitempty"`  // 表示知识空间可见性（公开空间 或 私有空间）, 可选值有: public: 公开空间, private: 私有空间
}

// createWikiSpaceResp ...
type createWikiSpaceResp struct {
	Code int64                `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string               `json:"msg,omitempty"`  // 错误描述
	Data *CreateWikiSpaceResp `json:"data,omitempty"`
}
