def generate(numRows):
    """
    :type numRows: int
    :rtype: List[List[int]]
    """
    result = [[1]]
    for i in range(1, numRows):
        result.append([1])
        for j in range(1, i):
            new_sum = result[i-1][j-1] + result[i-1][j]
            result[i].append(new_sum)

        result[i].append(1)
    
    return result


if __name__ == '__main__':
    print(generate(2))
    # [[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]]
    print(generate(5))
