package samples

import (
	"fmt"
	"time"

	"github.com/thuongtruong109/gouse"
)

/*
Description: Cache local values
*/
func CacheLocal() {
	newCache := gouse.NewLCache()
	newCache.SetLCache("key1", "local cache value 1")
	newCache.SetLCache("key2", "local cache value 2")
	newCache.SetLCache("key3", "local cache value 3")

	all := newCache.AllLCache()
	println("All local cache values:")
	for k, v := range all {
		fmt.Printf("\t%s: %s\n", k, v)
	}

	getted1, err := newCache.GetLCache("key1")
	if err != nil {
		panic(err)
	}
	fmt.Println("Getted key 1:", getted1)

	newCache.DelLCache("key2")

	all = newCache.AllLCache()
	println("All local cache values (after delete key 2):")
	for k, v := range all {
		fmt.Printf("\t%s: %s\n", k, v)
	}

	newCache.FlushLCache()

	all = newCache.AllLCache()
	fmt.Println("All local cache values (after flush):", all)
}

/*
Description: Cache temporary values (with expiration time)
*/
func CacheTmp() {
	newCache := gouse.NewTCache(gouse.ToSecond(3))
	newCache.SetTCache("key1", "temp cache value 1", gouse.ToSecond(3))
	newCache.SetTCache("key2", "temp cache value 2", gouse.ToSecond(6))
	newCache.SetTCache("key3", "temp cache value 3", gouse.ToSecond(6))

	all := newCache.AllTCache()
	println("All temp cache values:")
	for k, v := range all {
		fmt.Printf("\t%s: %v\n", k, v)
	}

	getted := newCache.GetTCache("key1")
	fmt.Println("Getted key 1 (before expires):", getted)

	time.Sleep(gouse.ToSecond(4))

	getted = newCache.GetTCache("key1")
	fmt.Println("Getted key 1 (after expires):", getted)

	newCache.DelTCache("key2")

	all = newCache.AllTCache()
	println("All temp cache values (after delete key 2):")
	for k, v := range all {
		fmt.Printf("\t%s: %v\n", k, v)
	}

	newCache.FlushTCache()

	all = newCache.AllTCache()
	fmt.Println("All temp cache values (after flush):", all)
}
