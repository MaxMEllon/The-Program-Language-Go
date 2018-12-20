// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 339.

package sexpr

import (
	"bytes"
	"fmt"
	"reflect"
)

//!+Marshal
// Marshal encodes a Go value in S-expression form.
func Marshal(v interface{}) ([]byte, error) {
	var buf bytes.Buffer
	buf.WriteByte('{')
	if err := encode(&buf, reflect.ValueOf(v)); err != nil {
		return nil, err
	}
	buf.WriteByte('}')
	return buf.Bytes(), nil
}

//!-Marshal

// encode writes to buf an S-expression representation of v.
//!+encode
func encode(buf *bytes.Buffer, v reflect.Value) error {
	switch v.Kind() {
	case reflect.Invalid:
		buf.WriteString("null")

	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		fmt.Fprintf(buf, "%d", v.Int())

	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		fmt.Fprintf(buf, "%d", v.Uint())

	/*
		json では 複素数をサポートしない
		case reflect.Complex64, reflect.Complex128:
			fmt.Fprintf(buf, "%%C (%1.1f %1.1f)", real(v.Complex()), imag(v.Complex()))
	*/

	case reflect.String:
		fmt.Fprintf(buf, "%q", v.String())

	case reflect.Ptr:
		return encode(buf, v.Elem())

	case reflect.Bool:
		fmt.Fprint(buf, v.Bool())

	case reflect.Float32, reflect.Float64:
		fmt.Fprintf(buf, "%f", v.Float())

	case reflect.Array, reflect.Slice: // [value, ...]
		buf.WriteByte('[')
		for i := 0; i < v.Len(); i++ {
			if i > 0 {
				buf.WriteByte(',')
			}
			if err := encode(buf, v.Index(i)); err != nil {
				return err
			}
		}
		buf.WriteByte(']')

	case reflect.Struct: // { name: { key: value } ...}
		fmt.Fprintf(buf, "\"%s\":", v.Type().Name())
		buf.WriteByte('{')
		for i := 0; i < v.NumField(); i++ {
			if i > 0 {
				buf.WriteByte(' ')
			}
			fmt.Fprintf(buf, "\"%s\":", v.Type().Field(i).Name)
			if err := encode(buf, v.Field(i)); err != nil {
				return err
			}
			if i != v.NumField()-1 {
				buf.WriteByte(',')
			}
		}
		buf.WriteByte('}')

	case reflect.Map: // { key: value }, ...
		buf.WriteByte('{')
		for i, key := range v.MapKeys() {
			if i > 0 {
				buf.WriteByte(' ')
			}
			if err := encode(buf, key); err != nil {
				return err
			}
			buf.WriteByte(':')
			if err := encode(buf, v.MapIndex(key)); err != nil {
				return err
			}
			if i != len(v.MapKeys())-1 {
				buf.WriteByte(',')
			}
		}
		buf.WriteByte('}')

	default: // chan, func, interface
		return fmt.Errorf("unsupported type: %s", v.Type())
	}
	return nil
}

//!-encode
