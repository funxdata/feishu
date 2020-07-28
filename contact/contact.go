package contact

import "github.com/funxdata/feishu/core"

const (
	uriGetScope = "https://open.feishu.cn/open-apis/contact/v1/scope/get"
)

// FeishuContact 飞书通讯录
type FeishuContact struct {
	*core.Context
}

// GetScope .
func (f *FeishuContact) GetScope() (map[string][]string, error) {
	tkn, err := f.GetInternalAppAccessToken()
	if err != nil {
		return nil, err
	}
	var (
		ret struct {
			core.FeishuResponse
			Data map[string][]string `json:"data"`
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
