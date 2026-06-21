class MinStack:
    def __init__(self):
        self.stack = []
        self.min_stack = []

    def push(self, val: int) -> None:
        """
        :type value: int
        :rtype: None
        """
        self.stack.append(val)
        # Push the new minimum or the current minimum
        new_min = val if not self.min_stack else min(self.min_stack[-1], val)
        self.min_stack.append(new_min)

    def pop(self) -> None:
        """
        :rtype: None
        """
        if self.stack:
            self.stack.pop()
            self.min_stack.pop()

    def top(self) -> int:
        """
        :rtype: int
        """
        return self.stack[-1]

    def getMin(self) -> int:
        """
        :rtype: int
        """
        if self.min_stack:
            return self.min_stack[-1]


if __name__ == '__main__':
    '''["MinStack","push","push","push","getMin","pop","top","getMin"]
    [[],[-2],[0],[-3],[],[],[],[]]
    '''
    minStack = MinStack()
    minStack.push(-2)
    minStack.push(0)
    minStack.push(-3)
    print(minStack.getMin())
    minStack.pop()
    print(minStack.top())
    print(minStack.getMin())
