// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package xlsx

import (
	"testing"

	"github.com/roboninc/xlsx/internal/ml"
	"github.com/roboninc/xlsx/internal/ml/primitives"
	pageOptions "github.com/roboninc/xlsx/types/options/page"
	printOptions "github.com/roboninc/xlsx/types/options/print"
	options "github.com/roboninc/xlsx/types/options/sheet"
	"github.com/stretchr/testify/require"
)

func TestSheetInfo(t *testing.T) {
	require.Equal(t, true, isCellEmpty(nil))
	require.Equal(t, true, isCellEmpty(&ml.Cell{}))
	require.Equal(t, true, isCellEmpty(&ml.Cell{Ref: "A10"}))
	require.Equal(t, false, isCellEmpty(&ml.Cell{Ref: "A10", Value: "1"}))

	require.Equal(t, true, isRowEmpty(nil))
	require.Equal(t, true, isRowEmpty(&ml.Row{}))
	require.Equal(t, true, isRowEmpty(&ml.Row{Ref: 1}))
	require.Equal(t, true, isRowEmpty(&ml.Row{Ref: 1, Cells: []*ml.Cell{}}))
	require.Equal(t, false, isRowEmpty(&ml.Row{Cells: []*ml.Cell{{}}}))
	require.Equal(t, false, isRowEmpty(&ml.Row{CustomHeight: true}))

	xl, err := Open("./test_files/example_simple.xlsx")
	if err != nil {
		panic(err)
	}

	defer xl.Close()
	sheet := xl.Sheet(0)

	//test options
	o := options.New(
		options.Visibility(options.VisibilityVeryHidden),
	)

	require.Equal(t, primitives.VisibilityType(0), xl.workbook.ml.Sheets[0].State)
	sheet.SetOptions(o)
	require.Equal(t, options.VisibilityVeryHidden, xl.workbook.ml.Sheets[0].State)

	//test print options
	pro := printOptions.New(
		printOptions.HorizontalCentered(true),
	)

	require.Nil(t, xl.sheets[0].ml.PrintOptions)
	sheet.SetPrintOptions(pro)
	require.Equal(t, true, xl.sheets[0].ml.PrintOptions.HorizontalCentered)

	//test page setup
	pgo := pageOptions.New(
		pageOptions.Orientation(pageOptions.OrientationLandscape),
	)

	require.Nil(t, xl.sheets[0].ml.PageSetup)
	sheet.SetPageSetup(pgo)
	require.Equal(t, pageOptions.OrientationLandscape, xl.sheets[0].ml.PageSetup.Orientation)

	//test set active
	require.Equal(t, 0, xl.workbook.ml.BookViews.Items[0].ActiveTab)
	sheet = xl.AddSheet("test")
	sheet.SetActive()
	require.Equal(t, 1, xl.workbook.ml.BookViews.Items[0].ActiveTab)
}
