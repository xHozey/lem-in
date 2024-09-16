package main

import (
	"fmt"
	"lem/utils"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 || !utils.CheckFileName(args[0]) {
		fmt.Fprintln(os.Stderr, "Usage: go run . filename.txt")
		return
	}

	fileData, err := utils.ReadFile(args[0])
	if err != nil {
		fmt.Fprintln(os.Stderr, "Please Write a valid file")
		return
	}

	fmt.Println(utils.ParseData(fileData))
}
