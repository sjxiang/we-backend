
package tencent

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	sdk "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20210111"
)

// 这个需要手动跑，也就是你需要在本地搞好这些环境变量
func TestSender(t *testing.T) {
	secretID, ok := os.LookupEnv("SMS_SECRET_ID")
	if !ok {
		t.Fatal()
	}
	secretKey, ok := os.LookupEnv("SMS_SECRET_KEY")
	if !ok {
		t.Fatal()
	}

	c, err := sdk.NewClient(common.NewCredential(secretID, secretKey),
		"ap-nanjing",
		profile.NewClientProfile())
	if err != nil {
		t.Fatal(err)
	}

	svc := NewTencentSMS(c, "xxxx", "惠惠科技")
	
	testCases := []struct {
		name    string
		tplID   string
		params  []string
		numbers []string
		wantErr error
	}{
		{
			name:   "发送验证码",
			tplID:  "1877556",
			params: []string{"123456"},
			// 改成你的手机号码
			numbers: []string{""},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			er := svc.Send(context.Background(), tc.tplID, tc.params, tc.numbers...)
			assert.Equal(t, tc.wantErr, er)
		})
	}
}