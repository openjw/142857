package main

import "fmt"

func main() {
	fmt.Printf("%v\n", twoSum([]int{2, 7, 11, 15}, 9))
}

func twoSum(nums []int, target int) []int {
	source := map[int]int{}
	for i, v := range nums {
		source[v] = i
	}
	for i, v := range nums {
		sec_key := target - v
		if s_v, ok := source[sec_key]; ok && s_v != i {
			return []int{i, s_v}
		}
	}
	return nil
}
