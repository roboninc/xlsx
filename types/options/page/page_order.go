// Copyright (c) 2024 ROBON INC.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package options

import (
	"github.com/roboninc/xlsx/internal/ml/primitives"
)

// List of all possible values for PageOrderType
const (
	_ primitives.PageOrderType = iota
	PageOrderDownThenOver
	PageOrderOverThenDown
)

func init() {
	primitives.FromPageOrderType = map[primitives.PageOrderType]string{
		PageOrderDownThenOver: "downThenOver",
		PageOrderOverThenDown: "overThenDown",
	}

	primitives.ToPageOrderType = make(map[string]primitives.PageOrderType, len(primitives.FromPageOrderType))
	for k, v := range primitives.FromPageOrderType {
		primitives.ToPageOrderType[v] = k
	}
}
