package fibonacci

import "fmt"

func GetFibonacciSequence(n int) []int {
	if n <= 0 {
		return []int{}
	}
	if n == 1 {
		return []int{0}
	}
	result := []int{0, 1}
	for i := 2; i <= n-1; i++ {
		result = append(result, result[i-2]+result[i-1])
	}
	return result
}

func main() {
	fmt.Println(GetFibonacciSequence(5))
}
