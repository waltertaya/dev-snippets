def removeDuplicates(nums):
    """
    :type nums: List[int]
    :rtype: int
    """

    # Base case
    if len(nums) <= 2:
        return len(nums)

    # TWO pointer
    left = 2
    for right in range(2, len(nums)):
        if nums[right] != nums[left - 2]:
            nums[left] = nums[right]
            left += 1
    
    return left, nums


if __name__ == '__main__':
    tests = [
        [1,1,1,2,2,3],
        [0,0,1,1,1,1,2,3,3],
        [1, 1, 1]
    ]

    ans = [
        5, 7, 2
    ]

    for idx in range(3):
        result, list_result = removeDuplicates(tests[idx])
        print(f'Result: {result}, Expected: {ans[idx]}')
        print(f'Nums output: {list_result}')
        if result == ans[idx]:
            print('PASSED')
        else:
            print('FAILED')
