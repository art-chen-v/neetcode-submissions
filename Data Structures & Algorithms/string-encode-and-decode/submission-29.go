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

// "5#Hello5#World"
func (s *Solution) Decode(encoded string) []string {
	res := make([]string, 0)
	l, r := 0, 0
	for r < len(encoded) {
		for encoded[r] != '#' {
			r++
			continue
		}
		length, _ := strconv.Atoi(encoded[l:r])
		l = r + 1
		r = l + length
		res = append(res, encoded[l:r])
		l = r
	}
	return res
}
