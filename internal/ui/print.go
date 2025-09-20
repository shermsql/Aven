package ui

import "fmt"

func Header(title string) {
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println(title)
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
}

func Field(label string, value string) {
	fmt.Printf("▶ %-7s: %s\n", label, value)
}

func Output(output string) {
	fmt.Printf("▶ Output : %s\n", output)
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
}