// Package structs provides a function to unmarshal properties to structs.
//
// Use Go struct tag to define property mappings. For example:
//
// 	type Data struct {
//		String   string  `cfg:"string"`
//		Integer  int     `cfg:"integer"`
//		Uinteger uint    `cfg:"uinteger"`
//		Bool     bool    `cfg:"bool"`
//		Float    float64 `cfg:"float"`
//	}
//
// See strconv package to see how string value is converted to int, bool
// or float.
package structs

import (
	"fmt"
	"reflect"
	"strconv"
)

const (
	tagName = "cfg"
)

// Unmarshal unmarshals properties p into v.
// v must be a pointer to a struct.
func Unmarshal(p map[string]string, v interface{}) error {
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Ptr {
		return fmt.Errorf("value must be pointer to a struct: %+v", reflect.TypeOf(v))
	}
	val := rv.Elem()

	for i := 0; i < val.NumField(); i++ {
		tag := val.Type().Field(i).Tag.Get(tagName)
		if tag == "" {
			continue
		}
		strVal := p[tag]
		if strVal == "" {
			continue
		}
		field := val.Field(i)
		switch field.Kind() {
		case reflect.String:
			field.SetString(strVal)
		case reflect.Int, reflect.Int64, reflect.Int32, reflect.Int16, reflect.Int8:
			x, err := strconv.ParseInt(strVal, 10, 0)
			if err != nil {
				return err
			}
			field.SetInt(x)
		case reflect.Uint, reflect.Uint64, reflect.Uint32, reflect.Uint16, reflect.Uint8:
			x, err := strconv.ParseUint(strVal, 10, 0)
			if err != nil {
				return err
			}
			field.SetUint(x)
		case reflect.Bool:
			x, err := strconv.ParseBool(strVal)
			if err != nil {
				return err
			}
			field.SetBool(x)
		case reflect.Float64, reflect.Float32:
			x, err := strconv.ParseFloat(strVal, 0)
			if err != nil {
				return err
			}
			field.SetFloat(x)
		}

	}
	return nil
}
