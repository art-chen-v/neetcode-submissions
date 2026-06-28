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

func (s *Solution) Decode(encoded string) []string {
	res := make([]string, 0)
	l := 0
	r := 0
	for r < len(encoded) {
		for encoded[r] != '#' {
			r++
		}
		length, _ := strconv.Atoi(encoded[l:r])
		l = r + 1
		r = l + length
		res = append(res, encoded[l:r])
		l = r
	}
	return res
}
