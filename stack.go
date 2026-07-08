package main

func stackTwo(top []string, bottom []string) []string {
	result := make([]string, len(top)+len(bottom))

	copy(result[:len(top)], top)
	copy(result[len(top):], bottom)
	return  result

}
func stackAll(str [][]string) []string {
	result := []string{}

	for _, j := range str {
		stackTwo(result, j)
	}
	return result
}
