class ListNode():
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


def addTwoNumbers(l1, l2):
    '''
    :type l1: Optional[ListNode]
    :type l2: Optional[ListNode]
    :rtype: Optional[ListNode]
    '''
    result = ListNode()
    head = result
    
    carry = 0
    while l1 or l2:
        l1_value = l1.val if l1 else 0
        l2_value = l2.val if l2 else 0

        sum = l1_value + l2_value + carry
        val = sum % 10
        carry = sum // 10

        current = ListNode(val, None)
        result.next = current
        result = result.next

        l1 = l1.next if l1 else l1
        l2 = l2.next if l2 else l2
    
    result.val += carry
    
    return head.next


if __name__ == '__main__':  
    l1 = ListNode(2, None)
    l1.next = ListNode(4, None)
    l1.next.next = ListNode(3, None)

    l2 = ListNode(5, None)
    l2.next = ListNode(6, None)
    l2.next.next = ListNode(4, None)

    result = addTwoNumbers(l1, l2)

    while result is not None:
        print(result.val)
        result = result.next
