package mapping

// KeywordField keyword, which is used for structured content such as IDs, email
// addresses, hostnames, status codes, zip codes, or tags.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/keyword.html#keyword-field-type
type KeywordField struct {
	BaseField                     `bson:",inline" json:",inline"`
	DocValuesParam                `bson:",inline" json:",inline"`
	EagerGlobalOrdinalsParam      `bson:",inline" json:",inline"`
	FieldsParam                   `bson:",inline" json:",inline"`
	IgnoreAboveParam              `bson:",inline" json:",inline"`
	IndexParam                    `bson:",inline" json:",inline"`
	IndexOptionsParam             `bson:",inline" json:",inline"`
	NormsParam                    `bson:",inline" json:",inline"`
	NullValueParam                `bson:",inline" json:",inline"`
	StoreParam                    `bson:",inline" json:",inline"`
	SimilarityParam               `bson:",inline" json:",inline"`
	NormalizerParam               `bson:",inline" json:",inline"`
	SplitQueriesOnWhitespaceParam `bson:",inline" json:",inline"`
	MetaParam                     `bson:",inline" json:",inline"`
}

func NewKeywordField() *KeywordField {
	return &KeywordField{BaseField: BaseField{MappingType: TypeKeyword}}
}