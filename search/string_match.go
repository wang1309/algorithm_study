package search

// bfSearch 暴力匹配算法
// m 主串，s 子串
func bfSearch(m, s string) int {
	if len(m) <= 0 || len(s) <= 0 {
		return -1
	}

	j := 0

	mLen, sLen := len(m), len(s)

	for i := 0; i <= mLen-sLen; i++ {
		for j = 0; j < sLen; j++ {
			if m[i+j] == s[j] {
				continue
			} else {
				break
			}
		}

		if j == sLen {
			return i
		}
	}

	return -1
}
