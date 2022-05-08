package utils_test

import (
	"election-service/internal/utils"
	"testing"

	"github.com/stretchr/testify/require"
)

type testcase struct {
	name     string
	request  interface{}
	expected map[string]interface{}
	err      string
}

func TestService_Transcode(t *testing.T) {
	assertions := require.New(t)

	t.Run("Transcode", func(t *testing.T) {
		cases := []testcase{
			{name: "JSON Success", request: map[string]interface{}{"test": "test"}, expected: map[string]interface{}{"test": "test"}, err: ""},
			{name: "NewDecoder error", request: "test", expected: nil, err: "json"},
		}

		for _, c := range cases {
			t.Run(c.name, func(t *testing.T) {
				var actual map[string]interface{}
				err := utils.Transcode(c.request, &actual)
				if len(c.err) > 0 {
					assertions.ErrorContains(err, c.err)
				}
				assertions.Equal(c.expected, actual)
			})
		}
	})
}

func TestService_JsonFilter(t *testing.T) {
	assertions := require.New(t)

	t.Run("JsonFilter", func(t *testing.T) {
		cases := []testcase{
			{name: "Success", request: map[string]interface{}{"test": "test"}, expected: map[string]interface{}{"test": "test"}, err: ""},
			{name: "NewDecoder error", request: "test", expected: nil, err: "json"},
		}

		for _, c := range cases {
			t.Run(c.name, func(t *testing.T) {
				var actual map[string]interface{}
				err := utils.JsonFilter(c.request, &actual)
				if len(c.err) > 0 {
					assertions.ErrorContains(err, c.err)
				}
				assertions.Equal(c.expected, actual)
			})
		}
	})
}
