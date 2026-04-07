package booking

import "time"
import "fmt"

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
    layout := "1/2/2006 15:04:05"
    t, err := time.Parse(layout, date)
    if (err != nil) {
        fmt.Println("Error parsing date:", err)
    }
    return t
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	now := time.Now()
    layout := "January 2, 2006 15:04:05"
    t, err := time.Parse(layout, date)
    if (err != nil) {
        fmt.Println("Error parsing date:", err)
    }
    return t.Before(now)
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	layout := "Monday, January 2, 2006 15:04:05"
    t, err := time.Parse(layout, date)
    if (err != nil) {
        fmt.Println("Error parsing date", err)
    }
    hour := t.Hour()

    return hour >= 12 && hour <= 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	layout := "1/2/2006 15:04:05"
    t, err := time.Parse(layout, date)
    if (err != nil) {
        fmt.Println("Error parsing date", err)
    }
    year, month, day := t.Date()
    hour, minute, _ := t.Clock()
    weekday := t.Weekday()
    greeting := fmt.Sprintf("You have an appointment on %v, %v %v, %v, at %v:%v.", weekday, month, day, year, hour, minute)
    return greeting
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	now := time.Now()
    currentYear := now.Year()
    date := fmt.Sprintf("%v-09-15 00:00:00", currentYear)
    layout := "2006-01-02 15:04:05"
    t, err := time.Parse(layout, date)
    if (err != nil) {
        fmt.Println("Error parsing date:", err)
    }
    return t
}
