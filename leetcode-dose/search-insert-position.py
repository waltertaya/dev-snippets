def searchInsert(nums, target):
    """
    :type nums: List[int]
    :type target: int
    :rtype: int
    """
    low, high = 0, len(nums) - 1

    mid = 0
    while low <= high:
        mid = low + (high - low) // 2

        if target == nums[mid]:
            return mid
        elif target < nums[mid]:
            high = mid - 1
        else:
            low = mid + 1

    return low


if __name__ == '__main__':
    print(searchInsert([1,3,5,6], 5))
    print(searchInsert([1,3,5,6], 2))
    print(searchInsert([1,3,5,6], 7))
