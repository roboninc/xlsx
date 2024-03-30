package primitives_test

import (
	"encoding/xml"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/roboninc/xlsx/internal/ml/primitives"
)

func TestRefMode(t *testing.T) {
	type Entity struct {
		Attribute primitives.RefMode `xml:"attribute,attr"`
	}

	list := map[primitives.RefMode]string{
		primitives.RefMode(0):  "",
		primitives.RefModeA1:   primitives.RefModeA1.String(),
		primitives.RefModeR1C1: primitives.RefModeR1C1.String(),
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
