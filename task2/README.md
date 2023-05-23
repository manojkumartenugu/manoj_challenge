# JSON Data Processor

This repository contains a JSON data processing program written in Go. The program takes a JSON input string, performs modifications on the data, and outputs the modified JSON.

## Functionality

### `main` 

- The `main`  function serves as the entry point of the program. It defines a JSON `input` string named input that represents a structured data object. The `jsonObj` function is called with the input string as an argument. It performs the processing and modification of the JSON data.

### `jsonObj(input string) (string, error)`

- The `jsonObj` function takes a JSON input string as its parameter and returns a modified JSON output string. It uses the `json.Unmarshal` function to parse the input string into a `map[string]interface{}` named `data`. An empty slice of `map[string]interface{}` named `dataModified` is created to store the modified data elements.

- The function then iterates over each key-value pair in the `data` map using a for loop. It skips fields with empty keys and calls the `modVal` function to modify the corresponding value. The modified key-value pairs are added to the `dataModified` slice as maps with a single key-value pair.

- The `sort.Slice` function is used to sort the `dataModified` slice based on the keys of the maps. Finally, the `json.Marshal` function is used to convert the sorted `dataModified` slice into a JSON output string.


## `modVal(value interface{}) (interface{}, error)`

- The `modVal` function is a recursive function that handles different value types (map, string, and list) and performs specific modifications. When the value type is a map, it calls the `modMap` function to recursively process the key-value pairs of the map and construct a modified map.

- When the value type is a string, the `string_mod` function is called. It trims leading and trailing whitespace and checks if the string is in RFC3339 format. If it is, the string is transformed into a Unix Epoch timestamp. If the string is empty, an error is returned.

- When the value type is a list, the `list_updated` function is called. It iterates over the list elements and recursively modifies each element.

## `modMap(data map[string]interface{}) (map[string]interface{}, error)`

- The `modMap` function is called when the value type is a map. It recursively processes the key-value pairs of the map and constructs a modified map.

## `string_mod(value string) (interface{}, error)`

- The `string_mod` function is called when the value type is a string. It trims leading and trailing whitespace and checks if the string is in RFC3339 format. If it is, the string is transformed into a Unix Epoch timestamp. If the string is empty, an error is returned.

## `list_updated(data []interface{}) ([]interface{}, error)`

- The `list_updated` function is called when the value type is a list. It iterates over the list elements and recursively modifies each element.

## `getKey(data map[string]interface{}) string`

- The `getKey` function retrieves the key from a map by iterating over the keys and returning the first key encountered. This is used to extract the key from a map when constructing the `dataModified` slice.

Finally, the sorted dataModified slice is returned as a JSON output string from the jsonObj function.