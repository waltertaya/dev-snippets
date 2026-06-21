def groupAnagrams(strs):
    """
    :type strs: List[str]
    :rtype: List[List[str]]
    """

    result = {}

    for s in strs:
        result.setdefault(''.join(sorted(s)), []).append(s)

    return [l for l in result.values()]


if __name__ == '__main__':
    print(groupAnagrams(["eat","tea","tan","ate","nat","bat"]))
    print(groupAnagrams([""]))
    print(groupAnagrams(["a"]))
