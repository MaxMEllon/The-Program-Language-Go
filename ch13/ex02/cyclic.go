package cyclic

import (
	"reflect"
	"unsafe"
)

func IsCyclic(ptr interface{}) bool {
	seen := make(map[unsafe.Pointer]bool)
	return isCyclic(reflect.ValueOf(ptr), seen)
}

func isCyclic(v reflect.Value, seen map[unsafe.Pointer]bool) bool {
	k := v.Kind()
	if v.CanAddr() && k != reflect.Struct && k != reflect.Array && k != reflect.Slice && k != reflect.Map {
		xptr := unsafe.Pointer(v.UnsafeAddr())

		if seen[xptr] {
			return true
		}
		seen[xptr] = true
	}

	switch k {
	case reflect.Ptr, reflect.Interface:
		return isCyclic(v.Elem(), seen)
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			if isCyclic(v.Field(i), seen) {
				return true
			}
		}
	case reflect.Map:
		for _, k := range v.MapKeys() {
			if isCyclic(v.MapIndex(k), seen) {
				return true
			}
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < v.Len(); i++ {
			if isCyclic(v.Index(i), seen) {
				return true
			}
		}
	}

	return false
}
