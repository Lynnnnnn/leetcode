class Solution(object):
    def isValid(self, s):
        """
        :type s: str
        :rtype: bool
        """
        if len(s) % 2 != 0:
            return False

        l = []

        for ch in s:
            list_len = len(l)

            if (list_len == 0):
                l.append(ch)
                continue

            if (l[-1] == '[' and ch == ']'):
                l.pop()
            elif (l[-1] == '{' and ch == '}'):
                l.pop()
            elif (l[-1] == '(' and ch == ')'):
                l.pop()
            else:
                l.append(ch)
        
        return len(l) == 0

print Solution().isValid('{[]}()')