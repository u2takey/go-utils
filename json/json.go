package json

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"

	"github.com/json-iterator/go"
)

// json 包的目的是简化替换 json 库的流程, 项目中直接搜索 "encoding/json" 替换为 "github.com/u2takey/go-utils/json" 就可以

type InvalidUnmarshalError = json.InvalidUnmarshalError
type UnmarshalTypeError = json.UnmarshalTypeError
type UnsupportedTypeError = json.UnsupportedTypeError
type MarshalerError = json.MarshalerError
type SyntaxError = json.SyntaxError

type RawMessage = json.RawMessage
type Number = json.Number

var jsoniterStd = jsoniter.ConfigCompatibleWithStandardLibrary

// Valid delegates to jsoniterStd
func Valid(data []byte) bool {
	return jsoniterStd.Valid(data)
}

// NewDecoder delegates to jsoniterStd
func NewDecoder(reader io.Reader) *jsoniter.Decoder {
	return jsoniterStd.NewDecoder(reader)
}

// NewEncoder delegates to jsoniterStd
func NewEncoder(w io.Writer) *jsoniter.Encoder {
	return jsoniter.NewEncoder(w)
}

// Marshal delegates to jsoniterStd
// It is only here so this package can be a drop-in for common encoding/json uses
func Marshal(v interface{}) ([]byte, error) {
	return jsoniterStd.Marshal(v)
}

// Unmarshal delegates to jsoniterStd
func Unmarshal(data []byte, v interface{}) error {
	return jsoniterStd.Unmarshal(data, v)
}

// limit recursive depth to prevent stack overflow errors
const maxDepth = 10000

// Unmarshal unmarshals the given data
// If v is a *map[string]interface{}, numbers are converted to int64 or float64
func UnmarshalV2(data []byte, v interface{}) error {
	switch v := v.(type) {
	case *map[string]interface{}:
		// Build a decoder from the given data
		decoder := json.NewDecoder(bytes.NewBuffer(data))
		// Preserve numbers, rather than casting to float64 automatically
		decoder.UseNumber()
		// Run the decode
		if err := decoder.Decode(v); err != nil {
			return err
		}
		// If the decode succeeds, post-process the map to convert json.Number objects to int64 or float64
		return convertMapNumbers(*v, 0)

	case *[]interface{}:
		// Build a decoder from the given data
		decoder := jsoniter.NewDecoder(bytes.NewBuffer(data))
		// Preserve numbers, rather than casting to float64 automatically
		decoder.UseNumber()
		// Run the decode
		if err := decoder.Decode(v); err != nil {
			return err
		}
		// If the decode succeeds, post-process the map to convert json.Number objects to int64 or float64
		return convertSliceNumbers(*v, 0)

	default:
		return jsoniterStd.Unmarshal(data, v)
	}
}

// convertMapNumbers traverses the map, converting any json.Number values to int64 or float64.
// values which are map[string]interface{} or []interface{} are recursively visited
func convertMapNumbers(m map[string]interface{}, depth int) error {
	if depth > maxDepth {
		return fmt.Errorf("exceeded max depth of %d", maxDepth)
	}

	var err error
	for k, v := range m {
		switch v := v.(type) {
		case json.Number:
			m[k], err = convertNumber(v)
		case map[string]interface{}:
			err = convertMapNumbers(v, depth+1)
		case []interface{}:
			err = convertSliceNumbers(v, depth+1)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// convertSliceNumbers traverses the slice, converting any json.Number values to int64 or float64.
// values which are map[string]interface{} or []interface{} are recursively visited
func convertSliceNumbers(s []interface{}, depth int) error {
	if depth > maxDepth {
		return fmt.Errorf("exceeded max depth of %d", maxDepth)
	}

	var err error
	for i, v := range s {
		switch v := v.(type) {
		case json.Number:
			s[i], err = convertNumber(v)
		case map[string]interface{}:
			err = convertMapNumbers(v, depth+1)
		case []interface{}:
			err = convertSliceNumbers(v, depth+1)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// convertNumber converts a json.Number to an int64 or float64, or returns an error
func convertNumber(n json.Number) (interface{}, error) {
	// Attempt to convert to an int64 first
	if i, err := n.Int64(); err == nil {
		return i, nil
	}
	// Return a float64 (default json.Decode() behavior)
	// An overflow will return an error
	return n.Float64()
}

type JsoniterSerializer struct{}

func (*JsoniterSerializer) Serialize(v interface{}) ([]byte, error) {
	return jsoniterStd.Marshal(v)
}

func (*JsoniterSerializer) Deserialize(data []byte, v interface{}) error {
	return jsoniterStd.Unmarshal(data, v)
}

func JsonAdapt(input, target interface{}) error {
	inputB, err := Marshal(input)
	if err != nil {
		return fmt.Errorf("fail to marshal input: %v, err: %s", input, err.Error())
	}
	if err := Unmarshal(inputB, target); err != nil {
		return fmt.Errorf("fail to unmarshal to target type, err: %s", err.Error())
	}
	return nil
}
