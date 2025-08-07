package main
func reverseVowels(s string) string {
    i := 0
    j := len(s) - 1

    // Convert string to a slice of runes to allow mutation
    chars := []rune(s)

    vowels := map[rune]bool{
        'a': true, 'A': true,
        'e': true, 'E': true,
        'i': true, 'I': true,
        'o': true, 'O': true,
        'u': true, 'U': true,
    }

    for i < j {
        if vowels[chars[i]] && vowels[chars[j]] {
            chars[i], chars[j] = chars[j], chars[i]
            i++
            j--
        } else if !vowels[chars[i]] && vowels[chars[j]] {
            i++
        } else if vowels[chars[i]] && !vowels[chars[j]] {
            j--
        } else {
            i++
            j--
        }
    }

    return string(chars)
}

func reverseVowelsInitial(s string) string {
    i := 0
    j := len(s) - 1

    // Convert string to a slice of runes to allow mutation
    chars := []rune(s)

    for i < j {
        if isVowel(chars[i]) && isVowel(chars[j]) {
            chars[i], chars[j] = chars[j], chars[i]
            i++
            j--
        } else if !isVowel(chars[i]) && isVowel(chars[j]) {
            i++
        } else if isVowel(chars[i]) && !isVowel(chars[j]) {
            j--
        } else {
            i++
            j--
        }
    }

    return string(chars)
}

func isVowel(vowel rune) bool {
    return vowel == 'a' || vowel == 'A' || vowel == 'e' || vowel == 'E' ||
        vowel == 'i' || vowel == 'I' || vowel == 'o' || vowel == 'O' ||
        vowel == 'u' || vowel == 'U'
}



