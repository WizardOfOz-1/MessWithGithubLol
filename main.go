package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Input struct {
	Year   int
	Draws  string
	Offset int
}

// ParseDate function to convert date string like "MM/DD" into a time.Time object
func ParseDate(monthDay string, year int) (time.Time, error) {
	parts := strings.Split(monthDay, "/")
	if len(parts) != 2 {
		return time.Time{}, fmt.Errorf("invalid date format")
	}

	month, err := strconv.Atoi(parts[0])
	if err != nil {
		return time.Time{}, fmt.Errorf("invalid month")
	}

	day, err := strconv.Atoi(parts[1])
	if err != nil {
		return time.Time{}, fmt.Errorf("invalid day")
	}

	// Create the time.Time object using the provided year
	date := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
	fmt.Println("Date Processed: ", date)
	// Check for valid date
	if date.Month() != time.Month(month) || date.Day() != day {
		return time.Time{}, fmt.Errorf("invalid date: %s", monthDay)
	}

	return date, nil
}

// GenerateDates function to handle both ranges and individual dates with an offset
func GenerateDates(draws string, year int, offset int) ([]time.Time, error) {
	var result []time.Time

	// Split by commas to handle multiple ranges or individual dates
	sections := strings.Split(draws, ",")
	for _, section := range sections {
		// Check if it's a range (contains a "-")
		if strings.Contains(section, "-") {
			// Split by the dash to get start and end dates
			dates := strings.Split(section, "-")
			if len(dates) != 2 {
				return nil, fmt.Errorf("invalid range format")
			}

			// Parse the start and end dates
			startDate, err := ParseDate(strings.TrimSpace(dates[0]), year)
			if err != nil {
				return nil, err
			}

			endDate, err := ParseDate(strings.TrimSpace(dates[1]), year)
			if err != nil {
				return nil, err
			}

			// Apply offset to start and end dates
			startDate = startDate.AddDate(0, 0, offset)
			endDate = endDate.AddDate(0, 0, offset)

			// Generate all dates from start to end
			for d := startDate; !d.After(endDate); d = d.AddDate(0, 0, 1) {
				result = append(result, d)
			}

		} else {
			// Handle individual dates with offset
			date, err := ParseDate(strings.TrimSpace(section), year)
			if err != nil {
				return nil, err
			}
			date = date.AddDate(0, 0, offset)
			result = append(result, date)
		}
	}

	return result, nil
}

func (i *Input) TakeInput() {
	// Take Year Input
	fmt.Print("Enter Year: ")
	fmt.Scanln(&i.Year)

	// Take Draws Input
	fmt.Print("Enter Draws (e.g., 2/5-2/9,2/11): ")
	fmt.Scanln(&i.Draws)

	// Take Offset Input
	fmt.Print("Enter Offset (e.g., 1 for next day, -1 for previous day): ")
	fmt.Scanln(&i.Offset)

	// Generate the dates from the Draws input
	generatedDates, err := GenerateDates(i.Draws, i.Year, i.Offset)
	if err != nil {
		fmt.Println("Error generating dates:", err)
		return
	}

	// Print the result
	for _, date := range generatedDates {
		formattedDate := date.Format("2006-01-02T15:04:05")
		fmt.Printf("GIT_AUTHOR_DATE=\"%s\" GIT_COMMITTER_DATE=\"%s\" git commit --allow-empty -m \"%s\"\n",
			formattedDate, formattedDate, formattedDate)
	}
}

func main() {
	var input Input
	input.TakeInput()
}
