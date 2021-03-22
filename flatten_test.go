package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var spaceReplacer = strings.NewReplacer("\n", "", "\t", "", " ", "")

//TestFlattenSimple makes sure that singleton objects are encoded as themselves
func TestFlattenSimple(t *testing.T) {
	testCases := []struct {
		desc  string
		input string
	}{
		{
			desc:  "simple_int",
			input: "{\"a\": 1}",
		},
		{
			desc:  "simple_bool",
			input: "{\"a\": true}",
		},
		{
			desc:  "simple_string",
			input: "{\"a\": \"hello world\"}",
		},
		{
			desc:  "simple_null",
			input: "{\"a\": null}",
		},
		{
			desc:  "simple_float",
			input: "{\"a\": 3.141592}",
		},
		{
			desc: "simple_all",
			input: `{
				"a": 1,
				"b": true,
				"c": "hello world",
				"d": null,
				"e": 3.141592
			}`,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			output, err := flattenJSON([]byte(tC.input))
			assert.Nil(t, err)
			assert.Equal(t, spaceReplacer.Replace(tC.input), spaceReplacer.Replace(output))
		})
	}
}
