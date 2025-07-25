def climbStairs(n):
    """
    :type n: int
    :rtype: int
    """
    if n == 1:
        return 1
    elif n == 2:
        return 2

    prev1, prev2 = 1, 2
    for _ in range(3, n + 1):
        prev1, prev2 = prev2, prev1 + prev2  # Fibonacci sequence
        print(f'Prev1 = {prev1}, Prev2 = {prev2}')

    return prev2


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

    for k, v in tests.items():
        soln = climbStairs(int(k))

        if soln == v:
            print(f'Test case passed for {k} = {v}')
        else:
            print(f'Test case FAILED for {k} != {soln}')
