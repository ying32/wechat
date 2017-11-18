package cache

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"

	"time"

 
)

type TSimpleCache struct {
	cacehFileName string
	accessKey     string
	expirsTime    time.Time
	Cache
}

// NewSimpleCache 本地缓存
func NewSimpleCache(name string) *TSimpleCache {
	f := new(TSimpleCache)
	file, _ := exec.LookPath(os.Args[0])
	f.cacehFileName = filepath.Dir(file) + "/.access_token_" + name
	bs, err := ioutil.ReadFile(f.cacehFileName)
	if err == nil && len(bs) > 0 {
		sps := strings.Split(string(bs), "|")
		if len(sps) == 2 {
			f.accessKey = sps[0]
			i64, err := strconv.ParseInt(sps[1], 10, 0)
			if err == nil {
				f.expirsTime = time.Unix(i64, 0)
			} else {
				f.accessKey = ""
			}
		}
	}
	return f
}

// Get 获取
func (f *TSimpleCache) Get(key string) interface{} {
	if f.IsExist(key) {
		return f.accessKey
	}
	return nil
}

// Set更新accesskey及过期时间
func (f *TSimpleCache) Set(key string, val interface{}, timeout time.Duration) error {
	f.accessKey = val.(string)
	f.expirsTime = time.Now().Add(timeout - time.Duration(5)*time.Second) // 提前5秒获取新的accesskey
	// 写到文件
	log.Println("更新accessKey，过期时间为：", f.expirsTime)
	return ioutil.WriteFile(f.cacehFileName, []byte(fmt.Sprintf("%s|%d", val.(string), f.expirsTime.Unix())), 0770)
}

// IsExist 如果accesskey为空，或者已经过期，表示不存在这个了
func (f *TSimpleCache) IsExist(key string) bool {
	if f.expirsTime.Unix() < time.Now().Unix() {
		return false
	}
	if f.accessKey == "" {
		return false
	}
	return true
}

// Delete 清空及删除缓存文件
func (f *TSimpleCache) Delete(key string) error {
	f.accessKey = ""
	f.expirsTime = time.Now()
	return os.Remove(f.cacehFileName)
}
