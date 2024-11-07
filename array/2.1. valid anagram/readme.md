### Challenge: Contains Duplicate

> Problem
Given two strings s and t, return true if t is an anagram of s, and false otherwise.
___

> Explanation

**intial thought process**

- return false, if length of both string are not equal
- create a map to count character occurrences
- increment counts for string s
- decrement counts for string t
- check if all counts are zero


**optimized thought process**
- return false, if length of both string are not equal
- Create a fixed-size array (26 elements for 'a' to 'z')
- Loop through the first string and count occurrences of each character
- Loop through the second string and decrement counts
    - If any count goes negative, strings are not anagrams


[source](https://leetcode.com/problems/valid-anagram/)
