// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package xlsx

import (
	"github.com/roboninc/ooxml"
	"github.com/roboninc/xlsx/internal"
	"github.com/roboninc/xlsx/internal/ml"
	"github.com/roboninc/xlsx/internal/ml/primitives"
)

// workbook is a higher level object that wraps ml.Workbook with functionality
type workbook struct {
	ml   ml.Workbook
	doc  *Spreadsheet
	file *ooxml.PackageFile
}

func newWorkbook(f interface{}, doc *Spreadsheet) *workbook {
	wb := &workbook{
		doc: doc,
	}

	doc.workbook = wb

	wb.file = ooxml.NewPackageFile(doc.pkg, f, &wb.ml, nil)
	wb.file.LoadIfRequired(nil)

	if wb.file.IsNew() {
		doc.pkg.ContentTypes().RegisterContent(wb.file.FileName(), internal.ContentTypeWorkbook)
		doc.pkg.Relationships().AddFile(internal.RelationTypeWorkbook, wb.file.FileName())
		wb.file.MarkAsUpdated()
	}

	return wb
}

// setFullCalcOnLoad は、CalcPr.FullCalcOnLoad の値を設定します。
func (wb *workbook) setFullCalcOnLoad(v bool) {
	pr := wb.ml.CalcPr
	if pr != nil && pr.FullCalcOnLoad != nil && *pr.FullCalcOnLoad == v {
		return
	}
	if pr == nil {
		wb.ml.CalcPr = &ml.CalcPr{}
	}
	if v {
		wb.ml.CalcPr.CalcMode = primitives.CalcModeAuto
	}
	wb.ml.CalcPr.FullCalcOnLoad = &v
	wb.file.MarkAsUpdated()
}
