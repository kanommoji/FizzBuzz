package fizzbuzz

import "strconv"

func ConvertNumberToFizzBuzz(number int) string {
	var stringNumber string
	fizzBuzz := []string{"Fizz", "Buzz"}
	condition := []int{3, 5}
	for index, c := range condition {
		if number%c == 0 {
			stringNumber += fizzBuzz[index]
		}
	}
	if stringNumber == "" {
		stringNumber = strconv.Itoa(number)
	}
	return stringNumber
}
