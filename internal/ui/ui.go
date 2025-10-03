package ui

import (
	"fmt"
	"os"
	"path/filepath"
)

func Help() {
	Path := filepath.Join("internal", "ui", "Help.txt")
	Content, err := os.ReadFile(Path)

	if err != nil {
		fmt.Println("Aven, Failed To Load Help Message!")
		return
	}
	fmt.Println(string(Content))
}