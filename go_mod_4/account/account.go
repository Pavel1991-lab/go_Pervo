package account

import (
	"errors"
	"math/rand/v2"
	"net/url"
	"time"

	"github.com/fatih/color"
)

var leterRunes = []rune("abcdefgh123")

type Account struct {
	login    string
	password string
	url      string
}

type AccountWithTimeStamp struct {
	creat  time.Time
	update time.Time
	Account
}

func NewAccountWithTimeStamp(login, password, urlString string) (*AccountWithTimeStamp, error) {

	if login == "" {
		return nil, errors.New("invallogin")
	}
	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("Invalid Url")
	}

	newAcc := &AccountWithTimeStamp{
		creat:  time.Now(),
		update: time.Now(),
		Account: Account{
			url:      urlString,
			login:    login,
			password: password,
		},
	}
	if password == "" {
		newAcc.generatePassword(4)
	}
	return newAcc, nil
}

func (acc Account) OutputPassword() {
	color.Cyan(acc.login)
	// fmt.Println(acc.login, acc.password, acc.url)
}

func (acc *Account) generatePassword(n int) {
	res := make([]rune, n)
	for i := range res {
		res[i] = leterRunes[rand.IntN(len(leterRunes))]
	}
	acc.password = string(res)
}
