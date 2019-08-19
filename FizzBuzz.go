package fizzbuzz

import "strconv"

func ConvertNumberToFizzBuzz(number int) string {
	fizzBuzzMaps := map[int]string{
		3: "Fizz",
		5: "Buzz",
	}

	var stringNumber string
	for keyMap, fizzBuzzMap := range fizzBuzzMaps {
		if number%keyMap == 0 {
			stringNumber += fizzBuzzMap
		}
	}
	if stringNumber == "" {
		stringNumber = strconv.Itoa(number)
	}

	return stringNumber
}
