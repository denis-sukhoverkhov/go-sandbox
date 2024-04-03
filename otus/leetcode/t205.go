package leetcode

// 205. Isomorphic Strings
// https://leetcode.com/problems/isomorphic-strings/

func isIsomorphic(s string, t string) bool {
	m := make(map[byte]byte, len(s))

	for i := 0; i < len(s); i++ {
		_, ok := m[s[i]]
		if !ok {
			m[s[i]] = t[i]
		}

		if m[s[i]] != t[i] {
			return false
		}
	}

	uniqueValues := make(map[byte]struct{})

	for _, v := range m {
		uniqueValues[v] = struct{}{}
	}

	return len(m) == len(uniqueValues)
}
