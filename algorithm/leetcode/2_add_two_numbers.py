from typing import List


# Definition for singly-linked list.


class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None


class Solution:
    def addTwoNumbers(self, l1: ListNode, l2: ListNode) -> ListNode:
        carry = 0
        r = ListNode(0)
        cur = r
        while l1 or l2 or carry:
            carry, now = divmod(
                carry + (l1.val if l1 else 0) + (l2.val if l2 else 0), 10)
            l1 = l1.next if l1 else None
            l2 = l2.next if l2 else None
            cur.next = ListNode(now)
            cur = cur.next

        self.printListNode(r.next)
        return r.next

    def buildListNode(self, l: List[int]) -> ListNode:
        r = ListNode(l[0])
        cur = r

        for v in l:
            cur.next = ListNode(v)
            cur = cur.next

        self.printListNode(r.next)
        return r.next

    def printListNode(self, l: ListNode):
        while l:
            if l.next:
                print("{} -> ".format(l.val), end="")
            else:
                print("{}".format(l.val))
            l = l.next


if __name__ == '__main__':
    for i in [
        ([2, 4, 3], [5, 6, 4]),
        ([5], [5]),
        ([1, 2, 3], [9, 7, 6, 0, 1]),
    ]:
        s = Solution()
        a = s.buildListNode(i[0])
        b = s.buildListNode(i[1])
        s.addTwoNumbers(a, b)
        print('---')
