package main

import (
	"fmt"
	"main/account"
	"main/files"
)

func main() {

	login := promptData("add log")
	password := promptData("add_pass")
	url := promptData("add url")

	myAccount, err := account.NewAccountWithTimeStamp(login, password, url)
	if err != nil {
		fmt.Println("Error url")
		return
	}
	myAccount.OutputPassword()
	files.WrightFile()
	fmt.Println(myAccount)
}

func promptData(prompt string) string {
	fmt.Println(prompt + ": ")
	var res string
	fmt.Scan(&res)
	return res
}
