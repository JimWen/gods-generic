// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/JimWen/gods-generic/maps/treebidimap"
	"github.com/JimWen/gods-generic/utils"
)

// TreeBidiMapExample to demonstrate basic usage of TreeBidiMap
func main() {
	m := treebidimap.NewWith[int, string](utils.NumberComparator[int], utils.StringComparator)
	m.Put(1, "x")        // 1->x
	m.Put(3, "b")        // 1->x, 3->b (ordered)
	m.Put(1, "a")        // 1->a, 3->b (ordered)
	m.Put(2, "b")        // 1->a, 2->b (ordered)
	_, _ = m.GetKey("a") // 1, true
	_, _ = m.Get(2)      // b, true
	_, _ = m.Get(3)      // nil, false
	_ = m.Values()       // []string{"a", "b"} (ordered)
	_ = m.Keys()         // []int{1, 2} (ordered)
	m.Remove(1)          // 2->b
	m.Clear()            // empty
	m.Empty()            // true
	m.Size()             // 0
}
