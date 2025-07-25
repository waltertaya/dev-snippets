# Taking n steps to reach top (using staircase)
# You can take one or two steps every to reach top:
# Example 13:
'''
steps to solve is the permutations of every possible rearrangement

1. Start arranging the arrays starting with 2 followed by one
 formula =   n! / (k1! * k2!)
 where n is the total number of elements in the list
        k1 is the number 1's
        k2 is the number 2's
'''
import math

def permutations(n, k1, k2):
    return math.factorial(n) / (math.factorial(k1) * math.factorial(k2))


def climbStairs(n):
    """
    :type n: int
    :rtype: int
    """
    k1 = int(n/2)
    k2 = int(n%2)

    n = k1 + k2

    result = permutations(n, k1, k2)

    while k1 > 0:
        k2 += 1
        k1 -= 1
        result += permutations(n, k1, k2)

    return int(result)


if __name__ == '__main__':

    tests = {
        "1": 1,
        "2": 2,
        "3": 3,
        "4": 5,
        "5": 8,
        "6": 13,
        "7": 21,
        "13": 377
    }

    for idx, (k, v) in enumerate(tests.items()):
        soln = climbStairs(int(k))

        if soln == v:
            print(f'Test case passed for {k} = {v}')
        else:
            print(f'Test case FAILED for {k} != {soln}')
