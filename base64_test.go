package formats

import (
	"reflect"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
)

func TestBASE64Marshaling(t *testing.T) {
	samples := []struct {
		input       interface{}
		result      []byte
		err         error
		inputResult bool
		resultInput bool
	}{
		{
			NewStringer("hello"),
			[]byte("aGVsbG8="),
			nil, true, true,
		},
		{
			NewStringer("hello"),
			[]byte(" aG V\tsb\nG8 = "),
			nil, false, true,
		},
	}

	enc := NewBASE64()
	for k, sample := range samples {
		msg := spew.Sdump(k, sample)

		if sample.inputResult {
			result, err := enc.Marshal(sample.input)
			assert.EqualValues(t, sample.err, err, msg)
			assert.EqualValues(t, sample.result, result, msg)
		}

		if sample.resultInput {
			v := reflect.New(reflect.TypeOf(sample.input)).Interface()
			err := enc.Unmarshal(sample.result, v)
			assert.EqualValues(t, sample.err, err, msg)
			assert.EqualValues(
				t, sample.input.(*Stringer).String(),
				reflect.ValueOf(v).Elem().Interface().(*Stringer).String(),
				msg,
			)
		}
	}
}
