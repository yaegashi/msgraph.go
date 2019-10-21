package jsonx

import "testing"

type MarshalExtra1 struct {
	X string
	Y int
	R map[string]interface{} `json:"-" jsonx:"extra"`
}

type MarshalExtra2 struct {
	R map[string]interface{} `json:"-" jsonx:"extra"`
	X string
	Y int
}

func TestMarshalExtra(t *testing.T) {
	tests := []struct {
		in   interface{}
		want string
		ok   bool
	}{
		{MarshalExtra1{X: "123", Y: 123, R: map[string]interface{}{"A": "123", "B": 123}}, `{"X":"123","Y":123,"A":"123","B":123}`, true},
		{MarshalExtra2{X: "123", Y: 123, R: map[string]interface{}{"A": "123", "B": 123}}, `{"A":"123","B":123,"X":"123","Y":123}`, true},
	}

	for i, tt := range tests {
		b, err := Marshal(tt.in)
		if ok := (err == nil); ok != tt.ok {
			if err != nil {
				t.Errorf("test %d, unexpected failure: %v", i, err)
			} else {
				t.Errorf("test %d, unexpected success", i)
			}
		}
		if got := string(b); got != tt.want {
			t.Errorf("test %d, Marshal(%#v) = %q, want %q", i, tt.in, got, tt.want)
		}
	}
}
