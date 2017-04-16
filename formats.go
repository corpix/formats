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
	"path/filepath"
	"strings"
)

const (
	// JSON is a JSON format name.
	JSON = "json"

	// YAML is a YAML format name.
	YAML = "yaml"
)

var (
	// Names lists supported set of formats
	// which could be used as argument to Name() and
	// also available as constants exported from the package.
	Names = []string{
		JSON,
		YAML,
	}

	// synonyms represents a format name synonyms mapping
	// into concrete format name.
	synonyms = map[string]string{
		"yml": YAML,
	}
)

// Format is a iterator that provides a data from source.
type Format interface {
	Marshal(v interface{}) ([]byte, error)
	Unmarshal(data []byte, v interface{}) error
}

// NewFromPath returns a new Format from file extension
// in the path argument.
// It respects the synonyms for format names,
// for example: yaml format files could have extensions
// yaml or yml.
func NewFromPath(path string) (Format, error) {
	name := strings.TrimPrefix(
		filepath.Ext(path),
		".",
	)
	if synonym, ok := synonyms[name]; ok {
		name = synonym
	}

	return New(name)
}

// New create a new format marshaler/unmarshaler from name.
func New(name string) (Format, error) {
	if name == "" {
		return nil, NewErrFormatNameIsEmpty()
	}

	switch name {
	case JSON:
		return NewJSON(), nil
	case YAML:
		return NewYAML(), nil
	default:
		return nil, NewErrNotSupported(name)
	}
}
