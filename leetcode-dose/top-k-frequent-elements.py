def topKFrequent(nums, k):
    """
    :type nums: List[int]
    :type k: int
    :rtype: List[int]
    """
    from collections import Counter
    import heapq

    hashmap = Counter(nums)
    heap = []
    result = []

    for num, count in hashmap.items():
        heapq.heappush(heap, (-count, num))
    
    for _ in range(k):
        result.append(heapq.heappop(heap)[1])
        # heapq.heapify(heap)

    return result


if __name__ == '__main__':
    print(topKFrequent([1,1,1,2,2,3], 2))
    print(topKFrequent([1], 1))
    print(topKFrequent([1,2,1,2,1,2,3,1,3,2], 2))
