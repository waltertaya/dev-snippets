# Definition for singly-linked list.
class ListNode(object):
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next
class Solution(object):
    def deleteDuplicates(self, head):
        """
        :type head: Optional[ListNode]
        :rtype: Optional[ListNode]
        """
        # result = ListNode()
        # result.val = head.val

        # while head.next != None:
        #     if head.next.val != result.val:
        #         result.next = head.next
        #         head = head.next
            
        #     else:
        #         head = head.next
        
        # return  result

        if head == None:
            return head
        current = head
        while current.next != None:
            if current.val == current.next.val:
                current.next = current.next.next
            else:
                current = current.next

        return head
    
if __name__ == '__main__':
    head = ListNode(1)
    head.next = ListNode(1)
    head.next.next = ListNode(2)
    head.next.next.next = ListNode(3)
    head.next.next.next.next = ListNode(3)

    s = Solution()
    result = s.deleteDuplicates(head)
    while result != None:
        print(result.val)
        result = result.next
