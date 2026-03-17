def find_odd(arr):
    ''' use xor if the odd occurence is just one in the array
    time complexity: O(n)
    space complexity: O(1)
    '''
    result = 0
    for num in arr:
        result ^= num
    return result

def find_odds(arr):
    '''
    use hash map if the odd occurences is more than one in arr
    time and space complexity: O(n)
    '''
    count = {}

    for num in arr:
        count[num] = count.get(num, 0) + 1
    
    return [k for k, v in count.items() if v % 2 != 0]

# test_cases_xor = [
#     # normal case
#     ([1, 2, 3, 2, 3, 1, 3], 3),

#     # single element
#     ([7], 7),

#     # element appears many odd times
#     ([10, 10, 10], 10),

#     # mix of numbers
#     ([4, 3, 4, 4, 4, 5, 5], 3),

#     # includes zero
#     ([0, 1, 0, 1, 0], 0),

#     # negative numbers
#     ([-1, -1, -1], -1),

#     # larger case
#     ([2, 2, 5, 5, 6, 6, 6], 6),
# ]

# test_cases_hash = [
#     # two odd elements
#     ([1, 2, 3, 1, 2, 3, 4, 5], [4, 5]),

#     # multiple odds
#     ([10, 20, 10, 30, 20, 40], [30, 40]),

#     # all elements odd
#     ([1, 2, 3], [1, 2, 3]),

#     # negative numbers
#     ([-1, -2, -2, -3], [-1, -3]),

#     # mix
#     ([5, 5, 6, 7, 7, 8], [6, 8]),
# ]

# def test_xor():
#     for arr, expected in test_cases_xor:
#         result = find_odd(arr)
#         print(f"Input: {arr} | Expected: {expected} | Got: {result}")

# def test_hash():
#     for arr, expected in test_cases_hash:
#         result = find_odds(arr)
#         print(f"Input: {arr} | Expected: {expected} | Got: {result}")

# test_xor()
# test_hash()
