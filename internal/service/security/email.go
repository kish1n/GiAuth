package security

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"sync"
	"time"

	"github.com/kish1n/GiAuth/internal/config"
	"gopkg.in/gomail.v2"
)

var emailListCode = make(map[string]string)

var mu sync.Mutex

func AddToEmailList(username string, code string) {
	mu.Lock()
	emailListCode[username] = code
	mu.Unlock()

	time.AfterFunc(180*time.Second, func() {
		mu.Lock()
		defer mu.Unlock()
		if _, exists := emailListCode[username]; exists {
			delete(emailListCode, username)
		}
	})
}

func CheckInEmailList(email string, code string) bool {
	mu.Lock()
	defer mu.Unlock()

	if storedCode, exists := emailListCode[email]; exists && storedCode == code {
		delete(emailListCode, email)
		return true
	}

	return false
}

func SendConfirmationEmail(to string, code string, cfg config.Config) error {
	m := gomail.NewMessage()

	m.SetHeader("From", cfg.Email().Address)
	m.SetHeader("To", to)
	m.SetHeader("Subject", "Address Confirmation")

	plainText := fmt.Sprintf("Your confirmation code: %s", code)

	htmlContent := fmt.Sprintf(`
		<html>
		<body style="font-family: Arial, sans-serif; background-color: #f4f4f4; padding: 20px;">
			<div style="max-width: 600px; margin: auto; background-color: #fff; padding: 20px; border-radius: 10px; box-shadow: 0px 0px 10px rgba(0,0,0,0.1);">
				<h2 style="text-align: center; color: #4CAF50;">Address Confirmation</h2>
				<p style="text-align: center; font-size: 18px;">Your confirmation code:</p>
				<div style="text-align: center; font-size: 24px; font-weight: bold; color: #333; padding: 20px 0; background-color: #f9f9f9; border-radius: 5px;">
					%s
				</div>
				<p style="text-align: center; font-size: 14px; color: #666;">This code is valid for the next 3 minutes.</p>
				<hr style="border: none; border-top: 1px solid #eee; margin: 20px 0;">
				<p style="text-align: center; font-size: 12px; color: #999;">If you did not request this code, please ignore this email.</p>
			</div>
		</body>
		</html>
	`, code)

	m.SetBody("text/plain", plainText)
	m.AddAlternative("text/html", htmlContent)

	d := gomail.NewDialer("smtp.gmail.com", 587, cfg.Email().Address, cfg.Email().Password)

	if err := d.DialAndSend(m); err != nil {
		return err
	}

	AddToEmailList(to, code)

	return nil
}

func SendLoginAttemptEmail(to string, code string, cfg config.Config) error {
	m := gomail.NewMessage()

	m.SetHeader("From", cfg.Email().Address)
	m.SetHeader("To", to)
	m.SetHeader("Subject", "Login Attempt Confirmation")

	htmlContent := fmt.Sprintf(`
		<html>
		<body style="font-family: Arial, sans-serif; background-color: #f4f4f4; padding: 20px;">
			<div style="max-width: 600px; margin: auto; background-color: #fff; padding: 20px; border-radius: 10px; box-shadow: 0px 0px 10px rgba(0,0,0,0.1);">
				<h2 style="text-align: center; color: #4CAF50;">Login Attempt</h2>
				<p style="text-align: center; font-size: 18px;">We detected a login attempt for your account. If this was you, please confirm:</p>
				<div style="text-align: center; font-size: 18px; padding: 20px;">
					<a href="http://localhost:8000/integrations/GiAuth/auth/%s/%s" style="padding: 10px 20px; background-color: #4CAF50; color: white; text-decoration: none; border-radius: 5px;">Confirm Login</a>
				</div>
				<p style="text-align: center; font-size: 14px; color: #666;">If this wasn't you, please ignore this email.</p>
			</div>
		</body>
		</html>
	`, to, code)

	plainText := fmt.Sprintf("We detected a login attempt for your account. To confirm, use the following link: http://localhost:8000/integrations/GiAuth/auth/%s/%s", to, code)

	m.SetBody("text/plain", plainText)
	m.AddAlternative("text/html", htmlContent)

	d := gomail.NewDialer("smtp.gmail.com", 587, cfg.Email().Address, cfg.Email().Password)

	if err := d.DialAndSend(m); err != nil {
		return err
	}

	AddToEmailList(to, code)
	return nil
}

func GenerateConfirmationCode() (string, error) {
	var number uint32
	err := binary.Read(rand.Reader, binary.BigEndian, &number)
	if err != nil {
		return "", err
	}
	code := number % 1000000
	return fmt.Sprintf("%06d", code), nil
}
