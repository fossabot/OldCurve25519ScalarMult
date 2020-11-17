// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build !amd64 gccgo appengine purego

package OldCurve25519ScalarMult

// OldScalarMult -> noasm
func OldScalarMult(out, in, base *[32]byte) {
	oldScalarMultGeneric(out, in, base)
}
