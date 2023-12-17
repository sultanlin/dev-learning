package main

import (
	"fmt"
	"os"

	"sultan.com/go/files/fileutils"
)

func main() {
	rootPath, _ := os.Getwd()
	filepath := rootPath + "/data/text.txt"
	content, err := fileutils.ReadTextFile(filepath)

	if err == nil {
		fmt.Println(content)
		newContent := fmt.Sprintf("Original: %v\n Double Original: %v%v\n", content, content, content)
		fileutils.WriteToFile(filepath+".output.txt", newContent)
	} else {
		fmt.Printf("ERROR PANIC!! %v", err)
	}
}
