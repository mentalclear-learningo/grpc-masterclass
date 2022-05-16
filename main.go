package main

import "fmt"

func main() {
	var sum int32
	count := 0

	nums := []int32{3, 5, 9, 54, 23}
	for _, num := range nums {
		sum += num
		count++
	}

	average := float64(sum) / float64(count)
	fmt.Println("Average:", average)

}
