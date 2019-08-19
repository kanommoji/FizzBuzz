package fizzbuzz

import "strconv"

func ConvertNumberToFizzBuzz(number int) string {
	var stringNumber string
	fizzBuzz := [2]string{"Fizz", "Buzz"}
	conditions := [2]int{3, 5}

	for index, condition := range conditions {
		if number%condition == 0 {
			stringNumber += fizzBuzz[index]
		}
	}
	if stringNumber == "" {
		stringNumber = strconv.Itoa(number)
	}
	return stringNumber
}
