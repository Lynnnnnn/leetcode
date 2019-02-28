class Solution(object):
    def containsDuplicate(self, nums):
        """
        :type nums: List[int]
        :rtype: bool
        """
        s = set(nums)

        return (len(s) != len(nums))

r = Solution().containsDuplicate([1, 2, 3])
print r