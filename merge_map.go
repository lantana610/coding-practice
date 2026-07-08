package main

func merge(me map[rune][]string, val map[rune][]string) map[rune][]string {
	result := map[rune][]string{}

	for i, v := range me {
		result[i] = v
	}

	for i, v := range val {
		result[i] = v
	}
	return result
}
