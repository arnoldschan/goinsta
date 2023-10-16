package main

import (
	"C"
)
import (
	"fmt"
	"os"

	"github.com/Davincible/goinsta/v3"
)

func login(usernamePtr string, passwordPtr string, totpPtr string) string {
	username := usernamePtr
	password := passwordPtr
	totp := totpPtr

	insta := goinsta.New(username, password, totp)
	return func() string {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("panic occurred:", err)
			}

		}()
		insta.Login()
		aa := insta.ExportConfig()
		return aa.HeaderOptions["Authorization"]
	}()
}

func main() {
	username := os.Args[1]
	password := os.Args[2]
	var totp string = ""

	if len(os.Args) == 4 {
		totp = os.Args[3]
	} else {
		totp = ""
	}
	headers := login(username, password, totp)
	fmt.Println(headers)
}
