### Challenge: Binary Search

> Explanation

This can simply be solved by creating 3 indicators

- low/start: which indicates the starting-point/position of the first element in the array 
- high/end: which indicates the position of the ending-point/last element in the array 
- middle/median: which indicates the position of the average element in the array 

> Logic 
The logic behind this algorithm is about having a continous iteration while the condition of start(0) remains less than end(8). then making use of if/else statement to check if the data is found or not. If not found and the median value position 

example: 

value: 8
arr := []int{1,2,3,4,5,6,7,8,9}

position: 0  1  2  3  4  5  6  7  8 
          |  |  |  |  |  |  |  |  |
array:    1  2  3  4  5  6  7 [8] 9 
          |           |           |
          S           M           E


start: 0                      // (instantiated to 0)
end: len(array) - 1           // (9-1 = 8)   

median: (start + end)/2       // (0+8/2 = 4)

array[median] == value        // return true 
array[median] != value        // 5 != 8 hence false so iteration continues

since array[median] < value   // we update start to be median + 1  so start = 5 
if    array[medium] > value   // then end would be median - 1 as we move towards narrowing our search 

this process would be repeated till we find our value or start > end and retun's false stating value wasn't found.


###### Write a method that searches for a given value in a sorted array.

 [source](https://www.educative.io/courses/data-structures-and-algorithms-go/RMwj6Vo8LIE)