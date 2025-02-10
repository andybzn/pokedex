package pokecache

import (
	"fmt"
	"testing"
	"time"
)

func TestCacheAddGet(t *testing.T) {
	const interval = 5 * time.Second
	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "https://andy.bz",
			val: []byte("example data"),
		},
		{
			key: "https://andy.uno",
			val: []byte("more example data"),
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("test case: %v", i), func(t *testing.T) {
			cache := NewCache(interval)
			cache.Add(c.key, c.val)
			value, exists := cache.Get(c.key)
			if !exists {
				t.Errorf("expected to find key")
				return
			}
			if string(value) != string(c.val) {
				t.Errorf("expected to find value")
				return
			}
		})
	}
}

func TestCacheReap(t *testing.T) {
	const baseTime = 5 * time.Millisecond
	const waitTime = baseTime + (5 * time.Millisecond)
	cache := NewCache(baseTime)
	cache.Add("https://andy.bz", []byte("example data"))

	_, ok := cache.Get("https://andy.bz")
	if !ok {
		t.Errorf("expected to find key")
		return
	}

	time.Sleep(waitTime)

	_, ok = cache.Get("https://andy.bz")
	if ok {
		t.Errorf("expected key to have been removed")
		return
	}
}
