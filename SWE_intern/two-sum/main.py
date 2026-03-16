# given an array and target, return indices of two numbers that sum to target
# hash map solution

def twoSums(nums, target):
    hashmap = {}

    for i, num in enumerate(nums):
        compliment = target - num

        if compliment in hashmap:
            return hashmap[compliment], i
        
        hashmap[num] = i

nums = [2, 8, 11, 15, 7]
target = 9

print(twoSums(nums, target))
