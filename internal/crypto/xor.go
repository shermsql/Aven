package crypto

import (
	"fmt"
	"aven/internal/ui"
)

func XOREncrypt(text string, key byte) {
	ui.Header("🔐 [Aven] XOR Encryption")
	ui.Field("Input", text)
	ui.Field("Key", fmt.Sprintf("%d", key))

	var result string
	for i := 0; i < len(text); i++ {
		result += string(text[i] ^ key)
	}
	ui.Output(result)
}

func XORDecrypt(text string, key byte) {
	ui.Header("🔓 [Aven] XOR Decryption")
	ui.Field("Input", text)
	ui.Field("Key", fmt.Sprintf("%d", key))

	var result string
	for i := 0; i < len(text); i++ {
		result += string(text[i] ^ key)
	}
	ui.Output(result)
}