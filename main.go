package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/zburk/oncallschedule/internal/generator"
)

func main() {
	doctors := []generator.Doctor{
		{
			Name:   "Dr. Smith",
			OffDay: []time.Weekday{time.Monday, time.Wednesday},
		},
		{
			Name:   "Dr. Johnson",
			OffDay: []time.Weekday{time.Tuesday, time.Thursday},
		},
	}

	currentDate := time.Now()
	nextMonth := currentDate.AddDate(0, 1, 0)

	schedule := generator.Schedule{}.GenerateSchedule(currentDate, nextMonth, doctors)

	PrettyPrintData(schedule)

	for _, doctor := range doctors {
		summary := schedule.GenerateSummaryForDoctor(doctor)
		PrettyPrintData(summary)
	}
}

func PrettyPrintData(data interface{}) {
	// Convert data to pretty-printed JSON.
	if prettyOutput, err := json.MarshalIndent(data, "", "  "); err == nil {
		fmt.Println(string(prettyOutput))
	} else {
		// Handle error
	}
}
