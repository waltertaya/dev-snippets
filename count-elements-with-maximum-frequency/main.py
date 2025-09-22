class Solution(object):
    def maxFrequencyElements(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        frequency = {}
        totalMax, maxFrequency = 0, 0

        for element in nums:
            if element in frequency:
                frequency[element] += 1
            else:
                frequency[element] = 1
            
            maxFrequency = frequency[element] if frequency[element] > maxFrequency else maxFrequency
        
        for key, val in frequency.items():
            if val == maxFrequency:
                totalMax += val
        
        return totalMax


if __name__ == '__main__':
    soln = Solution()

    # testcase
    nums = [1,2,2,3,1,4]
    print(soln.maxFrequencyElements(nums)) # 4

    nums = [1,2,3,4,5]
    print(soln.maxFrequencyElements(nums)) # 5