# given two strings s and t, return true if t and is an anagram of s, and false otherwise
def anagram(s, t):
    return sorted(s) == sorted(t)
