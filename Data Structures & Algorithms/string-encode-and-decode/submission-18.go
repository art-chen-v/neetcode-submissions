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
	result := make([]string, 0)

		i := 0
		strLen := 0

		for j := i; j < len(encoded); {
			if encoded[j] == '#' {
				strLen, _ = strconv.Atoi(encoded[i:j])
				i = j + 1
				result = append(result, encoded[i:i+strLen])
				i = i + strLen
				j = i
			} else {
				j++
			}
		}

	return result
}
