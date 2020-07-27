package core

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	cachepkg "github.com/funxdata/feishu/cache"
	"github.com/sirupsen/logrus"
)

// Context 需要用到的上下文资源
type Context struct {
	httpCli *http.Client
	cache   cachepkg.Cache

	AppID     string
	AppSecret string
}

// New .
func New(appID, appSecret string, cache cachepkg.Cache) *Context {
	c := &Context{
		AppID:     appID,
		AppSecret: appSecret,
		httpCli: &http.Client{
			Timeout: time.Second * 15,
		},
		cache: cache,
	}
	if c.cache == nil {
		c.cache = cachepkg.NewMemory()
	}
	return c
}

// Post .
func (c *Context) Post(url string, body, ret interface{}) error {
	buf := new(bytes.Buffer)
	err := json.NewEncoder(buf).Encode(body)
	if err != nil {
		return err
	}

	resp, err := c.httpCli.Post(url, "application/json", buf)
	if err != nil {
		return err
	}

	err = json.NewDecoder(resp.Body).Decode(ret)
	if err != nil {
		logrus.Errorf("decode body to %T failed, %s", ret, err)
		return err
	}
	return nil
}

// FeishuResponse 统一的错误代码
type FeishuResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// Err 是否返回错误
func (f FeishuResponse) Err(action ...string) error {
	if f.Code == 0 {
		return nil
	}
	if len(action) > 0 {
		return fmt.Errorf("%s failed, (%v) %s", strings.Join(action, " "), f.Code, f.Msg)
	}
	return fmt.Errorf("(%v) %s", f.Code, f.Msg)
}
