package main

import (
	"fmt"
	"sort"
)

func main() {
	results := groupingAnagram([]string{"kita", "atik", "tika", "aku", "kia", "makan", "kua"})
	fmt.Println(results)
}
func groupingAnagram(listString []string) [][]string {
	mapString := make(map[string][]string)

	for _, str := range listString {
		bytes := []byte(str)
		sort.Slice(bytes, func(i, j int) bool {
			return bytes[i] < bytes[j]
		})
		sortedStr := string(bytes)
		mapString[sortedStr] = append(mapString[sortedStr], sortedStr)
	}

	result := make([][]string, 0)
	for m := range mapString {
		result = append(result, mapString[m])
	}
	return result
}
