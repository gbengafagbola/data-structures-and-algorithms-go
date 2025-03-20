### Explaining `TournamentWinner` to a Newbie (with Visuals)  

---

### **Understanding the Problem**  
Imagine a tournament where teams play against each other, and we want to determine the overall winner based on match results.  

- Each match involves **two teams**.  
- The winning team gets **3 points**.  
- At the end, the team with the **highest score wins the tournament**.  

---

### **Code Breakdown with Visuals**  

#### **1. Function Signature**
```go
func TournamentWinner(competitions [][]string, results []int) string
```
- `competitions`: A list of matches where each match has **two teams**.  
- `results`: A list where `1` means the **first team wins**, and `0` means the **second team wins**.  
- The function **returns the winner** of the tournament.  

---

### **2. Sample Input**
```go
competitions := [][]string{
    {"TeamA", "TeamB"},
    {"TeamB", "TeamC"},
    {"TeamC", "TeamA"},
}
results := []int{1, 0, 0}
```

This means:
| Match | Winner (from `results`) |
|--------|----------------|
| TeamA vs TeamB | TeamA wins (1st team) |
| TeamB vs TeamC | TeamC wins (2nd team) |
| TeamC vs TeamA | TeamA wins (2nd team) |

---

### **3. Visualizing the Tournament**  

At the start, all teams have **0 points**:  
```
TeamA: 0
TeamB: 0
TeamC: 0
```

#### **First Match: "TeamA" vs "TeamB"**
- `results[0] == 1`, so **TeamA wins** and gets **3 points**.  

📊 **Updated Scores:**  
```
TeamA: 3
TeamB: 0
TeamC: 0
```

#### **Second Match: "TeamB" vs "TeamC"**
- `results[1] == 0`, so **TeamC wins** and gets **3 points**.  

📊 **Updated Scores:**  
```
TeamA: 3
TeamB: 0
TeamC: 3
```

#### **Third Match: "TeamC" vs "TeamA"**
- `results[2] == 0`, so **TeamA wins** and gets **3 more points**.  

📊 **Final Scores:**  
```
TeamA: 6
TeamB: 0
TeamC: 3
```
🎉 **Winner = TeamA (6 points)**  

---

### **4. Code Walkthrough with Steps**
```go
maxScore := -1
winner := ""
m := make(map[string]int)
```
✅ We initialize:  
- `maxScore` (to keep track of the highest points).  
- `winner` (to store the name of the winning team).  
- `m` (a **map** to store team scores).

---

#### **Loop through competitions**
```go
for idx, competition := range competitions {
```
👀 We go through **each match** using an index `idx`.  

---

#### **Get Teams from Match**
```go
firstTeam := competition[0]
secondTeam := competition[1]
```
🏆 Extract team names. Example:  
For `{"TeamA", "TeamB"}`, `firstTeam = "TeamA"` and `secondTeam = "TeamB"`.

---

#### **Find Winning Team**
```go
if results[idx] == 1 {
    m[firstTeam] += 3
    if m[firstTeam] > maxScore {
        maxScore = m[firstTeam]
        winner = firstTeam
    }
} else {
    m[secondTeam] += 3
    if m[secondTeam] > maxScore {
        maxScore = m[secondTeam]
        winner = secondTeam
    }
}
```
- If `results[idx] == 1` → `firstTeam` wins and gets **3 points**.  
- If `results[idx] == 0` → `secondTeam` wins and gets **3 points**.  
- Update the **maxScore** and **winner** if the team's score exceeds the current highest.  

---

### **5. Final Output**
```go
return winner
```
🎉 **Returns `"TeamA"`**, the tournament winner!  

---

### **6. Complexity Analysis**
- **Time Complexity**: 🚀 `O(N)`, since we loop through the matches once.  
- **Space Complexity**: 📦 `O(K)`, where `K` is the number of unique teams.  

---

### **Conclusion**
This function efficiently determines the winner of a tournament by:  
✅ Tracking team scores in a **map**.  
✅ Iterating through matches **once**.  
✅ Dynamically updating the **winner**.  