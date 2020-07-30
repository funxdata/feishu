package contact

import (
	"net/url"

	"github.com/funxdata/feishu/core"
)

const (
	uriBatchGetUser = "https://open.feishu.cn/open-apis/contact/v1/user/batch_get"
	uriListUserDept = "https://open.feishu.cn/open-apis/contact/v1/department/user/list"
)

// Value .
type Value struct {
	Value string `json:"value"`
}

// UserInfo 用户信息
type UserInfo struct {
	// Name 用户名
	Name string `json:"name"`
	// NamePy 用户名拼音
	NamePy string `json:"name_py"`
	// EnName 英文名
	EnName string `json:"en_name"`
	// EmployeeID 用户的 employee_id，申请了"获取用户 user_id"权限的应用返回该字段
	EmployeeID string `json:"employee_id"`
	// EmployeeNo 工号
	EmployeeNo string `json:"employee_no"`
	// OpenID 用户的 open_id
	OpenID string `json:"open_id"`
	// UnionID 用户统一ID，申请了"获取用户统一ID"权限后返回
	UnionID string `json:"union_id"`
	// Status 用户状态，bit0(最低位): 1冻结，0未冻结；bit1:1离职，0在职；bit2:1未激活，0已激活
	Status int `json:"status"`
	// EmployeeType 员工类型。1:正式员工；2:实习生；3:外包；4:劳务；5:顾问
	EmployeeType int `json:"employee_type"`
	// Avatar72 用户头像，72*72px
	Avatar72 string `json:"avatar_72"`
	// Avatar240 用户头像，240*240px
	Avatar240 string `json:"avatar_240"`
	// Avatar640 用户头像，640*640px
	Avatar640 string `json:"avatar_640"`
	// AvatarURL 用户头像，原始大小
	AvatarURL string `json:"avatar_url"`
	// Gender 性别，未设置不返回该字段。1:男；2:女
	Gender int `json:"gender"`
	// Email 用户邮箱地址，已申请"获取用户邮箱"权限返回该字段
	Email string `json:"email"`
	// Mobile 用户手机号，已申请"获取用户手机号"权限的企业自建应用返回该字段
	Mobile string `json:"mobile"`
	// Description 用户个人签名
	Description string `json:"description"`
	// Country 用户所在国家
	Country string `json:"country"`
	// City 用户所在城市
	City string `json:"city"`
	// WorkStation 工位
	WorkStation string `json:"work_station"`
	// IsTenantManager 是否是企业超级管理员
	IsTenantManager bool `json:"is_tenant_manager"`
	// JoinTime 入职时间，未设置不返回该字段
	JoinTime int64 `json:"join_time"`
	// UpdateTime 更新时间
	UpdateTime int64 `json:"update_time"`
	// LeaderEmployeeID 用户直接领导的 employee_id，企业自建应用返回，应用商店应用申请了 employee_id 权限时才返回
	LeaderEmployeeID string `json:"leader_employee_id"`
	// LeaderOpenID 用户直接领导的 open_id
	LeaderOpenID string `json:"leader_open_id"`
	// LeaderUnionID 用户直接领导的 union_id,申请了"获取用户统一ID"权限后返回
	LeaderUnionID string `json:"leader_union_id"`

	// Departments 用户所在部门自定义 ID列表，用户可能同时存在于多个部门
	Departments []string `json:"departments"`
	// OpenDepartments 用户所在部门 openID 列表，用户可能同时存在于多个部门
	OpenDepartments []string `json:"open_departments"`
	// CustomAttrs 用户的自定义属性信息。
	// 该字段返回的每一个属性包括自定义属性 ID 和自定义属性值。
	// 企业开放了自定义用户属性且为该用户设置了自定义属性的值，才会返回该字段
	CustomAttrs map[string]Value `json:"custom_attrs"`
}

// ListUsersEmpIDs 批量获取用户信息
// Docs. https://open.feishu.cn/document/ukTMukTMukTM/uIzNz4iM3MjLyczM
func (f *FeishuContact) ListUsersEmpIDs(empIDs ...string) ([]*UserInfo, error) {
	return f.listUsers("employee_ids", empIDs)
}

// ListUsersOpenIDs 批量获取用户信息
// Docs. https://open.feishu.cn/document/ukTMukTMukTM/uIzNz4iM3MjLyczM
func (f *FeishuContact) ListUsersOpenIDs(openIDs ...string) ([]*UserInfo, error) {
	return f.listUsers("open_ids", openIDs)
}

func (f *FeishuContact) listUsers(idType string, ids []string) ([]*UserInfo, error) {
	if len(ids) == 0 {
		return []*UserInfo{}, nil
	}
	tkn, err := f.GetInternalAppAccessToken()
	if err != nil {
		return nil, err
	}

	uq := make(url.Values)
	for _, id := range ids {
		uq.Add(idType, id)
	}

	var ret struct {
		core.FeishuResponse
		Data struct {
			UserInfos []*UserInfo `json:"user_infos"`
		} `json:"data"`
	}
	err = f.Get(uriBatchGetUser, tkn, uq, &ret)
	if err != nil {
		return nil, err
	}

	if err := ret.Err(); err != nil {
		return nil, err
	}

	return ret.Data.UserInfos, nil
}
