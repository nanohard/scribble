// Package json contains a codec to encode and decode entities in JSON format
package json

import (
	"encoding/json"
)

const name = "json"

// Codec that encodes to and decodes from JSON.
var Codec = new(jsonCodec)

type jsonCodec int

func (c jsonCodec) Marshal(v interface{}) ([]byte, error) {
	// marshal, with tabs
	b, err := json.MarshalIndent(v, "", "\t")
	// add newline to the end
	b = append(b, byte('\n'))
	return b, err
}

func (c jsonCodec) Unmarshal(b []byte, v interface{}) error {
	return json.Unmarshal(b, v)
}

func (c jsonCodec) Name() string {
	return name
}
