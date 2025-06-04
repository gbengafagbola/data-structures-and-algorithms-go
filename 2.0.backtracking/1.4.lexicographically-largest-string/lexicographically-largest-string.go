package main

func answerString(word string, numFriends int) string {
	var result string

	var backtrack func(start int, parts []string)
	backtrack = func(start int, parts []string) {
		// If we've made numFriends - 1 splits, the last part is fixed
		if len(parts) == numFriends-1 {
			if start >= len(word) {
				return // can't end with empty string
			}
			parts = append(parts, word[start:])
			for _, part := range parts {
				if part > result {
					result = part
				}
			}
			return
		}

		// Try all possible split points
		for i := start + 1; i < len(word); i++ {
			part := word[start:i]
			partsCopy := append([]string{}, parts...)
			partsCopy = append(partsCopy, part)
			backtrack(i, partsCopy)
		}
	}

	backtrack(0, []string{})
	return result
}