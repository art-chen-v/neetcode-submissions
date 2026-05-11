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
// 5#hello5#world
func (s *Solution) Decode(encoded string) []string {
	res := make([]string, 0)

	i := 0
	j := 0

	for i < len(encoded) {
		for encoded[j] != '#' {
			j++
		}
		strLen, _ := strconv.Atoi(encoded[i:j])
		i = j + 1
		res = append(res, encoded[i:i+strLen])
		i = i+strLen
		j = i
	}

	return res
}
