// Copyright (c) 2024 ROBON INC.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package primitives

import (
	"encoding/xml"
)

// OrientationType is a type to encode XSD ST_Orientation
type OrientationType byte

var (
	ToOrientationType   map[string]OrientationType
	FromOrientationType map[OrientationType]string
)

func (t OrientationType) String() string {
	return FromOrientationType[t]
}

// MarshalXMLAttr marshal PatternType
func (t *OrientationType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{Name: name}

	if v, ok := FromOrientationType[*t]; ok {
		attr.Value = v
	} else {
		attr = xml.Attr{}
	}

	return attr, nil
}

// UnmarshalXMLAttr unmarshal PatternType
func (t *OrientationType) UnmarshalXMLAttr(attr xml.Attr) error {
	if v, ok := ToOrientationType[attr.Value]; ok {
		*t = v
	}

	return nil
}
