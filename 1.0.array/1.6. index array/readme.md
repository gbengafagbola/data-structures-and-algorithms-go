## Challenge: Index Array.

In this challenge, I was given an array of size n, containing elements from indexes 0 to n-1. All values from 0 to n-1 are present in the array, and if a value is not in the array, then -1 is there to take its place. Arrange values of the array such that value i is stored at array[i]..

> Sample Input
arr := []int{ 8, -1, 6, 1, 9, 3, 2, 7, 4, -1 } <br />
size = len(arr)   --> 9

> Sample Output

 array = [ -1, 1, 2, 3, 4, -1, 6, 7, 8, 9 ]

> Explanation

**My thought process towards solving this algorithm was**
- swap the i<sup>th</sup> element to its rightful position that is arr[i] with arr[arr[i]] provided its greater than 0
- and of course continue swapping the  i<sup>th</sup> elements to their rightful place / index
- until the i<sup>th</sup> element is in its rightful place / index or the i<sup>th</sup> element is less than 0, then we move to the next item by incrementing i (i++)


######  Index Array.
 [source](https://www.educative.io/courses/data-structures-and-algorithms-go/q2zgOlMG1Ek)
