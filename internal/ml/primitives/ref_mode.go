package primitives

import (
	"encoding/xml"
)

// RefMode is Reference Mode
type RefMode byte

var (
	toRefMode   map[string]RefMode
	fromRefMode map[RefMode]string
)

//List of all possible values for RefMode
const (
	_ RefMode = iota
	RefModeA1
	RefModeR1C1
)

func init() {
	fromRefMode = map[RefMode]string{
		RefModeA1:   "A1",
		RefModeR1C1: "R1C1",
	}
	toRefMode = make(map[string]RefMode, len(fromRefMode))
	for k, v := range fromRefMode {
		toRefMode[v] = k
	}
}

func (e RefMode) String() string {
	return fromRefMode[e]
}

//MarshalXMLAttr marshal CalcMode
func (e *RefMode) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{Name: name}

	if v, ok := fromRefMode[*e]; ok {
		attr.Value = v
	} else {
		attr = xml.Attr{}
	}

	return attr, nil
}

//UnmarshalXMLAttr unmarshal CalcMode
func (e *RefMode) UnmarshalXMLAttr(attr xml.Attr) error {
	if v, ok := toRefMode[attr.Value]; ok {
		*e = v
	}

	return nil
}
