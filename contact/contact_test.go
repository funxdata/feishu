package contact

import (
	"os"
	"testing"

	"github.com/funxdata/feishu/core"
	"github.com/stretchr/testify/assert"
)

var (
	testCli *FeishuContact
	skip    = true
)

func init() {
	var (
		testID     = os.Getenv("TEST_FEISHU_APPID")
		testSecret = os.Getenv("TEST_FEISHU_APPSECRET")
	)
	testCli = &FeishuContact{Context: core.New(testID, testSecret, nil)}
	if testID != "" && testSecret != "" {
		skip = false
	}
}

func TestGetScope(t *testing.T) {
	if skip {
		return
	}
	scopes, err := testCli.GetScope()
	assert.Nil(t, err)
	t.Logf("%v", scopes)
}
