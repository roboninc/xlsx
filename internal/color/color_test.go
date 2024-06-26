// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package color_test

import (
	"testing"

	"github.com/roboninc/xlsx/internal/color"
	"github.com/roboninc/xlsx/internal/ml"
	"github.com/roboninc/xlsx/internal/ml/primitives"
	"github.com/stretchr/testify/require"
)

func TestColor(t *testing.T) {
	require.Equal(t, &ml.Color{Indexed: primitives.OptionalIndex(6)}, color.New("#FF00FF"))
	require.Equal(t, &ml.Color{RGB: "FF112233"}, color.New("#112233"))
}
