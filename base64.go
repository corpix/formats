package formats

import (
	"encoding/base64"
	"fmt"

	"github.com/corpix/reflect"
)

// BASE64Format is a BASE64 marshaler.
type BASE64Format struct{}

func (f *BASE64Format) indirectToFstPtr(v interface{}) reflect.Value {
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
func (b *BASE64Format) Marshal(v interface{}) ([]byte, error) {
	var (
		rv = b.indirectToFstPtr(v)
		vs fmt.Stringer
	)

	if rv.Type() != stringerType || !rv.IsValid() {
		return nil, reflect.NewErrCanNotAssertType(
			v,
			stringerType,
		)
	}

	vs = rv.Interface().(*Stringer)

	var (
		vv  = []byte(vs.String())
		buf = make([]byte, base64.StdEncoding.EncodedLen(len(vv)))
	)

	base64.StdEncoding.Encode(buf, vv)
	return buf, nil
}

// Unmarshal deserializes data represented by data byte slice into v.
func (b *BASE64Format) Unmarshal(data []byte, v interface{}) error {
	var (
		buf = make([]byte, base64.StdEncoding.DecodedLen(len(data)))
		rv  = b.indirectToFstPtr(v)
		n int
		err error
	)

	if rv.Kind() != reflect.Ptr {
		return reflect.NewErrWrongKind(reflect.Ptr, rv.Kind())
	}

	n, err = base64.StdEncoding.Decode(buf, TrimEmptySpace(data))
	if err != nil {
		return err
	}
	buf = buf[:n]

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
func (b *BASE64Format) Name() string {
	return BASE64
}

// NewBASE64 constructs a new BASE64 format marshaler.
func NewBASE64() *BASE64Format { return &BASE64Format{} }
