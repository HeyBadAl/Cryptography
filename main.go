// main.go
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello Encryption")
	const masterKey = "kjhgfdsaqwertyuioplkjhgfdsaqwert"
	const iv = "1234567812345678"

	test(masterKey, iv, "thePasswordOnMyLuggage")
}

