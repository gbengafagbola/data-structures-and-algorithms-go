### Challenge: Contains Duplicate

> Problem
Given two strings s and t, return true if t is an anagram of s, and false otherwise.
___

> Explanation

**thought process**

- return false, if length of both string are not equal
- create a counter to keep count of each element occurence in string S
- iterate over the array
- compare the generated slice of array with that of t
- at the end of the loop and no occurence return false


[source](https://leetcode.com/problems/valid-anagram/)
