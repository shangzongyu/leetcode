package main

import "testing"

func Test_TwoSum_5nums(t *testing.T) {
	var nums []int
	for i := 0; i < 5; i++ {
		nums = append(nums, i+1)
	}
	t.Log(twoSum(nums, 9))
}

func Test_TwoSum_50nums(t *testing.T) {
	var nums []int
	for i := 0; i < 50; i++ {
		nums = append(nums, i)
	}
	t.Log(twoSum(nums, 90))
}

func Test_TwoSum_500nums(t *testing.T) {
	var nums []int
	for i := 0; i < 500; i++ {
		nums = append(nums, i)
	}
	t.Log(twoSum(nums, 900))
}
