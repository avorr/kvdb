package compute

import (
	"errors"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		queryStr string

		expectedQuery Query
		expectedErr   error
	}{
		"empty query": {
			queryStr:    "",
			expectedErr: errors.New("invalid query"),
		},
		"empty query without tokens": {
			queryStr:    "   ",
			expectedErr: errors.New("invalid query"),
		},
		"query with UTF symbols": {
			queryStr:    "字文下",
			expectedErr: errors.New("invalid command"),
		},
		"invalid command": {
			queryStr:    "TRUNCATE",
			expectedErr: errors.New("invalid command"),
		},
		"invalid number arguments for set query": {
			queryStr:    "SET key",
			expectedErr: errors.New("invalid arguments"),
		},
		"invalid number arguments for get query": {
			queryStr:    "GET key value",
			expectedErr: errors.New("invalid arguments"),
		},
		"invalid number arguments for del query": {
			queryStr:    "GET key value",
			expectedErr: errors.New("invalid arguments"),
		},
		"set query": {
			queryStr:      "SET __key__\nvalue",
			expectedQuery: NewQuery("SET", []string{"__key__", "value"}),
		},
		"get query": {
			queryStr:      "GET\t1key2",
			expectedQuery: NewQuery("GET", []string{"1key2"}),
		},
		"del query": {
			queryStr:      "DEL  /key-",
			expectedQuery: NewQuery("DEL", []string{"/key-"}),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			query, err := Parser(test.queryStr)
			assert.Equal(t, test.expectedErr, err)
			assert.True(t, reflect.DeepEqual(test.expectedQuery, query))
		})
	}
}
