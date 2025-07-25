def removeDuplicates(nums):
    """
    :type nums: List[int]
    :rtype: int
    """

    # for counting
    n = 1
    index_to_be_deleted = []

    # iterate through the list
    for idx, val in enumerate(nums):
        # check if the element appears more than twice and delete from the list

        if idx == 0:
            continue
        # count the number of times an element repeat (i.e n > 3, remove the 3rd or more)
        if val == nums[idx - 1]:
            print(f'Prev: {nums[idx - 1]} == Current: {val}; n is incremented')
            n = n + 1
        else:
            print(f'Prev: {nums[idx - 1]} NOT EQUAL Current: {val}; n is reset to one')
            n = 1
        
        if n > 2:
            print(f'Enter the condition "if n > 2"')
            # print(f'Deleting the element {val} of index {idx}')
            n = n - 1
            index_to_be_deleted.append(idx)
            print(f'index: {idx} will be deleted')
            # print(f'After deleting: new list {nums}')

        print(f'n: {n}')
        
    print(f'Indexes to be deleted: {index_to_be_deleted}')
        
    for i in sorted(index_to_be_deleted, reverse=True):
            del nums[i - 1]
    
    return len(nums), nums


if __name__ == '__main__':
    tests = [
        [1,1,1,2,2,3],
        [0,0,1,1,1,1,2,3,3],
        [1, 1, 1]
    ]

    ans = [
        5, 7, 2
    ]

    for idx in range(3):
        result, list_result = removeDuplicates(tests[idx])
        print(f'Result: {result}, Expected: {ans[idx]}')
        print(f'Nums output: {list_result}')
        if result == ans[idx]:
            print('PASSED')
        else:
            print('FAILED')
