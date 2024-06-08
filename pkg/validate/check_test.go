package validate

import (
	"testing"
	"we-backend/pkg/types"
)

func Test_check(t *testing.T) {
	req := types.LoginRequest{
		Email:    "123qq@qq.com",
		Password: "1nidqkbdiovcf2",
	}
	
	err := Check(req)	
	if err != nil {
		t.Fatal(err)
	}

	t.Log(err)
}