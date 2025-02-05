package tiltfile

import "go.starlark.net/starlark"

// Wrapper around starlark.AsString
func AsString(x starlark.Value) (string, bool) {
	b, ok := x.(*blob)
	if ok {
		return b.text, true
	}
	return starlark.AsString(x)
}

// Unpack an argument that can either be expressed as
// a string or as a list of strings.
func AsStringOrStringList(x starlark.Value) ([]string, bool) {
	if x == nil {
		return []string{}, true
	}

	s, ok := AsString(x)
	if ok {
		return []string{s}, true
	}

	iterable, ok := x.(starlark.Iterable)
	if ok {
		result := []string{}
		iter := iterable.Iterate()
		defer iter.Done()
		var item starlark.Value
		for iter.Next(&item) {
			s, ok := AsString(item)
			if !ok {
				return nil, false
			}
			result = append(result, s)
		}
		return result, true
	}

	return nil, false
}
