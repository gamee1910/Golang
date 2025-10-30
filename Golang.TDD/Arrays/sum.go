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

func SumAll(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}
	return sums
}

// slice[low:high]
func SumAllTail(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:] //"take from 1 to the end" with numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}
	return sums
}
