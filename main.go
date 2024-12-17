func groupAnagrams(strs []string) [][]string {
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