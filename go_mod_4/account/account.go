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
	Login    string    `json:"login`
	Password string    `json:"password"`
	Url      string    `json:"url"`
	Creat    time.Time `json:"creat"`
	Update   time.Time `json:"update"`
}

func NewAccount(login, password, urlString string) (*Account, error) {

	if login == "" {
		return nil, errors.New("invallogin")
	}
	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("Invalid Url")
	}

	newAcc := &Account{
		Creat:    time.Now(),
		Update:   time.Now(),
		Url:      urlString,
		Login:    login,
		Password: password,
	}

	if password == "" {
		newAcc.generatePassword(4)
	}
	return newAcc, nil
}

func (acc *Account) Output() {
	color.Cyan(acc.Login)
	color.Cyan(acc.Password)
	color.Cyan(acc.Url)

	// fmt.Println(acc.login, acc.password, acc.url)
}

func (acc *Account) generatePassword(n int) {
	res := make([]rune, n)
	for i := range res {
		res[i] = leterRunes[rand.IntN(len(leterRunes))]
	}
	acc.Password = string(res)
}
