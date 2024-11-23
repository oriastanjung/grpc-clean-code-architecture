package smtp

import (
	"fmt"
	"net/smtp"
	"strings"

	"github.com/oriastanjung/grpc_server/internal/config"
)

// SendEmail mengirim email verifikasi dengan tombol dinamis
func SendEmailForgetPassword(to, forgetPasswordToken string, link string) error {
	cfg := config.LoadEnv()
	// Konfigurasi email
	from := cfg.GmailEmail             // Ganti dengan email pengirim Gmail
	password := cfg.GmailPassword      // Ganti dengan password email pengirim
	smtpServer := "smtp.gmail.com:587" // SMTP server Gmail dengan port TLS

	// Template email dengan tombol aktivasi
	subject := "Reset Password Account"
	body := fmt.Sprintf(` 
		<p>Hi,</p>
		<p>We received your request to reset your password. Click the button to start reset your password :</p>
		<a href="%s/%s" style="background-color: #4CAF50; color: white; padding: 15px 32px; text-align: center; text-decoration: none; display: inline-block; font-size: 16px; margin: 4px 2px; cursor: pointer; border-radius: 4px;">Reset Password</a>
		<p>If you did not register for this account, please ignore this email.</p>
		<p>Best regards,<br>Your Company</p>
	`, link, forgetPasswordToken)

	msg := []byte("To: " + to + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"Content-Type: text/html; charset=UTF-8\r\n" +
		"\r\n" +
		body)

	// Create a channel to receive errors
	errCh := make(chan error)

	// Start a goroutine to send the email
	go func() {
		auth := smtp.PlainAuth("", from, password, strings.Split(smtpServer, ":")[0])
		err := smtp.SendMail(smtpServer, auth, from, []string{to}, msg)
		if err != nil {
			errCh <- err
			return
		}
		errCh <- nil
	}()

	// Wait for the result from the goroutine
	err := <-errCh
	return err
}