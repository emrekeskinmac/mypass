package main

import (
	"fmt"
	"log"

	"github.com/emrekeskinmac/mypass"
	"github.com/emrekeskinmac/mypass/account"
	"github.com/emrekeskinmac/mypass/providers/local"
)

// Example useage:
// $ mypass login
// Enter a username:
// emremac
// Enter a password:
// 123
//
// $mypass save
// Choose a type:
// (1) Password
// 1
// Enter a name:
// facebook
// Enter a username:
// emremac
// Enter a password:
// 123
//
// $ mypass get [facebook]
// emremac - 123
//
// $ mypass delete [facebook]
// Deleted: emremac - 123
func main() {
	username := "emremac"
	password := "emremac123"
	dbPath := "/tmp/mypass.db"

	pv, err := local.New(username, password, dbPath)
	if err != nil {
		log.Fatal(err)
	}

	mp := mypass.New(username, password, pv)
	err = mp.Save(account.Account{
		Type:     account.Password,
		Name:     "twitter",
		Username: "emretwit",
		Password: "123",
	})
	if err != nil {
		log.Fatal(err)
	}

	accounts, err := mp.Find("twitter")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(accounts)
}
