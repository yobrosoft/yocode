package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/yobrosoft/yocode/go/editor"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Please provide a file path to open. Usage: yocode <filepath>")
	}

	filePath, err := filepath.Abs(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		log.Fatal("File does not exist: ", filePath)
	}

	// Create a new editor instance with default configuration
	config := editor.DefaultConfig()
	ed, err := editor.New(filePath, config)
	if err != nil {
		log.Fatal("Failed to create editor: ", err)
	}

	// Run the editor
	ed.Run()
}
