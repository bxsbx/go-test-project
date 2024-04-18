package testing

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestFmt(t *testing.T) {
	for i := 0; i < 10; i++ {
		fmt.Println("test")
	}
}

func TestFmtParallel(t *testing.T) {
	for i := 0; i < 10; i++ {
		name := strconv.Itoa(i)
		t.Run(name, func(t *testing.T) {
			t.Parallel() //将每个测试用例标记为能够彼此并行运行
			fmt.Println(name, "test")
		})
	}
}

func BenchmarkSplit(b *testing.B) {
	time.Sleep(5 * time.Second) // 假设需要做一些耗时的无关操作
	b.ResetTimer()              // 重置计时器

	for i := 0; i < b.N; i++ {
		split := strings.Split("a:b:c:d:e", ":")
		fmt.Println(split)
	}
}

func BenchmarkSplitParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			split := strings.Split("a:b:c:d:e", ":")
			fmt.Println(split)
		}
	})
}

func FuzzSplit(f *testing.F) {
	f.Add("Hello, world", "!12345") // Use f.Add to provide a seed corpus

	f.Fuzz(func(t *testing.T, a, b string) {
		split := strings.Split(a, b)
		fmt.Println(split)
	})
}
