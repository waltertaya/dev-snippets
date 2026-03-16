import heapq

def topKFrequent(words, k):
    count = {}
    heap = []
    results = []

    for word in words:
        count[word] = count.get(word, 0) + 1
    
    for w, f in count.items():
        heapq.heappush(heap, (-f, w))

    for _ in range(k):
        results.append(heapq.heappop(heap)[1])

    return results

words = ["i","love","leetcode","i","love","coding"]
k = 2
print(topKFrequent(words, k))

words = ["the","day","is","sunny","the","the","the","sunny","is","is"]
k = 4
print(topKFrequent(words, k))