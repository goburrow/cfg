package structs

import (
	"fmt"
	"math"
	"testing"
)

func TestUnmarshal(t *testing.T) {
	in := map[string]string{
		"string":   "a string",
		"integer":  "-1234",
		"uinteger": "1234",
		"bool":     "true",
		"float":    "1.234",
	}
	var out struct {
		String   string  `cfg:"string"`
		Integer  int     `cfg:"integer"`
		Uinteger uint    `cfg:"uinteger"`
		Bool     bool    `cfg:"bool"`
		Float    float64 `cfg:"float"`
	}

	err := Unmarshal(in, &out)
	if err != nil {
		t.Fatal(err)
	}
	if out.String != "a string" {
		t.Errorf("unexpected string: %v", out.String)
	}
	if out.Integer != -1234 {
		t.Errorf("unexpected integer: %v", out.Integer)
	}
	if out.Uinteger != 1234 {
		t.Errorf("unexpected uinteger: %v", out.Uinteger)
	}
	if out.Bool != true {
		t.Errorf("unexpected bool: %v", out.Bool)
	}
	if math.Abs(out.Float-1.234) > 0.001 {
		t.Errorf("unexpected float: %v", out.Float)
	}
}

func ExampleUnmarshal() {
	in := map[string]string{
		"name":    "Foo Bar",
		"age":     "54",
		"married": "true",
		"height":  "1.75",
	}
	var out struct {
		Name    string  `cfg:"name"`
		Age     int     `cfg:"age"`
		Married bool    `cfg:"married"`
		Height  float64 `cfg:"height"`
	}
	err := Unmarshal(in, &out)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", out)
	// Output:
	// {Name:Foo Bar Age:54 Married:true Height:1.75}
}

func TestUnmarshalInvalidInt(t *testing.T) {
	in := map[string]string{
		"integer": "abc",
	}
	var out struct {
		Integer int `cfg:"integer"`
	}
	err := Unmarshal(in, &out)
	if err == nil {
		t.Fatal("error expected")
	}
}
