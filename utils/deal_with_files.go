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

// Parse data from that we get from the file
func ParseData(data string) (Data, error) {
	splited := strings.Split(data, "\n")
	if len(splited) <= 5 {
		return Data{}, errors.New("ERROR: invalid data format")
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
	startTracker := 0
	endTracker := 0
	end := false
	xAndYTracker := map[string]bool{}

	for i, line := range splited {
		if i == 0 {
			num, err := strconv.Atoi(line)
			if err != nil || num == 0 {
				return Data{}, errors.New("ERROR: invalid data format, invalid number of Ants")
			}
			parsedData.Ants = num
			continue
		}

		if (start || end) && !roomsChecker(line) {
			return Data{}, errors.New("ERROR: invalid data format")
		} else if strings.TrimSpace(line) == "##start" {
			start = true
			startTracker++
			continue
		} else if strings.TrimSpace(line) == "##end" {
			end = true
			endTracker++
			continue
		} else if commentChecker(line) {
			continue
		}

		if roomsChecker(line) {
			if links {
				return Data{}, errors.New("ERROR: invalid data format")
			}
			room := extractRoom(line)
			xy := extractXY(line)

			if xAndYTracker[xy] {
				return Data{}, errors.New("ERROR: invalid data format")
			} else {
				xAndYTracker[xy] = true
			}

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
		} else {
			return Data{}, errors.New("ERROR: invalid data format")
		}

	}

	if !links || startTracker != 1 || endTracker != 1 {
		return Data{}, errors.New("ERROR: invalid data format")
	}

	return parsedData, nil
}

func commentChecker(line string) bool {
	re := regexp.MustCompile(`^#.+`)
	return re.MatchString(strings.TrimSpace(line))
}

func roomsChecker(line string) bool {
	trimedLine := strings.TrimSpace(line)
	first := trimedLine != "" && trimedLine[0] != 'L' && trimedLine[0] != '#'

	re := regexp.MustCompile(`^\w+\s+\d+\s+\d+$`)
	return re.MatchString(strings.TrimSpace(line)) && first
}

func linksChecker(line string) bool {
	re := regexp.MustCompile(`^\w+-\w+$`)
	return re.MatchString(strings.TrimSpace(line))
}

func extractRoom(line string) string {
	return strings.Split(line, " ")[0]
}

func extractXY(line string) string {
	return strings.Join(strings.Split(line, " ")[1:], "-")
}

func extractLinks(line string) []string {
	return strings.Split(line, "-")
}
