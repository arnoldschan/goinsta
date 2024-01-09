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
	"path/filepath"

	"github.com/Davincible/goinsta/v3"
)

func login(usernamePtr string, passwordPtr string, totpPtr string, proxy string, savePathPtr string, passBrowserEmulationPtr bool) string {
	username := usernamePtr
	password := passwordPtr
	totp := totpPtr
	passBrowserEmulation := passBrowserEmulationPtr
	insta := goinsta.New(username, password, totp)
	insta.SetPassBrowserEmulation(passBrowserEmulation)
	return func() string {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("panic occurred:", err)
			}

		}()
		var config goinsta.ConfigFile
		err := os.MkdirAll(savePathPtr, os.ModePerm)
		savePath := filepath.Join(savePathPtr, username)
		conf_file, err := os.Open(savePath)
		if err == nil {
			byteValue, _ := ioutil.ReadAll(conf_file)
			json.Unmarshal(byteValue, &config)
			insta, _ = goinsta.ImportConfig(config)
			defer conf_file.Close()
		}
		if proxy != "" {
			_ = insta.SetProxy(proxy, true, true)
		}
		err = insta.Login()
		if err != nil {
			panic(err)
		}
		if insta.Account == nil {
			err = insta.Login()
			insta.Export(savePath)
		}
		//  still error
		if insta.Account == nil {
			return ""
		}
		insta.Export(savePath)
		aa := insta.ExportConfig()
		return aa.HeaderOptions["Authorization"]
	}()
}

func main() {
	var username = flag.String("user", "", "account username")
	var password = flag.String("pass", "", "account password")
	var totp = flag.String("totp", "", "account 2FA code")
	var proxy = flag.String("proxy", "", "proxy URL of account")
	var path = flag.String("path", "", "proxy URL of account")
	var passBrowserEmulation = flag.Bool("skipBrowser", true, "skip browser emulation")
	flag.Parse()
	headers := login(*username, *password, *totp, *proxy, *path, *passBrowserEmulation)
	fmt.Println(headers)
}
