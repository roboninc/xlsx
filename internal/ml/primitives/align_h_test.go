// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package primitives_test

import (
	"encoding/xml"
	"fmt"
	"testing"

	"github.com/roboninc/xlsx/format/styles"
	"github.com/roboninc/xlsx/internal/ml/primitives"
	"github.com/stretchr/testify/require"
)

func TestAlignH(t *testing.T) {
	type Entity struct {
		Attribute primitives.HAlignType `xml:"attribute,attr"`
	}

	list := map[primitives.HAlignType]string{
		primitives.HAlignType(0):      "",
		styles.HAlignGeneral:          styles.HAlignGeneral.String(),
		styles.HAlignLeft:             styles.HAlignLeft.String(),
		styles.HAlignCenter:           styles.HAlignCenter.String(),
		styles.HAlignRight:            styles.HAlignRight.String(),
		styles.HAlignFill:             styles.HAlignFill.String(),
		styles.HAlignJustify:          styles.HAlignJustify.String(),
		styles.HAlignCenterContinuous: styles.HAlignCenterContinuous.String(),
		styles.HAlignDistributed:      styles.HAlignDistributed.String(),
	}

	for v, s := range list {
		t.Run(s, func(tt *testing.T) {
			entity := Entity{Attribute: v}
			encoded, err := xml.Marshal(&entity)

			require.Empty(tt, err)
			if v == 0 {
				require.Equal(tt, `<Entity></Entity>`, string(encoded))
			} else {
				require.Equal(tt, fmt.Sprintf(`<Entity attribute="%s"></Entity>`, s), string(encoded))
			}

			var decoded Entity
			err = xml.Unmarshal(encoded, &decoded)
			require.Empty(tt, err)

			require.Equal(tt, entity, decoded)
			require.Equal(tt, s, decoded.Attribute.String())
		})
	}
}
