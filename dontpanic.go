// Copyright 2022 phelmkamp. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

// Package dontpanic provides panic-free alternatives to common operations.
package dontpanic

import (
	"fmt"
)

// SetMapIndex sets m[k] = v.
// Returns an error if m is nil.
func SetMapIndex[K comparable, V any, M ~map[K]V](m M, k K, v V) (err error) {
	defer Recover(&err)
	m[k] = v
	return nil
}

// SetSliceIndex sets s[i] = v.
// Returns an error if i is out of range.
func SetSliceIndex[E any, S ~[]E](s S, i int, v E) (err error) {
	defer Recover(&err)
	s[i] = v
	return nil
}

// SliceIndex returns s[i].
// Returns an error if i is out of range.
func SliceIndex[E any, S ~[]E](s S, i int) (_ E, err error) {
	defer Recover(&err)
	return s[i], nil
}

// Slice returns s[i:j:k].
//
// Supports [:]
//  Slice(s)
// or [i:]
//  Slice(s, i)
// or [i:j]
//  Slice(s, i, j)
// or [i:j:k]
//  Slice(s, i, j, k)
// Returns an error if indexes are out of range or too many are specified.
func Slice[E any, S ~[]E](s S, ijk ...int) (_ S, err error) {
	defer Recover(&err)
	switch len(ijk) {
	case 0:
		return s[:], nil
	case 1:
		return s[ijk[0]:], nil
	case 2:
		return s[ijk[0]:ijk[1]], nil
	case 3:
		return s[ijk[0]:ijk[1]:ijk[2]], nil
	default:
		return nil, fmt.Errorf("ijk: expected 0-3 indexes; found %d", len(ijk))
	}
}

// Recover calls recover() and writes the result to err.
func Recover(err *error) {
	switch r := recover().(type) {
	case nil:
		return
	case error:
		*err = r
	default:
		*err = fmt.Errorf("%v", r)
	}
}
