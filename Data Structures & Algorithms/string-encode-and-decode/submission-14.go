
type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	var sb strings.Builder

	for _, s := range strs {
		sb.WriteString(strconv.Itoa(len([]rune(s))))
		sb.WriteString("#")
		sb.WriteString(s)
	}

	return sb.String()
}

func (s *Solution) Decode(encoded string) []string {
	var sb strings.Builder
	result := make([]string, 0)

	runes := []rune(encoded)

	for i := 0; i < len(runes);{
		if runes[i] == '#' {
			strLen, _ := strconv.Atoi(sb.String())
			sb.Reset()
			for i = i + 1; i < len(runes) && strLen > 0;{
				sb.WriteRune(runes[i])
				strLen--
				i++
			}
			result = append(result, sb.String())
			sb.Reset()
		} else {
			sb.WriteRune(runes[i])
			i++
		}
	}

	return result
}
