package security

import (
	"fmt"
	"net/url"

	"github.com/pquerna/otp/totp"
	"github.com/skip2/go-qrcode"
)

func GenerateTOTPSecret(username string) (string, error) {
	secret, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "YourAppName",
		AccountName: username,
	})
	if err != nil {
		return "", err
	}

	return secret.Secret(), nil
}

func GenerateTOTPQRCodeURL(secret, username string) string {
	uri := fmt.Sprintf("otpauth://totp/%s?secret=%s&issuer=%s", url.QueryEscape(username), url.QueryEscape(secret), "YourAppName")
	return uri
}

func GenerateQRCode(url string) ([]byte, error) {
	var png []byte
	png, err := qrcode.Encode(url, qrcode.Medium, 256) // Размер 256x256
	if err != nil {
		return nil, err
	}

	return png, nil
}

func ValidateTOTPCode(secret string, code string) bool {
	return totp.Validate(code, secret)
}
