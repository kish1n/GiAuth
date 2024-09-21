package security

import (
	"crypto/rand"
	"fmt"
	"gopkg.in/gomail.v2"
	"sync"
	"time"
)

var emailListCode = make(map[string]string)

var mu sync.Mutex

func AddToEmailList(username string, code string) {
	mu.Lock()
	emailListCode[username] = code
	mu.Unlock()

	time.AfterFunc(60*time.Second, func() {
		mu.Lock()
		defer mu.Unlock()
		if _, exists := emailListCode[username]; exists {
			delete(emailListCode, username)
		}
	})
}

func CheckInEmailList(username string, code string) bool {
	mu.Lock()
	defer mu.Unlock()

	if storedCode, exists := emailListCode[username]; exists && storedCode == code {
		delete(emailListCode, username)
		return true
	}

	return false
}

func SendConfirmationEmail(to string, code string) error {
	m := gomail.NewMessage()

	m.SetHeader("From", "testemail1488@proton.me")
	m.SetHeader("To", to)
	m.SetHeader("Subject", "Email Confirmation")

	m.SetBody("text/plain", fmt.Sprintf("Your confirmation code: %s", code))

	d := gomail.NewDialer("smtp.example.com", 587, "testemail1488@proton.me", "=APWBAz-%6MCzxz")

	if err := d.DialAndSend(m); err != nil {
		return err
	}

	AddToEmailList(to, code)

	return nil
}

func GenerateConfirmationCode() (string, error) {
	b := make([]byte, 6)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	code := fmt.Sprintf("%06d", b[0:3])
	return code, nil
}
