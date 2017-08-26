package xlsx_test

import (
	"fmt"
	"github.com/plandem/xlsx"
	"github.com/plandem/xlsx/format"
	"github.com/plandem/xlsx/types"
	"log"
	"os"
	"strings"
)

func ExampleNew() {
	xl := xlsx.New()

	//... add a new content

	xl.SaveAs("new_file.xlsx")
}

func ExampleOpen() {
	xl, err := xlsx.Open("./test_files/example_simple.xlsx")
	if err != nil {
		log.Fatal(err)
	}

	defer xl.Close()
}

func ExampleOpenStream() {
	zipFile, err := os.Open("./test_files/example_simple.xlsx")
	if err != nil {
		log.Fatal(err)
	}

	xl, err := xlsx.OpenStream(zipFile)
	if err != nil {
		log.Fatal(err)
	}

	_ = xl
}

func ExampleSave() {
	xl, err := xlsx.Open("./test_files/example_simple.xlsx")
	if err != nil {
		log.Fatal(err)
	}

	defer xl.Close()

	//... change content

	err = xl.Save()
	if err != nil {
		log.Fatal(err)
	}
}

func ExampleSaveAs() {
	xl, err := xlsx.Open("./test_files/example_simple.xlsx")
	if err != nil {
		log.Fatal(err)
	}

	defer xl.Close()

	//... change content

	err = xl.SaveAs("new_file.xlsx")
	if err != nil {
		log.Fatal(err)
	}
}

func ExampleNewColumnOptions() {
	options := xlsx.NewColumnOptions(
		xlsx.ColumnOption.OutlineLevel(5),
		xlsx.ColumnOption.Hidden(true),
		xlsx.ColumnOption.Phonetic(true),
		xlsx.ColumnOption.Formatting(12345),
		xlsx.ColumnOption.Width(45.5),
	)

	_ = options
}

func ExampleNewRowOptions() {
	options := xlsx.NewRowOptions(
		xlsx.RowOption.OutlineLevel(5),
		xlsx.RowOption.Hidden(true),
		xlsx.RowOption.Phonetic(true),
		xlsx.RowOption.Formatting(12345),
		xlsx.RowOption.Height(45.5),
	)

	_ = options
}

func ExampleSpreadsheet_GetSheetNames() {
	xl, err := xlsx.Open("./test_files/example_simple.xlsx")
	if err != nil {
		log.Fatal(err)
	}

	defer xl.Close()

	fmt.Println(xl.GetSheetNames())
	//Output:
	// [Sheet1]
}

func ExampleSpreadsheet_Sheet() {
	xl, err := xlsx.Open("./test_files/example_simple.xlsx")
	if err != nil {
		log.Fatal(err)
	}

	defer xl.Close()

	//nil, if there is no sheet with requested index
	if sheet := xl.Sheet(12345); sheet == nil {
		fmt.Println("Unknown sheet")
	}

	if sheet := xl.Sheet(0); sheet != nil {
		fmt.Println(sheet.Name())
	}

	//Output:
	// Unknown sheet
	// Sheet1
}

func ExampleSpreadsheet_AddSheet() {
	xl, err := xlsx.Open("./test_files/example_simple.xlsx")
	if err != nil {
		log.Fatal(err)
	}

	defer xl.Close()

	sheet := xl.AddSheet("New sheet")

	//now you can use sheet as always
	fmt.Println(sheet.Name())

	//Output:
	// New sheet
}

func ExampleSpreadsheet_SetActive() {
	xl, err := xlsx.Open("./test_files/example_simple.xlsx")
	if err != nil {
		log.Fatal(err)
	}

	defer xl.Close()

	//add a new sheet, next index is 1
	xl.AddSheet("New sheet")

	//set sheet with index 1 as active
	xl.SetActive(1)
}

func ExampleSpreadsheet_DeleteSheet() {
	xl, err := xlsx.Open("./test_files/example_simple.xlsx")
	if err != nil {
		log.Fatal(err)
	}

	defer xl.Close()

	//add a new sheet, next index is 1
	xl.AddSheet("New sheet")

	//delete a sheet with index 0
	xl.DeleteSheet(0)
}

func ExampleSpreadsheet_AddFormatting() {
	xl, err := xlsx.Open("./test_files/example_simple.xlsx")
	if err != nil {
		log.Fatal(err)
	}

	defer xl.Close()

	//create a new format for a bold font with red color and yellow solid background
	redBold := format.New(
		format.Font.Bold,
		format.Font.Color("#ff0000"),
		format.Fill.Background("#ffff00"),
		format.Fill.Type(format.PatternTypeSolid),
	)

	//add formatting to xlsx
	styleId := xl.AddFormatting(redBold)

	//now you can use this id wherever you need
	_ = styleId
}

func ExampleSheet_Cell() {
	xl, err := xlsx.Open("./test_files/example_simple.xlsx")
	if err != nil {
		log.Fatal(err)
	}

	defer xl.Close()

	sheet := xl.Sheet(0)

	//get cell by 0-based indexes, e.g.: 13,27 is same as N28
	cell := sheet.Cell(13, 27)

	fmt.Println(cell.Value())
	//Output:
	// last cell
}

func ExampleSheet_CellByRef() {
	xl, err := xlsx.Open("./test_files/example_simple.xlsx")
	if err != nil {
		log.Fatal(err)
	}

	defer xl.Close()

	sheet := xl.Sheet(0)

	//get cell by reference, e.g.: N28 is same as 13,27
	cell := sheet.CellByRef("N28")

	fmt.Println(cell.Value())
	//Output:
	// last cell
}

func Example_access() {
	xl, err := xlsx.Open("./test_files/example_simple.xlsx")
	if err != nil {
		log.Fatal(err)
	}

	defer xl.Close()

	//get sheet by 0-based index
	sheet := xl.Sheet(0)

	//get cell by 0-based indexes
	cell := sheet.Cell(13, 27)
	fmt.Println(cell.Value())

	//get cell by reference
	cell = sheet.CellByRef("N28")
	fmt.Println(cell.Value())

	//get all cells of row for 0-based index
	row := sheet.Row(9)
	fmt.Println(strings.Join(row.Values(), ","))

	//get all cells of col for 0-based index
	col := sheet.Col(3)
	fmt.Println(strings.Join(col.Values(), ","))

	//get all cells of range
	area := sheet.Range("D10:H13")
	fmt.Println(strings.Join(area.Values(), ","))

	//Output:
	// last cell
	// last cell
	// ,,,1,2,3,4,5,,,,,
	// ,,,,,,,,,1,6,11,16,,,,,,,,,,,,,,
	// 1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20
}

func Example_update() {
	xl, err := xlsx.Open("./test_files/example_simple.xlsx")
	if err != nil {
		log.Fatal(err)
	}

	defer xl.Close()

	sheet := xl.Sheet(0)

	//update value of cell
	cell := sheet.Cell(13, 27)
	fmt.Println(cell.Value())
	cell.SetValue("new value")
	fmt.Println(cell.Value())

	//update value of cells in row
	row := sheet.Row(9)
	fmt.Println(strings.Join(row.Values(), ","))
	for i, c := range row.Cells() {
		c.SetValue(i)
	}
	fmt.Println(strings.Join(row.Values(), ","))

	//update value of cells in col
	col := sheet.Col(3)
	fmt.Println(strings.Join(col.Values(), ","))
	for i, c := range col.Cells() {
		c.SetValue(i)
	}
	fmt.Println(strings.Join(col.Values(), ","))

	//update value of cells in range
	area := sheet.Range("D10:H13")
	fmt.Println(strings.Join(area.Values(), ","))
	for i, c := range area.Cells() {
		c.SetValue(i)
	}
	fmt.Println(strings.Join(area.Values(), ","))

	//Output:
	// last cell
	// new value
	// ,,,1,2,3,4,5,,,,,
	// 0,1,2,3,4,5,6,7,8,9,10,11,12
	// ,,,,,,,,,3,6,11,16,,,,,,,,,,,,,,
	// 0,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26
	// 9,4,5,6,7,10,7,8,9,10,11,12,13,14,15,12,17,18,19,20
	// 0,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19
}

func Example_formatting() {
	xl, err := xlsx.Open("./test_files/example_simple.xlsx")
	if err != nil {
		log.Fatal(err)
	}

	defer xl.Close()

	//create a new format for a bold font with red color and yellow solid background
	redBold := format.New(
		format.Font.Bold,
		format.Font.Color("#ff0000"),
		format.Fill.Background("#ffff00"),
		format.Fill.Type(format.PatternTypeSolid),
	)

	//add formatting to xlsx
	styleId := xl.AddFormatting(redBold)

	sheet := xl.Sheet(0)

	//set formatting for cell
	sheet.CellByRef("N28").SetFormatting(styleId)

	//set DEFAULT formatting for row. Affects cells not yet allocated in the row.
	//In other words, this style applies to new cells.
	sheet.Row(9).Set(xlsx.NewRowOptions(xlsx.RowOption.Formatting(styleId)))

	//set DEFAULT formatting for col. Affects cells not yet allocated in the col.
	//In other words, this style applies to new cells.
	sheet.Col(3).Set(xlsx.NewColumnOptions(xlsx.ColumnOption.Formatting(styleId)))

	//set formatting for all cells in range
	sheet.Range("D10:H13").SetFormatting(styleId)
}

func Example_visibility() {
	xl, err := xlsx.Open("./test_files/example_simple.xlsx")
	if err != nil {
		log.Fatal(err)
	}

	defer xl.Close()

	sheet := xl.Sheet(0)

	//hide row
	sheet.Row(9).Set(xlsx.NewRowOptions(xlsx.RowOption.Hidden(true)))

	//hide col
	sheet.Col(3).Set(xlsx.NewColumnOptions(xlsx.ColumnOption.Hidden(true)))

	//hide sheet
	sheet.SetState(types.VisibilityTypeHidden)
}

func Example_append() {
	xl, err := xlsx.Open("./test_files/example_simple.xlsx")
	if err != nil {
		log.Fatal(err)
	}

	defer xl.Close()

	sheet := xl.Sheet(0)

	//to append a new col/row, simple request it - sheet will be auto expanded.
	//E.g.: we have 14 cols, 28 rows
	fmt.Println(sheet.TotalCols(), "x", sheet.TotalRows())

	//append 72 rows
	sheet.Row(99)
	fmt.Println(sheet.TotalCols(), "x", sheet.TotalRows())

	//append 36 cols
	sheet.Col(49)
	fmt.Println(sheet.TotalCols(), "x", sheet.TotalRows())

	//append 3 sheet
	fmt.Println(strings.Join(xl.GetSheetNames(), ","))
	xl.AddSheet("new sheet")
	xl.AddSheet("new sheet")
	xl.AddSheet("new sheet")
	fmt.Println(strings.Join(xl.GetSheetNames(), ","))

	//Output:
	// 14 x 28
	// 14 x 100
	// 50 x 100
	// Sheet1
	// Sheet1,new sheet,new sheet1,new sheet2
}

func Example_insert() {
	xl, err := xlsx.Open("./test_files/example_simple.xlsx")
	if err != nil {
		log.Fatal(err)
	}

	defer xl.Close()

	sheet := xl.Sheet(0)

	fmt.Println(sheet.TotalCols(), "x", sheet.TotalRows())
	fmt.Println(strings.Join(sheet.Col(3).Values(), ","))

	//insert a new col
	sheet.InsertCol(3)
	fmt.Println(sheet.TotalCols(), "x", sheet.TotalRows())
	fmt.Println(strings.Join(sheet.Col(3).Values(), ","))
	fmt.Println(strings.Join(sheet.Col(4).Values(), ","))

	//insert a new row
	fmt.Println(strings.Join(sheet.Row(9).Values(), ","))
	sheet.InsertRow(3)
	fmt.Println(sheet.TotalCols(), "x", sheet.TotalRows())
	fmt.Println(strings.Join(sheet.Row(9).Values(), ","))
	fmt.Println(strings.Join(sheet.Row(10).Values(), ","))

	//Output:
	// 14 x 28
	// ,,,,,,,,,1,6,11,16,,,,,,,,,,,,,,
	// 15 x 28
	// ,,,,,,,,,,,,,,,,,,,,,,,,,,
	// ,,,,,,,,,1,6,11,16,,,,,,,,,,,,,,
	// ,,,,1,2,3,4,5,,,,,
	// 15 x 29
	// ,,,,,,,,,,,,,
	// ,,,,1,2,3,4,5,,,,,
}

func Example_delete() {
	xl, err := xlsx.Open("./test_files/example_simple.xlsx")
	if err != nil {
		log.Fatal(err)
	}

	defer xl.Close()

	sheet := xl.Sheet(0)

	fmt.Println(sheet.TotalCols(), "x", sheet.TotalRows())

	//delete col
	fmt.Println(strings.Join(sheet.Col(3).Values(), ","))
	sheet.DeleteCol(3)
	fmt.Println(sheet.TotalCols(), "x", sheet.TotalRows())
	fmt.Println(strings.Join(sheet.Col(3).Values(), ","))

	//delete row
	fmt.Println(strings.Join(sheet.Row(3).Values(), ","))
	sheet.DeleteRow(3)
	fmt.Println(sheet.TotalCols(), "x", sheet.TotalRows())
	fmt.Println(strings.Join(sheet.Row(3).Values(), ","))

	//delete sheet
	fmt.Println(strings.Join(xl.GetSheetNames(), ","))
	xl.DeleteSheet(0)
	fmt.Println(strings.Join(xl.GetSheetNames(), ","))
	//Output:
	// 14 x 28
	// ,,,,,,,,,1,6,11,16,,,,,,,,,,,,,,
	// 13 x 28
	// ,merged cols,,merged rows+cols,,,,,,2,7,12,17,,,,,,,,,,,,,,
	// ,,merged rows,merged rows+cols,,,,,,,,
	// 13 x 27
	// with trailing space   ,,merged rows,,,,,,,,,
	// Sheet1
	//
}