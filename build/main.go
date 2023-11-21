package main

import (
	"C"
)
import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/Davincible/goinsta/v3"
)

func login(usernamePtr string, passwordPtr string, totpPtr string, proxy string, savePathPtr string) string {
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
		err := os.MkdirAll(savePathPtr, os.ModePerm)
		savePath := filepath.Join(savePathPtr, username)
		conf_file, err := os.Open(savePath)
		if err == nil {
			byteValue, _ := ioutil.ReadAll(conf_file)
			json.Unmarshal(byteValue, &config)
			insta, _ = goinsta.ImportConfig(config)
			defer conf_file.Close()
		}
		_ = insta.SetProxy(proxy, false, true)
		insta.Login()
		insta.Export(savePath)
		aa := insta.ExportConfig()
		return aa.HeaderOptions["Authorization"]
	}()
}

func main() {
	cookie := login("dnubsxshh", "V11U7@4D1f#J", "CSLMLT6CRMGGLI2HHCVOOQTNNGSB2GOI", "http://geonode_kfLaf4FKJX:8a2c106c-80a3-4ff0-8d96-83a92fff8c74@premium-residential.geonode.com:10004", "exports")
	fmt.Println(cookie)
}
