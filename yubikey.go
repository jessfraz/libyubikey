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
import "unsafe"

// Parse function as declared in libyubikey/yubikey.h:76
func Parse(token *[16]byte, key *[16]byte, out *Token) {
	ctoken, _ := *(**C.uint8_t)(unsafe.Pointer(&token)), cgoAllocsUnknown
	ckey, _ := *(**C.uint8_t)(unsafe.Pointer(&key)), cgoAllocsUnknown
	cout, _ := out.PassRef()
	C.yubikey_parse(ctoken, ckey, cout)
}

// Generate function as declared in libyubikey/yubikey.h:81
func Generate(token *Token, key *[16]byte, out *[32]byte) {
	ctoken, _ := token.PassRef()
	ckey, _ := *(**C.uint8_t)(unsafe.Pointer(&key)), cgoAllocsUnknown
	cout, _ := *(**C.char)(unsafe.Pointer(&out)), cgoAllocsUnknown
	C.yubikey_generate(ctoken, ckey, cout)
}

// ModhexEncode function as declared in libyubikey/yubikey.h:100
func ModhexEncode(dst []byte, src string, srcsize uint) {
	cdst, _ := (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&dst)).Data)), cgoAllocsUnknown
	csrc, _ := unpackPCharString(src)
	csrcsize, _ := (C.size_t)(srcsize), cgoAllocsUnknown
	C.yubikey_modhex_encode(cdst, csrc, csrcsize)
}

// ModhexDecode function as declared in libyubikey/yubikey.h:106
func ModhexDecode(dst []byte, src string, dstsize uint) {
	cdst, _ := (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&dst)).Data)), cgoAllocsUnknown
	csrc, _ := unpackPCharString(src)
	cdstsize, _ := (C.size_t)(dstsize), cgoAllocsUnknown
	C.yubikey_modhex_decode(cdst, csrc, cdstsize)
}

// HexEncode function as declared in libyubikey/yubikey.h:110
func HexEncode(dst []byte, src string, srcsize uint) {
	cdst, _ := (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&dst)).Data)), cgoAllocsUnknown
	csrc, _ := unpackPCharString(src)
	csrcsize, _ := (C.size_t)(srcsize), cgoAllocsUnknown
	C.yubikey_hex_encode(cdst, csrc, csrcsize)
}

// HexDecode function as declared in libyubikey/yubikey.h:111
func HexDecode(dst []byte, src string, dstsize uint) {
	cdst, _ := (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&dst)).Data)), cgoAllocsUnknown
	csrc, _ := unpackPCharString(src)
	cdstsize, _ := (C.size_t)(dstsize), cgoAllocsUnknown
	C.yubikey_hex_decode(cdst, csrc, cdstsize)
}

// ModhexP function as declared in libyubikey/yubikey.h:115
func ModhexP(str string) int32 {
	cstr, _ := unpackPCharString(str)
	__ret := C.yubikey_modhex_p(cstr)
	__v := (int32)(__ret)
	return __v
}

// HexP function as declared in libyubikey/yubikey.h:116
func HexP(str string) int32 {
	cstr, _ := unpackPCharString(str)
	__ret := C.yubikey_hex_p(cstr)
	__v := (int32)(__ret)
	return __v
}

// Crc16 function as declared in libyubikey/yubikey.h:123
func Crc16(buf string, bufSize uint) uint16 {
	cbuf, _ := unpackPUint8String(buf)
	cbufSize, _ := (C.size_t)(bufSize), cgoAllocsUnknown
	__ret := C.yubikey_crc16(cbuf, cbufSize)
	__v := (uint16)(__ret)
	return __v
}

// AesDecrypt function as declared in libyubikey/yubikey.h:129
func AesDecrypt(state []byte, key string) {
	cstate, _ := (*C.uint8_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&state)).Data)), cgoAllocsUnknown
	ckey, _ := unpackPUint8String(key)
	C.yubikey_aes_decrypt(cstate, ckey)
}

// AesEncrypt function as declared in libyubikey/yubikey.h:130
func AesEncrypt(state []byte, key string) {
	cstate, _ := (*C.uint8_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&state)).Data)), cgoAllocsUnknown
	ckey, _ := unpackPUint8String(key)
	C.yubikey_aes_encrypt(cstate, ckey)
}
