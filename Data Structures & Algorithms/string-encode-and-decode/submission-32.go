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
	l, r := 0, 0
	res := make([]string, 0)

	for r < len(encoded) {
		for encoded[r] != '#' {
			r++
		}
		length, _ := strconv.Atoi(encoded[l:r])
		r++
		l = r
		r += length
		res = append(res, encoded[l:r])
		l = r
	}
	return res
}
