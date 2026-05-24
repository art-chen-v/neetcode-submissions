type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	var sb strings.Builder
	for i := range strs {
		sb.WriteString(strconv.Itoa(len(strs[i])))
		sb.WriteString("#")
		sb.WriteString(strs[i])
	}
	return sb.String()
}
// 5#hello5#world
func (s *Solution) Decode(encoded string) []string {
	res := make([]string, 0)
	l := 0
	for r := 0; r < len(encoded); {
		for encoded[r] != '#' {
			r++
		}
		strLen, _ := strconv.Atoi(encoded[l:r])
		l = r + 1
		res = append(res, encoded[l:l+strLen])
		r = l + strLen
		l = r
	}
	return res
}
