def binarySearch(nums, target):
    '''
    :type nums: List[int]
    :rtype: List[int]
    '''
    low = 0
    high = len(nums) - 1

    while low <= high:
        mid = low + (high - low) // 2
        if target == nums[mid]:
            return mid
        elif target < nums[mid]:
            high = mid - 1
        else:
            low = mid + 1
    
    return -1


if __name__ == '__main__':
    print(binarySearch([2, 5, 6, 7, 8, 9, 10, 19, 27, 31], 19))
