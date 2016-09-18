# libyubikey

[![GoDoc](https://godoc.org/github.com/jfrazelle/libyubikey?status.svg)](https://godoc.org/github.com/jfrazelle/libyubikey) [![Travis CI](https://travis-ci.org/jfrazelle/libyubikey.svg?branch=master)](https://travis-ci.org/jfrazelle/libyubikey)

Go bindings for libyubikey so you can write Go to interact with your yubikeys.
The C library lives at
[yubico/yubico-c](https://github.com/Yubico/yubico-c).

## C Libraries Required for Compilation

- libcrypto
- [liblibyubikey](https://github.com/Yubico/yubico-piv-tool/)
- [libpcsclite](https://pcsclite.alioth.debian.org/pcsclite.html)
- [ccid](https://pcsclite.alioth.debian.org/ccid.html)

## Example

```go
package main

import (
	"fmt"
	"log"

	"github.com/jfrazelle/libyubikey"
)

func main() {
	s := libyubikey.NewState()
	defer s.Free()

	// Let's get the readers
	readers := make([]byte, 2048)
	len := []uint{2048}
	log.Println("list")
	err := libyubikey.ListReaders([]libyubikey.State{*s}, readers, len)
	if err != 0 {
		log.Fatalf("%s: %#v", libyubikey.Strerror(err), err)
	}

	fmt.Printf("readers: %s\n", string(readers))
}
```

## Starting `pcscd`

Hopefully your operating system does this for you with a nice init script but
if not here you go:

```console
$ sudo LIBCCID_ifdLogLevel=0x000F /usr/sbin/pcscd --foreground --debug --apdu --color
$ sudo /usr/sbin/pcscd --hotplug
```
