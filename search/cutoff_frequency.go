package search

import (
	"encoding/json"
	"reflect"

	"github.com/chanced/dynamic"
)

type WithCutoffFrequency interface {
	CutoffFrequency() dynamic.Number
	SetCutoffFreuency(dynamic.Number)
}

type cutoffFrequencyParam struct {
	cutoffFrequency dynamic.Number
}

func (cf cutoffFrequencyParam) CutoffFrequency() dynamic.Number {
	return cf.cutoffFrequency
}
func (cf *cutoffFrequencyParam) SetCutoffFreuency(n dynamic.Number) {
	cf.cutoffFrequency = n
}

func unmarshalCutoffFrequencyParam(data dynamic.JSON, target interface{}) error {
	if a, ok := target.(WithCutoffFrequency); ok {
		var str string
		var err error
		switch {
		case data.IsNull():
		case data.IsString():
			err = json.Unmarshal(data, &str)
		case data.IsNumber():
			str = data.UnquotedString()
		default:
			err = &json.UnmarshalTypeError{
				Value: data.String(),
				Type:  reflect.TypeOf(dynamic.Number{}),
			}
		}
		if err != nil {
			return err
		}

		n, err := dynamic.NewNumber(str)
		if err != nil {
			return err
		}
		a.SetCutoffFreuency(n)
	}
	return nil
}
func marshalCutoffFrequencyParam(data dynamic.Map, source interface{}) (dynamic.Map, error) {
	if p, ok := source.(WithCutoffFrequency); ok {
		if p.CutoffFrequency().HasValue() {
			data[paramCutoffFrequency] = p.CutoffFrequency()
		}
	}
	return data, nil
}
