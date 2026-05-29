func groupAnagrams(strs []string) [][]string {
	res := make([][]string, 0)
	dict := make(map[[26]int][]string)
	for _, s := range strs {
		var counter [26]int
		for i := range s {
			counter[s[i] - 'a']++
		}
		dict[counter] = append(dict[counter], s)
	}
	for _, v := range dict {
		res = append(res, v)
	}
	return res
}
