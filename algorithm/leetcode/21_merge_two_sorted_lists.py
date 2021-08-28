# Definition for singly-linked list.
class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


class Solution:
    def mergeTwoLists(self, l1: ListNode, l2: ListNode) -> ListNode:
        r = ListNode()
        p = r

        a = l1
        b = l2

        while a and b:
            if a.val <= b.val:
                p.next = ListNode(a.val)
                a = a.next
            else:
                p.next = ListNode(b.val)
                b = b.next

            p = p.next

        # 合并后 l1 和 l2 最多只有一个还未被合并完，我们直接将链表末尾指向未合并完的链表即可
        p.next = a if a else b

        return r.next
