package json

import (
	"encoding/json"
	"fmt"
	"math"
	"reflect"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TT struct {
	A     string
	B     int
	C     map[string]string
	D     []int
	Extra json.RawMessage
}

type Extra struct {
	A string
	B int
	C map[string]string
	D []int
}

var testData = []byte(`
{"A":"a", "B":1, "C": {"a":"eserunt dolorconsequat eiusmod  elit occaecat. Qui ex amet commo do reprehenderit excepteur dduis", "b":"eserunt dolorconsequat eiusmod consectetur elit occaecat. Qui ex amet commo do reprehenderit excepteur dduis"}, "D": [1,2,3,4], "Extra": {"A":"eserunt dolorconsequat eiusmod consectetur elit occaecat. Qui ex amet commo do reprehenderit excepteur dduis", "B":1, "C": {"a":"eserunt dolorconsequat eiusmod consectetur elit occaecat. Qui ex amet commo do reprehenderit excepteur dduis", "b":"eserunt dolorconsequat eiusmod consectetur elit occaecat. Qui ex amet commo do reprehenderit excepteur dduis"}, "D": [1,2,3,4]}}
`)

func TestUnMarshal(t *testing.T) {
	var t1, t2 = &TT{}, &TT{}
	var e1, e2 = &Extra{}, &Extra{}
	var err error
	err = Unmarshal(testData, t1)
	assert.Nil(t, err)
	err = json.Unmarshal(testData, t2)
	assert.Nil(t, err)

	err = Unmarshal(t1.Extra, e1)
	assert.Nil(t, err)
	err = json.Unmarshal(t2.Extra, e2)
	assert.Nil(t, err)

	d, _ := json.Marshal(t1)
	assert.Equal(t, d[0], uint8(123))

	assert.Equal(t, e1, e2)
}

func BenchmarkUnmarshal(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var t1 = &TT{}
		var e1 = &Extra{}
		_ = Unmarshal(testData, t1)
		_ = Unmarshal(t1.Extra, e1)
	}
}

func BenchmarkStdUnmarshal(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var t1 = &TT{}
		var e1 = &Extra{}
		_ = json.Unmarshal(testData, t1)
		_ = json.Unmarshal(t1.Extra, e1)
	}
}

var t1 = &T{}
var protoData []byte

func init() {
	_ = json.Unmarshal(testData, t1)
	protoData, _ = t1.XXX_Marshal([]byte{}, true)
}

func BenchmarkProtoUnmarshal(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var t1 = &T{}
		_ = t1.XXX_Unmarshal(protoData)
	}
}

func TestEvaluateTypes(t *testing.T) {
	testCases := []struct {
		In   string
		Data interface{}
		Out  string
		Err  bool
	}{
		// Invalid syntaxes
		{
			In:  `x`,
			Err: true,
		},
		{
			In:  ``,
			Err: true,
		},

		// Null
		{
			In:   `null`,
			Data: nil,
			Out:  `null`,
		},
		// Booleans
		{
			In:   `true`,
			Data: true,
			Out:  `true`,
		},
		{
			In:   `false`,
			Data: false,
			Out:  `false`,
		},

		// Integers
		{
			In:   `0`,
			Data: int64(0),
			Out:  `0`,
		},
		{
			In:   `-0`,
			Data: int64(-0),
			Out:  `0`,
		},
		{
			In:   `1`,
			Data: int64(1),
			Out:  `1`,
		},
		{
			In:   `2147483647`,
			Data: int64(math.MaxInt32),
			Out:  `2147483647`,
		},
		{
			In:   `-2147483648`,
			Data: int64(math.MinInt32),
			Out:  `-2147483648`,
		},
		{
			In:   `9223372036854775807`,
			Data: int64(math.MaxInt64),
			Out:  `9223372036854775807`,
		},
		{
			In:   `-9223372036854775808`,
			Data: int64(math.MinInt64),
			Out:  `-9223372036854775808`,
		},

		// Int overflow
		{
			In:   `9223372036854775808`, // MaxInt64 + 1
			Data: float64(9223372036854775808),
			Out:  `9223372036854776000`,
		},
		{
			In:   `-9223372036854775809`, // MinInt64 - 1
			Data: float64(math.MinInt64),
			Out:  `-9223372036854776000`,
		},

		// Floats
		{
			In:   `0.0`,
			Data: float64(0),
			Out:  `0`,
		},
		{
			In:   `-0.0`,
			Data: float64(-0.0),
			Out:  `-0`,
		},
		{
			In:   `0.5`,
			Data: float64(0.5),
			Out:  `0.5`,
		},
		{
			In:   `1e3`,
			Data: float64(1e3),
			Out:  `1000`,
		},
		{
			In:   `1.5`,
			Data: float64(1.5),
			Out:  `1.5`,
		},
		{
			In:   `-0.3`,
			Data: float64(-.3),
			Out:  `-0.3`,
		},
		{
			// Largest representable float32
			In:   `3.40282346638528859811704183484516925440e+38`,
			Data: float64(math.MaxFloat32),
			Out:  strconv.FormatFloat(math.MaxFloat32, 'g', -1, 64),
		},
		{
			// Smallest float32 without losing precision
			In:   `1.175494351e-38`,
			Data: float64(1.175494351e-38),
			Out:  `1.175494351e-38`,
		},
		{
			// float32 closest to zero
			In:   `1.401298464324817070923729583289916131280e-45`,
			Data: float64(math.SmallestNonzeroFloat32),
			Out:  strconv.FormatFloat(math.SmallestNonzeroFloat32, 'g', -1, 64),
		},
		{
			// Largest representable float64
			In:   `1.797693134862315708145274237317043567981e+308`,
			Data: float64(math.MaxFloat64),
			Out:  strconv.FormatFloat(math.MaxFloat64, 'g', -1, 64),
		},
		{
			// Closest to zero without losing precision
			In:   `2.2250738585072014e-308`,
			Data: float64(2.2250738585072014e-308),
			Out:  `2.2250738585072014e-308`,
		},

		{
			// float64 closest to zero
			In:   `4.940656458412465441765687928682213723651e-324`,
			Data: float64(math.SmallestNonzeroFloat64),
			Out:  strconv.FormatFloat(math.SmallestNonzeroFloat64, 'g', -1, 64),
		},

		{
			// math.MaxFloat64 + 2 overflow
			In:  `1.7976931348623159e+308`,
			Err: true,
		},

		// Strings
		{
			In:   `""`,
			Data: string(""),
			Out:  `""`,
		},
		{
			In:   `"0"`,
			Data: string("0"),
			Out:  `"0"`,
		},
		{
			In:   `"A"`,
			Data: string("A"),
			Out:  `"A"`,
		},
		{
			In:   `"Iñtërnâtiônàlizætiøn"`,
			Data: string("Iñtërnâtiônàlizætiøn"),
			Out:  `"Iñtërnâtiônàlizætiøn"`,
		},

		// Arrays
		{
			In:   `[]`,
			Data: []interface{}{},
			Out:  `[]`,
		},
		{
			In: `[` + strings.Join([]string{
				`null`,
				`true`,
				`false`,
				`0`,
				`9223372036854775807`,
				`0.0`,
				`0.5`,
				`1.0`,
				`1.797693134862315708145274237317043567981e+308`,
				`"0"`,
				`"A"`,
				`"Iñtërnâtiônàlizætiøn"`,
				`[null,true,1,1.0,1.5]`,
				`{"boolkey":true,"floatkey":1.0,"intkey":1,"nullkey":null}`,
			}, ",") + `]`,
			Data: []interface{}{
				nil,
				true,
				false,
				int64(0),
				int64(math.MaxInt64),
				float64(0.0),
				float64(0.5),
				float64(1.0),
				float64(math.MaxFloat64),
				string("0"),
				string("A"),
				string("Iñtërnâtiônàlizætiøn"),
				[]interface{}{nil, true, int64(1), float64(1.0), float64(1.5)},
				map[string]interface{}{"nullkey": nil, "boolkey": true, "intkey": int64(1), "floatkey": float64(1.0)},
			},
			Out: `[` + strings.Join([]string{
				`null`,
				`true`,
				`false`,
				`0`,
				`9223372036854775807`,
				`0`,
				`0.5`,
				`1`,
				strconv.FormatFloat(math.MaxFloat64, 'g', -1, 64),
				`"0"`,
				`"A"`,
				`"Iñtërnâtiônàlizætiøn"`,
				`[null,true,1,1,1.5]`,
				`{"boolkey":true,"floatkey":1,"intkey":1,"nullkey":null}`, // gets alphabetized by Marshal
			}, ",") + `]`,
		},

		// Maps
		{
			In:   `{}`,
			Data: map[string]interface{}{},
			Out:  `{}`,
		},
		{
			In:   `{"boolkey":true,"floatkey":1.0,"intkey":1,"nullkey":null}`,
			Data: map[string]interface{}{"nullkey": nil, "boolkey": true, "intkey": int64(1), "floatkey": float64(1.0)},
			Out:  `{"boolkey":true,"floatkey":1,"intkey":1,"nullkey":null}`, // gets alphabetized by Marshal
		},
	}

	for _, tc := range testCases {
		inputJSON := fmt.Sprintf(`{"data":%s}`, tc.In)
		expectedJSON := fmt.Sprintf(`{"data":%s}`, tc.Out)
		m := map[string]interface{}{}
		err := UnmarshalV2([]byte(inputJSON), &m)
		if tc.Err && err != nil {
			// Expected error
			continue
		}
		if err != nil {
			t.Errorf("%s: error decoding: %v", tc.In, err)
			continue
		}
		if tc.Err {
			t.Errorf("%s: expected error, got none", tc.In)
			continue
		}
		data, ok := m["data"]
		if !ok {
			t.Errorf("%s: decoded object missing data key: %#v", tc.In, m)
			continue
		}
		if !reflect.DeepEqual(tc.Data, data) {
			t.Errorf("%s: expected\n\t%#v (%v), got\n\t%#v (%v)", tc.In, tc.Data, reflect.TypeOf(tc.Data), data, reflect.TypeOf(data))
			continue
		}

		outputJSON, err := Marshal(m)
		if err != nil {
			t.Errorf("%s: error encoding: %v", tc.In, err)
			continue
		}

		if expectedJSON != string(outputJSON) {
			t.Errorf("%s: expected\n\t%s, got\n\t%s", tc.In, expectedJSON, string(outputJSON))
			continue
		}
	}
}
