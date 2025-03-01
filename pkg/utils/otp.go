package utils

import (
	"crypto/rand"
	"fmt"
)

func GenerateOTP() (string, error) {
	const otpLength = 6
	const digits = "0123456789"
	otp := make([]byte, otpLength)

	_, err := rand.Read(otp)
	if err != nil {
		return "", fmt.Errorf("failed to generate OTP: %w", err)
	}

	for i := range otp {
		otp[i] = digits[otp[i]%10]
	}

	return string(otp), nil
}
