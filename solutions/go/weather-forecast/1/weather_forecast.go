// Package weather provides tools to check the weather in Goblinocus.
package weather

// CurrentCondition stores the current weather condition.
var CurrentCondition string
// CurrentLocation stores the current location in Goblinocus.
var CurrentLocation string

// Forecast will tell the current condition in the current location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
