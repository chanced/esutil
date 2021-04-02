package picker

import "encoding/json"

type RangeField interface {
	Field
	WithCoerce
	WithIndex
	WithStore
}

type rangeFieldParams struct {
	// Coercion attempts to clean up dirty values to fit the data type of a
	// field. (Optional, bool) Defaults to false.
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/coerce.html
	Coerce interface{} `json:"coerce,omitempty"`
	// Index controls whether field values are indexed. It accepts true or false
	// and defaults to true. Fields that are not indexed are not queryable.
	// (Optional, bool)
	Index interface{} `json:"index,omitempty"`
	// WithStore is a mapping with a store paramter.
	//
	// By default, field values are indexed to make them searchable, but they are
	// not stored. This means that the field can be queried, but the original field
	// value cannot be retrieved.
	//
	// Usually this doesn’t matter. The field value is already part of the _source
	// field, which is stored by default. If you only want to retrieve the value of
	// a single field or of a few fields, instead of the whole _source, then this
	// can be achieved with source filtering.
	//
	// In certain situations it can make sense to store a field. For instance, if
	// you have a document with a title, a date, and a very large content field, you
	// may want to retrieve just the title and the date without having to extract
	// those fields from a large _source field
	//
	// Stored fields returned as arrays
	//
	// For consistency, stored fields are always returned as an array because there
	// is no way of knowing if the original field value was a single value, multiple
	// values, or an empty array.
	//
	// The original value can be retrieved from the _source field instead.
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/mapping-store.html
	Store interface{} `json:"store,omitempty"`
}

// A RangeField type represent a continuous range of values between an upper and
// lower bound. For example, a range can represent any date in October or any
// integer from 0 to 9. They are defined using the operators gt or gte for the
// lower bound, and lt or lte for the upper bound. They can be used for
// querying, and have limited support for aggregations. The only supported
// aggregations are histogram, cardinality.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/range.html

type IntegerRangeFieldParams rangeFieldParams

func (p IntegerRangeFieldParams) Field() (Field, error) {
	return p.IntegerRange()
}
func (p IntegerRangeFieldParams) IntegerRange() (*IntegerRangeField, error) {
	f := &IntegerRangeField{}
	err := f.SetCoerce(p)
	if err != nil {
		return f, err
	}
	err = f.SetIndex(p.Index)
	if err != nil {
		return f, err
	}
	err = f.SetStore(p.Store)
	if err != nil {
		return f, err
	}
	return f, nil
}

func NewIntegerRangeField(params IntegerRangeFieldParams) (*IntegerRangeField, error) {
	return params.IntegerRange()
}

type IntegerRangeField struct {
	coerceParam
	indexParam
	storeParam
}

func (IntegerRangeField) Type() FieldType {
	return FieldTypeIntegerRange
}

func (r IntegerRangeField) MarshalJSON() ([]byte, error) {
	return json.Marshal(IntegerRangeFieldParams{
		Coerce: r.coerce.Value(),
		Index:  r.index.Value(),
		Store:  r.store.Value(),
	})
}

func (r *IntegerRangeField) UnmarshalJSON(data []byte) error {
	*r = IntegerRangeField{}
	p := IntegerRangeFieldParams{}
	err := json.Unmarshal(data, &p)
	if err != nil {
		return err
	}
	v, err := p.IntegerRange()
	if err != nil {
		return err
	}
	*r = *v
	return nil
}

type FloatRangeFieldParams rangeFieldParams

func (p FloatRangeFieldParams) Field() (Field, error) {
	return p.FloatRange()
}
func (p FloatRangeFieldParams) FloatRange() (*FloatRangeField, error) {
	f := &FloatRangeField{}
	err := f.SetCoerce(p)
	if err != nil {
		return f, err
	}
	err = f.SetIndex(p.Index)
	if err != nil {
		return f, err
	}
	err = f.SetStore(p.Store)
	if err != nil {
		return f, err
	}
	return f, nil
}

func NewFloatRangeField(params FloatRangeFieldParams) (*FloatRangeField, error) {
	return params.FloatRange()
}

// FloatRangeField is a range of single-precision 32-bit IEEE 754 floating point
// values.
type FloatRangeField struct {
	coerceParam
	indexParam
	storeParam
}

func (FloatRangeField) Type() FieldType {
	return FieldTypeFloatRange
}

func (r FloatRangeField) MarshalJSON() ([]byte, error) {
	return json.Marshal(FloatRangeFieldParams{
		Coerce: r.coerce.Value(),
		Index:  r.index.Value(),
		Store:  r.store.Value(),
	})
}

func (r *FloatRangeField) UnmarshalJSON(data []byte) error {
	*r = FloatRangeField{}
	p := FloatRangeFieldParams{}
	err := json.Unmarshal(data, &p)
	if err != nil {
		return err
	}
	v, err := p.FloatRange()
	if err != nil {
		return err
	}
	*r = *v
	return nil
}

type LongRangeFieldParams rangeFieldParams

func (p LongRangeFieldParams) Field() (Field, error) {
	return p.LongRange()
}
func (p LongRangeFieldParams) LongRange() (*LongRangeField, error) {
	f := &LongRangeField{}
	err := f.SetCoerce(p)
	if err != nil {
		return f, err
	}
	err = f.SetIndex(p.Index)
	if err != nil {
		return f, err
	}
	err = f.SetStore(p.Store)
	if err != nil {
		return f, err
	}
	return f, nil
}

func NewLongRangeField(params LongRangeFieldParams) (*LongRangeField, error) {
	return params.LongRange()
}

// LongRangeField is a range of signed 64-bit integers with a minimum value of
// -263 and maximum of 263-1.
type LongRangeField struct {
	coerceParam
	indexParam
	storeParam
}

func (LongRangeField) Type() FieldType {
	return FieldTypeLongRange
}

func (r LongRangeField) MarshalJSON() ([]byte, error) {
	return json.Marshal(LongRangeFieldParams{
		Coerce: r.coerce.Value(),
		Index:  r.index.Value(),
		Store:  r.store.Value(),
	})
}

func (r *LongRangeField) UnmarshalJSON(data []byte) error {
	*r = LongRangeField{}
	p := LongRangeFieldParams{}
	err := json.Unmarshal(data, &p)
	if err != nil {
		return err
	}
	v, err := p.LongRange()
	if err != nil {
		return err
	}
	*r = *v
	return nil
}

type DoubleRangeFieldParams rangeFieldParams

func (p DoubleRangeFieldParams) Field() (Field, error) {
	return p.DoubleRange()
}
func (p DoubleRangeFieldParams) DoubleRange() (*DoubleRangeField, error) {
	f := &DoubleRangeField{}
	err := f.SetCoerce(p)
	if err != nil {
		return f, err
	}
	err = f.SetIndex(p.Index)
	if err != nil {
		return f, err
	}
	err = f.SetStore(p.Store)
	if err != nil {
		return f, err
	}
	return f, nil
}

func NewDoubleRangeField(params DoubleRangeFieldParams) (*DoubleRangeField, error) {
	return params.DoubleRange()
}

// DoubleRangeField is a range of double-precision 64-bit IEEE 754 floating
// point values.
type DoubleRangeField struct {
	coerceParam
	indexParam
	storeParam
}

func (DoubleRangeField) Type() FieldType {
	return FieldTypeDoubleRange
}

func (r DoubleRangeField) MarshalJSON() ([]byte, error) {
	return json.Marshal(DoubleRangeFieldParams{
		Coerce: r.coerce.Value(),
		Index:  r.index.Value(),
		Store:  r.store.Value(),
	})
}

func (r *DoubleRangeField) UnmarshalJSON(data []byte) error {
	*r = DoubleRangeField{}
	p := DoubleRangeFieldParams{}
	err := json.Unmarshal(data, &p)
	if err != nil {
		return err
	}
	v, err := p.DoubleRange()
	if err != nil {
		return err
	}
	*r = *v
	return nil
}

type DateRangeFieldParams rangeFieldParams

func (p DateRangeFieldParams) Field() (Field, error) {
	return p.DateRange()
}
func (p DateRangeFieldParams) DateRange() (*DateRangeField, error) {
	f := &DateRangeField{}
	err := f.SetCoerce(p)
	if err != nil {
		return f, err
	}
	err = f.SetIndex(p.Index)
	if err != nil {
		return f, err
	}
	err = f.SetStore(p.Store)
	if err != nil {
		return f, err
	}
	return f, nil
}

func NewDateRangeField(params DateRangeFieldParams) (*DateRangeField, error) {
	return params.DateRange()
}

// DateRangeField is a range of date values. Date ranges support various date
// formats through the format mapping parameter. Regardless of the format used,
// date values are parsed into an unsigned 64-bit integer representing
// milliseconds since the Unix epoch in UTC. Values containing the now date math
// expression are not supported.
type DateRangeField struct {
	coerceParam
	indexParam
	storeParam
}

func (DateRangeField) Type() FieldType {
	return FieldTypeDateRange
}

func (r DateRangeField) MarshalJSON() ([]byte, error) {
	return json.Marshal(DateRangeFieldParams{
		Coerce: r.coerce.Value(),
		Index:  r.index.Value(),
		Store:  r.store.Value(),
	})
}

func (r *DateRangeField) UnmarshalJSON(data []byte) error {
	*r = DateRangeField{}
	p := DateRangeFieldParams{}
	err := json.Unmarshal(data, &p)
	if err != nil {
		return err
	}
	v, err := p.DateRange()
	if err != nil {
		return err
	}
	*r = *v
	return nil
}

type IPRangeFieldParams rangeFieldParams

func (p IPRangeFieldParams) Field() (Field, error) {
	return p.IPRange()
}
func (p IPRangeFieldParams) IPRange() (*IPRangeField, error) {
	f := &IPRangeField{}
	err := f.SetCoerce(p)
	if err != nil {
		return f, err
	}
	err = f.SetIndex(p.Index)
	if err != nil {
		return f, err
	}
	err = f.SetStore(p.Store)
	if err != nil {
		return f, err
	}
	return f, nil
}

func NewIPRangeField(params IPRangeFieldParams) (*IPRangeField, error) {
	return params.IPRange()
}

// IPRangeField is a range of ip values supporting either IPv4 or IPv6 (or
// mixed) addresses.
type IPRangeField struct {
	coerceParam
	indexParam
	storeParam
}

func (IPRangeField) Type() FieldType {
	return FieldTypeIPRange
}

func (r IPRangeField) MarshalJSON() ([]byte, error) {
	return json.Marshal(IPRangeFieldParams{
		Coerce: r.coerce.Value(),
		Index:  r.index.Value(),
		Store:  r.store.Value(),
	})
}

func (r *IPRangeField) UnmarshalJSON(data []byte) error {
	*r = IPRangeField{}
	p := IPRangeFieldParams{}
	err := json.Unmarshal(data, &p)
	if err != nil {
		return err
	}
	v, err := p.IPRange()
	if err != nil {
		return err
	}
	*r = *v
	return nil
}