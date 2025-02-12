package main

import (
	"fmt"
	"sort"
)

//	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}


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

func groupAnagramsOne(strs []string) [][]string {
	// To be returned at the end
	result := [][]string{}
	// Track anagrams using a map where key is the sorted version of the string
	mapTracker := make(map[string][]string)

	for i := 0; i < len(strs); i++ {
		// Sort the current string to find an anagram group
		chars := []rune(strs[i])
		sort.Slice(chars, func(a, b int) bool {
			return chars[a] < chars[b]
		})
		sortedStr := string(chars)
		// Append the current string to the sorted key group
		mapTracker[sortedStr] = append(mapTracker[sortedStr], strs[i])
	}

	// Now, extract all the groups from the map
	for _, anagramGroup := range mapTracker {
		result = append(result, anagramGroup)
	}

	return result
}

func groupAnagramsTwo(strs []string) [][]string {
	// Map to store groups of anagrams
	anagramMap := make(map[string][]string)

	// Iterate through each string
	for _, str := range strs {
		// Convert string to a slice of characters
		chars := []rune(str)
		// Sort the characters
		sort.Slice(chars, func(i, j int) bool {
			return chars[i] < chars[j]
		})
		// Create a key for the sorted string (anagram signature)
		key := string(chars)
		// Group the anagrams by the key
		anagramMap[key] = append(anagramMap[key], str)
	}

	// Prepare result by collecting all the values from the map
	var result [][]string
	for _, group := range anagramMap {
		result = append(result, group)
	}

	return result
}

func groupAnagramsThree(strs []string) [][]string {
    // Map to store groups of anagrams based on frequency count
    anagramMap := make(map[[26]int][]string)

    // Iterate through each string in the input
    for _, str := range strs {
        // Initialize a frequency array to count occurrences of each character
        freq := [26]int{}
        
        // Count the frequency of each character in the string
        for i := 0; i < len(str); i++ {
            freq[str[i] - 'a']++
        }

        // Group the strings (anagrams) by the frequency array as the key
        anagramMap[freq] = append(anagramMap[freq], str)
    }

    // Prepare the result by collecting all the anagram groups from the map
    var result [][]string
    for _, group := range anagramMap {
        result = append(result, group)
    }

    return result
}

func groupAnagramsFour(strs []string) [][]string {
    anagramMap := map[[26]int][]string{} // Map using character frequency array as the key
    for _, word := range strs {
        freq := [26]int{}
        for i := 0; i < len(word); i++ {
            freq[word[i] - 'a'] += 1 // Counting frequency of each character
        }
        anagramMap[freq] = append(anagramMap[freq], word) // Group by frequency array
    }
    result := [][]string{}
    for _, value := range anagramMap {
        result = append(result, value) // Collect the anagram groups
    }
    return result
}

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	result := groupAnagrams(strs)
	fmt.Println(result) // Output: [["eat", "tea", "ate"], ["tan", "nat"], ["bat"]]
}
