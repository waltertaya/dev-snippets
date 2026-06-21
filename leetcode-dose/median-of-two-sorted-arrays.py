def findMedianSortedArrays(nums1, nums2):
    '''
    :type nums1: List[int]
    :type nums2: List[int]
    :rtype: float
    '''
    mergedArray = nums1 + nums2
    mergedArray.sort()

    l = len(mergedArray)
    if l % 2 != 0:
        return mergedArray[l // 2]
    else:
        mid1 = mergedArray[l // 2]
        mid2 = mergedArray[(l - 1) // 2]
        return (float(mid1 + mid2) / 2)


if __name__ == '__main__':
    print(findMedianSortedArrays([1, 3], [2]))
    print(findMedianSortedArrays([1, 2], [3, 4]))
