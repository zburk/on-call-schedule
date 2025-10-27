package main

import (
	"encoding/json"
	"fmt"
	"math"
	"time"
)

type schedule struct {
	Dates []onCallDay
}

type onCallDay struct {
	Date   time.Time
	Doctor doctor
}

type doctor struct {
	Name   string
	OffDay []time.Weekday
}

type scheduleSummary struct {
	Doctor    doctor
	Monday    int
	Tuesday   int
	Wednesday int
	Thursday  int
	Friday    int
}

type summaryReport struct {
	Doctor    doctor
	Monday    int
	Tuesday   int
	Wednesday int
	Thursday  int
	Friday    int
}

func main() {
	doctors := []doctor{
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

	schedule := schedule{}

	for day := currentDate; day.Before(nextMonth); day = day.AddDate(0, 0, 1) {
		if day.Weekday() == time.Saturday || day.Weekday() == time.Sunday {
			continue
		}

		tryThisDoctor := assignDoctor(schedule, day, filterAvailableDoctors(day.Weekday(), doctors))
		if tryThisDoctor.Name == "" {
			panic("Try to assigned no doctor")
		}

		buildOnCallDay := onCallDay{
			Date:   day,
			Doctor: tryThisDoctor,
		}

		schedule.Dates = append(schedule.Dates, buildOnCallDay)
	}

	PrettyPrintData(schedule)

	for _, doctor := range doctors {
		summary := schedule.GenerateSummaryForDoctor(doctor)
		PrettyPrintData(summary)
	}
}

func filterAvailableDoctors(weekday time.Weekday, doctors []doctor) []doctor {
	availableDoctors := []doctor{}

	for _, doctor := range doctors {
		isWorking := true
		for _, offDay := range doctor.OffDay {
			if offDay == weekday {
				isWorking = false
				break
			}
		}

		if isWorking {
			availableDoctors = append(availableDoctors, doctor)
		}
	}

	return availableDoctors
}

func assignDoctor(schedule schedule, day time.Time, doctors []doctor) doctor {
	allSummaries := calculateSummaryForEachDoctor(schedule, doctors)

	doctorSummaryWithLowestCountForDay := scheduleSummary{
		Doctor:    doctor{},
		Monday:    math.MaxInt,
		Tuesday:   math.MaxInt,
		Wednesday: math.MaxInt,
		Thursday:  math.MaxInt,
		Friday:    math.MaxInt,
	}

	for _, summary := range allSummaries {
		switch day.Weekday() {
		case time.Monday:
			if summary.Monday < doctorSummaryWithLowestCountForDay.Monday {
				doctorSummaryWithLowestCountForDay = summary
			}
		case time.Tuesday:
			if summary.Tuesday < doctorSummaryWithLowestCountForDay.Tuesday {
				doctorSummaryWithLowestCountForDay = summary
			}
		case time.Wednesday:
			if summary.Wednesday < doctorSummaryWithLowestCountForDay.Wednesday {
				doctorSummaryWithLowestCountForDay = summary
			}
		case time.Thursday:
			if summary.Thursday < doctorSummaryWithLowestCountForDay.Thursday {
				doctorSummaryWithLowestCountForDay = summary
			}
		case time.Friday:
			if summary.Friday < doctorSummaryWithLowestCountForDay.Friday {
				doctorSummaryWithLowestCountForDay = summary
			}
		}
	}

	return doctorSummaryWithLowestCountForDay.Doctor
}

func calculateSummaryForEachDoctor(schedule schedule, doctors []doctor) []scheduleSummary {
	allSummaries := []scheduleSummary{}

	for _, doctor := range doctors {
		summary := schedule.CalculateSummaryForDoctor(doctor)
		allSummaries = append(allSummaries, summary)
	}

	return allSummaries
}

func (s schedule) CalculateSummaryForDoctor(doctor doctor) scheduleSummary {
	summary := scheduleSummary{
		Doctor: doctor,
	}

	for _, onCallDay := range s.Dates {
		if onCallDay.Doctor.Name == doctor.Name {
			switch onCallDay.Date.Weekday() {
			case time.Monday:
				summary.Monday++
			case time.Tuesday:
				summary.Tuesday++
			case time.Wednesday:
				summary.Wednesday++
			case time.Thursday:
				summary.Thursday++
			case time.Friday:
				summary.Friday++
			}
		}
	}

	return summary
}

func (s schedule) GenerateSummaryForDoctor(doctor doctor) summaryReport {
	summary := summaryReport{
		Doctor: doctor,
	}

	for _, onCallDay := range s.Dates {
		if onCallDay.Doctor.Name == doctor.Name {
			switch onCallDay.Date.Weekday() {
			case time.Monday:
				summary.Monday++
			case time.Tuesday:
				summary.Tuesday++
			case time.Wednesday:
				summary.Wednesday++
			case time.Thursday:
				summary.Thursday++
			case time.Friday:
				summary.Friday++
			}
		}
	}

	return summary
}

func PrettyPrintData(data interface{}) {
	// Convert data to pretty-printed JSON.
	if prettyOutput, err := json.MarshalIndent(data, "", "  "); err == nil {
		fmt.Println(string(prettyOutput))
	} else {
		// Handle error
	}
}
