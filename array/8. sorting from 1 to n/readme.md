## Challenge: Sorting From 1 to n.

In this challenge, I was given an array of size n, containing elements from indexes 1 to n. All values from 1 to n are present in the array, and if a value is not in the array, then -1 is there to take its place. Sort the array..

> Sample Input
arr := []int{ 8, 5, 6, 1, 9, 3, 2, 7, 4, 10 } <br />
size = len(arr) --> 9

> Sample Output

 array = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]

> Explanation

**My thought process towards solving this algorithm was**
- this challenge is quite similar to that of [7. index array](https://github.com/gbengafagbola/data-structures-and-algorithms-go/tree/main/array/7.%20index%20array)
- swap the i<sup>th</sup> element to its rightful position. but in this scenerio where we are sorting from 1 to n, so the value of arr[i] rightful place would be at the index of arr[arr[i]-1].
- and of course continue swapping the  i<sup>th</sup> elements to their rightful place / index
- until the i<sup>th</sup> element is in its rightful place / index. Then we increment i.


######  Index Array.
 [source](https://www.educative.io/courses/data-structures-and-algorithms-go/Bn25887jD8N)
