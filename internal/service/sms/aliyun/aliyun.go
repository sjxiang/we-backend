package aliyun

import (
	"context"
	"we-backend/internal/service/sms"
)

type aliyunSMS struct {

}

func NewAliyunSMS() sms.ShortMessageService {
	return &aliyunSMS{}
} 

func (impl *aliyunSMS) Send(ctx context.Context, tplId string, args []string, numbers ...string) error {
	return nil 
}