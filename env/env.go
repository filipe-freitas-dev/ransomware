package env

import (
	"crypto/aes"
	"crypto/cipher"
)

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

var (
	key   = "XY(!Ni>uD]'T/WUGjUm5_(N1w9bpHJ@m"
	Key   []byte
	Gcm   cipher.AEAD
	Block cipher.Block
	Nouce []byte
	err   error
)

func Load() {
	Key = []byte(key)
	Block, err = aes.NewCipher(Key)
	handleError(err)

	Gcm, err = cipher.NewGCM(Block)
	handleError(err)

	Nouce = make([]byte, Gcm.NonceSize())
}
