class ListNode(object):
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next



def reverseList(head):
    """
    :type head: Optional[ListNode]
    :rtype: Optional[ListNode]
    """
    curr = head
    prev = None

    while curr:
        next_temp = curr.next
        curr.next = prev
        prev = curr
        curr = next_temp
    
    return prev


if __name__ == '__main__':
    # [1,2,3,4,5]
    head = ListNode(1, None)
    head.next = ListNode(2, None)
    head.next.next = ListNode(3, None)
    head.next.next.next = ListNode(4, None)
    head.next.next.next.next = ListNode(5, None)

    result = reverseList(head)

    while result:
        print(result.val)
        result = result.next
    
    print("\n-----------\n")

    head = ListNode(1, None)
    head.next = ListNode(2, None)

    result = reverseList(head)

    while result:
        print(result.val)
        result = result.next

    print("\n-----------\n")

    head = ListNode()

    result = reverseList(head)

    while result:
        print(result.val)
        result = result.next