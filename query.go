package formats

import (
	"fmt"
	"net/url"

	"github.com/corpix/reflect"
)

// QUERYFormat is a QUERY marshaler.
type QUERYFormat struct{}

func (f *QUERYFormat) indirectToFstPtr(v interface{}) reflect.Value {
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
func (b *QUERYFormat) Marshal(v interface{}) ([]byte, error) {
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

	return []byte(url.QueryEscape(vs.String())), nil
}

// Unmarshal deserializes data represented by data byte slice into v.
func (b *QUERYFormat) Unmarshal(data []byte, v interface{}) error {
	rv := b.indirectToFstPtr(v)

	if rv.Kind() != reflect.Ptr {
		return reflect.NewErrWrongKind(reflect.Ptr, rv.Kind())
	}

	buf, err := url.QueryUnescape(string(TrimEmptySpace(data)))
	if err != nil {
		return err
	}

	if rv.Type() != stringerType {
		if rv.Elem().Kind() == reflect.Interface {
			rv.Elem().Set(reflect.ValueOf(*NewStringer(buf)))
			return nil
		}
		return reflect.NewErrCanNotAssertType(v, stringerType)
	}

	rv.Set(reflect.ValueOf(NewStringer(buf)))

	return nil
}

// Name returns a format name which is used to identify this format
// in the package.
func (b *QUERYFormat) Name() string {
	return QUERY
}

// NewQUERY constructs a new QUERY format marshaler.
func NewQUERY() *QUERYFormat { return &QUERYFormat{} }
