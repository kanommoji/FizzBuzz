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

func Test_ConvertNumberToFizzBuzz_Input_5_Should_Be_Buzz(t *testing.T) {
	expected := "Buzz"
	number := 5

	actual := ConvertNumberToFizzBuzz(number)

	if expected != actual {
		t.Errorf("Expect %v but got %v", expected, actual)
	}
}

func Test_ConvertNumberToFizzBuzz_Input_6_Should_Be_Fizz(t *testing.T) {
	expected := "Fizz"
	number := 6

	actual := ConvertNumberToFizzBuzz(number)

	if expected != actual {
		t.Errorf("Expect %v but got %v", expected, actual)
	}
}

func Test_ConvertNumberToFizzBuzz_Input_15_Should_Be_FizzBuzz(t *testing.T) {
	expected := "FizzBuzz"
	number := 15

	actual := ConvertNumberToFizzBuzz(number)

	if expected != actual {
		t.Errorf("Expect %v but got %v", expected, actual)
	}
}
