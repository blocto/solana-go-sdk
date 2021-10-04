package borsh

import (
	"io"
	"reflect"
)

type Encoder struct {
	w io.Writer
}

func NewEncoder(w io.Writer) *Encoder {
	return &Encoder{w: w}
}

func (e *Encoder) Encode(s interface{}) error {
	return serialize(reflect.ValueOf(s), e.w)
}

func (e *Encoder) Close() error {
	return nil
}
