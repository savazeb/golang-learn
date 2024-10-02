package main

import "fmt"

func main() {
	var nums = []int{2, 7, 11, 15}
	var target = 9

	var result [2]int 
	result = twosum(nums, target)

	fmt.Println(result[0], result[1])
}

func twosum(nums []int, target int) [2]int {
	var result = [2]int{0, 0}
	var numMap = make(map[int]int)

	for i, num := range nums {
		x := target - num
		if _, ok := numMap[x]; ok {
			result[0] = numMap[x]
			result[1] = i
			return result
		}
		numMap[num] = i
	}

	return result
}

