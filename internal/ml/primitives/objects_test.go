// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package primitives_test

import (
	"encoding/xml"
	"fmt"
	"testing"

	"github.com/roboninc/xlsx/internal/ml/primitives"
	"github.com/stretchr/testify/require"
)

func TestObjects(t *testing.T) {
	type Entity struct {
		Attribute primitives.ObjectsType `xml:"attribute,attr"`
	}

	list := map[primitives.ObjectsType]string{
		primitives.ObjectsType(0):          "",
		primitives.ObjectsTypeAll:          primitives.ObjectsTypeAll.String(),
		primitives.ObjectsTypePlaceholders: primitives.ObjectsTypePlaceholders.String(),
		primitives.ObjectsTypeNone:         primitives.ObjectsTypeNone.String(),
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
