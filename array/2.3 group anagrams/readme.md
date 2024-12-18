### Challenge: Group Anagrams

**Problem**  
Given an array of strings, group the strings that are anagrams of each other.  
An anagram is a word or phrase formed by rearranging the letters of another, using all the original letters exactly once.

---

### **Approaches**  

#### 1. **Check If Two Strings Are Anagrams**  
- Compare each string with others to verify if they are anagrams by counting the frequency of each character.  
- Use an array of size 26 (for lowercase English letters) to store character counts, incrementing for one string and decrementing for the other.  
- If the frequency array becomes invalid (negative count), the strings are not anagrams.  

- **Time Complexity**: \(O(n \times k)\), where \(n\) is the number of strings and \(k\) is the average length of a string.  
- **Space Complexity**: \(O(1)\).  

---

#### 2. **Group Anagrams Using Sorted Keys**  
- Sort each string alphabetically to create a unique key for anagrams.  
- Use a map to group strings by their sorted key.  
- The sorted key ensures all anagrams are grouped together.  

- **Time Complexity**: \(O(n \times k \log k)\), due to sorting each string.  
- **Space Complexity**: \(O(n \times k)\), for the map storing grouped anagrams.  

---

#### 3. **Group Anagrams Using Frequency Count**  
- Create a frequency count array (size 26) for each string as its unique key.  
- Use the frequency array to represent anagrams and group them in a map.  
- This avoids the sorting overhead by directly using character counts.  

- **Time Complexity**: \(O(n \times k)\), as we count characters for each string.  
- **Space Complexity**: \(O(n \times k)\), for the map storing grouped anagrams.  

---

#### 4. **Optimized Frequency Count for Keys**  
- Similar to the previous approach but uses a more compact representation of frequency arrays as map keys (e.g., arrays or encoded strings).  
- This ensures faster operations without changing the fundamental grouping logic.  

- **Time Complexity**: \(O(n \times k)\), counting characters for each string.  
- **Space Complexity**: \(O(n \times k)\), for the map and result.  

--- 

> **Summary**  
- Sorting is straightforward but less efficient for larger strings due to the \(k \log k\) sorting step.  
- Frequency count avoids sorting and offers an optimal solution for grouping anagrams.  


[source](https://leetcode.com/problems/group-anagrams/)