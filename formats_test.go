package formats

// Copyright © 2017 Dmitry Moskowski
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
)

func TestNewFromPath(t *testing.T) {
	samples := []struct {
		path   string
		format Format
		err    error
	}{

		// json

		{
			"foo/bar.json",
			NewJSON(),
			nil,
		},
		{
			".json",
			NewJSON(),
			nil,
		},
		{
			"foo/barjson",
			nil,
			NewErrFormatNameIsEmpty(),
		},

		// yaml

		{
			"foo/bar.yaml",
			NewYAML(),
			nil,
		},
		{
			"foo/bar.yml",
			NewYAML(),
			nil,
		},
		{
			"foo/baryaml",
			nil,
			NewErrFormatNameIsEmpty(),
		},

		// other

		{
			"foo/bar.lol",
			nil,
			NewErrNotSupported("lol"),
		},
	}

	for k, sample := range samples {
		msg := spew.Sdump(k, sample)

		format, err := NewFromPath(sample.path)

		assert.EqualValues(t, sample.err, err, msg)
		assert.EqualValues(t, sample.format, format, msg)
	}
}

func TestNew(t *testing.T) {
	samples := []struct {
		name   string
		format Format
		err    error
	}{

		// json

		{
			"json",
			NewJSON(),
			nil,
		},
		{
			"JSON",
			nil,
			NewErrNotSupported("JSON"),
		},
		{
			"",
			nil,
			NewErrFormatNameIsEmpty(),
		},

		// yaml
		{
			"yaml",
			NewYAML(),
			nil,
		},
		{
			"yml",
			nil,
			NewErrNotSupported("yml"),
		},
		{
			"YAML",
			nil,
			NewErrNotSupported("YAML"),
		},
		{
			"",
			nil,
			NewErrFormatNameIsEmpty(),
		},
	}

	for k, sample := range samples {
		msg := spew.Sdump(k, sample)

		format, err := New(sample.name)

		assert.EqualValues(t, sample.err, err, msg)
		assert.EqualValues(t, sample.format, format, msg)
	}
}
