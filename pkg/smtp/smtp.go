package smtp

import (
	"fmt"
	"reservation/config"
	"strings"

	"github.com/go-gomail/gomail"
)

func SendEmail(senderEmail, targetEmail, subjectMessage, otpCode string) error {
	host := config.LoadConfig().SmtpHost
	senderName := config.LoadConfig().SmtpSenderName
	username := config.LoadConfig().SmtpUsername
	password := config.LoadConfig().SmtpPassword
	port := config.LoadConfig().SmtpPort

	if senderEmail == "" {
		senderEmail = username
	}

	body := fmt.Sprintf("Your OTP code is: %s", otpCode)

	mailer := gomail.NewMessage()
	mailer.SetHeader("From", fmt.Sprintf("%s <%s>", senderName, senderEmail))
	mailer.SetHeader("To", strings.ToLower(targetEmail))
	mailer.SetHeader("Subject", subjectMessage)
	mailer.SetBody("text/plain", body)

	dialer := gomail.NewDialer(host, port, username, password)

	if err := dialer.DialAndSend(mailer); err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	return nil
}
