package formats

import (
	"reflect"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
)

func TestBASE64Marshaling(t *testing.T) {
	samples := []struct {
		input  interface{}
		result []byte
		err    error
	}{
		{
			NewStringer(`hello`),
			[]byte(`aGVsbG8=`),
			nil,
		},
	}

	base64 := NewBASE64()
	for k, sample := range samples {
		msg := spew.Sdump(k, sample)

		result, err := base64.Marshal(sample.input)
		assert.EqualValues(t, sample.err, err, msg)
		assert.EqualValues(t, sample.result, result, msg)

		v := reflect.New(reflect.TypeOf(sample.input)).Interface()
		err = base64.Unmarshal(result, v)
		assert.EqualValues(t, sample.err, err, msg)
		assert.EqualValues(
			t, sample.input.(*Stringer).String(),
			reflect.ValueOf(v).Elem().Interface().(*Stringer).String(),
			msg,
		)
	}
}
