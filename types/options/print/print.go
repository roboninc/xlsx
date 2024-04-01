// Copyright (c) 2024 ROBON INC.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package options

// Info hold advanced settings of sheet.
// N.B.: You should NOT mutate any value directly.
type Info struct {
	HorizontalCentered bool
	VerticalCentered   bool
	Headings           bool
	GridLines          bool
	GridLinesSet       *bool
}

// Option is helper type to set options for sheet
type Option func(co *Info)

// New create and returns option set for sheet
func New(settings ...Option) *Info {
	i := &Info{}
	i.Set(settings...)
	return i
}

// Set sets new options for option set
func (i *Info) Set(settings ...Option) {
	for _, o := range settings {
		o(i)
	}
}

// HorizontalCentered sets a flag to indicate horizontal center alignment of the page when printed.
func HorizontalCentered(horizontalCentered bool) Option {
	return func(i *Info) {
		i.HorizontalCentered = horizontalCentered
	}
}

// VerticalCentered sets a flag to indicate vertical center alignment of the page when printed.
func VerticalCentered(verticalCentered bool) Option {
	return func(i *Info) {
		i.VerticalCentered = verticalCentered
	}
}

// Headings sets a flag indicating that row and column headings are printed.
func Headings(headings bool) Option {
	return func(i *Info) {
		i.Headings = headings
	}
}

// GridLines sets a flag indicating that grid lines are printed, if both gridLines and gridLinesSet are true.
func GridLines(gridLines bool) Option {
	return func(i *Info) {
		i.GridLines = gridLines
	}
}

// GridLinesSet sets a flag indicating that grid lines are printed, if both gridLines and gridLinesSet are true.
func GridLinesSet(gridLinesSet bool) Option {
	return func(i *Info) {
		if !gridLinesSet {
			i.GridLinesSet = &gridLinesSet
		}
	}
}
