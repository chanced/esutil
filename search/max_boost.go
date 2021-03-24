package search

import (
	"math"

	"github.com/tidwall/gjson"
)

type WithMaxBoost interface {
	MaxBoost() float64
	SetMaxBoost(v float64)
}

type MaxBoostParam struct {
	MaxBoostValue *float64 `json:"max_boost,omitempty" bson:"max_boost,omitempty"`
}

func (mb MaxBoostParam) MaxBoost() float64 {
	if mb.MaxBoostValue == nil {
		return math.MaxFloat32
	}
	return *mb.MaxBoostValue
}

func (mb MaxBoostParam) SetMaxBoost(v float64) {
	if mb.MaxBoost() != v {
		mb.MaxBoostValue = &v
	}
}
func unmarshalMaxBoostParam(value gjson.Result, target interface{}) error {
	if a, ok := target.(WithMaxBoost); ok {
		a.SetMaxBoost(value.Float())
	}
	return nil
}