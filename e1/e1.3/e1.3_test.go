package e1_3

import (
	"strings"
	"testing"
)

// 使用 go test --bench=. e1_test.go -run=none 即可看到结果
func BenchmarkStringPlus(b *testing.B) {
	strs := make([]string, 0)
	for i := 0; i < 100000; i++ {
		strs = append(strs, "abc")
	}
	var output string
	b.ResetTimer()
	for _, str := range strs {
		output += str
	}
}

func BenchmarkStringJoin(b *testing.B) {
	strs := make([]string, 0)
	for i := 0; i < 100000; i++ {
		strs = append(strs, "abc")
	}
	b.ResetTimer()
	strings.Join(strs, "")
}
