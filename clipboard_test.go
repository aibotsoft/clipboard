package clipboard

import (
	"testing"
)

func BenchmarkReadAll(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		got, err := Get()
		if err != nil {
			b.Log(err)
			b.FailNow()
		}
		if got == "" {
			b.Log("clip is empty")
			b.FailNow()
		}
	}
}
