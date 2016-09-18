// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions
// are met:
//
// * Redistributions of source code must retain the above copyright
// notice, this list of conditions and the following disclaimer.
//
// * Redistributions in binary form must reproduce the above
// copyright notice, this list of conditions and the following
// disclaimer in the documentation and/or other materials provided
// with the distribution.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
// "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
// LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
// A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
// OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
// SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
// LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
// DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
// THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
//

// WARNING: This file has automatically been generated on Fri, 23 Sep 2016 20:10:24 UTC.
// By https://git.io/cgogen. DO NOT EDIT.

package yubikey

/*
#cgo LDFLAGS: -lyubikey
#include "yubikey.h"
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"
import (
	"sync"
	"unsafe"
)

// cgoAllocMap stores pointers to C allocated memory for future reference.
type cgoAllocMap struct {
	mux sync.RWMutex
	m   map[unsafe.Pointer]struct{}
}

var cgoAllocsUnknown = new(cgoAllocMap)

func (a *cgoAllocMap) Add(ptr unsafe.Pointer) {
	a.mux.Lock()
	if a.m == nil {
		a.m = make(map[unsafe.Pointer]struct{})
	}
	a.m[ptr] = struct{}{}
	a.mux.Unlock()
}

func (a *cgoAllocMap) IsEmpty() bool {
	a.mux.RLock()
	isEmpty := len(a.m) == 0
	a.mux.RUnlock()
	return isEmpty
}

func (a *cgoAllocMap) Borrow(b *cgoAllocMap) {
	if b == nil || b.IsEmpty() {
		return
	}
	b.mux.Lock()
	a.mux.Lock()
	for ptr := range b.m {
		if a.m == nil {
			a.m = make(map[unsafe.Pointer]struct{})
		}
		a.m[ptr] = struct{}{}
		delete(b.m, ptr)
	}
	a.mux.Unlock()
	b.mux.Unlock()
}

func (a *cgoAllocMap) Free() {
	a.mux.Lock()
	for ptr := range a.m {
		C.free(ptr)
		delete(a.m, ptr)
	}
	a.mux.Unlock()
}

// allocTokenStMemory allocates memory for type *C.yubikey_token_st in C.
// The caller is responsible for freeing the this memory via C.free.
func allocTokenStMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfTokenStValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfTokenStValue = unsafe.Sizeof([1]*C.yubikey_token_st{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *Token) Ref() *C.yubikey_token_st {
	if x == nil {
		return nil
	}
	return x.ref33654d9c
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *Token) Free() {
	if x != nil && x.allocs33654d9c != nil {
		x.allocs33654d9c.(*cgoAllocMap).Free()
		x.ref33654d9c = nil
	}
}

// NewTokenRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewTokenRef(ref unsafe.Pointer) *Token {
	if ref == nil {
		return nil
	}
	obj := new(Token)
	ref33654d9c := (*C.yubikey_token_t)(unsafe.Pointer(ref))
	obj.ref33654d9c = (*C.yubikey_token_st)(unsafe.Pointer(ref33654d9c))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *Token) PassRef() (*C.yubikey_token_st, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref33654d9c != nil {
		return x.ref33654d9c, nil
	}
	mem33654d9c := allocTokenStMemory(1)
	ref33654d9c := (*C.yubikey_token_st)(mem33654d9c)
	allocs33654d9c := new(cgoAllocMap)
	var cuid_allocs *cgoAllocMap
	ref33654d9c.uid, cuid_allocs = *(*[6]C.uint8_t)(unsafe.Pointer(&x.Uid)), cgoAllocsUnknown
	allocs33654d9c.Borrow(cuid_allocs)

	var cctr_allocs *cgoAllocMap
	ref33654d9c.ctr, cctr_allocs = (C.uint16_t)(x.Ctr), cgoAllocsUnknown
	allocs33654d9c.Borrow(cctr_allocs)

	var ctstpl_allocs *cgoAllocMap
	ref33654d9c.tstpl, ctstpl_allocs = (C.uint16_t)(x.Tstpl), cgoAllocsUnknown
	allocs33654d9c.Borrow(ctstpl_allocs)

	var ctstph_allocs *cgoAllocMap
	ref33654d9c.tstph, ctstph_allocs = (C.uint8_t)(x.Tstph), cgoAllocsUnknown
	allocs33654d9c.Borrow(ctstph_allocs)

	var cuse_allocs *cgoAllocMap
	ref33654d9c.use, cuse_allocs = (C.uint8_t)(x.Use), cgoAllocsUnknown
	allocs33654d9c.Borrow(cuse_allocs)

	var crnd_allocs *cgoAllocMap
	ref33654d9c.rnd, crnd_allocs = (C.uint16_t)(x.Rnd), cgoAllocsUnknown
	allocs33654d9c.Borrow(crnd_allocs)

	var ccrc_allocs *cgoAllocMap
	ref33654d9c.crc, ccrc_allocs = (C.uint16_t)(x.Crc), cgoAllocsUnknown
	allocs33654d9c.Borrow(ccrc_allocs)

	x.ref33654d9c = ref33654d9c
	x.allocs33654d9c = allocs33654d9c
	return ref33654d9c, allocs33654d9c

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x Token) PassValue() (C.yubikey_token_st, *cgoAllocMap) {
	if x.ref33654d9c != nil {
		return *x.ref33654d9c, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *Token) Deref() {
	if x.ref33654d9c == nil {
		return
	}
	x.Uid = *(*[6]byte)(unsafe.Pointer(&x.ref33654d9c.uid))
	x.Ctr = (uint16)(x.ref33654d9c.ctr)
	x.Tstpl = (uint16)(x.ref33654d9c.tstpl)
	x.Tstph = (byte)(x.ref33654d9c.tstph)
	x.Use = (byte)(x.ref33654d9c.use)
	x.Rnd = (uint16)(x.ref33654d9c.rnd)
	x.Crc = (uint16)(x.ref33654d9c.crc)
}

type sliceHeader struct {
	Data uintptr
	Len  int
	Cap  int
}

// unpackPCharString represents the data from Go string as *C.char and avoids copying.
func unpackPCharString(str string) (*C.char, *cgoAllocMap) {
	h := (*stringHeader)(unsafe.Pointer(&str))
	return (*C.char)(unsafe.Pointer(h.Data)), cgoAllocsUnknown
}

type stringHeader struct {
	Data uintptr
	Len  int
}

// unpackPUint8String represents the data from Go string as *C.uint8_t and avoids copying.
func unpackPUint8String(str string) (*C.uint8_t, *cgoAllocMap) {
	h := (*stringHeader)(unsafe.Pointer(&str))
	return (*C.uint8_t)(unsafe.Pointer(h.Data)), cgoAllocsUnknown
}
