package main

import (
	"fmt"
	"slices"
	"sort"
)

func twoSum(nums []int, target int) (both []int) {
	closest := nums[0]
	for v := range nums {
		if nums[v] < target && nums[v] > closest {
			closest = nums[v]
		}
	}
	sumNum := target - closest
	both = append(both, slices.Index(nums, closest), slices.Index(nums, sumNum))
	sort.Ints(both)
	return both
}

func main() {
	xi := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(xi, target))
}

/*
Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
You may assume that each input would have exactly one solution, and you may not use the same element twice.
You can return the answer in any order.

Example 1:

Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].

Example 2:

Input: nums = [3,2,4], target = 6
Output: [1,2]

Example 3:

Input: nums = [3,3], target = 6
Output: [0,1]
*/
