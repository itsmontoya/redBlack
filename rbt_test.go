package rbt

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/itsmontoya/rbt/testUtils"

	"github.com/missionMeteora/journaler"

	"github.com/OneOfOne/skiplist"
	cznic "github.com/cznic/b"
)

var (
	testSortedList  = testUtils.GetSorted(10000)
	testReverseList = testUtils.GetReverse(10000)
	testRandomList  = testUtils.GetRand(10000)

	testSortedListStr  = testUtils.GetStrSlice(testSortedList)
	testReverseListStr = testUtils.GetStrSlice(testReverseList)
	testRandomListStr  = testUtils.GetStrSlice(testRandomList)

	testVal      []byte
	testCznicVal interface{}
)

func TestForEach(t *testing.T) {
	w := New(1024)
	w.ForEach(func(k, v []byte) (end bool) {
		journaler.Debug("uh shit?")
		return
	})
}

func TestRootDelete(t *testing.T) {
	w := New(1024)
	k := []byte("hello")
	v := []byte("world")
	w.Put(k, v)

	if val := string(w.Get(k)); val != string(v) {
		t.Fatalf("invalid value, expected \"%s\" and received \"%s\"", string(v), val)
	}

	w.Delete(k)

	if val := string(w.Get(k)); len(val) != 0 {
		t.Fatalf("invalid value, expected \"%s\" and received \"%s\"", "", val)
	}
}

func TestDelete(t *testing.T) {
	var keys [][]byte
	// New problem threshold, 100
	for i := 1; i <= 1000; i++ {
		keys = append(keys, []byte(strconv.Itoa(i)))
	}
	w := New(1024)

	for _, key := range keys {
		w.Put(key, key)

		if val := string(w.Get(key)); val != string(key) {
			t.Fatalf("invalid value, expected \"%s\" and received \"%s\"", string(key), val)
		}
	}

	for _, key := range keys {
		if val := string(w.Get(key)); val != string(key) {
			t.Fatalf("invalid value, expected \"%s\" and received \"%s\"", string(key), val)
		}

		w.Delete(key)

		if val := string(w.Get(key)); len(val) != 0 {
			t.Fatalf("invalid value, expected \"%s\" and received \"%s\"", "", val)
		}
	}
}

func TestGrow(t *testing.T) {
	w := New(1024)
	k := []byte("hello")
	v := []byte("world")
	empty := []byte{0, 0, 0, 0, 0}
	w.Put(k, v)
	w.Grow(k, 10)

	if rv := w.Get(k); len(rv) != 10 {
		t.Fatalf("invalid value length, expected %d and received %d (%v)", 10, len(rv), rv)
	} else if string(rv[:5]) != string(v) {
		t.Fatalf("invalid value, expected %s and received %s", string(rv[:5]), string(v))
	} else if string(rv[5:]) != string(empty) {
		t.Fatalf("invalid value, expected an end of %v and received %v", empty, rv[5:])
	}
}

func TestBasic(t *testing.T) {
	var err error
	if err = os.MkdirAll("./test_data", 0755); err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll("./test_data")

	var tr *Tree
	if tr, err = NewMMAP("./test_data", "mmap.db", 64); err != nil {
		t.Fatal(err)
	}

	tr.Put([]byte("1"), []byte("1"))
	tr.Put([]byte("2"), []byte("2"))
	tr.Put([]byte("3"), []byte("3"))
	tr.Put([]byte("4"), []byte("4"))
	tr.Put([]byte("5"), []byte("5"))
	tr.Put([]byte("6"), []byte("6"))
	tr.Put([]byte("7"), []byte("7"))
	tr.Put([]byte("8"), []byte("8"))
	tr.Put([]byte("9"), []byte("9"))
	tr.Put([]byte("10"), []byte("10"))

	if val := string(tr.Get([]byte("1"))); val != "1" {
		t.Fatalf("invalid value, expected \"%v\" and received \"%v\"", "1", val)
	}

	if val := string(tr.Get([]byte("2"))); val != "2" {
		t.Fatalf("invalid value, expected \"%v\" and received \"%v\"", "1", val)
	}

	if val := string(tr.Get([]byte("3"))); val != "3" {
		t.Fatalf("invalid value, expected \"%v\" and received \"%v\"", "1", val)
	}

	if val := string(tr.Get([]byte("4"))); val != "4" {
		t.Fatalf("invalid value, expected \"%v\" and received \"%v\"", "1", val)
	}

	if val := string(tr.Get([]byte("5"))); val != "5" {
		t.Fatalf("invalid value, expected \"%v\" and received \"%v\"", "1", val)
	}

	if val := string(tr.Get([]byte("6"))); val != "6" {
		t.Fatalf("invalid value, expected \"%v\" and received \"%v\"", "1", val)
	}

	if val := string(tr.Get([]byte("7"))); val != "7" {
		t.Fatalf("invalid value, expected \"%v\" and received \"%v\"", "1", val)
	}

	if val := string(tr.Get([]byte("8"))); val != "8" {
		t.Fatalf("invalid value, expected \"%v\" and received \"%v\"", "1", val)
	}

	if val := string(tr.Get([]byte("9"))); val != "9" {
		t.Fatalf("invalid value, expected \"%v\" and received \"%v\"", "1", val)
	}

	if val := string(tr.Get([]byte("10"))); val != "10" {
		t.Fatalf("invalid value, expected \"%v\" and received \"%v\"", "1", val)
	}

	tr.Close()

	if tr, err = NewMMAP("./test_data", "mmap.db", 64); err != nil {
		t.Fatal(err)
	}

	if val := string(tr.Get([]byte("1"))); val != "1" {
		t.Fatalf("invalid value, expected \"%v\" and received \"%v\"", "1", val)
	}

	if val := string(tr.Get([]byte("2"))); val != "2" {
		t.Fatalf("invalid value, expected \"%v\" and received \"%v\"", "1", val)
	}

	if val := string(tr.Get([]byte("3"))); val != "3" {
		t.Fatalf("invalid value, expected \"%v\" and received \"%v\"", "1", val)
	}

	if val := string(tr.Get([]byte("4"))); val != "4" {
		t.Fatalf("invalid value, expected \"%v\" and received \"%v\"", "1", val)
	}

	if val := string(tr.Get([]byte("5"))); val != "5" {
		t.Fatalf("invalid value, expected \"%v\" and received \"%v\"", "1", val)
	}

	if val := string(tr.Get([]byte("6"))); val != "6" {
		t.Fatalf("invalid value, expected \"%v\" and received \"%v\"", "1", val)
	}

	if val := string(tr.Get([]byte("7"))); val != "7" {
		t.Fatalf("invalid value, expected \"%v\" and received \"%v\"", "1", val)
	}

	if val := string(tr.Get([]byte("8"))); val != "8" {
		t.Fatalf("invalid value, expected \"%v\" and received \"%v\"", "1", val)
	}

	if val := string(tr.Get([]byte("9"))); val != "9" {
		t.Fatalf("invalid value, expected \"%v\" and received \"%v\"", "1", val)
	}

	if val := string(tr.Get([]byte("10"))); val != "10" {
		t.Fatalf("invalid value, expected \"%v\" and received \"%v\"", "1", val)
	}

}

func TestSortedPut(t *testing.T) {
	testPut(t, testUtils.GetSorted(10))
}

func TestReversePut(t *testing.T) {
	testPut(t, testUtils.GetReverse(10))
}

func TestRandomPut(t *testing.T) {
	testPut(t, testUtils.GetRand(10))
}

func BenchmarkTreeGet(b *testing.B) {
	benchGet(b, testSortedListStr)
	b.ReportAllocs()
}

func BenchmarkTreeSortedGetPut(b *testing.B) {
	benchGetPut(b, testSortedListStr)
	b.ReportAllocs()
}

func BenchmarkTreeSortedPut(b *testing.B) {
	benchPut(b, testSortedListStr)
	b.ReportAllocs()
}

func BenchmarkTreeReversePut(b *testing.B) {
	benchPut(b, testReverseListStr)
	b.ReportAllocs()
}

func BenchmarkTreeRandomPut(b *testing.B) {
	benchPut(b, testRandomListStr)
	b.ReportAllocs()
}

func BenchmarkTreeForEach(b *testing.B) {
	benchForEach(b, testSortedListStr)
	b.ReportAllocs()
}

func BenchmarkTreeMMapGet(b *testing.B) {
	benchMMAPGet(b, testSortedListStr)
	b.ReportAllocs()
}

func BenchmarkTreeMMapSortedGetPut(b *testing.B) {
	benchMMAPGetPut(b, testSortedListStr)
	b.ReportAllocs()
}

func BenchmarkTreeMMapSortedPut(b *testing.B) {
	benchMMAPPut(b, testSortedListStr)
	b.ReportAllocs()
}

func BenchmarkTreeMMapReversePut(b *testing.B) {
	benchMMAPPut(b, testReverseListStr)
	b.ReportAllocs()
}

func BenchmarkTreeMMapRandomPut(b *testing.B) {
	benchMMAPPut(b, testRandomListStr)
	b.ReportAllocs()
}

func BenchmarkTreeMMapForEach(b *testing.B) {
	benchMMAPForEach(b, testSortedListStr)
	b.ReportAllocs()
}

func BenchmarkMapGet(b *testing.B) {
	benchMapGet(b, testSortedListStr)
	b.ReportAllocs()
}

func BenchmarkMapSortedGetPut(b *testing.B) {
	benchMapGetPut(b, testSortedListStr)
	b.ReportAllocs()
}

func BenchmarkMapSortedPut(b *testing.B) {
	benchMapPut(b, testSortedListStr)
	b.ReportAllocs()
}

func BenchmarkMapReversePut(b *testing.B) {
	benchMapPut(b, testReverseListStr)
	b.ReportAllocs()
}

func BenchmarkMapRandomPut(b *testing.B) {
	benchMapPut(b, testRandomListStr)
	b.ReportAllocs()
}

func BenchmarkMapForEach(b *testing.B) {
	benchMapForEach(b, testSortedListStr)
	b.ReportAllocs()
}
func BenchmarkSkiplistGet(b *testing.B) {
	benchSkiplistGet(b, testSortedListStr)
	b.ReportAllocs()
}

func BenchmarkSkiplistSortedGetPut(b *testing.B) {
	benchSkiplistGetPut(b, testSortedListStr)
	b.ReportAllocs()
}

func BenchmarkSkiplistSortedPut(b *testing.B) {
	benchSkiplistPut(b, testSortedListStr)
	b.ReportAllocs()
}

func BenchmarkSkiplistReversePut(b *testing.B) {
	benchSkiplistPut(b, testReverseListStr)
	b.ReportAllocs()
}

func BenchmarkSkiplistRandomPut(b *testing.B) {
	benchSkiplistPut(b, testRandomListStr)
	b.ReportAllocs()
}

func BenchmarkSkiplistForEach(b *testing.B) {
	benchSkiplistForEach(b, testSortedListStr)
	b.ReportAllocs()
}

func BenchmarkCznicGet(b *testing.B) {
	benchCznicGet(b, testSortedListStr)
	b.ReportAllocs()
}

func BenchmarkCznicSortedGetPut(b *testing.B) {
	benchCznicGetPut(b, testSortedListStr)
	b.ReportAllocs()
}

func BenchmarkCznicSortedPut(b *testing.B) {
	benchCznicPut(b, testSortedListStr)
	b.ReportAllocs()
}

func BenchmarkCznicReversePut(b *testing.B) {
	benchCznicPut(b, testReverseListStr)
	b.ReportAllocs()
}

func BenchmarkCznicRandomPut(b *testing.B) {
	benchCznicPut(b, testRandomListStr)
	b.ReportAllocs()
}

func BenchmarkCznicForEach(b *testing.B) {
	benchCznicForEach(b, testSortedListStr)
	b.ReportAllocs()
}

func testPut(t *testing.T, s []int) {
	cnt := len(s)
	tr := New(1024 * 1024)
	tm := make(map[string][]byte, cnt)

	for _, v := range s {
		key := fmt.Sprintf("%012d", v)
		val := []byte(strconv.Itoa(v))
		tr.Put([]byte(key), val)
		tm[key] = val
	}

	for key, mv := range tm {
		val := tr.Get([]byte(key))
		if !bytes.Equal(val, mv) {
			t.Fatalf("invalid value:\nKey: %s\nExpected: %v\nReturned: %v\n", key, mv, val)
		}
	}

	var fecnt int
	tr.ForEach(func(key, val []byte) (end bool) {
		if !bytes.Equal(val, tm[string(key)]) {
			t.Fatalf("invalid value:\nKey: %s\nExpected: %v\nReturned: %v\n", key, tm[string(key)], val)
		}

		fecnt++
		return
	})

	if fecnt != cnt {
		t.Fatalf("invalid ForEach iterations:\nExpected: %v\nActual: %v\n", cnt, fecnt)
	}
}

func benchGet(b *testing.B, s []testUtils.KV) {
	tr := New(1024 * 1024)
	for _, kv := range s {
		tr.Put(kv.Val, kv.Val)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, kv := range s {
			testVal = tr.Get(kv.Val)
		}
	}
}

func benchPut(b *testing.B, s []testUtils.KV) {
	tr := New(1024 * 1024)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, kv := range s {
			tr.Put(kv.Val, kv.Val)
		}
	}
}

func benchGetPut(b *testing.B, s []testUtils.KV) {
	tr := New(1024 * 1024)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, kv := range s {
			tr.Put(kv.Val, kv.Val)
			testVal = tr.Get(kv.Val)
		}
	}
}

func benchForEach(b *testing.B, s []testUtils.KV) {
	tr := New(1024 * 1024)

	for _, kv := range s {
		tr.Put(kv.Val, kv.Val)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		tr.ForEach(func(_, val []byte) (end bool) {
			testVal = val
			return
		})
	}
}

func benchMMAPGet(b *testing.B, s []testUtils.KV) {
	var err error
	if err = os.MkdirAll("./test_data", 0755); err != nil {
		b.Fatal(err)
	}
	defer os.RemoveAll("./test_data")

	var tr *Tree
	if tr, err = NewMMAP("./test_data", "mmap.db", 1024*1024); err != nil {
		b.Fatal(err)
	}

	for _, kv := range s {
		tr.Put(kv.Val, kv.Val)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, kv := range s {
			testVal = tr.Get(kv.Val)
		}
	}
}

func benchMMAPPut(b *testing.B, s []testUtils.KV) {
	var err error
	if err = os.MkdirAll("./test_data", 0755); err != nil {
		b.Fatal(err)
	}
	defer os.RemoveAll("./test_data")

	var tr *Tree
	if tr, err = NewMMAP("./test_data", "mmap.db", 1024*1024); err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, kv := range s {
			tr.Put(kv.Val, kv.Val)
		}
	}
}

func benchMMAPGetPut(b *testing.B, s []testUtils.KV) {
	var err error
	if err = os.MkdirAll("./test_data", 0755); err != nil {
		b.Fatal(err)
	}
	defer os.RemoveAll("./test_data")

	var tr *Tree
	if tr, err = NewMMAP("./test_data", "mmap.db", 1024*1024); err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, kv := range s {
			tr.Put(kv.Val, kv.Val)
			testVal = tr.Get(kv.Val)
		}
	}
}

func benchMMAPForEach(b *testing.B, s []testUtils.KV) {
	var err error
	if err = os.MkdirAll("./test_data", 0755); err != nil {
		b.Fatal(err)
	}
	defer os.RemoveAll("./test_data")

	var tr *Tree
	if tr, err = NewMMAP("./test_data", "mmap.db", 1024*1024); err != nil {
		b.Fatal(err)
	}

	for _, kv := range s {
		tr.Put(kv.Val, kv.Val)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		tr.ForEach(func(_, val []byte) (end bool) {
			testVal = val
			return
		})
	}
}

func benchMapGet(b *testing.B, s []testUtils.KV) {
	m := make(map[string][]byte)
	for _, kv := range s {
		m[kv.Key] = kv.Val
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, kv := range s {
			testVal = m[kv.Key]
		}
	}
}

func benchMapPut(b *testing.B, s []testUtils.KV) {
	b.ResetTimer()
	m := make(map[string][]byte)

	for i := 0; i < b.N; i++ {
		for _, kv := range s {
			m[kv.Key] = kv.Val
		}
	}
}

func benchMapGetPut(b *testing.B, s []testUtils.KV) {
	b.ResetTimer()
	m := make(map[string][]byte)

	for i := 0; i < b.N; i++ {
		for _, kv := range s {
			testVal = m[kv.Key]
			m[kv.Key] = kv.Val
		}
	}
}

func benchMapForEach(b *testing.B, s []testUtils.KV) {
	m := make(map[string][]byte)
	for _, kv := range s {
		m[kv.Key] = kv.Val
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, val := range m {
			testVal = val
		}
	}
}

func benchSkiplistGet(b *testing.B, s []testUtils.KV) {
	sl := skiplist.New(32)
	for _, kv := range s {
		sl.Set(kv.Key, kv.Val)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, kv := range s {
			testVal = sl.Get(kv.Key).([]byte)
		}
	}
}

func benchSkiplistPut(b *testing.B, s []testUtils.KV) {
	b.ResetTimer()
	sl := skiplist.New(32)
	for i := 0; i < b.N; i++ {
		for _, kv := range s {
			sl.Set(kv.Key, kv.Val)
		}
	}
}

func benchSkiplistGetPut(b *testing.B, s []testUtils.KV) {
	b.ResetTimer()
	sl := skiplist.New(32)

	for i := 0; i < b.N; i++ {
		for _, kv := range s {
			sl.Set(kv.Key, kv.Val)
			testVal = sl.Get(kv.Key).([]byte)
		}
	}
}

func benchSkiplistForEach(b *testing.B, s []testUtils.KV) {
	sl := skiplist.New(32)
	for _, kv := range s {
		sl.Set(kv.Key, kv.Val)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		sl.ForEach(func(_ string, val interface{}) bool {
			testVal = val.([]byte)
			return false
		})
	}
}

func benchCznicGet(b *testing.B, s []testUtils.KV) {
	tr := cznic.TreeNew(byteLess)
	for _, kv := range s {
		tr.Put(kv.Val, func(_ interface{}, exists bool) (interface{}, bool) {
			return kv.Val, true
		})
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, kv := range s {
			testCznicVal, _ = tr.Get(kv.Val)
		}
	}
}

func byteLess(a, b interface{}) int {
	return bytes.Compare(a.([]byte), b.([]byte))
}

func benchCznicPut(b *testing.B, s []testUtils.KV) {
	tr := cznic.TreeNew(byteLess)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, kv := range s {
			tr.Put(kv.Val, func(_ interface{}, exists bool) (interface{}, bool) {
				return kv.Val, true
			})
		}
	}
}

func benchCznicGetPut(b *testing.B, s []testUtils.KV) {
	tr := cznic.TreeNew(byteLess)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, kv := range s {
			tr.Put(kv.Val, func(_ interface{}, exists bool) (interface{}, bool) {
				return kv.Val, true
			})
			testCznicVal, _ = tr.Get(kv.Val)
		}
	}
}

func benchCznicForEach(b *testing.B, s []testUtils.KV) {
	tr := cznic.TreeNew(byteLess)

	for _, kv := range s {
		tr.Put(kv.Val, func(_ interface{}, exists bool) (interface{}, bool) {
			return kv.Val, true
		})
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		e, ok := tr.Seek([]byte(""))
		if !ok {
			b.Fatal("error calling iterator")
		}

		var (
			v   interface{}
			err error
			cnt int
		)

		for _, v, err = e.Next(); err == nil; _, v, err = e.Next() {
			testCznicVal = v
			cnt++
		}

		if cnt != len(s) {
			b.Fatalf("invalid count, expected %v and received %v (%v)", len(s), cnt, err)
		}
	}
}
