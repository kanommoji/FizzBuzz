package fizzbuzz

func ConvertNumberToFizzBuzz(number int) string {
	var stringNumber string
	stringNumber = "1"
	if number%3 == 0 {
		stringNumber = "Fizz"
	}
	return stringNumber
}
