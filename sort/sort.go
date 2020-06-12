package sort

import "sort"

//sorts the characters of a word
func sortCharactersInWord(s string)string {
	byteSlice := []byte(s)
	sort.Slice(byteSlice, func(i, j int) bool {
		return byteSlice[i] < byteSlice[j]
	})
	return string(byteSlice)
}

func sortStringArray(t []string) [] string {
	sort.Slice(t, func(i, j int) bool {
		return t[i] < t[j]
	})
	return t
}
