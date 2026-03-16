'''

Code
Testcase
Test Result
Test Result
438. Find All Anagrams in a String
Medium
Topics
premium lock icon
Companies
Given two strings s and p, return an array of all the start indices of p's anagrams in s. You may return the answer in any order.

 

Example 1:

Input: s = "cbaebabacd", p = "abc"
Output: [0,6]
Explanation:
The substring with start index = 0 is "cba", which is an anagram of "abc".
The substring with start index = 6 is "bac", which is an anagram of "abc".
Example 2:

Input: s = "abab", p = "ab"
Output: [0,1,2]
Explanation:
The substring with start index = 0 is "ab", which is an anagram of "ab".
The substring with start index = 1 is "ba", which is an anagram of "ab".
The substring with start index = 2 is "ab", which is an anagram of "ab".
 

Constraints:

1 <= s.length, p.length <= 3 * 104
s and p consist of lowercase English letters.
'''

def findAnagrams(s, p):
    results = []
    for i in range((len(s) - len(p)) + 1):
        window = []
        for j in range(i, i + len(p)):
            window.append(s[j])
        
        if sorted(window) == sorted(p):
            results.append(i)
    
    return results

s = "cbaebabacd"
p = "abc"
print(findAnagrams(s, p))

s = "abab"
p = "ab"
print(findAnagrams(s, p))
