def longestNiceSubarray(nums):
    """
    :type nums: List[int]
    :rtype: int
    """
    left = 0  # Left pointer
    current_bits = 0  # Tracks the active bits in the window
    max_length = 0

    for right in range(len(nums)):
        # If adding nums[right] violates the "nice" condition, move left
        while (current_bits & nums[right]) != 0:
            current_bits ^= nums[left]  # Remove nums[left] from bitmask
            left += 1  # Move left pointer

        # Add the new number to the bitmask
        current_bits |= nums[right]

        # Update max_length
        max_length = max(max_length, right - left + 1)

    return max_length


# Test Cases
if __name__ == '__main__':
    tests = [
        [1, 3, 8, 48, 10],   # Expected output: 3
        [3, 1, 5, 11, 13]     # Expected output: 1
    ]
    expected_results = [3, 1]

    for i, test in enumerate(tests):
        result = longestNiceSubarray(test)
        if result == expected_results[i]:
            print(f'PASSED return = {result}: Expected = {expected_results[i]}')
        else:
            print(f'FAILED return = {result}: Expected = {expected_results[i]}')
