package picker_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/chanced/picker"
	"github.com/stretchr/testify/require"
)

func TestConstantScore(t *testing.T) {
	assert := require.New(t)
	data := []byte(`{
		"query": {
			"constant_score": {
				"filter": {
					"term": { "user.id": { "value": "kimchy" } }
				},
				"boost": 1.2
			}
		}
	}`)
	s, err := picker.NewSearch(picker.SearchParams{
		Query: picker.QueryParams{
			ConstantScore: picker.ConstantScoreQueryParams{
				Filter: &picker.QueryParams{
					Term: picker.TermQueryParams{
						Field: "user.id",
						Value: "kimchy",
					},
				},
				Boost: 1.2,
			},
		},
	})

	assert.NoError(err)
	sd, err := json.MarshalIndent(s, "", "  ")
	fmt.Println(string(sd))
	assert.NoError(err)
	assert.NoError(compareJSONObject(data, sd))

	var sr *picker.Search
	err = json.Unmarshal(data, &sr)

	assert.NoError(err)
}
