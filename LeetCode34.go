package main

import "fmt"

func searchRange(nums []int, target int) []int {
	ret := []int{-1, -1}

	leftIndex :=extremeInsertionIndex(nums, target, true)

	if leftIndex == len(nums) || nums[leftIndex] != target {
		return ret
	}

	ret[0] = leftIndex
	ret[1] = extremeInsertionIndex(nums, target, false) - 1

	return ret

}

func extremeInsertionIndex(nums []int, target int, left bool) int {

	lo := 0
	hi := len(nums)
	for lo < hi {
		mid:= (lo+hi)/2

		if nums[mid] > target || (left && nums[mid] == target) {
			hi = mid
		} else {
			lo += 1
		}
	}

	return lo
}

func main() {
	nums := []int{5,7,7,8,8,10}; target := 8

	fmt.Println(searchRange(nums, target))
}