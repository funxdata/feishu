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

func TestListUsers(t *testing.T) {
	if skip {
		return
	}
	scopes, err := testCli.GetScope()
	assert.Nil(t, err)
	{
		t.Logf("openids: %v", scopes.OpenIDs)
		list, err := testCli.ListUsersOpenIDs(scopes.OpenIDs...)
		assert.Nil(t, err)
		t.Logf("%v", list)
		for _, v := range list {
			t.Logf("user_info: %v %v %v", v.Name, v.Email, v.Mobile)
		}
	}
	{
		t.Logf("openids: %v", scopes.EmployeeIDs)
		list, err := testCli.ListUsersEmpIDs(scopes.EmployeeIDs...)
		assert.Nil(t, err)
		t.Logf("%v", list)
	}
}

func TestListDepartmentUsers(t *testing.T) {
	if skip {
		return
	}
	{
		opt := &ListDepartmentUsersOption{
			FetchChild: true,
		}
		t.Logf("openids: %v", opt)
		list, err := testCli.ListDepartmentUsers(opt)
		assert.Nil(t, err)
		t.Logf("%v", list)
		for _, v := range list.UserList {
			t.Logf("user_info: %v", v)
		}
	}
}
