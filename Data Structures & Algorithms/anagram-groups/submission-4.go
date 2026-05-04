func groupAnagrams(strs []string) [][]string {
	group := make(map[[26]int][]string)

	for _, s := range strs {
		var counter [26]int
		for _, r := range s {
			counter[r-'a']++
		}
		group[counter] = append(group[counter], s)
	}
	
	res := make([][]string, 0)

	for _, v := range group {
		res = append(res, v)
	}

	return res
}
