# hash map + fixed sliding window

def findAnagrams(s, p):
    p_count = {}
    window = {}
    results = []

    for c in p:
        p_count[c] = p_count.get(c, 0) + 1
    
    for i in range(len(s)):
        window[s[i]] = window.get(s[i], 0) + 1
        if i >= len(p):
            # remove the left value
            left_char = s[i - len(p)]
            window[left_char] -= 1

            if window[left_char] == 0:
                del window[left_char]
        
        if window == p_count:
            start = i - len(p) + 1
            results.append(start)
        
    return results

s = "cbaebabacd"
p = "abc"
print(findAnagrams(s, p))

s = "abab"
p = "ab"
print(findAnagrams(s, p))
