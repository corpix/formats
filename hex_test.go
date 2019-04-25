package formats

import (
	"reflect"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
)

func TestHEXMarshal(t *testing.T) {
	samples := []struct {
		input       interface{}
		result      []byte
		err         error
		inputResult bool
		resultInput bool
	}{
		{
			NewStringer("d34db33f"),
			[]byte("6433346462333366"),
			nil, true, true,
		},
		{
			NewStringer("holy moly!"),
			[]byte("686f6c79206d6f6c7921"),
			nil, true, true,
		},
		{
			NewStringer("holy moly!"),
			[]byte(" 686f6c 7920\n\n6d6f\t6c7921  "),
			nil, false, true,
		},
	}

	enc := NewHEX()
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
