package cache

import (
	"testing"
	"time"
)

func TestSimpleCache(t *testing.T)  {
	cache := NewSimpleCache("test")
	t.Log("exists:accessToken_xxx",  cache.IsExist("accessToken_xxx"))
	cache.Set("accessToken_acc", "aaa", 1000*time.Second)
	cache.Set("accessToken_js", "bbb", 7200*time.Second)
	cache.Set("accessToken_xxx", "ccc", 2888*time.Second)

}
