package fizzbuzz

import "strconv"

func ConvertNumberToFizzBuzz(number int) string {
	var stringNumber string
	stringNumber = strconv.Itoa(number)
	if number%3 == 0 {
		stringNumber = "Fizz"
	}
	if number%5 == 0 {
		stringNumber = "Buzz"
	}
	if number%15 == 0 {
		stringNumber = "FizzBuzz"
	}
	return stringNumber

}
