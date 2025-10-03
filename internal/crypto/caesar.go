package crypto

import (
	"fmt"
	"unicode"
	"aven/internal/ui"
)

func CaesarEncrypt(text string, shift int) {
	ui.Header("Aven, Caesar Encryption")
	ui.Field("Input", text)
	ui.Field("Shift", fmt.Sprintf("%d", shift))

	var result string
	for _, c := range text {
		if unicode.IsLetter(c) {
			base := 'A'
			if unicode.IsLower(c) {
				base = 'a'
			}
			encrypted := (int(c)-int(base)+shift)%26 + int(base)
			result += string(rune(encrypted))
		} else {
			result += string(c)
		}
	}
	ui.Output(result)
}

func CaesarDecrypt(text string, shift int) {
	ui.Header("Aven, Caesar Decryption")
	ui.Field("Input", text)
	ui.Field("Shift", fmt.Sprintf("%d", shift))

	var result string
	for _, c := range text {
		if unicode.IsLetter(c) {
			base := 'A'
			if unicode.IsLower(c) {
				base = 'a'
			}
			decrypted := (int(c) - int(base) - shift + 26) %26 + int(base)
			result += string(rune(decrypted))
		} else {
			result += string(c)
		}
	}
	ui.Output(result)
}