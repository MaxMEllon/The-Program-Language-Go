package params

import (
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"
)

//!+Unpack

// Unpack populates the fields of the struct pointed to by ptr
// from the HTTP request parameters in req.
func Unpack(req *http.Request, ptr interface{}) error {
	if err := req.ParseForm(); err != nil {
		return err
	}

	// Build map of fields keyed by effective name.
	fields := make(map[string]reflect.Value)
	v := reflect.ValueOf(ptr).Elem() // the struct variable
	for i := 0; i < v.NumField(); i++ {
		fieldInfo := v.Type().Field(i) // a reflect.StructField
		tag := fieldInfo.Tag           // a reflect.StructTag
		name := tag.Get("http")
		if name == "" {
			name = strings.ToLower(fieldInfo.Name)
		}
		fields[name] = v.Field(i)
	}

	// Update struct field for each parameter in the request.
	for name, values := range req.Form {
		f := fields[name]
		if !f.IsValid() {
			continue // ignore unrecognized HTTP parameters
		}
		for _, value := range values {
			if f.Kind() == reflect.Slice {
				elem := reflect.New(f.Type().Elem()).Elem()
				if err := populate(elem, value); err != nil {
					return fmt.Errorf("%s: %v", name, err)
				}
				f.Set(reflect.Append(f, elem))
			} else {
				if err := populate(f, value); err != nil {
					return fmt.Errorf("%s: %v", name, err)
				}
			}
		}
	}
	return nil
}

func toString(c reflect.Value) string {
	switch c.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(c.Int(), 10)
	case reflect.String:
		return c.String()
	case reflect.Bool:
		return strconv.FormatBool(c.Bool())
	case reflect.Array, reflect.Slice:
		res := ""
		for i := 0; i < c.Len(); i++ {
			res += toString(c.Index(i))
			if i != c.Len()-1 {
				res += ","
			}
		}
		return res
	}
	return ""
}

//!-Unpack

func Pack(ptr interface{}) (string, error) {
	q := url.Values{}
	v := reflect.ValueOf(ptr).Elem()
	if v.Kind() != reflect.Struct {
		return "", fmt.Errorf("should struct types")
	}
	for i := 0; i < v.NumField(); i++ {
		child := v.Field(i)
		fieldInfo := v.Type().Field(i)
		key := fieldInfo.Tag.Get("http")
		key = strings.Trim(key, " ")
		//fmt.Println(key)
		if key == "" {
			return "", fmt.Errorf("not allow empty key")
		}
		switch child.Kind() {
		case reflect.Int, reflect.String, reflect.Bool:
			q.Add(key, toString(child))
		case reflect.Array, reflect.Slice:
			for i := 0; i < child.Len(); i++ {
				q.Add(key+"[]", toString(child.Index(i)))
			}
		case reflect.Map:
			for _, k := range child.MapKeys() {
				q.Add(key+"."+toString(k), toString(child.MapIndex(k)))
			}
		}
	}
	return q.Encode(), nil
}

//!+populate
func populate(v reflect.Value, value string) error {
	switch v.Kind() {
	case reflect.String:
		v.SetString(value)

	case reflect.Int:
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return err
		}
		v.SetInt(i)

	case reflect.Bool:
		b, err := strconv.ParseBool(value)
		if err != nil {
			return err
		}
		v.SetBool(b)

	default:
		return fmt.Errorf("unsupported kind %s", v.Type())
	}
	return nil
}

//!-populate
