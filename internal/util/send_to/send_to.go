package send_to

import (
	"bytes"
	"fmt"
	"html/template"
	"net/smtp"
	"strings"

	"github.com/nghiatk54/go_ecommerce_api/global"
	"go.uber.org/zap"
)

type EmailAddress struct {
	Address string `json:"address"`
	Name    string `json:"name"`
}

// Mail struct
type Mail struct {
	From    EmailAddress
	To      []string
	Subject string
	Body    string
}

// BuildMessage build message for email
func BuildMessage(mail Mail) string {
	message := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n"
	message += fmt.Sprintf("From: %s\n", mail.From.Address)
	message += fmt.Sprintf("To: %s\n", strings.Join(mail.To, ";"))
	message += fmt.Sprintf("Subject: %s\n", mail.Subject)
	message += fmt.Sprintf("\n%s\n", mail.Body)
	return message
}

// SendTextEmailOtp send text email otp
func SendTextEmailOtp(to []string, from string, otp string) error {
	contentEmail := Mail{
		From:    EmailAddress{Address: from, Name: "test"},
		To:      to,
		Subject: "Otp verification",
		Body:    fmt.Sprintf("Your Otp is %s. Please enter it to verify your account.", otp),
	}
	messageEmail := BuildMessage(contentEmail)
	// send smtp
	auth := smtp.PlainAuth("", global.Config.Smtp.Username, global.Config.Smtp.Password, global.Config.Smtp.Host)
	err := smtp.SendMail(fmt.Sprintf("%s:%d", global.Config.Smtp.Host, global.Config.Smtp.Port), auth, from, to, []byte(messageEmail))
	if err != nil {
		global.Logger.Error("Send email failed", zap.Error(err))
		return err
	}
	return nil
}

// SendTemplateEmail send template email
func SendTemplateEmailOtp(
	to []string,
	from string,
	nameTemplate string,
	dataTemplate map[string]interface{},
) error {
	htmlBody, err := getMailTemplate(nameTemplate, dataTemplate)
	if err != nil {
		return err
	}
	return send(to, from, htmlBody)
}

func getMailTemplate(nameTemplate string, dataTemplate map[string]interface{}) (string, error) {
	htmlTemplate := new(bytes.Buffer)
	t := template.Must(template.New(nameTemplate).ParseFiles("templates_email/" + nameTemplate))
	err := t.Execute(htmlTemplate, dataTemplate)
	if err != nil {
		return "", err
	}
	return htmlTemplate.String(), nil
}

// send send email
func send(to []string, from string, htmlTemplate string) error {
	contentEmail := Mail{
		From:    EmailAddress{Address: from, Name: "test"},
		To:      to,
		Subject: "Otp verification",
		Body:    htmlTemplate,
	}
	messageEmail := BuildMessage(contentEmail)
	// send smtp
	auth := smtp.PlainAuth("", global.Config.Smtp.Username, global.Config.Smtp.Password, global.Config.Smtp.Host)
	err := smtp.SendMail(fmt.Sprintf("%s:%d", global.Config.Smtp.Host, global.Config.Smtp.Port), auth, from, to, []byte(messageEmail))
	if err != nil {
		global.Logger.Error("Send email failed", zap.Error(err))
		return err
	}
	return nil
}
