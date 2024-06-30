package types

import (
	"testing"
	"we-backend/pkg/we"

	"github.com/stretchr/testify/require"
)

func Test_validate_password(t *testing.T) {
	
	if minSize, digit, special, letter := ValidatePassword("hasicoghwif*4YY"); !minSize || !digit || !special || !letter {
		t.Fatal("这个密码太弱了，不少于8个字符，必须包含英文字母、阿拉伯数字以及特殊符号")
	}

	t.Log("pass")
}


func TestRegisterInput_Sanitize(t *testing.T) {
	input := RegisterInput{
		Username:        " bob ",
		Email:           " BOB@gmail.com  ",
		Password:        "password",
		ConfirmPassword: "password",
	}

	want := RegisterInput{
		Username:        "bob",
		Email:           "bob@gmail.com",
		Password:        "password",
		ConfirmPassword: "password",
	}

	input.Sanitize()

	require.Equal(t, want, input)
}

func TestRegisterInput_Validate(t *testing.T) {
	testCases := []struct {
		name  string
		input RegisterInput
		err   error
	}{
		{
			name: "valid",
			input: RegisterInput{
				Username:        "bob",
				Email:           "bob@gmail.com",
				Password:        "password",
				ConfirmPassword: "password",
			},
			err: nil,
		},
		{
			name: "invalid email",
			input: RegisterInput{
				Username:        "bob",
				Email:           "bob",
				Password:        "password",
				ConfirmPassword: "password",
			},
			err: we.ErrInvalidParameter,
		},
		{
			name: "too short username",
			input: RegisterInput{
				Username:        "b",
				Email:           "bob@gmail.com",
				Password:        "password",
				ConfirmPassword: "password",
			},
			err: we.ErrInvalidParameter,
		},
		{
			name: "too short password",
			input: RegisterInput{
				Username:        "bob",
				Email:           "bob@gmail.com",
				Password:        "pass",
				ConfirmPassword: "pass",
			},
			err: we.ErrInvalidParameter,
		},
		{
			name: "confirm password doesn't match password",
			input: RegisterInput{
				Username:        "bob",
				Email:           "bob@gmail.com",
				Password:        "password",
				ConfirmPassword: "wrongpassword",
			},
			err: we.ErrInvalidParameter,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.input.Validate()

			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestLoginInput_Sanitize(t *testing.T) {
	input := LoginInput{
		Email:    " BOB@gmail.com  ",
		Password: "password",
	}

	want := LoginInput{
		Email:    "bob@gmail.com",
		Password: "password",
	}

	input.Sanitize()

	require.Equal(t, want, input)
}

func TestLoginInput_Validate(t *testing.T) {
	testCases := []struct {
		name  string
		input LoginInput
		err   error
	}{
		{
			name: "valid",
			input: LoginInput{
				Email:    "bob@gmail.com",
				Password: "password",
			},
			err: nil,
		},
		{
			name: "invalid email",
			input: LoginInput{
				Email:    "bob",
				Password: "password",
			},
			err: we.ErrInvalidParameter,
		},
		{
			name: "empty password",
			input: LoginInput{
				Email:    "bob@gmail.com",
				Password: "",
			},
			err: we.ErrInvalidParameter,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.input.Validate()

			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
