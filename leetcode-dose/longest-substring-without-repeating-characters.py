# def lengthOfLongestSubstring(s):
#     '''
#     :type s: str
#     :rtype: int
#     '''
#     # variable sliding window
#     result = 0
#     for i in range(len(s)):
#         window = s[i]
#         j = i + 1

#         while j < len(s) and s[j] not in window:
#             window += s[j]
#             j += 1

#         result = max(result, len(window))
    
#     return result

def lengthOfLongestSubstring(s):
    '''
    :type s: str
    :rtype: int
    '''
    seen = set()
    left = 0
    longest = 0

    for right in range(len(s)):
        while s[right] in seen:
            seen.remove(s[left])
            left += 1

        seen.add(s[right])
        longest = max(longest, right - left + 1)

    return longest


if __name__ == '__main__':
    print(lengthOfLongestSubstring('abcabccbb'))
    print(lengthOfLongestSubstring('bbbbb'))
    print(lengthOfLongestSubstring('pwwkew'))
