// Copyright 2026 Harald Albrecht.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may not
// use this file except in compliance with the License. You may obtain a copy
// of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

//go:build go1.27

package opt

// If returns a bool condition value that is to be chained with Then and finally
// Else, providing type safety for the “then” and “else” values.
func If(cond bool) ifthen { return ifthen(cond) }

type ifthen bool

// Then specifies the value eventually be returned by an If.Then.Else chain in
// case the condition value is true.
func (i ifthen) Then[T any](v T) orelse[T] {
	return orelse[T]{cond: bool(i), then: v}
}

type orelse[T any] struct {
	cond bool
	then T
}

// Else specifies the value eventually to be returned by an If.Then.Else chain
// in case the condition value is false.
func (e orelse[T]) Else(v T) T {
	if e.cond {
		return e.then
	}
	return v
}
