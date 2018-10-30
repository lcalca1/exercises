#!/usr/local/bin/python3
# coding=utf-8

# Description: 
'''
Given a string containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:
    1. Open brackets must be closed by the same type of brackets.
    2. Open brackets must be closed in the correct order.
Note that an empty string is also considered valid.
'''

class Solution:
    def __init__( self ):
        '''
        constructor funciton 
        '''
        self._parenthesis_match = {
            '}' : '{',
            ']' : '[',
            ')' : '(',
        }

        self._parenthesis_num = {
            '{' : 0,
            '[' : 0,
            '(' : 0,
        }

        self._present_parenthesis = -1

    def isValid( self, s ):
        '''
        :type s: str
        :rtype: bool
        '''
        for each_s in s:
            key_tmp = self._parenthesis_num.get( each_s )
            self._present_parenthesis = self._present_parenthesis + 1

            if ( self._parenthesis_num[ key_tmp ] < 0 ):
                match = self._parenthesis_match.get( each_s )

                if ( match != s[ self._present_parenthesis ]):
                    return False

                self._parenthesis_match[ match ] = self._parenthesis_match[ match ] - 1

                if ( self._parenthesis_match[ key_tmp1 ] < 0 ):
                    return False
            else:
                self._parenthesis_num[ key_tmp ] = self._parenthesis_num[ key_tmp ] + 1
