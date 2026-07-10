type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	var builder strings.Builder

	for _, s := range strs {
		builder.WriteString(strconv.Itoa(len(s)))
		builder.WriteString("#")
		builder.WriteString(s)
	}
	return builder.String()
}
// 5#hello5#world
func (s *Solution) Decode(encoded string) []string {
	res := make([]string, 0)
	l, r := 0, 0
	for r < len(encoded) {
		for encoded[r] != '#' {
			r++
		}
		strLen, _ := strconv.Atoi(encoded[l:r])
		r += 1
		res = append(res, encoded[r:r+strLen])
		r += strLen
		l = r
	}
	return res
}
