# Given an array of strings strs, group the anagrams together.
# you can return the answer in any order

def groupAnagrams(strs):
    hashmap = {}
    results = []

    for s in strs:
        if "".join(sorted(s)) in hashmap:
            hashmap["".join(sorted(s))].append(s)
        else:
            hashmap["".join(sorted(s))] = []
            hashmap["".join(sorted(s))].append(s)
    
    for k, v in hashmap.items():
        results.append(v)
    
    return results

strs = ["eat","tea","tan","ate","nat","bat"]
print(groupAnagrams(strs))

strs = [""]
print(groupAnagrams(strs))

strs = ["a"]
print(groupAnagrams(strs))
