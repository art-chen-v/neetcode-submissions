func groupAnagrams(strs []string) [][]string {
	res := make([][]string, 0)
	bucket := make(map[[26]int][]string)

	for _, s := range strs {
		counter := [26]int{}
		for i := range s {
			counter[s[i]-'a']++
		}
		bucket[counter] = append(bucket[counter], s)
	}
	for _, v := range bucket {
		res = append(res, v)
	}

	return res
}
