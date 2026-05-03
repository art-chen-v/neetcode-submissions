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

	for i := 0; i < len(encoded); {
		j := i

		for encoded[j] != '#' {
			j++
		}

		strLen, _ := strconv.Atoi(encoded[i:j])
		i = j + 1
		res = append(res, encoded[i:i+strLen])
		i = i + strLen
	}
	return res
}
