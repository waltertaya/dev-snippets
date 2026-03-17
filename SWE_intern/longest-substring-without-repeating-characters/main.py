def lengthOfLongestSubstring(s):
    char_set = set()
    left = 0
    max_length = 0

    for right in range(len(s)):
        while s[right] in char_set:
            char_set.remove(s[left])
            left += 1
        
        char_set.add(s[right])

        max_length = max(max_length, right - left + 1)

    return max_length

s = "abcabcbb" # 3
print(lengthOfLongestSubstring(s))
s = "bbbbb" # 1
print(lengthOfLongestSubstring(s))
s = "pwwkew" # 3
print(lengthOfLongestSubstring(s))
s = "" # 0
print(lengthOfLongestSubstring(s))
s = "a" # 1
print(lengthOfLongestSubstring(s))
s = "abba" # 2
print(lengthOfLongestSubstring(s))
s = "dvdf" # 3
print(lengthOfLongestSubstring(s))
