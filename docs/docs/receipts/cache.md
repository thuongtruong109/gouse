
# <Badge style='font-size: 1.8rem; text-shadow: 1px 1px 2px rgba(0, 0, 0, 0.3); padding: 0.35rem 0.75rem 0.35rem 0;' type='info' text='🔖 Cache' />


```go
import (
	"time"
	"github.com/thuongtruong109/gouse"
)
```

## 1. Cache local

Description: Cache local values<br>

```go
func CacheLocal() {
	newCache := gouse.NewLocalCache()
	newCache.SetLocalCache("key1", "local cache value 1")
	newCache.SetLocalCache("key2", "local cache value 2")
	newCache.SetLocalCache("key3", "local cache value 3")

	all := newCache.AllLocalCache()
	println("All local cache values:")
	for k, v := range all {
		gouse.Printf("\t%s: %s\n", k, v)
	}

	getted1, err := newCache.GetLocalCache("key1")
	if err != nil {
		panic(err)
	}
	gouse.Println("Getted key 1:", getted1)

	newCache.DelLocalCache("key2")

	all = newCache.AllLocalCache()
	println("All local cache values (after delete key 2):")
	for k, v := range all {
		gouse.Printf("\t%s: %s\n", k, v)
	}

	newCache.FlushLocalCache()

	all = newCache.AllLocalCache()
	gouse.Println("All local cache values (after flush):", all)
}
```

## 2. Cache tmp

Description: Cache temporary values (with expiration time)<br>

```go
func CacheTmp() {
	newCache := gouse.NewTmpCache(gouse.ToSecond(3))
	newCache.SetTmpCache("key1", "temp cache value 1", gouse.ToSecond(3))
	newCache.SetTmpCache("key2", "temp cache value 2", gouse.ToSecond(6))
	newCache.SetTmpCache("key3", "temp cache value 3", gouse.ToSecond(6))

	all := newCache.AllTmpCache()
	println("All temp cache values:")
	for k, v := range all {
		gouse.Printf("\t%s: %v\n", k, v)
	}

	getted := newCache.GetTmpCache("key1")
	gouse.Println("Getted key 1 (before expires):", getted)

	time.Sleep(gouse.ToSecond(4))

	getted = newCache.GetTmpCache("key1")
	gouse.Println("Getted key 1 (after expires):", getted)

	newCache.DelTmpCache("key2")

	all = newCache.AllTmpCache()
	println("All temp cache values (after delete key 2):")
	for k, v := range all {
		gouse.Printf("\t%s: %v\n", k, v)
	}

	newCache.FlushTmpCache()

	all = newCache.AllTmpCache()
	gouse.Println("All temp cache values (after flush):", all)
}
```
