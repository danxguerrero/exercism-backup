package luhn
import "strings"

func Valid(id string) bool {
    normalizedId := strings.ReplaceAll(id, " ", "")
    
	if len(normalizedId) < 2 || normalizedId == "0" {
        return false
    }

	runes := []rune(normalizedId)
    evenLength:= len(runes) % 2 == 0
    sum := 0

    for i := len(runes) - 1; i>=0; i-- {
        num := int(runes[i] - '0')
        if num >= 10 {
            return false
        }

        doubled := num * 2

        if evenLength {
            if i % 2 == 0 {
                if doubled > 9 {
                    num = doubled - 9
                } else {
                    num = doubled
                }
            }
        } else {
            if i % 2 != 0 {
                if doubled > 9 {
                	num = doubled - 9
                } else {
                    num = doubled
                }
            }
        }
        sum += num
	}
    return sum % 10 == 0
}
