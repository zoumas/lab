package walk

import (
	"reflect"
)

func Walk(x any, fn func(input string)) {
	v := getValue(x)

	switch v.Kind() {
	case reflect.String:
		fn(v.String())
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			field := v.Field(i)

			switch field.Kind() {
			case reflect.String:
				fn(field.String())
			case reflect.Struct:
				Walk(field.Interface(), fn)
			}
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < v.Len(); i++ {
			Walk(v.Index(i).Interface(), fn)
		}
	case reflect.Map:
		for _, key := range v.MapKeys() {
			Walk(v.MapIndex(key).Interface(), fn)
		}
	case reflect.Chan:
		for r, ok := v.Recv(); ok; r, ok = v.Recv() {
			Walk(r.Interface(), fn)
		}
	case reflect.Func:
		fnResult := v.Call(nil)
		for _, r := range fnResult {
			Walk(r.Interface(), fn)
		}
	}
}

func getValue(x any) reflect.Value {
	v := reflect.ValueOf(x)
	if v.Kind() == reflect.Pointer {
		v = v.Elem()
	}
	return v
}
