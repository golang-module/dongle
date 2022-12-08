// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package md2

import (
	"testing"
)

var bench, buf = New(), makeBuf()

func makeBuf() []byte {
	b := make([]byte, 8<<10)
	for i := range b {
		b[i] = byte(i)
	}
	return b
}

func BenchmarkHash1K(b *testing.B) {
	b.SetBytes(1024)
	for i := 0; i < b.N; i++ {
		bench.Write(buf[:1024])
	}
}

func BenchmarkHash8K(b *testing.B) {
	b.SetBytes(int64(len(buf)))
	for i := 0; i < b.N; i++ {
		bench.Write(buf)
	}
}
