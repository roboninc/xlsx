// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package rule

import (
	"github.com/roboninc/xlsx/format/styles"
	"github.com/roboninc/xlsx/internal/ml"
	"github.com/roboninc/xlsx/internal/ml/primitives"
)

type uniqueRule struct {
	baseRule
}

// Unique is helper object to set specific options for rule
var Unique uniqueRule

func (x uniqueRule) initIfRequired(r *Info) {
	if !r.initialized {
		r.initialized = true
		r.validator = Unique
		r.rule = &ml.ConditionalRule{
			Type: primitives.ConditionTypeUniqueValues,
		}
	}
}

func (x uniqueRule) Styles(s *styles.Info) Option {
	return func(r *Info) {
		x.initIfRequired(r)
		r.style = s
	}
}
