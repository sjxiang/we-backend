package tencent

import (
	"fmt"
	"context"

	sdk "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20210111"
	"github.com/rs/zerolog/log"
	
	"we-backend/internal/service/sms"
)

type tencentSMS struct {
	client        *sdk.Client
	appID         *string
	signatureName *string
}

func NewTencentSMS(client *sdk.Client, appID string, signName string) sms.ShortMessageService {
	return &tencentSMS{
		client:        client,
		appID:         &appID,
		signatureName: &signName,
	}
} 

func (impl *tencentSMS) Send(ctx context.Context, tplId string, args []string, numbers ...string) error {
	request := sdk.NewSendSmsRequest()
	request.SetContext(ctx)

	request.SmsSdkAppId      = impl.appID
	request.SignName         = impl.signatureName
	request.TemplateId       = &tplId
	request.TemplateParamSet = toPtrSlice(args)
	request.PhoneNumberSet   = toPtrSlice(numbers)
	
	response, err := impl.client.SendSms(request)

	log.Info().Any("request", request).Any("response", response).Msg("请求腾讯 SendSMS 接口")

	if err != nil {
		return err
	}

	for _, statusPtr := range response.Response.SendStatusSet {
		if statusPtr == nil {
			// 不可能进来这里
			continue
		}

		// 不要相信昨天的自己，检查是否是 nil，如果是，解引用会引发 panic
		status := *statusPtr
		if status.Code == nil || *(status.Code) != "Ok" {
			// 发送失败
			return fmt.Errorf("发送短信失败 code: %s, msg: %s", *status.Code, *status.Message)
		}
	}

	return nil 
}


func toPtrSlice(data []string) []*string {
	tmp := make([]*string, len(data))

	for i, s := range data {
		tmp[i] = &s
	}

	return tmp
}
