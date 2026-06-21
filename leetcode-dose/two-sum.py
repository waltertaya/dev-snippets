# def twoSum(nums, target):
#     """
#     :type nums: List[int]
#     :type target: int
#     :rtype: List[int]
#     """
#     #create hashmap / dictionary
#     hashmap = {}
#     for idx, val in enumerate(nums):
#         hashmap.setdefault(val, []).append(idx)
    
#     for i in nums:
#         compliment = target - i

#         if compliment == i and len(hashmap.get(compliment)) > 1:
#             return [hashmap[i][0], hashmap[i][1]]

#         if compliment != i and hashmap.get(compliment):
#             return [hashmap[i][0], hashmap[compliment][0]]

def twoSum(nums, target):
    """
    :type nums: List[int]
    :type target: int
    :rtype: List[int]
    """
    #create hashmap / dictionary
    hashmap = {}
    for idx, val in enumerate(nums):
        compliment = target - val
        if compliment in hashmap:
            return [hashmap[compliment], idx]
        
        hashmap[val] = idx


if __name__ == '__main__':
    print(twoSum([3, 2, 4], 6))
    print(twoSum([11, 2, 15, 7], 9))
    print(twoSum([3, 3], 6))
