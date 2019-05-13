class Solution(object):
    def fourSum(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: List[List[int]]
        """
        result = []
        sum_of_2_d = {}

        nums.sort()

        for i in range(0, len(nums)):
            for j in range(i + 1, len(nums)):
                sum2 = nums[i] + nums[j]

                if sum2 in sum_of_2_d:
                    sum_of_2_d[sum2].append([i, j])
                else:
                    sum_of_2_d[sum2] = [[i, j]]


        for i in range(0, len(nums)):
            for j in range(len(nums) - 1, i + 1, -1):
                sum2 = nums[i] + nums[j]
                left = target - sum2

                if left in sum_of_2_d:
                    for item in sum_of_2_d[left]:
                        if (item[0] > i and item[1] < j):
                            temp = [nums[item[0]], nums[item[1]], nums[i], nums[j]]
                            if temp not in result:
                                result.append(temp)
        return result

Solution().fourSum([-3,-2,-1,0,0,1,2,3], 0)
