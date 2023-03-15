## Challenge: Array Waveform.

where by we are given an array to transform in such a way that the neigbouring characters surronding the odd indicies are less than or equal to their neighboring elements at the even indices.

> Sample Input
arr := []int{8, 1, 2, 3, 4, 5, 6, 4, 2} <br />
k = 2

> Sample Output

array = [ 8, 1, 3, 2, 5, 4, 6, 2, 4 ] Or array = [ 2, 1, 3, 2, 4, 4, 6, 5, 8 ]

> Explanation

**My thought process towards solving this algorithm was**
- iterate over the odd character in the array (for optimization and easy comparison), made comparison with it negbouring characters (i.e the one before it and the one after it) to know which character to swap with
- if the index value of the odd character in the array is greater than the character that comes before it (arr[i] > arr[i-1]), then make a swap with the character that comes before it (arr[i-1]). provided its within the bounds of the array such that i - 1 >= 0.
- if the index value of the odd character in the array is greater than the character that comes after it (arr[i] > arr[i+1]), then make a swap with the character that comes afterit(arr[i+1]). provided its within the bounds of the array such that i <= len(arr).


######  Array Waveform.
 [source](https://www.educative.io/courses/data-structures-and-algorithms-go/g7WjBYrp2V6)
