## Challenge: Find the largest sum sub array.

where by looking for all possible subarrays starting from the first element and traversing the array for each possible sum combination.

arr := []int{1, -2, 3, 4, -4, 6, -14, 6, 2} <br />
ans := 9

> Explanation

The naive solution (**NaiveMaxSubArraySum function**) was a quick fix approach which shouldn’t be used even in a life & death situation but maybe an
exception were a bad answer is better than none. But the concept was a way to keep track of every iteration which was further developed in the next solution (**MaxSubArraySum**) by summing all possible subarrays to find the maximum.

For further clarification
example of expected sub-arrays are:
[1], [1,-2], [1,-2,3] ..... 

- sum := 0
>  initialized to 0 & keeps track of the maximum value of the current sub arrays

- max := 0
> nitialized to 0 & keeps track of  the maximum value across all sub arrays 

- then a foor loop is used to traverse the array to automate the sum

- if the sum is less than 0, the sum has to be reset to 0 (since we're keeping track of the largest +ve summation)

- if the sum is greater than the max we update the max

- and finally return the max at hte end of the array

**In a single scan, the max subarray in the array is discovered unlike that of the naive approach**

###### Write a method that returns largest sum subarray.

 [source](https://www.educative.io/courses/data-structures-and-algorithms-go/3Ylv3mNNkrr)
