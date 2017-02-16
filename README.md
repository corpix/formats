formats
----------

This package provides a consistent API to transform same data(represented by some `struct`) between arbitrary formats.

Supported formats:

- `JSON`
- `YAML`

# Compatibility

There is a compatibility layer for:

- `JSON`, which helps to [mitigate](https://github.com/go-yaml/yaml/issues/139) non string keys in maps before they will be marshaled into `JSON`
