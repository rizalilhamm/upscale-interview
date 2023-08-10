package main

import (
	"fmt"
	"strconv"
	"strings"
)

func convertTo24HourFormat(time12 string) (string, error) {
	time12 = strings.ToUpper(time12)
	var suffix string
	if strings.HasSuffix(time12, "AM") {
		suffix = "AM"
	} else if strings.HasSuffix(time12, "PM") {
		suffix = "PM"
	} else {
		return "", fmt.Errorf("invalid time format")
	}

	time12 = strings.TrimSuffix(time12, "AM")
	time12 = strings.TrimSuffix(time12, "PM")
	parts := strings.Split(time12, ":")
	if len(parts) != 2 {
		return "", fmt.Errorf("invalid time format")
	}

	hour, err := strconv.Atoi(parts[0])
	if err != nil {
		return "", fmt.Errorf("invalid hour")
	}

	minute, err := strconv.Atoi(parts[1])
	if err != nil {
		return "", fmt.Errorf("invalid minute")
	}

	if suffix == "PM" && hour != 12 {
		hour += 12
	} else if suffix == "AM" && hour == 12 {
		hour = 0
	}

	return fmt.Sprintf("%02d:%02d", hour, minute), nil
}

func main() {
	var time12 string
	fmt.Print("Enter time in 12-hour format (e.g., 02:30 PM): ")
	fmt.Scan(&time12)

	time24, err := convertTo24HourFormat(time12)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Time in 24-hour format:", time24)
}
