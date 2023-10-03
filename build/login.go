package main

import (
	"C"
)
import (
	"fmt"

	"github.com/Davincible/goinsta/v3"
)

//export login
func login(usernamePtr *C.char, passwordPtr *C.char, totpPtr *C.char) *C.char {
	username := C.GoString(usernamePtr)
	password := C.GoString(passwordPtr)
	totp := C.GoString(totpPtr)

	insta := goinsta.New(username, password, totp)
	return func() *C.char {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("panic occurred:", err)
			}

		}()
		insta.Login()
		aa := insta.ExportConfig()
		return C.CString(aa.HeaderOptions["Authorization"])
	}()
}

func main() {
}
