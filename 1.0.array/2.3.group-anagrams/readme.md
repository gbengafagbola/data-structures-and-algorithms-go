### Challenge: Group Anagrams

## Overview
This Go function efficiently groups anagrams from a list of strings using a frequency-based hashing technique.

## Code Implementation
```go
func groupAnagrams(strs []string) [][]string {
    res := make(map[[26]int][]string)

    for _, s := range strs {
        var count [26]int
        for _, c := range s {
            count[c-'a']++
        }
        res[count] = append(res[count], s)
    }

    var result [][]string
    for _, group := range res {
        result = append(result, group)
    }
    return result
}
```

## How It Works

### Input
A list of strings:
```go
strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
```

### Step-by-Step Execution
1. **Initialize a map** (`res`) to store anagrams using a 26-length array as a key.
2. **Iterate through each word**, convert it into a character frequency array.
3. **Group words by their frequency key**.
4. **Return the grouped anagrams**.

### Example Execution
| Word  | Letter Frequency Key  | Group in `res`  |
|--------|-----------------------------|-------------------|
| "eat"  | `[1,0,0,0,1,0,0,0,0,0,0,0,1,...]`  | `["eat"]`  |
| "tea"  | Same as "eat"  | `["eat", "tea"]`  |
| "tan"  | `[0,0,0,0,0,0,0,0,0,0,0,0,1,1,...]`  | `["tan"]`  |
| "ate"  | Same as "eat"  | `["eat", "tea", "ate"]`  |
| "nat"  | Same as "tan"  | `["tan", "nat"]`  |
| "bat"  | `[1,1,0,0,0,0,0,0,0,0,0,0,0,...]`  | `["bat"]`  |

### Final Output
```go
[["eat", "tea", "ate"], ["tan", "nat"], ["bat"]]
```

## Complexity Analysis
- **Time Complexity:** O(N * K), where N is the number of words and K is the maximum length of a word.
- **Space Complexity:** O(N * K) for storing groups in a map.

## Usage
You can use this function in Go projects that involve text processing, such as:
- Grouping words in dictionary applications
- Identifying anagram relationships in natural language processing
- Categorizing words based on character composition

## Contributions
Feel free to fork and enhance the implementation! 🚀

// extra


**Problem**  
Given an array of strings, group the strings that are anagrams of each other.  
An anagram is a word or phrase formed by rearranging the letters of another, using all the original letters exactly once.

---

### **Approaches**  

#### 1. **Check If Two Strings Are Anagrams**  
- Compare each string with others to verify if they are anagrams by counting the frequency of each character.  
- Use an array of size 26 (for lowercase English letters) to store character counts, incrementing for one string and decrementing for the other.  
- If the frequency array becomes invalid (negative count), the strings are not anagrams.  

- **Time Complexity**: O(n x k), where n is the number of strings and k is the average length of a string.  
- **Space Complexity**: O(1).  

---

#### 2. **Group Anagrams Using Sorted Keys**  
- Sort each string alphabetically to create a unique key for anagrams.  
- Use a map to group strings by their sorted key.  
- The sorted key ensures all anagrams are grouped together.  

- **Time Complexity**: O(n \times k \log k), due to sorting each string.  
- **Space Complexity**: O(n x k), for the map storing grouped anagrams.  

---

#### 3. **Group Anagrams Using Frequency Count**  
- Create a frequency count array (size 26) for each string as its unique key.  
- Use the frequency array to represent anagrams and group them in a map.  
- This avoids the sorting overhead by directly using character counts.  

- **Time Complexity**: O(n x k), as we count characters for each string.  
- **Space Complexity**: O(n x k), for the map storing grouped anagrams.  

---

#### 4. **Optimized Frequency Count for Keys**  
- Similar to the previous approach but uses a more compact representation of frequency arrays as map keys (e.g., arrays or encoded strings).  
- This ensures faster operations without changing the fundamental grouping logic.  

- **Time Complexity**: O(n x k), counting characters for each string.  
- **Space Complexity**: O(n x k), for the map and result.  

--- 

> **Summary**  
- Sorting is straightforward but less efficient for larger strings due to the k \log k sorting step.  
- Frequency count avoids sorting and offers an optimal solution for grouping anagrams.  


[source](https://leetcode.com/problems/group-anagrams/)


