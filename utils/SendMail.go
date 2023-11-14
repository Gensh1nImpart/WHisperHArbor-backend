package utils

import (
	"WHisperHArbor-backend/model"
	"fmt"
	"github.com/jordan-wright/email"
	"net/smtp"
)

func SendMail(to []string, verifyUrl string) error {
	e := email.NewEmail()
	e.From = fmt.Sprintf("WHHA <yrh6@qq.com>")
	e.To = to
	e.Subject = "WHHA注册验证"
	content := fmt.Sprintf(`
	<div>
		<div style="padding: 8px 40px 8px 50px;">
			<p>您于 %s 提交的邮箱验证，验证链接为<a href="%s">%s</a>，为了保证账号安全，请与24小时内点击链接验证。请确认为本人操作，切勿向他人泄露，感谢您的理解。</p>
		</div>
		<div>
			<p>此邮箱为系统邮箱，请勿回复。</p>
		</div>
	</div>`, to[0], verifyUrl, verifyUrl)
	e.HTML = []byte(content)
	err := e.Send(model.MyMail.Server+":"+model.MyMail.Port, smtp.PlainAuth("", model.MyMail.Account, model.MyMail.Passwd, model.MyMail.Server))
	if err != nil {
		return err
	}
	return nil
}
