package main

import (
	"log"
	"os"

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

	Args := os.Args[1]

	var (
		Username string
		Password string
		PV       *local.LocalProvider
	)

	if Args == "login" {
		Username := os.Args[2]
		Password := os.Args[3]
		dbPath := "/tmp/mypass.db"

		PV, err := local.New(Username, Password, dbPath)
		if err != nil {
			log.Fatal(err)
		}

	} else if Args == "save" {
		SaveUsername := os.Args[2]
		SavePass := os.Args[3]
		mp := mypass.New(Username, Password, PV)
		err2 := mp.Save(account.Account{
			Type:     account.Password,
			Name:     "uie",
			Username: SaveUsername,
			Password: SavePass,
		})
		if err2 != nil {
			log.Fatal(err2)
		}
	}

	/*

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
		fmt.Println(accounts) */
}
