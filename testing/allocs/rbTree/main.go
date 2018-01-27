package main

import (
	"fmt"
	"runtime"
	"strconv"

	"github.com/itsmontoya/rbTree"
)

var val interface{}

func main() {
	var start, end runtime.MemStats
	runtime.ReadMemStats(&start)
	t := populateN(1000000)
	runtime.GC()
	val = t.Get("1")
	runtime.ReadMemStats(&end)
	fmt.Println(end.Alloc - start.Alloc)
}

func populateN(n int) (t *rbTree.Tree) {
	t = rbTree.New(n)

	for i := 0; i < n; i++ {
		key := strconv.Itoa(i)
		t.Put(key, []byte(key))
	}

	return
}
