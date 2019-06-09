package formats

import (
	"bytes"
	"compress/flate"
	"io/ioutil"

	"github.com/corpix/reflect"
)

// FLATEFormat is a FLATE marshaler.
type FLATEFormat struct{}

func (f *FLATEFormat) indirectToFstPtr(v interface{}) reflect.Value {
	var (
		rv = reflect.ValueOf(v)
	)
	for {
		if rv.Kind() == reflect.Ptr && rv.Elem().Kind() == reflect.Ptr {
			rv = rv.Elem()
			continue
		}
		break
	}
	return rv
}

// Marshal serializes data represented by v into slice of bytes.
func (b *FLATEFormat) Marshal(v interface{}) ([]byte, error) {
	vs, ok := v.(*Stringer)
	if !ok {
		return nil, reflect.NewErrCanNotAssertType(
			v,
			stringerType,
		)
	}

	buf := bytes.NewBuffer(nil)
	w, err := flate.NewWriter(buf, 9)
	if err != nil {
		return nil, err
	}
	defer w.Close()

	_, err = w.Write([]byte(vs.String()))
	if err != nil {
		return nil, err
	}

	err = w.Flush()
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// Unmarshal deserializes data represented by data byte slice into v.
func (b *FLATEFormat) Unmarshal(data []byte, v interface{}) error {
	rv := b.indirectToFstPtr(v)

	if rv.Kind() != reflect.Ptr {
		return reflect.NewErrWrongKind(reflect.Ptr, rv.Kind())
	}

	rc := flate.NewReader(bytes.NewReader(TrimEmptySpace(data)))
	defer rc.Close()

	buf, err := ioutil.ReadAll(rc)
	if err != nil {
		return err
	}

	if rv.Type() != stringerType {
		if rv.Elem().Kind() == reflect.Interface {
			rv.Elem().Set(reflect.ValueOf(*NewStringer(string(buf))))
			return nil
		}
		return reflect.NewErrCanNotAssertType(v, stringerType)
	}

	rv.Set(reflect.ValueOf(NewStringer(string(buf))))

	return nil
}

// Name returns a format name which is used to identify this format
// in the package.
func (b *FLATEFormat) Name() string {
	return FLATE
}

// NewFLATE constructs a new FLATE format marshaler.
func NewFLATE() *FLATEFormat { return &FLATEFormat{} }
