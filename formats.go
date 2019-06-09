package formats

import (
	"path/filepath"
	"strings"
)

const (
	// JSON is a JSON format name.
	JSON = "json"

	// YAML is a YAML format name.
	YAML = "yaml"

	// TOML is a TOML format name.
	TOML = "toml"

	// HEX is a HEX format name.
	HEX = "hex"

	// BASE64 is a BASE64 format name.
	BASE64 = "base64"

	// QUERY is a QUERY format name.
	QUERY = "query"

	// FLATE is a FLATE format name.
	FLATE = "flate"
)

var (
	// Names lists supported set of formats
	// which could be used as argument to Name() and
	// also available as constants exported from the package.
	Names = []string{
		JSON,
		YAML,
		TOML,
		HEX,
		BASE64,
		QUERY,
		FLATE,
	}

	Descriptions = map[string]string{
		JSON:   "javascript object notation is a lightweight data-interchange format",
		YAML:   "yaml ain't markup language, human-readable data-serialization language",
		TOML:   "tom's obvious, minimal language, a configuration file format",
		HEX:    "base 16 encoding",
		BASE64: "base 64 encoding",
		QUERY:  "percent encoded string",
		FLATE:  "deflate compressed data format",
	}

	// synonyms represents a format name synonyms mapping
	// into concrete format name.
	synonyms = map[string]string{
		"yml": YAML,
		"b64": BASE64,
		"q":   QUERY,
		"f": FLATE,
	}
)

// Format is a iterator that provides a data from source.
type Format interface {
	Marshal(v interface{}) ([]byte, error)
	Unmarshal(data []byte, v interface{}) error
	Name() string
}

// NewFromPath returns a new Format from file extension
// in the path argument.
// It respects the synonyms for format names,
// for example: yaml format files could have extensions
// yaml or yml.
func NewFromPath(path string) (Format, error) {
	name := strings.TrimPrefix(filepath.Ext(path), ".")
	if synonym, ok := synonyms[name]; ok {
		name = synonym
	}

	return New(strings.ToLower(name))
}

// New create a new format marshaler/unmarshaler from name.
func New(name string) (Format, error) {
	if name == "" {
		return nil, NewErrFormatNameIsEmpty()
	}

	if synonym, ok := synonyms[name]; ok {
		name = synonym
	}

	switch name {
	case JSON:
		return NewJSON(), nil
	case YAML:
		return NewYAML(), nil
	case TOML:
		return NewTOML(), nil
	case HEX:
		return NewHEX(), nil
	case BASE64:
		return NewBASE64(), nil
	case QUERY:
		return NewQUERY(), nil
	case FLATE:
		return NewFLATE(), nil
	default:
		return nil, NewErrNotSupported(name)
	}
}
