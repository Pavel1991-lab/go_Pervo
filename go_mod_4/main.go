package main

import (
	"fmt"
	"main/account"
	"main/files"
	"main/output"
	"strings"

	"github.com/fatih/color"
)

var menu = map[string]func(*account.VaultWithDb){
	"1": createAccount,
	"2": findAccount,
	"3": deleteAccount,
}

func main() {
	fmt.Println("I am manager of password")
	vault := account.NewVault(files.NewJsonDb(("data.json")))
Menu:
	for {
		variant := promptData([]string{
			"1. createAccount",
			"2. findAccount",
			"3. deleteAccount",
			"4. exit",
			"5. chose variant",
		})
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

func findAccount(vault *account.VaultWithDb) {
	url := promptData([]string{"Add url"})
	accounts := vault.FindAccounts(url, check_url)
	if len(accounts) == 0 {
		color.Yellow("No account")
	}
	for _, account := range accounts {
		account.Output()
	}
}

func check_url(acc account.Account, str string) bool {
	return strings.Contains(acc.Url, str)
}

func deleteAccount(vault *account.VaultWithDb) {

	url := promptData([]string{"enter url for delete"})
	isDeleted := vault.DeleteAccoutByUrl(url)
	if isDeleted {
		color.Green("Delete")
	} else {
		output.PrintError("Do not found")
	}
}

func createAccount(vault *account.VaultWithDb) {
	login := promptData([]string{"Add login"})
	password := promptData([]string{"Add password"})
	url := promptData([]string{"Add url"})

	myAccount, err := account.NewAccount(login, password, url)
	if err != nil {
		output.PrintError("Wrong format")
		return
	}
	vault.AddAccount(*myAccount)
}

func promptData[T any](prompt []T) string {
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
