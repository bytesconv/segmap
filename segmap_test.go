package segmap

import (
	"strconv"
	"sync"
	"testing"
)

func BenchmarkSegmap_Store(b *testing.B) {
	s := New(13)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		go s.Store(strconv.Itoa(i), "1")
	}
}

func BenchmarkSyncMap_Store(b *testing.B) {
	s := sync.Map{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		go s.Store(strconv.Itoa(i), "1")
	}
}

func BenchmarkSegmap_Load(b *testing.B) {
	s := New(13)
	for i := 0; i < 512; i++ {
		s.Store(strconv.Itoa(i), strconv.Itoa(i))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		go s.Load(strconv.Itoa(i))
	}
}

func BenchmarkSyncMap_Load(b *testing.B) {
	s := sync.Map{}
	for i := 0; i < 512; i++ {
		s.Store(strconv.Itoa(i), strconv.Itoa(i))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		go s.Load(strconv.Itoa(i))
	}
}

func BenchmarkSegmap_LoadAndDelete(b *testing.B) {
	s := New(13)
	for i := 0; i < 512; i++ {
		s.Store(strconv.Itoa(i), strconv.Itoa(i))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		go s.LoadAndDelete(strconv.Itoa(i))
	}
}

func BenchmarkSyncMap_LoadAndDelete(b *testing.B) {
	s := sync.Map{}
	for i := 0; i < 512; i++ {
		s.Store(strconv.Itoa(i), strconv.Itoa(i))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		go s.Load(strconv.Itoa(i))
	}
}

func BenchmarkSegmap_Range(b *testing.B) {
	s := New(13)
	for i := 0; i < 512; i++ {
		s.Store(strconv.Itoa(i), strconv.Itoa(i))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Range(func(key interface{}, value interface{}) bool {
			return true
		})
	}
}

func BenchmarkSyncMap_Range(b *testing.B) {
	s := new(sync.Map)
	for i := 0; i < 512; i++ {
		s.Store(strconv.Itoa(i), strconv.Itoa(i))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Range(func(key interface{}, value interface{}) bool {
			return true
		})
	}
}