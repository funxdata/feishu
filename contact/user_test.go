package contact

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
