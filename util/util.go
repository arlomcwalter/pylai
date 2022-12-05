package util

import (
	"encoding/base32"
	"fmt"
	"github.com/pquerna/otp/totp"
	"regexp"
	"strings"
	"time"
)

const (
	otpPattern        = `^\d{3} ?\d{3}$`
	nicknameRegex     = "^[a-zA-Z0-9()\\s-_]*$"
	nicknameMaxLength = 16
)

func IsValidNickname(nickname string) bool {
	nickname = strings.TrimSpace(nickname)
	length := len(nickname)

	if length == 0 || length > nicknameMaxLength {
		return false
	}

	re := regexp.MustCompile(nicknameRegex)
	if !re.MatchString(nickname) {
		return false
	}

	return true
}

func IsValidSecret(secret string) bool {
	if len(secret) < 16 {
		return false
	}

	_, err := base32.StdEncoding.DecodeString(secret)
	if err != nil {
		return false
	}

	return true
}

func IsValidOtp(otp string) bool {
	return regexp.MustCompile(otpPattern).MatchString(otp)
}

func GenerateTotp(secret string) (string, error) {
	if !IsValidSecret(secret) {
		return "", fmt.Errorf("invalid secret")
	}
	return totp.GenerateCode(secret, time.Now())
}

func VerifyTotp(otp, secret string) bool {
	if !IsValidSecret(secret) {
		return false
	}
	return totp.Validate(otp, secret)
}
