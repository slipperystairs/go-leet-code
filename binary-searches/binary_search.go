package main

import "fmt"

// Given an array of integers nums which is sorted in ascending order, and an integer target,
// write a function to search target in nums. If target exists, then return its index. Otherwise, return -1.

// You must write an algorithm with O(log n) runtime complexity.

// Example 1:
// Input: nums = [-1,0,3,5,9,12], target = 9
// Output: 4
// Explanation: 9 exists in nums and its index is 4

// Example 2:
// Input: nums = [-1,0,3,5,9,12], target = 2
// Output: -1
// Explanation: 2 does not exist in nums so return -1
func Search(nums []int, target int) int {
	low := 0
	high := len(nums)

	for low < high {
		midpoint := (low + (high - low) / 2)
		value := nums[midpoint]
		fmt.Printf("value: %d\n", value)
		fmt.Printf("midpoint: %d\n", midpoint)

		if value == target {
			return midpoint
		} else if value > target {
			high = midpoint;
		} else {
			low = midpoint + 1;
		}

	}

	return -1
}
