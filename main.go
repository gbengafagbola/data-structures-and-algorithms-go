package main

func containDuplicate(arr []int) bool {
	duplicateMap := make(map[int]int)
	for i := 0; i < len(arr); i++ {
		// if a key dosen't exist in a map we get the value type's zero
		if duplicateMap[arr[i]] == 0 {
			duplicateMap[arr[i]] = 1
		} else {
			return false
		}
	}
	return true
}

func main() {
	arr := []int{1, 2, 3, 4}

	result := containDuplicate(arr)
	println(result)
}
