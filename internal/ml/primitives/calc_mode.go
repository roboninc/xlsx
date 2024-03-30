package primitives

import (
	"encoding/xml"
)

// CalcMode is Calculation Mode
type CalcMode byte

var (
	toCalcMode   map[string]CalcMode
	fromCalcMode map[CalcMode]string
)

//List of all possible values for CalcMode
const (
	_ CalcMode = iota
	CalcModeManual
	CalcModeAuto
	CalcModeAutoNoTable
)

func init() {
	fromCalcMode = map[CalcMode]string{
		CalcModeManual:      "manual",
		CalcModeAuto:        "auto",
		CalcModeAutoNoTable: "autoNoTable",
	}
	toCalcMode = make(map[string]CalcMode, len(fromCalcMode))
	for k, v := range fromCalcMode {
		toCalcMode[v] = k
	}
}

func (e CalcMode) String() string {
	return fromCalcMode[e]
}

//MarshalXMLAttr marshal CalcMode
func (e *CalcMode) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{Name: name}

	if v, ok := fromCalcMode[*e]; ok {
		attr.Value = v
	} else {
		attr = xml.Attr{}
	}

	return attr, nil
}

//UnmarshalXMLAttr unmarshal CalcMode
func (e *CalcMode) UnmarshalXMLAttr(attr xml.Attr) error {
	if v, ok := toCalcMode[attr.Value]; ok {
		*e = v
	}

	return nil
}
