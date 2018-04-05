package rbt

type color uint8

type childType uint8

type trunk struct {
	root int64
	cnt  int64
	tail int64
	cap  int64
}

// ForEachFn is used when calling ForEach from a Tree
type ForEachFn func(key, val []byte) (end bool)

// IterateFn is used when calling iterate from a Tree
type IterateFn func(b *Block) (end bool)

// GrowFn is used when calling grow internally
type GrowFn func(sz int64) (bs []byte)

// CloseFn is used when close is called
type CloseFn func() error

func getSize(cnt, keySize, valSize int64) (sz int64) {
	sz = keySize + valSize + BlockSize
	sz *= cnt
	sz += TrunkSize
	return
}
