package goutil

import (
	"encoding/base64"
	"fmt"
	"net/mail"
	"net/smtp"
)

// MailInfo 邮件
type MailInfo struct {
	SMTPHost   string
	SMTPPasswd string
	FromEmail  string
	FromName   string
	ToEmail    string
	ToName     string
	Subject    string
	Body       string
}

// SendMail 发送邮件
// @mailInfo
func SendMail(mailInfo MailInfo) {
	b64 := base64.NewEncoding("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/")
	// 邮件服务器信息
	from := mail.Address{Name: mailInfo.FromName, Address: mailInfo.FromEmail}
	to := mail.Address{Name: mailInfo.ToName, Address: mailInfo.ToEmail}
	// 头文件
	header := make(map[string]string)
	header["From"] = from.String()
	header["To"] = to.String()
	header["Subject"] = fmt.Sprintf("=?UTF-8?B?%s?=", b64.EncodeToString([]byte(mailInfo.Subject)))
	header["MIME-Version"] = "1.0"
	header["Content-Type"] = "text/html; charset=UTF-8"
	header["Content-Transfer-Encoding"] = "base64"
	// 主体信息
	message := ""
	for k, v := range header {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + b64.EncodeToString([]byte(mailInfo.Body))
	// 服务器身份验证
	auth := smtp.PlainAuth(
		"",
		mailInfo.FromEmail,
		mailInfo.SMTPPasswd,
		mailInfo.SMTPHost,
	)
	// 发送邮件
	err := smtp.SendMail(
		mailInfo.SMTPHost+":25",
		auth,
		mailInfo.FromEmail,
		[]string{to.Address},
		[]byte(message),
	)
	// 异常处理
	if err != nil {
		fmt.Println("邮件发送失败:" + err.Error())
		return
	}

	fmt.Println("已经发送邮件")
}
