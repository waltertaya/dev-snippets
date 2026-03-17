'''
Steps for RIGHT rotation:
- Reverse whole array
- Reverse first k elements
- Reverse remaining
'''
# time complexity:- O(n)
def rotate_right(arr, k):
    n = len(arr)
    k = k % n

    def reverse(l, r):
        while l < r:
            arr[l], arr[r] = arr[r], arr[l]
            l += 1
            r -= 1
        
    reverse(0, n-1)
    reverse(0, k-1)
    reverse(k, n-1)

    return arr


arr = [1,2,3,4,5]
k = 2
print(rotate_right(arr, k))