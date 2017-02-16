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
package formats

import (
	"errors"
	"fmt"

	"github.com/corpix/errcomposer"
)

var (
	// ErrNotSupported indicates that format with specified name is not supported.
	ErrNotSupported = errors.New("Format is not supported")
)

// NewErrNotSupported append a name of the source to the ErrNotSupported.
func NewErrNotSupported(format string) error {
	return errcomposer.Compose(
		ErrNotSupported,
		fmt.Errorf("Name: '%s'", format),
	)
}

// Format is a iterator that provides a data from source.
type Format interface {
	Marshal(v interface{}) ([]byte, error)
	Unmarshal(data []byte, v interface{}) error
}

// New create a new format marshaler/unmarshaler from name.
func New(name string) (Format, error) {
	switch name {
	case "yaml":
		return NewYAML(), nil
	case "json":
		return NewJSON(), nil
	default:
		return nil, NewErrNotSupported(name)
	}
}
