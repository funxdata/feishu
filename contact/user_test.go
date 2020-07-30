package contact

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListUsers(t *testing.T) {
	scopes, err := testCli.GetScope()
	assert.Nil(t, err)
	{
		t.Logf("openids: %v", scopes["authed_open_ids"])
		list, err := testCli.ListUsersOpenIDs(scopes["authed_open_ids"]...)
		assert.Nil(t, err)
		t.Logf("%v", list)
		for _, v := range list {
			t.Logf("user_info: %v %v %v", v.Name, v.Email, v.Mobile)
		}
	}
	{
		t.Logf("openids: %v", scopes["authed_employee_ids"])
		list, err := testCli.ListUsersEmpIDs(scopes["authed_employee_ids"]...)
		assert.Nil(t, err)
		t.Logf("%v", list)
	}
}
