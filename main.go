package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

const format = "\t\"%s%s\": %v,\n"

func flattenJSONHelper(jsonObj map[string]interface{}, namespace string) string {
	flattened := ""
	for k, v := range jsonObj {
		switch obj := v.(type) {
		case map[string]interface{}:
			//Recursively add children
			flattened += flattenJSONHelper(obj, namespace+k+".") //Recurse with a new namespace
		case nil:
			flattened += fmt.Sprintf(format, namespace, k, "null")
		case string:
			flattened += fmt.Sprintf(format, namespace, k, "\""+obj+"\"")
		default:
			flattened += fmt.Sprintf(format, namespace, k, v)
		}
	}
	return flattened
}

func flattenJSON(raw []byte) (string, error) {
	jsonObj := make(map[string]interface{})
	err := json.Unmarshal(raw, &jsonObj)
	if err != nil {
		return "", errors.New("Could not parse input")
	}
	flattened := flattenJSONHelper(jsonObj, "")
	flattened = flattened[:len(flattened)-2] //Remove the ,\n at the end of the output
	return "{\n" + flattened + "\n}", nil
}

func main() {
	buf := make([]byte, 0)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		buf = append(buf, scanner.Bytes()...)
	}

	output, err := flattenJSON(buf)
	if err != nil {
		panic(err)
	}
	fmt.Println(output)
}
