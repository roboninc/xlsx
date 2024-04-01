// Copyright (c) 2024 ROBON INC.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package primitives

import (
	"encoding/xml"
)

// PageOrderType is a type to encode XSD ST_PageOrder
type PageOrderType byte

var (
	ToPageOrderType   map[string]PageOrderType
	FromPageOrderType map[PageOrderType]string
)

func (t PageOrderType) String() string {
	return FromPageOrderType[t]
}

// MarshalXMLAttr marshal PatternType
func (t *PageOrderType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{Name: name}

	if v, ok := FromPageOrderType[*t]; ok {
		attr.Value = v
	} else {
		attr = xml.Attr{}
	}

	return attr, nil
}

// UnmarshalXMLAttr unmarshal PatternType
func (t *PageOrderType) UnmarshalXMLAttr(attr xml.Attr) error {
	if v, ok := ToPageOrderType[attr.Value]; ok {
		*t = v
	}

	return nil
}
