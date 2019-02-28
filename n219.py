class Solution(object):
    def containsNearbyDuplicate(self, nums, k):
        """
        :type nums: List[int]
        :type k: int
        :rtype: bool
        """
        d = {}
        for i, v in enumerate(nums):
            if (d.has_key(v)):
                cache_i = d.get(v)
                if (i - cache_i <= k):
                    return True
                else:
                    d[v] = i
            else:
                d[v] = i
        
        return False


r = Solution().containsNearbyDuplicate([1, 0, 1, 1], 1)
print r