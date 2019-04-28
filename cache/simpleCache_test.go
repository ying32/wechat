package cache

import (
	"testing"
)

func TestSimpleCache(t *testing.T)  {
	cache := NewSimpleCache("test")
	t.Log("exists:accessToken_xxx",  cache.IsExist("accessToken_xxx"))
	cache.Set("accessToken_acc", "aaa", 7200)
	cache.Set("accessToken_js", "bbb", 7300)
	cache.Set("accessToken_xxx", "ccc", 99999200)

}
