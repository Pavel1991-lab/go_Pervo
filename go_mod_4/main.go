package main

import (
	"fmt"
	"main/account"
	"main/files"

	"github.com/fatih/color"
)

func main() {
	fmt.Println("I am manager of password")
	vault := account.NewVault(files.NewJsonDb(("data.json")))
Menu:
	for {
		variant := getMenu()
		switch variant {
		case 1:
			createAccount(vault)
		case 2:
			findAccount(vault)
		case 3:
			deleteAccount(vault)
		default:
			break Menu

		}
	}

}

func getMenu() int {
	var variant int
	fmt.Println("Chose variant")
	fmt.Println("1 Create account")
	fmt.Println("2 Find accont")
	fmt.Println("3 Delete account")
	fmt.Println("4 Exit")
	fmt.Scan(&variant)
	return variant

}

func findAccount(vault *account.VaultWithDb) {
	url := promptData("Add url")
	accounts := vault.FindAccountsByUrl(url)
	if len(accounts) == 0 {
		color.Yellow("No account")
	}
	for _, account := range accounts {
		account.Output()
	}
}

func deleteAccount(vault *account.VaultWithDb) {

	url := promptData("Add url for delete")
	isDeleted := vault.DeleteAccoutByUrl(url)
	if isDeleted {
		color.Green("Delete")
	} else {
		color.Red("Not Found")
	}
}

func createAccount(vault *account.VaultWithDb) {
	login := promptData("add log")
	password := promptData("add_pass")
	url := promptData("add url")

	myAccount, err := account.NewAccount(login, password, url)
	if err != nil {
		fmt.Println("Error url")
		return
	}
	vault.AddAccount(*myAccount)
}

func promptData(prompt string) string {
	fmt.Println(prompt + ": ")
	var res string
	fmt.Scan(&res)
	return res
}
