package local

import (
	"fmt"
	"context"
	
	"we-backend/internal/service/sms"
)

type localSMS struct {
}

func NewLocalSMS() sms.ShortMessageService {
	return &localSMS{}
}

// 方便本地测试
func (impl *localSMS) Send(ctx context.Context, tplId string, args []string, numbers ...string) error {
	fmt.Println(args)
	return nil 
}
	