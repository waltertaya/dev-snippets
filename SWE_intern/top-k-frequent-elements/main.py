'''
347. Top K Frequent Elements
Solved
Medium
Topics
premium lock icon
Companies
Given an integer array nums and an integer k, return the k most frequent elements. You may return the answer in any order.

 

Example 1:

Input: nums = [1,1,1,2,2,3], k = 2

Output: [1,2]

Example 2:

Input: nums = [1], k = 1

Output: [1]

Example 3:

Input: nums = [1,2,1,2,1,2,3,1,3,2], k = 2

Output: [1,2]

 

Constraints:

1 <= nums.length <= 105
-104 <= nums[i] <= 104
k is in the range [1, the number of unique elements in the array].
It is guaranteed that the answer is unique.
 

Follow up: Your algorithm's time complexity must be better than O(n log n), where n is the array's size.
'''

def topKFrequent(nums, k):
    hashmap = {}
    results = []
    for num in nums:
        hashmap[num] = hashmap.get(num, 0) + 1

    for _ in range(k):
        largest = None
        for k, v in hashmap.items():
            if largest is None or v > hashmap[largest]:
                largest = k
        results.append(largest)
        hashmap[largest] = 0
    
    return results

nums = [1,1,1,2,2,3]
k = 2
print(topKFrequent(nums, k))

nums = [1,2,1,2,1,2,3,1,3,2]
k = 2
print(topKFrequent(nums, k))
