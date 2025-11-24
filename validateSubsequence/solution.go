package validateSubsequence

func IsSubsequence(s string, t string) bool {
	j := 0
	if len(s) < 1 {
		return true
	}
	for i := 0; i < len(t); i++ {
		if t[i] == s[j] {
			j++
		}
		if j == len(s) {
			return true
		}
	}
	return false
}
