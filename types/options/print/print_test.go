// Copyright (c) 2024 ROBON INC.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package options

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPrinntOptions(t *testing.T) {
	gridLinesSet := true
	o := New(
		HorizontalCentered(true),
		VerticalCentered(false),
		Headings(true),
		GridLines(false),
		GridLinesSet(gridLinesSet),
	)

	require.IsType(t, &Info{}, o)
	require.Equal(t, &Info{
		HorizontalCentered: true,
		VerticalCentered:   false,
		Headings:           true,
		GridLines:          false,
	}, o)

	gridLinesSet = false
	o = New(
		GridLinesSet(gridLinesSet),
	)

	require.IsType(t, &Info{}, o)
	require.Equal(t, &Info{
		GridLinesSet: &gridLinesSet,
	}, o)
}
