func groupAnagrams(strs []string) [][]string {
	res := make([][]string, 0)
	m := make(map[[26]int][]string)

	for _, s := range strs {
		counter := [26]int{}
		for _, r := range s {
			counter[r-'a']++
		}
		m[counter] = append(m[counter], s)
	}

	for _, v := range m {
		res = append(res, v)
	}

	return res
}
