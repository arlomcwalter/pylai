package util

import (
	"github.com/pquerna/otp/totp"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValid(t *testing.T) {
	testCases := []struct {
		secret        string
		expectedValid bool
	}{
		{
			secret:        "",
			expectedValid: false,
		},
		{
			secret:        "short",
			expectedValid: false,
		},
		{
			secret:        "invalidcharacters123",
			expectedValid: false,
		},
		{
			secret:        "G37GZNCM32ZCX5OG",
			expectedValid: true,
		},
	}

	for _, tc := range testCases {
		result := IsValidSecret(tc.secret)
		assert.Equal(t, tc.expectedValid, result)
	}
}

func TestGeneration(t *testing.T) {
	testCases := []string{
		"ABCDEFGJKLMNOPQR",
		"STUVWXYZ23467JKL",
		"WXYZABCDEFGJKLMN",
		"VWXYZABCDEFGJKLM",
		"MNOPQRSTUVWXYZAB",
		"CDEFGHIJKLMNOPQR",
		"UVWXYZABCDEFGJKL",
		"PQRSTUVWXYZABCDE",
		"JKLMNOPQRSTUVWXY",
		"ABCDEFGJKLMNOPQR",
		"STUVWXYZ23467JKL",
		"WXYZABCDEFGJKLMN",
		"VWXYZABCDEFGJKLM",
		"MNOPQRSTUVWXYZAB",
		"CDEFGHIJKLMNOPQR",
		"UVWXYZABCDEFGJKL",
		"PQRSTUVWXYZABCDE",
		"JKLMNOPQRSTUVWXY",
		"ABCDEFGJKLMNOPQR",
		"STUVWXYZ23467JKL",
		"WXYZABCDEFGJKLMN",
		"VWXYZABCDEFGJKLM",
		"MNOPQRSTUVWXYZAB",
		"CDEFGHIJKLMNOPQR",
		"UVWXYZABCDEFGJKL",
	}

	for _, tc := range testCases {
		code, err := GenerateTotp(tc)
		assert.NoError(t, err)
		assert.True(t, totp.Validate(code, tc))
	}
}

func TestIsValidOTP(t *testing.T) {
	testCases := []struct {
		otp           string
		expectedValid bool
	}{
		{
			otp:           "",
			expectedValid: false,
		},
		{
			otp:           "123456",
			expectedValid: true,
		},
		{
			otp:           "1234567",
			expectedValid: false,
		},
		{
			otp:           "123 456",
			expectedValid: true,
		},
		{
			otp:           "123 4567",
			expectedValid: false,
		},
	}

	for _, tc := range testCases {
		result := IsValidOtp(tc.otp)
		assert.Equal(t, tc.expectedValid, result)
	}
}
