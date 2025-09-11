package main

import (
	"fmt"
	"main/account"
	"main/encrypter"
	"main/files"
	"main/output"
	"strings"

	"github.com/fatih/color"
	"github.com/joho/godotenv"
)

var menu = map[string]func(*account.VaultWithDb){
	"1": createAccount,
	"2": findAccountByUrl,
	"3": findAccountByLogin,
	"4": deleteAccount,
}

var menuVariants = []string{
	"1. createAccount",
	"2. findAccount by url",
	"3. findAccount by login",
	"4. deleteAccount",
	"5. exit",
}

func manuCounter() func() {
	i := 0
	return func() {
		i++
		fmt.Println(i)
	}
}

func main() {
	fmt.Println("I am manager of password")
	err := godotenv.Load()
	if err != nil {
		output.PrintError("Do not find env file")
	}
	vault := account.NewVault(files.NewJsonDb("data.json"), *encrypter.NewEncrypter())
	conunter := manuCounter()
Menu:
	for {
		conunter()
		variant := promptData(
			menuVariants...,
		)
		menuFunc := menu[variant]
		if menuFunc == nil {
			break Menu
		}
		menuFunc(vault)
		// switch variant {
		// case "1":
		// 	createAccount(vault)
		// case "2":
		// 	findAccount(vault)
		// case "3":
		// 	deleteAccount(vault)
		// default:
		// 	break Menu

		// }
	}

}

func findAccountByUrl(vault *account.VaultWithDb) {
	url := promptData("Add url")
	accounts := vault.FindAccounts(url, func(acc account.Account, str string) bool {
		return strings.Contains(acc.Url, str)
	})
	outputResult(&accounts)

}

func findAccountByLogin(vault *account.VaultWithDb) {
	login := promptData("Add login")
	accounts := vault.FindAccounts(login, func(acc account.Account, str string) bool {
		return strings.Contains(acc.Login, str)
	})
	outputResult(&accounts)
}

func outputResult(accounts *[]account.Account) {
	if len(*accounts) == 0 {
		color.Yellow("No account")
	}
	for _, account := range *accounts {
		account.Output()
	}
}

func deleteAccount(vault *account.VaultWithDb) {

	url := promptData("enter url for delete")
	isDeleted := vault.DeleteAccoutByUrl(url)
	if isDeleted {
		color.Green("Delete")
	} else {
		output.PrintError("Do not found")
	}
}

func createAccount(vault *account.VaultWithDb) {
	login := promptData("Add login")
	password := promptData("Add password")
	url := promptData("Add url")

	myAccount, err := account.NewAccount(login, password, url)
	if err != nil {
		output.PrintError("Wrong format")
		return
	}
	vault.AddAccount(*myAccount)
}

func promptData(prompt ...string) string {
	for i, line := range prompt {
		if i == len(prompt)-1 {
			fmt.Printf("%v: ", line)
		} else {
			fmt.Println(line)
		}
	}
	var res string
	fmt.Scan(&res)
	return res
}
