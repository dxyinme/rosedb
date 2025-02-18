package cache

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func TestNewCache(t *testing.T) {
	cache := NewLruCache(64)
	assert.NotNil(t, cache)
}

func TestLruCache_Get(t *testing.T) {
	lru := NewLruCache(64)
	lru.Set([]byte("k-1"), []byte("ss"))
	lru.Set([]byte("k-1"), []byte("dd"))
	lru.Set([]byte("k-1"), []byte("ssss"))

	lru.Set([]byte("k-2"), []byte("bbb"))

	v, ok := lru.Get([]byte("k-1"))
	t.Log(string(v), ok)

	v1, ok := lru.Get([]byte("k-2"))
	t.Log(string(v1), ok)
}

func TestLruCache_Set(t *testing.T) {
	lru := NewLruCache(100)
	lru.Set(nil, nil)
}

func BenchmarkCache_Set(b *testing.B) {
	cache := NewLruCache(500)

	b.ResetTimer()
	b.ReportAllocs()

	//k, v := []byte("test-key"), []byte("test-value")
	for i := 0; i < b.N; i++ {
		cache.Set(GetKey(i), GetValue())
	}
}

func BenchmarkCache_Get(b *testing.B) {
	cache := NewLruCache(1024)
	for i := 0; i < 10000; i++ {
		cache.Set(GetKey(i), GetValue())
	}

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		cache.Get(GetKey(i))
	}
}

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func GetKey(n int) []byte {
	return []byte("test_key_" + fmt.Sprintf("%09d", n))
}

func GetValue() []byte {
	var str bytes.Buffer
	for i := 0; i < 12; i++ {
		str.WriteByte(alphabet[rand.Int()%26])
	}
	return []byte("test_val-" + strconv.FormatInt(time.Now().UnixNano(), 10) + str.String())
}
