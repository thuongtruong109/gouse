
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
	newCache := gouse.NewLCache()
	newCache.SetLCache("key1", "local cache value 1")
	newCache.SetLCache("key2", "local cache value 2")
	newCache.SetLCache("key3", "local cache value 3")

	all := newCache.AllLCache()
	println("All local cache values:")
	for k, v := range all {
		gouse.Printf("\t%s: %s\n", k, v)
	}

	getted1, err := newCache.GetLCache("key1")
	if err != nil {
		panic(err)
	}
	gouse.Println("Getted key 1:", getted1)

	newCache.DelLCache("key2")

	all = newCache.AllLCache()
	println("All local cache values (after delete key 2):")
	for k, v := range all {
		gouse.Printf("\t%s: %s\n", k, v)
	}

	newCache.FlushLCache()

	all = newCache.AllLCache()
	gouse.Println("All local cache values (after flush):", all)
}
```

## 2. Cache tmp

Description: Cache temporary values (with expiration time)<br>

```go
func CacheTmp() {
	newCache := gouse.NewTCache(gouse.ToSecond(3))
	newCache.SetTCache("key1", "temp cache value 1", gouse.ToSecond(3))
	newCache.SetTCache("key2", "temp cache value 2", gouse.ToSecond(6))
	newCache.SetTCache("key3", "temp cache value 3", gouse.ToSecond(6))

	all := newCache.AllTCache()
	println("All temp cache values:")
	for k, v := range all {
		gouse.Printf("\t%s: %v\n", k, v)
	}

	getted := newCache.GetTCache("key1")
	gouse.Println("Getted key 1 (before expires):", getted)

	time.Sleep(gouse.ToSecond(4))

	getted = newCache.GetTCache("key1")
	gouse.Println("Getted key 1 (after expires):", getted)

	newCache.DelTCache("key2")

	all = newCache.AllTCache()
	println("All temp cache values (after delete key 2):")
	for k, v := range all {
		gouse.Printf("\t%s: %v\n", k, v)
	}

	newCache.FlushTCache()

	all = newCache.AllTCache()
	gouse.Println("All temp cache values (after flush):", all)
}
```
