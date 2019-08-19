package fizzbuzz

import "testing"

func Test_ConvertNumberToFizzBuzz_Input_1_Should_Be_1(t *testing.T) {
	expected := "1"
	number := 1

	actual := ConvertNumberToFizzBuzz(number)

	if expected != actual {
		t.Errorf("Expect %v but got %v", expected, actual)
	}
}

func Test_ConvertNumberToFizzBuzz_Input_3_Should_Be_Fizz(t *testing.T) {
	expected := "Fizz"
	number := 3

	actual := ConvertNumberToFizzBuzz(number)

	if expected != actual {
		t.Errorf("Expect %v but got %v", expected, actual)
	}
}
