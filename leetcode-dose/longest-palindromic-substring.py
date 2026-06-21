# def longestPalindrome(s):
#     '''
#     :type s: str
#     :rtype: str
#     '''
#     if not s:
#         return ""
#     # sliding window
#     longest_palindrome = s[0]
#     for i in range(len(s)):
#         window = ""

#         for j in range(i, len(s)):
#             window += s[j]
#             if isPalindrome(window) and len(window) > len(longest_palindrome):
#                 longest_palindrome = window

#     return longest_palindrome

# def isPalindrome(s):
#     return s == s[::-1]

def longestPalindrome(s):
    '''
    :type s: str
    :rtype: str
    '''
    result = ""

    for i in range(len(s)):
        # odd case
        l, r = i, i

        while l >= 0 and r < len(s) and s[l] == s[r]:
            if len(result) < (r - l + 1):
                result = s[l:r+1]
            
            l -= 1
            r += 1

        # even case
        l, r = i, i + 1

        while l >= 0 and r < len(s) and s[l] == s[r]:
            if len(result) < (r - l + 1):
                result = s[l:r+1]
            
            l -= 1
            r += 1

    return result


if __name__ == '__main__':
    print(longestPalindrome("babad"))
    print(longestPalindrome("cbbd"))
    print(longestPalindrome("a"))
