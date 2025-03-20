package main

const HOME_TEAM_WON = 1

func TournamentWinner(competitions [][]string, results []int) string {
    if len(competitions) != len(results) {
        return ""
    }
    
    maxScore := -1
    winner := ""
    
    m := make(map[string]int)
    for idx, competition := range competitions {
        
        firstTeam := competition[0]
        secondTeam := competition[1]
        
        if results[idx] == 1 {
            m[firstTeam] += 3
            if m[firstTeam] > maxScore {
                maxScore = m[firstTeam]
                winner = firstTeam
            }
        } else if (results[idx] == 0) {
              m[secondTeam] += 3
              if m[secondTeam] > maxScore {
                maxScore = m[secondTeam]
                winner = secondTeam
            }
        }   
    }

    return winner
}

func TournamentWinner2(competitions [][]string, results []int) string {
	// Write your code here.

    final := ""
    result := make(map[string]int)    

    scores := [][]int{}
    
    for _, value := range results {
        if value != 0 {
            scores = append(scores, []int{3, 0}) 
        } else if value == 0 {
            scores = append(scores, []int{0, 3}) 
        }
    }
    
    for i := 0; i < len(competitions); i++ {
        for j := 0; j < len(competitions[i]); j++ { 
            team := competitions[i][j] 
    
            if _, ok := result[team]; ok {
                result[team] += scores[i][j] 
            } else {
                result[team] = scores[i][j]
            }
        }
    }

    ans := 0

    for key, value := range result {
        if value > ans {
            ans = value
            final = key
        }
    }
  
    return final
}
