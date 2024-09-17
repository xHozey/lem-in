package utils

import (
	"errors"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Data struct {
	Ants  int
	Start string
	End   string
	Rooms []string
	Links [][]string
}

func CheckFileName(name string) bool {
	return strings.HasSuffix(name, ".txt")
}

func ReadFile(name string) (string, error) {
	fileData, err := os.ReadFile(name)
	if err != nil {
		return "", err
	}

	return string(fileData), nil
}

func ParseData(data string) (Data, error) {
	splited := strings.Split(data, "\n")
	if len(splited) < 5 {
		return Data{}, errors.New("file is not valid!!")
	}

	parsedData := Data{
		Ants:  0,
		Start: "",
		End:   "",
		Rooms: make([]string, 0),
		Links: make([][]string, 0),
	}
	links := false
	start := false
	end := false

	for i, line := range splited {
		if i == 0 {
			num, err := strconv.Atoi(line)
			if err != nil {
				return Data{}, err
			}
			parsedData.Ants = num
		}

		if roomsChecker(line) {
			if links {
				return Data{}, errors.New("file isn't valid!!")
			}
			room := extractRoom(line)

			if start {
				parsedData.Start = room
				start = false
			} else if end {
				parsedData.End = room
				end = false
			}

			parsedData.Rooms = append(parsedData.Rooms, room)
		} else if linksChecker(line) {
			links = true
			parsedData.Links = append(parsedData.Links, extractLinks(line))
		} else if strings.TrimSpace(line) == "##start" {
			start = true
		} else if strings.TrimSpace(line) == "##end" {
			end = true
		}

	}

	return parsedData, nil
}

func roomsChecker(line string) bool {
	re := regexp.MustCompile(`^\w+\s+\d+\s\d+$`)
	return re.MatchString(strings.TrimSpace(line))
}

func linksChecker(line string) bool {
	re := regexp.MustCompile(`^\w+-\w+$`)
	return re.MatchString(strings.TrimSpace(line))
}

func extractRoom(line string) string {
	return strings.Split(line, " ")[0]
}

func extractLinks(line string) []string {
	return strings.Split(line, "-")
}