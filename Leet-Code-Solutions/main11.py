def majorityElement(nums):
    """
    :type nums: List[int]
    :rtype: int
    """
    majority = nums[0]
    for _,val in enumerate(set(nums)):
        if nums.count(val) > len(nums)/2:
            mojority = val
    
    return mojority


if __name__ == '__main__':
    tests = [
        [3,2,3],
        [2,2,1,1,1,2,2]
    ]
    ans = [
        3,2
    ]
    for idx, test in enumerate(tests):
        result = majorityElement(test)
        if result == ans[idx]:
            print('PASSED')
        else:
            print('FAILED')
