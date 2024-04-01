// Copyright (c) 2024 ROBON INC.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package options

import (
	"github.com/roboninc/xlsx/internal/ml/primitives"
)

// List of all possible values for OrientationType
const (
	_ primitives.OrientationType = iota
	OrientationDefault
	OrientationPortrait
	OrientationLandscape
)

func init() {
	primitives.FromOrientationType = map[primitives.OrientationType]string{
		OrientationDefault:   "default",
		OrientationPortrait:  "portrait",
		OrientationLandscape: "landscape",
	}

	primitives.ToOrientationType = make(map[string]primitives.OrientationType, len(primitives.FromOrientationType))
	for k, v := range primitives.FromOrientationType {
		primitives.ToOrientationType[v] = k
	}
}
