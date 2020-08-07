package contact

import (
	"fmt"
	"net/url"

	"github.com/funxdata/feishu/core"
)

const (
	uriListUserDept       = "https://open.feishu.cn/open-apis/contact/v1/department/user/list"
	uriListUserDetailDept = "https://open.feishu.cn/open-apis/contact/v1/department/user/detail/list"
)

// ListDepartmentUsersOption .
type ListDepartmentUsersOption struct {
	OpenDepartmentID string
	DepartmentID     string
	PageToken        string
	PageSize         int
	FetchChild       bool
}

// DepartmentUserList .
type DepartmentUserList struct {
	HasMore   bool              `json:"has_more"`
	PageToken string            `json:"page_token"`
	UserList  []*DepartmentUser `json:"user_list"`
}

// DepartmentUserDetailList .
type DepartmentUserDetailList struct {
	HasMore   bool        `json:"has_more"`
	PageToken string      `json:"page_token"`
	UserList  []*UserInfo `json:"user_list"`
}

// DepartmentUser 用户列表信息
type DepartmentUser struct {
	EmployeeID string `json:"employee_id"`
	OpenID     string `json:"open_id"`
	Name       string `json:"name"`
	EmployeeNo string `json:"employee_no"`
	UnionID    string `json:"union_id"`
}

// ListDepartmentUsers 获取部门用户列表
// Docs. https://open.feishu.cn/document/ukTMukTMukTM/uEzNz4SM3MjLxczM
func (f *FeishuContact) ListDepartmentUsers(opt *ListDepartmentUsersOption) (*DepartmentUserList, error) {
	uq := make(url.Values)
	if opt.DepartmentID != "" {
		uq.Set("department_id", opt.DepartmentID)
	} else if opt.OpenDepartmentID != "" {
		uq.Set("open_department_id", opt.OpenDepartmentID)
	} else {
		return nil, fmt.Errorf("invalid ListDepartmentUsersOption.")
	}
	if opt.PageSize > 0 {
		uq.Set("page_size", fmt.Sprint(opt.PageSize))
	} else {
		uq.Set("page_size", "10")
	}
	if opt.FetchChild {
		uq.Set("fetch_child", "true")
	} else {
		uq.Set("fetch_child", "false")
	}
	if opt.PageToken != "" {
		uq.Set("page_token", opt.PageToken)
	}
	tkn, err := f.GetInternalAppAccessToken()
	if err != nil {
		return nil, err
	}

	var ret struct {
		core.FeishuResponse
		Data *DepartmentUserList `json:"data"`
	}
	err = f.Get(uriListUserDept, tkn, uq, &ret)
	if err != nil {
		return nil, err
	}

	if err := ret.Err(); err != nil {
		return nil, err
	}

	return ret.Data, nil
}

// ListDepartmentUserDetails 获取部门用户列表
// Docs. https://open.feishu.cn/document/ukTMukTMukTM/uYzN3QjL2czN04iN3cDN
func (f *FeishuContact) ListDepartmentUserDetails(opt *ListDepartmentUsersOption) (*DepartmentUserDetailList, error) {
	uq := make(url.Values)
	if opt.DepartmentID != "" {
		uq.Set("department_id", opt.DepartmentID)
	} else if opt.OpenDepartmentID != "" {
		uq.Set("open_department_id", opt.OpenDepartmentID)
	} else {
		return nil, fmt.Errorf("invalid ListDepartmentUsersOption.")
	}
	if opt.PageSize > 0 {
		uq.Set("page_size", fmt.Sprint(opt.PageSize))
	} else {
		uq.Set("page_size", "10")
	}
	if opt.FetchChild {
		uq.Set("fetch_child", "true")
	} else {
		uq.Set("fetch_child", "false")
	}
	if opt.PageToken != "" {
		uq.Set("page_token", opt.PageToken)
	}
	tkn, err := f.GetInternalAppAccessToken()
	if err != nil {
		return nil, err
	}

	var ret struct {
		core.FeishuResponse
		Data *DepartmentUserDetailList `json:"data"`
	}
	err = f.Get(uriListUserDetailDept, tkn, uq, &ret)
	if err != nil {
		return nil, err
	}

	if err := ret.Err(); err != nil {
		return nil, err
	}

	return ret.Data, nil
}
