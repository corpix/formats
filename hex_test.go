package formats

import (
	"reflect"
	"testing"

	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
)

func TestHEXMarshal(t *testing.T) {
	samples := []struct {
		input  interface{}
		result []byte
		err    error
	}{
		{
			input:  NewStringer("d34db33f"),
			result: []byte(`6433346462333366`),
			err:    nil,
		},
		{
			input:  NewStringer("holy moly!"),
			result: []byte(`686f6c79206d6f6c7921`),
			err:    nil,
		},
	}

	hex := NewHEX()
	for k, sample := range samples {
		msg := spew.Sdump(k, sample)

		result, err := hex.Marshal(sample.input)
		assert.EqualValues(t, sample.err, err, msg)
		assert.EqualValues(t, sample.result, result, msg)

		v := reflect.New(reflect.TypeOf(sample.input)).Interface()
		err = hex.Unmarshal(result, v)
		assert.EqualValues(t, sample.err, err, fmt.Sprintf("%s", err))
		assert.EqualValues(t, sample.input, reflect.ValueOf(v).Elem().Interface(), msg)
	}
}
