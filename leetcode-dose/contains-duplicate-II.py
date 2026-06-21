def containsNearbyDuplicate(nums, k):
    """
    :type nums: List[int]
    :type k: int
    :rtype: bool
    """
    seen = {}

    for i in range(len(nums)):
        if nums[i] in seen and abs(seen[nums[i]] - i) <= k:
            return True
        
        seen[nums[i]] = i
        
    return False


if __name__ == '__main__':
    print(containsNearbyDuplicate([1,2,3,1], 3))
    print(containsNearbyDuplicate([1,0,1,1], 1))
    print(containsNearbyDuplicate([1,2,3,1,2,3], 2))
