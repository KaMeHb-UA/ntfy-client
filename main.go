package main

import (
	"bufio"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

type Event struct {
	Id string
	Time int64
	Expires int64
	Event string
	Topic string
	Message string
	Priority int8
	Title string
	Tags []string
}

var colorsByPriority = map[int8]int8 {
	1: 90,
	2: 36,
	3: 37,
	4: 33,
	5: 31,
}

func formatTime(timestamp int64) string {
	return time.Unix(timestamp, 0).Format("15:04:05, 02.01.2006")
}

func main() {
	resp, err := http.Get("https://ntfy.sh/" + os.Args[1] + "/json")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		var event Event
		json.Unmarshal(scanner.Bytes(), &event)
		if event.Event == "message" {
			color := strconv.FormatInt(int64(colorsByPriority[event.Priority]), 10)
			println("\x1b[" + color + "m" + formatTime(event.Time) + " | " + event.Message + "\x1b[0m")
		}
	}
}
