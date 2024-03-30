// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package xlsx

import (
	"github.com/roboninc/ooxml"

	//init enums for marshal/unmarshal
	_ "github.com/roboninc/xlsx/format/conditional"
	_ "github.com/roboninc/xlsx/format/conditional/rule"
	_ "github.com/roboninc/xlsx/format/styles"
	_ "github.com/roboninc/xlsx/internal/ml"
	_ "github.com/roboninc/xlsx/internal/ml/primitives"
	_ "github.com/roboninc/xlsx/types"
	_ "github.com/roboninc/xlsx/types/comment"
	_ "github.com/roboninc/xlsx/types/hyperlink"
	_ "github.com/roboninc/xlsx/types/options/column"
	_ "github.com/roboninc/xlsx/types/options/row"
	_ "github.com/roboninc/xlsx/types/options/sheet"
)

// 真偽値を変数で用意する。
// Open Office XML の仕様で、デフォルトが true の属性の省略時値を実現するため、bool のポインターを使用する。
// リテラル値のポインタは使用できないため、bool値のポインターを作るために共通で利用できる変数とする。
var (
	TRUE  = true
	FALSE = false
)

// Open opens a XLSX file with name or io.Reader
func Open(f interface{}) (*Spreadsheet, error) {
	doc, err := ooxml.Open(f, newSpreadsheet)
	if err != nil {
		return nil, err
	}

	if xlDoc, ok := doc.(*Spreadsheet); ok {
		return xlDoc, nil
	}

	return nil, ooxml.ErrorUnknownPackage(Spreadsheet{})
}

// New creates and returns a new XLSX document
func New() *Spreadsheet {
	if doc, err := newSpreadsheet(ooxml.NewPackage(nil)); err == nil {
		if xlDoc, ok := doc.(*Spreadsheet); ok {
			return xlDoc
		}
	}

	panic("Could not create a new XLSX document.")
}
