package generator

import (
	"math"
	"time"
)

type Schedule struct {
	Dates []onCallDay
}

type onCallDay struct {
	Date   time.Time
	Doctor Doctor
}

type Doctor struct {
	Name   string
	OffDay []time.Weekday
}

type scheduleSummary struct {
	Doctor    Doctor
	Monday    int
	Tuesday   int
	Wednesday int
	Thursday  int
	Friday    int
}

type summaryReport struct {
	Doctor    Doctor
	Monday    int
	Tuesday   int
	Wednesday int
	Thursday  int
	Friday    int
}

func (s Schedule) GenerateSchedule(startDate time.Time, endDate time.Time, doctors []Doctor) Schedule {
	for day := startDate; day.Before(endDate); day = day.AddDate(0, 0, 1) {
		if day.Weekday() == time.Saturday || day.Weekday() == time.Sunday {
			continue
		}

		tryThisDoctor := assignDoctor(s, day, filterAvailableDoctors(day.Weekday(), doctors))
		if tryThisDoctor.Name == "" {
			panic("Try to assigned no Doctor")
		}

		buildOnCallDay := onCallDay{
			Date:   day,
			Doctor: tryThisDoctor,
		}

		s.Dates = append(s.Dates, buildOnCallDay)
	}

	return s
}

func filterAvailableDoctors(weekday time.Weekday, doctors []Doctor) []Doctor {
	availableDoctors := []Doctor{}

	for _, Doctor := range doctors {
		isWorking := true
		for _, offDay := range Doctor.OffDay {
			if offDay == weekday {
				isWorking = false
				break
			}
		}

		if isWorking {
			availableDoctors = append(availableDoctors, Doctor)
		}
	}

	return availableDoctors
}

func assignDoctor(schedule Schedule, day time.Time, doctors []Doctor) Doctor {
	allSummaries := calculateSummaryForEachDoctor(schedule, doctors)

	doctorSummaryWithLowestCountForDay := scheduleSummary{
		Doctor:    Doctor{},
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

func calculateSummaryForEachDoctor(schedule Schedule, doctors []Doctor) []scheduleSummary {
	allSummaries := []scheduleSummary{}

	for _, Doctor := range doctors {
		summary := schedule.CalculateSummaryForDoctor(Doctor)
		allSummaries = append(allSummaries, summary)
	}

	return allSummaries
}

func (s Schedule) CalculateSummaryForDoctor(Doctor Doctor) scheduleSummary {
	summary := scheduleSummary{
		Doctor: Doctor,
	}

	for _, onCallDay := range s.Dates {
		if onCallDay.Doctor.Name == Doctor.Name {
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

func (s Schedule) GenerateSummaryForDoctor(Doctor Doctor) summaryReport {
	summary := summaryReport{
		Doctor: Doctor,
	}

	for _, onCallDay := range s.Dates {
		if onCallDay.Doctor.Name == Doctor.Name {
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
