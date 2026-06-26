func groupAnagrams(strs []string) [][]string {
	res := make([][]string, 0)
	wordsMap := make(map[[26]int][]string)
	for _, s := range strs {
		var key [26]int
		for _, r := range s {
			key[r - 'a']++
		}
		wordsMap[key] = append(wordsMap[key], s)
	}

	for _, v := range wordsMap {
		res = append(res, v)
	}
	return res
}
