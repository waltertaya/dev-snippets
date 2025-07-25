#!/usr/bin/env python

'''
[1,3,8,48,10]
'''

def longestNiceSubarray(nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        longest = 1

        # Base case
        if len(nums) == 1:
             return longest
        
        for i in range(len(nums)):
            for j in range(i, len(nums) - i):
                if nums[i] & nums[j]:
                     break
                for k in range(j, i + 1, -1):
                     if nums[k] & nums[j]:
                          break
                longest += 1

        return longest


if __name__ == '__main__':
     tests = [
        [1,3,8,48,10],
        [3,1,5,11,13]
     ]
     ans = [
          3, 1
     ]

     for r in range(2):
        if longestNiceSubarray(tests[r]) == ans[r]:
            print(f'PASSED return = {longestNiceSubarray(tests[r])}: Expected = {ans[r]}')
        else:
            print(f'FAILED return = {longestNiceSubarray(tests[r])}: Expected = {ans[r]}')
