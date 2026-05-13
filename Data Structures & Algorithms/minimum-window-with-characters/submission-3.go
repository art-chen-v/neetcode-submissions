func minWindow(s string, t string) string {
    window := make(map[byte]int)
    countT := make(map[byte]int)
    for i := range t {
        countT[t[i]]++
    }
    have, need := 0, len(countT)
    minLen := math.MaxInt
    res := [...]int{-1, -1}
    l := 0
    for r := range s {
        window[s[r]]++
        _, ok := countT[s[r]]
        if ok && window[s[r]] == countT[s[r]] {
            have++
        }
        for have == need {
            if (r-l+1) < minLen {
                minLen = r-l+1
                res[0], res[1] = l, r+1
            }
            window[s[l]]--
            _, ok := countT[s[l]]
            if ok && window[s[l]] < countT[s[l]] {
                have--
            }
            l++
        }
    }
    if res != [...]int{-1, -1} {
        return s[res[0]:res[1]]
    }
    return ""
}
