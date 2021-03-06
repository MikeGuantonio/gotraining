package sum

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum;
}

func SumAll(numbersToSum ...[]int) []int{
	var sums []int

	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))		
	}
	return sums
}

func SumAllTails(tailsToSum ...[]int) []int {
	var sums []int

	for _, numbers := range tailsToSum {
		if len(numbers) == 0 {
			sums = append(sums, Sum(numbers))
		} else {
			sums = append(sums, Sum(numbers[1:]))
		}
	}
	return sums
}