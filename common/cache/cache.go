package cache

import (
	"github.com/zeromicro/go-zero/core/collection"
	"time"
)

var cache *collection.Cache
var usecache bool //是否使用缓存
// 初始化 cache，其中 WithLimit 可以指定最大缓存的数量
func Init() {
	var err error
	usecache = true
	cache, err = collection.NewCache(5*time.Second, collection.WithLimit(10000))
	if err != nil {
		panic(err)
	}
}

func Set(k string, x interface{}) {
	cache.Set(k, x)
}
func Find(k string, val interface{}) bool {
	val, find := cache.Get(k)
	return find
}
func Get(k string) (interface{}, bool) {
	return cache.Get(k)
}

// 删除 cache
func Del(k string) {
	cache.Del(k)
}

// 获取缓存，如果 key 不存在的，则会调用 func 去生成缓存
func Take(k string, f func() (interface{}, error)) (interface{}, error) {
	if !usecache {
		return f()
	}
	return cache.Take(k, f)
}
