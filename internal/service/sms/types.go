package sms

import "context"


// 短信服务 _ short message service _ SMS
type ShortMessageService interface {
	Send(ctx context.Context, tplId string, args []string, numbers ...string) error
}


// 另一种实现方式，参数详细 or 粗略
// 适配器，间接调用

type Params struct {

}


/*

供应商
	腾讯、阿里云


设计

	发短信、发邮件

		验证码
		通知


要求太多，要备案，没戏！


 */
