package fizzbuzz

import "strconv"

func ConvertNumberToFizzBuzz(number int) string {
	fizzBuzzMaps := make(map[int]string)
	fizzBuzzMaps[3] = "Fizz"
	fizzBuzzMaps[5] = "Buzz"

	var stringNumber string
	if number%3 == 0 {
		stringNumber += fizzBuzzMaps[3]
	}
	if number%5 == 0 {
		stringNumber += fizzBuzzMaps[5]
	}
	if stringNumber == "" {
		stringNumber = strconv.Itoa(number)
	}

	return stringNumber
}
