func groupAnagrams(strs []string) [][]string {
	result := make(map[[26]int][]string)
	

	for _, s := range strs {
		counter := [26]int{}
		for _, c := range s {
			counter[c-'a']++
		}
		result[counter] = append(result[counter], s)
	}

	output := make([][]string, 0)
	for _, v := range result {
		output = append(output, v)
	}

	return output
}
