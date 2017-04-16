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
	"gopkg.in/yaml.v2"
)

// YAMLFormat is a YAML marshaler.
type YAMLFormat uint8

func (y *YAMLFormat) Marshal(v interface{}) ([]byte, error) {
	return yaml.Marshal(v)
}

func (y *YAMLFormat) Unmarshal(data []byte, v interface{}) error {
	return yaml.Unmarshal(data, v)
}

// Name returns a format name which is used to identify this format
// in the package.
func (y *YAMLFormat) Name() string {
	return YAML
}

// NewYAML constructs a new YAML format marshaler.
func NewYAML() *YAMLFormat { return new(YAMLFormat) }