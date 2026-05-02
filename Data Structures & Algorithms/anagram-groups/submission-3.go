func groupAnagrams(strs []string) [][]string {
	group := make(map[[26]int][]string)

	for _, s := range strs {
		counter := [26]int{}
		for _, r := range s {
			counter[r-'a']++
		}
		group[counter] = append(group[counter], s)
	}

	output := make([][]string, 0)
	for _, v := range group {
		output = append(output, v)
	}
	return output
}
