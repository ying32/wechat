package cache

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

type TSimpleCacheItem struct {
	Key string
	AccessKey     string
	ExpirsTime    int64
}

type TSimpleCache struct {
	cacehFileName string
	items map[string]TSimpleCacheItem
	Cache
}

// NewSimpleCache 本地缓存
func NewSimpleCache(name string) *TSimpleCache {
	f := new(TSimpleCache)
	file, _ := exec.LookPath(os.Args[0])
	// 文件格式为 simpleCache_appid
	f.cacehFileName = filepath.Dir(file) + "/.simpleCache_" + name
	f.items = make(map[string]TSimpleCacheItem, 0)
	bs, err := ioutil.ReadFile(f.cacehFileName)
	if err == nil && len(bs) > 0 {
		vals := []TSimpleCacheItem{}
		err := json.Unmarshal(bs, &vals)
		if err == nil {
			for _, v := range vals {
				f.items[v.Key] = v
			}
		} else {
			log.Println("读取simpleCahce错误。")
		}
		log.Println(f.items)
	}
	return f
}

// Get 获取
func (f *TSimpleCache) Get(key string) interface{} {
	if f.IsExist(key) {
		return f.items[key].AccessKey
	}
	return nil
}

// Set更新accesskey及过期时间
func (f *TSimpleCache) Set(key string, val interface{}, timeout time.Duration) error {
	item := TSimpleCacheItem{
		 Key:key,
		 AccessKey:val.(string),
		 ExpirsTime:time.Now().Add(timeout * time.Second - time.Duration(5)*time.Second).Unix(), // 提前5秒获取新的accesskey
	}
	f.items[key] = item
	// 写到文件
	log.Println("更新",key, "，过期时间为：", time.Unix(item.ExpirsTime, 0))

	vals := []TSimpleCacheItem{}
	for _, v := range f.items {
		vals = append(vals, v)
	}
	bs, err := json.MarshalIndent(&vals, "", "\t")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(f.cacehFileName, bs, 0770)
}

// IsExist 如果accesskey为空，或者已经过期，表示不存在这个了
func (f *TSimpleCache) IsExist(key string) bool {
	 if val, ok := f.items[key]; ok {
		 if val.ExpirsTime > time.Now().Unix() {
			 return true
		 }
	 }
	return false
}

// Delete 清空及删除缓存文件
func (f *TSimpleCache) Delete(key string) error {
	for key, _ := range f.items {
		delete(f.items, key)
	}
	return os.Remove(f.cacehFileName)
}
