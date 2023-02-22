## Challenge: Rotate an Array by K position.

where by we are given an array and asked to rotate its element by k times

> Sample Input
arr := []int{1, 2, 3, 4, 5, 6} <br />
k = 2

> Sample Output

array = {3, 4, 5, 6, 1, 2}

> Explanation

**My thought process towards solving this algorithm was**
- slice out character from the 0<sup>th</sup> to the K<sup>th</sup> and store in an array (let call this array 1)
- slice out character from the K<sup>th</sup> up to the last character in the array (lets call this array 2)
- concantinate the the two array (such as array2 + array 1)


######  Rotate an Array by K position.
 [source](https://www.educative.io/courses/data-structures-and-algorithms-go/YQ4OqmxoOKp)
