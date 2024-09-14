package ants

import (
	"log"
	"os"
)

func readFile() {
	args := os.Args[1:]
	if len(args) != 1 {
		log.Fatal("Enter FileName.txt")
	}
}
