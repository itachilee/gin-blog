package email

import (
	"crypto/tls"
	"gopkg.in/gomail.v2"
)

type Email struct {
	*SMTPInfo
}
type SMTPInfo struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	IsSSL    bool   `json:"is_ssl"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
	From     string `json:"from"`
}

func NewEmail(info *SMTPInfo) *Email {
	return &Email{
		SMTPInfo: info,
	}
}

func (e *Email) SendMail(to []string, subject, body string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", e.From)
	m.SetHeader("to", to...)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)
	dialer := gomail.NewDialer(e.Host, e.Port, e.UserName, e.Password)
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: e.IsSSL}
	return dialer.DialAndSend(m)
}
