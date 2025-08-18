#!/usr/bin/env python3
''' Restore IP Addresses '''

from typing import List

class Solution:
    ''' 
    A valid IP address consists of exactly 4 integers separated by a single dots.
    Each integer is between 0 and 255 (inclusive) and cannot have leading zeros.
    Given a string s containing only digits, return all possible valid IP addresses that canbe formed by inserting dots inot s.
    You are not allowed to reorder or remove any digits in s.
    You may return the valid IP in any order.
    '''   

    def restoreIpAddresses(self, s: str) -> List[str]:
        result = []

        def is_valid(segment: str) -> bool:
            if not segment:
                return False
            if segment[0] == "0" and len(segment) > 1:
                return False
            if len(segment) > 3 or int(segment) > 255:
                return False
            return True

        def backtrack(start: int, path: List[str]):
            if len(path) == 4:
                if start == len(s):
                    result.append(".".join(path))
                return

            for length in range(1, 4):
                if start + length > len(s):
                    break
                segment = s[start:start+length]
                if is_valid(segment):
                    backtrack(start + length, path + [segment])

        backtrack(0, [])
        return result
        

if __name__ == '__main__':
    ''' Test cases '''
    s = '0000'
    solution = Solution()
    print(solution.restoreIpAddresses(s))

    s = "25525511135"
    print(solution.restoreIpAddresses(s))

    s = "101023"
    print(solution.restoreIpAddresses(s))
