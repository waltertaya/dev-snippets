def merge(nums1, m, nums2, n):
    """
    :type nums1: List[int]
    :type m: int
    :type nums2: List[int]
    :type n: int
    :rtype: None Do not return anything, modify nums1 in-place instead.
    """
    nums1 = nums1[:m]
    nums2 = nums2[:n]

    result = nums1 + nums2

    return sorted(result)


if __name__ == '__main__':

    list1 = [1, 2, 3, 0, 0, 0]
    list2 = [2,5,6]
    print(merge(list1, 3, list2, 3))
    print(merge(list1, 3, list2, 3) == [1, 2, 2, 3, 5, 6])

    list3 = [1]
    list4 = []
    print(merge(list3, 1, list4, 0))
    print(merge(list3, 1, list4, 0) == [1])

    list5 = [0]
    list6 = [1]
    print(merge(list5, 0, list6, 1))
    print(merge(list5, 0, list6, 1) == [1])
