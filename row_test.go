// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package xlsx

import (
	"testing"

	"github.com/roboninc/xlsx/format/styles"
	"github.com/stretchr/testify/require"
)

func TestRow(t *testing.T) {
	xl, err := Open("./test_files/example_simple.xlsx")
	if err != nil {
		panic(err)
	}

	defer xl.Close()
	sheet := xl.Sheet(0)
	r := sheet.Row(5)

	o := options.New(
		options.Height(0),
		options.OutlineLevel(10),
		options.Hidden(true),
		options.Phonetic(true),
		options.Collapsed(true),
	)

	r.SetOptions(o)

	require.Equal(t, r.ml.OutlineLevel, o.OutlineLevel)
	require.Equal(t, r.ml.Hidden, o.Hidden)
	require.Equal(t, r.ml.Phonetic, o.Phonetic)
	require.Equal(t, r.ml.Collapsed, o.Collapsed)
	require.Equal(t, r.ml.Height, float32(0.0))
	require.Equal(t, r.ml.CustomHeight, false)
	require.Equal(t, r.ml.CustomFormat, false)
	require.Equal(t, r.ml.Style, styles.DirectStyleID(0))

	o = options.New(
		options.Height(100.0),
	)

	r.SetOptions(o)
	require.Equal(t, r.ml.OutlineLevel, o.OutlineLevel)
	require.Equal(t, r.ml.Hidden, o.Hidden)
	require.Equal(t, r.ml.Phonetic, o.Phonetic)
	require.Equal(t, r.ml.Collapsed, o.Collapsed)
	require.Equal(t, r.ml.Height, float32(100.0))
	require.Equal(t, r.ml.CustomHeight, true)
	require.Equal(t, r.ml.CustomFormat, false)
	require.Equal(t, r.ml.Style, styles.DirectStyleID(0))

	style := styles.New(
		styles.Font.Name("Calibri"),
		styles.Font.Size(12),
	)

	styleRef := xl.AddStyles(style)
	r.SetStyles(styleRef)

	require.Equal(t, r.ml.CustomFormat, true)
	require.Equal(t, r.ml.Style, styles.DirectStyleID(styleRef))
}
