package arrays

//https://gobyexample.com/range

//range lets you iterate over an array.
// On each iteration, range returns two values - the index and the value. We are choosing to ignore the index value by using _

func Sum(numbers []int) int {
	sum := 0

	for _, number := range numbers {
		sum += number
	}
	return sum
}
