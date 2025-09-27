// Given an integer array nums, return true if any
// value appears at least twice in the array, and
// return false if every element is distinct.
//
// Example 1:
// Input: nums = [1, 2, 3 ,1]
// Output: true
//
// Example 2:
// Input: nums = [1, 2, 3, 4]
// Output: false
//
// Example 3:
// Input: nums = [1, 1, 1, 1, 3, 3, 3, 4, 2, 4, 2]
// Output: true
package main

func ContainsDuplicate(nums []int) bool {
	seen := make(map[int]struct{})
	for _, num := range nums {
		// Ignore the actual value and check if it exist (T/F)
		_, found := seen[num]
		if found {
			return true
		}
		// Add it to the map
		seen[num] = struct{}{}
	}
	return false
}
