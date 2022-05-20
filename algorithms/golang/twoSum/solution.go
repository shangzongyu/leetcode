/**********************************************************************************
 * Source : https://leetcode.com/problems/two-sum/
 * Author : Tom Shine
 * Date   : 2018-09-08
 **********************************************************************************/

/**********************************************************************************
*
* Given an array of integers, find two numbers such that they add up to a specific target number.
*
* The function twoSum should return indices of the two numbers such that they add up to the target,
* where index1 must be less than index2. Please note that your returned answers (both index1 and index2)
* are not zero-based.
*
* You may assume that each input would have exactly one solution.
*
* Input: numbers={2, 7, 11, 15}, target=9
* Output: index1=1, index2=2
*
*
**********************************************************************************/

package main

func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)
	for index, num := range nums {
		if i, ok := numMap[target-num]; ok {
			return []int{i, index}
		}
		numMap[num] = index
	}
	return nil
}
