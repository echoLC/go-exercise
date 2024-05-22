package exercise

import (
	"strconv"
	"strings"
)
func IsValidISBN(isbn string) bool {
	if isbn == "" {
		return false
	}


	isbnArr := strings.Split(isbn, "-")

	numeralStr := strings.Join(isbnArr, "")
	strLen := len(numeralStr)

	if strLen != 10 {
		return false
	}

	sum := 0
	multiplyValue := 10 

	for i := 0; i < strLen; i++{
		realValue := 0
		value := string(numeralStr[i])
		if value == "X" {
			if i != strLen - 1 {
				return false
			} else {
				realValue = 10
			}
		} else {
			v, err := strconv.Atoi(value)

			if err != nil {
				return false
			}

			realValue = v
		}


		sum += realValue * multiplyValue

		multiplyValue -= 1
	}

	return sum % 11 == 0
}