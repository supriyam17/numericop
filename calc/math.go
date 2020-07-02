package calc

import "errors"

// Add returns sum of integers with error
func Add(numbers ...int) int {

	sum := 0

	if len(numbers) < 2 {
		return errors.New("provide more than 2 numbers"), sum
	} else {
		for _, num := range numbers {
			sum = sum + num
		}
		return nil, sum

	}
}
