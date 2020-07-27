package core

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	testCli *Context
	skip    = true
)

func init() {
	var (
		testID     = os.Getenv("TEST_FEISHU_APPID")
		testSecret = os.Getenv("TEST_FEISHU_APPSECRET")
	)
	testCli = New(testID, testSecret, nil)
	if testID != "" && testSecret != "" {
		skip = false
	}
}

func TestAuthen(t *testing.T) {
	if skip {
		return
	}
	{
		at, err := testCli.GetInternalAppAccessToken()
		assert.Nil(t, err)
		t.Logf("%v", at)
	}
	{
		at, err := testCli.GetInternalTenantAccessToken()
		assert.Nil(t, err)
		t.Logf("%v", at)
	}

}
