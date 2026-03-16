# heap + hashmap
import heapq

def topKFrequent(nums, k):
    count = {}
    results = []
    heap = []

    for num in nums:
        count[num] = count.get(num, 0) + 1
    
    for num, frequency in count.items():
        heapq.heappush(heap, (frequency, num))

        if len(heap) > k:
            heapq.heappop(heap)
        
    results = [num for frequency, num in heap]

    return results

nums = [1,1,1,2,2,3]
k = 2
print(topKFrequent(nums, k))

nums = [1,2,1,2,1,2,3,1,3,2]
k = 2
print(topKFrequent(nums, k))
