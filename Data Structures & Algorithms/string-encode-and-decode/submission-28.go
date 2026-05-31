type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	var sb strings.Builder
	for _, s := range strs {
		sb.WriteString(strconv.Itoa(len(s)))
		sb.WriteString("#")
		sb.WriteString(s)
	}
	return sb.String()
}

func (s *Solution) Decode(encoded string) []string {
	res := make([]string, 0)

	l, r := 0, 0
	for r < len(encoded) {
		for encoded[r] != '#' {
			r++
		}

		strLen, _ := strconv.Atoi(encoded[l:r])
		l = r + 1
		r = l + strLen
		res = append(res, encoded[l:r])
		l = r
	}
	return res
}
