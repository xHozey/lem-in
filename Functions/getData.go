package ants

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// Usage Error.
func errorMessage() {
	fmt.Println("ERROR: invalid data format")
	os.Exit(0)
}

// Get data from provided file in args.
func ReadFile() string {
	args := os.Args[1:]
	if len(args) != 1 {
		log.Fatal("Enter FileName.txt")
	}
	filename := args[0]
	if filepath.Ext(filename) != ".txt" {
		log.Fatal("please enter .txt file!")
	}
	data, err := os.ReadFile(filename)
	if err != nil {
		errorMessage()
	}
	if len(data) == 0 {
		errorMessage()
	}
	return string(data)
}

// Getting needed info, rooms, links, ants.
func GetData(g *Graph) (int, string, string) {
	fileData := strings.Split(ReadFile(), "\n")
	ants, err := strconv.Atoi(fileData[0])
	if err != nil {
		errorMessage()
	}

	start, end := handleDataFile(fileData)
	g.getRooms(fileData)

	return ants, start, end
}

// function to check wrong data input with extracting the start and end
func handleDataFile(str []string) (string, string) {
	start, end := "", ""
	roomsSet, linksSet := false, false

	for i, val := range str {
		if val == "" {
			errorMessage()
		}
		if val == "##start" || val == "##end" {
			if i+1 >= len(str) || str[i+1] == "" {
				errorMessage()
			}
			room := strings.Split(str[i+1], " ")
			if len(room) != 3 {
				errorMessage()
			}
			if val == "##start" {
				if roomsSet {
					errorMessage()
				}
				getRoom(room, &start)
				roomsSet = true
			} else if val == "##end" {
				if linksSet {
					errorMessage()
				}
				getRoom(room, &end)
				linksSet = true
			}
		}
	}

	if !roomsSet || !linksSet {
		errorMessage()
	}

	return start, end
}

// Function to get room with check coordinations
func getRoom(str []string, room *string) {
	checkCoordinates(str[1])
	checkCoordinates(str[2])
	*room = str[0]
}

// Function to extract all rooms from the data
func (g *Graph) getRooms(str []string) {
	lockRooms := false

	for _, val := range str[1:] {
		if strings.HasPrefix(val, "#") || strings.HasPrefix(val, "L") {
			continue
		}
		if strings.Contains(val, " ") {
			if lockRooms {
				errorMessage()
			}
			room := strings.Split(val, " ")
			if len(room) == 3 {
				checkCoordinates(room[1])
				checkCoordinates(room[2])
				g.AddVertix(room[0])
			} else {
				errorMessage()
			}
		} else if strings.Contains(val, "-") {
			links := strings.Split(val, "-")
			if len(links) == 2 {
				lockRooms = true
				g.AddIndirectedEdge(links[0], links[1])
			} else {
				errorMessage()
			}
		} else {
			errorMessage()
		}
	}
}

// Function to check cordinations
func checkCoordinates(s string) {
	_, err := strconv.Atoi(s)
	if err != nil {
		errorMessage()
	}
}
