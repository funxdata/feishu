package core

import "time"

const (
	cacheKeyAppAccessToken    = "feishu_app_access_token_"
	cacheKeyTenantAccessToken = "feishu_tenant_access_token_"

	uriGetInternalAppAccessToken    = "https://open.feishu.cn/open-apis/auth/v3/app_access_token/internal/"
	uriGetInternalTenantAccessToken = "https://open.feishu.cn/open-apis/auth/v3/tenant_access_token/internal/"
)

// GetInternalAppAccessToken 获取 app_access_token（企业自建应用）
func (c *Context) GetInternalAppAccessToken() (string, error) {
	key := cacheKeyAppAccessToken + c.AppID
	if str := c.cache.Get(key); str != nil {
		return str.(string), nil
	}
	var (
		reqBody = map[string]string{
			"app_id":     c.AppID,
			"app_secret": c.AppSecret,
		}
		ret struct {
			FeishuResponse
			AppAccessToken string `json:"app_access_token"`
			// Expire 过期时间，单位为秒（两小时失效）
			Expire int64 `json:"expire"`
		}
	)
	err := c.Post(uriGetInternalAppAccessToken, reqBody, &ret)
	if err != nil {
		return "", err
	}
	if err := ret.Err(); err != nil {
		return "", err
	}
	c.cache.Set(key, ret.AppAccessToken, time.Second*(time.Duration(ret.Expire/3)))

	return ret.AppAccessToken, nil
}

// GetInternalTenantAccessToken 获取 tenant_access_token（企业自建应用）
func (c *Context) GetInternalTenantAccessToken() (string, error) {
	key := cacheKeyTenantAccessToken + c.AppID
	if str := c.cache.Get(key); str != nil {
		return str.(string), nil
	}
	var (
		reqBody = map[string]string{
			"app_id":     c.AppID,
			"app_secret": c.AppSecret,
		}
		ret struct {
			FeishuResponse
			TenantAccessToken string `json:"tenant_access_token"`
			// Expire 过期时间，单位为秒（两小时失效）
			Expire int64 `json:"expire"`
		}
	)
	err := c.Post(uriGetInternalTenantAccessToken, reqBody, &ret)
	if err != nil {
		return "", err
	}
	if err := ret.Err(); err != nil {
		return "", err
	}
	c.cache.Set(key, ret.TenantAccessToken, time.Second*(time.Duration(ret.Expire/3)))

	return ret.TenantAccessToken, nil
}
