package account

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"
	"time"
)

var leterRunes = []rune("abcdefgh123")

type account struct {
	login    string
	password string
	url      string
}

type accountWithTimeStamp struct {
	creat  time.Time
	update time.Time
	account
}

func newAccountwithtimeStamp(login, password, urlString string) (*accountWithTimeStamp, error) {

	if login == "" {
		return nil, errors.New("invallogin")
	}
	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("Invalid Url")
	}

	newAcc := &accountWithTimeStamp{
		creat:  time.Now(),
		update: time.Now(),
		account: account{
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

func (acc account) outputPassword() {
	fmt.Println(acc.login, acc.password, acc.url)
}

func (acc *account) generatePassword(n int) {
	res := make([]rune, n)
	for i := range res {
		res[i] = leterRunes[rand.IntN(len(leterRunes))]
	}
	acc.password = string(res)
}
