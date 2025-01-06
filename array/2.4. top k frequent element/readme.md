### Challenge: Top K Frequent Elements

**Problem**  
Given an array of integers, return the k most frequent elements.

---

### **Approaches** 

#### 1.**Count Element Frequencies**
- Use a map to store the frequency of each element in the array.
- Each key in the map represents an integer from the array, and its value is the count of occurrences of that integer.
2. **Store Key-Value Pairs in a Slice**
- Create a slice to hold the key-value pairs from the frequency map.
- Each entry in the slice is a struct containing an element and its frequency.
3. **Sort the Slice by Frequency**
- Use Go’s sort.Slice function to sort the slice in descending order based on the frequency of elements.
4.**Extract the Top K Elements**
- After sorting, iterate over the slice to extract the top k elements based on their frequencies.

- **Time Complexity**:  
- -  Counting frequencies: O(n)
- -  Sorting the slice O (n log n) where n is the number of unique elements
- - Overall: O(n + nlogn)
- -
- **Space Complexity**: O(n), for the map and result.  

> **Summary**  
- Sorting is straightforward but less efficient for larger strings due to the \(k \log k\) sorting step.  
- Frequency count avoids sorting and offers an optimal solution for grouping anagrams.  


[source](https://neetcode.io/problems/top-k-elements-in-list/)
--- 