package email

import (
	"github.com/google/wire"
	"configs"
	"log"
)

type EMail interface {
	Send()
}

// mailSender MailSender接口实现
type eMail struct {
}

// Send 发邮件
func (e *eMail) Send() {
	log.Println("send email")
}

// NewMailSender provider: 根据邮件配置初始化 mailSender
func NewMailSender(m *configs.EMailConfig) *eMail {
	return &eMail{}
}

// MailSet 声明NewMailSender的返回值是MailSender接口类型
var EMailSet = wire.NewSet(NewMailSender, wire.Bind(new(EMail), new(*EMail)))