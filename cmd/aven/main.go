package main

import (
	"fmt"
	"aven/internal/args"
	"aven/internal/crypto"
	"aven/internal/ui"
)

func main() {
	Config := args.Parse()

	if Config.Mode == "" || Config.Algorithm == "" || Config.Text == "" || Config.Help {
		ui.Help()
		return
	}

	switch Config.Mode {
	case "Encrypt":
		crypto.Encrypt(Config)
	case "Decrypt":
		crypto.Decrypt(Config)
	default:
		fmt.Println("[Aven] Invalid Mode:", Config.Mode)
	}
}