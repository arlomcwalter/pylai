package database

import (
	"fmt"
	"github.com/arlomcwalter/pylai/util"
	"strings"
)

type Account struct {
	Nickname string
	Secret   string
}

func (a *Account) Save() error {
	a.Nickname = strings.TrimSpace(a.Nickname)

	if !util.IsValidSecret(a.Secret) {
		return fmt.Errorf("invalid secret key")
	}

	if !util.IsValidNickname(a.Nickname) {
		return fmt.Errorf("invalid nickname format or length")
	}

	existingAcc, err := Get(a.Nickname)
	if err == nil && existingAcc != nil {
		return fmt.Errorf("account with nickname '%s' already exists", a.Nickname)
	}

	return DB.Put([]byte(a.Nickname), []byte(a.Secret), nil)
}

func (a *Account) Delete() error {
	return DB.Delete([]byte(a.Nickname), nil)
}

func GetAll() ([]*Account, error) {
	accounts := make([]*Account, 0)

	iter := DB.NewIterator(nil, nil)
	for iter.Next() {
		key := iter.Key()
		value := iter.Value()

		a := &Account{
			Nickname: string(key),
			Secret:   string(value),
		}

		accounts = append(accounts, a)
	}

	return accounts, nil
}

func Get(nickname string) (*Account, error) {
	data, err := DB.Get([]byte(nickname), nil)
	if err != nil {
		return nil, fmt.Errorf("account with nickname '%s' not found", nickname)
	}

	acc := &Account{
		Nickname: nickname,
		Secret:   string(data),
	}

	return acc, nil
}
