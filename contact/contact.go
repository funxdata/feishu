package contact

import "github.com/funxdata/feishu/core"

const (
	uriGetScope = "https://open.feishu.cn/open-apis/contact/v1/scope/get"
)

// FeishuContact 飞书通讯录
type FeishuContact struct {
	*core.Context
}

// Scope 授权范围
type Scope struct {
	// Departments 已授权部门自定义 ID 列表，授权范围为全员可见时返回的是当前企业的所有一级部门列表
	Departments []string `json:"authed_departments"`
	// OpenDepartments 已授权部门 openID 列表，授权范围为全员可见时返回的是当前企业的所有一级部门列表
	OpenDepartments []string `json:"authed_open_departments"`
	// EmployeeIDs 已授权用户 employee_id 列表，应用申请了 获取用户user_id 权限时返回；当授权范围为全员可见时返回的是当前企业所有顶级部门用户列表
	EmployeeIDs []string `json:"authed_employee_ids"`
	// OpenIDs 已授权用户 open_id 列表；当授权范围为全员可见时返回的是当前企业所有顶级部门用户列表
	OpenIDs []string `json:"authed_open_ids"`
}

// GetScope 获取通讯录授权范围
// Docs https://open.feishu.cn/document/ukTMukTMukTM/ugjNz4CO2MjL4YzM
func (f *FeishuContact) GetScope() (*Scope, error) {
	tkn, err := f.GetInternalAppAccessToken()
	if err != nil {
		return nil, err
	}
	var (
		ret struct {
			core.FeishuResponse
			Data *Scope `json:"data"`
		}
	)
	if err := f.Post(uriGetScope, tkn, nil, &ret); err != nil {
		return nil, err
	}
	if err := ret.Err(); err != nil {
		return nil, err
	}

	return ret.Data, nil
}
