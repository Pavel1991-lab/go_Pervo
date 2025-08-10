package main

import (
	"fmt"
)

func main() {

	login := promtData("add log")
	password := promtData("add_pass")
	url := promtData("add url")

	myAccount, err := newAccountwithtimeStamp(login, password, url)
	if err != nil {
		fmt.Println("Error url")
		return
	}
	myAccount.outputPassword()
	fmt.Println(myAccount)

}

func promtData(promt string) string {
	fmt.Println(promt + ": ")
	var res string
	fmt.Scan(&res)
	return res
}
