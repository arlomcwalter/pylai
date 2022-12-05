package database

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

var testCases = []struct {
	account       *Account
	expectedError string
}{
	{
		account: &Account{
			Nickname: "",
			Secret:   "G37GZNCM32ZCX5OG",
		},
		expectedError: "invalid nickname format or length",
	},
	{
		account: &Account{
			Nickname: "test account !@#$%^&*()",
			Secret:   "G37GZNCM32ZCX5OG",
		},
		expectedError: "invalid nickname format or length",
	},
	{
		account: &Account{
			Nickname: "test account",
			Secret:   "ijefgoiwjej",
		},
		expectedError: "invalid secret key",
	},
	{
		account: &Account{
			Nickname: "test account",
			Secret:   "G37GZNCM32ZCX5OG",
		},
		expectedError: "",
	},
	{
		account: &Account{
			Nickname: "test account",
			Secret:   "G37GZNCM32ZCX5OG",
		},
		expectedError: "account with nickname 'test account' already exists",
	},
}

func TestDatabase(t *testing.T) {
	err := initDb("test_db")
	assert.NoError(t, err)

	for _, tc := range testCases {
		err := tc.account.Save()
		if tc.expectedError == "" {
			assert.NoError(t, err)
		} else {
			assert.EqualError(t, err, tc.expectedError)
		}
	}

	for _, tc := range testCases {
		account, err := Get(tc.account.Nickname)
		if tc.expectedError == "" {
			assert.NoError(t, err)
			assert.Equal(t, tc.account, account)
		}
	}

	for _, tc := range testCases {
		if tc.expectedError == "" {
			assert.NoError(t, tc.account.Delete())
		}
	}

	err = closeDb()
	assert.NoError(t, err)

	err = os.RemoveAll("test_db")
	assert.NoError(t, err)
}
