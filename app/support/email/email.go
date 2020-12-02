package email

import (
	"github.com/gogf/gf/frame/g"
	"gopkg.in/gomail.v2"
)

type Letter struct {
	From     string // 发件人
	FromName string // 收件人名称
	To       string // 收件人
	ToName   string // 收件人名称
	Cc       string // 抄送人
	CcName   string // 抄送人名称
	Subject  string // 邮件主题
	Body     string // 邮件内容
}

func (letter *Letter) Send() error {
	message := gomail.NewMessage()

	// 发件人
	letter.setFrom(message)

	// 收件人
	letter.setTo(message)

	// 抄送人
	letter.setCc(message)

	// 邮件主题
	message.SetHeader("Subject", letter.Subject)

	// 邮件内容
	message.SetBody("text/html", letter.Body)

	d := gomail.NewDialer(g.Cfg().GetString("email.Host"),
		g.Cfg().GetInt("email.Port"),
		g.Cfg().GetString("email.Username"),
		g.Cfg().GetString("email.Password"))

	return d.DialAndSend(message)
}

// 设置发件人
func (letter *Letter) setCc(message *gomail.Message) {
	if len(letter.CcName) > 0 {
		message.SetAddressHeader("Cc", letter.Cc, letter.CcName)
	} else if len(letter.Cc) > 0 {
		message.SetHeader("Cc", letter.Cc)
	}
}

// 设置发件人
func (letter *Letter) setFrom(message *gomail.Message) {
	if len(letter.ToName) > 0 {
		message.SetAddressHeader("From", letter.From, letter.FromName)
	} else {
		message.SetHeader("From", letter.From)
	}
}

// 设置收件人
func (letter *Letter) setTo(message *gomail.Message) {
	if len(letter.ToName) > 0 {
		message.SetAddressHeader("To", letter.To, letter.ToName)
	} else {
		message.SetHeader("To", letter.To)
	}
}
