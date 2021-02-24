package common

import (
	"fmt"
	"reflect"
)

func SerializeData(data interface{}) ([]byte, error) {
	return serializeData(reflect.ValueOf(data))
}

func serializeData(v reflect.Value) ([]byte, error) {
	switch v.Kind() {
	case reflect.Bool:
		return boolEncode(v)
	case reflect.Uint8:
		return uint8Encode(v)
	case reflect.Array:
		switch v.Type().Elem().Kind() {
		case reflect.Uint8:
			b := make([]byte, 0, v.Len())
			for i := 0; i < v.Len(); i++ {
				b = append(b, byte(v.Index(i).Uint()))
			}
			return b, nil
		}
		return nil, fmt.Errorf("unsupport type: %v, elem: %v", v.Kind(), v.Elem().Kind())
	case reflect.Struct:
		data := make([]byte, 0, 1024)
		for i := 0; i < v.NumField(); i++ {
			field := v.Field(i)
			d, err := serializeData(field)
			if err != nil {
				return nil, fmt.Errorf("unsupport type: %v", field)
			}
			data = append(data, d...)
		}
		return data, nil
	}
	return nil, fmt.Errorf("unsupport type: %v", v.Kind())
}

func boolEncode(v reflect.Value) ([]byte, error) {
	if v.Bool() {
		return []byte{1}, nil
	}
	return []byte{0}, nil
}

func uint8Encode(v reflect.Value) ([]byte, error) {
	return []byte{uint8(v.Uint())}, nil
}
