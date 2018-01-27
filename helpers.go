package rbTree

type color uint8

type childType uint8

// ForEachFn is used when calling ForEach from a Tree
type ForEachFn func(key, val []byte) (end bool)

// GrowFn is used when calling grow internally
type GrowFn func(sz int64) (bs []byte)

func getSize(cnt, keySize, valSize int64) (sz int64) {
	sz = keySize + valSize + blockSize
	sz *= cnt
	sz += trunkSize
	return
}
