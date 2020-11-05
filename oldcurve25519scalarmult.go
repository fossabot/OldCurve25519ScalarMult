// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package OldCurve25519ScalarMult

// OldScalarMult sets dst to the product scalar * point.
func OldScalarMult(dst, scalar, point *[32]byte) {
	oldScalarMult(dst, scalar, point)
}

// OldScalarBaseMult sets dst to the product scalar * base where base is the
// standard generator.
func OldScalarBaseMult(dst, scalar *[32]byte) {
	OldScalarMult(dst, scalar, &basePoint)
}

const (
	// ScalarSize is the size of the scalar input to X25519.
	ScalarSize = 32
	// PointSize is the size of the point input to X25519.
	PointSize = 32
)

// Basepoint is the canonical Curve25519 generator.
var Basepoint []byte

var basePoint = [32]byte{9, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

