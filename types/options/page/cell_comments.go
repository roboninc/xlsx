// Copyright (c) 2024 ROBON INC.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package options

import (
	"github.com/roboninc/xlsx/internal/ml/primitives"
)

// List of all possible values for CellCommentsType
const (
	_ primitives.CellCommentsType = iota
	CellCommentsNone
	CellCommentsAsDisplayed
	CellCommentsAtEnd
)

func init() {
	primitives.FromCellCommentsType = map[primitives.CellCommentsType]string{
		CellCommentsNone:        "none",
		CellCommentsAsDisplayed: "asDisplayed",
		CellCommentsAtEnd:       "atEnd",
	}

	primitives.ToCellCommentsType = make(map[string]primitives.CellCommentsType, len(primitives.FromCellCommentsType))
	for k, v := range primitives.FromCellCommentsType {
		primitives.ToCellCommentsType[v] = k
	}
}
