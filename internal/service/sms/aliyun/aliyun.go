package aliyun

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"context"
	
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	
	"we-backend/internal/service/sms"
)

type aliyunSMS struct {
	client   *dysmsapi.Client
	signName string
}

func NewAliyunSMS(client *dysmsapi.Client, signName string) sms.ShortMessageService {
	return &aliyunSMS{
		client:     client,
		signName: signName,
	}
}

// []string
// 这是验证码{code}，然后要求有一个叫做 code 的参数
func (impl *aliyunSMS) Send(ctx context.Context, tplId string, args []string, numbers ...string) error {
	req := dysmsapi.CreateSendSmsRequest()
	req.Scheme = "https"
	// 阿里云多个手机号为字符串逗号间隔
	req.PhoneNumbers = strings.Join(numbers, ",")
	req.SignName = impl.signName
	
	// 传的是 JSON
	argsMap := make(map[string]string, len(args))
	for idx, arg := range args {
		argsMap[strconv.Itoa(idx)] = arg
	}
	// 这意味着，你的模板必须是 你的短信验证码是{0}
	// 你的短信验证码是{code}
	bCode, err := json.Marshal(argsMap)
	if err != nil {
		return err
	}
	req.TemplateParam = string(bCode)
	req.TemplateCode = tplId

	var resp *dysmsapi.SendSmsResponse
	resp, err = impl.client.SendSms(req)
	if err != nil {
		return err
	}

	if resp.Code != "OK" {
		return fmt.Errorf("发送失败，code: %s, 原因：%s",
			resp.Code, resp.Message)
	}
	return nil
}
