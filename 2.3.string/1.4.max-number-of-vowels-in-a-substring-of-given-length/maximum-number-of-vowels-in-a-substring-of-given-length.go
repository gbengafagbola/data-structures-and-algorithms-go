package main

var vowels = map[byte]bool{
	'a': true,
	'e': true,
	'i': true,
	'o': true,
	'u': true,
}

func maxVowels(s string, k int) int {
	count := 0

	// Initial window
	for i := 0; i < k; i++ {
		if vowels[s[i]] {
			count++
		}
	}

	maxCount := count

	// Sliding window
	for i := k; i < len(s); i++ {
		if vowels[s[i]] {
			count++
		}
		if vowels[s[i-k]] {
			count--
		}
		if count > maxCount {
			maxCount = count
		}
	}

	return maxCount
}



func maxVowelsI(s string, k int) int {
    // Precompute vowels in a 26-length array (faster than map)
    vowels := make([]int, 26)
    vowels['a'-'a'] = 1
    vowels['e'-'a'] = 1
    vowels['i'-'a'] = 1
    vowels['o'-'a'] = 1
    vowels['u'-'a'] = 1

    count := 0

    // Count vowels in the first window
    for i := 0; i < k; i++ {
        count += vowels[s[i]-'a']
    }

    maxCount := count

    // Slide the window across the string
    for i := k; i < len(s); i++ {
        count -= vowels[s[i-k]-'a'] // remove left char
        count += vowels[s[i]-'a']   // add right char
        if count > maxCount {
            maxCount = count
        }
    }

    return maxCount
}
