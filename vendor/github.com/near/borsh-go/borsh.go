package borsh

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"math"
	"math/big"
	"reflect"
	"sort"
)

// Deserialize `data` according to the schema of `s`, and store the value into it. `s` must be a pointer type variable
// that points to the original schema of `data`.
func Deserialize(s interface{}, data []byte) error {
	reader := bytes.NewReader(data)
	v := reflect.ValueOf(s)
	if v.Kind() != reflect.Ptr {
		return errors.New("passed struct must be pointer")
	}
	result, err := deserialize(reflect.TypeOf(s).Elem(), reader)
	if err != nil {
		return err
	}
	v.Elem().Set(reflect.ValueOf(result))
	return nil
}

func read(r io.Reader, n int) ([]byte, error) {
	b := make([]byte, n)
	l, err := r.Read(b)
	if l != n {
		return nil, errors.New("failed to read required bytes")
	}
	if err != nil {
		return nil, err
	}
	return b, nil
}

func deserialize(t reflect.Type, r io.Reader) (interface{}, error) {
	if t.Kind() == reflect.Uint8 {
		tmp, err := read(r, 1)
		if err != nil {
			return nil, err
		}
		e := reflect.New(t)
		e.Elem().Set(reflect.ValueOf(uint8(tmp[0])).Convert(t))
		return e.Elem().Interface(), nil
	}

	switch t.Kind() {
	case reflect.Bool:
		tmp, err := read(r, 1)
		if err != nil {
			return nil, err
		}
		switch tmp[0] {
		case 0:
			return false, nil
		case 1:
			return true, nil
		default:
			return nil, fmt.Errorf("expected bool is 0 or 1, got %v", tmp[0])
		}
	case reflect.Int8:
		tmp, err := read(r, 1)
		if err != nil {
			return nil, err
		}
		return int8(tmp[0]), nil
	case reflect.Int16:
		tmp, err := read(r, 2)
		if err != nil {
			return nil, err
		}
		return int16(binary.LittleEndian.Uint16(tmp)), nil
	case reflect.Int32:
		tmp, err := read(r, 4)
		if err != nil {
			return nil, err
		}
		return int32(binary.LittleEndian.Uint32(tmp)), nil
	case reflect.Int64:
		tmp, err := read(r, 8)
		if err != nil {
			return nil, err
		}
		return int64(binary.LittleEndian.Uint64(tmp)), nil
	case reflect.Int:
		tmp, err := read(r, 8)
		if err != nil {
			return nil, err
		}
		return int(binary.LittleEndian.Uint64(tmp)), nil
	case reflect.Uint8:
		tmp, err := read(r, 1)
		if err != nil {
			return nil, err
		}
		return uint8(tmp[0]), nil
	case reflect.Uint16:
		tmp, err := read(r, 2)
		if err != nil {
			return nil, err
		}
		return uint16(binary.LittleEndian.Uint16(tmp)), nil
	case reflect.Uint32:
		tmp, err := read(r, 4)
		if err != nil {
			return nil, err
		}
		return uint32(binary.LittleEndian.Uint32(tmp)), nil
	case reflect.Uint64:
		tmp, err := read(r, 8)
		if err != nil {
			return nil, err
		}
		return uint64(binary.LittleEndian.Uint64(tmp)), nil
	case reflect.Uint:
		tmp, err := read(r, 8)
		if err != nil {
			return nil, err
		}
		return uint(binary.LittleEndian.Uint64(tmp)), nil
	case reflect.Float32:
		tmp, err := read(r, 4)
		if err != nil {
			return nil, err
		}
		bits := binary.LittleEndian.Uint32(tmp)
		f := math.Float32frombits(bits)
		if math.IsNaN(float64(f)) {
			return nil, errors.New("NaN for float not allowed")
		}
		return f, nil
	case reflect.Float64:
		tmp, err := read(r, 8)
		if err != nil {
			return nil, err
		}
		bits := binary.LittleEndian.Uint64(tmp)
		f := math.Float64frombits(bits)
		if math.IsNaN(f) {
			return nil, errors.New("NaN for float not allowed")
		}
		return f, nil
	case reflect.String:
		tmp, err := read(r, 4)
		if err != nil {
			return nil, err
		}
		l := int(binary.LittleEndian.Uint32(tmp))
		if l == 0 {
			return "", nil
		}
		tmp2, err := read(r, l)
		if err != nil {
			return nil, err
		}
		s := string(tmp2)
		return s, nil
	case reflect.Array:
		l := t.Len()
		a := reflect.New(t).Elem()
		for i := 0; i < l; i++ {
			av, err := deserialize(t.Elem(), r)
			if err != nil {
				return nil, err
			}
			a.Index(i).Set(reflect.ValueOf(av))
		}
		return a.Interface(), nil
	case reflect.Slice:
		tmp, err := read(r, 4)
		if err != nil {
			return nil, err
		}
		l := int(binary.LittleEndian.Uint32(tmp))
		a := reflect.New(t).Elem()
		if l == 0 {
			return a.Interface(), nil
		}
		for i := 0; i < l; i++ {
			av, err := deserialize(t.Elem(), r)
			if err != nil {
				return nil, err
			}
			a = reflect.Append(a, reflect.ValueOf(av))
		}
		return a.Interface(), nil
	case reflect.Map:
		tmp, err := read(r, 4)
		if err != nil {
			return nil, err
		}
		l := int(binary.LittleEndian.Uint32(tmp))
		m := reflect.MakeMap(t)
		if l == 0 {
			return m.Interface(), nil
		}
		for i := 0; i < l; i++ {
			k, err := deserialize(t.Key(), r)
			if err != nil {
				return nil, err
			}
			v, err := deserialize(t.Elem(), r)
			if err != nil {
				return nil, err
			}
			m.SetMapIndex(reflect.ValueOf(k), reflect.ValueOf(v))
		}
		return m.Interface(), nil
	case reflect.Ptr:
		tmp, err := read(r, 1)
		if err != nil {
			return nil, err
		}
		valid := uint8(tmp[0])
		if valid == 0 {
			p := reflect.New(t.Elem())
			return p.Interface(), nil
		} else {
			p := reflect.New(t.Elem())
			de, err := deserialize(t.Elem(), r)
			if err != nil {
				return nil, err
			}
			p.Elem().Set(reflect.ValueOf(de))
			return p.Interface(), nil
		}
	case reflect.Struct:
		if t == reflect.TypeOf(*big.NewInt(0)) {
			s, err := deserializeUint128(t, r)
			if err != nil {
				return nil, err
			}
			return s, nil
		} else {
			s, err := deserializeStruct(t, r)
			if err != nil {
				return nil, err
			}
			return s, nil
		}
	}

	return nil, nil
}

func deserializeComplexEnum(t reflect.Type, r io.Reader) (interface{}, error) {
	v := reflect.New(t).Elem()
	// read enum identifier
	tmp, err := read(r, 1)
	if err != nil {
		return nil, err
	}
	enum := Enum(tmp[0])
	v.Field(0).Set(reflect.ValueOf(enum))
	// read enum field, if necessary
	if int(enum)+1 >= t.NumField() {
		return nil, errors.New("complex enum too large")
	}
	fv, err := deserialize(t.Field(int(enum)+1).Type, r)
	if err != nil {
		return nil, err
	}
	v.Field(int(enum) + 1).Set(reflect.ValueOf(fv))

	return v.Interface(), nil
}

func deserializeStruct(t reflect.Type, r io.Reader) (interface{}, error) {
	// handle complex enum, if necessary
	if t.NumField() > 0 {
		// if the first field has type borsh.Enum and is flagged with "borsh_enum"
		// we have a complex enum
		firstField := t.Field(0)
		if firstField.Type.Kind() == reflect.Uint8 &&
			firstField.Tag.Get("borsh_enum") == "true" {
			return deserializeComplexEnum(t, r)
		}
	}

	v := reflect.New(t).Elem()

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag
		if tag.Get("borsh_skip") == "true" {
			continue
		}

		fv, err := deserialize(t.Field(i).Type, r)
		if err != nil {
			return nil, err
		}
		v.Field(i).Set(reflect.ValueOf(fv).Convert(field.Type))
	}

	return v.Interface(), nil
}

func deserializeUint128(t reflect.Type, r io.Reader) (interface{}, error) {
	d, err := read(r, 16)
	if err != nil {
		return nil, err
	}
	// make it big-endian
	for i, j := 0, 15; i < j; i, j = i+1, j-1 {
		d[i], d[j] = d[j], d[i]
	}
	var u big.Int
	u.SetBytes(d[:])
	return u, nil
}

// Serialize `s` into bytes according to Borsh's specification(https://borsh.io/).
//
// The type mapping can be found at https://github.com/near/borsh-go.
func Serialize(s interface{}) ([]byte, error) {
	result := new(bytes.Buffer)

	err := serialize(reflect.ValueOf(s), result)
	return result.Bytes(), err
}

func serializeComplexEnum(v reflect.Value, b io.Writer) error {
	t := v.Type()
	enum := Enum(v.Field(0).Uint())
	// write enum identifier
	if _, err := b.Write([]byte{byte(enum)}); err != nil {
		return err
	}
	// write enum field, if necessary
	if int(enum)+1 >= t.NumField() {
		return errors.New("complex enum too large")
	}
	field := v.Field(int(enum) + 1)
	if field.Kind() == reflect.Struct {
		return serializeStruct(field, b)
	}
	return nil
}

func serializeStruct(v reflect.Value, b io.Writer) error {
	t := v.Type()

	// handle complex enum, if necessary
	if t.NumField() > 0 {
		// if the first field has type borsh.Enum and is flagged with "borsh_enum"
		// we have a complex enum
		firstField := t.Field(0)
		if firstField.Type.Kind() == reflect.Uint8 &&
			firstField.Tag.Get("borsh_enum") == "true" {
			return serializeComplexEnum(v, b)
		}
	}

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Tag.Get("borsh_skip") == "true" {
			continue
		}
		err := serialize(v.Field(i), b)
		if err != nil {
			return err
		}
	}
	return nil
}

func serializeUint128(v reflect.Value, b io.Writer) error {
	u := v.Interface().(big.Int)
	buf := u.Bytes()
	if len(buf) > 16 {
		return errors.New("big.Int too large for u128")
	}
	// fill big-endian buffer
	var d [16]byte
	copy(d[16-len(buf):], buf)
	// make it little-endian
	for i, j := 0, 15; i < j; i, j = i+1, j-1 {
		d[i], d[j] = d[j], d[i]
	}
	_, err := b.Write(d[:])
	return err
}

func serialize(v reflect.Value, b io.Writer) error {
	var err error
	switch v.Kind() {
	case reflect.Bool:
		if v.Bool() {
			_, err = b.Write([]byte{1})
		} else {
			_, err = b.Write([]byte{0})
		}
	case reflect.Int8:
		_, err = b.Write([]byte{byte((v.Int()))})
	case reflect.Int16:
		tmp := make([]byte, 2)
		binary.LittleEndian.PutUint16(tmp, uint16(v.Int()))
		_, err = b.Write(tmp)
	case reflect.Int32:
		tmp := make([]byte, 4)
		binary.LittleEndian.PutUint32(tmp, uint32(v.Int()))
		_, err = b.Write(tmp)
	case reflect.Int64:
		tmp := make([]byte, 8)
		binary.LittleEndian.PutUint64(tmp, uint64(v.Int()))
		_, err = b.Write(tmp)
	case reflect.Int:
		tmp := make([]byte, 8)
		binary.LittleEndian.PutUint64(tmp, uint64(v.Interface().(int)))
		_, err = b.Write(tmp)
	case reflect.Uint8:
		// user-defined Enum type is also uint8, so can't directly assert type here
		_, err = b.Write([]byte{byte(v.Uint())})
	case reflect.Uint16:
		tmp := make([]byte, 2)
		binary.LittleEndian.PutUint16(tmp, uint16(v.Uint()))
		_, err = b.Write(tmp)
	case reflect.Uint32:
		tmp := make([]byte, 4)
		binary.LittleEndian.PutUint32(tmp, uint32(v.Uint()))
		_, err = b.Write(tmp)
	case reflect.Uint64, reflect.Uint:
		tmp := make([]byte, 8)
		binary.LittleEndian.PutUint64(tmp, v.Uint())
		_, err = b.Write(tmp)
	case reflect.Float32:
		tmp := make([]byte, 4)
		f := v.Float()
		if f == math.NaN() {
			return errors.New("NaN float value")
		}
		binary.LittleEndian.PutUint32(tmp, math.Float32bits(float32(f)))
		_, err = b.Write(tmp)
	case reflect.Float64:
		tmp := make([]byte, 8)
		f := v.Float()
		if f == math.NaN() {
			return errors.New("NaN float value")
		}
		binary.LittleEndian.PutUint64(tmp, math.Float64bits(f))
		_, err = b.Write(tmp)
	case reflect.String:
		tmp := make([]byte, 4)
		binary.LittleEndian.PutUint32(tmp, uint32(len(v.String())))
		_, err = b.Write(tmp)
		if err != nil {
			break
		}
		_, err = b.Write([]byte(v.String()))
	case reflect.Array:
		for i := 0; i < v.Len(); i++ {
			err = serialize(v.Index(i), b)
			if err != nil {
				break
			}
		}
	case reflect.Slice:
		tmp := make([]byte, 4)
		binary.LittleEndian.PutUint32(tmp, uint32(v.Len()))
		_, err = b.Write(tmp)
		if err != nil {
			break
		}
		for i := 0; i < v.Len(); i++ {
			err = serialize(v.Index(i), b)
			if err != nil {
				break
			}
		}
	case reflect.Map:
		tmp := make([]byte, 4)
		binary.LittleEndian.PutUint32(tmp, uint32(v.Len()))
		_, err = b.Write(tmp)
		if err != nil {
			break
		}
		keys := v.MapKeys()
		sort.Slice(keys, vComp(keys))
		for _, k := range keys {
			err = serialize(k, b)
			if err != nil {
				break
			}
			err = serialize(v.MapIndex(k), b)
		}
	case reflect.Ptr:
		if v.IsNil() {
			_, err = b.Write([]byte{0})
		} else {
			_, err = b.Write([]byte{1})
			if err != nil {
				break
			}
			err = serialize(v.Elem(), b)
		}
	case reflect.Struct:
		if v.Type() == reflect.TypeOf(*big.NewInt(0)) {
			err = serializeUint128(v, b)
		} else {
			err = serializeStruct(v, b)
		}
	}
	return err
}

func vComp(keys []reflect.Value) func(int, int) bool {
	return func(i int, j int) bool {
		a, b := keys[i], keys[j]
		if a.Kind() == reflect.Interface {
			a = a.Elem()
			b = b.Elem()
		}
		switch a.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32:
			return a.Int() < b.Int()
		case reflect.Int64:
			return a.Interface().(int64) < b.Interface().(int64)
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32:
			return a.Uint() < b.Uint()
		case reflect.Uint64:
			return a.Interface().(uint64) < b.Interface().(uint64)
		case reflect.Float32, reflect.Float64:
			return a.Float() < b.Float()
		case reflect.String:
			return a.String() < b.String()
		}
		panic("unsupported key compare")
	}
}
