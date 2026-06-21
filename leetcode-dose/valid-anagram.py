# def isAnagram(s, t):
#     """
#     :type s: str
#     :type t: str
#     :rtype: bool
#     """
#     # return sorted(list(s)) == sorted(list(t))
#     return sorted(s) == sorted(t)

from collections import Counter

def isAnagram(s, t):
    """
    :type s: str
    :type t: str
    :rtype: bool
    """
    # return sorted(list(s)) == sorted(list(t))
    return Counter(s) == Counter(t)


if __name__ == '__main__':
    print(isAnagram("anagram", "nagaram"))
    print(isAnagram("rat", "car"))
