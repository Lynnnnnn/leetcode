class Solution(object):
    def twoSum(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: List[int]
        """
        d = {}

        for i, v in enumerate(nums):
            left = target - v

            if left in d:
                return [i, d[left]]
            else:
                d[v] = i

print Solution().twoSum([2, 7, 11, 15], 9)