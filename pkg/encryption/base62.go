package encryption

import "strings"

const chars = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

func Encode(n int64) string {
	if n == 0 {
		return "0"
	}

	var sb strings.Builder
	for n > 0 {
		rem := n % 62
		sb.WriteByte(chars[rem])
		n /= 62
	}

	runes := []rune(sb.String())
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}
