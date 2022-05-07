// Copyright 2022 phelmkamp. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

// Package dontpanic provides panic-free alternatives to common operations.
package dontpanic

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// Deref invokes *p.
// Returns an error if p is nil.
func Deref[T any](p *T) (_ T, err error) {
	defer Recover(&err)
	return *p, nil
}

// Div returns x / y.
// Returns an error if y is zero.
func Div[N constraints.Integer | constraints.Float | constraints.Complex](x, y N) (_ N, err error) {
	defer Recover(&err)
	return x / y, nil
}

// Mod returns x % y.
// Returns an error if y is zero.
func Mod[N constraints.Integer](x, y N) (_ N, err error) {
	defer Recover(&err)
	return x % y, nil
}

// Send invokes ch <- v.
// Returns an error if ch is closed.
func Send[T any](ch chan T, v T) (err error) {
	defer Recover(&err)
	ch <- v
	return nil
}

// Close invokes close(ch).
// Returns an error if ch is nil or closed.
func Close[T any](ch chan T) (err error) {
	defer Recover(&err)
	close(ch)
	return nil
}

// SetMapIndex sets m[k] = v.
// Returns an error if m is nil.
func SetMapIndex[K comparable, V any, M ~map[K]V](m M, k K, v V) (err error) {
	defer Recover(&err)
	m[k] = v
	return nil
}

// Make invokes make(S, size...).
// Returns an error if length is negative or larger than capacity.
//
// Supports empty:
//  MakeSlice()
// or length n:
//  MakeSlice(n)
// or length n and capacity m:
//  MakeSlice(n, m)
func MakeSlice[E any, S ~[]E](size ...int) (_ S, err error) {
	defer Recover(&err)
	switch len(size) {
	case 0:
		return nil, nil
	case 1:
		return make([]E, size[0]), nil
	case 2:
		return make([]E, size[0], size[1]), nil
	default:
		return nil, fmt.Errorf("size: expected 0-2 arguments; found %d", len(size))
	}
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
// Returns an error if indexes are out of range or too many are specified.
//
// Supports s[:]
//  Slice(s)
// or s[i:]
//  Slice(s, i)
// or s[i:j]
//  Slice(s, i, j)
// or s[i:j:k]
//  Slice(s, i, j, k)
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

// StringIndex returns s[i].
// Returns an error if i is out of range.
func StringIndex(s string, i int) (_ byte, err error) {
	defer Recover(&err)
	return s[i], nil
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
