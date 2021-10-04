package borsh

import (
	"errors"
	"io"
	"reflect"
)

type Decoder struct {
	r io.Reader
}

func NewDecoder(r io.Reader) *Decoder {
	return &Decoder{r: r}
}

func (d *Decoder) Decode(s interface{}) error {
	t := reflect.TypeOf(s)
	if t.Kind() != reflect.Ptr {
		return errors.New("argument must be pointer")
	}
	val, err := deserialize(t, d.r)
	if err != nil {
		return nil
	}
	reflect.ValueOf(s).Elem().Set(reflect.ValueOf(val))
	return nil
}

func (d *Decoder) Close() error {
	return nil
}
