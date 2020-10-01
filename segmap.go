package segmap

import (
	"github.com/bytesconv/hashkit"
	"strconv"
	"sync"
)

const SegmCount = 13

type Segmap struct {
	buckets []sync.Map
	segmCount uint32
}

func New(segmCount uint32) *Segmap {
	if segmCount <= 0 {
		segmCount = SegmCount
	}

	buckets := make([]sync.Map, segmCount)
	for i := uint32(0); i < segmCount; i++ {
		buckets = append(buckets, sync.Map{})
	}

	return &Segmap{
		buckets: buckets,
		segmCount: segmCount,
	}
}

func (segmap *Segmap) index(key string) uint32 {
	return hashkit.Fnv32(key) % segmap.segmCount
}

func (segmap *Segmap) Store(key string, value interface{}) {
	segmap.buckets[segmap.index(key)].Store(key, value)
}

func (segmap *Segmap) LoadOrStore(key string, value interface{}) (actual interface{}, loaded bool) {
	actual, loaded = segmap.buckets[segmap.index(key)].LoadOrStore(key, value)
	return
}

func (segmap *Segmap) Load(key string) (value interface{}, ok bool)  {
	value, ok = segmap.buckets[segmap.index(key)].Load(key)
	return
}

func (segmap *Segmap) LoadAndDelete(key string) (value interface{}, loaded bool)  {
	value, loaded = segmap.buckets[segmap.index(key)].LoadAndDelete(key)
	return
}

func (segmap *Segmap) Delete(key string) {
	segmap.buckets[segmap.index(key)].Delete(key)
}

func (segmap *Segmap) Range(f func(key interface{}, value interface{}) bool) {
	for i := uint32(0); i < segmap.segmCount; i++ {
		segmap.buckets[segmap.index(strconv.FormatUint(uint64(i), 10))].Range(f)
	}
}