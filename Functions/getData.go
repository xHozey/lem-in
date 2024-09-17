package ants

import (
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

var AntsRepresentaion []string

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
		log.Fatal(err)
	}
	if len(data) == 0 {
		log.Fatal("ERROR: invalid data format")
	}
	return string(data)
}

func HandleData(data string) ([]string, string, string) {
	slicedData := strings.Split(data, "\n")
	ants, err := strconv.Atoi(slicedData[0])
	if err != nil {
		log.Fatal("ERROR: invalid data format")
	}
	roomsOn := false
	linksOn := false
	startRoom := ""
	endRoom := ""
	var roomsAndLinks []string
	for i := 0; i < len(slicedData); i++ {
		if slicedData[i] != "" {
			if slicedData[i] == "##start" {
				if !roomsOn {
					roomsOn = true
					if i < len(slicedData)-1 && slicedData[i+1] != "" {

						tempData := strings.Split(slicedData[i+1], " ")
						startRoom = tempData[0]
						_, err := strconv.Atoi(tempData[1])
						if err != nil {
							log.Fatal("ERROR: invalid data format")
						}
						_, err = strconv.Atoi(tempData[2])
						if err != nil {
							log.Fatal("ERROR: invalid data format")
						}
					} else {
						log.Fatal("ERROR: invalid data format")
					}

				} else {
					log.Fatal("ERROR: invalid data format")
				}
			}
			if slicedData[i] == "##end" {
				if !linksOn {
					linksOn = true
					if i < len(slicedData)-1 && slicedData[i+1] != "" {
						tempData := strings.Split(slicedData[i+1], " ")
						endRoom = tempData[0]
						_, err := strconv.Atoi(tempData[1])
						if err != nil {
							log.Fatal("ERROR: invalid data format")
						}
						_, err = strconv.Atoi(tempData[2])
						if err != nil {
							log.Fatal("ERROR: invalid data format")
						}
					} else {
						log.Fatal("ERROR: invalid data format")
					}
				} else {
					log.Fatal("ERROR: invalid data format")
				}
			}
			roomsAndLinks = append(roomsAndLinks, slicedData[i])
		} else {
			log.Fatal("ERROR: invalid data format")
		}
	}
	if !roomsOn || !linksOn {
		log.Fatal("ERROR: invalid data format")
	}
	for i := 1; i <= ants; i++ {
		AntsRepresentaion = append(AntsRepresentaion, "L"+strconv.Itoa(i))
	}

	return roomsAndLinks, startRoom, endRoom
}

func GetRoomsAndLinks(s []string) ([]string, []string, []string) {
	var linksFrom []string
	var linksTo []string
	var rooms []string
	lockRooms := false
	for _, val := range s {
		if !lockRooms && strings.Contains(val, " ") {
			var room []string
			val = strings.TrimSpace(val)
			room = strings.Split(val, " ")
			if len(room) == 3 {
				for _, val := range room[1] {
					if val < '0' || val > '9' {
						log.Fatal("ERROR: invalid data format")
					}
				}
				for _, val := range room[2] {
					if val < '0' || val > '9' {
						log.Fatal("ERROR: invalid data format")
					}
				}
				rooms = append(rooms, room[0])
			} else {
				log.Fatal("ERROR: invalid data format")
			}
		} else if strings.Contains(val, "-") {
			lockRooms = true
			var link []string
			link = strings.Split(val, "-")
			if len(link) == 2 {
				linksFrom = append(linksFrom, link[0])
				linksTo = append(linksTo, link[1])
			} else {
				log.Fatal("ERROR: invalid data format")
			}
		}
	}
	return linksFrom, linksTo, rooms
}
