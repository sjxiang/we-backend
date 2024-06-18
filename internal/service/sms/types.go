package sms

import "context"


// 短信服务 SMS
type ShortMessageService interface {
	Send(ctx context.Context, tplId string, args []string, numbers ...string) error
}


// 另一种实现方式
type Params struct {

}


/*

供应商
	腾讯、阿里云


要求太多，要备案，没戏！


 */
