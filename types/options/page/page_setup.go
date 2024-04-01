// Copyright (c) 2024 ROBON INC.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package options

import (
	"github.com/roboninc/xlsx/internal/ml/primitives"
)

// Info hold advanced settings of sheet.
// N.B.: You should NOT mutate any value directly.
type Info struct {
	PaperSize          *int
	Scale              *int
	FirstPageNumber    *int
	FitToWidth         *int
	FitToHeight        *int
	PageOrder          primitives.PageOrderType
	Orientation        primitives.OrientationType
	UsePrinterDefaults *bool
	BlackAndWhite      bool
	Draft              bool
	CellComments       primitives.CellCommentsType
	UseFirstPageNumber bool
	Errors             primitives.PrintErrorType
	HorizontalDpi      *int
	VerticalDpi        *int
	Copies             *int
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

// PaperSize sets a paper size.
func PaperSize(paperSize int) Option {
	return func(i *Info) {
		if paperSize > 1 {
			i.PaperSize = &paperSize
		}
	}
}

// Scale sets a print scaling
func Scale(scale int) Option {
	return func(i *Info) {
		if scale != 100 && 10 <= scale && scale <= 400 {
			i.Scale = &scale
		}
	}
}

// FirstPageNumber sets a page number for first printed page.
func FirstPageNumber(firstPageNumber int) Option {
	return func(i *Info) {
		if firstPageNumber > 1 {
			i.FirstPageNumber = &firstPageNumber
		}
	}
}

// FitToWidth sets a number of horizontal pages to fit on.
func FitToWidth(fitToWidth int) Option {
	return func(i *Info) {
		if fitToWidth > 1 {
			i.FitToWidth = &fitToWidth
		}
	}
}

// FitToHeight sets a number of vertical pages to fit on.
func FitToHeight(fitToHeight int) Option {
	return func(i *Info) {
		if fitToHeight > 1 {
			i.FitToHeight = &fitToHeight
		}
	}
}

// PageOrder sets a order of printed pages.
func PageOrder(pageOrder primitives.PageOrderType) Option {
	return func(i *Info) {
		if pageOrder != PageOrderDownThenOver {
			i.PageOrder = pageOrder
		}
	}
}

// Orientation sets an orientation of the page.
func Orientation(orientation primitives.OrientationType) Option {
	return func(i *Info) {
		if orientation != OrientationDefault {
			i.Orientation = orientation
		}
	}
}

// UsePrinterDefaults sets flag indicating to use the printer’s defaults settings for page setup values
func UsePrinterDefaults(usePrinterDefaults bool) Option {
	return func(i *Info) {
		if !usePrinterDefaults {
			i.UsePrinterDefaults = &usePrinterDefaults
		}
	}
}

// BlackAndWhite sets flag indicating to print black and white.
func BlackAndWhite(blackAndWhite bool) Option {
	return func(i *Info) {
		i.BlackAndWhite = blackAndWhite
	}
}

// Draft sets flag indicating to print without graphics.
func Draft(draft bool) Option {
	return func(i *Info) {
		i.Draft = draft
	}
}

// CellComments specifies how to print cell comments.
func CellComments(cellComments primitives.CellCommentsType) Option {
	return func(i *Info) {
		if cellComments != CellCommentsNone {
			i.CellComments = cellComments
		}
	}
}

// UseFirstPageNumber sets flag indicating to use firstPageNumber value for first page number,
func UseFirstPageNumber(useFirstPageNumber bool) Option {
	return func(i *Info) {
		i.UseFirstPageNumber = useFirstPageNumber
	}
}

// Errors specifies how to print cell values for cells with errors.
func Errors(errors primitives.PrintErrorType) Option {
	return func(i *Info) {
		if errors != PrintErrorDisplayed {
			i.Errors = errors
		}
	}
}

// HorizontalDpi sets a horizontal print resolution of the device.
func HorizontalDpi(horizontalDpi int) Option {
	return func(i *Info) {
		if horizontalDpi > 0 && horizontalDpi != 600 {
			i.HorizontalDpi = &horizontalDpi
		}
	}
}

// VerticalDpi sets a vertical print resolution of the device.
func VerticalDpi(verticalDpi int) Option {
	return func(i *Info) {
		if verticalDpi > 0 && verticalDpi != 600 {
			i.VerticalDpi = &verticalDpi
		}
	}
}

// Copies sets a number of copies to print.
func Copies(copies int) Option {
	return func(i *Info) {
		if copies > 1 {
			i.Copies = &copies
		}
	}
}
