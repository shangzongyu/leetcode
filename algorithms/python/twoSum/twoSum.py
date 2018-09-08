
# Source : https://oj.leetcode.com/problems/two-sum/
# Author : Noah Shine
# Date   : 2018-09-08

###################################################################################
# Given an array of integers, find two numbers such that they add up to a specific target number.
#
# The function twoSum should return indices of the two numbers such that they add up to the target,
# where index1 must be less than index2. Please note that your returned answers (both index1 and index2)
# are not zero-based.
#
# You may assume that each input would have exactly one solution.
#
# Input: numbers={2, 7, 11, 15}, target=9
# Output: index1=1, index2=2
#
#
###################################################################################

class Solution:
    def twoSum(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: List[int]
        """

        for value_i in nums:
            for value_j in nums:
                if value_i + value_j == target:
                    return [value_i, value_j]

        return None

if __name__ == "__main__":
    solution = Solution()
    print(solution.twoSum([1,2,3,4,5], 9))
