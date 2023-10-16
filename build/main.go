package main

import (
	"C"
)
import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/Davincible/goinsta/v3"
)

func login(usernamePtr string, passwordPtr string, totpPtr string, proxy string) string {
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
		var config goinsta.ConfigFile
		conf_file, err := os.Open(username)
		if err == nil {
			byteValue, _ := ioutil.ReadAll(conf_file)
			json.Unmarshal(byteValue, &config)
			insta, _ = goinsta.ImportConfig(config)
			defer conf_file.Close()
		}
		_ = insta.SetProxy(proxy, false, true)
		insta.Login()
		insta.Export(username)
		aa := insta.ExportConfig()
		return aa.HeaderOptions["Authorization"]
	}()
}

func main() {
	var username = flag.String("user", "", "account username")
	var password = flag.String("pass", "", "account password")
	var totp = flag.String("totp", "", "account 2FA code")
	var proxy = flag.String("proxy", "", "proxy URL of account")
	flag.Parse()
	headers := login(*username, *password, *totp, *proxy)
	fmt.Println(headers)
}
