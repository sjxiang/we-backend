package validate

import (
	"testing"
	
	"we-backend/internal/types"
)

func Test_validate(t *testing.T) {
	
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