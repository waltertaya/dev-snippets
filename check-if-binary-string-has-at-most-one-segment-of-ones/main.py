class Solution:
    '''
    Check if Binary String Has at most one segment of ones
    '''
    def checkOnesSegment(self, s: str) -> bool:
        '''
        Given a binary string s ​​​​​without leading zeros, return true​​​ if s contains at most one contiguous segment of ones. Otherwise, return false.

        Example 1:

            Input: s = "1001"
            Output: false
            Explanation: The ones do not form a contiguous segment.
            
        Example 2:

            Input: s = "110"
            Output: true
        '''
        if "01" in s:
            return False
        return True


if __name__ == '__main__':
    testCases = {"1001": False, "110": True, "10": True, "1": True, "101": False}

    soln = Solution()

    for k, v in testCases.items():
        if soln.checkOnesSegment(k) == v:
            print(f"Testcase {k} passed ✅")
        else:
            print(f"Testcase {k} failed ❌")
