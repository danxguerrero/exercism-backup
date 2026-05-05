package phonenumber
import (
    "fmt"
	"regexp"
)

func Number(phoneNumber string) (string, error) {
	re := regexp.MustCompile("[^0-9]+")
    num := re.ReplaceAllString(phoneNumber, "")

    if len(num) == 11 {
        if num[0] != '1' {
            return "", fmt.Errorf("11 digits must start with 1")
        }
        num = num[1:]
    }

    if len(num) != 10 {
        return "", fmt.Errorf("must be 10 digits")
    }

    if num[0] < '2' || num[3] < '2' {
        return "", fmt.Errorf("invalid NANP encoding")
    }
	
    return num, nil
}

func AreaCode(phoneNumber string) (string, error) {
	number,err := Number(phoneNumber)
    if err !=nil{
        return "", err
    }
    return number[0:3], nil
}

func Format(phoneNumber string) (string, error) {
    number, err := Number(phoneNumber)
    if err != nil {
        return "", err
    }
	return fmt.Sprintf("(%s) %s-%s", number[0:3], number[3:6], number[6:]), nil
}
