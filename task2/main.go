package main

import (
	"encoding/json"
	"fmt"
	"log"
	"sort"
	"strings"
	"time"
)

func main() {
	input := `"number_1": {
		"N": "1.50"
		},
		"string_1": {
		"S": "784498 "
		},
		"string_2": {
		"S": "2014-07-16T20:55:46Z"
		},
		"map_1": {
		"M": {
		"bool_1": {
		"BOOL": "truthy"
		},
		"null_1": {
		"NULL ": "true"
		},
		"list_1": {
		"L": [
		{
		"S": ""
		},
		{
		"N": "011"
		},
		{
		"N": "5215s"
		},
		{
		"BOOL": "f"
		},
		{
		"NULL": "0"
		}
		]
		}
		}
		},
		"list_2": {
		"L": "noop"
		},
		"list_3": {
		"L": [
		"noop"
		]
		},
		"": {
		"S": "noop"
		}
		}`
	output, err := jsonObj(input)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(output)
}

func jsonObj(input string) (string, error) {
	var data map[string]interface{}
	err := json.Unmarshal([]byte(input), &data)
	if err != nil {
		return "", err
	}
	var dataModified []map[string]interface{}
	for key, value := range data {
		// Skip fields with empty keys
		if key == "" {
			continue
		}
		modifyVal, err := modVal(value)
		if err != nil {
			// Skip fields with invalid values
			continue
		}
		dataModified = append(dataModified, map[string]interface{}{
			key: modifyVal,
		})
	}
	sort.Slice(dataModified, func(i, j int) bool {
		return getKey(dataModified[i]) < getKey(dataModified[j])
	})
	jsonOutput, err := json.Marshal(dataModified)
	if err != nil {
		return "", err
	}
	return string(jsonOutput), nil
}

func modVal(value interface{}) (interface{}, error) {
	switch v := value.(type) {
	case map[string]interface{}:
		return modMap(v)
	case string:
		return string_mod(v)
	case []interface{}:
		return list_updated(v)
	default:
		return nil, fmt.Errorf("unsupported data type")
	}
}

func modMap(data map[string]interface{}) (map[string]interface{}, error) {
	modMap1 := make(map[string]interface{})
	for key, value := range data {
		modifyVal, err := modVal(value)
		if err != nil {
			// Skip fields with unsupported data types
			continue
		}
		modMap1[key] = modifyVal
	}
	// Omit map if it is empty
	if len(modMap1) == 0 {
		return nil, fmt.Errorf("empty map")
	}
	return modMap1, nil
}

func string_mod(value string) (interface{}, error) {
	// Trim trailing and leading whitespace
	value = strings.TrimSpace(value)
	// Transform RFC3339 formatted strings to Unix Epoch
	t, err := time.Parse(time.RFC3339, value)
	if err == nil {
		return t.Unix(), nil
	}
	// Omit empty string values
	if value == "" {
		return nil, fmt.Errorf("empty string")
	}
	return value, nil
}

func list_updated(data []interface{}) ([]interface{}, error) {
	list1 := make([]interface{}, 0, len(data))
	for _, value := range data {
		modifyVal, err := modVal(value)
		if err != nil {
			// Skip fields with unsupported data types
			continue
		}
		list1 = append(list1, modifyVal)
	}
	// Omit list if it is empty
	if len(list1) == 0 {
		return nil, fmt.Errorf("empty list")
	}
	return list1, nil
}

func getKey(data map[string]interface{}) string {
	for key := range data {
		return key
	}
	return ""
}