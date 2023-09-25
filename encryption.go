package main 

import "fmt"

func debugEncryptDecrypt(masterKey, iv, password string) {
  encryptedPassword := encrypt(password, masterKey, iv)
  fmt.Printf("Encrypted Password: %v\n", encryptedPassword)

  decryptPassword :=decrypt(encryptedPassword, masterKey, iv)
  fmt.Printf("Decrypted Password: %v\n", decryptPassword)

}

func test(masterKey, iv, password string) {
  debugEncryptDecrypt(masterKey,iv, password)
  fmt.Println("============================================")
}
