// Copyright (c) 2024 ROBON INC.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package options

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPrinntOptions(t *testing.T) {
	o := New(
		PaperSize(2),
		Scale(200),
		FirstPageNumber(2),
		FitToWidth(2),
		FitToHeight(2),
		PageOrder(PageOrderOverThenDown),
		Orientation(OrientationLandscape),
		UsePrinterDefaults(false),
		BlackAndWhite(true),
		Draft(true),
		CellComments(CellCommentsAsDisplayed),
		UseFirstPageNumber(true),
		Errors(PrintErrorDash),
		HorizontalDpi(1200),
		VerticalDpi(1200),
		Copies(2),
	)

	require.IsType(t, &Info{}, o)
	require.Equal(t, &Info{
		PaperSize:          toIntPtr(2),
		Scale:              toIntPtr(200),
		FirstPageNumber:    toIntPtr(2),
		FitToWidth:         toIntPtr(2),
		FitToHeight:        toIntPtr(2),
		PageOrder:          PageOrderOverThenDown,
		Orientation:        OrientationLandscape,
		UsePrinterDefaults: toBoolPtr(false),
		BlackAndWhite:      true,
		Draft:              true,
		CellComments:       CellCommentsAsDisplayed,
		UseFirstPageNumber: true,
		Errors:             PrintErrorDash,
		HorizontalDpi:      toIntPtr(1200),
		VerticalDpi:        toIntPtr(1200),
		Copies:             toIntPtr(2),
	}, o)

	o = New(
		PaperSize(1),
		Scale(100),
		FirstPageNumber(1),
		FitToWidth(1),
		FitToHeight(1),
		PageOrder(PageOrderDownThenOver),
		Orientation(OrientationDefault),
		UsePrinterDefaults(true),
		BlackAndWhite(false),
		Draft(false),
		CellComments(CellCommentsNone),
		UseFirstPageNumber(false),
		Errors(PrintErrorDisplayed),
		HorizontalDpi(600),
		VerticalDpi(600),
		Copies(1),
	)

	require.IsType(t, &Info{}, o)
	require.Equal(t, &Info{}, o)
}

func toIntPtr(i int) *int {
	return &i
}

func toBoolPtr(b bool) *bool {
	return &b
}
