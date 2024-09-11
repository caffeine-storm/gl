// Copyright 2012 The go-gl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gl

import (
	"testing"
	"unsafe"
)

func TestPtr(t *testing.T) {
	// test nil
	if p, q := unsafe.Pointer(nil), ptr(nil); p != q {
		t.Fatalf("expected %#v, got %#v\n", p, q)
	}

	// test nil interface
	var r interface{}
	if p, q := unsafe.Pointer(nil), ptr(r); p != q {
		t.Fatalf("expected %#v, got %#v\n", p, q)
	}

	// test nil pointer
	var s *int
	if p, q := unsafe.Pointer(nil), ptr(s); p != q {
		t.Fatalf("expected %#v, got %#v\n", p, q)
	}

	// test uintptr panics
	t.Run("ptr(uintptr(x)) should panic", func(t *testing.T) {
		defer func() {
			if e := recover(); e != nil {
				// succcess case
			} else {
				panic("expected a panic")
			}
		}()

		ptr(uintptr(42))

		t.Fatalf("we must not reach this statement")
	})
}

func TestExportTypeDefs(t *testing.T) {
	var x GLchar
	x = GLchar(65) // 'A' == 65
	if x != 'A' {
		t.Fatalf("expected GLchar to behave like a C.char\n")
	}
}

func TestGetGlVersion(t *testing.T) {
	MustInit()
	ver := getVersion()

	if ver == "" {
		t.Fatalf("expected a non-empty GL version string\n")
	}
}
