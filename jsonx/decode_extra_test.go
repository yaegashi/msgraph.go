package jsonx

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
	"testing"
)

type UnmarshalExtra1 struct {
	X string
	Y int
	R map[string]interface{} `json:"-" jsonx:"true"`
}

type UnmarshalExtra2 struct {
	X string
	Y int
	R map[string]string `json:"-" jsonx:"true"`
}

type UnmarshalExtra3 struct {
	X string
	Y int
	R int `json:"-" jsonx:"true"`
}

var unmarshalExtraTests = []unmarshalTest{
	{
		in:  `{"X":"123","Y":123,"A":"123","B":123}`,
		ptr: new(UnmarshalExtra1),
		out: UnmarshalExtra1{X: "123", Y: 123, R: map[string]interface{}{"A": "123", "B": 123.0}},
	},
	{
		in:  `{"X":"123","Y":123,"A":"123","B":123}`,
		ptr: new(UnmarshalExtra2),
		err: fmt.Errorf("json: cannot unmarshal number into Go struct field UnmarshalExtra2.R of type string"),
	},
	{
		in:  `{"X":"123","Y":123,"A":"123","B":123}`,
		ptr: new(UnmarshalExtra3),
		out: UnmarshalExtra3{X: "123", Y: 123},
	},
}

func TestUnmarshalExtra(t *testing.T) {
	for i, tt := range unmarshalExtraTests {
		var scan scanner
		in := []byte(tt.in)
		if err := checkValid(in, &scan); err != nil {
			if !equalError(err, tt.err) {
				t.Errorf("#%d: checkValid: %#v", i, err)
				continue
			}
		}
		if tt.ptr == nil {
			continue
		}

		// v = new(right-type)
		v := reflect.New(reflect.TypeOf(tt.ptr).Elem())
		dec := NewDecoder(bytes.NewReader(in))
		if tt.useNumber {
			dec.UseNumber()
		}
		if tt.disallowUnknownFields {
			dec.DisallowUnknownFields()
		}
		if err := dec.Decode(v.Interface()); !equalError(err, tt.err) {
			t.Errorf("#%d: %v, want %v", i, err, tt.err)
			continue
		} else if err != nil {
			continue
		}
		if !reflect.DeepEqual(v.Elem().Interface(), tt.out) {
			t.Errorf("#%d: mismatch\nhave: %#+v\nwant: %#+v", i, v.Elem().Interface(), tt.out)
			data, _ := Marshal(v.Elem().Interface())
			println(string(data))
			data, _ = Marshal(tt.out)
			println(string(data))
			continue
		}

		// Check round trip also decodes correctly.
		if tt.err == nil {
			enc, err := Marshal(v.Interface())
			if err != nil {
				t.Errorf("#%d: error re-marshaling: %v", i, err)
				continue
			}
			if tt.golden && !bytes.Equal(enc, in) {
				t.Errorf("#%d: remarshal mismatch:\nhave: %s\nwant: %s", i, enc, in)
			}
			vv := reflect.New(reflect.TypeOf(tt.ptr).Elem())
			dec = NewDecoder(bytes.NewReader(enc))
			if tt.useNumber {
				dec.UseNumber()
			}
			if err := dec.Decode(vv.Interface()); err != nil {
				t.Errorf("#%d: error re-unmarshaling %#q: %v", i, enc, err)
				continue
			}
			if !reflect.DeepEqual(v.Elem().Interface(), vv.Elem().Interface()) {
				t.Errorf("#%d: mismatch\nhave: %#+v\nwant: %#+v", i, v.Elem().Interface(), vv.Elem().Interface())
				t.Errorf("     In: %q", strings.Map(noSpace, string(in)))
				t.Errorf("Marshal: %q", strings.Map(noSpace, string(enc)))
				continue
			}
		}
	}
}
