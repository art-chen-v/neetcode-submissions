type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	var str strings.Builder

	for _, s := range strs {
		str.WriteString(strconv.Itoa(len(s)))
		str.WriteString("#")
		str.WriteString(s)
	}
	return str.String()
}

func (s *Solution) Decode(encoded string) []string {
	res := make([]string, 0)
	l, r := 0, 0
	for l < len(encoded) {
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
