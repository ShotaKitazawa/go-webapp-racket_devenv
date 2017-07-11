package main

func remove(words []string, search string) []string {
	result := []string{}
	for _, word := range words {
		if word != search {
			result = append(result, word)
		}
	}
	return result
}
