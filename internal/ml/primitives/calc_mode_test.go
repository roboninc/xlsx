package primitives_test

import (
	"encoding/xml"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/roboninc/xlsx/internal/ml/primitives"
)

func TestCalcMode(t *testing.T) {
	type Entity struct {
		Attribute primitives.CalcMode `xml:"attribute,attr"`
	}

	list := map[primitives.CalcMode]string{
		primitives.CalcMode(0):         "",
		primitives.CalcModeManual:      primitives.CalcModeManual.String(),
		primitives.CalcModeAuto:        primitives.CalcModeAuto.String(),
		primitives.CalcModeAutoNoTable: primitives.CalcModeAutoNoTable.String(),
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
