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
// ["we","say",":","yes","!@#$%^&*()"]
// 2#we3#say1#:3#yes10#!@#$%^&*()
func (s *Solution) Decode(encoded string) []string {
	res := make([]string, 0)

	r := 0
	for l := 0; l < len(encoded); {
		if encoded[r] != '#' {
			r++
			continue
		} 
		strL, _ := strconv.Atoi(encoded[l:r])
		l = r + 1
		res = append(res, encoded[l:l+strL])
		l = l + strL
		r = l
	}
	return res
}
