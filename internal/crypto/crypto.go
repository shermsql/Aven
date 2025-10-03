package crypto

import (
	"fmt"
	"aven/internal/args"
)

func Encrypt(Config args.Config) {
	switch Config.Algorithm {
	case "Caesar":
		CaesarEncrypt(Config.Text, 3)
	case "XOR":
		XOREncrypt(Config.Text, 'K')
	default:
		fmt.Println("Aven, Unsupported Algorithm: ", Config.Algorithm)
	}
}

func Decrypt(Config args.Config) {
	switch Config.Algorithm {
	case "Caesar":
		CaesarDecrypt(Config.Text, 3)
	case "XOR":
		XORDecrypt(Config.Text, 'K')
	default:
		fmt.Println("Aven, Unsupported Algorithm: ", Config.Algorithm)
	}
}