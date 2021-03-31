package picker

import (
	"encoding/json"

	"github.com/chanced/dynamic"
)

type Fuzzier interface {
	Fuzzy() (*FuzzyClause, error)
}

// FuzzyQuery returns documents that contain terms similar to the search term,
// as measured by a Levenshtein edit distance.
//
// An edit distance is the number of one-character changes needed to turn one
// term into another. These changes can include:
//
//      - Changing a character (box → fox)
//
//      - Removing a character (black → lack)
//
//      - Inserting a character (sic → sick)
//
//      - Transposing two adjacent characters (act → cat)
//
// To find similar terms, the fuzzy query creates a set of all possible
// variations, or expansions, of the search term within a specified edit
// distance. The query then returns exact matches for each expansion.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-fuzzy-query.html
type FuzzyQuery struct {
	// Value or term to find in the provided <field>. (Required)
	Value string
	// Field which is being queried against.
	//
	// This will be ignored if set through a mutator
	Field string
	// Maximum edit distance allowed for matching. See Fuzziness for valid
	// values and more information. (Optional)
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/common-options.html#fuzziness
	Fuzziness string
	// PrefixLength is the number of beginning characters left unchanged when
	// creating expansions. Defaults to 0. (Optional)
	PrefixLength int
	// NoTranspositions indicates whether edits include transpositions of two
	// adjacent characters (ab → ba). (Optional)
	//
	// Setting NoTranspositions to true sets transpositions to false
	NoTranspositions bool
	// Rewrite method used to rewrite the query. (Optional)
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-multi-term-rewrite.html
	Rewrite Rewrite
	// Maximum number of variations created. Defaults to 50. (Optional)
	//
	// Avoid using a high value in the max_expansions parameter, especially if
	// the prefix_length parameter value is 0. High values in the max_expansions
	// parameter can cause poor performance due to the high number of variations
	// examined.
	MaxExpansions interface{}
	// Name of the query (Optional)
	Name string
}

func (f FuzzyQuery) Fuzzy() (*FuzzyClause, error) {
	q := &FuzzyClause{field: f.Field}
	err := q.setValue(f.Value)
	if err != nil {
		return q, NewQueryError(err, KindFuzzy, f.Field)
	}
	err = q.SetMaxExpansions(f.MaxExpansions)
	if err != nil {
		return q, NewQueryError(err, KindFuzzy, f.Field)
	}
	err = q.SetRewrite(f.Rewrite)
	if err != nil {
		return q, NewQueryError(err, KindFuzzy, f.Field)
	}
	q.SetTranspositions(!f.NoTranspositions)
	q.SetName(f.Name)
	q.SetFuzziness(f.Fuzziness)
	q.SetPrefixLength(f.PrefixLength)
	return q, nil
}

func (f FuzzyQuery) Clause() (Clause, error) {
	return f.Fuzzy()
}

// func NewFuzzyQuery(params Fuzzier) (*FuzzyQuery, error) {
// 	q, err := params.Fuzzy()
// 	if err != nil {
// 		return nil, NewQueryError(err, KindFuzzy, getField(q, nil))
// 	}
// 	err = checkField(q.field, KindFuzzy)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return q, nil
// }

// FuzzyClause returns documents that contain terms similar to the search term,
// as measured by a Levenshtein edit distance.
//
// An edit distance is the number of one-character changes needed to turn one
// term into another. These changes can include:
//
//      - Changing a character (box → fox)
//
//      - Removing a character (black → lack)
//
//      - Inserting a character (sic → sick)
//
//      - Transposing two adjacent characters (act → cat)
//
// To find similar terms, the fuzzy query creates a set of all possible
// variations, or expansions, of the search term within a specified edit
// distance. The query then returns exact matches for each expansion.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-fuzzy-query.html
type FuzzyClause struct {
	field string
	value string
	fuzzinessParam
	maxExpansionsParam
	prefixLengthParam
	transpositionsParam
	rewriteParam
	nameParam
	completeClause
}

func (f *FuzzyClause) Clause() (QueryClause, error) {
	return f, nil
}
func (f *FuzzyClause) Vale() string {
	return f.value
}

func (f FuzzyClause) Field() string {
	return f.field
}
func (f FuzzyClause) Kind() Kind {
	return KindFuzzy
}

func (f *FuzzyClause) IsEmpty() bool {
	return !(len(f.field) != 0 || len(f.value) != 0)
}

func (f *FuzzyClause) Set(field string, fuzzier Fuzzier) error {
	q, err := fuzzier.Fuzzy()
	if err != nil {
		return NewQueryError(err, KindFuzzy, field)
	}
	err = checkField(field, KindFuzzy)
	if err != nil {
		return err
	}
	q.field = field
	*f = *q
	return nil
}
func (f *FuzzyClause) setValue(v string) error {
	if len(v) == 0 {
		return ErrValueRequired
	}
	return nil
}

func (f FuzzyClause) MarshalJSON() ([]byte, error) {
	if f.IsEmpty() {
		return dynamic.Null, nil
	}
	data, err := f.marshalClauseJSON()
	if err != nil {
		return nil, err
	}
	return json.Marshal(dynamic.Map{f.field: data})
}

func (f FuzzyClause) marshalClauseJSON() (dynamic.JSON, error) {
	params, err := marshalClauseParams(&f)
	if err != nil {
		return nil, err
	}
	params["value"] = f.value
	return json.Marshal(params)
}

func (f *FuzzyClause) UnmarshalJSON(data []byte) error {
	*f = FuzzyClause{}

	d := map[string]dynamic.JSON{}
	err := json.Unmarshal(data, &d)
	if err != nil {
		return err
	}
	for k, v := range d {
		f.field = k
		return f.unmarshalClauseJSON(v)
	}
	return nil
}

func (f *FuzzyClause) unmarshalClauseJSON(data dynamic.JSON) error {
	fields, err := unmarshalClauseParams(data, f)
	if err != nil {
		return err
	}
	if v, ok := fields["value"]; ok {
		var value string
		err := json.Unmarshal(v, &value)
		if err != nil {
			return err
		}
		f.value = value
	}
	return nil
}

func (f *FuzzyClause) Clear() {
	*f = FuzzyClause{}
}