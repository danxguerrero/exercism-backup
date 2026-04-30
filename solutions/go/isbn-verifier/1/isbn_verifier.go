package isbnverifier
import (
    "strings"
    "unicode/utf8"
)

func IsValidISBN(isbn string) bool {
    isbn = strings.ReplaceAll(isbn, "-","")
    if utf8.RuneCountInString(isbn) != 10 {
        return false
    }

    sum := 0

    check := 10

    for i := 0; i < len(isbn); i++ {
        char := isbn[i]

        var num int
        if char == 'X' {
            if check != 1 {
                return false
            }
            num = 10
        } else if char >= '0' && char <= '9' {
        	num = int(char - '0')
        } else {
            return false
        }

    	sum += num * check
        check--

        if check < 0 { return false }
    }

    return sum % 11 == 0
}
