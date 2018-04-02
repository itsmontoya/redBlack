package backend

import (
	"github.com/Path94/atoms"
	"github.com/itsmontoya/rbt/allocator"
	"github.com/missionMeteora/journaler"
	"github.com/missionMeteora/toolkit/errors"
)

// New will return a new backend
func New(m *Multi) *Backend {
	var b Backend
	b.m = m
	m.a.OnGrow(b.onGrow)
	return &b
}

// Backend represents a data Backend
type Backend struct {
	m  *Multi
	s  allocator.Section
	bs []byte

	closed atoms.Bool
}

func (b *Backend) onGrow() (end bool) {
	if b.closed.Get() {
		return true
	}

	b.SetBytes()
	return
}

// SetBytes will refresh the bytes reference
func (b *Backend) SetBytes() {
	//	journaler.Debug("Setting bytes")
	b.bs = b.m.a.Get(b.s.Offset, b.s.Size)
	//	journaler.Debug("Bytes set! %p / %d", b, len(b.bs))
}

// Bytes are the current bytes
func (b *Backend) Bytes() []byte {
	//	journaler.Debug("Getting bytes! %p / %v", b, b == nil)
	return b.bs
}

// Section will return a section
func (b *Backend) Section() allocator.Section {
	return b.s
}

func (b *Backend) allocate(sz int64) (bs []byte) {
	var (
		ns   allocator.Section
		grew bool
	)

	if ns, grew = b.m.a.Allocate(sz); grew {
		b.SetBytes()
	}

	bs = b.m.a.Get(ns.Offset, ns.Size)

	if b.s.Size > 0 {
		// Copy old bytes to new byteslice
		copy(bs, b.bs)
		// Release old bytes to allocator
		b.m.a.Release(b.s)
	}

	b.s = ns
	b.bs = bs
	return
}

// Grow will grow the backend
func (b *Backend) Grow(sz int64) (bs []byte) {
	if !b.s.IsEmpty() {
		bs = b.m.a.Get(b.s.Offset, b.s.Size)
	}

	var cap int64
	if cap = nextCap(b.s.Size, sz); cap == -1 {
		return
	}

	bs = b.allocate(cap)
	return
}

// Notify will notify the parent
func (b *Backend) Notify() {
	b.m.Set(b)
}

// Dup will duplicate a backend
func (b *Backend) Dup() (out *Backend) {
	out = New(b.m)
	b.m.a.OnGrow(out.onGrow)
	out.Grow(b.s.Size)
	b.SetBytes()
	journaler.Debug("Oh yea? %#v / %d", b.s, len(b.bs), len(out.bs))
	copy(out.bs, b.bs)
	return
}

// Destroy will destroy a backend and it's contents
func (b *Backend) Destroy() (err error) {
	if !b.closed.Set(true) {
		return errors.ErrIsClosed
	}

	b.m.a.Release(b.s)
	b.m = nil
	return
}

// Close will close an Backend
func (b *Backend) Close() (err error) {
	if !b.closed.Set(true) {
		return errors.ErrIsClosed
	}

	b.m = nil
	return
}