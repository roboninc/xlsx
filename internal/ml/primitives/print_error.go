// Copyright (c) 2024 ROBON INC.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package primitives

import (
	"encoding/xml"
)

// PrintErrorType is a type to encode XSD ST_PrintError
type PrintErrorType byte

var (
	ToPrintErrorType   map[string]PrintErrorType
	FromPrintErrorType map[PrintErrorType]string
)

func (t PrintErrorType) String() string {
	return FromPrintErrorType[t]
}

// MarshalXMLAttr marshal PatternType
func (t *PrintErrorType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{Name: name}

	if v, ok := FromPrintErrorType[*t]; ok {
		attr.Value = v
	} else {
		attr = xml.Attr{}
	}

	return attr, nil
}

// UnmarshalXMLAttr unmarshal PatternType
func (t *PrintErrorType) UnmarshalXMLAttr(attr xml.Attr) error {
	if v, ok := ToPrintErrorType[attr.Value]; ok {
		*t = v
	}

	return nil
}
