package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	testCli *Context
	skip    = true
)

func init() {
	var (
		// testID     = os.Getenv("TEST_FEISHU_APPID")
		// testSecret = os.Getenv("TEST_FEISHU_APPSECRET")
		testID     = "cli_9f97971cd7a9d00c"
		testSecret = "h3hw9uBnKuK37kY9sGwir1oGDCxdmPFq"
	)
	testCli = New(testID, testSecret, nil)
	if testID != "" && testSecret != "" {
		skip = false
	}
}

func TestAuthen(t *testing.T) {
	at, err := testCli.GetInternalAppAccessToken()
	assert.Nil(t, err)

	t.Logf("%v", at)
}
