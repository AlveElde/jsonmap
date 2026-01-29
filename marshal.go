package jsonmap

import (
	"bytes"
	"encoding/json"
)

// MarshalJSON implements json.Marshaler interface.
// It marshals the map into JSON object.
//
//	data, err := json.Marshal(m)
func (m *Map) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	buf.WriteByte('{')
	enc := json.NewEncoder(&buf)
	enc.SetEscapeHTML(m.escapeHTML)
	first := true
	for el := m.First(); el != nil; el = el.Next() {
		if !first {
			buf.WriteByte(',')
		}
		first = false
		if err := enc.Encode(el.Key()); err != nil {
			return nil, err
		}
		buf.WriteByte(':')
		if err := enc.Encode(el.Value()); err != nil {
			return nil, err
		}
	}
	buf.WriteByte('}')
	return buf.Bytes(), nil
}
