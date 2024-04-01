// Copyright (c) 2024 ROBON INC.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package primitives

import (
	"encoding/xml"
)

// CellCommentsType is a type to encode XSD ST_CellComments
type CellCommentsType byte

var (
	ToCellCommentsType   map[string]CellCommentsType
	FromCellCommentsType map[CellCommentsType]string
)

func (t CellCommentsType) String() string {
	return FromCellCommentsType[t]
}

// MarshalXMLAttr marshal PatternType
func (t *CellCommentsType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{Name: name}

	if v, ok := FromCellCommentsType[*t]; ok {
		attr.Value = v
	} else {
		attr = xml.Attr{}
	}

	return attr, nil
}

// UnmarshalXMLAttr unmarshal PatternType
func (t *CellCommentsType) UnmarshalXMLAttr(attr xml.Attr) error {
	if v, ok := ToCellCommentsType[attr.Value]; ok {
		*t = v
	}

	return nil
}
