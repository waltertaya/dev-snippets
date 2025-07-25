# Non-negative integer x
# Return whole number sqrt -> the retrun should be non-negative

def mySqrt(x):
    '''
    :type x: int
    :rtype: int
    '''
    import math

    result = math.trunc(math.sqrt(x))
    return result

if __name__ == '__main__':

    tests = [
        4, 8
    ]
    for _, val in enumerate(tests):
        print(mySqrt(val))
