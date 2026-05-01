import "slices"

func groupAnagrams(strs []string) [][]string {
	result := make(map[string][]string)

	for _, v := range strs {
		sortS := sortString(v)
		result[sortS] = append(result[sortS], v)
	}

	group := make([][]string, 0)
	for _, v := range result {
		group = append(group, v)
	}

	return group
}

func sortString(s string) string {
	chars := []rune(s)
	slices.Sort(chars)
	return string(chars)
}