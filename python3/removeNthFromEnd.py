#!/usr/local/bin/python3
# coding=utf-8

# Given a linked list: 1->2->3->4->5, and n = 2,
# After removing the second node from the end, the link list becomes 1->2->3->5.
# Note that given n will always be valid.

class ListNode:
    def __init__( self, x ):
        self.val = x
        self.next = None

class Solution:
    def removeNthFromEnd( self, head, n ):
        '''
        :type head: ListNode
        :type n: int
        :rtype: ListNode
        '''
        last_node = head
        remove_node = head
        remove_node_pre = head
        sign = 1

        while ( n > 1 ):
            n -= 1
            last_node = last_node.next

        while ( last_node.next ):
            remove_node = remove_node.next
            last_node = last_node.next
            sign -= 1

            if ( sign < 0 ):
                remove_node_pre = remove_node_pre.next

        if ( remove_node == remove_node_pre ):
            remove_node_pre = remove_node_pre.next
        else:
            remove_node_pre.next = remove_node.next

        del remove_node

        return remove_node_pre


