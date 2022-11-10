package rand

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// StringN generate rand resource name, length: len(prefix) + n.
func StringN(prefix string, n int) string {
	return fmt.Sprintf("%s-%s", prefix, String(n))
}

// String generate a random string of n length.
func String(n int) string {
	if n <= 0 {
		return ""
	}

	var chars = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l",
		"m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}

	str := strings.Builder{}
	length := len(chars)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < n; i++ {
		str.WriteString(chars[rand.Intn(length)])
	}

	return str.String()
}

// Number generate a random number.
func Number() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(1000000000)
}
