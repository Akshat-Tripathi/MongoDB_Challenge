package main

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func mapEquals(a, b map[string]interface{}) bool {
	for k, v := range a {
		val, ok := b[k]
		if !ok {
			return false
		}
		switch obj := v.(type) {
		case map[string]interface{}:
			ok := mapEquals(obj, val.(map[string]interface{}))
			if !ok {
				return false
			}
		default:
			if val != v {
				return false
			}
		}
	}
	return true
}

func testEquality(input, output string) bool {
	jsonInput := make(map[string]interface{})
	err := json.Unmarshal([]byte(input), &jsonInput)
	if err != nil {
		return false
	}
	jsonOutput := make(map[string]interface{})
	err = json.Unmarshal([]byte(input), &jsonOutput)
	if err != nil {
		return false
	}
	return mapEquals(jsonInput, jsonOutput)
}

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
			assert.True(t, testEquality(tC.input, output))
		})
	}
}

func TestFlattenNested(t *testing.T) {
	testCases := []struct {
		desc, input, output string
	}{
		{
			desc:   "",
			input:  `{"a": {"a": {"a": {"a": {"a": {"a": {"a": 1}}}}}}}`,
			output: `{"a.a.a.a.a.a.a": 1}`,
		},
		{
			desc: "",
			input: `{
				"a": {
					"a": {
						"a": {
							"a": {
								"a": {
									"a": {"a": 1},
									"b": 3}
									}}}}}`,
			output: `{"a.a.a.a.a.a.a": 1, "a.a.a.a.a.b"}`,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			output, err := flattenJSON([]byte(tC.input))
			assert.Nil(t, err)
			assert.True(t, testEquality(tC.input, output))
		})
	}
}

func TestInvalidJSON(t *testing.T) {
	_, err := flattenJSON([]byte("{hello"))
	assert.NotNil(t, err)
}
