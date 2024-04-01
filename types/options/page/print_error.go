// Copyright (c) 2024 ROBON INC.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package options

import (
	"github.com/roboninc/xlsx/internal/ml/primitives"
)

// List of all possible values for PrintErrorType
const (
	_ primitives.PrintErrorType = iota
	PrintErrorDisplayed
	PrintErrorBlank
	PrintErrorDash
	PrintErrorNA
)

func init() {
	primitives.FromPrintErrorType = map[primitives.PrintErrorType]string{
		PrintErrorDisplayed: "displayed",
		PrintErrorBlank:     "blank",
		PrintErrorDash:      "dash",
		PrintErrorNA:        "NA",
	}

	primitives.ToPrintErrorType = make(map[string]primitives.PrintErrorType, len(primitives.FromPrintErrorType))
	for k, v := range primitives.FromPrintErrorType {
		primitives.ToPrintErrorType[v] = k
	}
}
