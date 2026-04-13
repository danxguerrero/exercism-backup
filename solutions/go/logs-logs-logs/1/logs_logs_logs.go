package logs
import "fmt"
import "strings"
import "unicode/utf8"

// Application identifies the application emitting the given log.
func Application(log string) string {
    logTypes := map[string]string{
        "U+2757" : "recommendation",
        "U+1F50D" : "search",
        "U+2600" : "weather",
    }
	for _, char := range log {
		decimalPoint := fmt.Sprintf("%U", char)
        if _, ok := logTypes[decimalPoint]; ok {
            return logTypes[decimalPoint]
        }
    }
    return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
    newLog := strings.Map(func(r rune) rune {
        if r == oldRune {
            return newRune
        }
        return r
    }, log)
    return newLog
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	logLength := utf8.RuneCountInString(log)
    return logLength <= limit
}
