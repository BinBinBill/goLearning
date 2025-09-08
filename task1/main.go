package main

import (
	"sort"
)

func twoSum(nums []int, target int) []int {
	numsMap := make(map[int]int)
	for i, num := range nums {
		result := target - num
		if v, exists := numsMap[result]; exists {
			return []int{v, i}
		}
		numsMap[num] = i
	}
	return nil
}
func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	merged := make([][]int, 0)
	merged = append(merged, intervals[0])
	for i := 1; i < len(intervals); i++ {
		last := merged[len(merged)-1]
		current := intervals[i]
		if current[0] <= last[1] {
			if current[1] > last[1] {
				last[1] = current[1]
			}
		} else {
			merged = append(merged, current)
		}
	}
	return merged
}

// 26. 删除有序数组中的重复项：给你一个有序数组 nums ，请你原地删除重复出现的元素，使每个元素只出现一次，返回删除后数组的新长度。
// 不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。可以使用双指针法，一个慢指针 i 用于记录不重复元素的位置，
// 一个快指针 j 用于遍历数组，当 nums[i] 与 nums[j] 不相等时，将 nums[j] 赋值给 nums[i + 1]，并将 i 后移一位
func removeDuplicates(nums []int) int {
	i := 0
	for j := 1; j < len(nums); j++ {
		if nums[i] != nums[j] {
			i++
			nums[i] = nums[j]
		}
	}
	return i + 1
}
func plusOne(digits []int) []int {
	length := len(digits)
	result := make([]int, 0)
	for i := 0; i < length; i++ {
		if i == length-1 && digits[i] == 9 {
			result = append(result, 1)
			result = append(result, 0)
		} else if i == length-1 && digits[i] != 9 {
			result = append(result, digits[i]+1)
		} else {
			result = append(result, digits[i])
		}
	}
	return result
}
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 || len(strs) == 1 {
		return ""
	}
	baseStr := strs[0]
	for i := 0; i < len(baseStr); i++ {
		for j := 1; j < len(strs); j++ {
			if baseStr[i] != strs[j][i] || i > len(strs[j]) {
				return baseStr[:i]
			}
		}
	}
	return baseStr
}

// 有效的括号
func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}
	coupleMap := map[string]string{
		"(": ")",
		"[": "]",
		"{": "}",
	}
	stack := make([]string, 0)
	for _, char := range s {
		if char == '(' || char == '[' || char == '{' {
			stack = append(stack, string(char))
		} else {
			if len(stack) == 0 {
				return false
			}
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if coupleMap[top] != string(char) {
				return false
			}
		}
	}
	return len(stack) == 0
}

// 只出现一次的数字
func singleNumber(nums []int) int {
	numMap := make(map[int]int)
	for _, num := range nums {
		numMap[num]++
	}
	for num, count := range numMap {
		if count == 1 {
			return num
		}
	}
	return -1
}

// 回文数
func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	rev := 0
	for x > rev {
		rev = rev%10 + rev*10
		x /= 10
	}

	return x == rev || x == rev/10
}

func isPalindrome1(x int) bool {
	if x < 0 {
		return false
	}
	arr := make([]int, 0)
	for x > 0 {
		arr = append(arr, x%10)
		x /= 10
	}
	arrLength := len(arr)
	for i := range arr {
		if arr[i] != arr[arrLength-i-1] {
			return false
		}
	}
	return true
}
